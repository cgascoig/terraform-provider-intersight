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

// TelemetryDruidDataSourceMetadataRequestAllOf struct for TelemetryDruidDataSourceMetadataRequestAllOf
type TelemetryDruidDataSourceMetadataRequestAllOf struct {
	DataSource           TelemetryDruidDataSource    `json:"dataSource"`
	Context              *TelemetryDruidQueryContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidDataSourceMetadataRequestAllOf TelemetryDruidDataSourceMetadataRequestAllOf

// NewTelemetryDruidDataSourceMetadataRequestAllOf instantiates a new TelemetryDruidDataSourceMetadataRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidDataSourceMetadataRequestAllOf(dataSource TelemetryDruidDataSource) *TelemetryDruidDataSourceMetadataRequestAllOf {
	this := TelemetryDruidDataSourceMetadataRequestAllOf{}
	this.DataSource = dataSource
	return &this
}

// NewTelemetryDruidDataSourceMetadataRequestAllOfWithDefaults instantiates a new TelemetryDruidDataSourceMetadataRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidDataSourceMetadataRequestAllOfWithDefaults() *TelemetryDruidDataSourceMetadataRequestAllOf {
	this := TelemetryDruidDataSourceMetadataRequestAllOf{}
	return &this
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidDataSourceMetadataRequestAllOf) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDataSourceMetadataRequestAllOf) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidDataSourceMetadataRequestAllOf) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidDataSourceMetadataRequestAllOf) GetContext() TelemetryDruidQueryContext {
	if o == nil || o.Context == nil {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDataSourceMetadataRequestAllOf) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidDataSourceMetadataRequestAllOf) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidDataSourceMetadataRequestAllOf) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidDataSourceMetadataRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataSource"] = o.DataSource
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidDataSourceMetadataRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidDataSourceMetadataRequestAllOf := _TelemetryDruidDataSourceMetadataRequestAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidDataSourceMetadataRequestAllOf); err == nil {
		*o = TelemetryDruidDataSourceMetadataRequestAllOf(varTelemetryDruidDataSourceMetadataRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidDataSourceMetadataRequestAllOf struct {
	value *TelemetryDruidDataSourceMetadataRequestAllOf
	isSet bool
}

func (v NullableTelemetryDruidDataSourceMetadataRequestAllOf) Get() *TelemetryDruidDataSourceMetadataRequestAllOf {
	return v.value
}

func (v *NullableTelemetryDruidDataSourceMetadataRequestAllOf) Set(val *TelemetryDruidDataSourceMetadataRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidDataSourceMetadataRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidDataSourceMetadataRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidDataSourceMetadataRequestAllOf(val *TelemetryDruidDataSourceMetadataRequestAllOf) *NullableTelemetryDruidDataSourceMetadataRequestAllOf {
	return &NullableTelemetryDruidDataSourceMetadataRequestAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidDataSourceMetadataRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidDataSourceMetadataRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
