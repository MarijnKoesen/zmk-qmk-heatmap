/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zmk-heatmap",
	Short: "A heatmap generator for ZMK Keyboards",
	Long: `A heatmap generator for ZMK Keyboards.

Generating a heatmap is done in two steps:
1) Collect your keystrokes: this is the process of listening to all of your keystrokes and
   keeping of the amount of times each key is pressed and storing this in a file.
   Run the command as:
     $ zmk-heatmap collect -k /dev/tty.usbmodem142101
2) Generating the heatmap: taking all the keystrokes into account the heatmap can now be
   created:
     $ zmk-heatmap generate'`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(collectCmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zmk-heatmap.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
