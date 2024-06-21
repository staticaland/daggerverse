# Daggerverse

[Dagger documentation - Quickstart](https://docs.dagger.io/quickstart/daggerize)

## How do I create a new module?

```sh
dagger init --sdk=go --name=renovate renovate
```

## How do I run a module?

Check out the `Makefile` in each module for examples.

## How do I upgrade Dagger?

```sh
brew install dagger/tap/dagger
```

## How do I safely pass secrets to a module?

```sh
?
```

## How do I run a module with GitHub Actions?

Use the [`dagger/dagger-for-github`](https://github.com/dagger/dagger-for-github) action. It's just a composite action.

```yaml
name: Call a Dagger Function

on:
  push:
    branches: [main]

jobs:

  build:

    name: Call a Dagger Function
    runs-on: ubuntu-latest

    steps:

      - name: Checkout
        uses: actions/checkout@692973e3d937129bcbf40652eb9f2f61becf3332 # v4.1.7


      - name: Call Dagger Function
        uses: dagger/dagger-for-github@6210aa04aaf743e52a7315449e213bb85cd828ce # v5.10.0
        with:
          version: "latest"
          module: github.com/staticaland/daggerverse/cloc@main
          verb: call
          args: cloc-dir --directory-arg='.' --yaml-output
```
