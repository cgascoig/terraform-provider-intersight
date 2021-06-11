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
)

// KubernetesClusterCertificateConfigurationAllOf Definition of the list of properties defined in 'kubernetes.ClusterCertificateConfiguration', excluding properties defined in parent classes.
type KubernetesClusterCertificateConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Certificate for the Kubernetes API server.
	CaCert *string `json:"CaCert,omitempty"`
	// Private Key for the Kubernetes API server.
	CaKey *string `json:"CaKey,omitempty"`
	// Certificate for the etcd cluster.
	EtcdCert          *string  `json:"EtcdCert,omitempty"`
	EtcdEncryptionKey []string `json:"EtcdEncryptionKey,omitempty"`
	// Private key for the etcd cluster.
	EtcdKey *string `json:"EtcdKey,omitempty"`
	// Certificate for the front proxy to support Kubernetes API aggregators.
	FrontProxyCert *string `json:"FrontProxyCert,omitempty"`
	// Private key for the front proxy to support Kubernetes API aggregators.
	FrontProxyKey *string `json:"FrontProxyKey,omitempty"`
	// Service account private key used by Kubernetes Token Controller to sign generated service account tokens.
	SaPrivateKey *string `json:"SaPrivateKey,omitempty"`
	// Service account public key used by Kubernetes API Server to validate service account tokens generated by the Token Controller.
	SaPublicKey          *string `json:"SaPublicKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterCertificateConfigurationAllOf KubernetesClusterCertificateConfigurationAllOf

// NewKubernetesClusterCertificateConfigurationAllOf instantiates a new KubernetesClusterCertificateConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterCertificateConfigurationAllOf(classId string, objectType string) *KubernetesClusterCertificateConfigurationAllOf {
	this := KubernetesClusterCertificateConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesClusterCertificateConfigurationAllOfWithDefaults instantiates a new KubernetesClusterCertificateConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterCertificateConfigurationAllOfWithDefaults() *KubernetesClusterCertificateConfigurationAllOf {
	this := KubernetesClusterCertificateConfigurationAllOf{}
	var classId string = "kubernetes.ClusterCertificateConfiguration"
	this.ClassId = classId
	var objectType string = "kubernetes.ClusterCertificateConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesClusterCertificateConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesClusterCertificateConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesClusterCertificateConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesClusterCertificateConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCaCert returns the CaCert field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetCaCert() string {
	if o == nil || o.CaCert == nil {
		var ret string
		return ret
	}
	return *o.CaCert
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetCaCertOk() (*string, bool) {
	if o == nil || o.CaCert == nil {
		return nil, false
	}
	return o.CaCert, true
}

// HasCaCert returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) HasCaCert() bool {
	if o != nil && o.CaCert != nil {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given string and assigns it to the CaCert field.
func (o *KubernetesClusterCertificateConfigurationAllOf) SetCaCert(v string) {
	o.CaCert = &v
}

// GetCaKey returns the CaKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetCaKey() string {
	if o == nil || o.CaKey == nil {
		var ret string
		return ret
	}
	return *o.CaKey
}

// GetCaKeyOk returns a tuple with the CaKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetCaKeyOk() (*string, bool) {
	if o == nil || o.CaKey == nil {
		return nil, false
	}
	return o.CaKey, true
}

// HasCaKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) HasCaKey() bool {
	if o != nil && o.CaKey != nil {
		return true
	}

	return false
}

// SetCaKey gets a reference to the given string and assigns it to the CaKey field.
func (o *KubernetesClusterCertificateConfigurationAllOf) SetCaKey(v string) {
	o.CaKey = &v
}

// GetEtcdCert returns the EtcdCert field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdCert() string {
	if o == nil || o.EtcdCert == nil {
		var ret string
		return ret
	}
	return *o.EtcdCert
}

// GetEtcdCertOk returns a tuple with the EtcdCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdCertOk() (*string, bool) {
	if o == nil || o.EtcdCert == nil {
		return nil, false
	}
	return o.EtcdCert, true
}

// HasEtcdCert returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) HasEtcdCert() bool {
	if o != nil && o.EtcdCert != nil {
		return true
	}

	return false
}

// SetEtcdCert gets a reference to the given string and assigns it to the EtcdCert field.
func (o *KubernetesClusterCertificateConfigurationAllOf) SetEtcdCert(v string) {
	o.EtcdCert = &v
}

// GetEtcdEncryptionKey returns the EtcdEncryptionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdEncryptionKey() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EtcdEncryptionKey
}

// GetEtcdEncryptionKeyOk returns a tuple with the EtcdEncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdEncryptionKeyOk() (*[]string, bool) {
	if o == nil || o.EtcdEncryptionKey == nil {
		return nil, false
	}
	return &o.EtcdEncryptionKey, true
}

// HasEtcdEncryptionKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) HasEtcdEncryptionKey() bool {
	if o != nil && o.EtcdEncryptionKey != nil {
		return true
	}

	return false
}

// SetEtcdEncryptionKey gets a reference to the given []string and assigns it to the EtcdEncryptionKey field.
func (o *KubernetesClusterCertificateConfigurationAllOf) SetEtcdEncryptionKey(v []string) {
	o.EtcdEncryptionKey = v
}

// GetEtcdKey returns the EtcdKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdKey() string {
	if o == nil || o.EtcdKey == nil {
		var ret string
		return ret
	}
	return *o.EtcdKey
}

// GetEtcdKeyOk returns a tuple with the EtcdKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdKeyOk() (*string, bool) {
	if o == nil || o.EtcdKey == nil {
		return nil, false
	}
	return o.EtcdKey, true
}

// HasEtcdKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) HasEtcdKey() bool {
	if o != nil && o.EtcdKey != nil {
		return true
	}

	return false
}

// SetEtcdKey gets a reference to the given string and assigns it to the EtcdKey field.
func (o *KubernetesClusterCertificateConfigurationAllOf) SetEtcdKey(v string) {
	o.EtcdKey = &v
}

// GetFrontProxyCert returns the FrontProxyCert field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetFrontProxyCert() string {
	if o == nil || o.FrontProxyCert == nil {
		var ret string
		return ret
	}
	return *o.FrontProxyCert
}

// GetFrontProxyCertOk returns a tuple with the FrontProxyCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetFrontProxyCertOk() (*string, bool) {
	if o == nil || o.FrontProxyCert == nil {
		return nil, false
	}
	return o.FrontProxyCert, true
}

// HasFrontProxyCert returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) HasFrontProxyCert() bool {
	if o != nil && o.FrontProxyCert != nil {
		return true
	}

	return false
}

// SetFrontProxyCert gets a reference to the given string and assigns it to the FrontProxyCert field.
func (o *KubernetesClusterCertificateConfigurationAllOf) SetFrontProxyCert(v string) {
	o.FrontProxyCert = &v
}

// GetFrontProxyKey returns the FrontProxyKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetFrontProxyKey() string {
	if o == nil || o.FrontProxyKey == nil {
		var ret string
		return ret
	}
	return *o.FrontProxyKey
}

// GetFrontProxyKeyOk returns a tuple with the FrontProxyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetFrontProxyKeyOk() (*string, bool) {
	if o == nil || o.FrontProxyKey == nil {
		return nil, false
	}
	return o.FrontProxyKey, true
}

// HasFrontProxyKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) HasFrontProxyKey() bool {
	if o != nil && o.FrontProxyKey != nil {
		return true
	}

	return false
}

// SetFrontProxyKey gets a reference to the given string and assigns it to the FrontProxyKey field.
func (o *KubernetesClusterCertificateConfigurationAllOf) SetFrontProxyKey(v string) {
	o.FrontProxyKey = &v
}

// GetSaPrivateKey returns the SaPrivateKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetSaPrivateKey() string {
	if o == nil || o.SaPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.SaPrivateKey
}

// GetSaPrivateKeyOk returns a tuple with the SaPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetSaPrivateKeyOk() (*string, bool) {
	if o == nil || o.SaPrivateKey == nil {
		return nil, false
	}
	return o.SaPrivateKey, true
}

// HasSaPrivateKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) HasSaPrivateKey() bool {
	if o != nil && o.SaPrivateKey != nil {
		return true
	}

	return false
}

// SetSaPrivateKey gets a reference to the given string and assigns it to the SaPrivateKey field.
func (o *KubernetesClusterCertificateConfigurationAllOf) SetSaPrivateKey(v string) {
	o.SaPrivateKey = &v
}

// GetSaPublicKey returns the SaPublicKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetSaPublicKey() string {
	if o == nil || o.SaPublicKey == nil {
		var ret string
		return ret
	}
	return *o.SaPublicKey
}

// GetSaPublicKeyOk returns a tuple with the SaPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) GetSaPublicKeyOk() (*string, bool) {
	if o == nil || o.SaPublicKey == nil {
		return nil, false
	}
	return o.SaPublicKey, true
}

// HasSaPublicKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfigurationAllOf) HasSaPublicKey() bool {
	if o != nil && o.SaPublicKey != nil {
		return true
	}

	return false
}

// SetSaPublicKey gets a reference to the given string and assigns it to the SaPublicKey field.
func (o *KubernetesClusterCertificateConfigurationAllOf) SetSaPublicKey(v string) {
	o.SaPublicKey = &v
}

func (o KubernetesClusterCertificateConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CaCert != nil {
		toSerialize["CaCert"] = o.CaCert
	}
	if o.CaKey != nil {
		toSerialize["CaKey"] = o.CaKey
	}
	if o.EtcdCert != nil {
		toSerialize["EtcdCert"] = o.EtcdCert
	}
	if o.EtcdEncryptionKey != nil {
		toSerialize["EtcdEncryptionKey"] = o.EtcdEncryptionKey
	}
	if o.EtcdKey != nil {
		toSerialize["EtcdKey"] = o.EtcdKey
	}
	if o.FrontProxyCert != nil {
		toSerialize["FrontProxyCert"] = o.FrontProxyCert
	}
	if o.FrontProxyKey != nil {
		toSerialize["FrontProxyKey"] = o.FrontProxyKey
	}
	if o.SaPrivateKey != nil {
		toSerialize["SaPrivateKey"] = o.SaPrivateKey
	}
	if o.SaPublicKey != nil {
		toSerialize["SaPublicKey"] = o.SaPublicKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesClusterCertificateConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesClusterCertificateConfigurationAllOf := _KubernetesClusterCertificateConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesClusterCertificateConfigurationAllOf); err == nil {
		*o = KubernetesClusterCertificateConfigurationAllOf(varKubernetesClusterCertificateConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CaCert")
		delete(additionalProperties, "CaKey")
		delete(additionalProperties, "EtcdCert")
		delete(additionalProperties, "EtcdEncryptionKey")
		delete(additionalProperties, "EtcdKey")
		delete(additionalProperties, "FrontProxyCert")
		delete(additionalProperties, "FrontProxyKey")
		delete(additionalProperties, "SaPrivateKey")
		delete(additionalProperties, "SaPublicKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesClusterCertificateConfigurationAllOf struct {
	value *KubernetesClusterCertificateConfigurationAllOf
	isSet bool
}

func (v NullableKubernetesClusterCertificateConfigurationAllOf) Get() *KubernetesClusterCertificateConfigurationAllOf {
	return v.value
}

func (v *NullableKubernetesClusterCertificateConfigurationAllOf) Set(val *KubernetesClusterCertificateConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterCertificateConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterCertificateConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterCertificateConfigurationAllOf(val *KubernetesClusterCertificateConfigurationAllOf) *NullableKubernetesClusterCertificateConfigurationAllOf {
	return &NullableKubernetesClusterCertificateConfigurationAllOf{value: val, isSet: true}
}

func (v NullableKubernetesClusterCertificateConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterCertificateConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
