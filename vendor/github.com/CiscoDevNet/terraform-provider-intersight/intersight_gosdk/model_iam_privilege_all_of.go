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
)

// IamPrivilegeAllOf Definition of the list of properties defined in 'iam.Privilege', excluding properties defined in parent classes.
type IamPrivilegeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The hostname prefix of the resource corresponding to this privilege. For example \\'sentry\\' in https://sentry.intersight.com .
	HostnamePrefix *string `json:"HostnamePrefix,omitempty"`
	// The API method on the rest resource corresponding to privilege. For example READ, CREATE, UPDATE etc.
	Method *string `json:"Method,omitempty"`
	// The name of the privilege reported by microservice.
	Name *string `json:"Name,omitempty"`
	// The REST API path of the resource corresponding to this privilege. For example /v1/iam/Accounts, /v1/iam/Sessions.
	RestPath *string `json:"RestPath,omitempty"`
	// The URL path prefix of the resource corresponding to this privilege. For example /devops/kibana, /devops/grafana etc.
	UrlPrefix            *string                 `json:"UrlPrefix,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	System               *IamSystemRelationship  `json:"System,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamPrivilegeAllOf IamPrivilegeAllOf

// NewIamPrivilegeAllOf instantiates a new IamPrivilegeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPrivilegeAllOf(classId string, objectType string) *IamPrivilegeAllOf {
	this := IamPrivilegeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamPrivilegeAllOfWithDefaults instantiates a new IamPrivilegeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPrivilegeAllOfWithDefaults() *IamPrivilegeAllOf {
	this := IamPrivilegeAllOf{}
	var classId string = "iam.Privilege"
	this.ClassId = classId
	var objectType string = "iam.Privilege"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamPrivilegeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamPrivilegeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamPrivilegeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamPrivilegeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamPrivilegeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamPrivilegeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostnamePrefix returns the HostnamePrefix field value if set, zero value otherwise.
func (o *IamPrivilegeAllOf) GetHostnamePrefix() string {
	if o == nil || o.HostnamePrefix == nil {
		var ret string
		return ret
	}
	return *o.HostnamePrefix
}

// GetHostnamePrefixOk returns a tuple with the HostnamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeAllOf) GetHostnamePrefixOk() (*string, bool) {
	if o == nil || o.HostnamePrefix == nil {
		return nil, false
	}
	return o.HostnamePrefix, true
}

// HasHostnamePrefix returns a boolean if a field has been set.
func (o *IamPrivilegeAllOf) HasHostnamePrefix() bool {
	if o != nil && o.HostnamePrefix != nil {
		return true
	}

	return false
}

// SetHostnamePrefix gets a reference to the given string and assigns it to the HostnamePrefix field.
func (o *IamPrivilegeAllOf) SetHostnamePrefix(v string) {
	o.HostnamePrefix = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *IamPrivilegeAllOf) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeAllOf) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *IamPrivilegeAllOf) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *IamPrivilegeAllOf) SetMethod(v string) {
	o.Method = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamPrivilegeAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamPrivilegeAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamPrivilegeAllOf) SetName(v string) {
	o.Name = &v
}

// GetRestPath returns the RestPath field value if set, zero value otherwise.
func (o *IamPrivilegeAllOf) GetRestPath() string {
	if o == nil || o.RestPath == nil {
		var ret string
		return ret
	}
	return *o.RestPath
}

// GetRestPathOk returns a tuple with the RestPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeAllOf) GetRestPathOk() (*string, bool) {
	if o == nil || o.RestPath == nil {
		return nil, false
	}
	return o.RestPath, true
}

// HasRestPath returns a boolean if a field has been set.
func (o *IamPrivilegeAllOf) HasRestPath() bool {
	if o != nil && o.RestPath != nil {
		return true
	}

	return false
}

// SetRestPath gets a reference to the given string and assigns it to the RestPath field.
func (o *IamPrivilegeAllOf) SetRestPath(v string) {
	o.RestPath = &v
}

// GetUrlPrefix returns the UrlPrefix field value if set, zero value otherwise.
func (o *IamPrivilegeAllOf) GetUrlPrefix() string {
	if o == nil || o.UrlPrefix == nil {
		var ret string
		return ret
	}
	return *o.UrlPrefix
}

// GetUrlPrefixOk returns a tuple with the UrlPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeAllOf) GetUrlPrefixOk() (*string, bool) {
	if o == nil || o.UrlPrefix == nil {
		return nil, false
	}
	return o.UrlPrefix, true
}

// HasUrlPrefix returns a boolean if a field has been set.
func (o *IamPrivilegeAllOf) HasUrlPrefix() bool {
	if o != nil && o.UrlPrefix != nil {
		return true
	}

	return false
}

// SetUrlPrefix gets a reference to the given string and assigns it to the UrlPrefix field.
func (o *IamPrivilegeAllOf) SetUrlPrefix(v string) {
	o.UrlPrefix = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamPrivilegeAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamPrivilegeAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamPrivilegeAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *IamPrivilegeAllOf) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeAllOf) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *IamPrivilegeAllOf) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *IamPrivilegeAllOf) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

func (o IamPrivilegeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HostnamePrefix != nil {
		toSerialize["HostnamePrefix"] = o.HostnamePrefix
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RestPath != nil {
		toSerialize["RestPath"] = o.RestPath
	}
	if o.UrlPrefix != nil {
		toSerialize["UrlPrefix"] = o.UrlPrefix
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamPrivilegeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamPrivilegeAllOf := _IamPrivilegeAllOf{}

	if err = json.Unmarshal(bytes, &varIamPrivilegeAllOf); err == nil {
		*o = IamPrivilegeAllOf(varIamPrivilegeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HostnamePrefix")
		delete(additionalProperties, "Method")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RestPath")
		delete(additionalProperties, "UrlPrefix")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "System")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamPrivilegeAllOf struct {
	value *IamPrivilegeAllOf
	isSet bool
}

func (v NullableIamPrivilegeAllOf) Get() *IamPrivilegeAllOf {
	return v.value
}

func (v *NullableIamPrivilegeAllOf) Set(val *IamPrivilegeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPrivilegeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPrivilegeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPrivilegeAllOf(val *IamPrivilegeAllOf) *NullableIamPrivilegeAllOf {
	return &NullableIamPrivilegeAllOf{value: val, isSet: true}
}

func (v NullableIamPrivilegeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPrivilegeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
