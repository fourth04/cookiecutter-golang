package cmd

import (
        "fmt"
        "os"

		"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"
		"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/logger"
        "github.com/spf13/cobra"
)

var log = logger.Log()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
        Use:   "{{cookiecutter.app_name}}",
        Short: "A brief description of your application",
        Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
        // Uncomment the following line if your bare application
        // has an action associated with it:
        //      Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}

func initFlag() {
	rootCmd.PersistentFlags().StringP("configfile", "c", "", "config file")
	rootCmd.PersistentFlags().StringP("logfile", "f", "", "log file")
	rootCmd.PersistentFlags().StringP("loglevel", "l", "info", "log level")
	rootCmd.PersistentFlags().BoolP("json_logs", "j", false, "json logs")
}

func initConfig() {
	cfg := config.ConfigPtr()
	config.ReloadConfigFromFlagSet(cfg, rootCmd.PersistentFlags(), "configfile")
}

func initLog() {
	l := logger.LogPtr()
	cfg := config.Config()
	logger.ReloadLogrusLogger(l, cfg)
}

func init() {
	initFlag()
	cobra.OnInitialize(initConfig, initLog)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
