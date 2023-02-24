package uniqueids

import (
	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "unique-ids",
		RunE: func(cmd *cobra.Command, args []string) error {
			node := maelstrom.NewNode()
			node.Handle("generate", func(msg maelstrom.Message) error {
				body := map[string]any{
					"type": "generate_ok",
					"id":   uuid.NewString(),
				}

				return node.Reply(msg, body)
			})
			return node.Run()
		},
	}
	return cmd
}
