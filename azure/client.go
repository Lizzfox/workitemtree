package azure

import (
	"context"
	"fmt"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"

	"github.com/Lizzfox/workingitemtree/tree"
)

const AzureURL = "https://dev.azure.com/"

func GetProjects(orgName, PAT string) {
	orgURL := fmt.Sprintf("%s%s", AzureURL, orgName)
	connection := azuredevops.NewPatConnection(orgURL, PAT)

	ctx := context.Background()

	// Create a client to interact with the Core area
	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	// Get first page of the list of team projects for your organization
	responseValue, err := coreClient.GetProjects(ctx, core.GetProjectsArgs{})
	if err != nil {
		log.Fatal(err)
	}

	index := 0
	for responseValue != nil {
		// Log the page of team project names
		for _, teamProjectReference := range (*responseValue).Value {
			log.Printf("Name[%v] = %v", index, *teamProjectReference.Name)
			index++
		}

		// if continuationToken has a value, then there is at least one more page of projects to get
		if responseValue.ContinuationToken != "" {
			// Get next page of team projects
			projectArgs := core.GetProjectsArgs{
				ContinuationToken: &responseValue.ContinuationToken,
			}
			responseValue, err = coreClient.GetProjects(ctx, projectArgs)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			responseValue = nil
		}
	}
}

func GetWorkItems(orgName, PAT string) {
	orgURL := fmt.Sprintf("%s%s", AzureURL, orgName)
	connection := azuredevops.NewPatConnection(orgURL, PAT)
	ctx := context.Background()

	coreClient, err := workitemtracking.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	queryText := `
								select [System.Id], [System.WorkItemType], [System.Title]
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

	workItems, err := coreClient.GetWorkItems(ctx, workitemtracking.GetWorkItemsArgs{Ids: &tree.ItemIDs})

	if err != nil {
		log.Fatal(err)
	}

	log.Print(workItems)

	tree.Show()
}
