# Register Shoot Agent Sandbox Extension in Shoot Clusters

## Introduction

Within a Shoot cluster, it is possible to enable the agent sandbox. It is necessary that the Gardener installation your Shoot cluster runs in is equipped with a `shoot-agent-sandbox` extension. Please ask your Gardener operator if the extension is available in your environment.

## Shoot Feature Gate

As the `shoot-agent-sandbox` extension is not intended to be enabled globally, it must be configured per Shoot cluster. Please adapt the Shoot specification by the configuration shown below to activate the extension individually.

```yaml
kind: Shoot
...
spec:
  extensions:
    - type: shoot-agent-sandbox
...
```