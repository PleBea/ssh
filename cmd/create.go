package cmd

import (
	"fmt"

	"github.com/plebea/ssh/libs/config"
	"github.com/plebea/ssh/libs/prompt"
	"github.com/spf13/cobra"
)

func Create() {
	host := prompt.GetInput(prompt.InputPrompt{
		Label: "What is the name of the connection?",
	})

	hostName := prompt.GetInput(prompt.InputPrompt{
		Label: "What is the host?",
	})

	user := prompt.GetInput(prompt.InputPrompt{
		Label: "What is the username?",
	})

	port := prompt.GetInput(prompt.InputPrompt{
		Label:   "What is the port?",
		Default: "22",
	})

	newConnection := config.Connection{
		Host:     host,
		HostName: hostName,
		User:     user,
		Port:     port,
	}

	config.Create(newConnection)

	fmt.Println("Connection created successfully")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new SSH connection",
	Run: func(cmd *cobra.Command, args []string) {
		Create()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
