# [Gardener Extension for Agent Sandbox](https://gardener.cloud)

[![REUSE status](https://api.reuse.software/badge/github.com/gardener/gardener-extension-shoot-agent-sandbox)](https://api.reuse.software/info/github.com/gardener/gardener-extension-shoot-agent-sandbox)
[![Build](https://github.com/gardener/gardener-extension-shoot-agent-sandbox/actions/workflows/non-release.yaml/badge.svg)](https://github.com/gardener/gardener-extension-shoot-agent-sandbox/actions/workflows/non-release.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gardener/gardener-extension-shoot-agent-sandbox)](https://goreportcard.com/report/github.com/gardener/gardener-extension-shoot-agent-sandbox)

⚠️ This extension is still in alpha state. 

Project Gardener implements the automated management and operation of [Kubernetes](https://kubernetes.io/) clusters as a service. Its main principle is to leverage Kubernetes concepts for all of its tasks.

This repository contains a Gardener extension to automatically deploy agent sandbox controller into Shoot clusters. It implements Gardener's extension contract for the `shoot-agent-sandbox` extension. If you want to know what a Gardener extension is, have a look at [GEP-1](https://github.com/gardener/gardener/blob/master/docs/proposals/01-extensibility.md).

An example for a `ControllerRegistration` resource that can be used to register this controller to Gardener can be found [here](example/controller-registration.yaml).

## Extension Resources

Currently there is nothing to specify in the extension spec.

Example extension resource:

```yaml
apiVersion: extensions.gardener.cloud/v1alpha1
kind: Extension
metadata:
  name: extension-shoot-agent-sandbox
  namespace: shoot--project--abc
spec:
```

Please note, this extension controller relies on the [Gardener-Resource-Manager](https://github.com/gardener/gardener/blob/master/docs/concepts/resource-manager.md) to deploy K8S resources to Seed and Shoot clusters.

## How to start using or developing this extension controller locally

You can run the controller locally on your machine by executing `make start`.

We are using Go modules for Golang package dependency management and [Ginkgo](https://github.com/onsi/ginkgo)/[Gomega](https://github.com/onsi/gomega) for testing.

## Feedback and Support

Feedback and contributions are always welcome!

Please report bugs or suggestions as [GitHub issues](https://github.com/gardener/gardener-extension-shoot-agent-sandbox/issues) or reach out on [Slack](https://gardener-cloud.slack.com/) (join the workspace [here](https://gardener.cloud/community/community-bio/)).

## Learn more!

Please find further resources about out project here:

* [Our landing page gardener.cloud](https://gardener.cloud/)
* ["Gardener, the Kubernetes Botanist" blog on kubernetes.io](https://kubernetes.io/blog/2018/05/17/gardener/)
* ["Gardener Project Update" blog on kubernetes.io](https://kubernetes.io/blog/2019/12/02/gardener-project-update/)
* [GEP-1 (Gardener Enhancement Proposal) on extensibility](https://github.com/gardener/gardener/blob/master/docs/proposals/01-extensibility.md)
* [Extensibility API documentation](https://github.com/gardener/gardener/tree/master/docs/extensions)
* [Gardener Extensions Golang library](https://godoc.org/github.com/gardener/gardener/extensions/pkg)
* [Gardener API Reference](https://github.com/gardener/gardener/blob/master/docs/api-reference/README.md)
