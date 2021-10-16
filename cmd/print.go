package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevegood/sunsets/pkg/print"
	"strings"
)

// cmdPrint represents the print command
var cmdPrint = &cobra.Command{
	Use:   "print [string to print]",
	Short: "Print anything to the screen",
	Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		print.Message(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(cmdPrint)
}
