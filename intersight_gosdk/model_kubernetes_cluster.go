/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-24T06:47:07Z.
 *
 * API version: 1.0.9-3824
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// KubernetesCluster Inventories a Kubernetes cluster state. A Cluster object is automatically created when a Kubernetes API server is configured for a cluster.
type KubernetesCluster struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of the endpoint connection of this Kubernetes cluster. * `` - The target details have been persisted but Intersight has not yet attempted to connect to the target. * `Connected` - Intersight is able to establish a connection to the target and initiate management activities. * `NotConnected` - Intersight is unable to establish a connection to the target. * `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established. * `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// Kubeconfig for the cluster to collect inventory for.
	KubeConfig *string `json:"KubeConfig,omitempty"`
	// Name of the Kubernetes cluster.
	Name               *string                               `json:"Name,omitempty"`
	Var0ClusterProfile *KubernetesClusterProfileRelationship `json:"_0_ClusterProfile,omitempty"`
	Organization       *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to assetDeviceRegistration resources.
	RegisteredDevices    []AssetDeviceRegistrationRelationship `json:"RegisteredDevices,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesCluster KubernetesCluster

// NewKubernetesCluster instantiates a new KubernetesCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCluster(classId string, objectType string) *KubernetesCluster {
	this := KubernetesCluster{}
	this.ClassId = classId
	this.ObjectType = objectType
	var connectionStatus string = ""
	this.ConnectionStatus = &connectionStatus
	return &this
}

// NewKubernetesClusterWithDefaults instantiates a new KubernetesCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterWithDefaults() *KubernetesCluster {
	this := KubernetesCluster{}
	var classId string = "kubernetes.Cluster"
	this.ClassId = classId
	var objectType string = "kubernetes.Cluster"
	this.ObjectType = objectType
	var connectionStatus string = ""
	this.ConnectionStatus = &connectionStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesCluster) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesCluster) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesCluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesCluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *KubernetesCluster) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *KubernetesCluster) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *KubernetesCluster) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetKubeConfig returns the KubeConfig field value if set, zero value otherwise.
func (o *KubernetesCluster) GetKubeConfig() string {
	if o == nil || o.KubeConfig == nil {
		var ret string
		return ret
	}
	return *o.KubeConfig
}

// GetKubeConfigOk returns a tuple with the KubeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetKubeConfigOk() (*string, bool) {
	if o == nil || o.KubeConfig == nil {
		return nil, false
	}
	return o.KubeConfig, true
}

// HasKubeConfig returns a boolean if a field has been set.
func (o *KubernetesCluster) HasKubeConfig() bool {
	if o != nil && o.KubeConfig != nil {
		return true
	}

	return false
}

// SetKubeConfig gets a reference to the given string and assigns it to the KubeConfig field.
func (o *KubernetesCluster) SetKubeConfig(v string) {
	o.KubeConfig = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesCluster) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesCluster) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesCluster) SetName(v string) {
	o.Name = &v
}

// GetVar0ClusterProfile returns the Var0ClusterProfile field value if set, zero value otherwise.
func (o *KubernetesCluster) GetVar0ClusterProfile() KubernetesClusterProfileRelationship {
	if o == nil || o.Var0ClusterProfile == nil {
		var ret KubernetesClusterProfileRelationship
		return ret
	}
	return *o.Var0ClusterProfile
}

// GetVar0ClusterProfileOk returns a tuple with the Var0ClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetVar0ClusterProfileOk() (*KubernetesClusterProfileRelationship, bool) {
	if o == nil || o.Var0ClusterProfile == nil {
		return nil, false
	}
	return o.Var0ClusterProfile, true
}

// HasVar0ClusterProfile returns a boolean if a field has been set.
func (o *KubernetesCluster) HasVar0ClusterProfile() bool {
	if o != nil && o.Var0ClusterProfile != nil {
		return true
	}

	return false
}

// SetVar0ClusterProfile gets a reference to the given KubernetesClusterProfileRelationship and assigns it to the Var0ClusterProfile field.
func (o *KubernetesCluster) SetVar0ClusterProfile(v KubernetesClusterProfileRelationship) {
	o.Var0ClusterProfile = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesCluster) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesCluster) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesCluster) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetRegisteredDevices returns the RegisteredDevices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesCluster) GetRegisteredDevices() []AssetDeviceRegistrationRelationship {
	if o == nil {
		var ret []AssetDeviceRegistrationRelationship
		return ret
	}
	return o.RegisteredDevices
}

// GetRegisteredDevicesOk returns a tuple with the RegisteredDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCluster) GetRegisteredDevicesOk() (*[]AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevices == nil {
		return nil, false
	}
	return &o.RegisteredDevices, true
}

// HasRegisteredDevices returns a boolean if a field has been set.
func (o *KubernetesCluster) HasRegisteredDevices() bool {
	if o != nil && o.RegisteredDevices != nil {
		return true
	}

	return false
}

// SetRegisteredDevices gets a reference to the given []AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevices field.
func (o *KubernetesCluster) SetRegisteredDevices(v []AssetDeviceRegistrationRelationship) {
	o.RegisteredDevices = v
}

func (o KubernetesCluster) MarshalJSON() ([]byte, error) {
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
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.KubeConfig != nil {
		toSerialize["KubeConfig"] = o.KubeConfig
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Var0ClusterProfile != nil {
		toSerialize["_0_ClusterProfile"] = o.Var0ClusterProfile
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.RegisteredDevices != nil {
		toSerialize["RegisteredDevices"] = o.RegisteredDevices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesCluster) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesClusterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of the endpoint connection of this Kubernetes cluster. * `` - The target details have been persisted but Intersight has not yet attempted to connect to the target. * `Connected` - Intersight is able to establish a connection to the target and initiate management activities. * `NotConnected` - Intersight is unable to establish a connection to the target. * `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established. * `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
		// Kubeconfig for the cluster to collect inventory for.
		KubeConfig *string `json:"KubeConfig,omitempty"`
		// Name of the Kubernetes cluster.
		Name               *string                               `json:"Name,omitempty"`
		Var0ClusterProfile *KubernetesClusterProfileRelationship `json:"_0_ClusterProfile,omitempty"`
		Organization       *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to assetDeviceRegistration resources.
		RegisteredDevices []AssetDeviceRegistrationRelationship `json:"RegisteredDevices,omitempty"`
	}

	varKubernetesClusterWithoutEmbeddedStruct := KubernetesClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesClusterWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesCluster := _KubernetesCluster{}
		varKubernetesCluster.ClassId = varKubernetesClusterWithoutEmbeddedStruct.ClassId
		varKubernetesCluster.ObjectType = varKubernetesClusterWithoutEmbeddedStruct.ObjectType
		varKubernetesCluster.ConnectionStatus = varKubernetesClusterWithoutEmbeddedStruct.ConnectionStatus
		varKubernetesCluster.KubeConfig = varKubernetesClusterWithoutEmbeddedStruct.KubeConfig
		varKubernetesCluster.Name = varKubernetesClusterWithoutEmbeddedStruct.Name
		varKubernetesCluster.Var0ClusterProfile = varKubernetesClusterWithoutEmbeddedStruct.Var0ClusterProfile
		varKubernetesCluster.Organization = varKubernetesClusterWithoutEmbeddedStruct.Organization
		varKubernetesCluster.RegisteredDevices = varKubernetesClusterWithoutEmbeddedStruct.RegisteredDevices
		*o = KubernetesCluster(varKubernetesCluster)
	} else {
		return err
	}

	varKubernetesCluster := _KubernetesCluster{}

	err = json.Unmarshal(bytes, &varKubernetesCluster)
	if err == nil {
		o.MoBaseMo = varKubernetesCluster.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "KubeConfig")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "_0_ClusterProfile")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RegisteredDevices")

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

type NullableKubernetesCluster struct {
	value *KubernetesCluster
	isSet bool
}

func (v NullableKubernetesCluster) Get() *KubernetesCluster {
	return v.value
}

func (v *NullableKubernetesCluster) Set(val *KubernetesCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCluster(val *KubernetesCluster) *NullableKubernetesCluster {
	return &NullableKubernetesCluster{value: val, isSet: true}
}

func (v NullableKubernetesCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
