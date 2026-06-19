/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/jesseck3013/cbdb-tool/internal/formatter"
	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/spf13/cobra"
)

// personCmd represents the person command
var personCmd = &cobra.Command{
	Use:   "person",
	Short: "Search profiles of person from CBDB",
	Long:  `Search profiles of person from CBDB`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			idStr := args[0]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Fatal("id should be number")

			}

			person, err := model.GetPersonByID(cmd.Context(), db, int64(id))

			switch err {
			case sql.ErrNoRows:
			case nil:
				formatter.PrintPerson(person)
			default:
				log.Fatalf("error: %v", err)
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(personCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// personCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// personCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
