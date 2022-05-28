package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "saidb",
	Short: "Simple and Insecure Database",
	Long: `                                                                                                    
                                                                                bbbbbbbb            
   SSSSSSSSSSSSSSS              AAA               IIIIIIIIIIDDDDDDDDDDDDD       b::::::b            
 SS:::::::::::::::S            A:::A              I::::::::ID::::::::::::DDD    b::::::b            
S:::::SSSSSS::::::S           A:::::A             I::::::::ID:::::::::::::::DD  b::::::b            
S:::::S     SSSSSSS          A:::::::A            II::::::IIDDD:::::DDDDD:::::D  b:::::b            
S:::::S                     A:::::::::A             I::::I    D:::::D    D:::::D b:::::bbbbbbbbb    
S:::::S                    A:::::A:::::A            I::::I    D:::::D     D:::::Db::::::::::::::bb  
 S::::SSSS                A:::::A A:::::A           I::::I    D:::::D     D:::::Db::::::::::::::::b 
  SS::::::SSSSS          A:::::A   A:::::A          I::::I    D:::::D     D:::::Db:::::bbbbb:::::::b
    SSS::::::::SS       A:::::A     A:::::A         I::::I    D:::::D     D:::::Db:::::b    b::::::b
       SSSSSS::::S     A:::::AAAAAAAAA:::::A        I::::I    D:::::D     D:::::Db:::::b     b:::::b
            S:::::S   A:::::::::::::::::::::A       I::::I    D:::::D     D:::::Db:::::b     b:::::b
            S:::::S  A:::::AAAAAAAAAAAAA:::::A      I::::I    D:::::D    D:::::D b:::::b     b:::::b
SSSSSSS     S:::::S A:::::A             A:::::A   II::::::IIDDD:::::DDDDD:::::D  b:::::bbbbbb::::::b
S::::::SSSSSS:::::SA:::::A               A:::::A  I::::::::ID:::::::::::::::DD   b::::::::::::::::b 
S:::::::::::::::SSA:::::A                 A:::::A I::::::::ID::::::::::::DDD     b:::::::::::::::b  
 SSSSSSSSSSSSSSS AAAAAAA                   AAAAAAAIIIIIIIIIIDDDDDDDDDDDDD        bbbbbbbbbbbbbbbb   

`,
	//Run: func(cmd *cobra.Command, args []string) {
	//commands := map[string]interface{}{}
	//
	//reader := bufio.NewScanner(os.Stdin)
	//fmt.Print("saidb > ")
	//
	//for reader.Scan() {
	//	text := strings.TrimSpace(reader.Text())
	//	text = strings.ToLower(text)
	//
	//	if command, exists := commands[text]; exists {
	//		// Call a hardcoded function
	//		command.(func())()
	//	} else if strings.EqualFold("exit", text) {
	//		return
	//	} else {
	//		fmt.Printf("command not found: %s \n", text)
	//	}
	//
	//	fmt.Print("saidb > ")
	//}
	//
	//fmt.Println()
	//},
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
