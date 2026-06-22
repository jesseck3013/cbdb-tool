/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/jesseck3013/cbdb-tool/internal/controller"
	"github.com/jesseck3013/cbdb-tool/internal/repository"
	"github.com/spf13/cobra"
)

var db *repository.Queries
var c *controller.CLI

var rootCmd = &cobra.Command{
	Use:   "cbdb",
	Short: "A query tool for CBDB",
	Long: `A query tool for the sqlite databse of China Biographical
Database Project`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(cli *controller.CLI) {
	c = cli
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
}
