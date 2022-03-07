package apiclient

import "testing"

func TestClient_GetNetworkPolicyRules(t *testing.T) {
	items, err := apiClient.GetNetworkPolicyRules("base-test-env")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Got %d Items\n", len(items))
}
