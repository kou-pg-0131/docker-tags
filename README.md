[![CircleCI](https://circleci.com/gh/kou-pg-0131/docker-tags/tree/main.svg?style=shield)](https://circleci.com/gh/kou-pg-0131/docker-tags/tree/main)
[![Maintainability](https://api.codeclimate.com/v1/badges/009887ace72c96268bb1/maintainability)](https://codeclimate.com/github/kou-pg-0131/docker-tags/maintainability)
[![codecov](https://codecov.io/gh/kou-pg-0131/docker-tags/branch/main/graph/badge.svg?token=CUunJTFrLa)](https://codecov.io/gh/kou-pg-0131/docker-tags)
[![LICENSE](https://img.shields.io/github/license/kou-pg-0131/docker-tags?style=plastic)](./LICENSE)
[![Twitter Follow](https://img.shields.io/twitter/follow/kou_pg_0131?style=social)](https://twitter.com/kou_pg_0131)

# Overview

A CLI tool that displays a list of tags for a specific Docker image.

# Installation

```bash
$ go get -u github.com/kou-pg-0131/docker-tags
```

# Usage

```bash
$ docker-tags <IMAGE>
```

**Example**

```bash
$ docker-tags node
node:0
node:0-onbuild
node:0-slim
node:0-wheezy
node:0.10
node:0.10-onbuild
node:0.10-slim
node:0.10-wheezy
node:0.10.28
...
```

# LICENSE

[MIT](./LICENSE)
