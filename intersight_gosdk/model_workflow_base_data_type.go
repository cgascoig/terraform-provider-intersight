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

// WorkflowBaseDataType The base data type that is used to derive all the other data types for inputs and outputs.
type WorkflowBaseDataType struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string                       `json:"ObjectType"`
	Default    NullableWorkflowDefaultValue `json:"Default,omitempty"`
	// Provide a detailed description of the data type.
	Description *string                     `json:"Description,omitempty"`
	DisplayMeta NullableWorkflowDisplayMeta `json:"DisplayMeta,omitempty"`
	// JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'.
	InputParameters interface{} `json:"InputParameters,omitempty"`
	// Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character.
	Label *string `json:"Label,omitempty"`
	// Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.
	Name *string `json:"Name,omitempty"`
	// Specifies whether this parameter is required. The field is applicable for task and workflow.
	Required             *bool `json:"Required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowBaseDataType WorkflowBaseDataType

// NewWorkflowBaseDataType instantiates a new WorkflowBaseDataType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBaseDataType(classId string, objectType string) *WorkflowBaseDataType {
	this := WorkflowBaseDataType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowBaseDataTypeWithDefaults instantiates a new WorkflowBaseDataType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBaseDataTypeWithDefaults() *WorkflowBaseDataType {
	this := WorkflowBaseDataType{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowBaseDataType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowBaseDataType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowBaseDataType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowBaseDataType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowBaseDataType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowBaseDataType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBaseDataType) GetDefault() WorkflowDefaultValue {
	if o == nil || o.Default.Get() == nil {
		var ret WorkflowDefaultValue
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBaseDataType) GetDefaultOk() (*WorkflowDefaultValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *WorkflowBaseDataType) HasDefault() bool {
	if o != nil && o.Default.IsSet() {
		return true
	}

	return false
}

// SetDefault gets a reference to the given NullableWorkflowDefaultValue and assigns it to the Default field.
func (o *WorkflowBaseDataType) SetDefault(v WorkflowDefaultValue) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil
func (o *WorkflowBaseDataType) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil
func (o *WorkflowBaseDataType) UnsetDefault() {
	o.Default.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowBaseDataType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBaseDataType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowBaseDataType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowBaseDataType) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayMeta returns the DisplayMeta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBaseDataType) GetDisplayMeta() WorkflowDisplayMeta {
	if o == nil || o.DisplayMeta.Get() == nil {
		var ret WorkflowDisplayMeta
		return ret
	}
	return *o.DisplayMeta.Get()
}

// GetDisplayMetaOk returns a tuple with the DisplayMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBaseDataType) GetDisplayMetaOk() (*WorkflowDisplayMeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayMeta.Get(), o.DisplayMeta.IsSet()
}

// HasDisplayMeta returns a boolean if a field has been set.
func (o *WorkflowBaseDataType) HasDisplayMeta() bool {
	if o != nil && o.DisplayMeta.IsSet() {
		return true
	}

	return false
}

// SetDisplayMeta gets a reference to the given NullableWorkflowDisplayMeta and assigns it to the DisplayMeta field.
func (o *WorkflowBaseDataType) SetDisplayMeta(v WorkflowDisplayMeta) {
	o.DisplayMeta.Set(&v)
}

// SetDisplayMetaNil sets the value for DisplayMeta to be an explicit nil
func (o *WorkflowBaseDataType) SetDisplayMetaNil() {
	o.DisplayMeta.Set(nil)
}

// UnsetDisplayMeta ensures that no value is present for DisplayMeta, not even an explicit nil
func (o *WorkflowBaseDataType) UnsetDisplayMeta() {
	o.DisplayMeta.Unset()
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBaseDataType) GetInputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBaseDataType) GetInputParametersOk() (*interface{}, bool) {
	if o == nil || o.InputParameters == nil {
		return nil, false
	}
	return &o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *WorkflowBaseDataType) HasInputParameters() bool {
	if o != nil && o.InputParameters != nil {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given interface{} and assigns it to the InputParameters field.
func (o *WorkflowBaseDataType) SetInputParameters(v interface{}) {
	o.InputParameters = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowBaseDataType) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBaseDataType) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowBaseDataType) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowBaseDataType) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowBaseDataType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBaseDataType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowBaseDataType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowBaseDataType) SetName(v string) {
	o.Name = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *WorkflowBaseDataType) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBaseDataType) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *WorkflowBaseDataType) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *WorkflowBaseDataType) SetRequired(v bool) {
	o.Required = &v
}

func (o WorkflowBaseDataType) MarshalJSON() ([]byte, error) {
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
	if o.Default.IsSet() {
		toSerialize["Default"] = o.Default.Get()
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DisplayMeta.IsSet() {
		toSerialize["DisplayMeta"] = o.DisplayMeta.Get()
	}
	if o.InputParameters != nil {
		toSerialize["InputParameters"] = o.InputParameters
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Required != nil {
		toSerialize["Required"] = o.Required
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowBaseDataType) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowBaseDataTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string                       `json:"ObjectType"`
		Default    NullableWorkflowDefaultValue `json:"Default,omitempty"`
		// Provide a detailed description of the data type.
		Description *string                     `json:"Description,omitempty"`
		DisplayMeta NullableWorkflowDisplayMeta `json:"DisplayMeta,omitempty"`
		// JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'.
		InputParameters interface{} `json:"InputParameters,omitempty"`
		// Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character.
		Label *string `json:"Label,omitempty"`
		// Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.
		Name *string `json:"Name,omitempty"`
		// Specifies whether this parameter is required. The field is applicable for task and workflow.
		Required *bool `json:"Required,omitempty"`
	}

	varWorkflowBaseDataTypeWithoutEmbeddedStruct := WorkflowBaseDataTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowBaseDataTypeWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowBaseDataType := _WorkflowBaseDataType{}
		varWorkflowBaseDataType.ClassId = varWorkflowBaseDataTypeWithoutEmbeddedStruct.ClassId
		varWorkflowBaseDataType.ObjectType = varWorkflowBaseDataTypeWithoutEmbeddedStruct.ObjectType
		varWorkflowBaseDataType.Default = varWorkflowBaseDataTypeWithoutEmbeddedStruct.Default
		varWorkflowBaseDataType.Description = varWorkflowBaseDataTypeWithoutEmbeddedStruct.Description
		varWorkflowBaseDataType.DisplayMeta = varWorkflowBaseDataTypeWithoutEmbeddedStruct.DisplayMeta
		varWorkflowBaseDataType.InputParameters = varWorkflowBaseDataTypeWithoutEmbeddedStruct.InputParameters
		varWorkflowBaseDataType.Label = varWorkflowBaseDataTypeWithoutEmbeddedStruct.Label
		varWorkflowBaseDataType.Name = varWorkflowBaseDataTypeWithoutEmbeddedStruct.Name
		varWorkflowBaseDataType.Required = varWorkflowBaseDataTypeWithoutEmbeddedStruct.Required
		*o = WorkflowBaseDataType(varWorkflowBaseDataType)
	} else {
		return err
	}

	varWorkflowBaseDataType := _WorkflowBaseDataType{}

	err = json.Unmarshal(bytes, &varWorkflowBaseDataType)
	if err == nil {
		o.MoBaseComplexType = varWorkflowBaseDataType.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Default")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "DisplayMeta")
		delete(additionalProperties, "InputParameters")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Required")

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

type NullableWorkflowBaseDataType struct {
	value *WorkflowBaseDataType
	isSet bool
}

func (v NullableWorkflowBaseDataType) Get() *WorkflowBaseDataType {
	return v.value
}

func (v *NullableWorkflowBaseDataType) Set(val *WorkflowBaseDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBaseDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBaseDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBaseDataType(val *WorkflowBaseDataType) *NullableWorkflowBaseDataType {
	return &NullableWorkflowBaseDataType{value: val, isSet: true}
}

func (v NullableWorkflowBaseDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBaseDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
