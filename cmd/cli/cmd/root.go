/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/jesseck3013/cbdb-tool/internal/query"
	"github.com/jesseck3013/cbdb-tool/internal/repository"
	"github.com/spf13/cobra"
)

var (
	dbPath string
	db     *repository.Queries
	sqlite *sql.DB
)

var rootCmd = &cobra.Command{
	Use:   "cbdb",
	Short: "A query tool for CBDB",
	Long: `A query tool for the sqlite databse of China Biographical
Database Project`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Annotations["skip-db"] == "true" {
			return nil
		}

		err := query.BuildIndex(dbPath)
		if err != nil {
			return err
		}

		sqlite, err := query.OpenDB(dbPath, true)
		if err != nil {
			return err
		}

		db = repository.New(sqlite)
		return nil
	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if sqlite != nil {
			sqlite.Close()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	path, err := os.UserCacheDir()
	if err != nil {
		log.Println(err)
	}
	path = filepath.Join(path, "/cbdb-tool/cbdb.sqlite3")
	rootCmd.PersistentFlags().StringVar(&dbPath, "db", path, "Specify the sqlite file path")
}
