# RecordMetadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The date of the last change formatted as yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39; | [optional] [readonly] |
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The date of the record creation formatted as yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39; | [optional] [readonly] |
|**State** | Pointer to [**ProvisioningState**](ProvisioningState.md) |  | [optional] |
|**Fqdn** | Pointer to **string** | A fully qualified domain name. FQDN consists of two parts - the hostname and the domain name. | [optional] [readonly] |
|**ZoneId** | Pointer to **string** | The ID (UUID) of the DNS zone of which record belongs to. | [optional] [readonly] |

## Methods

### NewRecordMetadata

`func NewRecordMetadata() *RecordMetadata`

NewRecordMetadata instantiates a new RecordMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordMetadataWithDefaults

`func NewRecordMetadataWithDefaults() *RecordMetadata`

NewRecordMetadataWithDefaults instantiates a new RecordMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModifiedDate

`func (o *RecordMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *RecordMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *RecordMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *RecordMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetCreatedDate

`func (o *RecordMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *RecordMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *RecordMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *RecordMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetState

`func (o *RecordMetadata) GetState() ProvisioningState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RecordMetadata) GetStateOk() (*ProvisioningState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RecordMetadata) SetState(v ProvisioningState)`

SetState sets State field to given value.

### HasState

`func (o *RecordMetadata) HasState() bool`

HasState returns a boolean if a field has been set.

### GetFqdn

`func (o *RecordMetadata) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *RecordMetadata) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *RecordMetadata) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *RecordMetadata) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetZoneId

`func (o *RecordMetadata) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *RecordMetadata) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *RecordMetadata) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *RecordMetadata) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


