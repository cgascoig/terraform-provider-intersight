/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4870
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidJoinDataSource Join datasources allow you to do a SQL-style join of two datasources. Stacking joins on top of each other allows you to join arbitrarily many datasources. Joins are implemented with a broadcast hash-join algorithm. This means that all tables other than the leftmost \"base\" table must fit in memory. It also means that the join condition must be an equality. This feature is intended mainly to allow joining regular Druid tables with lookup, inline, and query datasources.
type TelemetryDruidJoinDataSource struct {
	// The type of data source.
	Type string `json:"type"`
	// Left-hand datasource. Must be of type table, join, lookup, query, or inline. Placing another join as the left datasource allows you to join arbitrarily many datasources.
	Left string `json:"left"`
	// Right-hand datasource. Must be of type lookup, query, or inline.
	Right string `json:"right"`
	// String prefix that will be applied to all columns from the right-hand datasource, to prevent them from colliding with columns from the left-hand datasource. Can be any string, so long as it is nonempty and is not be a prefix of the string __time. Any columns from the left-hand side that start with your rightPrefix will be shadowed. It is up to you to provide a prefix that will not shadow any important columns from the left side.
	RightPrefix string `json:"rightPrefix"`
	// Expression that must be an equality where one side is an expression of the left-hand side, and the other side is a simple column reference to the right-hand side. The right-hand reference must be a simple column reference.
	Condition            string `json:"condition"`
	JoinType             string `json:"joinType"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidJoinDataSource TelemetryDruidJoinDataSource

// NewTelemetryDruidJoinDataSource instantiates a new TelemetryDruidJoinDataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidJoinDataSource(type_ string, left string, right string, rightPrefix string, condition string, joinType string) *TelemetryDruidJoinDataSource {
	this := TelemetryDruidJoinDataSource{}
	this.Type = type_
	this.Left = left
	this.Right = right
	this.RightPrefix = rightPrefix
	this.Condition = condition
	this.JoinType = joinType
	return &this
}

// NewTelemetryDruidJoinDataSourceWithDefaults instantiates a new TelemetryDruidJoinDataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidJoinDataSourceWithDefaults() *TelemetryDruidJoinDataSource {
	this := TelemetryDruidJoinDataSource{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidJoinDataSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidJoinDataSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidJoinDataSource) SetType(v string) {
	o.Type = v
}

// GetLeft returns the Left field value
func (o *TelemetryDruidJoinDataSource) GetLeft() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Left
}

// GetLeftOk returns a tuple with the Left field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidJoinDataSource) GetLeftOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Left, true
}

// SetLeft sets field value
func (o *TelemetryDruidJoinDataSource) SetLeft(v string) {
	o.Left = v
}

// GetRight returns the Right field value
func (o *TelemetryDruidJoinDataSource) GetRight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Right
}

// GetRightOk returns a tuple with the Right field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidJoinDataSource) GetRightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Right, true
}

// SetRight sets field value
func (o *TelemetryDruidJoinDataSource) SetRight(v string) {
	o.Right = v
}

// GetRightPrefix returns the RightPrefix field value
func (o *TelemetryDruidJoinDataSource) GetRightPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RightPrefix
}

// GetRightPrefixOk returns a tuple with the RightPrefix field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidJoinDataSource) GetRightPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RightPrefix, true
}

// SetRightPrefix sets field value
func (o *TelemetryDruidJoinDataSource) SetRightPrefix(v string) {
	o.RightPrefix = v
}

// GetCondition returns the Condition field value
func (o *TelemetryDruidJoinDataSource) GetCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidJoinDataSource) GetConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *TelemetryDruidJoinDataSource) SetCondition(v string) {
	o.Condition = v
}

// GetJoinType returns the JoinType field value
func (o *TelemetryDruidJoinDataSource) GetJoinType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinType
}

// GetJoinTypeOk returns a tuple with the JoinType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidJoinDataSource) GetJoinTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinType, true
}

// SetJoinType sets field value
func (o *TelemetryDruidJoinDataSource) SetJoinType(v string) {
	o.JoinType = v
}

func (o TelemetryDruidJoinDataSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["left"] = o.Left
	}
	if true {
		toSerialize["right"] = o.Right
	}
	if true {
		toSerialize["rightPrefix"] = o.RightPrefix
	}
	if true {
		toSerialize["condition"] = o.Condition
	}
	if true {
		toSerialize["joinType"] = o.JoinType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidJoinDataSource) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidJoinDataSource := _TelemetryDruidJoinDataSource{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidJoinDataSource); err == nil {
		*o = TelemetryDruidJoinDataSource(varTelemetryDruidJoinDataSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "left")
		delete(additionalProperties, "right")
		delete(additionalProperties, "rightPrefix")
		delete(additionalProperties, "condition")
		delete(additionalProperties, "joinType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidJoinDataSource struct {
	value *TelemetryDruidJoinDataSource
	isSet bool
}

func (v NullableTelemetryDruidJoinDataSource) Get() *TelemetryDruidJoinDataSource {
	return v.value
}

func (v *NullableTelemetryDruidJoinDataSource) Set(val *TelemetryDruidJoinDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidJoinDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidJoinDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidJoinDataSource(val *TelemetryDruidJoinDataSource) *NullableTelemetryDruidJoinDataSource {
	return &NullableTelemetryDruidJoinDataSource{value: val, isSet: true}
}

func (v NullableTelemetryDruidJoinDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidJoinDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
