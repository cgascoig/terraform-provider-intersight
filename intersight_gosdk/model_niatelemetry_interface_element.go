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
	"reflect"
	"strings"
)

// NiatelemetryInterfaceElement Stores information related to ports of a device.
type NiatelemetryInterfaceElement struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Return value of name of the port.
	Name *string `json:"Name,omitempty"`
	// Return value of operState attribute.
	OperState *string `json:"OperState,omitempty"`
	// Return whether sfp is present or not.
	XcvrPresent          *string `json:"XcvrPresent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryInterfaceElement NiatelemetryInterfaceElement

// NewNiatelemetryInterfaceElement instantiates a new NiatelemetryInterfaceElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryInterfaceElement(classId string, objectType string) *NiatelemetryInterfaceElement {
	this := NiatelemetryInterfaceElement{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryInterfaceElementWithDefaults instantiates a new NiatelemetryInterfaceElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryInterfaceElementWithDefaults() *NiatelemetryInterfaceElement {
	this := NiatelemetryInterfaceElement{}
	var classId string = "niatelemetry.InterfaceElement"
	this.ClassId = classId
	var objectType string = "niatelemetry.InterfaceElement"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryInterfaceElement) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryInterfaceElement) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryInterfaceElement) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryInterfaceElement) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryInterfaceElement) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryInterfaceElement) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryInterfaceElement) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInterfaceElement) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryInterfaceElement) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryInterfaceElement) SetName(v string) {
	o.Name = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *NiatelemetryInterfaceElement) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInterfaceElement) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *NiatelemetryInterfaceElement) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *NiatelemetryInterfaceElement) SetOperState(v string) {
	o.OperState = &v
}

// GetXcvrPresent returns the XcvrPresent field value if set, zero value otherwise.
func (o *NiatelemetryInterfaceElement) GetXcvrPresent() string {
	if o == nil || o.XcvrPresent == nil {
		var ret string
		return ret
	}
	return *o.XcvrPresent
}

// GetXcvrPresentOk returns a tuple with the XcvrPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInterfaceElement) GetXcvrPresentOk() (*string, bool) {
	if o == nil || o.XcvrPresent == nil {
		return nil, false
	}
	return o.XcvrPresent, true
}

// HasXcvrPresent returns a boolean if a field has been set.
func (o *NiatelemetryInterfaceElement) HasXcvrPresent() bool {
	if o != nil && o.XcvrPresent != nil {
		return true
	}

	return false
}

// SetXcvrPresent gets a reference to the given string and assigns it to the XcvrPresent field.
func (o *NiatelemetryInterfaceElement) SetXcvrPresent(v string) {
	o.XcvrPresent = &v
}

func (o NiatelemetryInterfaceElement) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.XcvrPresent != nil {
		toSerialize["XcvrPresent"] = o.XcvrPresent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryInterfaceElement) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryInterfaceElementWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Return value of name of the port.
		Name *string `json:"Name,omitempty"`
		// Return value of operState attribute.
		OperState *string `json:"OperState,omitempty"`
		// Return whether sfp is present or not.
		XcvrPresent *string `json:"XcvrPresent,omitempty"`
	}

	varNiatelemetryInterfaceElementWithoutEmbeddedStruct := NiatelemetryInterfaceElementWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryInterfaceElementWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryInterfaceElement := _NiatelemetryInterfaceElement{}
		varNiatelemetryInterfaceElement.ClassId = varNiatelemetryInterfaceElementWithoutEmbeddedStruct.ClassId
		varNiatelemetryInterfaceElement.ObjectType = varNiatelemetryInterfaceElementWithoutEmbeddedStruct.ObjectType
		varNiatelemetryInterfaceElement.Name = varNiatelemetryInterfaceElementWithoutEmbeddedStruct.Name
		varNiatelemetryInterfaceElement.OperState = varNiatelemetryInterfaceElementWithoutEmbeddedStruct.OperState
		varNiatelemetryInterfaceElement.XcvrPresent = varNiatelemetryInterfaceElementWithoutEmbeddedStruct.XcvrPresent
		*o = NiatelemetryInterfaceElement(varNiatelemetryInterfaceElement)
	} else {
		return err
	}

	varNiatelemetryInterfaceElement := _NiatelemetryInterfaceElement{}

	err = json.Unmarshal(bytes, &varNiatelemetryInterfaceElement)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryInterfaceElement.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "XcvrPresent")

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

type NullableNiatelemetryInterfaceElement struct {
	value *NiatelemetryInterfaceElement
	isSet bool
}

func (v NullableNiatelemetryInterfaceElement) Get() *NiatelemetryInterfaceElement {
	return v.value
}

func (v *NullableNiatelemetryInterfaceElement) Set(val *NiatelemetryInterfaceElement) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryInterfaceElement) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryInterfaceElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryInterfaceElement(val *NiatelemetryInterfaceElement) *NullableNiatelemetryInterfaceElement {
	return &NullableNiatelemetryInterfaceElement{value: val, isSet: true}
}

func (v NullableNiatelemetryInterfaceElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryInterfaceElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
