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
	"time"
)

// SoftwareDownloadHistoryAllOf Definition of the list of properties defined in 'software.DownloadHistory', excluding properties defined in parent classes.
type SoftwareDownloadHistoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of software which was downloaded.
	Name *string `json:"Name,omitempty"`
	// The product type of the downloaded software.
	Product *string `json:"Product,omitempty"`
	// The download time of the software image.
	Timestamp *time.Time `json:"Timestamp,omitempty"`
	// The version of software which was downloaded.
	Version              *string                                `json:"Version,omitempty"`
	Account              *IamAccountRelationship                `json:"Account,omitempty"`
	Image                *FirmwareBaseDistributableRelationship `json:"Image,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareDownloadHistoryAllOf SoftwareDownloadHistoryAllOf

// NewSoftwareDownloadHistoryAllOf instantiates a new SoftwareDownloadHistoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareDownloadHistoryAllOf(classId string, objectType string) *SoftwareDownloadHistoryAllOf {
	this := SoftwareDownloadHistoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwareDownloadHistoryAllOfWithDefaults instantiates a new SoftwareDownloadHistoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareDownloadHistoryAllOfWithDefaults() *SoftwareDownloadHistoryAllOf {
	this := SoftwareDownloadHistoryAllOf{}
	var classId string = "software.DownloadHistory"
	this.ClassId = classId
	var objectType string = "software.DownloadHistory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareDownloadHistoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareDownloadHistoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareDownloadHistoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareDownloadHistoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SoftwareDownloadHistoryAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistoryAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SoftwareDownloadHistoryAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SoftwareDownloadHistoryAllOf) SetName(v string) {
	o.Name = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *SoftwareDownloadHistoryAllOf) GetProduct() string {
	if o == nil || o.Product == nil {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistoryAllOf) GetProductOk() (*string, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *SoftwareDownloadHistoryAllOf) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *SoftwareDownloadHistoryAllOf) SetProduct(v string) {
	o.Product = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SoftwareDownloadHistoryAllOf) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistoryAllOf) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SoftwareDownloadHistoryAllOf) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *SoftwareDownloadHistoryAllOf) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SoftwareDownloadHistoryAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistoryAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SoftwareDownloadHistoryAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SoftwareDownloadHistoryAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SoftwareDownloadHistoryAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistoryAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SoftwareDownloadHistoryAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *SoftwareDownloadHistoryAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *SoftwareDownloadHistoryAllOf) GetImage() FirmwareBaseDistributableRelationship {
	if o == nil || o.Image == nil {
		var ret FirmwareBaseDistributableRelationship
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistoryAllOf) GetImageOk() (*FirmwareBaseDistributableRelationship, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *SoftwareDownloadHistoryAllOf) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given FirmwareBaseDistributableRelationship and assigns it to the Image field.
func (o *SoftwareDownloadHistoryAllOf) SetImage(v FirmwareBaseDistributableRelationship) {
	o.Image = &v
}

func (o SoftwareDownloadHistoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Product != nil {
		toSerialize["Product"] = o.Product
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Image != nil {
		toSerialize["Image"] = o.Image
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwareDownloadHistoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwareDownloadHistoryAllOf := _SoftwareDownloadHistoryAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwareDownloadHistoryAllOf); err == nil {
		*o = SoftwareDownloadHistoryAllOf(varSoftwareDownloadHistoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Product")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Image")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwareDownloadHistoryAllOf struct {
	value *SoftwareDownloadHistoryAllOf
	isSet bool
}

func (v NullableSoftwareDownloadHistoryAllOf) Get() *SoftwareDownloadHistoryAllOf {
	return v.value
}

func (v *NullableSoftwareDownloadHistoryAllOf) Set(val *SoftwareDownloadHistoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareDownloadHistoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareDownloadHistoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareDownloadHistoryAllOf(val *SoftwareDownloadHistoryAllOf) *NullableSoftwareDownloadHistoryAllOf {
	return &NullableSoftwareDownloadHistoryAllOf{value: val, isSet: true}
}

func (v NullableSoftwareDownloadHistoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareDownloadHistoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
