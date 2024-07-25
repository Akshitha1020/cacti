/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the CordappDeploymentConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CordappDeploymentConfig{}

// CordappDeploymentConfig struct for CordappDeploymentConfig
type CordappDeploymentConfig struct {
	SshCredentials CordaNodeSshCredentials `json:"sshCredentials"`
	RpcCredentials CordaRpcCredentials `json:"rpcCredentials"`
	// The shell command to execute in order to start back up a Corda node after having placed new jars in the cordapp directory of said node.
	CordaNodeStartCmd string `json:"cordaNodeStartCmd"`
	// The absolute file system path where the Corda Node is expecting deployed Cordapp jar files to be placed.
	CordappDir string `json:"cordappDir"`
	// The absolute file system path where the corda.jar file of the node can be found. This is used to execute database schema migrations where applicable (H2 database in use in development environments).
	CordaJarPath string `json:"cordaJarPath"`
	// The absolute file system path where the base directory of the Corda node can be found. This is used to pass in to corda.jar when being invoked for certain tasks such as executing database schema migrations for a deployed contract.
	NodeBaseDirPath string `json:"nodeBaseDirPath"`
}

// NewCordappDeploymentConfig instantiates a new CordappDeploymentConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCordappDeploymentConfig(sshCredentials CordaNodeSshCredentials, rpcCredentials CordaRpcCredentials, cordaNodeStartCmd string, cordappDir string, cordaJarPath string, nodeBaseDirPath string) *CordappDeploymentConfig {
	this := CordappDeploymentConfig{}
	this.SshCredentials = sshCredentials
	this.RpcCredentials = rpcCredentials
	this.CordaNodeStartCmd = cordaNodeStartCmd
	this.CordappDir = cordappDir
	this.CordaJarPath = cordaJarPath
	this.NodeBaseDirPath = nodeBaseDirPath
	return &this
}

// NewCordappDeploymentConfigWithDefaults instantiates a new CordappDeploymentConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCordappDeploymentConfigWithDefaults() *CordappDeploymentConfig {
	this := CordappDeploymentConfig{}
	return &this
}

// GetSshCredentials returns the SshCredentials field value
func (o *CordappDeploymentConfig) GetSshCredentials() CordaNodeSshCredentials {
	if o == nil {
		var ret CordaNodeSshCredentials
		return ret
	}

	return o.SshCredentials
}

// GetSshCredentialsOk returns a tuple with the SshCredentials field value
// and a boolean to check if the value has been set.
func (o *CordappDeploymentConfig) GetSshCredentialsOk() (*CordaNodeSshCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SshCredentials, true
}

// SetSshCredentials sets field value
func (o *CordappDeploymentConfig) SetSshCredentials(v CordaNodeSshCredentials) {
	o.SshCredentials = v
}

// GetRpcCredentials returns the RpcCredentials field value
func (o *CordappDeploymentConfig) GetRpcCredentials() CordaRpcCredentials {
	if o == nil {
		var ret CordaRpcCredentials
		return ret
	}

	return o.RpcCredentials
}

// GetRpcCredentialsOk returns a tuple with the RpcCredentials field value
// and a boolean to check if the value has been set.
func (o *CordappDeploymentConfig) GetRpcCredentialsOk() (*CordaRpcCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RpcCredentials, true
}

// SetRpcCredentials sets field value
func (o *CordappDeploymentConfig) SetRpcCredentials(v CordaRpcCredentials) {
	o.RpcCredentials = v
}

// GetCordaNodeStartCmd returns the CordaNodeStartCmd field value
func (o *CordappDeploymentConfig) GetCordaNodeStartCmd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CordaNodeStartCmd
}

// GetCordaNodeStartCmdOk returns a tuple with the CordaNodeStartCmd field value
// and a boolean to check if the value has been set.
func (o *CordappDeploymentConfig) GetCordaNodeStartCmdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CordaNodeStartCmd, true
}

// SetCordaNodeStartCmd sets field value
func (o *CordappDeploymentConfig) SetCordaNodeStartCmd(v string) {
	o.CordaNodeStartCmd = v
}

// GetCordappDir returns the CordappDir field value
func (o *CordappDeploymentConfig) GetCordappDir() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CordappDir
}

// GetCordappDirOk returns a tuple with the CordappDir field value
// and a boolean to check if the value has been set.
func (o *CordappDeploymentConfig) GetCordappDirOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CordappDir, true
}

// SetCordappDir sets field value
func (o *CordappDeploymentConfig) SetCordappDir(v string) {
	o.CordappDir = v
}

// GetCordaJarPath returns the CordaJarPath field value
func (o *CordappDeploymentConfig) GetCordaJarPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CordaJarPath
}

// GetCordaJarPathOk returns a tuple with the CordaJarPath field value
// and a boolean to check if the value has been set.
func (o *CordappDeploymentConfig) GetCordaJarPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CordaJarPath, true
}

// SetCordaJarPath sets field value
func (o *CordappDeploymentConfig) SetCordaJarPath(v string) {
	o.CordaJarPath = v
}

// GetNodeBaseDirPath returns the NodeBaseDirPath field value
func (o *CordappDeploymentConfig) GetNodeBaseDirPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeBaseDirPath
}

// GetNodeBaseDirPathOk returns a tuple with the NodeBaseDirPath field value
// and a boolean to check if the value has been set.
func (o *CordappDeploymentConfig) GetNodeBaseDirPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeBaseDirPath, true
}

// SetNodeBaseDirPath sets field value
func (o *CordappDeploymentConfig) SetNodeBaseDirPath(v string) {
	o.NodeBaseDirPath = v
}

func (o CordappDeploymentConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CordappDeploymentConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sshCredentials"] = o.SshCredentials
	toSerialize["rpcCredentials"] = o.RpcCredentials
	toSerialize["cordaNodeStartCmd"] = o.CordaNodeStartCmd
	toSerialize["cordappDir"] = o.CordappDir
	toSerialize["cordaJarPath"] = o.CordaJarPath
	toSerialize["nodeBaseDirPath"] = o.NodeBaseDirPath
	return toSerialize, nil
}

type NullableCordappDeploymentConfig struct {
	value *CordappDeploymentConfig
	isSet bool
}

func (v NullableCordappDeploymentConfig) Get() *CordappDeploymentConfig {
	return v.value
}

func (v *NullableCordappDeploymentConfig) Set(val *CordappDeploymentConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCordappDeploymentConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCordappDeploymentConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCordappDeploymentConfig(val *CordappDeploymentConfig) *NullableCordappDeploymentConfig {
	return &NullableCordappDeploymentConfig{value: val, isSet: true}
}

func (v NullableCordappDeploymentConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCordappDeploymentConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


