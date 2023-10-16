package cmd

import (
	"github.com/spf13/cobra"
	"vehicle-log-grpc-api/app/nsqd"
)

const (
	FlagTypeNSqd = "types"
)

var nSqdCmd = &cobra.Command{
	Use: "nSqd",
	Run: func(cmd *cobra.Command, args []string) {
		types, _ := cmd.Flags().GetString(FlagTypeNSqd)
		n := nsqd.New()
		n.Start(types)
	},
}
