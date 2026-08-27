/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
		ps := controller.NewPersonService(db, view.JSONRenderer{})
		c := controller.NewWEB(ps)
		http.HandleFunc("/person/{id}", c.PersonByID)
		http.HandleFunc("/person/", c.PersonByName)

		port, err := cmd.Flags().GetString("port")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("API server starts listening to port :%v", port)
		addr := fmt.Sprintf(":%v", port)
		log.Fatal(http.ListenAndServe(addr, nil))
	},
}

func init() {
	serverCmd.Flags().String("port", "8080", "specify an http port: --port=8080")
	rootCmd.AddCommand(serverCmd)
}
