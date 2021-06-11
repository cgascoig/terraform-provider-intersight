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
	"time"
)

// UcsdBackupInfo List of backup images available for target end device for restore operation.
type UcsdBackupInfo struct {
	RecoveryAbstractBackupInfo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Auto generated backup File Name with combination of file prefix given an user input and the timestamp.
	BackupFileName *string `json:"BackupFileName,omitempty"`
	// Backup location that contains the backup images for end device which can be used for restore operation.
	BackupLocation *string `json:"BackupLocation,omitempty"`
	// Backup server where backup images are maintained.
	BackupServerIp *string `json:"BackupServerIp,omitempty"`
	// Size of the backup image in bytes.
	BackupSize *int64              `json:"BackupSize,omitempty"`
	Connectors []UcsdConnectorPack `json:"Connectors,omitempty"`
	// Time taken for the backup to be completed.
	Duration *int64 `json:"Duration,omitempty"`
	// The key used for encrypting the backup file.
	EncryptionKey *string `json:"EncryptionKey,omitempty"`
	// Reason for backup failure.
	FailureReason *string `json:"FailureReason,omitempty"`
	// Backup image got purged or not. The backup images get purged based on the retention count set by the user in the backup config policy.
	IsPurged *bool `json:"IsPurged,omitempty"`
	// Last modified time when this backup record got updated.
	LastModified *time.Time `json:"LastModified,omitempty"`
	// Backup current precentage completion status information.
	PercentageCompletion *int64 `json:"PercentageCompletion,omitempty"`
	// The end device product version when the backup image was taken.
	ProductVersion *string `json:"ProductVersion,omitempty"`
	// Protocol used for the remote backup. possible values are FTP, SCP and SFTP. Not applicable for the localhost (127.0.0.1).
	Protocol *string `json:"Protocol,omitempty"`
	// Backup current status stage information.
	StageCompletion *string `json:"StageCompletion,omitempty"`
	// Start time of backup when it got initiated.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// Current status of Backup current.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UcsdBackupInfo UcsdBackupInfo

// NewUcsdBackupInfo instantiates a new UcsdBackupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUcsdBackupInfo(classId string, objectType string) *UcsdBackupInfo {
	this := UcsdBackupInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewUcsdBackupInfoWithDefaults instantiates a new UcsdBackupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUcsdBackupInfoWithDefaults() *UcsdBackupInfo {
	this := UcsdBackupInfo{}
	var classId string = "ucsd.BackupInfo"
	this.ClassId = classId
	var objectType string = "ucsd.BackupInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UcsdBackupInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UcsdBackupInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *UcsdBackupInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UcsdBackupInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackupFileName returns the BackupFileName field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetBackupFileName() string {
	if o == nil || o.BackupFileName == nil {
		var ret string
		return ret
	}
	return *o.BackupFileName
}

// GetBackupFileNameOk returns a tuple with the BackupFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetBackupFileNameOk() (*string, bool) {
	if o == nil || o.BackupFileName == nil {
		return nil, false
	}
	return o.BackupFileName, true
}

// HasBackupFileName returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasBackupFileName() bool {
	if o != nil && o.BackupFileName != nil {
		return true
	}

	return false
}

// SetBackupFileName gets a reference to the given string and assigns it to the BackupFileName field.
func (o *UcsdBackupInfo) SetBackupFileName(v string) {
	o.BackupFileName = &v
}

// GetBackupLocation returns the BackupLocation field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetBackupLocation() string {
	if o == nil || o.BackupLocation == nil {
		var ret string
		return ret
	}
	return *o.BackupLocation
}

// GetBackupLocationOk returns a tuple with the BackupLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetBackupLocationOk() (*string, bool) {
	if o == nil || o.BackupLocation == nil {
		return nil, false
	}
	return o.BackupLocation, true
}

// HasBackupLocation returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasBackupLocation() bool {
	if o != nil && o.BackupLocation != nil {
		return true
	}

	return false
}

// SetBackupLocation gets a reference to the given string and assigns it to the BackupLocation field.
func (o *UcsdBackupInfo) SetBackupLocation(v string) {
	o.BackupLocation = &v
}

// GetBackupServerIp returns the BackupServerIp field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetBackupServerIp() string {
	if o == nil || o.BackupServerIp == nil {
		var ret string
		return ret
	}
	return *o.BackupServerIp
}

// GetBackupServerIpOk returns a tuple with the BackupServerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetBackupServerIpOk() (*string, bool) {
	if o == nil || o.BackupServerIp == nil {
		return nil, false
	}
	return o.BackupServerIp, true
}

