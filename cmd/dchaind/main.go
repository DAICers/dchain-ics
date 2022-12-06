package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/DAICers/dchain-ics/app"
	"github.com/DAICers/dchain-ics/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		"dchain",
		"cosmos",
		app.DefaultNodeHome,
		"dchain",
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
