// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AgentSandbox contains the configuration for the agent-sandbox controller.
type AgentSandbox struct {
	metav1.TypeMeta `json:",inline"`

	// EnableExtensions enables extension CRDs (SandboxClaim, SandboxTemplate, SandboxWarmPool) and their RBAC
	// +optional
	EnableExtensions *bool `json:"enableExtensions,omitempty"`
}
