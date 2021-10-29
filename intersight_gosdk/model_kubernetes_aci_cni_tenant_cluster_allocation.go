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

// KubernetesAciCniTenantClusterAllocation Internally generated parameter allocations for a k8s cluster using a particular ACI CNI profile.
type KubernetesAciCniTenantClusterAllocation struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// CIDR allocated for ACI node service IPs in this tenant cluster.
	NodeSvcIpSubnet *string `json:"NodeSvcIpSubnet,omitempty"`
	// CIDR allocated for pod IPs in this tenant cluster.
	PodIpSubnet *string `json:"PodIpSubnet,omitempty"`
	// End of VLAN range allocated to this tenant cluster.
	VlanEnd *string `json:"VlanEnd,omitempty"`
	// Start of VLAN range allocated to this tenant cluster.
	VlanStart            *string                               `json:"VlanStart,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAciCniTenantClusterAllocation KubernetesAciCniTenantClusterAllocation

// NewKubernetesAciCniTenantClusterAllocation instantiates a new KubernetesAciCniTenantClusterAllocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAciCniTenantClusterAllocation(classId string, objectType string) *KubernetesAciCniTenantClusterAllocation {
	this := KubernetesAciCniTenantClusterAllocation{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAciCniTenantClusterAllocationWithDefaults instantiates a new KubernetesAciCniTenantClusterAllocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAciCniTenantClusterAllocationWithDefaults() *KubernetesAciCniTenantClusterAllocation {
	this := KubernetesAciCniTenantClusterAllocation{}
	var classId string = "kubernetes.AciCniTenantClusterAllocation"
	this.ClassId = classId
	var objectType string = "kubernetes.AciCniTenantClusterAllocation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAciCniTenantClusterAllocation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAciCniTenantClusterAllocation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAciCniTenantClusterAllocation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAciCniTenantClusterAllocation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNodeSvcIpSubnet returns the NodeSvcIpSubnet field value if set, zero value otherwise.
func (o *KubernetesAciCniTenantClusterAllocation) GetNodeSvcIpSubnet() string {
	if o == nil || o.NodeSvcIpSubnet == nil {
		var ret string
		return ret
	}
	return *o.NodeSvcIpSubnet
}

// GetNodeSvcIpSubnetOk returns a tuple with the NodeSvcIpSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocation) GetNodeSvcIpSubnetOk() (*string, bool) {
	if o == nil || o.NodeSvcIpSubnet == nil {
		return nil, false
	}
	return o.NodeSvcIpSubnet, true
}

// HasNodeSvcIpSubnet returns a boolean if a field has been set.
func (o *KubernetesAciCniTenantClusterAllocation) HasNodeSvcIpSubnet() bool {
	if o != nil && o.NodeSvcIpSubnet != nil {
		return true
	}

	return false
}

// SetNodeSvcIpSubnet gets a reference to the given string and assigns it to the NodeSvcIpSubnet field.
func (o *KubernetesAciCniTenantClusterAllocation) SetNodeSvcIpSubnet(v string) {
	o.NodeSvcIpSubnet = &v
}

// GetPodIpSubnet returns the PodIpSubnet field value if set, zero value otherwise.
func (o *KubernetesAciCniTenantClusterAllocation) GetPodIpSubnet() string {
	if o == nil || o.PodIpSubnet == nil {
		var ret string
		return ret
	}
	return *o.PodIpSubnet
}

// GetPodIpSubnetOk returns a tuple with the PodIpSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocation) GetPodIpSubnetOk() (*string, bool) {
	if o == nil || o.PodIpSubnet == nil {
		return nil, false
	}
	return o.PodIpSubnet, true
}

// HasPodIpSubnet returns a boolean if a field has been set.
func (o *KubernetesAciCniTenantClusterAllocation) HasPodIpSubnet() bool {
	if o != nil && o.PodIpSubnet != nil {
		return true
	}

	return false
}

// SetPodIpSubnet gets a reference to the given string and assigns it to the PodIpSubnet field.
func (o *KubernetesAciCniTenantClusterAllocation) SetPodIpSubnet(v string) {
	o.PodIpSubnet = &v
}

// GetVlanEnd returns the VlanEnd field value if set, zero value otherwise.
func (o *KubernetesAciCniTenantClusterAllocation) GetVlanEnd() string {
	if o == nil || o.VlanEnd == nil {
		var ret string
		return ret
	}
	return *o.VlanEnd
}

// GetVlanEndOk returns a tuple with the VlanEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocation) GetVlanEndOk() (*string, bool) {
	if o == nil || o.VlanEnd == nil {
		return nil, false
	}
	return o.VlanEnd, true
}

// HasVlanEnd returns a boolean if a field has been set.
func (o *KubernetesAciCniTenantClusterAllocation) HasVlanEnd() bool {
	if o != nil && o.VlanEnd != nil {
		return true
	}

	return false
}

// SetVlanEnd gets a reference to the given string and assigns it to the VlanEnd field.
func (o *KubernetesAciCniTenantClusterAllocation) SetVlanEnd(v string) {
	o.VlanEnd = &v
}

// GetVlanStart returns the VlanStart field value if set, zero value otherwise.
func (o *KubernetesAciCniTenantClusterAllocation) GetVlanStart() string {
	if o == nil || o.VlanStart == nil {
		var ret string
		return ret
	}
	return *o.VlanStart
}

// GetVlanStartOk returns a tuple with the VlanStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocation) GetVlanStartOk() (*string, bool) {
	if o == nil || o.VlanStart == nil {
		return nil, false
	}
	return o.VlanStart, true
}

// HasVlanStart returns a boolean if a field has been set.
func (o *KubernetesAciCniTenantClusterAllocation) HasVlanStart() bool {
	if o != nil && o.VlanStart != nil {
		return true
	}

	return false
}

// SetVlanStart gets a reference to the given string and assigns it to the VlanStart field.
func (o *KubernetesAciCniTenantClusterAllocation) SetVlanStart(v string) {
	o.VlanStart = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesAciCniTenantClusterAllocation) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesAciCniTenantClusterAllocation) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesAciCniTenantClusterAllocation) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesAciCniTenantClusterAllocation) MarshalJSON() ([]byte, error) {
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
	if o.NodeSvcIpSubnet != nil {
		toSerialize["NodeSvcIpSubnet"] = o.NodeSvcIpSubnet
	}
	if o.PodIpSubnet != nil {
		toSerialize["PodIpSubnet"] = o.PodIpSubnet
	}
	if o.VlanEnd != nil {
		toSerialize["VlanEnd"] = o.VlanEnd
	}
	if o.VlanStart != nil {
		toSerialize["VlanStart"] = o.VlanStart
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAciCniTenantClusterAllocation) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesAciCniTenantClusterAllocationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// CIDR allocated for ACI node service IPs in this tenant cluster.
		NodeSvcIpSubnet *string `json:"NodeSvcIpSubnet,omitempty"`
		// CIDR allocated for pod IPs in this tenant cluster.
		PodIpSubnet *string `json:"PodIpSubnet,omitempty"`
		// End of VLAN range allocated to this tenant cluster.
		VlanEnd *string `json:"VlanEnd,omitempty"`
		// Start of VLAN range allocated to this tenant cluster.
		VlanStart    *string                               `json:"VlanStart,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varKubernetesAciCniTenantClusterAllocationWithoutEmbeddedStruct := KubernetesAciCniTenantClusterAllocationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesAciCniTenantClusterAllocationWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesAciCniTenantClusterAllocation := _KubernetesAciCniTenantClusterAllocation{}
		varKubernetesAciCniTenantClusterAllocation.ClassId = varKubernetesAciCniTenantClusterAllocationWithoutEmbeddedStruct.ClassId
		varKubernetesAciCniTenantClusterAllocation.ObjectType = varKubernetesAciCniTenantClusterAllocationWithoutEmbeddedStruct.ObjectType
		varKubernetesAciCniTenantClusterAllocation.NodeSvcIpSubnet = varKubernetesAciCniTenantClusterAllocationWithoutEmbeddedStruct.NodeSvcIpSubnet
		varKubernetesAciCniTenantClusterAllocation.PodIpSubnet = varKubernetesAciCniTenantClusterAllocationWithoutEmbeddedStruct.PodIpSubnet
		varKubernetesAciCniTenantClusterAllocation.VlanEnd = varKubernetesAciCniTenantClusterAllocationWithoutEmbeddedStruct.VlanEnd
		varKubernetesAciCniTenantClusterAllocation.VlanStart = varKubernetesAciCniTenantClusterAllocationWithoutEmbeddedStruct.VlanStart
		varKubernetesAciCniTenantClusterAllocation.Organization = varKubernetesAciCniTenantClusterAllocationWithoutEmbeddedStruct.Organization
		*o = KubernetesAciCniTenantClusterAllocation(varKubernetesAciCniTenantClusterAllocation)
	} else {
		return err
	}

	varKubernetesAciCniTenantClusterAllocation := _KubernetesAciCniTenantClusterAllocation{}

	err = json.Unmarshal(bytes, &varKubernetesAciCniTenantClusterAllocation)
	if err == nil {
		o.MoBaseMo = varKubernetesAciCniTenantClusterAllocation.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NodeSvcIpSubnet")
		delete(additionalProperties, "PodIpSubnet")
		delete(additionalProperties, "VlanEnd")
		delete(additionalProperties, "VlanStart")
		delete(additionalProperties, "Organization")

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

type NullableKubernetesAciCniTenantClusterAllocation struct {
	value *KubernetesAciCniTenantClusterAllocation
	isSet bool
}

func (v NullableKubernetesAciCniTenantClusterAllocation) Get() *KubernetesAciCniTenantClusterAllocation {
	return v.value
}

func (v *NullableKubernetesAciCniTenantClusterAllocation) Set(val *KubernetesAciCniTenantClusterAllocation) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAciCniTenantClusterAllocation) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAciCniTenantClusterAllocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAciCniTenantClusterAllocation(val *KubernetesAciCniTenantClusterAllocation) *NullableKubernetesAciCniTenantClusterAllocation {
	return &NullableKubernetesAciCniTenantClusterAllocation{value: val, isSet: true}
}

func (v NullableKubernetesAciCniTenantClusterAllocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAciCniTenantClusterAllocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
