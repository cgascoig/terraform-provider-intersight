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

// VnicIscsiBootPolicyAllOf Definition of the list of properties defined in 'vnic.IscsiBootPolicy', excluding properties defined in parent classes.
type VnicIscsiBootPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Auto target interface that is represented via the Initiator name or the DHCP vendor ID. The vendor ID can be up to 32 alphanumeric characters.
	AutoTargetvendorName *string                      `json:"AutoTargetvendorName,omitempty"`
	Chap                 NullableVnicIscsiAuthProfile `json:"Chap,omitempty"`
	// Source Type of Initiator IP Address - Auto/Static/Pool. * `DHCP` - The IP address is assigned using DHCP, if available. * `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area. * `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool.
	InitiatorIpSource *string `json:"InitiatorIpSource,omitempty"`
	// Static IP address provided for iSCSI Initiator.
	InitiatorStaticIpV4Address *string                      `json:"InitiatorStaticIpV4Address,omitempty"`
	InitiatorStaticIpV4Config  NullableIppoolIpV4Config     `json:"InitiatorStaticIpV4Config,omitempty"`
	MutualChap                 NullableVnicIscsiAuthProfile `json:"MutualChap,omitempty"`
	// Source Type of Targets - Auto/Static. * `Static` - Type indicates that static target interface is assigned to iSCSI boot. * `Auto` - Type indicates that the system selects the target interface automatically during iSCSI boot.
	TargetSourceType      *string                                  `json:"TargetSourceType,omitempty"`
	InitiatorIpPool       *IppoolPoolRelationship                  `json:"InitiatorIpPool,omitempty"`
	IscsiAdapterPolicy    *VnicIscsiAdapterPolicyRelationship      `json:"IscsiAdapterPolicy,omitempty"`
	Organization          *OrganizationOrganizationRelationship    `json:"Organization,omitempty"`
	PrimaryTargetPolicy   *VnicIscsiStaticTargetPolicyRelationship `json:"PrimaryTargetPolicy,omitempty"`
	SecondaryTargetPolicy *VnicIscsiStaticTargetPolicyRelationship `json:"SecondaryTargetPolicy,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _VnicIscsiBootPolicyAllOf VnicIscsiBootPolicyAllOf

// NewVnicIscsiBootPolicyAllOf instantiates a new VnicIscsiBootPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicIscsiBootPolicyAllOf(classId string, objectType string) *VnicIscsiBootPolicyAllOf {
	this := VnicIscsiBootPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var initiatorIpSource string = "DHCP"
	this.InitiatorIpSource = &initiatorIpSource
	var targetSourceType string = "Static"
	this.TargetSourceType = &targetSourceType
	return &this
}

// NewVnicIscsiBootPolicyAllOfWithDefaults instantiates a new VnicIscsiBootPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicIscsiBootPolicyAllOfWithDefaults() *VnicIscsiBootPolicyAllOf {
	this := VnicIscsiBootPolicyAllOf{}
	var classId string = "vnic.IscsiBootPolicy"
	this.ClassId = classId
	var objectType string = "vnic.IscsiBootPolicy"
	this.ObjectType = objectType
	var initiatorIpSource string = "DHCP"
	this.InitiatorIpSource = &initiatorIpSource
	var targetSourceType string = "Static"
	this.TargetSourceType = &targetSourceType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicIscsiBootPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicIscsiBootPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicIscsiBootPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicIscsiBootPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutoTargetvendorName returns the AutoTargetvendorName field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyAllOf) GetAutoTargetvendorName() string {
	if o == nil || o.AutoTargetvendorName == nil {
		var ret string
		return ret
	}
	return *o.AutoTargetvendorName
}

// GetAutoTargetvendorNameOk returns a tuple with the AutoTargetvendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyAllOf) GetAutoTargetvendorNameOk() (*string, bool) {
	if o == nil || o.AutoTargetvendorName == nil {
		return nil, false
	}
	return o.AutoTargetvendorName, true
}

// HasAutoTargetvendorName returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasAutoTargetvendorName() bool {
	if o != nil && o.AutoTargetvendorName != nil {
		return true
	}

	return false
}

// SetAutoTargetvendorName gets a reference to the given string and assigns it to the AutoTargetvendorName field.
func (o *VnicIscsiBootPolicyAllOf) SetAutoTargetvendorName(v string) {
	o.AutoTargetvendorName = &v
}

// GetChap returns the Chap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiBootPolicyAllOf) GetChap() VnicIscsiAuthProfile {
	if o == nil || o.Chap.Get() == nil {
		var ret VnicIscsiAuthProfile
		return ret
	}
	return *o.Chap.Get()
}

// GetChapOk returns a tuple with the Chap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiBootPolicyAllOf) GetChapOk() (*VnicIscsiAuthProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.Chap.Get(), o.Chap.IsSet()
}

// HasChap returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasChap() bool {
	if o != nil && o.Chap.IsSet() {
		return true
	}

	return false
}

// SetChap gets a reference to the given NullableVnicIscsiAuthProfile and assigns it to the Chap field.
func (o *VnicIscsiBootPolicyAllOf) SetChap(v VnicIscsiAuthProfile) {
	o.Chap.Set(&v)
}

// SetChapNil sets the value for Chap to be an explicit nil
func (o *VnicIscsiBootPolicyAllOf) SetChapNil() {
	o.Chap.Set(nil)
}

// UnsetChap ensures that no value is present for Chap, not even an explicit nil
func (o *VnicIscsiBootPolicyAllOf) UnsetChap() {
	o.Chap.Unset()
}

// GetInitiatorIpSource returns the InitiatorIpSource field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyAllOf) GetInitiatorIpSource() string {
	if o == nil || o.InitiatorIpSource == nil {
		var ret string
		return ret
	}
	return *o.InitiatorIpSource
}

// GetInitiatorIpSourceOk returns a tuple with the InitiatorIpSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyAllOf) GetInitiatorIpSourceOk() (*string, bool) {
	if o == nil || o.InitiatorIpSource == nil {
		return nil, false
	}
	return o.InitiatorIpSource, true
}

// HasInitiatorIpSource returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasInitiatorIpSource() bool {
	if o != nil && o.InitiatorIpSource != nil {
		return true
	}

	return false
}

// SetInitiatorIpSource gets a reference to the given string and assigns it to the InitiatorIpSource field.
func (o *VnicIscsiBootPolicyAllOf) SetInitiatorIpSource(v string) {
	o.InitiatorIpSource = &v
}

// GetInitiatorStaticIpV4Address returns the InitiatorStaticIpV4Address field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyAllOf) GetInitiatorStaticIpV4Address() string {
	if o == nil || o.InitiatorStaticIpV4Address == nil {
		var ret string
		return ret
	}
	return *o.InitiatorStaticIpV4Address
}

// GetInitiatorStaticIpV4AddressOk returns a tuple with the InitiatorStaticIpV4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyAllOf) GetInitiatorStaticIpV4AddressOk() (*string, bool) {
	if o == nil || o.InitiatorStaticIpV4Address == nil {
		return nil, false
	}
	return o.InitiatorStaticIpV4Address, true
}

// HasInitiatorStaticIpV4Address returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasInitiatorStaticIpV4Address() bool {
	if o != nil && o.InitiatorStaticIpV4Address != nil {
		return true
	}

	return false
}

// SetInitiatorStaticIpV4Address gets a reference to the given string and assigns it to the InitiatorStaticIpV4Address field.
func (o *VnicIscsiBootPolicyAllOf) SetInitiatorStaticIpV4Address(v string) {
	o.InitiatorStaticIpV4Address = &v
}

// GetInitiatorStaticIpV4Config returns the InitiatorStaticIpV4Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiBootPolicyAllOf) GetInitiatorStaticIpV4Config() IppoolIpV4Config {
	if o == nil || o.InitiatorStaticIpV4Config.Get() == nil {
		var ret IppoolIpV4Config
		return ret
	}
	return *o.InitiatorStaticIpV4Config.Get()
}

// GetInitiatorStaticIpV4ConfigOk returns a tuple with the InitiatorStaticIpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiBootPolicyAllOf) GetInitiatorStaticIpV4ConfigOk() (*IppoolIpV4Config, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitiatorStaticIpV4Config.Get(), o.InitiatorStaticIpV4Config.IsSet()
}

// HasInitiatorStaticIpV4Config returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasInitiatorStaticIpV4Config() bool {
	if o != nil && o.InitiatorStaticIpV4Config.IsSet() {
		return true
	}

	return false
}

// SetInitiatorStaticIpV4Config gets a reference to the given NullableIppoolIpV4Config and assigns it to the InitiatorStaticIpV4Config field.
func (o *VnicIscsiBootPolicyAllOf) SetInitiatorStaticIpV4Config(v IppoolIpV4Config) {
	o.InitiatorStaticIpV4Config.Set(&v)
}

// SetInitiatorStaticIpV4ConfigNil sets the value for InitiatorStaticIpV4Config to be an explicit nil
func (o *VnicIscsiBootPolicyAllOf) SetInitiatorStaticIpV4ConfigNil() {
	o.InitiatorStaticIpV4Config.Set(nil)
}

// UnsetInitiatorStaticIpV4Config ensures that no value is present for InitiatorStaticIpV4Config, not even an explicit nil
func (o *VnicIscsiBootPolicyAllOf) UnsetInitiatorStaticIpV4Config() {
	o.InitiatorStaticIpV4Config.Unset()
}

// GetMutualChap returns the MutualChap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiBootPolicyAllOf) GetMutualChap() VnicIscsiAuthProfile {
	if o == nil || o.MutualChap.Get() == nil {
		var ret VnicIscsiAuthProfile
		return ret
	}
	return *o.MutualChap.Get()
}

// GetMutualChapOk returns a tuple with the MutualChap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiBootPolicyAllOf) GetMutualChapOk() (*VnicIscsiAuthProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.MutualChap.Get(), o.MutualChap.IsSet()
}

// HasMutualChap returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasMutualChap() bool {
	if o != nil && o.MutualChap.IsSet() {
		return true
	}

	return false
}

// SetMutualChap gets a reference to the given NullableVnicIscsiAuthProfile and assigns it to the MutualChap field.
func (o *VnicIscsiBootPolicyAllOf) SetMutualChap(v VnicIscsiAuthProfile) {
	o.MutualChap.Set(&v)
}

// SetMutualChapNil sets the value for MutualChap to be an explicit nil
func (o *VnicIscsiBootPolicyAllOf) SetMutualChapNil() {
	o.MutualChap.Set(nil)
}

// UnsetMutualChap ensures that no value is present for MutualChap, not even an explicit nil
func (o *VnicIscsiBootPolicyAllOf) UnsetMutualChap() {
	o.MutualChap.Unset()
}

// GetTargetSourceType returns the TargetSourceType field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyAllOf) GetTargetSourceType() string {
	if o == nil || o.TargetSourceType == nil {
		var ret string
		return ret
	}
	return *o.TargetSourceType
}

// GetTargetSourceTypeOk returns a tuple with the TargetSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyAllOf) GetTargetSourceTypeOk() (*string, bool) {
	if o == nil || o.TargetSourceType == nil {
		return nil, false
	}
	return o.TargetSourceType, true
}

// HasTargetSourceType returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasTargetSourceType() bool {
	if o != nil && o.TargetSourceType != nil {
		return true
	}

	return false
}

// SetTargetSourceType gets a reference to the given string and assigns it to the TargetSourceType field.
func (o *VnicIscsiBootPolicyAllOf) SetTargetSourceType(v string) {
	o.TargetSourceType = &v
}

// GetInitiatorIpPool returns the InitiatorIpPool field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyAllOf) GetInitiatorIpPool() IppoolPoolRelationship {
	if o == nil || o.InitiatorIpPool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.InitiatorIpPool
}

// GetInitiatorIpPoolOk returns a tuple with the InitiatorIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyAllOf) GetInitiatorIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.InitiatorIpPool == nil {
		return nil, false
	}
	return o.InitiatorIpPool, true
}

// HasInitiatorIpPool returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasInitiatorIpPool() bool {
	if o != nil && o.InitiatorIpPool != nil {
		return true
	}

	return false
}

// SetInitiatorIpPool gets a reference to the given IppoolPoolRelationship and assigns it to the InitiatorIpPool field.
func (o *VnicIscsiBootPolicyAllOf) SetInitiatorIpPool(v IppoolPoolRelationship) {
	o.InitiatorIpPool = &v
}

// GetIscsiAdapterPolicy returns the IscsiAdapterPolicy field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyAllOf) GetIscsiAdapterPolicy() VnicIscsiAdapterPolicyRelationship {
	if o == nil || o.IscsiAdapterPolicy == nil {
		var ret VnicIscsiAdapterPolicyRelationship
		return ret
	}
	return *o.IscsiAdapterPolicy
}

// GetIscsiAdapterPolicyOk returns a tuple with the IscsiAdapterPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyAllOf) GetIscsiAdapterPolicyOk() (*VnicIscsiAdapterPolicyRelationship, bool) {
	if o == nil || o.IscsiAdapterPolicy == nil {
		return nil, false
	}
	return o.IscsiAdapterPolicy, true
}

// HasIscsiAdapterPolicy returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasIscsiAdapterPolicy() bool {
	if o != nil && o.IscsiAdapterPolicy != nil {
		return true
	}

	return false
}

// SetIscsiAdapterPolicy gets a reference to the given VnicIscsiAdapterPolicyRelationship and assigns it to the IscsiAdapterPolicy field.
func (o *VnicIscsiBootPolicyAllOf) SetIscsiAdapterPolicy(v VnicIscsiAdapterPolicyRelationship) {
	o.IscsiAdapterPolicy = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicIscsiBootPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetPrimaryTargetPolicy returns the PrimaryTargetPolicy field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyAllOf) GetPrimaryTargetPolicy() VnicIscsiStaticTargetPolicyRelationship {
	if o == nil || o.PrimaryTargetPolicy == nil {
		var ret VnicIscsiStaticTargetPolicyRelationship
		return ret
	}
	return *o.PrimaryTargetPolicy
}

// GetPrimaryTargetPolicyOk returns a tuple with the PrimaryTargetPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyAllOf) GetPrimaryTargetPolicyOk() (*VnicIscsiStaticTargetPolicyRelationship, bool) {
	if o == nil || o.PrimaryTargetPolicy == nil {
		return nil, false
	}
	return o.PrimaryTargetPolicy, true
}

// HasPrimaryTargetPolicy returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasPrimaryTargetPolicy() bool {
	if o != nil && o.PrimaryTargetPolicy != nil {
		return true
	}

	return false
}

// SetPrimaryTargetPolicy gets a reference to the given VnicIscsiStaticTargetPolicyRelationship and assigns it to the PrimaryTargetPolicy field.
func (o *VnicIscsiBootPolicyAllOf) SetPrimaryTargetPolicy(v VnicIscsiStaticTargetPolicyRelationship) {
	o.PrimaryTargetPolicy = &v
}

// GetSecondaryTargetPolicy returns the SecondaryTargetPolicy field value if set, zero value otherwise.
func (o *VnicIscsiBootPolicyAllOf) GetSecondaryTargetPolicy() VnicIscsiStaticTargetPolicyRelationship {
	if o == nil || o.SecondaryTargetPolicy == nil {
		var ret VnicIscsiStaticTargetPolicyRelationship
		return ret
	}
	return *o.SecondaryTargetPolicy
}

// GetSecondaryTargetPolicyOk returns a tuple with the SecondaryTargetPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiBootPolicyAllOf) GetSecondaryTargetPolicyOk() (*VnicIscsiStaticTargetPolicyRelationship, bool) {
	if o == nil || o.SecondaryTargetPolicy == nil {
		return nil, false
	}
	return o.SecondaryTargetPolicy, true
}

// HasSecondaryTargetPolicy returns a boolean if a field has been set.
func (o *VnicIscsiBootPolicyAllOf) HasSecondaryTargetPolicy() bool {
	if o != nil && o.SecondaryTargetPolicy != nil {
		return true
	}

	return false
}

// SetSecondaryTargetPolicy gets a reference to the given VnicIscsiStaticTargetPolicyRelationship and assigns it to the SecondaryTargetPolicy field.
func (o *VnicIscsiBootPolicyAllOf) SetSecondaryTargetPolicy(v VnicIscsiStaticTargetPolicyRelationship) {
	o.SecondaryTargetPolicy = &v
}

func (o VnicIscsiBootPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AutoTargetvendorName != nil {
		toSerialize["AutoTargetvendorName"] = o.AutoTargetvendorName
	}
	if o.Chap.IsSet() {
		toSerialize["Chap"] = o.Chap.Get()
	}
	if o.InitiatorIpSource != nil {
		toSerialize["InitiatorIpSource"] = o.InitiatorIpSource
	}
	if o.InitiatorStaticIpV4Address != nil {
		toSerialize["InitiatorStaticIpV4Address"] = o.InitiatorStaticIpV4Address
	}
	if o.InitiatorStaticIpV4Config.IsSet() {
		toSerialize["InitiatorStaticIpV4Config"] = o.InitiatorStaticIpV4Config.Get()
	}
	if o.MutualChap.IsSet() {
		toSerialize["MutualChap"] = o.MutualChap.Get()
	}
	if o.TargetSourceType != nil {
		toSerialize["TargetSourceType"] = o.TargetSourceType
	}
	if o.InitiatorIpPool != nil {
		toSerialize["InitiatorIpPool"] = o.InitiatorIpPool
	}
	if o.IscsiAdapterPolicy != nil {
		toSerialize["IscsiAdapterPolicy"] = o.IscsiAdapterPolicy
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.PrimaryTargetPolicy != nil {
		toSerialize["PrimaryTargetPolicy"] = o.PrimaryTargetPolicy
	}
	if o.SecondaryTargetPolicy != nil {
		toSerialize["SecondaryTargetPolicy"] = o.SecondaryTargetPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicIscsiBootPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicIscsiBootPolicyAllOf := _VnicIscsiBootPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varVnicIscsiBootPolicyAllOf); err == nil {
		*o = VnicIscsiBootPolicyAllOf(varVnicIscsiBootPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoTargetvendorName")
		delete(additionalProperties, "Chap")
		delete(additionalProperties, "InitiatorIpSource")
		delete(additionalProperties, "InitiatorStaticIpV4Address")
		delete(additionalProperties, "InitiatorStaticIpV4Config")
		delete(additionalProperties, "MutualChap")
		delete(additionalProperties, "TargetSourceType")
		delete(additionalProperties, "InitiatorIpPool")
		delete(additionalProperties, "IscsiAdapterPolicy")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "PrimaryTargetPolicy")
		delete(additionalProperties, "SecondaryTargetPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicIscsiBootPolicyAllOf struct {
	value *VnicIscsiBootPolicyAllOf
	isSet bool
}

func (v NullableVnicIscsiBootPolicyAllOf) Get() *VnicIscsiBootPolicyAllOf {
	return v.value
}

func (v *NullableVnicIscsiBootPolicyAllOf) Set(val *VnicIscsiBootPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicIscsiBootPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicIscsiBootPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicIscsiBootPolicyAllOf(val *VnicIscsiBootPolicyAllOf) *NullableVnicIscsiBootPolicyAllOf {
	return &NullableVnicIscsiBootPolicyAllOf{value: val, isSet: true}
}

func (v NullableVnicIscsiBootPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicIscsiBootPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
