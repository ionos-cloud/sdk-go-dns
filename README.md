# Go API client for ionoscloud

Cloud DNS service helps IONOS Cloud customers to automate DNS Zone and Record management.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.15.4
- Package version: 1.2.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://docs.ionos.com/support/general-information/contact-information](https://docs.ionos.com/support/general-information/contact-information)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import ionoscloud "github.com/ionos-cloud/go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), ionoscloud.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), ionoscloud.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

## Documentation for API Endpoints

All URIs are relative to *https://dns.de-fra.ionos.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DNSSECApi* | [**ZonesKeysDelete**](docs/api/DNSSECApi.md#zoneskeysdelete) | **Delete** /zones/{zoneId}/keys | Delete a DNSSEC key
*DNSSECApi* | [**ZonesKeysGet**](docs/api/DNSSECApi.md#zoneskeysget) | **Get** /zones/{zoneId}/keys | Retrieve a DNSSEC key
*DNSSECApi* | [**ZonesKeysPost**](docs/api/DNSSECApi.md#zoneskeyspost) | **Post** /zones/{zoneId}/keys | Create a DNSSEC key
*QuotaApi* | [**QuotaGet**](docs/api/QuotaApi.md#quotaget) | **Get** /quota | Retrieve resources quota
*RecordsApi* | [**RecordsGet**](docs/api/RecordsApi.md#recordsget) | **Get** /records | Retrieve all records from primary zones
*RecordsApi* | [**SecondaryzonesRecordsGet**](docs/api/RecordsApi.md#secondaryzonesrecordsget) | **Get** /secondaryzones/{secondaryZoneId}/records | Retrieve records for a secondary zone
*RecordsApi* | [**ZonesRecordsDelete**](docs/api/RecordsApi.md#zonesrecordsdelete) | **Delete** /zones/{zoneId}/records/{recordId} | Delete a record
*RecordsApi* | [**ZonesRecordsFindById**](docs/api/RecordsApi.md#zonesrecordsfindbyid) | **Get** /zones/{zoneId}/records/{recordId} | Retrieve a record
*RecordsApi* | [**ZonesRecordsGet**](docs/api/RecordsApi.md#zonesrecordsget) | **Get** /zones/{zoneId}/records | Retrieve records
*RecordsApi* | [**ZonesRecordsPost**](docs/api/RecordsApi.md#zonesrecordspost) | **Post** /zones/{zoneId}/records | Create a record
*RecordsApi* | [**ZonesRecordsPut**](docs/api/RecordsApi.md#zonesrecordsput) | **Put** /zones/{zoneId}/records/{recordId} | Update a record
*ReverseRecordsApi* | [**ReverserecordsDelete**](docs/api/ReverseRecordsApi.md#reverserecordsdelete) | **Delete** /reverserecords/{reverserecordId} | Delete a reverse DNS record
*ReverseRecordsApi* | [**ReverserecordsFindById**](docs/api/ReverseRecordsApi.md#reverserecordsfindbyid) | **Get** /reverserecords/{reverserecordId} | Retrieve a reverse DNS record
*ReverseRecordsApi* | [**ReverserecordsGet**](docs/api/ReverseRecordsApi.md#reverserecordsget) | **Get** /reverserecords | Retrieves existing reverse DNS records
*ReverseRecordsApi* | [**ReverserecordsPost**](docs/api/ReverseRecordsApi.md#reverserecordspost) | **Post** /reverserecords | Create a reverse DNS record
*ReverseRecordsApi* | [**ReverserecordsPut**](docs/api/ReverseRecordsApi.md#reverserecordsput) | **Put** /reverserecords/{reverserecordId} | Update a reverse DNS record
*SecondaryZonesApi* | [**SecondaryzonesAxfrGet**](docs/api/SecondaryZonesApi.md#secondaryzonesaxfrget) | **Get** /secondaryzones/{secondaryZoneId}/axfr | Get status of zone transfer
*SecondaryZonesApi* | [**SecondaryzonesAxfrPut**](docs/api/SecondaryZonesApi.md#secondaryzonesaxfrput) | **Put** /secondaryzones/{secondaryZoneId}/axfr | Start zone transfer
*SecondaryZonesApi* | [**SecondaryzonesDelete**](docs/api/SecondaryZonesApi.md#secondaryzonesdelete) | **Delete** /secondaryzones/{secondaryZoneId} | Delete a secondary zone
*SecondaryZonesApi* | [**SecondaryzonesFindById**](docs/api/SecondaryZonesApi.md#secondaryzonesfindbyid) | **Get** /secondaryzones/{secondaryZoneId} | Retrieve a secondary zone
*SecondaryZonesApi* | [**SecondaryzonesGet**](docs/api/SecondaryZonesApi.md#secondaryzonesget) | **Get** /secondaryzones | Retrieve secondary zones
*SecondaryZonesApi* | [**SecondaryzonesPost**](docs/api/SecondaryZonesApi.md#secondaryzonespost) | **Post** /secondaryzones | Create a secondary zone
*SecondaryZonesApi* | [**SecondaryzonesPut**](docs/api/SecondaryZonesApi.md#secondaryzonesput) | **Put** /secondaryzones/{secondaryZoneId} | Update a secondary zone
*ZoneFilesApi* | [**ZonesZonefileGet**](docs/api/ZoneFilesApi.md#zoneszonefileget) | **Get** /zones/{zoneId}/zonefile | Retrieve a zone file
*ZoneFilesApi* | [**ZonesZonefilePut**](docs/api/ZoneFilesApi.md#zoneszonefileput) | **Put** /zones/{zoneId}/zonefile | Updates a zone with a file
*ZonesApi* | [**ZonesDelete**](docs/api/ZonesApi.md#zonesdelete) | **Delete** /zones/{zoneId} | Delete a zone
*ZonesApi* | [**ZonesFindById**](docs/api/ZonesApi.md#zonesfindbyid) | **Get** /zones/{zoneId} | Retrieve a zone
*ZonesApi* | [**ZonesGet**](docs/api/ZonesApi.md#zonesget) | **Get** /zones | Retrieve zones
*ZonesApi* | [**ZonesPost**](docs/api/ZonesApi.md#zonespost) | **Post** /zones | Create a zone
*ZonesApi* | [**ZonesPut**](docs/api/ZonesApi.md#zonesput) | **Put** /zones/{zoneId} | Update a zone


## Documentation For Models

 - [Algorithm](docs/models/Algorithm.md)
 - [CommonZone](docs/models/CommonZone.md)
 - [CommonZoneRead](docs/models/CommonZoneRead.md)
 - [CommonZoneReadList](docs/models/CommonZoneReadList.md)
 - [DnssecKey](docs/models/DnssecKey.md)
 - [DnssecKeyCreate](docs/models/DnssecKeyCreate.md)
 - [DnssecKeyParameters](docs/models/DnssecKeyParameters.md)
 - [DnssecKeyReadCreation](docs/models/DnssecKeyReadCreation.md)
 - [DnssecKeyReadList](docs/models/DnssecKeyReadList.md)
 - [DnssecKeyReadListMetadata](docs/models/DnssecKeyReadListMetadata.md)
 - [DnssecKeyReadListProperties](docs/models/DnssecKeyReadListProperties.md)
 - [DnssecKeyReadListPropertiesKeyParameters](docs/models/DnssecKeyReadListPropertiesKeyParameters.md)
 - [DnssecKeyReadListPropertiesNsecParameters](docs/models/DnssecKeyReadListPropertiesNsecParameters.md)
 - [Error](docs/models/Error.md)
 - [ErrorMessages](docs/models/ErrorMessages.md)
 - [KeyData](docs/models/KeyData.md)
 - [KeyParameters](docs/models/KeyParameters.md)
 - [KskBits](docs/models/KskBits.md)
 - [Links](docs/models/Links.md)
 - [Metadata](docs/models/Metadata.md)
 - [MetadataForSecondaryZoneRecords](docs/models/MetadataForSecondaryZoneRecords.md)
 - [MetadataWithStateFqdnZoneId](docs/models/MetadataWithStateFqdnZoneId.md)
 - [MetadataWithStateFqdnZoneIdAllOf](docs/models/MetadataWithStateFqdnZoneIdAllOf.md)
 - [MetadataWithStateNameservers](docs/models/MetadataWithStateNameservers.md)
 - [MetadataWithStateNameserversAllOf](docs/models/MetadataWithStateNameserversAllOf.md)
 - [NsecMode](docs/models/NsecMode.md)
 - [NsecParameters](docs/models/NsecParameters.md)
 - [ProvisioningState](docs/models/ProvisioningState.md)
 - [Quota](docs/models/Quota.md)
 - [QuotaDetail](docs/models/QuotaDetail.md)
 - [Record](docs/models/Record.md)
 - [RecordCreate](docs/models/RecordCreate.md)
 - [RecordEnsure](docs/models/RecordEnsure.md)
 - [RecordRead](docs/models/RecordRead.md)
 - [RecordReadList](docs/models/RecordReadList.md)
 - [ReverseRecord](docs/models/ReverseRecord.md)
 - [ReverseRecordCreate](docs/models/ReverseRecordCreate.md)
 - [ReverseRecordEnsure](docs/models/ReverseRecordEnsure.md)
 - [ReverseRecordRead](docs/models/ReverseRecordRead.md)
 - [ReverseRecordsReadList](docs/models/ReverseRecordsReadList.md)
 - [SecondaryZone](docs/models/SecondaryZone.md)
 - [SecondaryZoneAllOf](docs/models/SecondaryZoneAllOf.md)
 - [SecondaryZoneCreate](docs/models/SecondaryZoneCreate.md)
 - [SecondaryZoneEnsure](docs/models/SecondaryZoneEnsure.md)
 - [SecondaryZoneRead](docs/models/SecondaryZoneRead.md)
 - [SecondaryZoneReadAllOf](docs/models/SecondaryZoneReadAllOf.md)
 - [SecondaryZoneReadList](docs/models/SecondaryZoneReadList.md)
 - [SecondaryZoneReadListAllOf](docs/models/SecondaryZoneReadListAllOf.md)
 - [SecondaryZoneRecordRead](docs/models/SecondaryZoneRecordRead.md)
 - [SecondaryZoneRecordReadList](docs/models/SecondaryZoneRecordReadList.md)
 - [SecondaryZoneRecordReadListMetadata](docs/models/SecondaryZoneRecordReadListMetadata.md)
 - [Zone](docs/models/Zone.md)
 - [ZoneAllOf](docs/models/ZoneAllOf.md)
 - [ZoneCreate](docs/models/ZoneCreate.md)
 - [ZoneEnsure](docs/models/ZoneEnsure.md)
 - [ZoneRead](docs/models/ZoneRead.md)
 - [ZoneReadAllOf](docs/models/ZoneReadAllOf.md)
 - [ZoneReadList](docs/models/ZoneReadList.md)
 - [ZoneReadListAllOf](docs/models/ZoneReadListAllOf.md)
 - [ZoneTransferPrimaryIpStatus](docs/models/ZoneTransferPrimaryIpStatus.md)
 - [ZoneTransferPrimaryIpsStatus](docs/models/ZoneTransferPrimaryIpsStatus.md)
 - [ZskBits](docs/models/ZskBits.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### tokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@cloud.ionos.com

