// Copyright 2021-2022 The phy-api-go authors
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

// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package v1

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 専用グローバルネットワーク 一覧
	// (GET /dedicated_subnets/)
	ListDedicatedSubnets(c *gin.Context, params ListDedicatedSubnetsParams)
	// 専用グローバルネットワーク
	// (GET /dedicated_subnets/{dedicated_subnet_id}/)
	ReadDedicatedSubnet(c *gin.Context, dedicatedSubnetId DedicatedSubnetId, params ReadDedicatedSubnetParams)
	// ローカルネットワーク 一覧
	// (GET /private_networks/)
	ListPrivateNetworks(c *gin.Context, params ListPrivateNetworksParams)
	// ローカルネットワーク 詳細
	// (GET /private_networks/{private_network_id}/)
	ReadPrivateNetwork(c *gin.Context, privateNetworkId PrivateNetworkId)
	// サーバー一覧
	// (GET /servers/)
	ListServers(c *gin.Context, params ListServersParams)
	// サーバー
	// (GET /servers/{server_id}/)
	ReadServer(c *gin.Context, serverId ServerId)
	// インストール可能OS一覧
	// (GET /servers/{server_id}/os_images/)
	ListOSImages(c *gin.Context, serverId ServerId)
	// OSインストールの実行
	// (POST /servers/{server_id}/os_install/)
	OSInstall(c *gin.Context, serverId ServerId, params OSInstallParams)
	// ポートチャネル状態取得
	// (GET /servers/{server_id}/port_channels/{port_channel_id}/)
	ReadServerPortChannel(c *gin.Context, serverId ServerId, portChannelId PortChannelId)
	// ポートチャネル ボンディング設定
	// (POST /servers/{server_id}/port_channels/{port_channel_id}/configure_bonding/)
	ServerConfigureBonding(c *gin.Context, serverId ServerId, portChannelId PortChannelId, params ServerConfigureBondingParams)
	// ポート情報取得
	// (GET /servers/{server_id}/ports/{port_id}/)
	ReadServerPort(c *gin.Context, serverId ServerId, portId PortId)
	// ポート名称設定
	// (PATCH /servers/{server_id}/ports/{port_id}/)
	UpdateServerPort(c *gin.Context, serverId ServerId, portId PortId, params UpdateServerPortParams)
	// ネットワーク接続設定の変更
	// (POST /servers/{server_id}/ports/{port_id}/assign_network/)
	ServerAssignNetwork(c *gin.Context, serverId ServerId, portId PortId, params ServerAssignNetworkParams)
	// ポート有効/無効設定
	// (POST /servers/{server_id}/ports/{port_id}/enable/)
	EnableServerPort(c *gin.Context, serverId ServerId, portId PortId, params EnableServerPortParams)
	// トラフィックデータ取得
	// (GET /servers/{server_id}/ports/{port_id}/traffic_graph/)
	ReadServerTrafficByPort(c *gin.Context, serverId ServerId, portId PortId, params ReadServerTrafficByPortParams)
	// サーバーの電源操作
	// (POST /servers/{server_id}/power_control/)
	ServerPowerControl(c *gin.Context, serverId ServerId, params ServerPowerControlParams)
	// サーバーの電源情報を取得する
	// (GET /servers/{server_id}/power_status/)
	ReadServerPowerStatus(c *gin.Context, serverId ServerId)
	// サーバーのRAID状態を取得
	// (GET /servers/{server_id}/raid_status/)
	ReadRAIDStatus(c *gin.Context, serverId ServerId, params ReadRAIDStatusParams)
	// サービス一覧
	// (GET /services/)
	ListServices(c *gin.Context, params ListServicesParams)
	// サービス 詳細
	// (GET /services/{service_id}/)
	ReadService(c *gin.Context, serviceId ServiceId)
	// サービスの名称・説明の変更
	// (PATCH /services/{service_id}/)
	UpdateService(c *gin.Context, serviceId ServiceId, params UpdateServiceParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// ListDedicatedSubnets operation middleware
func (siw *ServerInterfaceWrapper) ListDedicatedSubnets(c *gin.Context) {

	var err error

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListDedicatedSubnetsParams

	// ------------- Optional query parameter "tag" -------------
	if paramValue := c.Query("tag"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "tag", c.Request.URL.Query(), &params.Tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter tag: %s", err)})
		return
	}

	// ------------- Optional query parameter "free_word" -------------
	if paramValue := c.Query("free_word"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "free_word", c.Request.URL.Query(), &params.FreeWord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter free_word: %s", err)})
		return
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := c.Query("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter limit: %s", err)})
		return
	}

	// ------------- Optional query parameter "offset" -------------
	if paramValue := c.Query("offset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "offset", c.Request.URL.Query(), &params.Offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter offset: %s", err)})
		return
	}

	// ------------- Optional query parameter "ordering" -------------
	if paramValue := c.Query("ordering"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "ordering", c.Request.URL.Query(), &params.Ordering)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter ordering: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ListDedicatedSubnets(c, params)
}

