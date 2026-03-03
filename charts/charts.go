// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package charts

import (
	"embed"
)

// Internal contains the internal charts
//
//go:embed internal
var Internal embed.FS

const (
	// ChartsPath is the path to the charts
	ChartsPath = "internal"

	// AgentSandboxChartPath is the path to the agent-sandbox chart
	AgentSandboxChartPath = "agent-sandbox"
)
