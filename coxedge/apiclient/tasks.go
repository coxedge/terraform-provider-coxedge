/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package apiclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"net/http"
	"time"
)

const TaskSuccess = "SUCCESS"
const TaskFailed = "FAILURE"
const TaskPending = "PENDING"

func (c *Client) GetTaskStatus(taskId string) (*TaskStatus, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/tasks/"+taskId, nil)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct TaskStatus
	fmt.Println(string(respBytes))
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct, nil
}

func (c *Client) AwaitTaskResolve(upstreamCtx context.Context, taskId string, attemptCount int, interval time.Duration, timeout time.Duration) (*TaskStatus, error) {
	//Let's create a content
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	//Create channels for data
	errChan := make(chan error)
	dataChan := make(chan TaskStatus)

	go func() {
		tflog.Info(upstreamCtx, "Waiting on Resource to Create.")
		for i := 0; i < attemptCount; i++ {
			taskRes, err := c.GetTaskStatus(taskId)
			if err != nil {
				errChan <- err
				break
			}

			tflog.Info(upstreamCtx, "Status of task is "+taskRes.Data.TaskStatus)

			//Check to see if we are done
			if taskRes.Data.TaskStatus == TaskFailed {
				errChan <- errors.New("task failed")
				break
			} else if taskRes.Data.TaskStatus == TaskSuccess {
				dataChan <- *taskRes
				break
			} else {
				time.Sleep(interval)
			}

		}
	}()

	select {
	case err := <-errChan:
		return nil, err
	case res := <-dataChan:
		return &res, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c *Client) AwaitTaskResolveWithDefaults(ctx context.Context, taskId string) (*TaskStatus, error) {
	return c.AwaitTaskResolve(ctx, taskId, 100, 5*time.Second, 10*time.Minute)
}

func (c *Client) AwaitTaskResolveWithCustomTimeout(ctx context.Context, taskId string,timeout time.Duration) (*TaskStatus, error) {
	return c.AwaitTaskResolve(ctx, taskId, 100, 5*time.Second, timeout)
}
