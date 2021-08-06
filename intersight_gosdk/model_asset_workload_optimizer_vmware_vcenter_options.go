/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-21T16:37:30Z.
 *
 * API version: 1.0.9-4403
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// AssetWorkloadOptimizerVmwareVcenterOptions Captures configuration specific to the VMware Vcenter target for the Workload Optimizer service.
type AssetWorkloadOptimizerVmwareVcenterOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// DatastoreBrowsingEnabled controls whether Workload Optimizer scans vCenter datastores to identify files which are not used and can be deleted to reclaim space and improve actual disk utilization. For example orphaned VMDK files.
	DatastoreBrowsingEnabled *bool `json:"DatastoreBrowsingEnabled,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _AssetWorkloadOptimizerVmwareVcenterOptions AssetWorkloadOptimizerVmwareVcenterOptions

// NewAssetWorkloadOptimizerVmwareVcenterOptions instantiates a new AssetWorkloadOptimizerVmwareVcenterOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerVmwareVcenterOptions(classId string, objectType string) *AssetWorkloadOptimizerVmwareVcenterOptions {
	this := AssetWorkloadOptimizerVmwareVcenterOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerVmwareVcenterOptionsWithDefaults instantiates a new AssetWorkloadOptimizerVmwareVcenterOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerVmwareVcenterOptionsWithDefaults() *AssetWorkloadOptimizerVmwareVcenterOptions {
	this := AssetWorkloadOptimizerVmwareVcenterOptions{}
	var classId string = "asset.WorkloadOptimizerVmwareVcenterOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerVmwareVcenterOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerVmwareVcenterOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerVmwareVcenterOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDatastoreBrowsingEnabled returns the DatastoreBrowsingEnabled field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetDatastoreBrowsingEnabled() bool {
	if o == nil || o.DatastoreBrowsingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DatastoreBrowsingEnabled
}

// GetDatastoreBrowsingEnabledOk returns a tuple with the DatastoreBrowsingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetDatastoreBrowsingEnabledOk() (*bool, bool) {
	if o == nil || o.DatastoreBrowsingEnabled == nil {
		return nil, false
	}
	return o.DatastoreBrowsingEnabled, true
}

// HasDatastoreBrowsingEnabled returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerVmwareVcenterOptions) HasDatastoreBrowsingEnabled() bool {
	if o != nil && o.DatastoreBrowsingEnabled != nil {
		return true
	}

	return false
}

// SetDatastoreBrowsingEnabled gets a reference to the given bool and assigns it to the DatastoreBrowsingEnabled field.
func (o *AssetWorkloadOptimizerVmwareVcenterOptions) SetDatastoreBrowsingEnabled(v bool) {
	o.DatastoreBrowsingEnabled = &v
}

func (o AssetWorkloadOptimizerVmwareVcenterOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DatastoreBrowsingEnabled != nil {
		toSerialize["DatastoreBrowsingEnabled"] = o.DatastoreBrowsingEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerVmwareVcenterOptions) UnmarshalJSON(bytes []byte) (err error) {
	type AssetWorkloadOptimizerVmwareVcenterOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// DatastoreBrowsingEnabled controls whether Workload Optimizer scans vCenter datastores to identify files which are not used and can be deleted to reclaim space and improve actual disk utilization. For example orphaned VMDK files.
		DatastoreBrowsingEnabled *bool `json:"DatastoreBrowsingEnabled,omitempty"`
	}

	varAssetWorkloadOptimizerVmwareVcenterOptionsWithoutEmbeddedStruct := AssetWorkloadOptimizerVmwareVcenterOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerVmwareVcenterOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetWorkloadOptimizerVmwareVcenterOptions := _AssetWorkloadOptimizerVmwareVcenterOptions{}
		varAssetWorkloadOptimizerVmwareVcenterOptions.ClassId = varAssetWorkloadOptimizerVmwareVcenterOptionsWithoutEmbeddedStruct.ClassId
		varAssetWorkloadOptimizerVmwareVcenterOptions.ObjectType = varAssetWorkloadOptimizerVmwareVcenterOptionsWithoutEmbeddedStruct.ObjectType
		varAssetWorkloadOptimizerVmwareVcenterOptions.DatastoreBrowsingEnabled = varAssetWorkloadOptimizerVmwareVcenterOptionsWithoutEmbeddedStruct.DatastoreBrowsingEnabled
		*o = AssetWorkloadOptimizerVmwareVcenterOptions(varAssetWorkloadOptimizerVmwareVcenterOptions)
	} else {
		return err
	}

	varAssetWorkloadOptimizerVmwareVcenterOptions := _AssetWorkloadOptimizerVmwareVcenterOptions{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerVmwareVcenterOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetWorkloadOptimizerVmwareVcenterOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DatastoreBrowsingEnabled")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerVmwareVcenterOptions struct {
	value *AssetWorkloadOptimizerVmwareVcenterOptions
	isSet bool
}

func (v NullableAssetWorkloadOptimizerVmwareVcenterOptions) Get() *AssetWorkloadOptimizerVmwareVcenterOptions {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerVmwareVcenterOptions) Set(val *AssetWorkloadOptimizerVmwareVcenterOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerVmwareVcenterOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerVmwareVcenterOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerVmwareVcenterOptions(val *AssetWorkloadOptimizerVmwareVcenterOptions) *NullableAssetWorkloadOptimizerVmwareVcenterOptions {
	return &NullableAssetWorkloadOptimizerVmwareVcenterOptions{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerVmwareVcenterOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerVmwareVcenterOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
