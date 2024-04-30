package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"

	"github.com/toozej/sn2ssg-go/pkg/man"
	"github.com/toozej/sn2ssg-go/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:              "sn2ssg-go",
	Short:            "Simplenote to SSG Exporter",
	Long:             `A Golang application to export notes from Simplenote and format them for use in a variety of static-site-generators.`,
	PersistentPreRun: rootCmdPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.Config.ConfigVar)

	},
}

func rootCmdPreRun(cmd *cobra.Command, args []string) {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return
	}
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	// create rootCmd-level flags
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug-level logging")

	// add sub-commands
	rootCmd.AddCommand(
		man.NewManCmd(),
		version.Command(),
	)
}
