// Copyright 2021-2025 The phy-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fake

import (
	"testing"
	"time"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

func TestDataStore_PrivateNetworks(t *testing.T) {
	ds := &Engine{
		PrivateNetworks: []*v1.PrivateNetwork{
			{
				PrivateNetworkId: "100000000001",
				ServerCount:      1,
				Service: v1.ServiceQuiet{
					Activated: time.Now(),
					Nickname:  "private-network01",
					ServiceId: "100000000001",
				},
				VlanId: 1,
				Zone: v1.Zone{
					Region: "is",
					ZoneId: 302,
				},
			},
			{
				PrivateNetworkId: "100000000002",
				ServerCount:      1,
				Service: v1.ServiceQuiet{
					Activated: time.Now(),
					Nickname:  "private-network02",
					ServiceId: "100000000002",
				},
				VlanId: 1,
				Zone: v1.Zone{
					Region: "is",
					ZoneId: 302,
				},
			},
		},
	}

	t.Run("select all", func(t *testing.T) {
		networks, err := ds.ListPrivateNetworks(v1.ListPrivateNetworksParams{})
		require.NoError(t, err)
		require.Equal(t, len(ds.PrivateNetworks), networks.Meta.Count)
		require.Len(t, networks.PrivateNetworks, len(ds.PrivateNetworks))
	})

	t.Run("select by id", func(t *testing.T) {
		network, err := ds.ReadPrivateNetwork("100000000002")
		require.NoError(t, err)
		require.Equal(t, "private-network02", network.Service.Nickname)
	})
}
