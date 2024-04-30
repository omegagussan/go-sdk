/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://openfga.dev/community
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"bytes"

	"encoding/json"
)

// UsersetUser struct for UsersetUser
type UsersetUser struct {
	Type     string `json:"type"yaml:"type"`
	Id       string `json:"id"yaml:"id"`
	Relation string `json:"relation"yaml:"relation"`
}

// NewUsersetUser instantiates a new UsersetUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersetUser(type_ string, id string, relation string) *UsersetUser {
	this := UsersetUser{}
	this.Type = type_
	this.Id = id
	this.Relation = relation
	return &this
}

// NewUsersetUserWithDefaults instantiates a new UsersetUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersetUserWithDefaults() *UsersetUser {
	this := UsersetUser{}
	return &this
}

// GetType returns the Type field value
func (o *UsersetUser) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UsersetUser) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UsersetUser) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *UsersetUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UsersetUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UsersetUser) SetId(v string) {
	o.Id = v
}

// GetRelation returns the Relation field value
func (o *UsersetUser) GetRelation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Relation
}

// GetRelationOk returns a tuple with the Relation field value
// and a boolean to check if the value has been set.
func (o *UsersetUser) GetRelationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relation, true
}

// SetRelation sets field value
func (o *UsersetUser) SetRelation(v string) {
	o.Relation = v
}

func (o UsersetUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["relation"] = o.Relation
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableUsersetUser struct {
	value *UsersetUser
	isSet bool
}

func (v NullableUsersetUser) Get() *UsersetUser {
	return v.value
}

func (v *NullableUsersetUser) Set(val *UsersetUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersetUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersetUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersetUser(val *UsersetUser) *NullableUsersetUser {
	return &NullableUsersetUser{value: val, isSet: true}
}

func (v NullableUsersetUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersetUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}