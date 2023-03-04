//go:build rocksdb

package car

import (
	"github.com/spf13/cobra"
	"github.com/jumpsiegel/radiance/cmd/radiance/car/create"
	"github.com/jumpsiegel/radiance/cmd/radiance/car/dump"
)

var Cmd = cobra.Command{
	Use:   "car",
	Short: "Manage IPLD Content-addressable ARchives",
	Long:  "https://ipld.io/specs/transport/car/",
}

func init() {
	Cmd.AddCommand(
		&create.Cmd,
		&dump.Cmd,
	)
}
