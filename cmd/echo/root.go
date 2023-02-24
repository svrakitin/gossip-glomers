package echo

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "echo",
		RunE: func(cmd *cobra.Command, args []string) error {
			node := maelstrom.NewNode()
			node.Handle("echo", func(msg maelstrom.Message) error {
				var body map[string]any
				if err := json.Unmarshal(msg.Body, &body); err != nil {
					return err
				}

				body["type"] = "echo_ok"

				return node.Reply(msg, body)
			})
			return node.Run()
		},
	}
	return cmd
}
