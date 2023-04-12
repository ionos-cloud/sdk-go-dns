# ZoneResponseMetadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The date of the last change formatted as yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;. | [optional] [readonly] |
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The date of creation of the zone formatted as yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;. | [optional] [readonly] |
|**State** | Pointer to [**ProvisioningState**](ProvisioningState.md) |  | [optional] |
|**Nameservers** | Pointer to **[]string** | The list of nameservers associated to the zone | [optional] |

## Methods

### NewZoneResponseMetadata

`func NewZoneResponseMetadata() *ZoneResponseMetadata`

NewZoneResponseMetadata instantiates a new ZoneResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneResponseMetadataWithDefaults

`func NewZoneResponseMetadataWithDefaults() *ZoneResponseMetadata`

NewZoneResponseMetadataWithDefaults instantiates a new ZoneResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModifiedDate

`func (o *ZoneResponseMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *ZoneResponseMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *ZoneResponseMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *ZoneResponseMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ZoneResponseMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ZoneResponseMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ZoneResponseMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ZoneResponseMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetState

`func (o *ZoneResponseMetadata) GetState() ProvisioningState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ZoneResponseMetadata) GetStateOk() (*ProvisioningState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ZoneResponseMetadata) SetState(v ProvisioningState)`

SetState sets State field to given value.

### HasState

`func (o *ZoneResponseMetadata) HasState() bool`

HasState returns a boolean if a field has been set.

### GetNameservers

`func (o *ZoneResponseMetadata) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *ZoneResponseMetadata) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *ZoneResponseMetadata) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *ZoneResponseMetadata) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.


