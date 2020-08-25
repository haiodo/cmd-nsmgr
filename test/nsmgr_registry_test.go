// Copyright (c) 2020 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"net"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/peer"

	"github.com/networkservicemesh/api/pkg/api/registry"
)

func (f *NsmgrTestSuite) TestNSMgrEndpointRegister() {
	t := f.T()
	setup := newSetup(t)
	setup.Start()
	defer setup.Stop()

	regCtx := peer.NewContext(setup.ctx,
		&peer.Peer{
			Addr: &net.UnixAddr{
				Name: "/var/run/nse-1.sock",
				Net:  "unix",
			},
			AuthInfo: nil,
		})

	regClient := setup.NewRegistryClient(regCtx)

	regResponse, err := regClient.Register(regCtx, &registry.NetworkServiceEndpoint{
		Name: "my-nse",
	})
	require.Nil(t, err)
	require.NotNil(t, regResponse)
	require.Equal(t, setup.configuration.ListenOn[0].String(), regResponse.Url)
}
