package apiclient

import (
	"context"
	"testing"
)

func TestUpdateWAFSettings(t *testing.T) {
	var c bool
	b := true
	new := WAFSettings{
		EnvironmentName: "test-codecraft",
		OwaspThreats: WAFOwaspThreats{SQLInjection: &b,
			XSSAttack: &c}}
	taskResp, err := apiClient.UpdateWAFSettings("352cdc1e-c071-49ad-bddd-371094880507", new)
	if err != nil {
		t.Error(err)
	}
	taskResult, err := apiClient.AwaitTaskResolveWithDefaults(context.TODO(), taskResp.TaskId)
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(taskResult)
	t.Logf("Got %v wafsettings", taskResult.Data)
}
