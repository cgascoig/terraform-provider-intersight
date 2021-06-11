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
	"reflect"
	"strings"
)

// StorageBaseReplicationBlackout The range of time during which replication will be suspended. System disables replication during this interval.
type StorageBaseReplicationBlackout struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The end time of day for replication blackout window. Example: 17:00:01 which is 17 hours, 0 minutes, 1 seconds.
	End *string `json:"End,omitempty"`
	// The start time of day when replication blackout is active. When replication blackout is active, the storage array temporarily disables replication. Example: 15:04:03.123 which is 15 hours, 4 minutes, 3 seconds and 123 milliseconds. The fractional seconds are written using the standard decimal notation which can be used for setting milliseconds and microseconds.
	Start                *string `json:"Start,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseReplicationBlackout StorageBaseReplicationBlackout

// NewStorageBaseReplicationBlackout instantiates a new StorageBaseReplicationBlackout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseReplicationBlackout(classId string, objectType string) *StorageBaseReplicationBlackout {
	this := StorageBaseReplicationBlackout{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseReplicationBlackoutWithDefaults instantiates a new StorageBaseReplicationBlackout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseReplicationBlackoutWithDefaults() *StorageBaseReplicationBlackout {
	this := StorageBaseReplicationBlackout{}
	var classId string = "storage.PureReplicationBlackout"
	this.ClassId = classId
	var objectType string = "storage.PureReplicationBlackout"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseReplicationBlackout) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationBlackout) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseReplicationBlackout) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseReplicationBlackout) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationBlackout) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseReplicationBlackout) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *StorageBaseReplicationBlackout) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationBlackout) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *StorageBaseReplicationBlackout) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *StorageBaseReplicationBlackout) SetEnd(v string) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *StorageBaseReplicationBlackout) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationBlackout) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *StorageBaseReplicationBlackout) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *StorageBaseReplicationBlackout) SetStart(v string) {
	o.Start = &v
}

func (o StorageBaseReplicationBlackout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.End != nil {
		toSerialize["End"] = o.End
	}
	if o.Start != nil {
		toSerialize["Start"] = o.Start
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseReplicationBlackout) UnmarshalJSON(bytes []byte) (err error) {
	type StorageBaseReplicationBlackoutWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The end time of day for replication blackout window. Example: 17:00:01 which is 17 hours, 0 minutes, 1 seconds.
		End *string `json:"End,omitempty"`
		// The start time of day when replication blackout is active. When replication blackout is active, the storage array temporarily disables replication. Example: 15:04:03.123 which is 15 hours, 4 minutes, 3 seconds and 123 milliseconds. The fractional seconds are written using the standard decimal notation which can be used for setting milliseconds and microseconds.
		Start *string `json:"Start,omitempty"`
	}

	varStorageBaseReplicationBlackoutWithoutEmbeddedStruct := StorageBaseReplicationBlackoutWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageBaseReplicationBlackoutWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseReplicationBlackout := _StorageBaseReplicationBlackout{}
		varStorageBaseReplicationBlackout.ClassId = varStorageBaseReplicationBlackoutWithoutEmbeddedStruct.ClassId
		varStorageBaseReplicationBlackout.ObjectType = varStorageBaseReplicationBlackoutWithoutEmbeddedStruct.ObjectType
		varStorageBaseReplicationBlackout.End = varStorageBaseReplicationBlackoutWithoutEmbeddedStruct.End
		varStorageBaseReplicationBlackout.Start = varStorageBaseReplicationBlackoutWithoutEmbeddedStruct.Start
		*o = StorageBaseReplicationBlackout(varStorageBaseReplicationBlackout)
	} else {
		return err
	}

	varStorageBaseReplicationBlackout := _StorageBaseReplicationBlackout{}

	err = json.Unmarshal(bytes, &varStorageBaseReplicationBlackout)
	if err == nil {
		o.MoBaseComplexType = varStorageBaseReplicationBlackout.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "End")
		delete(additionalProperties, "Start")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableStorageBaseReplicationBlackout struct {
	value *StorageBaseReplicationBlackout
	isSet bool
}

func (v NullableStorageBaseReplicationBlackout) Get() *StorageBaseReplicationBlackout {
	return v.value
}

func (v *NullableStorageBaseReplicationBlackout) Set(val *StorageBaseReplicationBlackout) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseReplicationBlackout) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseReplicationBlackout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseReplicationBlackout(val *StorageBaseReplicationBlackout) *NullableStorageBaseReplicationBlackout {
	return &NullableStorageBaseReplicationBlackout{value: val, isSet: true}
}

func (v NullableStorageBaseReplicationBlackout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseReplicationBlackout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
