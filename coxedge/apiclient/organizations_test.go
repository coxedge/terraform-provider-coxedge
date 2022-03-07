package apiclient

import (
	"os"
	"testing"
)

var apiKey string
var apiClient Client

var firstOrgId string
var firstOrgRoleId string

func TestMain(m *testing.M) {
	apiKey = os.Getenv("TEST_API_KEY")
	apiClient = NewClient(apiKey)
	code := m.Run()
	os.Exit(code)
}

func TestGetOrganizations(t *testing.T) {
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
	t.Logf("Got %d Organizations\n", len(orgs))
}

func TestGetOrganization(t *testing.T) {
	org, err := apiClient.GetOrganization(firstOrgId)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Got organization with ID: %s\n", org.Id)
}
