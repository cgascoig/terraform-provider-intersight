/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-05-25T18:18:54Z.
 *
 * API version: 1.0.9-4305
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidNumericTopNMetricSpec The simplest metric specification is a String value indicating the metric to sort topN results by.
type TelemetryDruidNumericTopNMetricSpec struct {
	// The dimension spec type.
	Type string `json:"type"`
	// The name of the metric to sort topN results.
	Metric               string `json:"metric"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidNumericTopNMetricSpec TelemetryDruidNumericTopNMetricSpec

// NewTelemetryDruidNumericTopNMetricSpec instantiates a new TelemetryDruidNumericTopNMetricSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidNumericTopNMetricSpec(type_ string, metric string) *TelemetryDruidNumericTopNMetricSpec {
	this := TelemetryDruidNumericTopNMetricSpec{}
	this.Type = type_
	this.Metric = metric
	return &this
}

// NewTelemetryDruidNumericTopNMetricSpecWithDefaults instantiates a new TelemetryDruidNumericTopNMetricSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidNumericTopNMetricSpecWithDefaults() *TelemetryDruidNumericTopNMetricSpec {
	this := TelemetryDruidNumericTopNMetricSpec{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidNumericTopNMetricSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidNumericTopNMetricSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidNumericTopNMetricSpec) SetType(v string) {
	o.Type = v
}

// GetMetric returns the Metric field value
func (o *TelemetryDruidNumericTopNMetricSpec) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidNumericTopNMetricSpec) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *TelemetryDruidNumericTopNMetricSpec) SetMetric(v string) {
	o.Metric = v
}

func (o TelemetryDruidNumericTopNMetricSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["metric"] = o.Metric
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidNumericTopNMetricSpec) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidNumericTopNMetricSpec := _TelemetryDruidNumericTopNMetricSpec{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidNumericTopNMetricSpec); err == nil {
		*o = TelemetryDruidNumericTopNMetricSpec(varTelemetryDruidNumericTopNMetricSpec)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "metric")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidNumericTopNMetricSpec struct {
	value *TelemetryDruidNumericTopNMetricSpec
	isSet bool
}

func (v NullableTelemetryDruidNumericTopNMetricSpec) Get() *TelemetryDruidNumericTopNMetricSpec {
	return v.value
}

func (v *NullableTelemetryDruidNumericTopNMetricSpec) Set(val *TelemetryDruidNumericTopNMetricSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidNumericTopNMetricSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidNumericTopNMetricSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidNumericTopNMetricSpec(val *TelemetryDruidNumericTopNMetricSpec) *NullableTelemetryDruidNumericTopNMetricSpec {
	return &NullableTelemetryDruidNumericTopNMetricSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidNumericTopNMetricSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidNumericTopNMetricSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
