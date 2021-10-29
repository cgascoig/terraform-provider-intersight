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
	"reflect"
	"strings"
)

// NiatelemetryNexusDashboardDetails Details of NexusDashboard.
type NiatelemetryNexusDashboardDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the nexus dashboard cluster.
	ClusterName *string `json:"ClusterName,omitempty"`
	// Model of the nexus dashboard cluster.
	DeviceModel *string `json:"DeviceModel,omitempty"`
	// Name of the NexusDashboard.
	NexusDashboardName *string `json:"NexusDashboardName,omitempty"`
	// Serial number of NexusDashboard.
	NexusDashboardSerialNumber *string `json:"NexusDashboardSerialNumber,omitempty"`
	// Node type of the nexus dashboard cluster.
	Type                 *string                              `json:"Type,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNexusDashboardDetails NiatelemetryNexusDashboardDetails

// NewNiatelemetryNexusDashboardDetails instantiates a new NiatelemetryNexusDashboardDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNexusDashboardDetails(classId string, objectType string) *NiatelemetryNexusDashboardDetails {
	this := NiatelemetryNexusDashboardDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNexusDashboardDetailsWithDefaults instantiates a new NiatelemetryNexusDashboardDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNexusDashboardDetailsWithDefaults() *NiatelemetryNexusDashboardDetails {
	this := NiatelemetryNexusDashboardDetails{}
	var classId string = "niatelemetry.NexusDashboardDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.NexusDashboardDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNexusDashboardDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNexusDashboardDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNexusDashboardDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNexusDashboardDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardDetails) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetails) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardDetails) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *NiatelemetryNexusDashboardDetails) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardDetails) GetDeviceModel() string {
	if o == nil || o.DeviceModel == nil {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetails) GetDeviceModelOk() (*string, bool) {
	if o == nil || o.DeviceModel == nil {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardDetails) HasDeviceModel() bool {
	if o != nil && o.DeviceModel != nil {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *NiatelemetryNexusDashboardDetails) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetNexusDashboardName returns the NexusDashboardName field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardDetails) GetNexusDashboardName() string {
	if o == nil || o.NexusDashboardName == nil {
		var ret string
		return ret
	}
	return *o.NexusDashboardName
}

// GetNexusDashboardNameOk returns a tuple with the NexusDashboardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetails) GetNexusDashboardNameOk() (*string, bool) {
	if o == nil || o.NexusDashboardName == nil {
		return nil, false
	}
	return o.NexusDashboardName, true
}

// HasNexusDashboardName returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardDetails) HasNexusDashboardName() bool {
	if o != nil && o.NexusDashboardName != nil {
		return true
	}

	return false
}

// SetNexusDashboardName gets a reference to the given string and assigns it to the NexusDashboardName field.
func (o *NiatelemetryNexusDashboardDetails) SetNexusDashboardName(v string) {
	o.NexusDashboardName = &v
}

// GetNexusDashboardSerialNumber returns the NexusDashboardSerialNumber field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardDetails) GetNexusDashboardSerialNumber() string {
	if o == nil || o.NexusDashboardSerialNumber == nil {
		var ret string
		return ret
	}
	return *o.NexusDashboardSerialNumber
}

// GetNexusDashboardSerialNumberOk returns a tuple with the NexusDashboardSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetails) GetNexusDashboardSerialNumberOk() (*string, bool) {
	if o == nil || o.NexusDashboardSerialNumber == nil {
		return nil, false
	}
	return o.NexusDashboardSerialNumber, true
}

// HasNexusDashboardSerialNumber returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardDetails) HasNexusDashboardSerialNumber() bool {
	if o != nil && o.NexusDashboardSerialNumber != nil {
		return true
	}

	return false
}

// SetNexusDashboardSerialNumber gets a reference to the given string and assigns it to the NexusDashboardSerialNumber field.
func (o *NiatelemetryNexusDashboardDetails) SetNexusDashboardSerialNumber(v string) {
	o.NexusDashboardSerialNumber = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardDetails) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetails) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardDetails) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NiatelemetryNexusDashboardDetails) SetType(v string) {
	o.Type = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNexusDashboardDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNexusDashboardDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if o.DeviceModel != nil {
		toSerialize["DeviceModel"] = o.DeviceModel
	}
	if o.NexusDashboardName != nil {
		toSerialize["NexusDashboardName"] = o.NexusDashboardName
	}
	if o.NexusDashboardSerialNumber != nil {
		toSerialize["NexusDashboardSerialNumber"] = o.NexusDashboardSerialNumber
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNexusDashboardDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the nexus dashboard cluster.
		ClusterName *string `json:"ClusterName,omitempty"`
		// Model of the nexus dashboard cluster.
		DeviceModel *string `json:"DeviceModel,omitempty"`
		// Name of the NexusDashboard.
		NexusDashboardName *string `json:"NexusDashboardName,omitempty"`
		// Serial number of NexusDashboard.
		NexusDashboardSerialNumber *string `json:"NexusDashboardSerialNumber,omitempty"`
		// Node type of the nexus dashboard cluster.
		Type             *string                              `json:"Type,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct := NiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNexusDashboardDetails := _NiatelemetryNexusDashboardDetails{}
		varNiatelemetryNexusDashboardDetails.ClassId = varNiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryNexusDashboardDetails.ObjectType = varNiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryNexusDashboardDetails.ClusterName = varNiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct.ClusterName
		varNiatelemetryNexusDashboardDetails.DeviceModel = varNiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct.DeviceModel
		varNiatelemetryNexusDashboardDetails.NexusDashboardName = varNiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct.NexusDashboardName
		varNiatelemetryNexusDashboardDetails.NexusDashboardSerialNumber = varNiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct.NexusDashboardSerialNumber
		varNiatelemetryNexusDashboardDetails.Type = varNiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct.Type
		varNiatelemetryNexusDashboardDetails.RegisteredDevice = varNiatelemetryNexusDashboardDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryNexusDashboardDetails(varNiatelemetryNexusDashboardDetails)
	} else {
		return err
	}

	varNiatelemetryNexusDashboardDetails := _NiatelemetryNexusDashboardDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryNexusDashboardDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryNexusDashboardDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "DeviceModel")
		delete(additionalProperties, "NexusDashboardName")
		delete(additionalProperties, "NexusDashboardSerialNumber")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableNiatelemetryNexusDashboardDetails struct {
	value *NiatelemetryNexusDashboardDetails
	isSet bool
}

func (v NullableNiatelemetryNexusDashboardDetails) Get() *NiatelemetryNexusDashboardDetails {
	return v.value
}

func (v *NullableNiatelemetryNexusDashboardDetails) Set(val *NiatelemetryNexusDashboardDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNexusDashboardDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNexusDashboardDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNexusDashboardDetails(val *NiatelemetryNexusDashboardDetails) *NullableNiatelemetryNexusDashboardDetails {
	return &NullableNiatelemetryNexusDashboardDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryNexusDashboardDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNexusDashboardDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
