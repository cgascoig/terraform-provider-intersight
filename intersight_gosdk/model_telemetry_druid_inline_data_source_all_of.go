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

// TelemetryDruidInlineDataSourceAllOf struct for TelemetryDruidInlineDataSourceAllOf
type TelemetryDruidInlineDataSourceAllOf struct {
	// the column names.
	ColumnNames []string `json:"columnNames"`
	// an array of rows.
	Rows                 [][]string `json:"rows"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidInlineDataSourceAllOf TelemetryDruidInlineDataSourceAllOf

// NewTelemetryDruidInlineDataSourceAllOf instantiates a new TelemetryDruidInlineDataSourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidInlineDataSourceAllOf(columnNames []string, rows [][]string) *TelemetryDruidInlineDataSourceAllOf {
	this := TelemetryDruidInlineDataSourceAllOf{}
	this.ColumnNames = columnNames
	this.Rows = rows
	return &this
}

// NewTelemetryDruidInlineDataSourceAllOfWithDefaults instantiates a new TelemetryDruidInlineDataSourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidInlineDataSourceAllOfWithDefaults() *TelemetryDruidInlineDataSourceAllOf {
	this := TelemetryDruidInlineDataSourceAllOf{}
	return &this
}

// GetColumnNames returns the ColumnNames field value
func (o *TelemetryDruidInlineDataSourceAllOf) GetColumnNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ColumnNames
}

// GetColumnNamesOk returns a tuple with the ColumnNames field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInlineDataSourceAllOf) GetColumnNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnNames, true
}

// SetColumnNames sets field value
func (o *TelemetryDruidInlineDataSourceAllOf) SetColumnNames(v []string) {
	o.ColumnNames = v
}

// GetRows returns the Rows field value
func (o *TelemetryDruidInlineDataSourceAllOf) GetRows() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInlineDataSourceAllOf) GetRowsOk() (*[][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rows, true
}

// SetRows sets field value
func (o *TelemetryDruidInlineDataSourceAllOf) SetRows(v [][]string) {
	o.Rows = v
}

func (o TelemetryDruidInlineDataSourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["columnNames"] = o.ColumnNames
	}
	if true {
		toSerialize["rows"] = o.Rows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidInlineDataSourceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidInlineDataSourceAllOf := _TelemetryDruidInlineDataSourceAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidInlineDataSourceAllOf); err == nil {
		*o = TelemetryDruidInlineDataSourceAllOf(varTelemetryDruidInlineDataSourceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "columnNames")
		delete(additionalProperties, "rows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidInlineDataSourceAllOf struct {
	value *TelemetryDruidInlineDataSourceAllOf
	isSet bool
}

func (v NullableTelemetryDruidInlineDataSourceAllOf) Get() *TelemetryDruidInlineDataSourceAllOf {
	return v.value
}

func (v *NullableTelemetryDruidInlineDataSourceAllOf) Set(val *TelemetryDruidInlineDataSourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidInlineDataSourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidInlineDataSourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidInlineDataSourceAllOf(val *TelemetryDruidInlineDataSourceAllOf) *NullableTelemetryDruidInlineDataSourceAllOf {
	return &NullableTelemetryDruidInlineDataSourceAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidInlineDataSourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidInlineDataSourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
