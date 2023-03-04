//go:build rocksdb

package blockstore

import (
	"github.com/spf13/cobra"
	"github.com/jumpsiegel/radiance/cmd/radiance/blockstore/dumpshreds"
	"github.com/jumpsiegel/radiance/cmd/radiance/blockstore/verifydata"
	"github.com/jumpsiegel/radiance/cmd/radiance/blockstore/yaml"
)

var Cmd = cobra.Command{
	Use:   "blockstore",
	Short: "Access blockstore database",
}

func init() {
	Cmd.AddCommand(
		&dumpshreds.Cmd,
		&verifydata.Cmd,
		&yaml.Cmd,
	)
}
