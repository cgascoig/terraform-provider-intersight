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

// TelemetryDruidQueryDataSource This is used for nested groupBys and is only currently supported for groupBys.
type TelemetryDruidQueryDataSource struct {
	// The type of data source.
	Type                 string                       `json:"type"`
	Query                TelemetryDruidGroupByRequest `json:"query"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidQueryDataSource TelemetryDruidQueryDataSource

// NewTelemetryDruidQueryDataSource instantiates a new TelemetryDruidQueryDataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidQueryDataSource(type_ string, query TelemetryDruidGroupByRequest) *TelemetryDruidQueryDataSource {
	this := TelemetryDruidQueryDataSource{}
	this.Type = type_
	this.Query = query
	return &this
}

// NewTelemetryDruidQueryDataSourceWithDefaults instantiates a new TelemetryDruidQueryDataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidQueryDataSourceWithDefaults() *TelemetryDruidQueryDataSource {
	this := TelemetryDruidQueryDataSource{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidQueryDataSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryDataSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidQueryDataSource) SetType(v string) {
	o.Type = v
}

// GetQuery returns the Query field value
func (o *TelemetryDruidQueryDataSource) GetQuery() TelemetryDruidGroupByRequest {
	if o == nil {
		var ret TelemetryDruidGroupByRequest
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryDataSource) GetQueryOk() (*TelemetryDruidGroupByRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *TelemetryDruidQueryDataSource) SetQuery(v TelemetryDruidGroupByRequest) {
	o.Query = v
}

func (o TelemetryDruidQueryDataSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidQueryDataSource) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidQueryDataSource := _TelemetryDruidQueryDataSource{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidQueryDataSource); err == nil {
		*o = TelemetryDruidQueryDataSource(varTelemetryDruidQueryDataSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "query")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidQueryDataSource struct {
	value *TelemetryDruidQueryDataSource
	isSet bool
}

func (v NullableTelemetryDruidQueryDataSource) Get() *TelemetryDruidQueryDataSource {
	return v.value
}

func (v *NullableTelemetryDruidQueryDataSource) Set(val *TelemetryDruidQueryDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidQueryDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidQueryDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidQueryDataSource(val *TelemetryDruidQueryDataSource) *NullableTelemetryDruidQueryDataSource {
	return &NullableTelemetryDruidQueryDataSource{value: val, isSet: true}
}

func (v NullableTelemetryDruidQueryDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidQueryDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
