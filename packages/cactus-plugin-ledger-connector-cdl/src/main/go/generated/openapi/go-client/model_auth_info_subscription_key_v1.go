/*
Hyperledger Cacti Plugin - Connector CDL

Can perform basic tasks on Fujitsu CDL service.

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-cdl

import (
	"encoding/json"
)

// checks if the AuthInfoSubscriptionKeyV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthInfoSubscriptionKeyV1{}

// AuthInfoSubscriptionKeyV1 struct for AuthInfoSubscriptionKeyV1
type AuthInfoSubscriptionKeyV1 struct {
	SubscriptionKey string `json:"subscriptionKey"`
	TrustAgentId string `json:"trustAgentId"`
	TrustAgentRole string `json:"trustAgentRole"`
	TrustUserId string `json:"trustUserId"`
	TrustUserRole string `json:"trustUserRole"`
}

// NewAuthInfoSubscriptionKeyV1 instantiates a new AuthInfoSubscriptionKeyV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthInfoSubscriptionKeyV1(subscriptionKey string, trustAgentId string, trustAgentRole string, trustUserId string, trustUserRole string) *AuthInfoSubscriptionKeyV1 {
	this := AuthInfoSubscriptionKeyV1{}
	this.SubscriptionKey = subscriptionKey
	this.TrustAgentId = trustAgentId
	this.TrustAgentRole = trustAgentRole
	this.TrustUserId = trustUserId
	this.TrustUserRole = trustUserRole
	return &this
}

// NewAuthInfoSubscriptionKeyV1WithDefaults instantiates a new AuthInfoSubscriptionKeyV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthInfoSubscriptionKeyV1WithDefaults() *AuthInfoSubscriptionKeyV1 {
	this := AuthInfoSubscriptionKeyV1{}
	return &this
}

// GetSubscriptionKey returns the SubscriptionKey field value
func (o *AuthInfoSubscriptionKeyV1) GetSubscriptionKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionKey
}

// GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field value
// and a boolean to check if the value has been set.
func (o *AuthInfoSubscriptionKeyV1) GetSubscriptionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionKey, true
}

// SetSubscriptionKey sets field value
func (o *AuthInfoSubscriptionKeyV1) SetSubscriptionKey(v string) {
	o.SubscriptionKey = v
}

// GetTrustAgentId returns the TrustAgentId field value
func (o *AuthInfoSubscriptionKeyV1) GetTrustAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrustAgentId
}

// GetTrustAgentIdOk returns a tuple with the TrustAgentId field value
// and a boolean to check if the value has been set.
func (o *AuthInfoSubscriptionKeyV1) GetTrustAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustAgentId, true
}

// SetTrustAgentId sets field value
func (o *AuthInfoSubscriptionKeyV1) SetTrustAgentId(v string) {
	o.TrustAgentId = v
}

// GetTrustAgentRole returns the TrustAgentRole field value
func (o *AuthInfoSubscriptionKeyV1) GetTrustAgentRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrustAgentRole
}

// GetTrustAgentRoleOk returns a tuple with the TrustAgentRole field value
// and a boolean to check if the value has been set.
func (o *AuthInfoSubscriptionKeyV1) GetTrustAgentRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustAgentRole, true
}

// SetTrustAgentRole sets field value
func (o *AuthInfoSubscriptionKeyV1) SetTrustAgentRole(v string) {
	o.TrustAgentRole = v
}

// GetTrustUserId returns the TrustUserId field value
func (o *AuthInfoSubscriptionKeyV1) GetTrustUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrustUserId
}

// GetTrustUserIdOk returns a tuple with the TrustUserId field value
// and a boolean to check if the value has been set.
func (o *AuthInfoSubscriptionKeyV1) GetTrustUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustUserId, true
}

// SetTrustUserId sets field value
func (o *AuthInfoSubscriptionKeyV1) SetTrustUserId(v string) {
	o.TrustUserId = v
}

// GetTrustUserRole returns the TrustUserRole field value
func (o *AuthInfoSubscriptionKeyV1) GetTrustUserRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrustUserRole
}

// GetTrustUserRoleOk returns a tuple with the TrustUserRole field value
// and a boolean to check if the value has been set.
func (o *AuthInfoSubscriptionKeyV1) GetTrustUserRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustUserRole, true
}

// SetTrustUserRole sets field value
func (o *AuthInfoSubscriptionKeyV1) SetTrustUserRole(v string) {
	o.TrustUserRole = v
}

func (o AuthInfoSubscriptionKeyV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthInfoSubscriptionKeyV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionKey"] = o.SubscriptionKey
	toSerialize["trustAgentId"] = o.TrustAgentId
	toSerialize["trustAgentRole"] = o.TrustAgentRole
	toSerialize["trustUserId"] = o.TrustUserId
	toSerialize["trustUserRole"] = o.TrustUserRole
	return toSerialize, nil
}

type NullableAuthInfoSubscriptionKeyV1 struct {
	value *AuthInfoSubscriptionKeyV1
	isSet bool
}

func (v NullableAuthInfoSubscriptionKeyV1) Get() *AuthInfoSubscriptionKeyV1 {
	return v.value
}

func (v *NullableAuthInfoSubscriptionKeyV1) Set(val *AuthInfoSubscriptionKeyV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthInfoSubscriptionKeyV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthInfoSubscriptionKeyV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthInfoSubscriptionKeyV1(val *AuthInfoSubscriptionKeyV1) *NullableAuthInfoSubscriptionKeyV1 {
	return &NullableAuthInfoSubscriptionKeyV1{value: val, isSet: true}
}

func (v NullableAuthInfoSubscriptionKeyV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthInfoSubscriptionKeyV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


