package main

import (
	"github.com/Lizzfox/workitemtree/azure"
	"github.com/thatisuday/commando"
)

func main() {
	commando.SetExecutableName("workitemtree").
		SetVersion("1.0.0").
		SetDescription("The program traverses the Parent-Child relationship until there are no more children \nand prints out Item Id and Title for each level")

	commando.
		Register(nil).
		AddArgument("organization", "the name of you Azure DevOps organization", "").
		AddArgument("token", "the personal access token", "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			orgName := args["organization"].Value
			token := args["token"].Value
			azure.PrintWorkItems(orgName, token)
		})

		commando.Parse(nil)
}
