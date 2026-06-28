/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jesseck3013/cbdb-tool/internal/controller"
	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/jesseck3013/cbdb-tool/internal/view"
	"github.com/spf13/cobra"
)

// personCmd represents the person command
var personCmd = &cobra.Command{
	Use:   "person",
	Short: "Search profiles of person from CBDB",
	Long:  `Search profiles of person from CBDB`,
	Run: func(cmd *cobra.Command, args []string) {
		isJSONOutput, err := cmd.Flags().GetBool("json")
		if err != nil {
			msg := fmt.Sprintf("Developer Error: flag json is not defined %v", err)
			panic(msg)
		}

		var c *controller.CLI
		switch isJSONOutput {
		case true:
			ps := controller.NewPersonService(db, view.JSONRenderer{})
			c = controller.NewCLI(ps)
		default:
			ps := controller.NewPersonService(db, view.TextRenderer{})
			c = controller.NewCLI(ps)
		}
		err = c.Person(cmd.Context(), args, cmd.Flags())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	personCmd.Flags().Bool("json", false, "output the query result in JSON format")
	personCmd.Flags().Bool(string(model.All), false, "select all fields")
	personCmd.Flags().Bool(string(model.AltName), false, "alternative names")
	personCmd.Flags().Bool(string(model.Association), false, "non-kinship association")
	personCmd.Flags().Bool(string(model.BasicInfo), true, "basic info")
	personCmd.Flags().Bool(string(model.Entry), false, "modes of entry")
	personCmd.Flags().Bool(string(model.Institution), false, "institution")
	personCmd.Flags().Bool(string(model.KinShip), false, "kinship")
	personCmd.Flags().Bool(string(model.Place), false, "biographical places")
	personCmd.Flags().Bool(string(model.Posting), false, "postings and offices")
	personCmd.Flags().Bool(string(model.Status), false, "social status")
	personCmd.Flags().Bool(string(model.Text), false, "text")

	rootCmd.AddCommand(personCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// personCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// personCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
