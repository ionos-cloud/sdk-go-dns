# ZoneResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The zone ID (UUID). | [optional] |
|**Metadata** | Pointer to [**ZoneResponseMetadata**](ZoneResponseMetadata.md) |  | [optional] |
|**Properties** | Pointer to [**ZoneResponseProperties**](ZoneResponseProperties.md) |  | [optional] |

## Methods

### NewZoneResponse

`func NewZoneResponse() *ZoneResponse`

NewZoneResponse instantiates a new ZoneResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneResponseWithDefaults

`func NewZoneResponseWithDefaults() *ZoneResponse`

NewZoneResponseWithDefaults instantiates a new ZoneResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ZoneResponse) GetMetadata() ZoneResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ZoneResponse) GetMetadataOk() (*ZoneResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ZoneResponse) SetMetadata(v ZoneResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ZoneResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *ZoneResponse) GetProperties() ZoneResponseProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ZoneResponse) GetPropertiesOk() (*ZoneResponseProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ZoneResponse) SetProperties(v ZoneResponseProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ZoneResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


