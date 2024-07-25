/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-satp-hermes

import (
	"encoding/json"
)

// checks if the CommitFinalV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommitFinalV1Request{}

// CommitFinalV1Request struct for CommitFinalV1Request
type CommitFinalV1Request struct {
	SessionID string `json:"sessionID"`
	MessageType string `json:"messageType"`
	ClientIdentityPubkey string `json:"clientIdentityPubkey"`
	ServerIdentityPubkey string `json:"serverIdentityPubkey"`
	CommitFinalClaim string `json:"commitFinalClaim"`
	CommitFinalClaimFormat map[string]interface{} `json:"commitFinalClaimFormat,omitempty"`
	HashCommitPrepareAck string `json:"hashCommitPrepareAck"`
	ClientTransferNumber NullableInt32 `json:"clientTransferNumber,omitempty"`
	Signature string `json:"signature"`
	SequenceNumber float32 `json:"sequenceNumber"`
}

// NewCommitFinalV1Request instantiates a new CommitFinalV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitFinalV1Request(sessionID string, messageType string, clientIdentityPubkey string, serverIdentityPubkey string, commitFinalClaim string, hashCommitPrepareAck string, signature string, sequenceNumber float32) *CommitFinalV1Request {
	this := CommitFinalV1Request{}
	this.SessionID = sessionID
	this.MessageType = messageType
	this.ClientIdentityPubkey = clientIdentityPubkey
	this.ServerIdentityPubkey = serverIdentityPubkey
	this.CommitFinalClaim = commitFinalClaim
	this.HashCommitPrepareAck = hashCommitPrepareAck
	this.Signature = signature
	this.SequenceNumber = sequenceNumber
	return &this
}

// NewCommitFinalV1RequestWithDefaults instantiates a new CommitFinalV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitFinalV1RequestWithDefaults() *CommitFinalV1Request {
	this := CommitFinalV1Request{}
	return &this
}

// GetSessionID returns the SessionID field value
func (o *CommitFinalV1Request) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *CommitFinalV1Request) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *CommitFinalV1Request) SetSessionID(v string) {
	o.SessionID = v
}

// GetMessageType returns the MessageType field value
func (o *CommitFinalV1Request) GetMessageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *CommitFinalV1Request) GetMessageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *CommitFinalV1Request) SetMessageType(v string) {
	o.MessageType = v
}

// GetClientIdentityPubkey returns the ClientIdentityPubkey field value
func (o *CommitFinalV1Request) GetClientIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientIdentityPubkey
}

// GetClientIdentityPubkeyOk returns a tuple with the ClientIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *CommitFinalV1Request) GetClientIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientIdentityPubkey, true
}

// SetClientIdentityPubkey sets field value
func (o *CommitFinalV1Request) SetClientIdentityPubkey(v string) {
	o.ClientIdentityPubkey = v
}

// GetServerIdentityPubkey returns the ServerIdentityPubkey field value
func (o *CommitFinalV1Request) GetServerIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerIdentityPubkey
}

// GetServerIdentityPubkeyOk returns a tuple with the ServerIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *CommitFinalV1Request) GetServerIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerIdentityPubkey, true
}

// SetServerIdentityPubkey sets field value
func (o *CommitFinalV1Request) SetServerIdentityPubkey(v string) {
	o.ServerIdentityPubkey = v
}

// GetCommitFinalClaim returns the CommitFinalClaim field value
func (o *CommitFinalV1Request) GetCommitFinalClaim() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitFinalClaim
}

// GetCommitFinalClaimOk returns a tuple with the CommitFinalClaim field value
// and a boolean to check if the value has been set.
func (o *CommitFinalV1Request) GetCommitFinalClaimOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitFinalClaim, true
}

// SetCommitFinalClaim sets field value
func (o *CommitFinalV1Request) SetCommitFinalClaim(v string) {
	o.CommitFinalClaim = v
}

// GetCommitFinalClaimFormat returns the CommitFinalClaimFormat field value if set, zero value otherwise.
func (o *CommitFinalV1Request) GetCommitFinalClaimFormat() map[string]interface{} {
	if o == nil || IsNil(o.CommitFinalClaimFormat) {
		var ret map[string]interface{}
		return ret
	}
	return o.CommitFinalClaimFormat
}

// GetCommitFinalClaimFormatOk returns a tuple with the CommitFinalClaimFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitFinalV1Request) GetCommitFinalClaimFormatOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CommitFinalClaimFormat) {
		return map[string]interface{}{}, false
	}
	return o.CommitFinalClaimFormat, true
}

