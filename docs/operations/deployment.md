# Agent Sandbox for Shoots

## Introduction

Save yourself the hassle of deploying an agent sandbox controller helm chart into your cluster and have Gardener do it by requesting this extension in the Shoot manifest.
To support this the Gardener must be installed with the `shoot-agent-sandbox` extension.

## Configuration

To generally enable the automatic deployment of agent sandbox for Shoots, the `shoot-agent-sandbox` extension must be registered by providing an appropriate [extension registration](../../example/controller-registration.yaml) in the garden cluster.

The extension must be separately enabled per shoot.

### ControllerRegistration

An example of a `ControllerRegistration` for the `shoot-agent-sandbox` can be found at [controller-registration.yaml](https://github.com/gardener/gardener-extension-shoot-agent-sandbox/blob/master/example/controller-registration.yaml).

The `ControllerRegistration` contains a Helm chart which eventually deploys the `shoot-agent-sandbox` to seed clusters

```yaml
apiVersion: core.gardener.cloud/v1beta1
kind: ControllerDeployment
...
  values:
```

### Enablement for a Shoot

To request the agent sandbox extension to work for a Shoot, it must be requested per shoot. To enable the service for a shoot, the shoot manifest must explicitly add the `shoot-agent-sandbox` extension.

```yaml
apiVersion: core.gardener.cloud/v1beta1
kind: Shoot
...
spec:
  extensions:
    - type: shoot-agent-sandbox
...
```

### Configuration of the Agent Sandbox in a Shoot

The agent sandbox deployment can be configured by providing a `providerConfig` to the extension section in the Shoot manifest:

```yaml
apiVersion: core.gardener.cloud/v1beta1
kind: Shoot
...
spec:
  extensions:
    - type: shoot-agent-sandbox
      providerConfig:
        apiVersion: agent-sandbox.extensions.gardener.cloud/v1alpha1
        kind: AgentSandbox
        enableExtensions: true
```