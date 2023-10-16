package cmd

import (
	"github.com/spf13/cobra"
	"vehicle-log-grpc-api/app/grpc"
)

var gRpcCmd = &cobra.Command{
	Use: "gRpc",
	Run: func(cmd *cobra.Command, args []string) {
		grpcServer := grpc.NewGrpc()
		grpcServer.Start()
	},
}
