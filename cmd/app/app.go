package app

import "github.com/spf13/cobra"

var appCommand = &cobra.Command{
	Use:  "jvm",
	Long: "Simple jvm on golang!",
}

func Execute() error {
	return appCommand.Execute()
}

func init() {
	appCommand.AddCommand(versionCommand)
	appCommand.AddCommand(jvmCommand)
}
