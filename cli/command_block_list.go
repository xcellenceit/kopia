package cli

import (
	"fmt"
	"sort"

	"github.com/kopia/kopia/block"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	blockListCommand = blockCommands.Command("list", "List blocks").Alias("ls")
	blockListLong    = blockListCommand.Flag("long", "Long output").Short('l').Bool()
	blockListPrefix  = blockListCommand.Flag("prefix", "Prefix").String()
	blockListSort    = blockListCommand.Flag("sort", "Sort order").Default("name").Enum("name", "size", "time", "none", "pack")
	blockListReverse = blockListCommand.Flag("reverse", "Reverse sort").Short('r').Bool()
	blockListSummary = blockListCommand.Flag("summary", "Summarize the list").Short('s').Bool()
)

func runListBlocksAction(context *kingpin.ParseContext) error {
	rep := mustOpenRepository(nil)
	defer rep.Close()

	blocks, err := rep.Blocks.ListBlocks(*blockListPrefix)
	if err != nil {
		return err
	}

	maybeReverse := func(b bool) bool { return b }

	if *blockListReverse {
		maybeReverse = func(b bool) bool { return !b }
	}

	switch *blockListSort {
	case "name":
		sort.Slice(blocks, func(i, j int) bool { return maybeReverse(blocks[i].BlockID < blocks[j].BlockID) })
	case "size":
		sort.Slice(blocks, func(i, j int) bool { return maybeReverse(blocks[i].Length < blocks[j].Length) })
	case "time":
		sort.Slice(blocks, func(i, j int) bool { return maybeReverse(blocks[i].Timestamp.Before(blocks[j].Timestamp)) })
	case "pack":
		sort.Slice(blocks, func(i, j int) bool { return maybeReverse(comparePacks(blocks[i], blocks[j])) })
	}

	var count int
	var totalSize int64
	uniquePacks := map[string]bool{}
	for _, b := range blocks {
		totalSize += b.Length
		count++
		if b.PackBlockID != "" {
			uniquePacks[b.PackBlockID] = true
		}
		if *blockListLong {
			if b.PackBlockID != "" {
				fmt.Printf("%-34v %10v %v in %v offset %v\n", b.BlockID, b.Length, b.Timestamp.Local().Format(timeFormat), b.PackBlockID, b.PackOffset)
			} else {
				fmt.Printf("%-34v %10v %v\n", b.BlockID, b.Length, b.Timestamp.Local().Format(timeFormat))
			}
		} else {
			fmt.Printf("%v\n", b.BlockID)
		}
	}

	if *blockListSummary {
		fmt.Printf("Total: %v blocks, %v packs, %v bytes\n", count, len(uniquePacks), totalSize)
	}

	return nil
}

func comparePacks(a, b block.Info) bool {
	if a, b := a.PackBlockID, b.PackBlockID; a != b {
		return a < b
	}

	return a.PackOffset < b.PackOffset
}

func init() {
	blockListCommand.Action(runListBlocksAction)
}
