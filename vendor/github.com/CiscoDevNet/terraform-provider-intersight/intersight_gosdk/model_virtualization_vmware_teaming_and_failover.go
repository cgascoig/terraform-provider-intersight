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

// VirtualizationVmwareTeamingAndFailover The teams can then either share the load of traffic between physical and virtual networks among some or all of its members, or provide passive failover in the event of a hardware failure or a network outage.
type VirtualizationVmwareTeamingAndFailover struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType     string   `json:"ObjectType"`
	ActiveAdapters []string `json:"ActiveAdapters,omitempty"`
	// By default, a failback policy is enabled on a NIC team. If a failed physical NIC returns online, the network component sets the NIC back to active by replacing the standby NIC that took over its slot.
	Failback *bool `json:"Failback,omitempty"`
	// Determines how network traffic is distributed between the network adapters in a NIC team. * `loadbalanceIP` - Load balance based on IP hash. * `loadbalanceSrcmac` - Route based on source MAC hash. * `loadbalanceSrcid` - Route based on originating virtual port. * `failoverExplicit` - Use explicit failover order. * `loadbalanceSrcnic` - Route based on physical NIC load.
	LoadBalancing *string `json:"LoadBalancing,omitempty"`
	// Name of the network component, example dvswitch, dvnetwork, vswitch or standard network.
	Name *string `json:"Name,omitempty"`
	// Methods used by network component for failover detection. * `linkStatus` - This option detects failures such as removed cables and physical switch power failures. * `beaconProbing` - Sends out and listens for beacon probes on all NICs in the team, and uses this information, in addition to link status, to determine link failure. ESXi sends beacon packets every second.
	NetworkFailureDetection *string `json:"NetworkFailureDetection,omitempty"`
	// Determines how network traffic is distributed between the network adapters in a NIC team.
	NotifySwitches       *bool    `json:"NotifySwitches,omitempty"`
	StandbyAdapters      []string `json:"StandbyAdapters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareTeamingAndFailover VirtualizationVmwareTeamingAndFailover

// NewVirtualizationVmwareTeamingAndFailover instantiates a new VirtualizationVmwareTeamingAndFailover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareTeamingAndFailover(classId string, objectType string) *VirtualizationVmwareTeamingAndFailover {
	this := VirtualizationVmwareTeamingAndFailover{}
	this.ClassId = classId
	this.ObjectType = objectType
	var loadBalancing string = "loadbalanceIP"
	this.LoadBalancing = &loadBalancing
	var networkFailureDetection string = "linkStatus"
	this.NetworkFailureDetection = &networkFailureDetection
	return &this
}

// NewVirtualizationVmwareTeamingAndFailoverWithDefaults instantiates a new VirtualizationVmwareTeamingAndFailover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareTeamingAndFailoverWithDefaults() *VirtualizationVmwareTeamingAndFailover {
	this := VirtualizationVmwareTeamingAndFailover{}
	var classId string = "virtualization.VmwareTeamingAndFailover"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareTeamingAndFailover"
	this.ObjectType = objectType
	var loadBalancing string = "loadbalanceIP"
	this.LoadBalancing = &loadBalancing
	var networkFailureDetection string = "linkStatus"
	this.NetworkFailureDetection = &networkFailureDetection
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareTeamingAndFailover) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareTeamingAndFailover) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareTeamingAndFailover) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareTeamingAndFailover) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareTeamingAndFailover) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareTeamingAndFailover) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveAdapters returns the ActiveAdapters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareTeamingAndFailover) GetActiveAdapters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ActiveAdapters
}

// GetActiveAdaptersOk returns a tuple with the ActiveAdapters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareTeamingAndFailover) GetActiveAdaptersOk() (*[]string, bool) {
	if o == nil || o.ActiveAdapters == nil {
		return nil, false
	}
	return &o.ActiveAdapters, true
}

// HasActiveAdapters returns a boolean if a field has been set.
func (o *VirtualizationVmwareTeamingAndFailover) HasActiveAdapters() bool {
	if o != nil && o.ActiveAdapters != nil {
		return true
	}

	return false
}

// SetActiveAdapters gets a reference to the given []string and assigns it to the ActiveAdapters field.
func (o *VirtualizationVmwareTeamingAndFailover) SetActiveAdapters(v []string) {
	o.ActiveAdapters = v
}

// GetFailback returns the Failback field value if set, zero value otherwise.
func (o *VirtualizationVmwareTeamingAndFailover) GetFailback() bool {
	if o == nil || o.Failback == nil {
		var ret bool
		return ret
	}
	return *o.Failback
}

// GetFailbackOk returns a tuple with the Failback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareTeamingAndFailover) GetFailbackOk() (*bool, bool) {
	if o == nil || o.Failback == nil {
		return nil, false
	}
	return o.Failback, true
}

// HasFailback returns a boolean if a field has been set.
func (o *VirtualizationVmwareTeamingAndFailover) HasFailback() bool {
	if o != nil && o.Failback != nil {
		return true
	}

	return false
}

// SetFailback gets a reference to the given bool and assigns it to the Failback field.
func (o *VirtualizationVmwareTeamingAndFailover) SetFailback(v bool) {
	o.Failback = &v
}

// GetLoadBalancing returns the LoadBalancing field value if set, zero value otherwise.
func (o *VirtualizationVmwareTeamingAndFailover) GetLoadBalancing() string {
	if o == nil || o.LoadBalancing == nil {
		var ret string
		return ret
	}
	return *o.LoadBalancing
}

// GetLoadBalancingOk returns a tuple with the LoadBalancing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareTeamingAndFailover) GetLoadBalancingOk() (*string, bool) {
	if o == nil || o.LoadBalancing == nil {
		return nil, false
	}
	return o.LoadBalancing, true
}

// HasLoadBalancing returns a boolean if a field has been set.
func (o *VirtualizationVmwareTeamingAndFailover) HasLoadBalancing() bool {
	if o != nil && o.LoadBalancing != nil {
		return true
	}

	return false
}

// SetLoadBalancing gets a reference to the given string and assigns it to the LoadBalancing field.
func (o *VirtualizationVmwareTeamingAndFailover) SetLoadBalancing(v string) {
	o.LoadBalancing = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationVmwareTeamingAndFailover) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareTeamingAndFailover) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationVmwareTeamingAndFailover) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationVmwareTeamingAndFailover) SetName(v string) {
	o.Name = &v
}

// GetNetworkFailureDetection returns the NetworkFailureDetection field value if set, zero value otherwise.
func (o *VirtualizationVmwareTeamingAndFailover) GetNetworkFailureDetection() string {
	if o == nil || o.NetworkFailureDetection == nil {
		var ret string
		return ret
	}
	return *o.NetworkFailureDetection
}

// GetNetworkFailureDetectionOk returns a tuple with the NetworkFailureDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareTeamingAndFailover) GetNetworkFailureDetectionOk() (*string, bool) {
	if o == nil || o.NetworkFailureDetection == nil {
		return nil, false
	}
	return o.NetworkFailureDetection, true
}

// HasNetworkFailureDetection returns a boolean if a field has been set.
func (o *VirtualizationVmwareTeamingAndFailover) HasNetworkFailureDetection() bool {
	if o != nil && o.NetworkFailureDetection != nil {
		return true
	}

	return false
}

// SetNetworkFailureDetection gets a reference to the given string and assigns it to the NetworkFailureDetection field.
func (o *VirtualizationVmwareTeamingAndFailover) SetNetworkFailureDetection(v string) {
	o.NetworkFailureDetection = &v
}

// GetNotifySwitches returns the NotifySwitches field value if set, zero value otherwise.
func (o *VirtualizationVmwareTeamingAndFailover) GetNotifySwitches() bool {
	if o == nil || o.NotifySwitches == nil {
		var ret bool
		return ret
	}
	return *o.NotifySwitches
}

// GetNotifySwitchesOk returns a tuple with the NotifySwitches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareTeamingAndFailover) GetNotifySwitchesOk() (*bool, bool) {
	if o == nil || o.NotifySwitches == nil {
		return nil, false
	}
	return o.NotifySwitches, true
}

// HasNotifySwitches returns a boolean if a field has been set.
func (o *VirtualizationVmwareTeamingAndFailover) HasNotifySwitches() bool {
	if o != nil && o.NotifySwitches != nil {
		return true
	}

	return false
}

// SetNotifySwitches gets a reference to the given bool and assigns it to the NotifySwitches field.
func (o *VirtualizationVmwareTeamingAndFailover) SetNotifySwitches(v bool) {
	o.NotifySwitches = &v
}

// GetStandbyAdapters returns the StandbyAdapters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareTeamingAndFailover) GetStandbyAdapters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.StandbyAdapters
}

// GetStandbyAdaptersOk returns a tuple with the StandbyAdapters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareTeamingAndFailover) GetStandbyAdaptersOk() (*[]string, bool) {
	if o == nil || o.StandbyAdapters == nil {
		return nil, false
	}
	return &o.StandbyAdapters, true
}

// HasStandbyAdapters returns a boolean if a field has been set.
func (o *VirtualizationVmwareTeamingAndFailover) HasStandbyAdapters() bool {
	if o != nil && o.StandbyAdapters != nil {
		return true
	}

	return false
}

// SetStandbyAdapters gets a reference to the given []string and assigns it to the StandbyAdapters field.
func (o *VirtualizationVmwareTeamingAndFailover) SetStandbyAdapters(v []string) {
	o.StandbyAdapters = v
}

func (o VirtualizationVmwareTeamingAndFailover) MarshalJSON() ([]byte, error) {
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
	if o.ActiveAdapters != nil {
		toSerialize["ActiveAdapters"] = o.ActiveAdapters
	}
	if o.Failback != nil {
		toSerialize["Failback"] = o.Failback
	}
	if o.LoadBalancing != nil {
		toSerialize["LoadBalancing"] = o.LoadBalancing
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NetworkFailureDetection != nil {
		toSerialize["NetworkFailureDetection"] = o.NetworkFailureDetection
	}
	if o.NotifySwitches != nil {
		toSerialize["NotifySwitches"] = o.NotifySwitches
	}
	if o.StandbyAdapters != nil {
		toSerialize["StandbyAdapters"] = o.StandbyAdapters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareTeamingAndFailover) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType     string   `json:"ObjectType"`
		ActiveAdapters []string `json:"ActiveAdapters,omitempty"`
		// By default, a failback policy is enabled on a NIC team. If a failed physical NIC returns online, the network component sets the NIC back to active by replacing the standby NIC that took over its slot.
		Failback *bool `json:"Failback,omitempty"`
		// Determines how network traffic is distributed between the network adapters in a NIC team. * `loadbalanceIP` - Load balance based on IP hash. * `loadbalanceSrcmac` - Route based on source MAC hash. * `loadbalanceSrcid` - Route based on originating virtual port. * `failoverExplicit` - Use explicit failover order. * `loadbalanceSrcnic` - Route based on physical NIC load.
		LoadBalancing *string `json:"LoadBalancing,omitempty"`
		// Name of the network component, example dvswitch, dvnetwork, vswitch or standard network.
		Name *string `json:"Name,omitempty"`
		// Methods used by network component for failover detection. * `linkStatus` - This option detects failures such as removed cables and physical switch power failures. * `beaconProbing` - Sends out and listens for beacon probes on all NICs in the team, and uses this information, in addition to link status, to determine link failure. ESXi sends beacon packets every second.
		NetworkFailureDetection *string `json:"NetworkFailureDetection,omitempty"`
		// Determines how network traffic is distributed between the network adapters in a NIC team.
		NotifySwitches  *bool    `json:"NotifySwitches,omitempty"`
		StandbyAdapters []string `json:"StandbyAdapters,omitempty"`
	}

	varVirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct := VirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareTeamingAndFailover := _VirtualizationVmwareTeamingAndFailover{}
		varVirtualizationVmwareTeamingAndFailover.ClassId = varVirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareTeamingAndFailover.ObjectType = varVirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareTeamingAndFailover.ActiveAdapters = varVirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct.ActiveAdapters
		varVirtualizationVmwareTeamingAndFailover.Failback = varVirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct.Failback
		varVirtualizationVmwareTeamingAndFailover.LoadBalancing = varVirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct.LoadBalancing
		varVirtualizationVmwareTeamingAndFailover.Name = varVirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct.Name
		varVirtualizationVmwareTeamingAndFailover.NetworkFailureDetection = varVirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct.NetworkFailureDetection
		varVirtualizationVmwareTeamingAndFailover.NotifySwitches = varVirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct.NotifySwitches
		varVirtualizationVmwareTeamingAndFailover.StandbyAdapters = varVirtualizationVmwareTeamingAndFailoverWithoutEmbeddedStruct.StandbyAdapters
		*o = VirtualizationVmwareTeamingAndFailover(varVirtualizationVmwareTeamingAndFailover)
	} else {
		return err
	}

	varVirtualizationVmwareTeamingAndFailover := _VirtualizationVmwareTeamingAndFailover{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareTeamingAndFailover)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationVmwareTeamingAndFailover.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveAdapters")
		delete(additionalProperties, "Failback")
		delete(additionalProperties, "LoadBalancing")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NetworkFailureDetection")
		delete(additionalProperties, "NotifySwitches")
		delete(additionalProperties, "StandbyAdapters")

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

type NullableVirtualizationVmwareTeamingAndFailover struct {
	value *VirtualizationVmwareTeamingAndFailover
	isSet bool
}

func (v NullableVirtualizationVmwareTeamingAndFailover) Get() *VirtualizationVmwareTeamingAndFailover {
	return v.value
}

func (v *NullableVirtualizationVmwareTeamingAndFailover) Set(val *VirtualizationVmwareTeamingAndFailover) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareTeamingAndFailover) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareTeamingAndFailover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareTeamingAndFailover(val *VirtualizationVmwareTeamingAndFailover) *NullableVirtualizationVmwareTeamingAndFailover {
	return &NullableVirtualizationVmwareTeamingAndFailover{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareTeamingAndFailover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareTeamingAndFailover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
