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

// NiatelemetryNvePacketCounters Stores the packet count per switch.
type NiatelemetryNvePacketCounters struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Return mcast in packet count.
	McastInpkts *int64 `json:"McastInpkts,omitempty"`
	// Return mcast outbytes count.
	McastOutbytes *int64 `json:"McastOutbytes,omitempty"`
	// Return ucast in packet count.
	UcastInpkts *int64 `json:"UcastInpkts,omitempty"`
	// Return ucast out packet count.
	UcastOutpkts         *int64 `json:"UcastOutpkts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNvePacketCounters NiatelemetryNvePacketCounters

// NewNiatelemetryNvePacketCounters instantiates a new NiatelemetryNvePacketCounters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNvePacketCounters(classId string, objectType string) *NiatelemetryNvePacketCounters {
	this := NiatelemetryNvePacketCounters{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNvePacketCountersWithDefaults instantiates a new NiatelemetryNvePacketCounters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNvePacketCountersWithDefaults() *NiatelemetryNvePacketCounters {
	this := NiatelemetryNvePacketCounters{}
	var classId string = "niatelemetry.NvePacketCounters"
	this.ClassId = classId
	var objectType string = "niatelemetry.NvePacketCounters"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNvePacketCounters) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNvePacketCounters) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNvePacketCounters) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNvePacketCounters) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNvePacketCounters) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNvePacketCounters) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMcastInpkts returns the McastInpkts field value if set, zero value otherwise.
func (o *NiatelemetryNvePacketCounters) GetMcastInpkts() int64 {
	if o == nil || o.McastInpkts == nil {
		var ret int64
		return ret
	}
	return *o.McastInpkts
}

// GetMcastInpktsOk returns a tuple with the McastInpkts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNvePacketCounters) GetMcastInpktsOk() (*int64, bool) {
	if o == nil || o.McastInpkts == nil {
		return nil, false
	}
	return o.McastInpkts, true
}

// HasMcastInpkts returns a boolean if a field has been set.
func (o *NiatelemetryNvePacketCounters) HasMcastInpkts() bool {
	if o != nil && o.McastInpkts != nil {
		return true
	}

	return false
}

// SetMcastInpkts gets a reference to the given int64 and assigns it to the McastInpkts field.
func (o *NiatelemetryNvePacketCounters) SetMcastInpkts(v int64) {
	o.McastInpkts = &v
}

// GetMcastOutbytes returns the McastOutbytes field value if set, zero value otherwise.
func (o *NiatelemetryNvePacketCounters) GetMcastOutbytes() int64 {
	if o == nil || o.McastOutbytes == nil {
		var ret int64
		return ret
	}
	return *o.McastOutbytes
}

// GetMcastOutbytesOk returns a tuple with the McastOutbytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNvePacketCounters) GetMcastOutbytesOk() (*int64, bool) {
	if o == nil || o.McastOutbytes == nil {
		return nil, false
	}
	return o.McastOutbytes, true
}

// HasMcastOutbytes returns a boolean if a field has been set.
func (o *NiatelemetryNvePacketCounters) HasMcastOutbytes() bool {
	if o != nil && o.McastOutbytes != nil {
		return true
	}

	return false
}

// SetMcastOutbytes gets a reference to the given int64 and assigns it to the McastOutbytes field.
func (o *NiatelemetryNvePacketCounters) SetMcastOutbytes(v int64) {
	o.McastOutbytes = &v
}

// GetUcastInpkts returns the UcastInpkts field value if set, zero value otherwise.
func (o *NiatelemetryNvePacketCounters) GetUcastInpkts() int64 {
	if o == nil || o.UcastInpkts == nil {
		var ret int64
		return ret
	}
	return *o.UcastInpkts
}

// GetUcastInpktsOk returns a tuple with the UcastInpkts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNvePacketCounters) GetUcastInpktsOk() (*int64, bool) {
	if o == nil || o.UcastInpkts == nil {
		return nil, false
	}
	return o.UcastInpkts, true
}

// HasUcastInpkts returns a boolean if a field has been set.
func (o *NiatelemetryNvePacketCounters) HasUcastInpkts() bool {
	if o != nil && o.UcastInpkts != nil {
		return true
	}

	return false
}

// SetUcastInpkts gets a reference to the given int64 and assigns it to the UcastInpkts field.
func (o *NiatelemetryNvePacketCounters) SetUcastInpkts(v int64) {
	o.UcastInpkts = &v
}

// GetUcastOutpkts returns the UcastOutpkts field value if set, zero value otherwise.
func (o *NiatelemetryNvePacketCounters) GetUcastOutpkts() int64 {
	if o == nil || o.UcastOutpkts == nil {
		var ret int64
		return ret
	}
	return *o.UcastOutpkts
}

// GetUcastOutpktsOk returns a tuple with the UcastOutpkts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNvePacketCounters) GetUcastOutpktsOk() (*int64, bool) {
	if o == nil || o.UcastOutpkts == nil {
		return nil, false
	}
	return o.UcastOutpkts, true
}

// HasUcastOutpkts returns a boolean if a field has been set.
func (o *NiatelemetryNvePacketCounters) HasUcastOutpkts() bool {
	if o != nil && o.UcastOutpkts != nil {
		return true
	}

	return false
}

// SetUcastOutpkts gets a reference to the given int64 and assigns it to the UcastOutpkts field.
func (o *NiatelemetryNvePacketCounters) SetUcastOutpkts(v int64) {
	o.UcastOutpkts = &v
}

func (o NiatelemetryNvePacketCounters) MarshalJSON() ([]byte, error) {
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
	if o.McastInpkts != nil {
		toSerialize["McastInpkts"] = o.McastInpkts
	}
	if o.McastOutbytes != nil {
		toSerialize["McastOutbytes"] = o.McastOutbytes
	}
	if o.UcastInpkts != nil {
		toSerialize["UcastInpkts"] = o.UcastInpkts
	}
	if o.UcastOutpkts != nil {
		toSerialize["UcastOutpkts"] = o.UcastOutpkts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNvePacketCounters) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryNvePacketCountersWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Return mcast in packet count.
		McastInpkts *int64 `json:"McastInpkts,omitempty"`
		// Return mcast outbytes count.
		McastOutbytes *int64 `json:"McastOutbytes,omitempty"`
		// Return ucast in packet count.
		UcastInpkts *int64 `json:"UcastInpkts,omitempty"`
		// Return ucast out packet count.
		UcastOutpkts *int64 `json:"UcastOutpkts,omitempty"`
	}

	varNiatelemetryNvePacketCountersWithoutEmbeddedStruct := NiatelemetryNvePacketCountersWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryNvePacketCountersWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNvePacketCounters := _NiatelemetryNvePacketCounters{}
		varNiatelemetryNvePacketCounters.ClassId = varNiatelemetryNvePacketCountersWithoutEmbeddedStruct.ClassId
		varNiatelemetryNvePacketCounters.ObjectType = varNiatelemetryNvePacketCountersWithoutEmbeddedStruct.ObjectType
		varNiatelemetryNvePacketCounters.McastInpkts = varNiatelemetryNvePacketCountersWithoutEmbeddedStruct.McastInpkts
		varNiatelemetryNvePacketCounters.McastOutbytes = varNiatelemetryNvePacketCountersWithoutEmbeddedStruct.McastOutbytes
		varNiatelemetryNvePacketCounters.UcastInpkts = varNiatelemetryNvePacketCountersWithoutEmbeddedStruct.UcastInpkts
		varNiatelemetryNvePacketCounters.UcastOutpkts = varNiatelemetryNvePacketCountersWithoutEmbeddedStruct.UcastOutpkts
		*o = NiatelemetryNvePacketCounters(varNiatelemetryNvePacketCounters)
	} else {
		return err
	}

	varNiatelemetryNvePacketCounters := _NiatelemetryNvePacketCounters{}

	err = json.Unmarshal(bytes, &varNiatelemetryNvePacketCounters)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryNvePacketCounters.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "McastInpkts")
		delete(additionalProperties, "McastOutbytes")
		delete(additionalProperties, "UcastInpkts")
		delete(additionalProperties, "UcastOutpkts")

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

type NullableNiatelemetryNvePacketCounters struct {
	value *NiatelemetryNvePacketCounters
	isSet bool
}

func (v NullableNiatelemetryNvePacketCounters) Get() *NiatelemetryNvePacketCounters {
	return v.value
}

func (v *NullableNiatelemetryNvePacketCounters) Set(val *NiatelemetryNvePacketCounters) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNvePacketCounters) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNvePacketCounters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNvePacketCounters(val *NiatelemetryNvePacketCounters) *NullableNiatelemetryNvePacketCounters {
	return &NullableNiatelemetryNvePacketCounters{value: val, isSet: true}
}

func (v NullableNiatelemetryNvePacketCounters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNvePacketCounters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
