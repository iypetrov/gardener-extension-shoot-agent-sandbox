// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package agentsandbox

import (
	"fmt"
	"path/filepath"

	"github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/extensions/pkg/util"
	"github.com/gardener/gardener/pkg/chartrenderer"

	"github.com/gardener/gardener-extension-shoot-agent-sandbox/charts"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/apis/operator"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/constants"
)

var (
	agentSandboxChartPath = filepath.Join(charts.ChartsPath, charts.AgentSandboxChartPath)
)

// RenderAgentSandboxChart renders the agent-sandbox chart with the provided configuration.
func RenderAgentSandboxChart(cluster *controller.Cluster, config *operator.AgentSandbox) (*chartrenderer.RenderedChart, error) {
	renderer, err := util.NewChartRendererForShoot(cluster.Shoot.Spec.Kubernetes.Version)
	if err != nil {
		return nil, fmt.Errorf("could not create chart renderer: %w", err)
	}

	values := map[string]any{
		"enableExtensions": true,
	}

	// Override enableExtensions if specified in config
	if config != nil && config.EnableExtensions != nil {
		values["enableExtensions"] = *config.EnableExtensions
	}

	renderedChart, err := renderer.RenderEmbeddedFS(charts.Internal, agentSandboxChartPath, constants.ReleaseAgentSandbox, constants.NamespaceAgentSandbox, values)
	if err != nil {
		return nil, fmt.Errorf("could not render agent-sandbox chart: %w", err)
	}

	return renderedChart, nil
}
