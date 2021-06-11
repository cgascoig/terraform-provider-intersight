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

// UcsdConnectorPack Information about the connector packs present in the backup file.
type UcsdConnectorPack struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of the connector pack whether it is enabled or disabled.
	ConnectorFeature *string  `json:"ConnectorFeature,omitempty"`
	DependencyNames  []string `json:"DependencyNames,omitempty"`
	// Version of the connector pack that is last downloaded successfully to UCS Director.
	DownloadedVersion *string `json:"DownloadedVersion,omitempty"`
	// Name of the connector pack running on the UCS Director.
	Name     *string  `json:"Name,omitempty"`
	Services []string `json:"Services,omitempty"`
	// State of the connector pack whether it is enabled or disabled.
	State *string `json:"State,omitempty"`
	// Version of the connector pack.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UcsdConnectorPack UcsdConnectorPack

// NewUcsdConnectorPack instantiates a new UcsdConnectorPack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUcsdConnectorPack(classId string, objectType string) *UcsdConnectorPack {
	this := UcsdConnectorPack{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewUcsdConnectorPackWithDefaults instantiates a new UcsdConnectorPack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUcsdConnectorPackWithDefaults() *UcsdConnectorPack {
	this := UcsdConnectorPack{}
	var classId string = "ucsd.ConnectorPack"
	this.ClassId = classId
	var objectType string = "ucsd.ConnectorPack"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UcsdConnectorPack) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UcsdConnectorPack) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *UcsdConnectorPack) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UcsdConnectorPack) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectorFeature returns the ConnectorFeature field value if set, zero value otherwise.
func (o *UcsdConnectorPack) GetConnectorFeature() string {
	if o == nil || o.ConnectorFeature == nil {
		var ret string
		return ret
	}
	return *o.ConnectorFeature
}

// GetConnectorFeatureOk returns a tuple with the ConnectorFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetConnectorFeatureOk() (*string, bool) {
	if o == nil || o.ConnectorFeature == nil {
		return nil, false
	}
	return o.ConnectorFeature, true
}

// HasConnectorFeature returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasConnectorFeature() bool {
	if o != nil && o.ConnectorFeature != nil {
		return true
	}

	return false
}

// SetConnectorFeature gets a reference to the given string and assigns it to the ConnectorFeature field.
func (o *UcsdConnectorPack) SetConnectorFeature(v string) {
	o.ConnectorFeature = &v
}

// GetDependencyNames returns the DependencyNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UcsdConnectorPack) GetDependencyNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DependencyNames
}

// GetDependencyNamesOk returns a tuple with the DependencyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UcsdConnectorPack) GetDependencyNamesOk() (*[]string, bool) {
	if o == nil || o.DependencyNames == nil {
		return nil, false
	}
	return &o.DependencyNames, true
}

// HasDependencyNames returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasDependencyNames() bool {
	if o != nil && o.DependencyNames != nil {
		return true
	}

	return false
}

// SetDependencyNames gets a reference to the given []string and assigns it to the DependencyNames field.
func (o *UcsdConnectorPack) SetDependencyNames(v []string) {
	o.DependencyNames = v
}

// GetDownloadedVersion returns the DownloadedVersion field value if set, zero value otherwise.
func (o *UcsdConnectorPack) GetDownloadedVersion() string {
	if o == nil || o.DownloadedVersion == nil {
		var ret string
		return ret
	}
	return *o.DownloadedVersion
}

// GetDownloadedVersionOk returns a tuple with the DownloadedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetDownloadedVersionOk() (*string, bool) {
	if o == nil || o.DownloadedVersion == nil {
		return nil, false
	}
	return o.DownloadedVersion, true
}

// HasDownloadedVersion returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasDownloadedVersion() bool {
	if o != nil && o.DownloadedVersion != nil {
		return true
	}

	return false
}

