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

// Node struct for Node
type Node struct {
	Name         *string                `json:"name,omitempty"yaml:"name,omitempty"`
	Leaf         *Leaf                  `json:"leaf,omitempty"yaml:"leaf,omitempty"`
	Difference   *UsersetTreeDifference `json:"difference,omitempty"yaml:"difference,omitempty"`
	Union        *Nodes                 `json:"union,omitempty"yaml:"union,omitempty"`
	Intersection *Nodes                 `json:"intersection,omitempty"yaml:"intersection,omitempty"`
}

// NewNode instantiates a new Node object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNode() *Node {
	this := Node{}
	return &this
}

// NewNodeWithDefaults instantiates a new Node object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeWithDefaults() *Node {
	this := Node{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Node) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Node) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Node) SetName(v string) {
	o.Name = &v
}

// GetLeaf returns the Leaf field value if set, zero value otherwise.
func (o *Node) GetLeaf() Leaf {
	if o == nil || o.Leaf == nil {
		var ret Leaf
		return ret
	}
	return *o.Leaf
}

// GetLeafOk returns a tuple with the Leaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetLeafOk() (*Leaf, bool) {
	if o == nil || o.Leaf == nil {
		return nil, false
	}
	return o.Leaf, true
}

// HasLeaf returns a boolean if a field has been set.
func (o *Node) HasLeaf() bool {
	if o != nil && o.Leaf != nil {
		return true
	}

	return false
}

// SetLeaf gets a reference to the given Leaf and assigns it to the Leaf field.
func (o *Node) SetLeaf(v Leaf) {
	o.Leaf = &v
}

// GetDifference returns the Difference field value if set, zero value otherwise.
func (o *Node) GetDifference() UsersetTreeDifference {
	if o == nil || o.Difference == nil {
		var ret UsersetTreeDifference
		return ret
	}
	return *o.Difference
}

// GetDifferenceOk returns a tuple with the Difference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetDifferenceOk() (*UsersetTreeDifference, bool) {
	if o == nil || o.Difference == nil {
		return nil, false
	}
	return o.Difference, true
}

// HasDifference returns a boolean if a field has been set.
func (o *Node) HasDifference() bool {
	if o != nil && o.Difference != nil {
		return true
	}

	return false
}

// SetDifference gets a reference to the given UsersetTreeDifference and assigns it to the Difference field.
func (o *Node) SetDifference(v UsersetTreeDifference) {
	o.Difference = &v
}

// GetUnion returns the Union field value if set, zero value otherwise.
func (o *Node) GetUnion() Nodes {
	if o == nil || o.Union == nil {
		var ret Nodes
		return ret
	}
	return *o.Union
}

// GetUnionOk returns a tuple with the Union field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetUnionOk() (*Nodes, bool) {
	if o == nil || o.Union == nil {
		return nil, false
	}
	return o.Union, true
}

// HasUnion returns a boolean if a field has been set.
func (o *Node) HasUnion() bool {
	if o != nil && o.Union != nil {
		return true
	}

	return false
}

// SetUnion gets a reference to the given Nodes and assigns it to the Union field.
func (o *Node) SetUnion(v Nodes) {
	o.Union = &v
}

// GetIntersection returns the Intersection field value if set, zero value otherwise.
func (o *Node) GetIntersection() Nodes {
	if o == nil || o.Intersection == nil {
		var ret Nodes
		return ret
	}
	return *o.Intersection
}

// GetIntersectionOk returns a tuple with the Intersection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetIntersectionOk() (*Nodes, bool) {
	if o == nil || o.Intersection == nil {
		return nil, false
	}
	return o.Intersection, true
}

// HasIntersection returns a boolean if a field has been set.
func (o *Node) HasIntersection() bool {
	if o != nil && o.Intersection != nil {
		return true
	}

	return false
}

// SetIntersection gets a reference to the given Nodes and assigns it to the Intersection field.
func (o *Node) SetIntersection(v Nodes) {
	o.Intersection = &v
}

func (o Node) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Leaf != nil {
		toSerialize["leaf"] = o.Leaf
	}
	if o.Difference != nil {
		toSerialize["difference"] = o.Difference
	}
	if o.Union != nil {
		toSerialize["union"] = o.Union
	}
	if o.Intersection != nil {
		toSerialize["intersection"] = o.Intersection
	}
	return json.Marshal(toSerialize)
}

type NullableNode struct {
	value *Node
	isSet bool
}

func (v NullableNode) Get() *Node {
	return v.value
}

func (v *NullableNode) Set(val *Node) {
	v.value = val
	v.isSet = true
}

func (v NullableNode) IsSet() bool {
	return v.isSet
}

func (v *NullableNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNode(val *Node) *NullableNode {
	return &NullableNode{value: val, isSet: true}
}

func (v NullableNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
