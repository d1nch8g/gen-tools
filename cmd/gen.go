package cmd

import (
	"errors"
	"os"
	"strings"

	"dancheg97.ru/templates/gen-tools/templates"
	"dancheg97.ru/templates/gen-tools/templates/arch"
	"dancheg97.ru/templates/gen-tools/templates/devops"
	"dancheg97.ru/templates/gen-tools/templates/golang"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:     "gen",
	Short:   "📃 Generate template components",
	Run:     Gen,
	Example: "drone ",
}

func Gen(cmd *cobra.Command, args []string) {
	setLogFormat()

	for _, arg := range args {
		switch arg {

		// OVERALL
		case "drone":
			WriteFile(".drone.yml", templates.DroneYml)
		case "make":
			WriteFile("Makefile", templates.Makefile)
		case "gpl":
			WriteFile("LICENSE", templates.LicenseGPLv3)
		case "mit":
			WriteFile("LICENSE", templates.LicenseMIT)
		case "lego":
			WriteFile("lego.sh", templates.LegoSh)

		// DEVOPS
		case "compose-drone":
			AppendToCompose(devops.DroneYaml)
		case "compose-gitea":
			AppendToCompose(devops.GiteaYaml)
			WriteFile(`gitea/gitea/templates/home.tmpl`, devops.GiteaHomeTmpl)
			WriteFile(`gitea/gitea/templates/custom/body_outer_pre.tmpl`, devops.GiteaThemeParkTmpl)
			WriteFile(`gitea/gitea/public/css/theme-earl-grey.css`, devops.GiteaEarlGrayCss)
		case "compose-nginx":
			AppendToCompose(devops.NginxYaml)
			WriteFile(`nginx/nginx.conf`, devops.NginxConf)
		case "compose-pacman":
			AppendToCompose(devops.PacmanYaml)
		case "pocketbase-pacman":
			AppendToCompose(devops.PocketbaseYaml)

		// GOLANG
		case "go-cli":
			WriteFile("main.go", golang.CliMainGo)
			WriteFile("cmd/flags.go", golang.CliFlagsGo)
			WriteFile("cmd/run.go", golang.CliRunGo)
			WriteFile("cmd/root.go", golang.CliRootGo)
		case "go-lint":
			WriteFile(".golangci.yml", golang.GolangCiYml)
		case "go-grpc":
			WriteFile("buf.yaml", golang.BufYaml)
			WriteFile("buf.gen.yaml", golang.BufGenYaml)
		case "go-docker":
			WriteFile("Dockerfile", golang.Dockerfile)
			WriteFile("docker-compose.yml", golang.DockerCompose)
		case "go-pg":
			WriteFile("sqlc.yaml", golang.SqlcYaml)
			WriteFile("sqlc.sql", golang.SqlcSql)
			WriteFile("migrations/0001_ini.sql", golang.MigrationSql)
			WriteFile("postgres/postgres.go", golang.PostgresGo)
		case "go-redis":
			WriteFile("redis/redis.go", golang.RedisGo)
		case "go-nats":
			WriteFile("nats/consumer.go", golang.NatsConsumerGo)
			WriteFile("nats/producer.go", golang.NatsProducerGo)

		// ARCH
		case "pkgbuild":
			WriteFile("PKGBUILD", arch.Pkgbuild)
		// UNKNOWN
		default:
			logrus.Error("unknown arguement: ", arg)
		}
	}

	logrus.Info("template generation finished")
}

func WriteFile(file string, content string) {
	PrepareDir(file)
	err := os.WriteFile(file, []byte(content), 0o600)
	checkErr(err)
	logrus.Info("File generated: ", file)
}

func AppendToFile(file string, content string) {
	PrepareDir(file)
	f, err := os.Open(file)
	checkErr(err)
	_, err = f.Write([]byte(content))
	checkErr(err)
	logrus.Info("File modified: ", file)
}

func AppendToCompose(content string) {
	const compose = `docker-compose.yml`
	if _, err := os.Stat(compose); errors.Is(err, os.ErrNotExist) {
		WriteFile(compose, "services:\n")
	}
	AppendToFile(compose, content)
}

func PrepareDir(filePath string) {
	if len(strings.Split(filePath, `/`)) != 1 {
		splitted := strings.Split(filePath, `/`)
		path := strings.Join(splitted[0:len(splitted)-1], `/`)
		err := os.MkdirAll(path, os.ModePerm)
		checkErr(err)
	}
}
