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

// HyperflexWitnessConfiguration The witness configuration of the HyperFlex cluster. Cisco HyperFlex Edge deployments use a witness in their HA arbitration implementations to continue functioning in case of single node failure or network partitions between nodes. HyperFlex administrators have the option to use two different types of witnesses, the Cisco Intersight Invisible Cloud Witness or a local witness deployed within the customers environment. The type of witness and its corresponding configuration is configured within the HX Connect Device Connector UI component and configuration is stored locally on the device. Changes made locally in the device are reflected in Intersight on this object.
type HyperflexWitnessConfiguration struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The detailed connection error to the external witness. Empty if status is connected.
	ConnectionError *string `json:"ConnectionError,omitempty"`
	// Custom witness has been configured by user.
	CustomWitnessEnabled *bool `json:"CustomWitnessEnabled,omitempty"`
	// The fingerprint of the witness server, identifies the revision of the witness servers database. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty.
	Fingerprint *string `json:"Fingerprint,omitempty"`
	// Status of the devices connection to the witness. Device will report status as either 'Connected' or 'NotConnected'.
	Status *string `json:"Status,omitempty"`
	// The version of the custom witness server. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty.
	Version *string `json:"Version,omitempty"`
	// URL of the witness endpoint, including IP/host and path. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty.
	WitnessUrl           *string                       `json:"WitnessUrl,omitempty"`
	Cluster              *HyperflexClusterRelationship `json:"Cluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexWitnessConfiguration HyperflexWitnessConfiguration

// NewHyperflexWitnessConfiguration instantiates a new HyperflexWitnessConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexWitnessConfiguration(classId string, objectType string) *HyperflexWitnessConfiguration {
	this := HyperflexWitnessConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexWitnessConfigurationWithDefaults instantiates a new HyperflexWitnessConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexWitnessConfigurationWithDefaults() *HyperflexWitnessConfiguration {
	this := HyperflexWitnessConfiguration{}
	var classId string = "hyperflex.WitnessConfiguration"
	this.ClassId = classId
	var objectType string = "hyperflex.WitnessConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexWitnessConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexWitnessConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexWitnessConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexWitnessConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexWitnessConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexWitnessConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionError returns the ConnectionError field value if set, zero value otherwise.
func (o *HyperflexWitnessConfiguration) GetConnectionError() string {
	if o == nil || o.ConnectionError == nil {
		var ret string
		return ret
	}
	return *o.ConnectionError
}

// GetConnectionErrorOk returns a tuple with the ConnectionError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexWitnessConfiguration) GetConnectionErrorOk() (*string, bool) {
	if o == nil || o.ConnectionError == nil {
		return nil, false
	}
	return o.ConnectionError, true
}

// HasConnectionError returns a boolean if a field has been set.
func (o *HyperflexWitnessConfiguration) HasConnectionError() bool {
	if o != nil && o.ConnectionError != nil {
		return true
	}

	return false
}

// SetConnectionError gets a reference to the given string and assigns it to the ConnectionError field.
func (o *HyperflexWitnessConfiguration) SetConnectionError(v string) {
	o.ConnectionError = &v
}

// GetCustomWitnessEnabled returns the CustomWitnessEnabled field value if set, zero value otherwise.
func (o *HyperflexWitnessConfiguration) GetCustomWitnessEnabled() bool {
	if o == nil || o.CustomWitnessEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CustomWitnessEnabled
}

// GetCustomWitnessEnabledOk returns a tuple with the CustomWitnessEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexWitnessConfiguration) GetCustomWitnessEnabledOk() (*bool, bool) {
	if o == nil || o.CustomWitnessEnabled == nil {
		return nil, false
	}
	return o.CustomWitnessEnabled, true
}

// HasCustomWitnessEnabled returns a boolean if a field has been set.
func (o *HyperflexWitnessConfiguration) HasCustomWitnessEnabled() bool {
	if o != nil && o.CustomWitnessEnabled != nil {
		return true
	}

	return false
}

// SetCustomWitnessEnabled gets a reference to the given bool and assigns it to the CustomWitnessEnabled field.
func (o *HyperflexWitnessConfiguration) SetCustomWitnessEnabled(v bool) {
	o.CustomWitnessEnabled = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *HyperflexWitnessConfiguration) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexWitnessConfiguration) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *HyperflexWitnessConfiguration) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *HyperflexWitnessConfiguration) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexWitnessConfiguration) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexWitnessConfiguration) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexWitnessConfiguration) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexWitnessConfiguration) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexWitnessConfiguration) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexWitnessConfiguration) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexWitnessConfiguration) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexWitnessConfiguration) SetVersion(v string) {
	o.Version = &v
}

// GetWitnessUrl returns the WitnessUrl field value if set, zero value otherwise.
func (o *HyperflexWitnessConfiguration) GetWitnessUrl() string {
	if o == nil || o.WitnessUrl == nil {
		var ret string
		return ret
	}
	return *o.WitnessUrl
}

// GetWitnessUrlOk returns a tuple with the WitnessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexWitnessConfiguration) GetWitnessUrlOk() (*string, bool) {
	if o == nil || o.WitnessUrl == nil {
		return nil, false
	}
	return o.WitnessUrl, true
}

// HasWitnessUrl returns a boolean if a field has been set.
func (o *HyperflexWitnessConfiguration) HasWitnessUrl() bool {
	if o != nil && o.WitnessUrl != nil {
		return true
	}

	return false
}

// SetWitnessUrl gets a reference to the given string and assigns it to the WitnessUrl field.
func (o *HyperflexWitnessConfiguration) SetWitnessUrl(v string) {
	o.WitnessUrl = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexWitnessConfiguration) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexWitnessConfiguration) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexWitnessConfiguration) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexWitnessConfiguration) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

func (o HyperflexWitnessConfiguration) MarshalJSON() ([]byte, error) {
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
	if o.ConnectionError != nil {
		toSerialize["ConnectionError"] = o.ConnectionError
	}
	if o.CustomWitnessEnabled != nil {
		toSerialize["CustomWitnessEnabled"] = o.CustomWitnessEnabled
	}
	if o.Fingerprint != nil {
		toSerialize["Fingerprint"] = o.Fingerprint
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.WitnessUrl != nil {
		toSerialize["WitnessUrl"] = o.WitnessUrl
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexWitnessConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexWitnessConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The detailed connection error to the external witness. Empty if status is connected.
		ConnectionError *string `json:"ConnectionError,omitempty"`
		// Custom witness has been configured by user.
		CustomWitnessEnabled *bool `json:"CustomWitnessEnabled,omitempty"`
		// The fingerprint of the witness server, identifies the revision of the witness servers database. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty.
		Fingerprint *string `json:"Fingerprint,omitempty"`
		// Status of the devices connection to the witness. Device will report status as either 'Connected' or 'NotConnected'.
		Status *string `json:"Status,omitempty"`
		// The version of the custom witness server. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty.
		Version *string `json:"Version,omitempty"`
		// URL of the witness endpoint, including IP/host and path. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty.
		WitnessUrl *string                       `json:"WitnessUrl,omitempty"`
		Cluster    *HyperflexClusterRelationship `json:"Cluster,omitempty"`
	}

	varHyperflexWitnessConfigurationWithoutEmbeddedStruct := HyperflexWitnessConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexWitnessConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexWitnessConfiguration := _HyperflexWitnessConfiguration{}
		varHyperflexWitnessConfiguration.ClassId = varHyperflexWitnessConfigurationWithoutEmbeddedStruct.ClassId
		varHyperflexWitnessConfiguration.ObjectType = varHyperflexWitnessConfigurationWithoutEmbeddedStruct.ObjectType
		varHyperflexWitnessConfiguration.ConnectionError = varHyperflexWitnessConfigurationWithoutEmbeddedStruct.ConnectionError
		varHyperflexWitnessConfiguration.CustomWitnessEnabled = varHyperflexWitnessConfigurationWithoutEmbeddedStruct.CustomWitnessEnabled
		varHyperflexWitnessConfiguration.Fingerprint = varHyperflexWitnessConfigurationWithoutEmbeddedStruct.Fingerprint
		varHyperflexWitnessConfiguration.Status = varHyperflexWitnessConfigurationWithoutEmbeddedStruct.Status
		varHyperflexWitnessConfiguration.Version = varHyperflexWitnessConfigurationWithoutEmbeddedStruct.Version
		varHyperflexWitnessConfiguration.WitnessUrl = varHyperflexWitnessConfigurationWithoutEmbeddedStruct.WitnessUrl
		varHyperflexWitnessConfiguration.Cluster = varHyperflexWitnessConfigurationWithoutEmbeddedStruct.Cluster
		*o = HyperflexWitnessConfiguration(varHyperflexWitnessConfiguration)
	} else {
		return err
	}

	varHyperflexWitnessConfiguration := _HyperflexWitnessConfiguration{}

	err = json.Unmarshal(bytes, &varHyperflexWitnessConfiguration)
	if err == nil {
		o.MoBaseMo = varHyperflexWitnessConfiguration.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionError")
		delete(additionalProperties, "CustomWitnessEnabled")
		delete(additionalProperties, "Fingerprint")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "WitnessUrl")
		delete(additionalProperties, "Cluster")

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

type NullableHyperflexWitnessConfiguration struct {
	value *HyperflexWitnessConfiguration
	isSet bool
}

func (v NullableHyperflexWitnessConfiguration) Get() *HyperflexWitnessConfiguration {
	return v.value
}

func (v *NullableHyperflexWitnessConfiguration) Set(val *HyperflexWitnessConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexWitnessConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexWitnessConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexWitnessConfiguration(val *HyperflexWitnessConfiguration) *NullableHyperflexWitnessConfiguration {
	return &NullableHyperflexWitnessConfiguration{value: val, isSet: true}
}

func (v NullableHyperflexWitnessConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexWitnessConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
