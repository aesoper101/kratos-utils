package bootstrap

import (
	"github.com/spf13/cobra"
	"os"
	"sync"
)

type RootCommand struct {
	config      *Config
	rootCommand *cobra.Command
	runCommand  *cobra.Command
	once        sync.Once
}

func NewRootCommand(names ...string) *RootCommand {
	commandName := os.Args[0]
	if len(names) > 0 {
		commandName = names[0]
	}
	return &RootCommand{
		config: &Config{},
		rootCommand: &cobra.Command{
			Use: commandName,
		},
		runCommand: &cobra.Command{
			Use:   "run",
			Short: "run app",
			Long:  `run app`,
		},
	}
}

func (rc *RootCommand) Init() {
	rc.runCommand.Flags().StringVarP(&rc.config.Env, "env", "e", "dev", "runtime environment, eg: -env dev")
	rc.runCommand.Flags().StringVarP(&rc.config.ConfigPath, "config", "c", "../../configs", "config path, eg: --conf bootstrap.yaml")
	rc.runCommand.Flags().StringVarP(&rc.config.ConfigType, "type", "t", "", "config provider, eg: --type consul")
	rc.runCommand.Flags().StringSliceVarP(&rc.config.ConfigHost, "servers", "s", []string{}, "config server host, eg: --serers 127.0.0.1:8500")
	rc.runCommand.Flags().StringVarP(&rc.config.Username, "username", "u", "", "config server's username")
	rc.runCommand.Flags().StringVarP(&rc.config.Password, "password", "p", "", "config server's password")
}

func (rc *RootCommand) Execute() error {
	rc.once.Do(func() {
		rc.rootCommand.AddCommand(rc.runCommand)
	})
	return rc.rootCommand.Execute()
}

func (rc *RootCommand) AddCommand(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		if cmd.Use == "run" {
			panic("run command cant not be set.")
		}
		rc.rootCommand.AddCommand(cmd)
	}
}

func (rc *RootCommand) SetRunCommandAction(runner func(cfg Config) error) {
	rc.runCommand.RunE = func(cmd *cobra.Command, args []string) error {
		return runner(*rc.config)
	}
}
