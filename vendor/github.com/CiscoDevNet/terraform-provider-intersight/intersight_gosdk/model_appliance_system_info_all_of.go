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
)

// ApplianceSystemInfoAllOf Definition of the list of properties defined in 'appliance.SystemInfo', excluding properties defined in parent classes.
type ApplianceSystemInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Connection state of the Intersight Appliance to the Intersight. * `` - The target details have been persisted but Intersight has not yet attempted to connect to the target. * `Connected` - Intersight is able to establish a connection to the target and initiate management activities. * `NotConnected` - Intersight is unable to establish a connection to the target. * `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established. * `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.
	CloudConnStatus *string `json:"CloudConnStatus,omitempty"`
	// Current running deployment size for the Intersight Appliance cluster. Eg. small, medium, large etc.
	DeploymentSize *string `json:"DeploymentSize,omitempty"`
	// Publicly accessible FQDN or IP address of the Intersight Appliance.
	Hostname *string `json:"Hostname,omitempty"`
	// Indicates that the setup initialization process has been completed.
	InitDone *bool `json:"InitDone,omitempty"`
	// Operational status of the Intersight Appliance cluster. * `Unknown` - Operational status of the Intersight Appliance entity is Unknown. * `Operational` - Operational status of the Intersight Appliance entity is Operational. * `Impaired` - Operational status of the Intersight Appliance entity is Impaired. * `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded.
	OperationalStatus *string `json:"OperationalStatus,omitempty"`
	// SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. It is a unique UUID string, and serialId will not change for the life time of the Intersight Appliance.
	SerialId *string `json:"SerialId,omitempty"`
	// Timezone of the Intersight Appliance. * `Pacific/Niue` -  * `Pacific/Pago_Pago` -  * `Pacific/Honolulu` -  * `Pacific/Rarotonga` -  * `Pacific/Tahiti` -  * `Pacific/Marquesas` -  * `America/Anchorage` -  * `Pacific/Gambier` -  * `America/Los_Angeles` -  * `America/Tijuana` -  * `America/Vancouver` -  * `America/Whitehorse` -  * `Pacific/Pitcairn` -  * `America/Dawson_Creek` -  * `America/Denver` -  * `America/Edmonton` -  * `America/Hermosillo` -  * `America/Mazatlan` -  * `America/Phoenix` -  * `America/Yellowknife` -  * `America/Belize` -  * `America/Chicago` -  * `America/Costa_Rica` -  * `America/El_Salvador` -  * `America/Guatemala` -  * `America/Managua` -  * `America/Mexico_City` -  * `America/Regina` -  * `America/Tegucigalpa` -  * `America/Winnipeg` -  * `Pacific/Galapagos` -  * `America/Bogota` -  * `America/Cancun` -  * `America/Cayman` -  * `America/Guayaquil` -  * `America/Havana` -  * `America/Iqaluit` -  * `America/Jamaica` -  * `America/Lima` -  * `America/Nassau` -  * `America/New_York` -  * `America/Panama` -  * `America/Port-au-Prince` -  * `America/Rio_Branco` -  * `America/Toronto` -  * `Pacific/Easter` -  * `America/Caracas` -  * `America/Asuncion` -  * `America/Barbados` -  * `America/Boa_Vista` -  * `America/Campo_Grande` -  * `America/Cuiaba` -  * `America/Curacao` -  * `America/Grand_Turk` -  * `America/Guyana` -  * `America/Halifax` -  * `America/La_Paz` -  * `America/Manaus` -  * `America/Martinique` -  * `America/Port_of_Spain` -  * `America/Porto_Velho` -  * `America/Puerto_Rico` -  * `America/Santo_Domingo` -  * `America/Thule` -  * `Atlantic/Bermuda` -  * `America/St_Johns` -  * `America/Araguaina` -  * `America/Argentina/Buenos_Aires` -  * `America/Bahia` -  * `America/Belem` -  * `America/Cayenne` -  * `America/Fortaleza` -  * `America/Godthab` -  * `America/Maceio` -  * `America/Miquelon` -  * `America/Montevideo` -  * `America/Paramaribo` -  * `America/Recife` -  * `America/Santiago` -  * `America/Sao_Paulo` -  * `Antarctica/Palmer` -  * `Antarctica/Rothera` -  * `Atlantic/Stanley` -  * `America/Noronha` -  * `Atlantic/South_Georgia` -  * `America/Scoresbysund` -  * `Atlantic/Azores` -  * `Atlantic/Cape_Verde` -  * `Africa/Abidjan` -  * `Africa/Accra` -  * `Africa/Bissau` -  * `Africa/Casablanca` -  * `Africa/El_Aaiun` -  * `Africa/Monrovia` -  * `America/Danmarkshavn` -  * `Atlantic/Canary` -  * `Atlantic/Faroe` -  * `Atlantic/Reykjavik` -  * `Etc/GMT` -  * `Europe/Dublin` -  * `Europe/Lisbon` -  * `Europe/London` -  * `Africa/Algiers` -  * `Africa/Ceuta` -  * `Africa/Lagos` -  * `Africa/Ndjamena` -  * `Africa/Tunis` -  * `Africa/Windhoek` -  * `Europe/Amsterdam` -  * `Europe/Andorra` -  * `Europe/Belgrade` -  * `Europe/Berlin` -  * `Europe/Brussels` -  * `Europe/Budapest` -  * `Europe/Copenhagen` -  * `Europe/Gibraltar` -  * `Europe/Luxembourg` -  * `Europe/Madrid` -  * `Europe/Malta` -  * `Europe/Monaco` -  * `Europe/Oslo` -  * `Europe/Paris` -  * `Europe/Prague` -  * `Europe/Rome` -  * `Europe/Stockholm` -  * `Europe/Tirane` -  * `Europe/Vienna` -  * `Europe/Warsaw` -  * `Europe/Zurich` -  * `Africa/Cairo` -  * `Africa/Johannesburg` -  * `Africa/Maputo` -  * `Africa/Tripoli` -  * `Asia/Amman` -  * `Asia/Beirut` -  * `Asia/Damascus` -  * `Asia/Gaza` -  * `Asia/Jerusalem` -  * `Asia/Nicosia` -  * `Europe/Athens` -  * `Europe/Bucharest` -  * `Europe/Chisinau` -  * `Europe/Helsinki` -  * `Europe/Istanbul` -  * `Europe/Kaliningrad` -  * `Europe/Kiev` -  * `Europe/Riga` -  * `Europe/Sofia` -  * `Europe/Tallinn` -  * `Europe/Vilnius` -  * `Africa/Khartoum` -  * `Africa/Nairobi` -  * `Antarctica/Syowa` -  * `Asia/Baghdad` -  * `Asia/Qatar` -  * `Asia/Riyadh` -  * `Europe/Minsk` -  * `Europe/Moscow` -  * `Asia/Tehran` -  * `Asia/Baku` -  * `Asia/Dubai` -  * `Asia/Tbilisi` -  * `Asia/Yerevan` -  * `Europe/Samara` -  * `Indian/Mahe` -  * `Indian/Mauritius` -  * `Indian/Reunion` -  * `Asia/Kabul` -  * `Antarctica/Mawson` -  * `Asia/Aqtau` -  * `Asia/Aqtobe` -  * `Asia/Ashgabat` -  * `Asia/Dushanbe` -  * `Asia/Karachi` -  * `Asia/Tashkent` -  * `Asia/Yekaterinburg` -  * `Indian/Kerguelen` -  * `Indian/Maldives` -  * `Asia/Calcutta` -  * `Asia/Kolkata` -  * `Asia/Colombo` -  * `Asia/Katmandu` -  * `Antarctica/Vostok` -  * `Asia/Almaty` -  * `Asia/Bishkek` -  * `Asia/Dhaka` -  * `Asia/Omsk` -  * `Asia/Thimphu` -  * `Indian/Chagos` -  * `Asia/Rangoon` -  * `Indian/Cocos` -  * `Antarctica/Davis` -  * `Asia/Bangkok` -  * `Asia/Hovd` -  * `Asia/Jakarta` -  * `Asia/Krasnoyarsk` -  * `Asia/Saigon` -  * `Indian/Christmas` -  * `Antarctica/Casey` -  * `Asia/Brunei` -  * `Asia/Choibalsan` -  * `Asia/Hong_Kong` -  * `Asia/Irkutsk` -  * `Asia/Kuala_Lumpur` -  * `Asia/Macau` -  * `Asia/Makassar` -  * `Asia/Manila` -  * `Asia/Shanghai` -  * `Asia/Singapore` -  * `Asia/Taipei` -  * `Asia/Ulaanbaatar` -  * `Australia/Perth` -  * `Asia/Pyongyang` -  * `Asia/Dili` -  * `Asia/Jayapura` -  * `Asia/Seoul` -  * `Asia/Tokyo` -  * `Asia/Yakutsk` -  * `Pacific/Palau` -  * `Australia/Adelaide` -  * `Australia/Darwin` -  * `Antarctica/DumontDUrville` -  * `Asia/Magadan` -  * `Asia/Vladivostok` -  * `Australia/Brisbane` -  * `Australia/Hobart` -  * `Australia/Sydney` -  * `Pacific/Chuuk` -  * `Pacific/Guam` -  * `Pacific/Port_Moresby` -  * `Pacific/Efate` -  * `Pacific/Guadalcanal` -  * `Pacific/Kosrae` -  * `Pacific/Norfolk` -  * `Pacific/Noumea` -  * `Pacific/Pohnpei` -  * `Asia/Kamchatka` -  * `Pacific/Auckland` -  * `Pacific/Fiji` -  * `Pacific/Funafuti` -  * `Pacific/Kwajalein` -  * `Pacific/Majuro` -  * `Pacific/Nauru` -  * `Pacific/Tarawa` -  * `Pacific/Wake` -  * `Pacific/Wallis` -  * `Pacific/Apia` -  * `Pacific/Enderbury` -  * `Pacific/Fakaofo` -  * `Pacific/Tongatapu` -  * `Pacific/Kiritimati` -
	TimeZone *string `json:"TimeZone,omitempty"`
	// Current software version of the Intersight Appliance.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceSystemInfoAllOf ApplianceSystemInfoAllOf

