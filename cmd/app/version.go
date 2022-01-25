package app

import (
	"fmt"
	jvmongo "github.com/rodbate/jvm-on-go"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:  "version",
	Long: "show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: ", jvmongo.Version)
	},
}
