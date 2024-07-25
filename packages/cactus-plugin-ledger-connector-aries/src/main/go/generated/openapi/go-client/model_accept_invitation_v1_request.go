/*
Hyperledger Cacti Plugin - Connector Aries

Can communicate with other Aries agents and Cacti Aries connectors

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-aries

import (
	"encoding/json"
)

// checks if the AcceptInvitationV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcceptInvitationV1Request{}

// AcceptInvitationV1Request Request for AcceptInvitation endpoint.
type AcceptInvitationV1Request struct {
	// Aries label of an agent to be used to connect using URL
	AgentName string `json:"agentName"`
	// Invitation URL generated by another aries agent.
	InvitationUrl string `json:"invitationUrl"`
}

// NewAcceptInvitationV1Request instantiates a new AcceptInvitationV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptInvitationV1Request(agentName string, invitationUrl string) *AcceptInvitationV1Request {
	this := AcceptInvitationV1Request{}
	this.AgentName = agentName
	this.InvitationUrl = invitationUrl
	return &this
}

// NewAcceptInvitationV1RequestWithDefaults instantiates a new AcceptInvitationV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptInvitationV1RequestWithDefaults() *AcceptInvitationV1Request {
	this := AcceptInvitationV1Request{}
	return &this
}

// GetAgentName returns the AgentName field value
func (o *AcceptInvitationV1Request) GetAgentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value
// and a boolean to check if the value has been set.
func (o *AcceptInvitationV1Request) GetAgentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentName, true
}

// SetAgentName sets field value
func (o *AcceptInvitationV1Request) SetAgentName(v string) {
	o.AgentName = v
}

// GetInvitationUrl returns the InvitationUrl field value
func (o *AcceptInvitationV1Request) GetInvitationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvitationUrl
}

// GetInvitationUrlOk returns a tuple with the InvitationUrl field value
// and a boolean to check if the value has been set.
func (o *AcceptInvitationV1Request) GetInvitationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvitationUrl, true
}

// SetInvitationUrl sets field value
func (o *AcceptInvitationV1Request) SetInvitationUrl(v string) {
	o.InvitationUrl = v
}

func (o AcceptInvitationV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptInvitationV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agentName"] = o.AgentName
	toSerialize["invitationUrl"] = o.InvitationUrl
	return toSerialize, nil
}

type NullableAcceptInvitationV1Request struct {
	value *AcceptInvitationV1Request
	isSet bool
}

func (v NullableAcceptInvitationV1Request) Get() *AcceptInvitationV1Request {
	return v.value
}

func (v *NullableAcceptInvitationV1Request) Set(val *AcceptInvitationV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptInvitationV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptInvitationV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptInvitationV1Request(val *AcceptInvitationV1Request) *NullableAcceptInvitationV1Request {
	return &NullableAcceptInvitationV1Request{value: val, isSet: true}
}

func (v NullableAcceptInvitationV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptInvitationV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


