package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "application",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("use -h to show available commands")
	},
}

func Run() {
	rootCmd.AddCommand(nSqdCmd)
	nSqdCmd.Flags().String(FlagTypeNSqd, "", "to select nSqd types will run")

	rootCmd.AddCommand(gRpcCmd)

	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	migrateCmd.AddCommand(migrateFreshCmd)
	rootCmd.AddCommand(migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
