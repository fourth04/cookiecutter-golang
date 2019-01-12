package cmd

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/logger"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"
	"github.com/spf13/cobra"
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
	log.Info(cmd.Flags())
	cfg := config.ConfigPtr()
	config.ReloadConfigFromFlagSet(cfg, cmd.Flags(), "configfile")
	log.Info(cfg.AllSettings())
}
