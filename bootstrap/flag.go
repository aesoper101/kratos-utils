package bootstrap

import (
	"github.com/spf13/cobra"
	"os"
)

type FlagCommand struct {
	config *ConfigProviderData
	*cobra.Command
}

func NewFlagCommand(names ...string) *FlagCommand {
	commandName := os.Args[0]
	if len(names) > 0 {
		commandName = names[0]
	}
	return &FlagCommand{
		config: &ConfigProviderData{},
		Command: &cobra.Command{
			Use: commandName,
		},
	}
}

func (fc *FlagCommand) Init() {
	fc.Command.Flags().StringVarP(&fc.config.Env, "env", "e", "dev", "runtime environment, eg: -env dev")
	fc.Command.Flags().StringVarP(&fc.config.ConfigPath, "config", "c", "../../configs", "config path, eg: --conf bootstrap.yaml")
	fc.Command.Flags().StringVarP(&fc.config.ConfigType, "type", "t", "", "config provider, eg: --type consul")
	fc.Command.Flags().StringSliceVarP(&fc.config.ConfigHost, "servers", "s", []string{}, "config server host, eg: --serers 127.0.0.1:8500")
	fc.Command.Flags().StringVarP(&fc.config.Username, "username", "u", "", "config server's username")
	fc.Command.Flags().StringVarP(&fc.config.Password, "password", "p", "", "config server's password")
}

func (fc *FlagCommand) GetConfig() *ConfigProviderData {
	return fc.config
}
