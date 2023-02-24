package cmd

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/svrakitin/gossip-glomers/cmd/echo"
	"github.com/svrakitin/gossip-glomers/cmd/uniqueids"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "maelstrom-node",
	}
	return cmd
}

func Execute() error {
	ctx := context.TODO()

	signalCtx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	cmd := newRootCmd()
	cmd.AddCommand(echo.NewCmd())
	cmd.AddCommand(uniqueids.NewCmd())
	return cmd.ExecuteContext(signalCtx)
}
