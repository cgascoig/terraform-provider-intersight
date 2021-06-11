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

// WorkflowWebApi This models a single Web API request within a batch of requests that get executed within a single workflow task.
type WorkflowWebApi struct {
	WorkflowApi
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Collection of key value pairs to set in the request header as Cookie list.
	Cookies interface{} `json:"Cookies,omitempty"`
	// If the target type is Endpoint, this property determines whether the request is to be handled as internal request or external request by the device connector. * `Internal` - The endpoint API executed is an internal request handled by the device connector plugin. * `External` - The endpoint API request is passed through by the device connector.
	EndpointRequestType *string `json:"EndpointRequestType,omitempty"`
	// Collection of key value pairs to set in the request header.
	Headers interface{} `json:"Headers,omitempty"`
	// The HTTP method to be executed in the given URL (GET, POST, PUT, etc). If the value is not specified, GET will be used as default. The supported values are GET, POST, PUT, DELETE, PATCH, HEAD.
	Method *string `json:"Method,omitempty"`
	// The type of the intersight object for which API request is to be made. The property is valid in case of Intersight API calls and the base url is automatically prepended based on the value.
	MoType *string `json:"MoType,omitempty"`
	// The accepted web protocol values are http and https.
	Protocol *string `json:"Protocol,omitempty"`
	// If the web API is to be executed in a remote device connected to the Intersight through device connector, 'Endpoint' is expected as the value whereas if the API is an Intersight API, 'Local' is expected as the value.
	TargetType *string `json:"TargetType,omitempty"`
	// The URL of the resource in the target to which the API request is made.
	Url                  *string `json:"Url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWebApi WorkflowWebApi

// NewWorkflowWebApi instantiates a new WorkflowWebApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWebApi(classId string, objectType string) *WorkflowWebApi {
	this := WorkflowWebApi{}
	this.ClassId = classId
	this.ObjectType = objectType
	var endpointRequestType string = "Internal"
	this.EndpointRequestType = &endpointRequestType
	return &this
}

// NewWorkflowWebApiWithDefaults instantiates a new WorkflowWebApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWebApiWithDefaults() *WorkflowWebApi {
	this := WorkflowWebApi{}
	var classId string = "workflow.WebApi"
	this.ClassId = classId
	var objectType string = "workflow.WebApi"
	this.ObjectType = objectType
	var endpointRequestType string = "Internal"
	this.EndpointRequestType = &endpointRequestType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWebApi) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWebApi) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWebApi) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWebApi) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCookies returns the Cookies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWebApi) GetCookies() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Cookies
}

// GetCookiesOk returns a tuple with the Cookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWebApi) GetCookiesOk() (*interface{}, bool) {
	if o == nil || o.Cookies == nil {
		return nil, false
	}
	return &o.Cookies, true
}

// HasCookies returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasCookies() bool {
	if o != nil && o.Cookies != nil {
		return true
	}

	return false
}

// SetCookies gets a reference to the given interface{} and assigns it to the Cookies field.
func (o *WorkflowWebApi) SetCookies(v interface{}) {
	o.Cookies = v
}

// GetEndpointRequestType returns the EndpointRequestType field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetEndpointRequestType() string {
	if o == nil || o.EndpointRequestType == nil {
		var ret string
		return ret
	}
	return *o.EndpointRequestType
}

// GetEndpointRequestTypeOk returns a tuple with the EndpointRequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetEndpointRequestTypeOk() (*string, bool) {
	if o == nil || o.EndpointRequestType == nil {
		return nil, false
	}
	return o.EndpointRequestType, true
}

// HasEndpointRequestType returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasEndpointRequestType() bool {
	if o != nil && o.EndpointRequestType != nil {
		return true
	}

	return false
}

// SetEndpointRequestType gets a reference to the given string and assigns it to the EndpointRequestType field.
func (o *WorkflowWebApi) SetEndpointRequestType(v string) {
	o.EndpointRequestType = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWebApi) GetHeaders() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWebApi) GetHeadersOk() (*interface{}, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given interface{} and assigns it to the Headers field.
func (o *WorkflowWebApi) SetHeaders(v interface{}) {
	o.Headers = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *WorkflowWebApi) SetMethod(v string) {
	o.Method = &v
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetMoType() string {
	if o == nil || o.MoType == nil {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetMoTypeOk() (*string, bool) {
	if o == nil || o.MoType == nil {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasMoType() bool {
	if o != nil && o.MoType != nil {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *WorkflowWebApi) SetMoType(v string) {
	o.MoType = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *WorkflowWebApi) SetProtocol(v string) {
	o.Protocol = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *WorkflowWebApi) SetTargetType(v string) {
	o.TargetType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WorkflowWebApi) SetUrl(v string) {
	o.Url = &v
}

func (o WorkflowWebApi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowApi, errWorkflowApi := json.Marshal(o.WorkflowApi)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	errWorkflowApi = json.Unmarshal([]byte(serializedWorkflowApi), &toSerialize)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cookies != nil {
		toSerialize["Cookies"] = o.Cookies
	}
	if o.EndpointRequestType != nil {
		toSerialize["EndpointRequestType"] = o.EndpointRequestType
	}
	if o.Headers != nil {
		toSerialize["Headers"] = o.Headers
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.MoType != nil {
		toSerialize["MoType"] = o.MoType
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWebApi) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowWebApiWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Collection of key value pairs to set in the request header as Cookie list.
		Cookies interface{} `json:"Cookies,omitempty"`
		// If the target type is Endpoint, this property determines whether the request is to be handled as internal request or external request by the device connector. * `Internal` - The endpoint API executed is an internal request handled by the device connector plugin. * `External` - The endpoint API request is passed through by the device connector.
		EndpointRequestType *string `json:"EndpointRequestType,omitempty"`
		// Collection of key value pairs to set in the request header.
		Headers interface{} `json:"Headers,omitempty"`
		// The HTTP method to be executed in the given URL (GET, POST, PUT, etc). If the value is not specified, GET will be used as default. The supported values are GET, POST, PUT, DELETE, PATCH, HEAD.
		Method *string `json:"Method,omitempty"`
		// The type of the intersight object for which API request is to be made. The property is valid in case of Intersight API calls and the base url is automatically prepended based on the value.
		MoType *string `json:"MoType,omitempty"`
		// The accepted web protocol values are http and https.
		Protocol *string `json:"Protocol,omitempty"`
		// If the web API is to be executed in a remote device connected to the Intersight through device connector, 'Endpoint' is expected as the value whereas if the API is an Intersight API, 'Local' is expected as the value.
		TargetType *string `json:"TargetType,omitempty"`
		// The URL of the resource in the target to which the API request is made.
		Url *string `json:"Url,omitempty"`
	}

	varWorkflowWebApiWithoutEmbeddedStruct := WorkflowWebApiWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowWebApiWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowWebApi := _WorkflowWebApi{}
		varWorkflowWebApi.ClassId = varWorkflowWebApiWithoutEmbeddedStruct.ClassId
		varWorkflowWebApi.ObjectType = varWorkflowWebApiWithoutEmbeddedStruct.ObjectType
		varWorkflowWebApi.Cookies = varWorkflowWebApiWithoutEmbeddedStruct.Cookies
		varWorkflowWebApi.EndpointRequestType = varWorkflowWebApiWithoutEmbeddedStruct.EndpointRequestType
		varWorkflowWebApi.Headers = varWorkflowWebApiWithoutEmbeddedStruct.Headers
		varWorkflowWebApi.Method = varWorkflowWebApiWithoutEmbeddedStruct.Method
		varWorkflowWebApi.MoType = varWorkflowWebApiWithoutEmbeddedStruct.MoType
		varWorkflowWebApi.Protocol = varWorkflowWebApiWithoutEmbeddedStruct.Protocol
		varWorkflowWebApi.TargetType = varWorkflowWebApiWithoutEmbeddedStruct.TargetType
		varWorkflowWebApi.Url = varWorkflowWebApiWithoutEmbeddedStruct.Url
		*o = WorkflowWebApi(varWorkflowWebApi)
	} else {
		return err
	}

	varWorkflowWebApi := _WorkflowWebApi{}

	err = json.Unmarshal(bytes, &varWorkflowWebApi)
	if err == nil {
		o.WorkflowApi = varWorkflowWebApi.WorkflowApi
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cookies")
		delete(additionalProperties, "EndpointRequestType")
		delete(additionalProperties, "Headers")
		delete(additionalProperties, "Method")
		delete(additionalProperties, "MoType")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "TargetType")
		delete(additionalProperties, "Url")

		// remove fields from embedded structs
		reflectWorkflowApi := reflect.ValueOf(o.WorkflowApi)
		for i := 0; i < reflectWorkflowApi.Type().NumField(); i++ {
			t := reflectWorkflowApi.Type().Field(i)

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

type NullableWorkflowWebApi struct {
	value *WorkflowWebApi
	isSet bool
}

func (v NullableWorkflowWebApi) Get() *WorkflowWebApi {
	return v.value
}

func (v *NullableWorkflowWebApi) Set(val *WorkflowWebApi) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWebApi) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWebApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWebApi(val *WorkflowWebApi) *NullableWorkflowWebApi {
	return &NullableWorkflowWebApi{value: val, isSet: true}
}

func (v NullableWorkflowWebApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWebApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
