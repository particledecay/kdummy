<p align="center">
  <h1 align="center">kdummy</h1>
  <p align="center">The greatest do nothing app.</p>
  <p align="center">
    <a href="https://github.com/particledecay/kdummy/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/particledecay/kdummy"></a>
    <a href="https://github.com/particledecay/kdummy/actions/workflows/release.yml"><img alt="GitHub Workflow Status" src="https://github.com/particledecay/kdummy/actions/workflows/release.yml/badge.svg"></a>
    <a href="https://codeclimate.com/github/particledecay/kdummy/maintainability"><img src="https://api.codeclimate.com/v1/badges/884a35e038544e847098/maintainability" /></a>
  </p>
</p>



## Description

kdummy is mostly useless as an application, but does several things that make it seem like a real app:

- emits log messages at regular intervals
- responds to health checks
- records metrics
- allows you to configure the logging interval live

The default ports it uses:

- 8080 for the `/heart/{rate}` configurable endpoint
- 9090 for the internal endpoints (`/metrics` and `/healthz`)

## Usage

The easiest way to use kdummy is via the Helm chart:

```sh
helm install kdummy chart
```

If you need to modify something, you may provide your own `values.yaml` file:

```sh
helm install kdummy chart -f my-values.yaml
```

You can also run kdummy on its own to launch the web server locally:

```sh
kdummy
```

## Why? Why not use an app that actually does something?

Because sometimes you want to just have a go-to application you want to launch into a k8s cluster when you're doing things like testing log shipping, metrics collection, autoscaling, ingress configs, DNS, Helm chart development, etc.

Sure, you can install something else, but this app won't interfere with whatever you're doing, and you can be sure it will be extremely fast and lightweight and stay out of your way so you can work.

## Known Issues

Check out the [Issues](https://github.com/particledecay/kdummy/issues) section or specifically [issues created by me](https://github.com/particledecay/kdummy/issues?q=is:issue+is:open+sort:updated-desc+author:particledecay)
