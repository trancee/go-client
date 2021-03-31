/*
 * NFT Storage API
 *
 * # API Reference 
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nftstorage

import (
	"encoding/json"
)

// GetResponse struct for GetResponse
type GetResponse struct {
	Ok *bool `json:"ok,omitempty"`
	Value *NFT `json:"value,omitempty"`
}

// NewGetResponse instantiates a new GetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetResponse() *GetResponse {
	this := GetResponse{}
	var ok bool = true
	this.Ok = &ok
	return &this
}

// NewGetResponseWithDefaults instantiates a new GetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetResponseWithDefaults() *GetResponse {
	this := GetResponse{}
	var ok bool = true
	this.Ok = &ok
	return &this
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *GetResponse) GetOk() bool {
	if o == nil || o.Ok == nil {
		var ret bool
		return ret
	}
	return *o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResponse) GetOkOk() (*bool, bool) {
	if o == nil || o.Ok == nil {
		return nil, false
	}
	return o.Ok, true
}

// HasOk returns a boolean if a field has been set.
func (o *GetResponse) HasOk() bool {
	if o != nil && o.Ok != nil {
		return true
	}

	return false
}

// SetOk gets a reference to the given bool and assigns it to the Ok field.
func (o *GetResponse) SetOk(v bool) {
	o.Ok = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetResponse) GetValue() NFT {
	if o == nil || o.Value == nil {
		var ret NFT
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResponse) GetValueOk() (*NFT, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetResponse) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given NFT and assigns it to the Value field.
func (o *GetResponse) SetValue(v NFT) {
	o.Value = &v
}

func (o GetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ok != nil {
		toSerialize["ok"] = o.Ok
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableGetResponse struct {
	value *GetResponse
	isSet bool
}

func (v NullableGetResponse) Get() *GetResponse {
	return v.value
}

func (v *NullableGetResponse) Set(val *GetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetResponse(val *GetResponse) *NullableGetResponse {
	return &NullableGetResponse{value: val, isSet: true}
}

func (v NullableGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

