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

package phy

import (
	"context"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

type ServerAPI interface {
	List(ctx context.Context, params *v1.ListServersParams) (*v1.Servers, error)
	Read(ctx context.Context, serverId v1.ServerId) (*v1.Server, error)
	ListOSImages(ctx context.Context, serverId v1.ServerId) ([]*v1.OsImage, error)
	OSInstall(ctx context.Context, serverId v1.ServerId, params v1.OsInstallParameter) error
	ReadPortChannel(ctx context.Context, serverId v1.ServerId, portChannelId v1.PortChannelId) (*v1.PortChannel, error)
	ConfigureBonding(ctx context.Context, serverId v1.ServerId, portChannelId v1.PortChannelId, params v1.ConfigureBondingParameter) (*v1.PortChannel, error)
	ReadPort(ctx context.Context, serverId v1.ServerId, portId v1.PortId) (*v1.InterfacePort, error)
	UpdatePort(ctx context.Context, serverId v1.ServerId, portId v1.PortId, params v1.UpdateServerPortParameter) (*v1.InterfacePort, error)
	EnablePort(ctx context.Context, serverId v1.ServerId, portId v1.PortId, enable bool) (*v1.InterfacePort, error)
	AssignNetwork(ctx context.Context, serverId v1.ServerId, portId v1.PortId, params v1.AssignNetworkParameter) (*v1.InterfacePort, error)
	ReadTrafficByPort(ctx context.Context, serverId v1.ServerId, portId v1.PortId, params v1.ReadServerTrafficByPortParams) (*v1.TrafficGraph, error)
	PowerControl(ctx context.Context, serverId v1.ServerId, operation v1.ServerPowerOperations) error
	ReadPowerStatus(ctx context.Context, serverId v1.ServerId) (*v1.ServerPowerStatus, error)
	ReadRAIDStatus(ctx context.Context, serverId v1.ServerId, refresh bool) (*v1.RaidStatus, error)
}

type ServerOp struct {
	client *Client
}

func NewServerOp(client *Client) ServerAPI {
	return &ServerOp{client: client}
}

func (op *ServerOp) List(ctx context.Context, params *v1.ListServersParams) (*v1.Servers, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ListServersWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return response.Result()
}

func (op *ServerOp) Read(ctx context.Context, serverId v1.ServerId) (*v1.Server, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ReadServerWithResponse(ctx, serverId)
	if err != nil {
		return nil, err
	}
	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.Server, nil
}

func (op *ServerOp) ListOSImages(ctx context.Context, serverId v1.ServerId) ([]*v1.OsImage, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ListOSImagesWithResponse(ctx, serverId)
	if err != nil {
		return nil, err
	}
	result, err := response.Result()
	if err != nil {
		return nil, err
	}

	var images []*v1.OsImage
	for i := range result.OsImages {
		images = append(images, &result.OsImages[i])
	}
	return images, nil
}

func (op *ServerOp) OSInstall(ctx context.Context, serverId v1.ServerId, params v1.OsInstallParameter) error {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return err
	}
	headers := &v1.OSInstallParams{
		XRequestedWith: v1.OSInstallParamsXRequestedWith(v1.XMLHttpRequest),
	}
	_, err = apiClient.OSInstall(ctx, serverId, headers, params)
	return err
}

func (op *ServerOp) ReadPortChannel(ctx context.Context, serverId v1.ServerId, portChannelId v1.PortChannelId) (*v1.PortChannel, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ReadServerPortChannelWithResponse(ctx, serverId, portChannelId)
	if err != nil {
		return nil, err
	}
	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.PortChannel, nil
}

func (op *ServerOp) ConfigureBonding(ctx context.Context, serverId v1.ServerId, portChannelId v1.PortChannelId, params v1.ConfigureBondingParameter) (*v1.PortChannel, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	headers := &v1.ServerConfigureBondingParams{
		XRequestedWith: v1.ServerConfigureBondingParamsXRequestedWith(v1.XMLHttpRequest),
	}
	response, err := apiClient.ServerConfigureBondingWithResponse(ctx, serverId, portChannelId, headers, params)
	if err != nil {
		return nil, err
	}

	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.PortChannel, nil
}

func (op *ServerOp) ReadPort(ctx context.Context, serverId v1.ServerId, portId v1.PortId) (*v1.InterfacePort, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ReadServerPortWithResponse(ctx, serverId, portId)
	if err != nil {
		return nil, err
	}
	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.Port, nil
}

func (op *ServerOp) UpdatePort(ctx context.Context, serverId v1.ServerId, portId v1.PortId, params v1.UpdateServerPortParameter) (*v1.InterfacePort, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	headers := &v1.UpdateServerPortParams{
		XRequestedWith: v1.UpdateServerPortParamsXRequestedWith(v1.XMLHttpRequest),
	}
	response, err := apiClient.UpdateServerPortWithResponse(ctx, serverId, portId, headers, params)
	if err != nil {
		return nil, err
	}

	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.Port, nil
}

func (op *ServerOp) EnablePort(ctx context.Context, serverId v1.ServerId, portId v1.PortId, enable bool) (*v1.InterfacePort, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	headers := &v1.EnableServerPortParams{
		XRequestedWith: v1.EnableServerPortParamsXRequestedWith(v1.XMLHttpRequest),
	}
	response, err := apiClient.EnableServerPortWithResponse(ctx, serverId, portId, headers, v1.EnableServerPortJSONRequestBody{Enable: enable})
	if err != nil {
		return nil, err
	}

	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.Port, nil
}

func (op *ServerOp) AssignNetwork(ctx context.Context, serverId v1.ServerId, portId v1.PortId, params v1.AssignNetworkParameter) (*v1.InterfacePort, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	headers := &v1.ServerAssignNetworkParams{
		XRequestedWith: v1.ServerAssignNetworkParamsXRequestedWith(v1.XMLHttpRequest),
	}
	response, err := apiClient.ServerAssignNetworkWithResponse(ctx, serverId, portId, headers, params)
	if err != nil {
		return nil, err
	}

	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.Port, nil
}

func (op *ServerOp) ReadTrafficByPort(ctx context.Context, serverId v1.ServerId, portId v1.PortId, params v1.ReadServerTrafficByPortParams) (*v1.TrafficGraph, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ReadServerTrafficByPortWithResponse(ctx, serverId, portId, &params)
	if err != nil {
		return nil, err
	}
	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.TrafficGraph, nil
}

func (op *ServerOp) PowerControl(ctx context.Context, serverId v1.ServerId, operation v1.ServerPowerOperations) error {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return err
	}
	headers := &v1.ServerPowerControlParams{
		XRequestedWith: v1.ServerPowerControlParamsXRequestedWith(v1.XMLHttpRequest),
	}
	_, err = apiClient.ServerPowerControl(ctx, serverId, headers, v1.ServerPowerControlJSONRequestBody{Operation: operation})
	return err
}

func (op *ServerOp) ReadPowerStatus(ctx context.Context, serverId v1.ServerId) (*v1.ServerPowerStatus, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ReadServerPowerStatusWithResponse(ctx, serverId)
	if err != nil {
		return nil, err
	}
	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.PowerStatus, nil
}

func (op *ServerOp) ReadRAIDStatus(ctx context.Context, serverId v1.ServerId, refresh bool) (*v1.RaidStatus, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ReadRAIDStatusWithResponse(ctx, serverId, &v1.ReadRAIDStatusParams{Refresh: &refresh})
	if err != nil {
		return nil, err
	}
	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.RaidStatus, nil
}
