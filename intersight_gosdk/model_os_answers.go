/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-05T15:05:56Z.
 *
 * API version: 1.0.9-3562
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// OsAnswers This MO captures the values for the most common set of fields in OS specific answer files. The values provided in this MO are used to construct the OS specific answer files (kickstart, seed, unattended xml) by replacing the fields/placeholders in selected os.ConfigurationFile content with the values of this MO properties.
type OsAnswers struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If the source of the answers is a static file, the content of the file is stored as value in this property. The value is mandatory only when the 'Source' property has been set to 'File'.
	AnswerFile *string `json:"AnswerFile,omitempty"`
	// Hostname to be configured for the server in the OS.
	Hostname *string `json:"Hostname,omitempty"`
	// IP configuration type. Values are Static or Dynamic configuration of IP. In case of static IP configuration, IP address, gateway and other details need to be populated. In case of dynamic the IP configuration is obtained dynamically from DHCP. * `static` - In case of static IP configuraton, provide the details such as IP address, netmask, and gateway. * `DHCP` - In case of dynamic IP configuration, the IP address, netmask and gateway detailsare obtained from DHCP.
	IpConfigType    *string                   `json:"IpConfigType,omitempty"`
	IpConfiguration NullableOsIpConfiguration `json:"IpConfiguration,omitempty"`
	// Indicates whether the value of the 'answerFile' property has been set.
	IsAnswerFileSet *bool `json:"IsAnswerFileSet,omitempty"`
	// Enable to indicate Root Password provided is encrypted.
	IsRootPasswordCrypted *bool `json:"IsRootPasswordCrypted,omitempty"`
	// Indicates whether the value of the 'rootPassword' property has been set.
	IsRootPasswordSet *bool `json:"IsRootPasswordSet,omitempty"`
	// IP address of the name server to be configured in the OS.
	Nameserver *string `json:"Nameserver,omitempty"`
	// The product key to be used for a specific version of Windows installation.
	ProductKey *string `json:"ProductKey,omitempty"`
	// Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password. Intersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password. For more details on encrypting passwords, see Help Center.
	RootPassword *string `json:"RootPassword,omitempty"`
	// Answer values can be provided from three sources - Embedded in OS image, static file, or as placeholder values for an answer file template. Source of the answers is given as value, Embedded/File/Template. 'Embedded' option indicates that the answer file is embedded within the OS Image. 'File' option indicates that the answers are provided as a file. 'Template' indicates that the placeholders in the selected os.ConfigurationFile MO are replaced with values provided as os.Answers MO. * `None` - Indicates that answers is not sent and values must be populated from Install Template.   * `Embedded` - Indicates that the answer file is embedded within OS image. * `File` - Indicates that the answer file is a static content that has all thevalues populated. * `Template` - Indicates that the given answers are used to populate the answer filetemplate. The template allows the users to refer some server specificanswers as fields/placeholders and replace these placeholders with theactual values for each Server during OS installation using 'Answers' and'AdditionalParameters' properties in os.Install MO.The answer file templates can be created by users as os.ConfigurationFile objects.
	Source               *string `json:"Source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsAnswers OsAnswers

// NewOsAnswers instantiates a new OsAnswers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsAnswers(classId string, objectType string) *OsAnswers {
	this := OsAnswers{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ipConfigType string = "static"
	this.IpConfigType = &ipConfigType
	var isAnswerFileSet bool = false
	this.IsAnswerFileSet = &isAnswerFileSet
	var isRootPasswordSet bool = false
	this.IsRootPasswordSet = &isRootPasswordSet
	var source string = "None"
	this.Source = &source
	return &this
}

// NewOsAnswersWithDefaults instantiates a new OsAnswers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsAnswersWithDefaults() *OsAnswers {
	this := OsAnswers{}
	var classId string = "os.Answers"
	this.ClassId = classId
	var objectType string = "os.Answers"
	this.ObjectType = objectType
	var ipConfigType string = "static"
	this.IpConfigType = &ipConfigType
	var isAnswerFileSet bool = false
	this.IsAnswerFileSet = &isAnswerFileSet
	var isRootPasswordSet bool = false
	this.IsRootPasswordSet = &isRootPasswordSet
	var source string = "None"
	this.Source = &source
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsAnswers) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsAnswers) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsAnswers) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsAnswers) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAnswerFile returns the AnswerFile field value if set, zero value otherwise.
func (o *OsAnswers) GetAnswerFile() string {
	if o == nil || o.AnswerFile == nil {
		var ret string
		return ret
	}
	return *o.AnswerFile
}

// GetAnswerFileOk returns a tuple with the AnswerFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetAnswerFileOk() (*string, bool) {
	if o == nil || o.AnswerFile == nil {
		return nil, false
	}
	return o.AnswerFile, true
}

// HasAnswerFile returns a boolean if a field has been set.
func (o *OsAnswers) HasAnswerFile() bool {
	if o != nil && o.AnswerFile != nil {
		return true
	}

	return false
}

// SetAnswerFile gets a reference to the given string and assigns it to the AnswerFile field.
func (o *OsAnswers) SetAnswerFile(v string) {
	o.AnswerFile = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *OsAnswers) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *OsAnswers) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *OsAnswers) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpConfigType returns the IpConfigType field value if set, zero value otherwise.
func (o *OsAnswers) GetIpConfigType() string {
	if o == nil || o.IpConfigType == nil {
		var ret string
		return ret
	}
	return *o.IpConfigType
}

// GetIpConfigTypeOk returns a tuple with the IpConfigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetIpConfigTypeOk() (*string, bool) {
	if o == nil || o.IpConfigType == nil {
		return nil, false
	}
	return o.IpConfigType, true
}

// HasIpConfigType returns a boolean if a field has been set.
func (o *OsAnswers) HasIpConfigType() bool {
	if o != nil && o.IpConfigType != nil {
		return true
	}

	return false
}

// SetIpConfigType gets a reference to the given string and assigns it to the IpConfigType field.
func (o *OsAnswers) SetIpConfigType(v string) {
	o.IpConfigType = &v
}

// GetIpConfiguration returns the IpConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsAnswers) GetIpConfiguration() OsIpConfiguration {
	if o == nil || o.IpConfiguration.Get() == nil {
		var ret OsIpConfiguration
		return ret
	}
	return *o.IpConfiguration.Get()
}

// GetIpConfigurationOk returns a tuple with the IpConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsAnswers) GetIpConfigurationOk() (*OsIpConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpConfiguration.Get(), o.IpConfiguration.IsSet()
}

// HasIpConfiguration returns a boolean if a field has been set.
func (o *OsAnswers) HasIpConfiguration() bool {
	if o != nil && o.IpConfiguration.IsSet() {
		return true
	}

	return false
}

// SetIpConfiguration gets a reference to the given NullableOsIpConfiguration and assigns it to the IpConfiguration field.
func (o *OsAnswers) SetIpConfiguration(v OsIpConfiguration) {
	o.IpConfiguration.Set(&v)
}

// SetIpConfigurationNil sets the value for IpConfiguration to be an explicit nil
func (o *OsAnswers) SetIpConfigurationNil() {
	o.IpConfiguration.Set(nil)
}

// UnsetIpConfiguration ensures that no value is present for IpConfiguration, not even an explicit nil
func (o *OsAnswers) UnsetIpConfiguration() {
	o.IpConfiguration.Unset()
}

// GetIsAnswerFileSet returns the IsAnswerFileSet field value if set, zero value otherwise.
func (o *OsAnswers) GetIsAnswerFileSet() bool {
	if o == nil || o.IsAnswerFileSet == nil {
		var ret bool
		return ret
	}
	return *o.IsAnswerFileSet
}

// GetIsAnswerFileSetOk returns a tuple with the IsAnswerFileSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetIsAnswerFileSetOk() (*bool, bool) {
	if o == nil || o.IsAnswerFileSet == nil {
		return nil, false
	}
	return o.IsAnswerFileSet, true
}

// HasIsAnswerFileSet returns a boolean if a field has been set.
func (o *OsAnswers) HasIsAnswerFileSet() bool {
	if o != nil && o.IsAnswerFileSet != nil {
		return true
	}

	return false
}

// SetIsAnswerFileSet gets a reference to the given bool and assigns it to the IsAnswerFileSet field.
func (o *OsAnswers) SetIsAnswerFileSet(v bool) {
	o.IsAnswerFileSet = &v
}

// GetIsRootPasswordCrypted returns the IsRootPasswordCrypted field value if set, zero value otherwise.
func (o *OsAnswers) GetIsRootPasswordCrypted() bool {
	if o == nil || o.IsRootPasswordCrypted == nil {
		var ret bool
		return ret
	}
	return *o.IsRootPasswordCrypted
}

// GetIsRootPasswordCryptedOk returns a tuple with the IsRootPasswordCrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetIsRootPasswordCryptedOk() (*bool, bool) {
	if o == nil || o.IsRootPasswordCrypted == nil {
		return nil, false
	}
	return o.IsRootPasswordCrypted, true
}

// HasIsRootPasswordCrypted returns a boolean if a field has been set.
func (o *OsAnswers) HasIsRootPasswordCrypted() bool {
	if o != nil && o.IsRootPasswordCrypted != nil {
		return true
	}

	return false
}

// SetIsRootPasswordCrypted gets a reference to the given bool and assigns it to the IsRootPasswordCrypted field.
func (o *OsAnswers) SetIsRootPasswordCrypted(v bool) {
	o.IsRootPasswordCrypted = &v
}

// GetIsRootPasswordSet returns the IsRootPasswordSet field value if set, zero value otherwise.
func (o *OsAnswers) GetIsRootPasswordSet() bool {
	if o == nil || o.IsRootPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsRootPasswordSet
}

// GetIsRootPasswordSetOk returns a tuple with the IsRootPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetIsRootPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsRootPasswordSet == nil {
		return nil, false
	}
	return o.IsRootPasswordSet, true
}

// HasIsRootPasswordSet returns a boolean if a field has been set.
func (o *OsAnswers) HasIsRootPasswordSet() bool {
	if o != nil && o.IsRootPasswordSet != nil {
		return true
	}

	return false
}

// SetIsRootPasswordSet gets a reference to the given bool and assigns it to the IsRootPasswordSet field.
func (o *OsAnswers) SetIsRootPasswordSet(v bool) {
	o.IsRootPasswordSet = &v
}

// GetNameserver returns the Nameserver field value if set, zero value otherwise.
func (o *OsAnswers) GetNameserver() string {
	if o == nil || o.Nameserver == nil {
		var ret string
		return ret
	}
	return *o.Nameserver
}

// GetNameserverOk returns a tuple with the Nameserver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetNameserverOk() (*string, bool) {
	if o == nil || o.Nameserver == nil {
		return nil, false
	}
	return o.Nameserver, true
}

// HasNameserver returns a boolean if a field has been set.
func (o *OsAnswers) HasNameserver() bool {
	if o != nil && o.Nameserver != nil {
		return true
	}

	return false
}

// SetNameserver gets a reference to the given string and assigns it to the Nameserver field.
func (o *OsAnswers) SetNameserver(v string) {
	o.Nameserver = &v
}

// GetProductKey returns the ProductKey field value if set, zero value otherwise.
func (o *OsAnswers) GetProductKey() string {
	if o == nil || o.ProductKey == nil {
		var ret string
		return ret
	}
	return *o.ProductKey
}

// GetProductKeyOk returns a tuple with the ProductKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetProductKeyOk() (*string, bool) {
	if o == nil || o.ProductKey == nil {
		return nil, false
	}
	return o.ProductKey, true
}

// HasProductKey returns a boolean if a field has been set.
func (o *OsAnswers) HasProductKey() bool {
	if o != nil && o.ProductKey != nil {
		return true
	}

	return false
}

// SetProductKey gets a reference to the given string and assigns it to the ProductKey field.
func (o *OsAnswers) SetProductKey(v string) {
	o.ProductKey = &v
}

// GetRootPassword returns the RootPassword field value if set, zero value otherwise.
func (o *OsAnswers) GetRootPassword() string {
	if o == nil || o.RootPassword == nil {
		var ret string
		return ret
	}
	return *o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetRootPasswordOk() (*string, bool) {
	if o == nil || o.RootPassword == nil {
		return nil, false
	}
	return o.RootPassword, true
}

// HasRootPassword returns a boolean if a field has been set.
func (o *OsAnswers) HasRootPassword() bool {
	if o != nil && o.RootPassword != nil {
		return true
	}

	return false
}

// SetRootPassword gets a reference to the given string and assigns it to the RootPassword field.
func (o *OsAnswers) SetRootPassword(v string) {
	o.RootPassword = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *OsAnswers) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *OsAnswers) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *OsAnswers) SetSource(v string) {
	o.Source = &v
}

func (o OsAnswers) MarshalJSON() ([]byte, error) {
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
	if o.AnswerFile != nil {
		toSerialize["AnswerFile"] = o.AnswerFile
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.IpConfigType != nil {
		toSerialize["IpConfigType"] = o.IpConfigType
	}
	if o.IpConfiguration.IsSet() {
		toSerialize["IpConfiguration"] = o.IpConfiguration.Get()
	}
	if o.IsAnswerFileSet != nil {
		toSerialize["IsAnswerFileSet"] = o.IsAnswerFileSet
	}
	if o.IsRootPasswordCrypted != nil {
		toSerialize["IsRootPasswordCrypted"] = o.IsRootPasswordCrypted
	}
	if o.IsRootPasswordSet != nil {
		toSerialize["IsRootPasswordSet"] = o.IsRootPasswordSet
	}
	if o.Nameserver != nil {
		toSerialize["Nameserver"] = o.Nameserver
	}
	if o.ProductKey != nil {
		toSerialize["ProductKey"] = o.ProductKey
	}
	if o.RootPassword != nil {
		toSerialize["RootPassword"] = o.RootPassword
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsAnswers) UnmarshalJSON(bytes []byte) (err error) {
	type OsAnswersWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If the source of the answers is a static file, the content of the file is stored as value in this property. The value is mandatory only when the 'Source' property has been set to 'File'.
		AnswerFile *string `json:"AnswerFile,omitempty"`
		// Hostname to be configured for the server in the OS.
		Hostname *string `json:"Hostname,omitempty"`
		// IP configuration type. Values are Static or Dynamic configuration of IP. In case of static IP configuration, IP address, gateway and other details need to be populated. In case of dynamic the IP configuration is obtained dynamically from DHCP. * `static` - In case of static IP configuraton, provide the details such as IP address, netmask, and gateway. * `DHCP` - In case of dynamic IP configuration, the IP address, netmask and gateway detailsare obtained from DHCP.
		IpConfigType    *string                   `json:"IpConfigType,omitempty"`
		IpConfiguration NullableOsIpConfiguration `json:"IpConfiguration,omitempty"`
		// Indicates whether the value of the 'answerFile' property has been set.
		IsAnswerFileSet *bool `json:"IsAnswerFileSet,omitempty"`
		// Enable to indicate Root Password provided is encrypted.
		IsRootPasswordCrypted *bool `json:"IsRootPasswordCrypted,omitempty"`
		// Indicates whether the value of the 'rootPassword' property has been set.
		IsRootPasswordSet *bool `json:"IsRootPasswordSet,omitempty"`
		// IP address of the name server to be configured in the OS.
		Nameserver *string `json:"Nameserver,omitempty"`
		// The product key to be used for a specific version of Windows installation.
		ProductKey *string `json:"ProductKey,omitempty"`
		// Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password. Intersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password. For more details on encrypting passwords, see Help Center.
		RootPassword *string `json:"RootPassword,omitempty"`
		// Answer values can be provided from three sources - Embedded in OS image, static file, or as placeholder values for an answer file template. Source of the answers is given as value, Embedded/File/Template. 'Embedded' option indicates that the answer file is embedded within the OS Image. 'File' option indicates that the answers are provided as a file. 'Template' indicates that the placeholders in the selected os.ConfigurationFile MO are replaced with values provided as os.Answers MO. * `None` - Indicates that answers is not sent and values must be populated from Install Template.   * `Embedded` - Indicates that the answer file is embedded within OS image. * `File` - Indicates that the answer file is a static content that has all thevalues populated. * `Template` - Indicates that the given answers are used to populate the answer filetemplate. The template allows the users to refer some server specificanswers as fields/placeholders and replace these placeholders with theactual values for each Server during OS installation using 'Answers' and'AdditionalParameters' properties in os.Install MO.The answer file templates can be created by users as os.ConfigurationFile objects.
		Source *string `json:"Source,omitempty"`
	}

	varOsAnswersWithoutEmbeddedStruct := OsAnswersWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOsAnswersWithoutEmbeddedStruct)
	if err == nil {
		varOsAnswers := _OsAnswers{}
		varOsAnswers.ClassId = varOsAnswersWithoutEmbeddedStruct.ClassId
		varOsAnswers.ObjectType = varOsAnswersWithoutEmbeddedStruct.ObjectType
		varOsAnswers.AnswerFile = varOsAnswersWithoutEmbeddedStruct.AnswerFile
		varOsAnswers.Hostname = varOsAnswersWithoutEmbeddedStruct.Hostname
		varOsAnswers.IpConfigType = varOsAnswersWithoutEmbeddedStruct.IpConfigType
		varOsAnswers.IpConfiguration = varOsAnswersWithoutEmbeddedStruct.IpConfiguration
		varOsAnswers.IsAnswerFileSet = varOsAnswersWithoutEmbeddedStruct.IsAnswerFileSet
		varOsAnswers.IsRootPasswordCrypted = varOsAnswersWithoutEmbeddedStruct.IsRootPasswordCrypted
		varOsAnswers.IsRootPasswordSet = varOsAnswersWithoutEmbeddedStruct.IsRootPasswordSet
		varOsAnswers.Nameserver = varOsAnswersWithoutEmbeddedStruct.Nameserver
		varOsAnswers.ProductKey = varOsAnswersWithoutEmbeddedStruct.ProductKey
		varOsAnswers.RootPassword = varOsAnswersWithoutEmbeddedStruct.RootPassword
		varOsAnswers.Source = varOsAnswersWithoutEmbeddedStruct.Source
		*o = OsAnswers(varOsAnswers)
	} else {
		return err
	}

	varOsAnswers := _OsAnswers{}

	err = json.Unmarshal(bytes, &varOsAnswers)
	if err == nil {
		o.MoBaseComplexType = varOsAnswers.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AnswerFile")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "IpConfigType")
		delete(additionalProperties, "IpConfiguration")
		delete(additionalProperties, "IsAnswerFileSet")
		delete(additionalProperties, "IsRootPasswordCrypted")
		delete(additionalProperties, "IsRootPasswordSet")
		delete(additionalProperties, "Nameserver")
		delete(additionalProperties, "ProductKey")
		delete(additionalProperties, "RootPassword")
		delete(additionalProperties, "Source")

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

type NullableOsAnswers struct {
	value *OsAnswers
	isSet bool
}

func (v NullableOsAnswers) Get() *OsAnswers {
	return v.value
}

func (v *NullableOsAnswers) Set(val *OsAnswers) {
	v.value = val
	v.isSet = true
}

func (v NullableOsAnswers) IsSet() bool {
	return v.isSet
}

func (v *NullableOsAnswers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsAnswers(val *OsAnswers) *NullableOsAnswers {
	return &NullableOsAnswers{value: val, isSet: true}
}

func (v NullableOsAnswers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsAnswers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
