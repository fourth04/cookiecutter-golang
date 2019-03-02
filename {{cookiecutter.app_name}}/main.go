package main

import (
	{% if cookiecutter.use_cobra_cmd == "y" %}
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/cmd"
	{% else %}
	"fmt"
	"os"
	"github.com/spf13/pflag"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version"
	{% if cookiecutter.use_logrus_logging == "y" %}"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/logger"{% endif %}
	{% if cookiecutter.use_viper_config == "y" %}"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"{% endif %}
	{% endif %}
)

{% if cookiecutter.use_cobra_cmd == "y" %}
func main() {

    cmd.Execute()

}
{% else %}
	{% if cookiecutter.use_logrus_logging == "y" %}
func initFlag() {
	{% if cookiecutter.use_viper_config == "y" %}pflag.StringP("configfile", "c", "", "config file"){% endif %}
	pflag.StringP("logfile", "f", "", "log file")
	pflag.StringP("loglevel", "l", "info", "log level")
	pflag.BoolP("json_logs", "j", false, "json logs")
	versionFlag := pflag.BoolP("version", "v", false, "Version")
	pflag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("Version:", version.Version)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("OS / Arch:", version.OsArch)
		os.Exit(0)
	}
}
		{% if cookiecutter.use_viper_config == "y" %}
func initConfig() {
	cfg := config.ConfigPtr()
	config.ReloadConfigFromFlagSet(cfg, pflag.CommandLine, "configfile")
}

func initLog() {
	l := logger.LogPtr()
	cfg := config.Config()
	logger.ReloadLogrusLoggerFromConfig(l, cfg)
}

func init() {
	initFlag()
	initConfig()
	initLog()
}
		{% else %}
func initLog() {
	l := logger.LogPtr()
	logger.ReloadLogrusLoggerFromFlagSet(l, pflag.CommandLine)
}

func init() {
	initFlag()
	initLog()
}
		{% endif %}
func main() {

	fmt.Println("Hello.")

}
	{% else %}
		{% if cookiecutter.use_viper_config == "y" %}
func initFlag() {
	pflag.StringP("configfile", "c", "", "config file")
	versionFlag := pflag.BoolP("version", "v", false, "Version")
	pflag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("Version:", version.Version)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("OS / Arch:", version.OsArch)
		os.Exit(0)
	}
}

func initConfig() {
	cfg := config.ConfigPtr()
	config.ReloadConfigFromFlagSet(cfg, pflag.CommandLine, "configfile")
}

func init() {
	initFlag()
	initConfig()
}

func main() {

	fmt.Println("Hello.")

}
		{% else %}
func main() {

	versionFlag := pflag.BoolP("version", "v", false, "Version")
	pflag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
        fmt.Println("Git Commit:", version.GitCommit)
        fmt.Println("Version:", version.Version)
        fmt.Println("Go Version:", version.GoVersion)
        fmt.Println("OS / Arch:", version.OsArch)
		os.Exit(1)
	}
	fmt.Println("Hello.")
}
		{% endif %}
	{% endif %}
{% endif %}
