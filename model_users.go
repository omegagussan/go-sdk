/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"encoding/json"
)

// Users struct for Users
type Users struct {
	Users *[]string `json:"users,omitempty"yaml:"users,omitempty"`
}

// NewUsers instantiates a new Users object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsers() *Users {
	this := Users{}
	return &this
}

// NewUsersWithDefaults instantiates a new Users object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersWithDefaults() *Users {
	this := Users{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Users) GetUsers() []string {
	if o == nil || o.Users == nil {
		var ret []string
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Users) GetUsersOk() (*[]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Users) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *Users) SetUsers(v []string) {
	o.Users = &v
}

func (o Users) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableUsers struct {
	value *Users
	isSet bool
}

func (v NullableUsers) Get() *Users {
	return v.value
}

func (v *NullableUsers) Set(val *Users) {
	v.value = val
	v.isSet = true
}

func (v NullableUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsers(val *Users) *NullableUsers {
	return &NullableUsers{value: val, isSet: true}
}

func (v NullableUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