// NewApplianceSystemInfoAllOf instantiates a new ApplianceSystemInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceSystemInfoAllOf(classId string, objectType string) *ApplianceSystemInfoAllOf {
	this := ApplianceSystemInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cloudConnStatus string = ""
	this.CloudConnStatus = &cloudConnStatus
	var operationalStatus string = "Unknown"
	this.OperationalStatus = &operationalStatus
	var timeZone string = "Pacific/Niue"
	this.TimeZone = &timeZone
	return &this
}

// NewApplianceSystemInfoAllOfWithDefaults instantiates a new ApplianceSystemInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceSystemInfoAllOfWithDefaults() *ApplianceSystemInfoAllOf {
	this := ApplianceSystemInfoAllOf{}
	var classId string = "appliance.SystemInfo"
	this.ClassId = classId
	var objectType string = "appliance.SystemInfo"
	this.ObjectType = objectType
	var cloudConnStatus string = ""
	this.CloudConnStatus = &cloudConnStatus
	var operationalStatus string = "Unknown"
	this.OperationalStatus = &operationalStatus
	var timeZone string = "Pacific/Niue"
	this.TimeZone = &timeZone
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceSystemInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceSystemInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceSystemInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceSystemInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCloudConnStatus returns the CloudConnStatus field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetCloudConnStatus() string {
	if o == nil || o.CloudConnStatus == nil {
		var ret string
		return ret
	}
	return *o.CloudConnStatus
}