// ReadDedicatedSubnet operation middleware
func (siw *ServerInterfaceWrapper) ReadDedicatedSubnet(c *gin.Context) {

	var err error

	// ------------- Path parameter "dedicated_subnet_id" -------------
	var dedicatedSubnetId DedicatedSubnetId

	err = runtime.BindStyledParameter("simple", false, "dedicated_subnet_id", c.Param("dedicated_subnet_id"), &dedicatedSubnetId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter dedicated_subnet_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ReadDedicatedSubnetParams

	// ------------- Optional query parameter "refresh" -------------
	if paramValue := c.Query("refresh"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "refresh", c.Request.URL.Query(), &params.Refresh)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter refresh: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ReadDedicatedSubnet(c, dedicatedSubnetId, params)
}

// ListPrivateNetworks operation middleware
func (siw *ServerInterfaceWrapper) ListPrivateNetworks(c *gin.Context) {

	var err error

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListPrivateNetworksParams

	// ------------- Optional query parameter "tag" -------------
	if paramValue := c.Query("tag"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "tag", c.Request.URL.Query(), &params.Tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter tag: %s", err)})
		return
	}

	// ------------- Optional query parameter "free_word" -------------
	if paramValue := c.Query("free_word"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "free_word", c.Request.URL.Query(), &params.FreeWord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter free_word: %s", err)})
		return
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := c.Query("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter limit: %s", err)})
		return
	}

	// ------------- Optional query parameter "offset" -------------
	if paramValue := c.Query("offset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "offset", c.Request.URL.Query(), &params.Offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter offset: %s", err)})
		return
	}

	// ------------- Optional query parameter "ordering" -------------
	if paramValue := c.Query("ordering"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "ordering", c.Request.URL.Query(), &params.Ordering)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter ordering: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ListPrivateNetworks(c, params)
}

// ReadPrivateNetwork operation middleware
func (siw *ServerInterfaceWrapper) ReadPrivateNetwork(c *gin.Context) {

	var err error

	// ------------- Path parameter "private_network_id" -------------
	var privateNetworkId PrivateNetworkId

	err = runtime.BindStyledParameter("simple", false, "private_network_id", c.Param("private_network_id"), &privateNetworkId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter private_network_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ReadPrivateNetwork(c, privateNetworkId)
}

// ListServers operation middleware
func (siw *ServerInterfaceWrapper) ListServers(c *gin.Context) {

	var err error

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListServersParams

	// ------------- Optional query parameter "power_status" -------------
	if paramValue := c.Query("power_status"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "power_status", c.Request.URL.Query(), &params.PowerStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter power_status: %s", err)})
		return
	}

	// ------------- Optional query parameter "internet" -------------
	if paramValue := c.Query("internet"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "internet", c.Request.URL.Query(), &params.Internet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter internet: %s", err)})
		return
	}

	// ------------- Optional query parameter "private_network" -------------
	if paramValue := c.Query("private_network"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "private_network", c.Request.URL.Query(), &params.PrivateNetwork)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter private_network: %s", err)})
		return
	}

	// ------------- Optional query parameter "tag" -------------
	if paramValue := c.Query("tag"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "tag", c.Request.URL.Query(), &params.Tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter tag: %s", err)})
		return
	}

	// ------------- Optional query parameter "free_word" -------------
	if paramValue := c.Query("free_word"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "free_word", c.Request.URL.Query(), &params.FreeWord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter free_word: %s", err)})
		return
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := c.Query("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter limit: %s", err)})
		return
	}

	// ------------- Optional query parameter "offset" -------------
	if paramValue := c.Query("offset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "offset", c.Request.URL.Query(), &params.Offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter offset: %s", err)})
		return
	}

	// ------------- Optional query parameter "ordering" -------------
	if paramValue := c.Query("ordering"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "ordering", c.Request.URL.Query(), &params.Ordering)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter ordering: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ListServers(c, params)
}

