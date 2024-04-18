package cmd

import (
	"fmt"

	"github.com/plebea/ssh/libs/config"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all saved SSH connections",
	Run: func(cmd *cobra.Command, args []string) {
		connections := config.Get()

		for _, connection := range connections {
			fmt.Println(connection.Host + "(" + connection.User + "@" + connection.HostName + ")")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
