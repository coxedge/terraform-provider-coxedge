package apiclient

import "testing"

func TestGetOriginSettings(t *testing.T) {
	res, err := apiClient.GetOriginSettings("test-codecraft", "352cdc1e-c071-49ad-bddd-371094880507")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Got %v origin", res.Id)
}

func TestUpdateOriginSettings(t *testing.T) {

	org := OriginSettingsOrigin{
		Address: "cc.coxedge.com",
	}
	orga := OriginSettings{
		EnvironmentName: "test-codecraft",
		Origin:          org,
	}
	res, err := apiClient.UpdateOriginSettings("352cdc1e-c071-49ad-bddd-371094880507", orga)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Got %v origin", res.Id)
}
