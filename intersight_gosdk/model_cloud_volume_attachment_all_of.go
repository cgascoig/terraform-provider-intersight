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
	"time"
)

// CloudVolumeAttachmentAllOf Definition of the list of properties defined in 'cloud.VolumeAttachment', excluding properties defined in parent classes.
type CloudVolumeAttachmentAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time stamp when the attachment of volume to virtual machine is initiated.
	AttachedTime *time.Time `json:"AttachedTime,omitempty"`
	// If true, volume is deleted on virtual machine termination.
	AutoDelete *bool `json:"AutoDelete,omitempty"`
	// The time stamp when the detached of volume to virtual machine is initiated.
	DetachedTime *time.Time `json:"DetachedTime,omitempty"`
	// The device name.For example, /dev/sdh or xvdh in case of AWS cloud.
	DeviceName *string `json:"DeviceName,omitempty"`
	// The internally generated identity of this volume.
	Identity *string `json:"Identity,omitempty"`
	// The position of the volume attachment in the virtual machine.
	Index *int64 `json:"Index,omitempty"`
	// If set to true, then it is the root volume.
	IsRoot               *bool `json:"IsRoot,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudVolumeAttachmentAllOf CloudVolumeAttachmentAllOf

// NewCloudVolumeAttachmentAllOf instantiates a new CloudVolumeAttachmentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudVolumeAttachmentAllOf(classId string, objectType string) *CloudVolumeAttachmentAllOf {
	this := CloudVolumeAttachmentAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudVolumeAttachmentAllOfWithDefaults instantiates a new CloudVolumeAttachmentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudVolumeAttachmentAllOfWithDefaults() *CloudVolumeAttachmentAllOf {
	this := CloudVolumeAttachmentAllOf{}
	var classId string = "cloud.VolumeAttachment"
	this.ClassId = classId
	var objectType string = "cloud.VolumeAttachment"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudVolumeAttachmentAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachmentAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudVolumeAttachmentAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudVolumeAttachmentAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachmentAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudVolumeAttachmentAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAttachedTime returns the AttachedTime field value if set, zero value otherwise.
func (o *CloudVolumeAttachmentAllOf) GetAttachedTime() time.Time {
	if o == nil || o.AttachedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AttachedTime
}

// GetAttachedTimeOk returns a tuple with the AttachedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachmentAllOf) GetAttachedTimeOk() (*time.Time, bool) {
	if o == nil || o.AttachedTime == nil {
		return nil, false
	}
	return o.AttachedTime, true
}

// HasAttachedTime returns a boolean if a field has been set.
func (o *CloudVolumeAttachmentAllOf) HasAttachedTime() bool {
	if o != nil && o.AttachedTime != nil {
		return true
	}

	return false
}

// SetAttachedTime gets a reference to the given time.Time and assigns it to the AttachedTime field.
func (o *CloudVolumeAttachmentAllOf) SetAttachedTime(v time.Time) {
	o.AttachedTime = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *CloudVolumeAttachmentAllOf) GetAutoDelete() bool {
	if o == nil || o.AutoDelete == nil {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachmentAllOf) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || o.AutoDelete == nil {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *CloudVolumeAttachmentAllOf) HasAutoDelete() bool {
	if o != nil && o.AutoDelete != nil {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *CloudVolumeAttachmentAllOf) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetDetachedTime returns the DetachedTime field value if set, zero value otherwise.
func (o *CloudVolumeAttachmentAllOf) GetDetachedTime() time.Time {
	if o == nil || o.DetachedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DetachedTime
}

// GetDetachedTimeOk returns a tuple with the DetachedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachmentAllOf) GetDetachedTimeOk() (*time.Time, bool) {
	if o == nil || o.DetachedTime == nil {
		return nil, false
	}
	return o.DetachedTime, true
}

// HasDetachedTime returns a boolean if a field has been set.
func (o *CloudVolumeAttachmentAllOf) HasDetachedTime() bool {
	if o != nil && o.DetachedTime != nil {
		return true
	}

	return false
}

// SetDetachedTime gets a reference to the given time.Time and assigns it to the DetachedTime field.
func (o *CloudVolumeAttachmentAllOf) SetDetachedTime(v time.Time) {
	o.DetachedTime = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *CloudVolumeAttachmentAllOf) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachmentAllOf) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *CloudVolumeAttachmentAllOf) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *CloudVolumeAttachmentAllOf) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudVolumeAttachmentAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachmentAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudVolumeAttachmentAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudVolumeAttachmentAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *CloudVolumeAttachmentAllOf) GetIndex() int64 {
	if o == nil || o.Index == nil {
		var ret int64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachmentAllOf) GetIndexOk() (*int64, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *CloudVolumeAttachmentAllOf) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int64 and assigns it to the Index field.
func (o *CloudVolumeAttachmentAllOf) SetIndex(v int64) {
	o.Index = &v
}

// GetIsRoot returns the IsRoot field value if set, zero value otherwise.
func (o *CloudVolumeAttachmentAllOf) GetIsRoot() bool {
	if o == nil || o.IsRoot == nil {
		var ret bool
		return ret
	}
	return *o.IsRoot
}

// GetIsRootOk returns a tuple with the IsRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachmentAllOf) GetIsRootOk() (*bool, bool) {
	if o == nil || o.IsRoot == nil {
		return nil, false
	}
	return o.IsRoot, true
}

// HasIsRoot returns a boolean if a field has been set.
func (o *CloudVolumeAttachmentAllOf) HasIsRoot() bool {
	if o != nil && o.IsRoot != nil {
		return true
	}

	return false
}

// SetIsRoot gets a reference to the given bool and assigns it to the IsRoot field.
func (o *CloudVolumeAttachmentAllOf) SetIsRoot(v bool) {
	o.IsRoot = &v
}

func (o CloudVolumeAttachmentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AttachedTime != nil {
		toSerialize["AttachedTime"] = o.AttachedTime
	}
	if o.AutoDelete != nil {
		toSerialize["AutoDelete"] = o.AutoDelete
	}
	if o.DetachedTime != nil {
		toSerialize["DetachedTime"] = o.DetachedTime
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Index != nil {
		toSerialize["Index"] = o.Index
	}
	if o.IsRoot != nil {
		toSerialize["IsRoot"] = o.IsRoot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudVolumeAttachmentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudVolumeAttachmentAllOf := _CloudVolumeAttachmentAllOf{}

	if err = json.Unmarshal(bytes, &varCloudVolumeAttachmentAllOf); err == nil {
		*o = CloudVolumeAttachmentAllOf(varCloudVolumeAttachmentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AttachedTime")
		delete(additionalProperties, "AutoDelete")
		delete(additionalProperties, "DetachedTime")
		delete(additionalProperties, "DeviceName")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Index")
		delete(additionalProperties, "IsRoot")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudVolumeAttachmentAllOf struct {
	value *CloudVolumeAttachmentAllOf
	isSet bool
}

func (v NullableCloudVolumeAttachmentAllOf) Get() *CloudVolumeAttachmentAllOf {
	return v.value
}

func (v *NullableCloudVolumeAttachmentAllOf) Set(val *CloudVolumeAttachmentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudVolumeAttachmentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudVolumeAttachmentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudVolumeAttachmentAllOf(val *CloudVolumeAttachmentAllOf) *NullableCloudVolumeAttachmentAllOf {
	return &NullableCloudVolumeAttachmentAllOf{value: val, isSet: true}
}

func (v NullableCloudVolumeAttachmentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudVolumeAttachmentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
