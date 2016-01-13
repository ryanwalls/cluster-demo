package commands

import (
	"github.com/ryanwalls/cluster-demo/services"
	"github.com/spf13/cobra"
	"os"
)

// ClusterCmd is the root command
var ClusterCmd = &cobra.Command{
	Use: "cluster-demo",
}

//Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	addCommands()
	if err := ClusterCmd.Execute(); err != nil {
		// the err is already logged by Cobra
		os.Exit(-1)
	}
}

func addCommands() {
	ClusterCmd.AddCommand(serviceACmd)
	ClusterCmd.AddCommand(serviceBCmd)
}

var serviceACmd = &cobra.Command{
	Use:   "service-a",
	Short: "Starts service A who delegates calls to service B",
	Long:  `Starts service A who delegates calls to service B`,
	Run: func(cmd *cobra.Command, args []string) {
		services.StartServiceA()
	},
}

var serviceBCmd = &cobra.Command{
	Use:   "service-b",
	Short: "Starts service B",
	Long:  `Starts service B`,
	Run: func(cmd *cobra.Command, args []string) {
		services.StartServiceB()
	},
}
