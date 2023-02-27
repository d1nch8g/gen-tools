package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "📑 Get a list of all available tools",
	Run:   List,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func List(cmd *cobra.Command, args []string) {
	fmt.Print(list)
}

const list = `

🧰 Options you can specify under 'gen' command:

- drone - includes drone.yml template for project CI-CD
- pkgbuild - arch format PKGBUILD for packaging to AUR

📃 Available licenses:

- license-gpl - GNU General Public License is a series of widely used free software licenses that guarantee end users the four freedoms to run, study, share, and modify the software
- license-mit - permissive free software license originating at the Massachusetts Institute of Technology
- license-apache - permissive free software license written by the Apache Software Foundation

🐳 Compose file options:

- compose-nginx - adds nginx to your compose file, and generates template of config in directory (reuses certificates generated by lego)
- compose-postgres - adds postgres to compose file, and goose tool for migrations aswell
- compose-redis - adds redis to compose
- compose-nats - adds single nats node to compose, and nats-ui for visualization
- compose-gitea - adds gitea with some predefined parameters and configuartion files, additional themes, description, logo
- compose-drone - adds drone server and runner to compose, with predifined parameters to connect to gitea
- compose-pacman - adds self-hosted pacman repository for arch packages
- compose-pocketbase - adds pocketbase template to compose (self-hosted real-time backend)
- compose-mkdocs - adds mkdocs with theme, custom css and some preconfigured stuff
- compose-kuma - adds kuma, for tracking and visualizing services stability
- compose-dozzle - adds dozzle, for for viewing services logs

🐰 Go code options:

- go-cli - includes cobra and viper
- go-lint - includes golanglint-ci linter for go code
- go-grpc - includes proto and buf files for generation
- go-docker - includes 2 stage Dockerfile and compose for ease of development
- go-pg - includes pgx module in porstgres, sqlc for generation and goose for migrations
- go-redis - includes redis template
- go-nats - includes consumer and producer nats template

`