// SetDownloadedVersion gets a reference to the given string and assigns it to the DownloadedVersion field.
func (o *UcsdConnectorPack) SetDownloadedVersion(v string) {
	o.DownloadedVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UcsdConnectorPack) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UcsdConnectorPack) SetName(v string) {
	o.Name = &v
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UcsdConnectorPack) GetServices() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UcsdConnectorPack) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *UcsdConnectorPack) SetServices(v []string) {
	o.Services = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UcsdConnectorPack) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UcsdConnectorPack) SetState(v string) {
	o.State = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UcsdConnectorPack) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UcsdConnectorPack) SetVersion(v string) {
	o.Version = &v
}

func (o UcsdConnectorPack) MarshalJSON() ([]byte, error) {
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
	if o.ConnectorFeature != nil {
		toSerialize["ConnectorFeature"] = o.ConnectorFeature
	}
	if o.DependencyNames != nil {
		toSerialize["DependencyNames"] = o.DependencyNames
	}
	if o.DownloadedVersion != nil {
		toSerialize["DownloadedVersion"] = o.DownloadedVersion
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UcsdConnectorPack) UnmarshalJSON(bytes []byte) (err error) {
	type UcsdConnectorPackWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// State of the connector pack whether it is enabled or disabled.
		ConnectorFeature *string  `json:"ConnectorFeature,omitempty"`
		DependencyNames  []string `json:"DependencyNames,omitempty"`
		// Version of the connector pack that is last downloaded successfully to UCS Director.
		DownloadedVersion *string `json:"DownloadedVersion,omitempty"`
		// Name of the connector pack running on the UCS Director.
		Name     *string  `json:"Name,omitempty"`
		Services []string `json:"Services,omitempty"`
		// State of the connector pack whether it is enabled or disabled.
		State *string `json:"State,omitempty"`
		// Version of the connector pack.
		Version *string `json:"Version,omitempty"`
	}

	varUcsdConnectorPackWithoutEmbeddedStruct := UcsdConnectorPackWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUcsdConnectorPackWithoutEmbeddedStruct)
	if err == nil {
		varUcsdConnectorPack := _UcsdConnectorPack{}
		varUcsdConnectorPack.ClassId = varUcsdConnectorPackWithoutEmbeddedStruct.ClassId
		varUcsdConnectorPack.ObjectType = varUcsdConnectorPackWithoutEmbeddedStruct.ObjectType
		varUcsdConnectorPack.ConnectorFeature = varUcsdConnectorPackWithoutEmbeddedStruct.ConnectorFeature
		varUcsdConnectorPack.DependencyNames = varUcsdConnectorPackWithoutEmbeddedStruct.DependencyNames
		varUcsdConnectorPack.DownloadedVersion = varUcsdConnectorPackWithoutEmbeddedStruct.DownloadedVersion
		varUcsdConnectorPack.Name = varUcsdConnectorPackWithoutEmbeddedStruct.Name
		varUcsdConnectorPack.Services = varUcsdConnectorPackWithoutEmbeddedStruct.Services
		varUcsdConnectorPack.State = varUcsdConnectorPackWithoutEmbeddedStruct.State
		varUcsdConnectorPack.Version = varUcsdConnectorPackWithoutEmbeddedStruct.Version
		*o = UcsdConnectorPack(varUcsdConnectorPack)
	} else {
		return err
	}

	varUcsdConnectorPack := _UcsdConnectorPack{}

	err = json.Unmarshal(bytes, &varUcsdConnectorPack)
	if err == nil {
		o.MoBaseComplexType = varUcsdConnectorPack.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectorFeature")
		delete(additionalProperties, "DependencyNames")
		delete(additionalProperties, "DownloadedVersion")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Services")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Version")

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

type NullableUcsdConnectorPack struct {
	value *UcsdConnectorPack
	isSet bool
}

func (v NullableUcsdConnectorPack) Get() *UcsdConnectorPack {
	return v.value
}

func (v *NullableUcsdConnectorPack) Set(val *UcsdConnectorPack) {
	v.value = val
	v.isSet = true
}

func (v NullableUcsdConnectorPack) IsSet() bool {
	return v.isSet
}

func (v *NullableUcsdConnectorPack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcsdConnectorPack(val *UcsdConnectorPack) *NullableUcsdConnectorPack {
	return &NullableUcsdConnectorPack{value: val, isSet: true}
}

func (v NullableUcsdConnectorPack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcsdConnectorPack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
