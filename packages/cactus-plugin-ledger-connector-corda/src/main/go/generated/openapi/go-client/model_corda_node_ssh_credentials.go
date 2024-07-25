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

// checks if the CordaNodeSshCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CordaNodeSshCredentials{}

// CordaNodeSshCredentials struct for CordaNodeSshCredentials
type CordaNodeSshCredentials struct {
	HostKeyEntry string `json:"hostKeyEntry"`
	Username string `json:"username"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	Port int32 `json:"port"`
}

// NewCordaNodeSshCredentials instantiates a new CordaNodeSshCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCordaNodeSshCredentials(hostKeyEntry string, username string, password string, hostname string, port int32) *CordaNodeSshCredentials {
	this := CordaNodeSshCredentials{}
	this.HostKeyEntry = hostKeyEntry
	this.Username = username
	this.Password = password
	this.Hostname = hostname
	this.Port = port
	return &this
}

// NewCordaNodeSshCredentialsWithDefaults instantiates a new CordaNodeSshCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCordaNodeSshCredentialsWithDefaults() *CordaNodeSshCredentials {
	this := CordaNodeSshCredentials{}
	return &this
}

// GetHostKeyEntry returns the HostKeyEntry field value
func (o *CordaNodeSshCredentials) GetHostKeyEntry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostKeyEntry
}

// GetHostKeyEntryOk returns a tuple with the HostKeyEntry field value
// and a boolean to check if the value has been set.
func (o *CordaNodeSshCredentials) GetHostKeyEntryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostKeyEntry, true
}

// SetHostKeyEntry sets field value
func (o *CordaNodeSshCredentials) SetHostKeyEntry(v string) {
	o.HostKeyEntry = v
}

// GetUsername returns the Username field value
func (o *CordaNodeSshCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CordaNodeSshCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CordaNodeSshCredentials) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *CordaNodeSshCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CordaNodeSshCredentials) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CordaNodeSshCredentials) SetPassword(v string) {
	o.Password = v
}

// GetHostname returns the Hostname field value
func (o *CordaNodeSshCredentials) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *CordaNodeSshCredentials) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *CordaNodeSshCredentials) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value
func (o *CordaNodeSshCredentials) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CordaNodeSshCredentials) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CordaNodeSshCredentials) SetPort(v int32) {
	o.Port = v
}

func (o CordaNodeSshCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CordaNodeSshCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hostKeyEntry"] = o.HostKeyEntry
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	toSerialize["hostname"] = o.Hostname
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

type NullableCordaNodeSshCredentials struct {
	value *CordaNodeSshCredentials
	isSet bool
}

func (v NullableCordaNodeSshCredentials) Get() *CordaNodeSshCredentials {
	return v.value
}

func (v *NullableCordaNodeSshCredentials) Set(val *CordaNodeSshCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableCordaNodeSshCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableCordaNodeSshCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCordaNodeSshCredentials(val *CordaNodeSshCredentials) *NullableCordaNodeSshCredentials {
	return &NullableCordaNodeSshCredentials{value: val, isSet: true}
}

func (v NullableCordaNodeSshCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCordaNodeSshCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


