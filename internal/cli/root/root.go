package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type AsciiArt string

const (
	cliBanner AsciiArt = `
 __    __     ______     ______     __  __     ______     __  __       
/\ "-./  \   /\  __ \   /\  ___\   /\ \/ /    /\  ___\   /\ \_\ \      
\ \ \-./\ \  \ \ \/\ \  \ \ \____  \ \  _"-.  \ \___  \  \ \____ \     
 \ \_\ \ \_\  \ \_____\  \ \_____\  \ \_\ \_\  \/\_____\  \/\_____\    
  \/_/  \/_/   \/_____/   \/_____/   \/_/\/_/   \/_____/   \/_____/ 

`
)

func initApp(cmd *cobra.Command, args []string) {
	fmt.Print(cliBanner)
	fmt.Println("Starting Mocksy")
	// Start the app
}

var rootCmd = &cobra.Command{
	Use:   "mocksy",
	Short: "Start Mocksy",
	Long:  "Starts the Mocksy server and initializes the app",
	Run:   initApp,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
