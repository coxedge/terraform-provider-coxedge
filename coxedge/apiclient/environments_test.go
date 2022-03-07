package apiclient

import (
	"fmt"
	"testing"
)

var firstEnvironmentId string
var createdEnvironmentId string
var serviceConnectionId string

func TestEnvironmentCreate(t *testing.T) {
	//Prep
	orgs, err := apiClient.GetOrganizations()
	if err != nil {
		t.Error(err)
	}
	if len(orgs) == 0 {
		t.Error("Got 0 organizations")
	}

	firstOrgId = orgs[0].Id
	for _, sc := range orgs[0].ServiceConnections {
		if sc.Name == "Edge Compute" {
			serviceConnectionId = sc.Id
		}
	}

	//Test
	environment := EnvironmentCreateRequest{
		EnvironmentName:   "test-env-for-tfrunner",
		Description:       "This was created by the Golang API Test",
		ServiceConnection: IdOnlyHelper{Id: serviceConnectionId},
		Organization:      IdOnlyHelper{Id: firstOrgId},
	}
	fmt.Println()

	newEnvironment, err := apiClient.CreateEnvironment(environment)
	if err != nil {
		t.Error(err)
	} else {
		createdEnvironmentId = newEnvironment.Id
		t.Logf("Created Environment with ID: %s\n", createdEnvironmentId)
	}
}

func TestGetEnvironments(t *testing.T) {
	orgs, err := apiClient.GetEnvironments()
	if err != nil {
		t.Error(err)
	}

	firstEnvironmentId = orgs[0].Id
	t.Logf("Got %d Environments\n", len(orgs))
}

func TestGetEnvironment(t *testing.T) {
	org, err := apiClient.GetEnvironment(firstEnvironmentId)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Got Environment with ID: %s\n", org.Id)
}

func TestEnvironmentDelete(t *testing.T) {
	err := apiClient.DeleteEnvironment(createdEnvironmentId)
	if err != nil {
		t.Error(err)
	}
}
