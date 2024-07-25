/*
Hyperledger Cactus Plugin - Connector Xdai

Can perform basic tasks on a Xdai ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-xdai

import (
	"encoding/json"
)

// checks if the XdaiTransactionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XdaiTransactionConfig{}

// XdaiTransactionConfig struct for XdaiTransactionConfig
type XdaiTransactionConfig struct {
	RawTransaction *string `json:"rawTransaction,omitempty"`
	From *XdaiTransactionConfigFrom `json:"from,omitempty"`
	To *XdaiTransactionConfigTo `json:"to,omitempty"`
	Value *XdaiTransactionConfigFrom `json:"value,omitempty"`
	Gas *XdaiTransactionConfigFrom `json:"gas,omitempty"`
	GasPrice *XdaiTransactionConfigFrom `json:"gasPrice,omitempty"`
	Nonce *float32 `json:"nonce,omitempty"`
	Data *XdaiTransactionConfigTo `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _XdaiTransactionConfig XdaiTransactionConfig

// NewXdaiTransactionConfig instantiates a new XdaiTransactionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXdaiTransactionConfig() *XdaiTransactionConfig {
	this := XdaiTransactionConfig{}
	return &this
}

// NewXdaiTransactionConfigWithDefaults instantiates a new XdaiTransactionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXdaiTransactionConfigWithDefaults() *XdaiTransactionConfig {
	this := XdaiTransactionConfig{}
	return &this
}

// GetRawTransaction returns the RawTransaction field value if set, zero value otherwise.
func (o *XdaiTransactionConfig) GetRawTransaction() string {
	if o == nil || IsNil(o.RawTransaction) {
		var ret string
		return ret
	}
	return *o.RawTransaction
}

// GetRawTransactionOk returns a tuple with the RawTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XdaiTransactionConfig) GetRawTransactionOk() (*string, bool) {
	if o == nil || IsNil(o.RawTransaction) {
		return nil, false
	}
	return o.RawTransaction, true
}

// HasRawTransaction returns a boolean if a field has been set.
func (o *XdaiTransactionConfig) HasRawTransaction() bool {
	if o != nil && !IsNil(o.RawTransaction) {
		return true
	}

	return false
}

// SetRawTransaction gets a reference to the given string and assigns it to the RawTransaction field.
func (o *XdaiTransactionConfig) SetRawTransaction(v string) {
	o.RawTransaction = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *XdaiTransactionConfig) GetFrom() XdaiTransactionConfigFrom {
	if o == nil || IsNil(o.From) {
		var ret XdaiTransactionConfigFrom
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XdaiTransactionConfig) GetFromOk() (*XdaiTransactionConfigFrom, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *XdaiTransactionConfig) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given XdaiTransactionConfigFrom and assigns it to the From field.
func (o *XdaiTransactionConfig) SetFrom(v XdaiTransactionConfigFrom) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *XdaiTransactionConfig) GetTo() XdaiTransactionConfigTo {
	if o == nil || IsNil(o.To) {
		var ret XdaiTransactionConfigTo
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XdaiTransactionConfig) GetToOk() (*XdaiTransactionConfigTo, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *XdaiTransactionConfig) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given XdaiTransactionConfigTo and assigns it to the To field.
func (o *XdaiTransactionConfig) SetTo(v XdaiTransactionConfigTo) {
	o.To = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *XdaiTransactionConfig) GetValue() XdaiTransactionConfigFrom {
	if o == nil || IsNil(o.Value) {
		var ret XdaiTransactionConfigFrom
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XdaiTransactionConfig) GetValueOk() (*XdaiTransactionConfigFrom, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *XdaiTransactionConfig) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given XdaiTransactionConfigFrom and assigns it to the Value field.
func (o *XdaiTransactionConfig) SetValue(v XdaiTransactionConfigFrom) {
	o.Value = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *XdaiTransactionConfig) GetGas() XdaiTransactionConfigFrom {
	if o == nil || IsNil(o.Gas) {
		var ret XdaiTransactionConfigFrom
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XdaiTransactionConfig) GetGasOk() (*XdaiTransactionConfigFrom, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *XdaiTransactionConfig) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given XdaiTransactionConfigFrom and assigns it to the Gas field.
func (o *XdaiTransactionConfig) SetGas(v XdaiTransactionConfigFrom) {
	o.Gas = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *XdaiTransactionConfig) GetGasPrice() XdaiTransactionConfigFrom {
	if o == nil || IsNil(o.GasPrice) {
		var ret XdaiTransactionConfigFrom
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XdaiTransactionConfig) GetGasPriceOk() (*XdaiTransactionConfigFrom, bool) {
	if o == nil || IsNil(o.GasPrice) {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *XdaiTransactionConfig) HasGasPrice() bool {
	if o != nil && !IsNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given XdaiTransactionConfigFrom and assigns it to the GasPrice field.
func (o *XdaiTransactionConfig) SetGasPrice(v XdaiTransactionConfigFrom) {
	o.GasPrice = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *XdaiTransactionConfig) GetNonce() float32 {
	if o == nil || IsNil(o.Nonce) {
		var ret float32
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XdaiTransactionConfig) GetNonceOk() (*float32, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *XdaiTransactionConfig) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given float32 and assigns it to the Nonce field.
func (o *XdaiTransactionConfig) SetNonce(v float32) {
	o.Nonce = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *XdaiTransactionConfig) GetData() XdaiTransactionConfigTo {
	if o == nil || IsNil(o.Data) {
		var ret XdaiTransactionConfigTo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XdaiTransactionConfig) GetDataOk() (*XdaiTransactionConfigTo, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *XdaiTransactionConfig) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given XdaiTransactionConfigTo and assigns it to the Data field.
func (o *XdaiTransactionConfig) SetData(v XdaiTransactionConfigTo) {
	o.Data = &v
}

func (o XdaiTransactionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XdaiTransactionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RawTransaction) {
		toSerialize["rawTransaction"] = o.RawTransaction
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	if !IsNil(o.GasPrice) {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *XdaiTransactionConfig) UnmarshalJSON(bytes []byte) (err error) {
	varXdaiTransactionConfig := _XdaiTransactionConfig{}

	if err = json.Unmarshal(bytes, &varXdaiTransactionConfig); err == nil {
		*o = XdaiTransactionConfig(varXdaiTransactionConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "rawTransaction")
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		delete(additionalProperties, "value")
		delete(additionalProperties, "gas")
		delete(additionalProperties, "gasPrice")
		delete(additionalProperties, "nonce")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableXdaiTransactionConfig struct {
	value *XdaiTransactionConfig
	isSet bool
}

func (v NullableXdaiTransactionConfig) Get() *XdaiTransactionConfig {
	return v.value
}

func (v *NullableXdaiTransactionConfig) Set(val *XdaiTransactionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableXdaiTransactionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableXdaiTransactionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXdaiTransactionConfig(val *XdaiTransactionConfig) *NullableXdaiTransactionConfig {
	return &NullableXdaiTransactionConfig{value: val, isSet: true}
}

func (v NullableXdaiTransactionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXdaiTransactionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


