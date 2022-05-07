/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package apiclient

import (
	"context"
	"fmt"
	"testing"
)

var firstWorkloadId string

func TestGetWorkloads(t *testing.T) {
	items, err := apiClient.GetWorkloads("base-test-env")
	if err != nil {
		t.Error(err)
	}
	//if len(items) == 0 {
	//	t.Error("Got 0 items")
	//}

	if len(items) > 0 {
		firstWorkloadId = items[0].Id
	}
	t.Logf("Got %d Items\n", len(items))
}

func TestGetWorkload(t *testing.T) {
	if firstWorkloadId != "" {
		org, err := apiClient.GetWorkload("base-test-env", firstWorkloadId)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Got workload with ID: %s\n", org.Id)
	}
}

func TestCreateWorkload(t *testing.T) {
	newWorkload := WorkloadCreateRequest{
		EnvironmentName: "base-test-env",
		Name:            "test2",
		Deployments: []WorkloadAutoscaleDeployment{
			{
				Name:               "test",
				Pops:               []string{"MIA"},
				EnableAutoScaling:  false,
				InstancesPerPop:    1,
				MaxInstancesPerPop: 0,
				MinInstancesPerPop: 0,
				CPUUtilization:     0,
			},
		},
		Image: "ubuntu:latest",
		Specs: "SP-1",
		Type:  "CONTAINER",
	}
	taskResp, err := apiClient.CreateWorkload(newWorkload)
	if err != nil {
		t.Error(err)
	}
	taskResult, err := apiClient.AwaitTaskResolveWithDefaults(context.TODO(), taskResp.TaskId)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(taskResult)
}
