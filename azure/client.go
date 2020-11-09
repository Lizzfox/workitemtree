package azure

import (
	"context"
	"fmt"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"

	"github.com/Lizzfox/workitemtree/tree"
)

// AzureURL is the Azure DevOps URL
const AzureURL = "https://dev.azure.com/"

// PrintWorkItems prints work items for organization  
func PrintWorkItems(orgName, PAT string) {
	orgURL := fmt.Sprintf("%s%s", AzureURL, orgName)
	connection := azuredevops.NewPatConnection(orgURL, PAT)
	ctx := context.Background()

	coreClient, err := workitemtracking.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	queryText := `
								select [System.Id], [System.Title]
								from WorkItemLinks
								where ([System.Links.LinkType] = 'System.LinkTypes.Hierarchy-Forward')
								order by [System.Id]
								mode (Recursive, ReturnMatchingChildren)
								`

	query := workitemtracking.QueryByWiqlArgs{
		Wiql: &workitemtracking.Wiql{
			Query: &queryText,
		},
	}

	resp, err := coreClient.QueryByWiql(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	if resp == nil {
		log.Print("Backlog is empty!")
		return
	}

	tree := tree.NewWorkItemTree(*(*resp).WorkItemRelations)

	queryArgs := workitemtracking.GetWorkItemsArgs{
		Ids: &tree.ItemIDs, 
		Fields: &[]string{"System.Id", "System.Title"},
	}

	workItems, err := coreClient.GetWorkItems(ctx, queryArgs)

	if err != nil {
		log.Fatal(err)
	}

	tree.MergeTitles(*workItems)
	tree.Show()
}