// HasBackupServerIp returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasBackupServerIp() bool {
	if o != nil && o.BackupServerIp != nil {
		return true
	}

	return false
}

// SetBackupServerIp gets a reference to the given string and assigns it to the BackupServerIp field.
func (o *UcsdBackupInfo) SetBackupServerIp(v string) {
	o.BackupServerIp = &v
}

// GetBackupSize returns the BackupSize field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetBackupSize() int64 {
	if o == nil || o.BackupSize == nil {
		var ret int64
		return ret
	}
	return *o.BackupSize
}

// GetBackupSizeOk returns a tuple with the BackupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetBackupSizeOk() (*int64, bool) {
	if o == nil || o.BackupSize == nil {
		return nil, false
	}
	return o.BackupSize, true
}

// HasBackupSize returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasBackupSize() bool {
	if o != nil && o.BackupSize != nil {
		return true
	}

	return false
}

// SetBackupSize gets a reference to the given int64 and assigns it to the BackupSize field.
func (o *UcsdBackupInfo) SetBackupSize(v int64) {
	o.BackupSize = &v
}

// GetConnectors returns the Connectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UcsdBackupInfo) GetConnectors() []UcsdConnectorPack {
	if o == nil {
		var ret []UcsdConnectorPack
		return ret
	}
	return o.Connectors
}

// GetConnectorsOk returns a tuple with the Connectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UcsdBackupInfo) GetConnectorsOk() (*[]UcsdConnectorPack, bool) {
	if o == nil || o.Connectors == nil {
		return nil, false
	}
	return &o.Connectors, true
}

// HasConnectors returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasConnectors() bool {
	if o != nil && o.Connectors != nil {
		return true
	}

	return false
}

// SetConnectors gets a reference to the given []UcsdConnectorPack and assigns it to the Connectors field.
func (o *UcsdBackupInfo) SetConnectors(v []UcsdConnectorPack) {
	o.Connectors = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *UcsdBackupInfo) SetDuration(v int64) {
	o.Duration = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || o.EncryptionKey == nil {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey != nil {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *UcsdBackupInfo) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *UcsdBackupInfo) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetIsPurged returns the IsPurged field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetIsPurged() bool {
	if o == nil || o.IsPurged == nil {
		var ret bool
		return ret
	}
	return *o.IsPurged
}

// GetIsPurgedOk returns a tuple with the IsPurged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetIsPurgedOk() (*bool, bool) {
	if o == nil || o.IsPurged == nil {
		return nil, false
	}
	return o.IsPurged, true
}

// HasIsPurged returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasIsPurged() bool {
	if o != nil && o.IsPurged != nil {
		return true
	}

	return false
}

