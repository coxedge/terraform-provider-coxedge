package apiclient

import "testing"

func TestUpdateCDNSettings(t *testing.T) {
	b := true

	cdn := CDNSettings{
		EnvironmentName:     "test-codecraft",
		SiteId:              "0e91e079-9f01-4a31-b71e-a7e8fb12ee3a",
		Http2SupportEnabled: &b}
	res, err := apiClient.UpdateCDNSettings("0e91e079-9f01-4a31-b71e-a7e8fb12ee3a", cdn)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Got %v cdn", res.TaskId)
}
