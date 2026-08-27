/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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
		dsn := "file:" + dbPath + "?mode=ro"

		var err error
		sqlite, err := sql.Open("sqlite", dsn)
		if err != nil {
			return fmt.Errorf("Failed to open DB: %w", err)
		}

		err = sqlite.Ping()
		if err != nil {
			return fmt.Errorf("Failed to connect to DB: %w", err)
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
	path = path + "/cbdb-tools/cbdb.sqlite3"
	rootCmd.PersistentFlags().StringVar(&dbPath, "db", path, "Specify the sqlite file path")
}
