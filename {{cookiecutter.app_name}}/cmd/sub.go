package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	{% if cookiecutter.use_viper_config == "y" %}"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"{% endif %}
)

// subCmd represents the parse command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "sub",
	Long:  `sub`,
	Run:   subCmdMain,
}

var input, output string

func initSubFlag() {
	subCmd.Flags().StringVarP(&input, "input", "i", "", "input filepath or directory")
	subCmd.Flags().StringVarP(&output, "output", "o", "", "output filepath")
}

func init() {
	initSubFlag()
	rootCmd.AddCommand(subCmd)
}

func subCmdMain(cmd *cobra.Command, args []string) {
	{% if cookiecutter.use_viper_config == "y" %}
	// binding sub command flags to cfg
	cfg := config.ConfigPtr()
	config.ReloadConfigFromFlagSet(cfg, cmd.Flags(), "")
	{% endif %}
	fmt.Println("Hello.")
}
