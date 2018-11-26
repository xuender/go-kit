package kit

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CmdString 读取配置String
func CmdString(cmd *cobra.Command, name string) string {
	f := cmd.Flag(name)
	// 命令行优先
	if f.Changed {
		return f.Value.String()
	}
	ret := viper.GetString(name)
	if ret == "" {
		return f.Value.String()
	}
	return ret
}

// CmdBool 读取配置Bool
func CmdBool(cmd *cobra.Command, name string) bool {
	f := cmd.Flag(name)
	b, _ := strconv.ParseBool(f.Value.String())
	// 命令行优先
	if f.Changed {
		return b
	}
	ret := viper.GetString(name)
	if ret == "" {
		return b
	}
	return viper.GetBool(name)
}
