/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// NiatelemetryJobDetailAllOf Definition of the list of properties defined in 'niatelemetry.JobDetail', excluding properties defined in parent classes.
type NiatelemetryJobDetailAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the id of the job.
	JobId *int64 `json:"JobId,omitempty"`
	// Returns the status of the job.
	UpgStatus            *string `json:"UpgStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryJobDetailAllOf NiatelemetryJobDetailAllOf

// NewNiatelemetryJobDetailAllOf instantiates a new NiatelemetryJobDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryJobDetailAllOf(classId string, objectType string) *NiatelemetryJobDetailAllOf {
	this := NiatelemetryJobDetailAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryJobDetailAllOfWithDefaults instantiates a new NiatelemetryJobDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryJobDetailAllOfWithDefaults() *NiatelemetryJobDetailAllOf {
	this := NiatelemetryJobDetailAllOf{}
	var classId string = "niatelemetry.JobDetail"
	this.ClassId = classId
	var objectType string = "niatelemetry.JobDetail"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryJobDetailAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryJobDetailAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryJobDetailAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryJobDetailAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryJobDetailAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryJobDetailAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *NiatelemetryJobDetailAllOf) GetJobId() int64 {
	if o == nil || o.JobId == nil {
		var ret int64
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryJobDetailAllOf) GetJobIdOk() (*int64, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *NiatelemetryJobDetailAllOf) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int64 and assigns it to the JobId field.
func (o *NiatelemetryJobDetailAllOf) SetJobId(v int64) {
	o.JobId = &v
}

// GetUpgStatus returns the UpgStatus field value if set, zero value otherwise.
func (o *NiatelemetryJobDetailAllOf) GetUpgStatus() string {
	if o == nil || o.UpgStatus == nil {
		var ret string
		return ret
	}
	return *o.UpgStatus
}

// GetUpgStatusOk returns a tuple with the UpgStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryJobDetailAllOf) GetUpgStatusOk() (*string, bool) {
	if o == nil || o.UpgStatus == nil {
		return nil, false
	}
	return o.UpgStatus, true
}

// HasUpgStatus returns a boolean if a field has been set.
func (o *NiatelemetryJobDetailAllOf) HasUpgStatus() bool {
	if o != nil && o.UpgStatus != nil {
		return true
	}

	return false
}

// SetUpgStatus gets a reference to the given string and assigns it to the UpgStatus field.
func (o *NiatelemetryJobDetailAllOf) SetUpgStatus(v string) {
	o.UpgStatus = &v
}

func (o NiatelemetryJobDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.JobId != nil {
		toSerialize["JobId"] = o.JobId
	}
	if o.UpgStatus != nil {
		toSerialize["UpgStatus"] = o.UpgStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryJobDetailAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryJobDetailAllOf := _NiatelemetryJobDetailAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryJobDetailAllOf); err == nil {
		*o = NiatelemetryJobDetailAllOf(varNiatelemetryJobDetailAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "JobId")
		delete(additionalProperties, "UpgStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryJobDetailAllOf struct {
	value *NiatelemetryJobDetailAllOf
	isSet bool
}

func (v NullableNiatelemetryJobDetailAllOf) Get() *NiatelemetryJobDetailAllOf {
	return v.value
}

func (v *NullableNiatelemetryJobDetailAllOf) Set(val *NiatelemetryJobDetailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryJobDetailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryJobDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryJobDetailAllOf(val *NiatelemetryJobDetailAllOf) *NullableNiatelemetryJobDetailAllOf {
	return &NullableNiatelemetryJobDetailAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryJobDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryJobDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