// GetCloudConnStatusOk returns a tuple with the CloudConnStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetCloudConnStatusOk() (*string, bool) {
	if o == nil || o.CloudConnStatus == nil {
		return nil, false
	}
	return o.CloudConnStatus, true
}

// HasCloudConnStatus returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasCloudConnStatus() bool {
	if o != nil && o.CloudConnStatus != nil {
		return true
	}

	return false
}

// SetCloudConnStatus gets a reference to the given string and assigns it to the CloudConnStatus field.
func (o *ApplianceSystemInfoAllOf) SetCloudConnStatus(v string) {
	o.CloudConnStatus = &v
}

// GetDeploymentSize returns the DeploymentSize field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetDeploymentSize() string {
	if o == nil || o.DeploymentSize == nil {
		var ret string
		return ret
	}
	return *o.DeploymentSize
}

// GetDeploymentSizeOk returns a tuple with the DeploymentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetDeploymentSizeOk() (*string, bool) {
	if o == nil || o.DeploymentSize == nil {
		return nil, false
	}
	return o.DeploymentSize, true
}

// HasDeploymentSize returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasDeploymentSize() bool {
	if o != nil && o.DeploymentSize != nil {
		return true
	}

	return false
}

// SetDeploymentSize gets a reference to the given string and assigns it to the DeploymentSize field.
func (o *ApplianceSystemInfoAllOf) SetDeploymentSize(v string) {
	o.DeploymentSize = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApplianceSystemInfoAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetInitDone returns the InitDone field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetInitDone() bool {
	if o == nil || o.InitDone == nil {
		var ret bool
		return ret
	}
	return *o.InitDone
}

// GetInitDoneOk returns a tuple with the InitDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetInitDoneOk() (*bool, bool) {
	if o == nil || o.InitDone == nil {
		return nil, false
	}
	return o.InitDone, true
}

