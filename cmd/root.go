package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "saidb",
	Short: "Simple and Insecure Database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`
     ___      )¯¯,¯\ ° |¯¯¯| |\¯¯¯|\¯¯\  |\¯¯¯,¯)  
  _(   __\   /__/'\__\ |___| \/___|/__/| \/___'__)'
/____)__| |__ |/\|__|'|___| |___ |__'|/‘ |______|‘
|____|‘     '               ‘           °        ‘            

`)
		commands := map[string]interface{}{}

		reader := bufio.NewScanner(os.Stdin)
		fmt.Print("saidb > ")

		for reader.Scan() {
			text := strings.TrimSpace(reader.Text())
			text = strings.ToLower(text)

			if text != "" {
				if command, exists := commands[text]; exists {
					// Call a hardcoded function
					command.(func())()
				} else if strings.EqualFold("exit", text) {
					return
				} else {
					fmt.Printf("command not found: %s \n", text)
				}
			}

			fmt.Print("saidb > ")
		}

		fmt.Println()
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.saidb.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
