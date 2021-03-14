package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "service-generator",
	Short: "Service generator is used to generate a new Golang project structure",
	Long: `Service generator is a CLI library which allows you to create
a new Golang project with a predefined structure that follows
DDD and clean architecture principles. It's intended to be used in
micro-service architecture, but it's not tied to it.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