// SetIsPurged gets a reference to the given bool and assigns it to the IsPurged field.
func (o *UcsdBackupInfo) SetIsPurged(v bool) {
	o.IsPurged = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetLastModified() time.Time {
	if o == nil || o.LastModified == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || o.LastModified == nil {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasLastModified() bool {
	if o != nil && o.LastModified != nil {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *UcsdBackupInfo) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetPercentageCompletion returns the PercentageCompletion field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetPercentageCompletion() int64 {
	if o == nil || o.PercentageCompletion == nil {
		var ret int64
		return ret
	}
	return *o.PercentageCompletion
}

// GetPercentageCompletionOk returns a tuple with the PercentageCompletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetPercentageCompletionOk() (*int64, bool) {
	if o == nil || o.PercentageCompletion == nil {
		return nil, false
	}
	return o.PercentageCompletion, true
}

// HasPercentageCompletion returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasPercentageCompletion() bool {
	if o != nil && o.PercentageCompletion != nil {
		return true
	}

	return false
}

// SetPercentageCompletion gets a reference to the given int64 and assigns it to the PercentageCompletion field.
func (o *UcsdBackupInfo) SetPercentageCompletion(v int64) {
	o.PercentageCompletion = &v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetProductVersion() string {
	if o == nil || o.ProductVersion == nil {
		var ret string
		return ret
	}
	return *o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetProductVersionOk() (*string, bool) {
	if o == nil || o.ProductVersion == nil {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasProductVersion() bool {
	if o != nil && o.ProductVersion != nil {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given string and assigns it to the ProductVersion field.
func (o *UcsdBackupInfo) SetProductVersion(v string) {
	o.ProductVersion = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *UcsdBackupInfo) SetProtocol(v string) {
	o.Protocol = &v
}

// GetStageCompletion returns the StageCompletion field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetStageCompletion() string {
	if o == nil || o.StageCompletion == nil {
		var ret string
		return ret
	}
	return *o.StageCompletion
}

// GetStageCompletionOk returns a tuple with the StageCompletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetStageCompletionOk() (*string, bool) {
	if o == nil || o.StageCompletion == nil {
		return nil, false
	}
	return o.StageCompletion, true
}

// HasStageCompletion returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasStageCompletion() bool {
	if o != nil && o.StageCompletion != nil {
		return true
	}

	return false
}

// SetStageCompletion gets a reference to the given string and assigns it to the StageCompletion field.
func (o *UcsdBackupInfo) SetStageCompletion(v string) {
	o.StageCompletion = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *UcsdBackupInfo) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UcsdBackupInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UcsdBackupInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UcsdBackupInfo) SetStatus(v string) {
	o.Status = &v
}

func (o UcsdBackupInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRecoveryAbstractBackupInfo, errRecoveryAbstractBackupInfo := json.Marshal(o.RecoveryAbstractBackupInfo)
	if errRecoveryAbstractBackupInfo != nil {
		return []byte{}, errRecoveryAbstractBackupInfo
	}
	errRecoveryAbstractBackupInfo = json.Unmarshal([]byte(serializedRecoveryAbstractBackupInfo), &toSerialize)
	if errRecoveryAbstractBackupInfo != nil {
		return []byte{}, errRecoveryAbstractBackupInfo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BackupFileName != nil {
		toSerialize["BackupFileName"] = o.BackupFileName
	}
	if o.BackupLocation != nil {
		toSerialize["BackupLocation"] = o.BackupLocation
	}
	if o.BackupServerIp != nil {
		toSerialize["BackupServerIp"] = o.BackupServerIp
	}
	if o.BackupSize != nil {
		toSerialize["BackupSize"] = o.BackupSize
	}
	if o.Connectors != nil {
		toSerialize["Connectors"] = o.Connectors
	}
	if o.Duration != nil {
		toSerialize["Duration"] = o.Duration
	}
	if o.EncryptionKey != nil {
		toSerialize["EncryptionKey"] = o.EncryptionKey
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.IsPurged != nil {
		toSerialize["IsPurged"] = o.IsPurged
	}
	if o.LastModified != nil {
		toSerialize["LastModified"] = o.LastModified
	}
	if o.PercentageCompletion != nil {
		toSerialize["PercentageCompletion"] = o.PercentageCompletion
	}
	if o.ProductVersion != nil {
		toSerialize["ProductVersion"] = o.ProductVersion
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.StageCompletion != nil {
		toSerialize["StageCompletion"] = o.StageCompletion
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UcsdBackupInfo) UnmarshalJSON(bytes []byte) (err error) {
	type UcsdBackupInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Auto generated backup File Name with combination of file prefix given an user input and the timestamp.
		BackupFileName *string `json:"BackupFileName,omitempty"`
		// Backup location that contains the backup images for end device which can be used for restore operation.
		BackupLocation *string `json:"BackupLocation,omitempty"`
		// Backup server where backup images are maintained.
		BackupServerIp *string `json:"BackupServerIp,omitempty"`
		// Size of the backup image in bytes.
		BackupSize *int64              `json:"BackupSize,omitempty"`
		Connectors []UcsdConnectorPack `json:"Connectors,omitempty"`
		// Time taken for the backup to be completed.
		Duration *int64 `json:"Duration,omitempty"`
		// The key used for encrypting the backup file.
		EncryptionKey *string `json:"EncryptionKey,omitempty"`
		// Reason for backup failure.
		FailureReason *string `json:"FailureReason,omitempty"`
		// Backup image got purged or not. The backup images get purged based on the retention count set by the user in the backup config policy.
		IsPurged *bool `json:"IsPurged,omitempty"`
		// Last modified time when this backup record got updated.
		LastModified *time.Time `json:"LastModified,omitempty"`
		// Backup current precentage completion status information.
		PercentageCompletion *int64 `json:"PercentageCompletion,omitempty"`
		// The end device product version when the backup image was taken.
		ProductVersion *string `json:"ProductVersion,omitempty"`
		// Protocol used for the remote backup. possible values are FTP, SCP and SFTP. Not applicable for the localhost (127.0.0.1).
		Protocol *string `json:"Protocol,omitempty"`
		// Backup current status stage information.
		StageCompletion *string `json:"StageCompletion,omitempty"`
		// Start time of backup when it got initiated.
		StartTime *time.Time `json:"StartTime,omitempty"`
		// Current status of Backup current.
		Status *string `json:"Status,omitempty"`
	}

	varUcsdBackupInfoWithoutEmbeddedStruct := UcsdBackupInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUcsdBackupInfoWithoutEmbeddedStruct)
	if err == nil {
		varUcsdBackupInfo := _UcsdBackupInfo{}
		varUcsdBackupInfo.ClassId = varUcsdBackupInfoWithoutEmbeddedStruct.ClassId
		varUcsdBackupInfo.ObjectType = varUcsdBackupInfoWithoutEmbeddedStruct.ObjectType
		varUcsdBackupInfo.BackupFileName = varUcsdBackupInfoWithoutEmbeddedStruct.BackupFileName
		varUcsdBackupInfo.BackupLocation = varUcsdBackupInfoWithoutEmbeddedStruct.BackupLocation
		varUcsdBackupInfo.BackupServerIp = varUcsdBackupInfoWithoutEmbeddedStruct.BackupServerIp
		varUcsdBackupInfo.BackupSize = varUcsdBackupInfoWithoutEmbeddedStruct.BackupSize
		varUcsdBackupInfo.Connectors = varUcsdBackupInfoWithoutEmbeddedStruct.Connectors
		varUcsdBackupInfo.Duration = varUcsdBackupInfoWithoutEmbeddedStruct.Duration
		varUcsdBackupInfo.EncryptionKey = varUcsdBackupInfoWithoutEmbeddedStruct.EncryptionKey
		varUcsdBackupInfo.FailureReason = varUcsdBackupInfoWithoutEmbeddedStruct.FailureReason
		varUcsdBackupInfo.IsPurged = varUcsdBackupInfoWithoutEmbeddedStruct.IsPurged
		varUcsdBackupInfo.LastModified = varUcsdBackupInfoWithoutEmbeddedStruct.LastModified
		varUcsdBackupInfo.PercentageCompletion = varUcsdBackupInfoWithoutEmbeddedStruct.PercentageCompletion
		varUcsdBackupInfo.ProductVersion = varUcsdBackupInfoWithoutEmbeddedStruct.ProductVersion
		varUcsdBackupInfo.Protocol = varUcsdBackupInfoWithoutEmbeddedStruct.Protocol
		varUcsdBackupInfo.StageCompletion = varUcsdBackupInfoWithoutEmbeddedStruct.StageCompletion
		varUcsdBackupInfo.StartTime = varUcsdBackupInfoWithoutEmbeddedStruct.StartTime
		varUcsdBackupInfo.Status = varUcsdBackupInfoWithoutEmbeddedStruct.Status
		*o = UcsdBackupInfo(varUcsdBackupInfo)
	} else {
		return err
	}

	varUcsdBackupInfo := _UcsdBackupInfo{}

	err = json.Unmarshal(bytes, &varUcsdBackupInfo)
	if err == nil {
		o.RecoveryAbstractBackupInfo = varUcsdBackupInfo.RecoveryAbstractBackupInfo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupFileName")
		delete(additionalProperties, "BackupLocation")
		delete(additionalProperties, "BackupServerIp")
		delete(additionalProperties, "BackupSize")
		delete(additionalProperties, "Connectors")
		delete(additionalProperties, "Duration")
		delete(additionalProperties, "EncryptionKey")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "IsPurged")
		delete(additionalProperties, "LastModified")
		delete(additionalProperties, "PercentageCompletion")
		delete(additionalProperties, "ProductVersion")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "StageCompletion")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")

		// remove fields from embedded structs
		reflectRecoveryAbstractBackupInfo := reflect.ValueOf(o.RecoveryAbstractBackupInfo)
		for i := 0; i < reflectRecoveryAbstractBackupInfo.Type().NumField(); i++ {
			t := reflectRecoveryAbstractBackupInfo.Type().Field(i)

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

type NullableUcsdBackupInfo struct {
	value *UcsdBackupInfo
	isSet bool
}

func (v NullableUcsdBackupInfo) Get() *UcsdBackupInfo {
	return v.value
}

func (v *NullableUcsdBackupInfo) Set(val *UcsdBackupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUcsdBackupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUcsdBackupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcsdBackupInfo(val *UcsdBackupInfo) *NullableUcsdBackupInfo {
	return &NullableUcsdBackupInfo{value: val, isSet: true}
}

func (v NullableUcsdBackupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcsdBackupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
