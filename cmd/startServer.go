/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fasttrack_api/api"

	"github.com/spf13/cobra"
)

var port string

// startServerCmd represents the startServer command
var startServerCmd = &cobra.Command{
	Use:   "startServer",
	Short: "This command will start the api server",
	Long:  `This command will start the api server.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	DisableFlagsInUseLine: true,
	Example: `start server:
	fasttrack_api startServer or
	fasttrack_api startServer -p 8080 or 
	fasttrack_api startServer --port=8080`,

	Run: func(cmd *cobra.Command, args []string) {
		api.SetPortFlag(port)
		api.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
	startServerCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "This flag sets the port of the server")
}
