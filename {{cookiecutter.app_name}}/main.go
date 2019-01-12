package main

import (
	{% if cookiecutter.use_cobra_cmd == "y" %}
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/cmd"
	{% else %}
	"fmt"
	"os"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version"
	{% if ((cookiecutter.use_viper_config == "y") and
			(cookiecutter.use_logrus_logging == "y")) %}
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/logger"
	"github.com/spf13/pflag"
	{% else %}
	"flag"
	{% endif %}
	{% endif %}
)

{% if cookiecutter.use_cobra_cmd == "y" %}
func main() {

    cmd.Execute()

}
{% else %}
{% if ((cookiecutter.use_viper_config == "y") and
			(cookiecutter.use_logrus_logging == "y")) %}
func initFlag() {
	pflag.StringP("configfile", "c", "", "config file")
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

func initConfig() {
	cfg := config.ConfigPtr()
	config.ReloadConfigFromFlagSet(cfg, pflag.CommandLine, "configfile")
}

func initLog() {
	l := logger.LogPtr()
	cfg := config.Config()
	logger.ReloadLogrusLogger(l, cfg)
}

func init() {
	initFlag()
	initConfig()
	initLog()
}

func main() {

	fmt.Println("Hello.")

}
{% else %}
func main() {

	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

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
