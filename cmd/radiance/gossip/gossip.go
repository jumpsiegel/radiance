package gossip

import (
	"github.com/spf13/cobra"
	"github.com/jumpsiegel/radiance/cmd/radiance/gossip/ping"
	"github.com/jumpsiegel/radiance/cmd/radiance/gossip/pull"
)

var Cmd = cobra.Command{
	Use:   "gossip",
	Short: "Interact with Solana gossip networks",
}

func init() {
	Cmd.AddCommand(
		&ping.Cmd,
		&pull.Cmd,
	)
}
