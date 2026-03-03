// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

//go:generate sh -c "bash ${GARDENER_HACK_DIR}/generate-controller-registration.sh shoot-agent-sandbox . $(cat ../../VERSION) ../../example/controller-registration.yaml Extension:shoot-agent-sandbox"

// Package chart enables go:generate support for generating the correct controller registration.
package chart
