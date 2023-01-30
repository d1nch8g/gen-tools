<p align="center">
<img style="align: center; padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="238px" height="238px" src="https://gitea.dancheg97.ru/repo-avatars/67-4297f15da3e76c29478ec89973007622" />
</p>

<h2 align="center">Golden go</h2>

[![Generic badge](https://img.shields.io/badge/LICENSE-GPLv3-red.svg)](https://gitea.dancheg97.ru/templates/golden-go/src/branch/main/LICENSE)
[![Generic badge](https://img.shields.io/badge/GITEA-REPO-orange.svg)](https://gitea.dancheg97.ru/templates/golden-go)
[![Generic badge](https://img.shields.io/badge/DOCKER-CONTAINER-blue.svg)](https://gitea.dancheg97.ru/templates/-/packages/container/golden-go/latest)
[![Build Status](https://drone.dancheg97.ru/api/badges/templates/golden-go/status.svg)](https://drone.dancheg97.ru/templates/golden-go)

Docker image with collection of formatters, linters and generators that can be
used in CI pipelines and for local development.

---

You can pull docker container via command:

```sh
docker pull gitea.dancheg97.ru/templates/golden-go:latest
```

---

Collection includes following tooling:

- [gofumpt](https://github.com/mvdan/gofumpt) - tool for formatting go code

```sh
docker run --rm -it -v ${pwd}:/wd -w /wd gitea.dancheg97.ru/templates/golden-go:latest gofumpt --help
```

- [golanglint-ci](https://golangci-lint.run/) - tool for linting go code

```sh
docker run --rm -it -v ${pwd}:/wd -w /wd gitea.dancheg97.ru/templates/golden-go:latest golanglint-ci --help
```

- [buf](https://docs.buf.build/introduction) - tool for helping with protocol buffers and gRPC

```sh
docker run --rm -it -v ${pwd}:/wd -w /wd gitea.dancheg97.ru/templates/golden-go:latest buf --help
```

- [sqlc](https://docs.sqlc.dev/en/stable) - tool for helping with protocol buffers and gRPC

```sh
docker run --rm -it -v ${pwd}:/wd -w /wd gitea.dancheg97.ru/templates/golden-go:latest sqlc --help
```
