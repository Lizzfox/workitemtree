package main

import (
	"github.com/Lizzfox/workingitemtree/azure"
	"github.com/thatisuday/commando"
)

func main() {
	commando.SetExecutableName("tree").
		SetVersion("1.0.0").
		SetDescription("The program traverses the Parent-Child relationship until there are no more children \nand prints out Item Id and Title for each level")

	commando.
		Register(nil).
		AddArgument("organization", "the name of organization", "").
		AddArgument("token", "the personal access token", "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			orgName := args["organization"].Value
			token := args["token"].Value
			azure.GetWorkItems(orgName, token)
		})

		commando.Parse(nil)
}