// HasInitDone returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasInitDone() bool {
	if o != nil && o.InitDone != nil {
		return true
	}

	return false
}

// SetInitDone gets a reference to the given bool and assigns it to the InitDone field.
func (o *ApplianceSystemInfoAllOf) SetInitDone(v bool) {
	o.InitDone = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceSystemInfoAllOf) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetSerialId returns the SerialId field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetSerialId() string {
	if o == nil || o.SerialId == nil {
		var ret string
		return ret
	}
	return *o.SerialId
}

// GetSerialIdOk returns a tuple with the SerialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetSerialIdOk() (*string, bool) {
	if o == nil || o.SerialId == nil {
		return nil, false
	}
	return o.SerialId, true
}

// HasSerialId returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasSerialId() bool {
	if o != nil && o.SerialId != nil {
		return true
	}

	return false
}

// SetSerialId gets a reference to the given string and assigns it to the SerialId field.
func (o *ApplianceSystemInfoAllOf) SetSerialId(v string) {
	o.SerialId = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *ApplianceSystemInfoAllOf) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApplianceSystemInfoAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o ApplianceSystemInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CloudConnStatus != nil {
		toSerialize["CloudConnStatus"] = o.CloudConnStatus
	}
	if o.DeploymentSize != nil {
		toSerialize["DeploymentSize"] = o.DeploymentSize
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.InitDone != nil {
		toSerialize["InitDone"] = o.InitDone
	}
	if o.OperationalStatus != nil {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	if o.SerialId != nil {
		toSerialize["SerialId"] = o.SerialId
	}
	if o.TimeZone != nil {
		toSerialize["TimeZone"] = o.TimeZone
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceSystemInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceSystemInfoAllOf := _ApplianceSystemInfoAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceSystemInfoAllOf); err == nil {
		*o = ApplianceSystemInfoAllOf(varApplianceSystemInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CloudConnStatus")
		delete(additionalProperties, "DeploymentSize")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "InitDone")
		delete(additionalProperties, "OperationalStatus")
		delete(additionalProperties, "SerialId")
		delete(additionalProperties, "TimeZone")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceSystemInfoAllOf struct {
	value *ApplianceSystemInfoAllOf
	isSet bool
}

func (v NullableApplianceSystemInfoAllOf) Get() *ApplianceSystemInfoAllOf {
	return v.value
}

func (v *NullableApplianceSystemInfoAllOf) Set(val *ApplianceSystemInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceSystemInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceSystemInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceSystemInfoAllOf(val *ApplianceSystemInfoAllOf) *NullableApplianceSystemInfoAllOf {
	return &NullableApplianceSystemInfoAllOf{value: val, isSet: true}
}

func (v NullableApplianceSystemInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceSystemInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