// ReadServer operation middleware
func (siw *ServerInterfaceWrapper) ReadServer(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ReadServer(c, serverId)
}

// ListOSImages operation middleware
func (siw *ServerInterfaceWrapper) ListOSImages(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ListOSImages(c, serverId)
}

// OSInstall operation middleware
func (siw *ServerInterfaceWrapper) OSInstall(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params OSInstallParams

	headers := c.Request.Header

	// ------------- Required header parameter "X-Requested-With" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Requested-With")]; found {
		var XRequestedWith OSInstallParamsXRequestedWith
		n := len(valueList)
		if n != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Expected one value for X-Requested-With, got %d", n)})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Requested-With", runtime.ParamLocationHeader, valueList[0], &XRequestedWith)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter X-Requested-With: %s", err)})
			return
		}

		params.XRequestedWith = XRequestedWith

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Header parameter X-Requested-With is required, but not found: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.OSInstall(c, serverId, params)
}

// ReadServerPortChannel operation middleware
func (siw *ServerInterfaceWrapper) ReadServerPortChannel(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	// ------------- Path parameter "port_channel_id" -------------
	var portChannelId PortChannelId

	err = runtime.BindStyledParameter("simple", false, "port_channel_id", c.Param("port_channel_id"), &portChannelId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter port_channel_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ReadServerPortChannel(c, serverId, portChannelId)
}

// ServerConfigureBonding operation middleware
func (siw *ServerInterfaceWrapper) ServerConfigureBonding(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	// ------------- Path parameter "port_channel_id" -------------
	var portChannelId PortChannelId

	err = runtime.BindStyledParameter("simple", false, "port_channel_id", c.Param("port_channel_id"), &portChannelId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter port_channel_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ServerConfigureBondingParams

	headers := c.Request.Header

	// ------------- Required header parameter "X-Requested-With" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Requested-With")]; found {
		var XRequestedWith ServerConfigureBondingParamsXRequestedWith
		n := len(valueList)
		if n != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Expected one value for X-Requested-With, got %d", n)})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Requested-With", runtime.ParamLocationHeader, valueList[0], &XRequestedWith)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter X-Requested-With: %s", err)})
			return
		}

		params.XRequestedWith = XRequestedWith

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Header parameter X-Requested-With is required, but not found: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServerConfigureBonding(c, serverId, portChannelId, params)
}

// ReadServerPort operation middleware
func (siw *ServerInterfaceWrapper) ReadServerPort(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	// ------------- Path parameter "port_id" -------------
	var portId PortId

	err = runtime.BindStyledParameter("simple", false, "port_id", c.Param("port_id"), &portId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter port_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ReadServerPort(c, serverId, portId)
}

// UpdateServerPort operation middleware
func (siw *ServerInterfaceWrapper) UpdateServerPort(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	// ------------- Path parameter "port_id" -------------
	var portId PortId

	err = runtime.BindStyledParameter("simple", false, "port_id", c.Param("port_id"), &portId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter port_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params UpdateServerPortParams

	headers := c.Request.Header

	// ------------- Required header parameter "X-Requested-With" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Requested-With")]; found {
		var XRequestedWith UpdateServerPortParamsXRequestedWith
		n := len(valueList)
		if n != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Expected one value for X-Requested-With, got %d", n)})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Requested-With", runtime.ParamLocationHeader, valueList[0], &XRequestedWith)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter X-Requested-With: %s", err)})
			return
		}

		params.XRequestedWith = XRequestedWith

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Header parameter X-Requested-With is required, but not found: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.UpdateServerPort(c, serverId, portId, params)
}

// ServerAssignNetwork operation middleware
func (siw *ServerInterfaceWrapper) ServerAssignNetwork(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	// ------------- Path parameter "port_id" -------------
	var portId PortId

	err = runtime.BindStyledParameter("simple", false, "port_id", c.Param("port_id"), &portId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter port_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ServerAssignNetworkParams

	headers := c.Request.Header

	// ------------- Required header parameter "X-Requested-With" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Requested-With")]; found {
		var XRequestedWith ServerAssignNetworkParamsXRequestedWith
		n := len(valueList)
		if n != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Expected one value for X-Requested-With, got %d", n)})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Requested-With", runtime.ParamLocationHeader, valueList[0], &XRequestedWith)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter X-Requested-With: %s", err)})
			return
		}

		params.XRequestedWith = XRequestedWith

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Header parameter X-Requested-With is required, but not found: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServerAssignNetwork(c, serverId, portId, params)
}

