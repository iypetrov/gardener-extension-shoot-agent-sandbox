// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/ptr"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaults_AgentSandbox sets default values for AgentSandbox objects.
func SetDefaults_AgentSandbox(obj *AgentSandbox) {
	if obj == nil {
		return
	}

	// Default EnableExtensions to true
	if obj.EnableExtensions == nil {
		obj.EnableExtensions = ptr.To(true)
	}
}
