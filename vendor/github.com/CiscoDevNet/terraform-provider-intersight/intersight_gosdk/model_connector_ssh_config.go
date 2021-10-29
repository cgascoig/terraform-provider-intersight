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

// ConnectorSshConfig Carries the SSH session details for opening a new connection. Sent by cloud services with the OpenSession message.
type ConnectorSshConfig struct {
	ConnectorBaseMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A jump host for establishing a connection to a server. Plugin will first establish a connection to this server, then create a tunneled connection to the target host.
	JumpHost *string `json:"JumpHost,omitempty"`
	// Optional passphrase if provided while creating the private key.
	Passphrase *string `json:"Passphrase,omitempty"`
	// Password to use in the connection credentials (If empty the private key will be used).
	Password *string `json:"Password,omitempty"`
	// The private key to use in the connection credentials (Optional if password is given).
	Pkey *string `json:"Pkey,omitempty"`
	// The remote server to connect to.
	Target *string `json:"Target,omitempty"`
	// Username for the remote connection.
	User                 *string `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorSshConfig ConnectorSshConfig

// NewConnectorSshConfig instantiates a new ConnectorSshConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorSshConfig(classId string, objectType string) *ConnectorSshConfig {
	this := ConnectorSshConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorSshConfigWithDefaults instantiates a new ConnectorSshConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorSshConfigWithDefaults() *ConnectorSshConfig {
	this := ConnectorSshConfig{}
	var classId string = "connector.SshConfig"
	this.ClassId = classId
	var objectType string = "connector.SshConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorSshConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorSshConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorSshConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorSshConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetJumpHost returns the JumpHost field value if set, zero value otherwise.
func (o *ConnectorSshConfig) GetJumpHost() string {
	if o == nil || o.JumpHost == nil {
		var ret string
		return ret
	}
	return *o.JumpHost
}

// GetJumpHostOk returns a tuple with the JumpHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfig) GetJumpHostOk() (*string, bool) {
	if o == nil || o.JumpHost == nil {
		return nil, false
	}
	return o.JumpHost, true
}

// HasJumpHost returns a boolean if a field has been set.
func (o *ConnectorSshConfig) HasJumpHost() bool {
	if o != nil && o.JumpHost != nil {
		return true
	}

	return false
}

// SetJumpHost gets a reference to the given string and assigns it to the JumpHost field.
func (o *ConnectorSshConfig) SetJumpHost(v string) {
	o.JumpHost = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *ConnectorSshConfig) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfig) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *ConnectorSshConfig) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *ConnectorSshConfig) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ConnectorSshConfig) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfig) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ConnectorSshConfig) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ConnectorSshConfig) SetPassword(v string) {
	o.Password = &v
}

// GetPkey returns the Pkey field value if set, zero value otherwise.
func (o *ConnectorSshConfig) GetPkey() string {
	if o == nil || o.Pkey == nil {
		var ret string
		return ret
	}
	return *o.Pkey
}

// GetPkeyOk returns a tuple with the Pkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfig) GetPkeyOk() (*string, bool) {
	if o == nil || o.Pkey == nil {
		return nil, false
	}
	return o.Pkey, true
}

// HasPkey returns a boolean if a field has been set.
func (o *ConnectorSshConfig) HasPkey() bool {
	if o != nil && o.Pkey != nil {
		return true
	}

	return false
}

// SetPkey gets a reference to the given string and assigns it to the Pkey field.
func (o *ConnectorSshConfig) SetPkey(v string) {
	o.Pkey = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ConnectorSshConfig) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfig) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ConnectorSshConfig) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ConnectorSshConfig) SetTarget(v string) {
	o.Target = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ConnectorSshConfig) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfig) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ConnectorSshConfig) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ConnectorSshConfig) SetUser(v string) {
	o.User = &v
}

func (o ConnectorSshConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorBaseMessage, errConnectorBaseMessage := json.Marshal(o.ConnectorBaseMessage)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	errConnectorBaseMessage = json.Unmarshal([]byte(serializedConnectorBaseMessage), &toSerialize)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.JumpHost != nil {
		toSerialize["JumpHost"] = o.JumpHost
	}
	if o.Passphrase != nil {
		toSerialize["Passphrase"] = o.Passphrase
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Pkey != nil {
		toSerialize["Pkey"] = o.Pkey
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorSshConfig) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorSshConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A jump host for establishing a connection to a server. Plugin will first establish a connection to this server, then create a tunneled connection to the target host.
		JumpHost *string `json:"JumpHost,omitempty"`
		// Optional passphrase if provided while creating the private key.
		Passphrase *string `json:"Passphrase,omitempty"`
		// Password to use in the connection credentials (If empty the private key will be used).
		Password *string `json:"Password,omitempty"`
		// The private key to use in the connection credentials (Optional if password is given).
		Pkey *string `json:"Pkey,omitempty"`
		// The remote server to connect to.
		Target *string `json:"Target,omitempty"`
		// Username for the remote connection.
		User *string `json:"User,omitempty"`
	}

	varConnectorSshConfigWithoutEmbeddedStruct := ConnectorSshConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorSshConfigWithoutEmbeddedStruct)
	if err == nil {
		varConnectorSshConfig := _ConnectorSshConfig{}
		varConnectorSshConfig.ClassId = varConnectorSshConfigWithoutEmbeddedStruct.ClassId
		varConnectorSshConfig.ObjectType = varConnectorSshConfigWithoutEmbeddedStruct.ObjectType
		varConnectorSshConfig.JumpHost = varConnectorSshConfigWithoutEmbeddedStruct.JumpHost
		varConnectorSshConfig.Passphrase = varConnectorSshConfigWithoutEmbeddedStruct.Passphrase
		varConnectorSshConfig.Password = varConnectorSshConfigWithoutEmbeddedStruct.Password
		varConnectorSshConfig.Pkey = varConnectorSshConfigWithoutEmbeddedStruct.Pkey
		varConnectorSshConfig.Target = varConnectorSshConfigWithoutEmbeddedStruct.Target
		varConnectorSshConfig.User = varConnectorSshConfigWithoutEmbeddedStruct.User
		*o = ConnectorSshConfig(varConnectorSshConfig)
	} else {
		return err
	}

	varConnectorSshConfig := _ConnectorSshConfig{}

	err = json.Unmarshal(bytes, &varConnectorSshConfig)
	if err == nil {
		o.ConnectorBaseMessage = varConnectorSshConfig.ConnectorBaseMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "JumpHost")
		delete(additionalProperties, "Passphrase")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Pkey")
		delete(additionalProperties, "Target")
		delete(additionalProperties, "User")

		// remove fields from embedded structs
		reflectConnectorBaseMessage := reflect.ValueOf(o.ConnectorBaseMessage)
		for i := 0; i < reflectConnectorBaseMessage.Type().NumField(); i++ {
			t := reflectConnectorBaseMessage.Type().Field(i)

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

type NullableConnectorSshConfig struct {
	value *ConnectorSshConfig
	isSet bool
}

func (v NullableConnectorSshConfig) Get() *ConnectorSshConfig {
	return v.value
}

func (v *NullableConnectorSshConfig) Set(val *ConnectorSshConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorSshConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorSshConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorSshConfig(val *ConnectorSshConfig) *NullableConnectorSshConfig {
	return &NullableConnectorSshConfig{value: val, isSet: true}
}

func (v NullableConnectorSshConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorSshConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
