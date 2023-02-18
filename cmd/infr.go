package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var infraCmd = &cobra.Command{
	Use:   "infr",
	Short: "🐋 Generate all infrastructure in single command",
	Run:   Infr,
}

func init() {
	rootCmd.AddCommand(infraCmd)
}

func Infr(cmd *cobra.Command, args []string) {
	PreventDefault(`mail`, viper.GetString(`mail`), `mail@example.com`)
	PreventDefault(`domain`, viper.GetString(`domain`), `example.com`)
	PreventDefault(`user`, viper.GetString(`user`), `admin`)
	PreventDefault(`pass`, viper.GetString(`pass`), `password`)

	Gen(cmd, []string{
		"compose-nginx",
		"compose-gitea",
		"compose-drone",
		"compose-mkdocs",
		"compose-kuma",
		"compose-dozzle",
	})
	logrus.Info("to obtain certificates run: sh certs.sh")
	logrus.Info("to run infrastructure run: docker compose up")
}

func PreventDefault(name string, actual string, initial string) {
	logrus.Warn(name, actual, initial)
	if actual == initial {
		logrus.Error(name + " was not applied, value is: " + actual)
		os.Exit(1)
	}
}