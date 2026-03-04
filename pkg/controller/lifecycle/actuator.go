// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package lifecycle

import (
	"context"
	_ "embed"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/extensions/pkg/controller/extension"
	"github.com/gardener/gardener/extensions/pkg/util"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/chartrenderer"
	"github.com/gardener/gardener/pkg/extensions"
	"github.com/gardener/gardener/pkg/utils/managedresources"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/gardener/gardener-extension-shoot-agent-sandbox/charts"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/agentsandbox"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/apis/config"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/apis/operator"
	api "github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/apis/operator"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/constants"
)

// NewActuator returns an actuator responsible for Extension resources.
func NewActuator(mgr manager.Manager, config config.Configuration) extension.Actuator {
	return &actuator{
		client:        mgr.GetClient(),
		config:        mgr.GetConfig(),
		decoder:       serializer.NewCodecFactory(mgr.GetScheme(), serializer.EnableStrict).UniversalDecoder(),
		serviceConfig: config,
	}
}

type actuator struct {
	client        client.Client
	config        *rest.Config
	decoder       runtime.Decoder
	serviceConfig config.Configuration
}

// Reconcile the Extension resource.
func (a *actuator) Reconcile(ctx context.Context, log logr.Logger, ex *extensionsv1alpha1.Extension) error {
	namespace := ex.GetNamespace()

	cluster, err := controller.GetCluster(ctx, a.client, namespace)
	if err != nil {
		return fmt.Errorf("unable to fetch cluster resource: %w", err)
	}

	if controller.IsHibernated(cluster) {
		return nil
	}

	config, err := decodeAgentSandboxConfig(a.decoder, ex)
	if err != nil {
		return fmt.Errorf("unable to extract/decode agent-sandbox config: %w", err)
	}

	return a.reconcile(ctx, cluster, namespace, config)
}

func (a *actuator) reconcile(ctx context.Context, cluster *extensions.Cluster, namespace string, config *operator.AgentSandbox) error {
	shootResources, err := a.getShootResources(cluster, config)
	if err != nil {
		return err
	}

	if err := managedresources.CreateForShoot(ctx, a.client, namespace, constants.ManagedResourceNamesAgentSandbox, constants.ExtensionName, false, shootResources); err != nil {
		return err
	}

	return nil
}

func decodeAgentSandboxConfig(decoder runtime.Decoder, ex *extensionsv1alpha1.Extension) (*operator.AgentSandbox, error) {
	config := &api.AgentSandbox{}
	if ex.Spec.ProviderConfig != nil {
		_, _, err := decoder.Decode(ex.Spec.ProviderConfig.Raw, nil, config)
		if err != nil {
			return nil, fmt.Errorf("failed to decode provider config: %w", err)
		}
	}

	return config, nil
}

func (a *actuator) getShootResources(cluster *controller.Cluster, config *operator.AgentSandbox) (map[string][]byte, error) {
	renderedChart, err := agentsandbox.RenderAgentSandboxChart(cluster, config)
	if err != nil {
		return nil, err
	}

	data := map[string][]byte{
		"agent-sandbox.yaml": renderedChart.Manifest(),
	}

	return data, nil
}

// Delete the Extension resource.
func (a *actuator) Delete(ctx context.Context, log logr.Logger, ex *extensionsv1alpha1.Extension) error {
	namespace := ex.GetNamespace()

	err := a.deleteShootResources(ctx, log, namespace, false)
	if err != nil {
		return err
	}

	return nil
}

// ForceDelete the Extension resource.
func (a *actuator) ForceDelete(ctx context.Context, log logr.Logger, ex *extensionsv1alpha1.Extension) error {
	namespace := ex.GetNamespace()

	err := a.deleteShootResources(ctx, log, namespace, true)
	if err != nil {
		return err
	}

	return nil
}

func (a *actuator) deleteShootResources(ctx context.Context, log logr.Logger, namespace string, forceDelete bool) error {
	log.Info("Deleting managed resource for shoot", "namespace", namespace)
	if err := managedresources.DeleteForShoot(ctx, a.client, namespace, constants.ManagedResourceNamesAgentSandbox); err != nil {
		return err
	}

	// We don't need to wait for the shoot managed resource deletion because managed resources are finalized by gardenlet
	// in later step in the Shoot force deletion flow.
	if forceDelete {
		return nil
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	if err := managedresources.WaitUntilDeleted(timeoutCtx, a.client, namespace, constants.ManagedResourceNamesAgentSandbox); err != nil {
		return err
	}

	return nil
}

// Restore the Extension resource.
func (a *actuator) Restore(ctx context.Context, log logr.Logger, ex *extensionsv1alpha1.Extension) error {
	return a.Reconcile(ctx, log, ex)
}

// Migrate the Extension resource.
func (a *actuator) Migrate(ctx context.Context, log logr.Logger, ex *extensionsv1alpha1.Extension) error {
	// Keep objects for shoot managed resources so that they are not deleted from the shoot during the migration
	if err := managedresources.SetKeepObjects(ctx, a.client, ex.GetNamespace(), constants.ManagedResourceNamesAgentSandbox, true); err != nil {
		return err
	}

	return a.Delete(ctx, log, ex)
}

var (
	agentSandboxChartPath    = filepath.Join(charts.ChartsPath, charts.AgentSandboxChartPath)
	shootComponentsChartPath = filepath.Join(charts.ChartsPath, charts.ShootComponentsChartPath)
)

// RenderAgentSandboxChart renders the agent-sandbox chart with the provided configuration.
func RenderAgentSandboxChart(cluster *controller.Cluster, values any) (*chartrenderer.RenderedChart, error) {
	renderer, err := util.NewChartRendererForShoot(cluster.Shoot.Spec.Kubernetes.Version)
	if err != nil {
		return nil, fmt.Errorf("could not create chart renderer: %w", err)
	}

	renderedChart, err := renderer.RenderEmbeddedFS(charts.Internal, agentSandboxChartPath, constants.ReleaseAgentSandbox, constants.NamespaceAgentSandbox, values)
	if err != nil {
		return nil, fmt.Errorf("could not render agent-sandbox chart: %w", err)
	}

	return renderedChart, nil
}
