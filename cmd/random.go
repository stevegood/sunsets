package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/stevegood/sunsets/pkg/sunset"
)

// cmdRandom represents the random command
var cmdRandom = &cobra.Command{
	Use:   "random",
	Short: "Generate random sunset SVG",
	Long:  `sunset generates and outputs a random sunset SVG.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		sunset.Random(os.Stdout, 1280, 720)
	},
}

func init() {
	rootCmd.AddCommand(cmdRandom)
}
