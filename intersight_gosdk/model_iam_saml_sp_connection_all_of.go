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

// IamSamlSpConnectionAllOf Definition of the list of properties defined in 'iam.SamlSpConnection', excluding properties defined in parent classes.
type IamSamlSpConnectionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// SingleLogout Services binding in IdP Metadata.
	IdentityProviderSloBinding *string `json:"IdentityProviderSloBinding,omitempty"`
	// SingleLogOut Services location in IdP Metadata.
	IdentityProviderSloUrl *string `json:"IdentityProviderSloUrl,omitempty"`
	// SingleSignOn Services binding in IdP Metadata.
	IdentityProviderSsoBinding *string `json:"IdentityProviderSsoBinding,omitempty"`
	// SingleSignOn Services location in IdP Metadata.
	IdentityProviderSsoUrl *string `json:"IdentityProviderSsoUrl,omitempty"`
	// Decoded Certificate from IdP Metatdata.
	IdpCertificateStore interface{} `json:"IdpCertificateStore,omitempty"`
	// WantAuthnRequestsSigned from IdP Metadata.
	SignAuthnRequests    *bool `json:"SignAuthnRequests,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamSamlSpConnectionAllOf IamSamlSpConnectionAllOf

// NewIamSamlSpConnectionAllOf instantiates a new IamSamlSpConnectionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSamlSpConnectionAllOf(classId string, objectType string) *IamSamlSpConnectionAllOf {
	this := IamSamlSpConnectionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamSamlSpConnectionAllOfWithDefaults instantiates a new IamSamlSpConnectionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSamlSpConnectionAllOfWithDefaults() *IamSamlSpConnectionAllOf {
	this := IamSamlSpConnectionAllOf{}
	var classId string = "iam.SamlSpConnection"
	this.ClassId = classId
	var objectType string = "iam.SamlSpConnection"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamSamlSpConnectionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnectionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamSamlSpConnectionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamSamlSpConnectionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnectionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamSamlSpConnectionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdentityProviderSloBinding returns the IdentityProviderSloBinding field value if set, zero value otherwise.
func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSloBinding() string {
	if o == nil || o.IdentityProviderSloBinding == nil {
		var ret string
		return ret
	}
	return *o.IdentityProviderSloBinding
}

// GetIdentityProviderSloBindingOk returns a tuple with the IdentityProviderSloBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSloBindingOk() (*string, bool) {
	if o == nil || o.IdentityProviderSloBinding == nil {
		return nil, false
	}
	return o.IdentityProviderSloBinding, true
}

// HasIdentityProviderSloBinding returns a boolean if a field has been set.
func (o *IamSamlSpConnectionAllOf) HasIdentityProviderSloBinding() bool {
	if o != nil && o.IdentityProviderSloBinding != nil {
		return true
	}

	return false
}

// SetIdentityProviderSloBinding gets a reference to the given string and assigns it to the IdentityProviderSloBinding field.
func (o *IamSamlSpConnectionAllOf) SetIdentityProviderSloBinding(v string) {
	o.IdentityProviderSloBinding = &v
}

// GetIdentityProviderSloUrl returns the IdentityProviderSloUrl field value if set, zero value otherwise.
func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSloUrl() string {
	if o == nil || o.IdentityProviderSloUrl == nil {
		var ret string
		return ret
	}
	return *o.IdentityProviderSloUrl
}

// GetIdentityProviderSloUrlOk returns a tuple with the IdentityProviderSloUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSloUrlOk() (*string, bool) {
	if o == nil || o.IdentityProviderSloUrl == nil {
		return nil, false
	}
	return o.IdentityProviderSloUrl, true
}

// HasIdentityProviderSloUrl returns a boolean if a field has been set.
func (o *IamSamlSpConnectionAllOf) HasIdentityProviderSloUrl() bool {
	if o != nil && o.IdentityProviderSloUrl != nil {
		return true
	}

	return false
}

// SetIdentityProviderSloUrl gets a reference to the given string and assigns it to the IdentityProviderSloUrl field.
func (o *IamSamlSpConnectionAllOf) SetIdentityProviderSloUrl(v string) {
	o.IdentityProviderSloUrl = &v
}

// GetIdentityProviderSsoBinding returns the IdentityProviderSsoBinding field value if set, zero value otherwise.
func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSsoBinding() string {
	if o == nil || o.IdentityProviderSsoBinding == nil {
		var ret string
		return ret
	}
	return *o.IdentityProviderSsoBinding
}

// GetIdentityProviderSsoBindingOk returns a tuple with the IdentityProviderSsoBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSsoBindingOk() (*string, bool) {
	if o == nil || o.IdentityProviderSsoBinding == nil {
		return nil, false
	}
	return o.IdentityProviderSsoBinding, true
}

// HasIdentityProviderSsoBinding returns a boolean if a field has been set.
func (o *IamSamlSpConnectionAllOf) HasIdentityProviderSsoBinding() bool {
	if o != nil && o.IdentityProviderSsoBinding != nil {
		return true
	}

	return false
}

// SetIdentityProviderSsoBinding gets a reference to the given string and assigns it to the IdentityProviderSsoBinding field.
func (o *IamSamlSpConnectionAllOf) SetIdentityProviderSsoBinding(v string) {
	o.IdentityProviderSsoBinding = &v
}

// GetIdentityProviderSsoUrl returns the IdentityProviderSsoUrl field value if set, zero value otherwise.
func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSsoUrl() string {
	if o == nil || o.IdentityProviderSsoUrl == nil {
		var ret string
		return ret
	}
	return *o.IdentityProviderSsoUrl
}

// GetIdentityProviderSsoUrlOk returns a tuple with the IdentityProviderSsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnectionAllOf) GetIdentityProviderSsoUrlOk() (*string, bool) {
	if o == nil || o.IdentityProviderSsoUrl == nil {
		return nil, false
	}
	return o.IdentityProviderSsoUrl, true
}

// HasIdentityProviderSsoUrl returns a boolean if a field has been set.
func (o *IamSamlSpConnectionAllOf) HasIdentityProviderSsoUrl() bool {
	if o != nil && o.IdentityProviderSsoUrl != nil {
		return true
	}

	return false
}

// SetIdentityProviderSsoUrl gets a reference to the given string and assigns it to the IdentityProviderSsoUrl field.
func (o *IamSamlSpConnectionAllOf) SetIdentityProviderSsoUrl(v string) {
	o.IdentityProviderSsoUrl = &v
}

// GetIdpCertificateStore returns the IdpCertificateStore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamSamlSpConnectionAllOf) GetIdpCertificateStore() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.IdpCertificateStore
}

// GetIdpCertificateStoreOk returns a tuple with the IdpCertificateStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamSamlSpConnectionAllOf) GetIdpCertificateStoreOk() (*interface{}, bool) {
	if o == nil || o.IdpCertificateStore == nil {
		return nil, false
	}
	return &o.IdpCertificateStore, true
}

// HasIdpCertificateStore returns a boolean if a field has been set.
func (o *IamSamlSpConnectionAllOf) HasIdpCertificateStore() bool {
	if o != nil && o.IdpCertificateStore != nil {
		return true
	}

	return false
}

// SetIdpCertificateStore gets a reference to the given interface{} and assigns it to the IdpCertificateStore field.
func (o *IamSamlSpConnectionAllOf) SetIdpCertificateStore(v interface{}) {
	o.IdpCertificateStore = v
}

// GetSignAuthnRequests returns the SignAuthnRequests field value if set, zero value otherwise.
func (o *IamSamlSpConnectionAllOf) GetSignAuthnRequests() bool {
	if o == nil || o.SignAuthnRequests == nil {
		var ret bool
		return ret
	}
	return *o.SignAuthnRequests
}

// GetSignAuthnRequestsOk returns a tuple with the SignAuthnRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSamlSpConnectionAllOf) GetSignAuthnRequestsOk() (*bool, bool) {
	if o == nil || o.SignAuthnRequests == nil {
		return nil, false
	}
	return o.SignAuthnRequests, true
}

// HasSignAuthnRequests returns a boolean if a field has been set.
func (o *IamSamlSpConnectionAllOf) HasSignAuthnRequests() bool {
	if o != nil && o.SignAuthnRequests != nil {
		return true
	}

	return false
}

// SetSignAuthnRequests gets a reference to the given bool and assigns it to the SignAuthnRequests field.
func (o *IamSamlSpConnectionAllOf) SetSignAuthnRequests(v bool) {
	o.SignAuthnRequests = &v
}

func (o IamSamlSpConnectionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IdentityProviderSloBinding != nil {
		toSerialize["IdentityProviderSloBinding"] = o.IdentityProviderSloBinding
	}
	if o.IdentityProviderSloUrl != nil {
		toSerialize["IdentityProviderSloUrl"] = o.IdentityProviderSloUrl
	}
	if o.IdentityProviderSsoBinding != nil {
		toSerialize["IdentityProviderSsoBinding"] = o.IdentityProviderSsoBinding
	}
	if o.IdentityProviderSsoUrl != nil {
		toSerialize["IdentityProviderSsoUrl"] = o.IdentityProviderSsoUrl
	}
	if o.IdpCertificateStore != nil {
		toSerialize["IdpCertificateStore"] = o.IdpCertificateStore
	}
	if o.SignAuthnRequests != nil {
		toSerialize["SignAuthnRequests"] = o.SignAuthnRequests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamSamlSpConnectionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamSamlSpConnectionAllOf := _IamSamlSpConnectionAllOf{}

	if err = json.Unmarshal(bytes, &varIamSamlSpConnectionAllOf); err == nil {
		*o = IamSamlSpConnectionAllOf(varIamSamlSpConnectionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IdentityProviderSloBinding")
		delete(additionalProperties, "IdentityProviderSloUrl")
		delete(additionalProperties, "IdentityProviderSsoBinding")
		delete(additionalProperties, "IdentityProviderSsoUrl")
		delete(additionalProperties, "IdpCertificateStore")
		delete(additionalProperties, "SignAuthnRequests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamSamlSpConnectionAllOf struct {
	value *IamSamlSpConnectionAllOf
	isSet bool
}

func (v NullableIamSamlSpConnectionAllOf) Get() *IamSamlSpConnectionAllOf {
	return v.value
}

func (v *NullableIamSamlSpConnectionAllOf) Set(val *IamSamlSpConnectionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSamlSpConnectionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSamlSpConnectionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSamlSpConnectionAllOf(val *IamSamlSpConnectionAllOf) *NullableIamSamlSpConnectionAllOf {
	return &NullableIamSamlSpConnectionAllOf{value: val, isSet: true}
}

func (v NullableIamSamlSpConnectionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSamlSpConnectionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
