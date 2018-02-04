package cli

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	blockRepackCommand       = blockCommands.Command("repack", "Repackage small blocks into bigger ones")
	blockRepackSizeThreshold = blockRepackCommand.Flag("max-size", "Max size of block to re-pack").Default("500000").Uint64()
)

func runBlockRepackAction(context *kingpin.ParseContext) error {
	rep := mustOpenRepository(nil)
	defer rep.Close()

	if err := rep.Blocks.Repackage(*blockRepackSizeThreshold); err != nil {
		return err
	}

	return nil
}

func init() {
	blockRepackCommand.Action(runBlockRepackAction)
}
