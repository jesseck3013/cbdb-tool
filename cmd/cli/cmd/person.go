/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/spf13/cobra"
)

// personCmd represents the person command
var personCmd = &cobra.Command{
	Use:   "person",
	Short: "Search profiles of person from CBDB",
	Long:  `Search profiles of person from CBDB`,
	Run: func(cmd *cobra.Command, args []string) {
		err := c.Person(cmd.Context(), args, cmd.Flags())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	personCmd.Flags().Bool(string(model.AltName), false, "select basic info")
	personCmd.Flags().Bool(string(model.Association), false, "select basic info")
	personCmd.Flags().Bool(string(model.BasicInfo), true, "select basic info")
	personCmd.Flags().Bool(string(model.Entry), false, "select basic info")
	personCmd.Flags().Bool(string(model.Institution), false, "select basic info")
	personCmd.Flags().Bool(string(model.KinShip), false, "select basic info")
	personCmd.Flags().Bool(string(model.Place), false, "select basic info")
	personCmd.Flags().Bool(string(model.Posting), false, "select basic info")
	personCmd.Flags().Bool(string(model.Status), false, "select basic info")
	personCmd.Flags().Bool(string(model.Text), false, "select basic info")

	rootCmd.AddCommand(personCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// personCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// personCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