// HasCommitFinalClaimFormat returns a boolean if a field has been set.
func (o *CommitFinalV1Request) HasCommitFinalClaimFormat() bool {
	if o != nil && !IsNil(o.CommitFinalClaimFormat) {
		return true
	}

	return false
}

// SetCommitFinalClaimFormat gets a reference to the given map[string]interface{} and assigns it to the CommitFinalClaimFormat field.
func (o *CommitFinalV1Request) SetCommitFinalClaimFormat(v map[string]interface{}) {
	o.CommitFinalClaimFormat = v
}

// GetHashCommitPrepareAck returns the HashCommitPrepareAck field value
func (o *CommitFinalV1Request) GetHashCommitPrepareAck() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HashCommitPrepareAck
}

// GetHashCommitPrepareAckOk returns a tuple with the HashCommitPrepareAck field value
// and a boolean to check if the value has been set.
func (o *CommitFinalV1Request) GetHashCommitPrepareAckOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashCommitPrepareAck, true
}

// SetHashCommitPrepareAck sets field value
func (o *CommitFinalV1Request) SetHashCommitPrepareAck(v string) {
	o.HashCommitPrepareAck = v
}

// GetClientTransferNumber returns the ClientTransferNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommitFinalV1Request) GetClientTransferNumber() int32 {
	if o == nil || IsNil(o.ClientTransferNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.ClientTransferNumber.Get()
}

// GetClientTransferNumberOk returns a tuple with the ClientTransferNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommitFinalV1Request) GetClientTransferNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientTransferNumber.Get(), o.ClientTransferNumber.IsSet()
}

// HasClientTransferNumber returns a boolean if a field has been set.
func (o *CommitFinalV1Request) HasClientTransferNumber() bool {
	if o != nil && o.ClientTransferNumber.IsSet() {
		return true
	}

	return false
}

// SetClientTransferNumber gets a reference to the given NullableInt32 and assigns it to the ClientTransferNumber field.
func (o *CommitFinalV1Request) SetClientTransferNumber(v int32) {
	o.ClientTransferNumber.Set(&v)
}
// SetClientTransferNumberNil sets the value for ClientTransferNumber to be an explicit nil
func (o *CommitFinalV1Request) SetClientTransferNumberNil() {
	o.ClientTransferNumber.Set(nil)
}

// UnsetClientTransferNumber ensures that no value is present for ClientTransferNumber, not even an explicit nil
func (o *CommitFinalV1Request) UnsetClientTransferNumber() {
	o.ClientTransferNumber.Unset()
}

// GetSignature returns the Signature field value
func (o *CommitFinalV1Request) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *CommitFinalV1Request) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *CommitFinalV1Request) SetSignature(v string) {
	o.Signature = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *CommitFinalV1Request) GetSequenceNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *CommitFinalV1Request) GetSequenceNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *CommitFinalV1Request) SetSequenceNumber(v float32) {
	o.SequenceNumber = v
}

func (o CommitFinalV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitFinalV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionID"] = o.SessionID
	toSerialize["messageType"] = o.MessageType
	toSerialize["clientIdentityPubkey"] = o.ClientIdentityPubkey
	toSerialize["serverIdentityPubkey"] = o.ServerIdentityPubkey
	toSerialize["commitFinalClaim"] = o.CommitFinalClaim
	if !IsNil(o.CommitFinalClaimFormat) {
		toSerialize["commitFinalClaimFormat"] = o.CommitFinalClaimFormat
	}
	toSerialize["hashCommitPrepareAck"] = o.HashCommitPrepareAck
	if o.ClientTransferNumber.IsSet() {
		toSerialize["clientTransferNumber"] = o.ClientTransferNumber.Get()
	}
	toSerialize["signature"] = o.Signature
	toSerialize["sequenceNumber"] = o.SequenceNumber
	return toSerialize, nil
}

type NullableCommitFinalV1Request struct {
	value *CommitFinalV1Request
	isSet bool
}

func (v NullableCommitFinalV1Request) Get() *CommitFinalV1Request {
	return v.value
}

func (v *NullableCommitFinalV1Request) Set(val *CommitFinalV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitFinalV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitFinalV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitFinalV1Request(val *CommitFinalV1Request) *NullableCommitFinalV1Request {
	return &NullableCommitFinalV1Request{value: val, isSet: true}
}

func (v NullableCommitFinalV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitFinalV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


