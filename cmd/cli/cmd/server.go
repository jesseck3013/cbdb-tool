/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/jesseck3013/cbdb-tool/internal/controller"
	"github.com/jesseck3013/cbdb-tool/internal/view"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start the web API server",
	Long:  `start the web API server`,
	Run: func(cmd *cobra.Command, args []string) {
		c := controller.NewWEB(db, view.JSONRenderer{})
		http.HandleFunc("/person/{id}", c.PersonByID)
		log.Fatal(http.ListenAndServe(":8080", nil))
	},
}

func init() {
	serverCmd.Flags().String("port", "8080", "specify a tcp port: --port=8080")

	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
