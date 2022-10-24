package bootstrap

import (
	"github.com/spf13/pflag"
)

var defaultBootstrapFlags = &ConfigFlags{
	Env:        "dev",
	ConfigPath: "../../configs",
}

func RegisterFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&defaultBootstrapFlags.Env, "env", "e", "dev", "runtime environment, eg: -env dev")
	flags.StringVarP(&defaultBootstrapFlags.ConfigPath, "config", "c", "../../configs", "config path, eg: --conf bootstrap.yaml")
	flags.StringVarP(&defaultBootstrapFlags.ConfigType, "type", "t", "", "config provider, eg: --type consul")
	flags.StringSliceVarP(&defaultBootstrapFlags.ConfigHost, "servers", "s", []string{}, "config server host, eg: --serers 127.0.0.1:8500")
	flags.StringVarP(&defaultBootstrapFlags.Username, "username", "u", "", "config server's username")
	flags.StringVarP(&defaultBootstrapFlags.Password, "password", "p", "", "config server's password")
}

func GetFlags() *ConfigFlags {
	return defaultBootstrapFlags
}
