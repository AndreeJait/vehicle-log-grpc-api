package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"vehicle-log-grpc-api/app/migrate"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("use -h to show available commands")
	},
}

var migrateUpCmd = &cobra.Command{
	Use: "up",
	Run: func(cmd *cobra.Command, args []string) {
		startMigrate(migrate.MigrationTypeUp)
	},
}

var migrateDownCmd = &cobra.Command{
	Use: "down",
	Run: func(cmd *cobra.Command, args []string) {
		startMigrate(migrate.MigrationTypeDown)
	},
}

var migrateFreshCmd = &cobra.Command{
	Use: "fresh",
	Run: func(cmd *cobra.Command, args []string) {
		startMigrate(migrate.MigrationTypeFresh)
	},
}

func startMigrate(migrateType string) {
	m := migrate.New()
	m.Start(migrateType)
}
