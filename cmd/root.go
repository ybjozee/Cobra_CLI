package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cipher_cli",
	Short: "Encrypt and decrypt secret messages in seconds!!!",
	Long: `This application encrypts and decrypts secret messages with ease. 

Try out the Caesar and Bacon Cipher options to generate secret messages and share with your inner circle

cipher_cli encrypt "Welcome to the hallowed chambers" --algorithm=caesar --key=54

cipher_cli encrypt "Welcome to the hallowed chambers" --algorithm=bacon`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().
	//	StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cipher_cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("key", "k", "", "The key to pass for the algorithm")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find project directory.
		home, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in project directory with name ".cipher_cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cipher_cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