// EnableServerPort operation middleware
func (siw *ServerInterfaceWrapper) EnableServerPort(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	// ------------- Path parameter "port_id" -------------
	var portId PortId

	err = runtime.BindStyledParameter("simple", false, "port_id", c.Param("port_id"), &portId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter port_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params EnableServerPortParams

	headers := c.Request.Header

	// ------------- Required header parameter "X-Requested-With" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Requested-With")]; found {
		var XRequestedWith EnableServerPortParamsXRequestedWith
		n := len(valueList)
		if n != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Expected one value for X-Requested-With, got %d", n)})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Requested-With", runtime.ParamLocationHeader, valueList[0], &XRequestedWith)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter X-Requested-With: %s", err)})
			return
		}

		params.XRequestedWith = XRequestedWith

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Header parameter X-Requested-With is required, but not found: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.EnableServerPort(c, serverId, portId, params)
}

// ReadServerTrafficByPort operation middleware
func (siw *ServerInterfaceWrapper) ReadServerTrafficByPort(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	// ------------- Path parameter "port_id" -------------
	var portId PortId

	err = runtime.BindStyledParameter("simple", false, "port_id", c.Param("port_id"), &portId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter port_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ReadServerTrafficByPortParams

	// ------------- Optional query parameter "since" -------------
	if paramValue := c.Query("since"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "since", c.Request.URL.Query(), &params.Since)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter since: %s", err)})
		return
	}

	// ------------- Optional query parameter "until" -------------
	if paramValue := c.Query("until"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "until", c.Request.URL.Query(), &params.Until)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter until: %s", err)})
		return
	}

	// ------------- Optional query parameter "step" -------------
	if paramValue := c.Query("step"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "step", c.Request.URL.Query(), &params.Step)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter step: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ReadServerTrafficByPort(c, serverId, portId, params)
}

// ServerPowerControl operation middleware
func (siw *ServerInterfaceWrapper) ServerPowerControl(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ServerPowerControlParams

	headers := c.Request.Header

	// ------------- Required header parameter "X-Requested-With" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Requested-With")]; found {
		var XRequestedWith ServerPowerControlParamsXRequestedWith
		n := len(valueList)
		if n != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Expected one value for X-Requested-With, got %d", n)})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Requested-With", runtime.ParamLocationHeader, valueList[0], &XRequestedWith)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter X-Requested-With: %s", err)})
			return
		}

		params.XRequestedWith = XRequestedWith

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Header parameter X-Requested-With is required, but not found: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServerPowerControl(c, serverId, params)
}

// ReadServerPowerStatus operation middleware
func (siw *ServerInterfaceWrapper) ReadServerPowerStatus(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ReadServerPowerStatus(c, serverId)
}

// ReadRAIDStatus operation middleware
func (siw *ServerInterfaceWrapper) ReadRAIDStatus(c *gin.Context) {

	var err error

	// ------------- Path parameter "server_id" -------------
	var serverId ServerId

	err = runtime.BindStyledParameter("simple", false, "server_id", c.Param("server_id"), &serverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter server_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ReadRAIDStatusParams

	// ------------- Optional query parameter "refresh" -------------
	if paramValue := c.Query("refresh"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "refresh", c.Request.URL.Query(), &params.Refresh)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter refresh: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ReadRAIDStatus(c, serverId, params)
}

// ListServices operation middleware
func (siw *ServerInterfaceWrapper) ListServices(c *gin.Context) {

	var err error

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListServicesParams

	// ------------- Optional query parameter "product_category" -------------
	if paramValue := c.Query("product_category"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "product_category", c.Request.URL.Query(), &params.ProductCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter product_category: %s", err)})
		return
	}

	// ------------- Optional query parameter "tag" -------------
	if paramValue := c.Query("tag"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "tag", c.Request.URL.Query(), &params.Tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter tag: %s", err)})
		return
	}

	// ------------- Optional query parameter "free_word" -------------
	if paramValue := c.Query("free_word"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "free_word", c.Request.URL.Query(), &params.FreeWord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter free_word: %s", err)})
		return
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := c.Query("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter limit: %s", err)})
		return
	}

	// ------------- Optional query parameter "offset" -------------
	if paramValue := c.Query("offset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "offset", c.Request.URL.Query(), &params.Offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter offset: %s", err)})
		return
	}

	// ------------- Optional query parameter "ordering" -------------
	if paramValue := c.Query("ordering"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "ordering", c.Request.URL.Query(), &params.Ordering)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter ordering: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ListServices(c, params)
}

