package cmd

import (
	"fmt"
	"os"

	"dvs_api/config"

	"dvs_api/server"

	"github.com/spf13/cobra"
)

var cfgFile, logFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "dvs_api",
	Short:   "dvs api",
	Long:    `dvs api is dvs backend.`,
	Version: "dvs/0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test")
		server.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.yml)")
	// rootCmd.PersistentFlags().StringVar(&logFile, "log", "", "log file (default is ./log/dvs.log)")
}

// init
func initConfig() {
	config.Init(cfgFile)
	// log.Init(logFile)
}
