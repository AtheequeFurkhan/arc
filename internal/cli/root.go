package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "arc",
	Short: "Universal cloud migration and automation tool",
	Long: `Arc is a powerful CLI tool for cloud migrations and automation.

It supports multiple cloud providers including Google Drive, AWS S3, 
Dropbox, OneDrive, and more. Perform bulk operations, migrations, 
and analytics across all your cloud storage.

Examples:
  arc auth login google          # Authenticate with Google
  arc ls google:/                # List files in Google Drive
  arc stats google:/             # Show file statistics
  arc cp google:/file s3://bucket/file  # Copy between providers`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.arc/config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")

	// Bind flags to viper
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Search config in home directory with name ".arc/config"
		viper.AddConfigPath(home + "/.arc")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	// Read in environment variables that match ARC_*
	viper.SetEnvPrefix("ARC")
	viper.AutomaticEnv()

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
}