// ReadService operation middleware
func (siw *ServerInterfaceWrapper) ReadService(c *gin.Context) {

	var err error

	// ------------- Path parameter "service_id" -------------
	var serviceId ServiceId

	err = runtime.BindStyledParameter("simple", false, "service_id", c.Param("service_id"), &serviceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter service_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ReadService(c, serviceId)
}

// UpdateService operation middleware
func (siw *ServerInterfaceWrapper) UpdateService(c *gin.Context) {

	var err error

	// ------------- Path parameter "service_id" -------------
	var serviceId ServiceId

	err = runtime.BindStyledParameter("simple", false, "service_id", c.Param("service_id"), &serviceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter service_id: %s", err)})
		return
	}

	c.Set(Account_api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params UpdateServiceParams

	headers := c.Request.Header

	// ------------- Required header parameter "X-Requested-With" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Requested-With")]; found {
		var XRequestedWith UpdateServiceParamsXRequestedWith
		n := len(valueList)
		if n != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Expected one value for X-Requested-With, got %d", n)})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Requested-With", runtime.ParamLocationHeader, valueList[0], &XRequestedWith)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter X-Requested-With: %s", err)})
			return
		}

		params.XRequestedWith = XRequestedWith

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Header parameter X-Requested-With is required, but not found: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.UpdateService(c, serviceId, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.GET(options.BaseURL+"/dedicated_subnets/", wrapper.ListDedicatedSubnets)

	router.GET(options.BaseURL+"/dedicated_subnets/:dedicated_subnet_id/", wrapper.ReadDedicatedSubnet)

	router.GET(options.BaseURL+"/private_networks/", wrapper.ListPrivateNetworks)

	router.GET(options.BaseURL+"/private_networks/:private_network_id/", wrapper.ReadPrivateNetwork)

	router.GET(options.BaseURL+"/servers/", wrapper.ListServers)

	router.GET(options.BaseURL+"/servers/:server_id/", wrapper.ReadServer)

	router.GET(options.BaseURL+"/servers/:server_id/os_images/", wrapper.ListOSImages)

	router.POST(options.BaseURL+"/servers/:server_id/os_install/", wrapper.OSInstall)

	router.GET(options.BaseURL+"/servers/:server_id/port_channels/:port_channel_id/", wrapper.ReadServerPortChannel)

	router.POST(options.BaseURL+"/servers/:server_id/port_channels/:port_channel_id/configure_bonding/", wrapper.ServerConfigureBonding)

	router.GET(options.BaseURL+"/servers/:server_id/ports/:port_id/", wrapper.ReadServerPort)

	router.PATCH(options.BaseURL+"/servers/:server_id/ports/:port_id/", wrapper.UpdateServerPort)

	router.POST(options.BaseURL+"/servers/:server_id/ports/:port_id/assign_network/", wrapper.ServerAssignNetwork)

	router.POST(options.BaseURL+"/servers/:server_id/ports/:port_id/enable/", wrapper.EnableServerPort)

	router.GET(options.BaseURL+"/servers/:server_id/ports/:port_id/traffic_graph/", wrapper.ReadServerTrafficByPort)

	router.POST(options.BaseURL+"/servers/:server_id/power_control/", wrapper.ServerPowerControl)

	router.GET(options.BaseURL+"/servers/:server_id/power_status/", wrapper.ReadServerPowerStatus)

	router.GET(options.BaseURL+"/servers/:server_id/raid_status/", wrapper.ReadRAIDStatus)

	router.GET(options.BaseURL+"/services/", wrapper.ListServices)

	router.GET(options.BaseURL+"/services/:service_id/", wrapper.ReadService)

	router.PATCH(options.BaseURL+"/services/:service_id/", wrapper.UpdateService)

	return router
}
