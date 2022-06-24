package apiclient

import "testing"

func TestCreateSite(t *testing.T) {
	newSite := SiteCreateRequest{
		EnvironmentName: "test-codecraft",
		Domain:          "www.cc.com",
		Hostname:        "199.250.204.212",
		Protocol:        "HTTPS",
		Services:        []string{"CDN", "SERVERLESS_EDGE_ENGINE", "WAF"},
	}
	res, err := apiClient.CreateSite(newSite)
	if err != nil {
		t.Error(err)
	}
	t.Logf("created %v", res.TaskStatus)
}
