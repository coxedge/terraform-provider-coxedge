/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package apiclient

import (
	"net/http"
	"testing"
)

var firstUserId string
var createdUserId string

func TestGetUsers(t *testing.T) {
	orgs, err := apiClient.GetUsers()
	if err != nil {
		t.Error(err)
	}
	if len(orgs) == 0 {
		t.Error("Got 0 users")
	}

	firstUserId = orgs[0].Id
	t.Logf("Got %d Users\n", len(orgs))
}

func TestGetUser(t *testing.T) {
	org, err := apiClient.GetUser(firstUserId)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Got user with ID: %s\n", org.Id)
}

func TestUserCreate(t *testing.T) {
	user := UserCreateRequest{
		UserName:       "testuser",
		FirstName:      "Test",
		LastName:       "User",
		Email:          "testuser@harpooncorp.io",
		OrganizationId: IdOnlyHelper{Id: firstOrgId},
	}

	newUser, err := apiClient.CreateUser(user)
	if err != nil {
		t.Error(err)
	}
	createdUserId = newUser.Id
	t.Logf("Created user with ID: %s\n", createdUserId)
}

func TestUserDelete(t *testing.T) {
	err := apiClient.DeleteUser(createdUserId)
	if err != nil {
		t.Error(err)
	}
}

//TestUnlockUser Unlock user in account by id
func (c *Client) TestUnlockUser(id string) error {
	//Create the request
	request, err := http.NewRequest("POST", CoxEdgeAPIBase+"/users/"+id+"/unlock", nil)
	if err != nil {
		return err
	}

	//Execute request
	_, err = c.doRequest(request)
	if err != nil {
		return err
	}
	return nil
}
