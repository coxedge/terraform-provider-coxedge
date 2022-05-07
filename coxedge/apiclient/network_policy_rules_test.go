/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package apiclient

import "testing"

func TestClient_GetNetworkPolicyRules(t *testing.T) {
	items, err := apiClient.GetNetworkPolicyRules("base-test-env")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Got %d Items\n", len(items))
}
