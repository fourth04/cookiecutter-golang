package cmd

import (
	"fmt"
	"os"
	{% if cookiecutter.use_logrus_logging == "y" %}"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/logger"{% endif %}
	{% if cookiecutter.use_viper_config == "y" %}"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"{% endif %}
	"github.com/spf13/cobra"
)

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
{% if cookiecutter.use_logrus_logging == "y" %}
var log = logger.Log()

func initFlag() {
	{% if cookiecutter.use_viper_config == "y" %}rootCmd.PersistentFlags().StringP("configfile", "c", "", "config file"){% endif %}
	rootCmd.PersistentFlags().StringP("logfile", "f", "", "log file")
	rootCmd.PersistentFlags().StringP("loglevel", "l", "info", "log level")
	rootCmd.PersistentFlags().BoolP("json_logs", "j", false, "json logs")
}
	{% if cookiecutter.use_viper_config == "y" %}
func initConfig() {
	cfg := config.ConfigPtr()
	config.ReloadConfigFromFlagSet(cfg, rootCmd.PersistentFlags(), "configfile")
}

func initLog() {
	l := logger.LogPtr()
	cfg := config.Config()
	logger.ReloadLogrusLoggerFromConfig(l, cfg)
}

func init() {
	initFlag()
	cobra.OnInitialize(initConfig, initLog)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
	{% else %}
func initLog() {
	l := logger.LogPtr()
	logger.ReloadLogrusLoggerFromFlagSet(l, rootCmd.PersistentFlags())
}

func init() {
	initFlag()
	cobra.OnInitialize(initLog)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
	{% endif %}
{% else %}
	{% if cookiecutter.use_viper_config == "y" %}
func initFlag() {
	rootCmd.PersistentFlags().StringP("configfile", "c", "", "config file")
}

func initConfig() {
	cfg := config.ConfigPtr()
	config.ReloadConfigFromFlagSet(cfg, rootCmd.PersistentFlags(), "configfile")
}

func init() {
	initFlag()
	cobra.OnInitialize(initConfig)
}
	{% endif %}
{% endif %}
