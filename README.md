<p align="center">
<img style="align: center; padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="238px" height="238px" src="https://dancheg97.ru/repo-avatars/67-4297f15da3e76c29478ec89973007622" />
</p>

<h2 align="center">Template generation tools</h2>

[![Generic badge](https://img.shields.io/badge/LICENSE-GPLv3-orange.svg)](https://dancheg97.ru/templates/gen-tools/src/branch/main/LICENSE)
[![Generic badge](https://img.shields.io/badge/GITEA-REPO-red.svg)](https://dancheg97.ru/templates/gen-tools)
[![Generic badge](https://img.shields.io/badge/GITHUB-REPO-white.svg)](https://github.com/dancheg97/gen-tools)
[![Generic badge](https://img.shields.io/badge/DOCKER-REGISTRY-blue.svg)](https://dancheg97.ru/templates/-/packages/container/gen-tools/latest)
[![Generic badge](https://img.shields.io/badge/AUR-REPO-cyan.svg)](https://aur.archlinux.org/packages/gen-tools)
[![Build Status](https://drone.dancheg97.ru/api/badges/templates/gen-tools/status.svg)](https://drone.dancheg97.ru/templates/gen-tools)

CLI Tool for generating different project modules for templating complex systems
with ease. Dockerized version also provides set of tools for code generation and
linting.

---

## Options:

Options you can specify under 'gen' command:

- **drone** - includes drone.yml template for project CI-CD
- **mit** - adds MIT license to project
- **gpl** - adds GPLv3 license to project
- **make** - adds Makefile to project
- **pkgbuild** - arch format PKGBUILD for packaging to AUR
- **lego** - generates templates of commands for lego, to simplify process of obtaining certificates

Compose file options:

- **compose-nginx** - adds nginx to your compose file, and generates template of config in directory (reuses certificates generated by lego)
- **compose-postgres** - adds postgres to compose file, and goose tool for migrations aswell
- **compose-redis** - adds redis to compose
- **compose-nats** - adds single nats node to compose, and nats-ui for visualization
- **compose-gitea** - adds gitea with some predefined parameters and configuartion files, additional themes, description, logo
- **compose-drone** - adds drone server and runner to compose, with predifined parameters to connect to gitea
- **compose-pacman** - adds self-hosted pacman repository for arch packages
- **compose-pocketbase** - adds pocketbase template to compose (self-hosted real-time backend)
- **compose-mkdocs** - adds mkdocs with theme, custom css and some preconfigured stuff
- **compose-kuma** - adds kuma, for tracking and visualizing services stability
- **compose-dozzle** - adds dozzle, for for viewing services logs

Go code options:

- **go-cli** - includes cobra and viper
- **go-lint** - includes golanglint-ci linter for go code
- **go-grpc** - includes proto and buf files for generation
- **go-docker** - includes 2 stage Dockerfile and compose for ease of development
- **go-pg** - includes pgx module in porstgres, sqlc for generation and goose for migrations
- **go-redis** - includes redis template
- **go-nats** - includes consumer and producer nats template

---

### Installation:

- docker

```
docker pull dancheg97.ru/templates/gen-tools:latest
```

- go

```
go install dancheg97.ru/templates/gen-tools@latest
```

- yay

```
yay -Sy gen-tools
```

---

### Examples:

```
gen-tools gen drone gpl
```

```
gen-tools go --repo myrepo.com/me/tool
```

```
gen-tools infr --name Nice --domain nice.org --user admin --pass SeCReT --email he@he.org
```

```
gen-tools infr --name Nice --domain nice.org --user admin --pass SeCReT --email he@he.org
```

- [gen-tools](README.md) - tool for generating go project templates

```sh
docker run --rm -it -v $(pwd):/wd -w /wd dancheg97.ru/templates/gen-tools:latest gen-tools --help
```

- [gofumpt](https://github.com/mvdan/gofumpt) - tool for formatting go code

```sh
docker run --rm -it -v $(pwd):/wd -w /wd dancheg97.ru/templates/gen-tools:latest gofumpt --help
```

- [golanglint-ci](https://golangci-lint.run/) - tool for linting go code, [config template](.golangci.yml)

```sh
docker run --rm -it -v $(pwd):/wd -w /wd dancheg97.ru/templates/gen-tools:latest golanglint-ci --help
```

- [buf](https://docs.buf.build/introduction) - tool for helping with protocol buffers and gRPC, [buf example](buf.yaml), [buf gen example](buf.gen.yaml)

```sh
docker run --rm -it -v $(pwd):/wd -w /wd dancheg97.ru/templates/gen-tools:latest buf --help
```

- [sqlc](https://docs.sqlc.dev/en/stable) - tool for generating type-safe go code from sql queries, [sqlc.sql example](sqlc.sql), config - [sqlc.yaml example](sqlc.yaml)

```sh
docker run --rm -it -v $(pwd):/wd -w /wd dancheg97.ru/templates/gen-tools:latest sqlc --help
```

- [go-swag](https://github.com/swaggo/swag) - tool for generating `swagger.yaml` from code annotations.

```sh
docker run --rm -it -v $(pwd):/wd -w /wd dancheg97.ru/templates/gen-tools:latest swag --help
```

- [go-lego](https://github.com/go-acme/lego) - tool for obtaining sertificates from ACME.

```sh
docker run --rm -it -v $(pwd):/wd -w /wd dancheg97.ru/templates/gen-tools:latest lego --help
```

- [go-lego](https://github.com/pressly/goose) - tool for migrating database (eg postgres).

```sh
docker run --rm -it -v $(pwd):/wd -w /wd dancheg97.ru/templates/gen-tools:latest goose --help
```
