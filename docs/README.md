# Introduction

## Overview

The IONOS Cloud SDK for GO provides you with access to the IONOS Cloud DNS. The client library supports both simple and complex requests. It is designed for developers who are building applications in GO. All API operations are performed over SSL and authenticated using your IONOS Cloud portal credentials. The API can be accessed within an instance running in IONOS Cloud or directly over the Internet from any application that can send an HTTPS request and receive an HTTPS response.

DNS allows users to publish DNS Zones of their domains and subdomains on public Name Servers.
IONOS Name Server Infrastructure is distributed across 14 point-of-presence (POP) locations in Europe and U.S.A. to ensure fast and reliable DNS resolution for users wherever they are located.
With DNS, you can manage your DNS Zones and Records programmatically via API.
More details about features and concepts can be read [here](https://docs.ionos.com/dns-as-a-service/readme/overview).

**NOTE:** Cloud DNS is currently in the **Early Access (EA)** phase. We recommend keeping usage and testing to non-production critical applications. Please contact your sales representative or support for more information.

## Getting Started

Before you begin you will need to have [signed-up](https://www.ionos.com/enterprise-cloud/signup) for a IONOS Cloud account. The credentials you setup during sign-up will be used to authenticate against the API.

### Installation

Install the Go language from the official [Go installation](https://golang.org/doc/install) guide.

The `GOPATH` environment variable specifies the location of your Go workspace. It is likely the only environment variable you will need to set when developing Go code. This is an example of pointing to a workspace configured under your home directory:

```text
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

The following `go` command will download `sdk-go-dns` to your configured `GOPATH`:

```go
go get "github.com/ionos-cloud/sdk-go-dns"
```

The source code of the package will be located here:

```text
$GOPATH/src/github.com/ionos-cloud/sdk-go
```

Create main package file _example.go_:

```go
package main

import (
    "fmt"
)

func main() {
}
```

Include the IONOS Cloud SDK for Go under the list of imports.

```go
import(
    "fmt"

    "github.com/ionos-cloud/sdk-go-dns"
)
```

### Authentication

The authentication token can be manually specified when initializing the SDK client:

```go
client := ionoscloud.NewAPIClient(ionoscloud.NewConfiguration(token, apiUrl))
```

Environment variables can also be used. The sdk uses the following variables:

* IONOS\_TOKEN - if an authentication token is being used
* IONOS\_API\_URL - to overwrite the api endpoint: `dns.de-fra.ionos.com` - if it is not set, the default value will be used

In this case, the client configuration needs to be initialized using `NewConfigurationFromEnv()`

```go
client := ionoscloud.NewAPIClient(ionoscloud.NewConfigurationFromEnv())
```

**Caution**: You will want to ensure you follow security best practices when using credentials within your code or stored in a file.

### Environment Variables

| Environment Variable | Description                                                                                                                                                                                                                           |
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `IONOS_TOKEN`        | Specify the token used to login.                                                                                                                                                                                                      |
| `IONOS_API_URL`      | Specify the API URL. It will overwrite the API endpoint default value `dns.de-fra.ionos.com`. Note: the host URL does not contain the `/cloudapi/v5` path, so it should _not_ be included in the `IONOS_API_URL` environment variable | 


### Changing the base URL

Changing the base URL for the HTTP operation is possible by using the following function:

```go
requestProperties.SetURL("https://dns.de-fra.ionos.com")
```

## Feature Reference

The IONOS Cloud SDK for GO aims to offer access to all resources in the IONOS Cloud API and also offers some additional features that make the integration easier:

* authentication for API calls
* handling of asynchronous requests 

## FAQ

1. How can I open a bug/feature request? 

Bugs & feature requests can be open on the repository issues: [https://github.com/ionos-cloud/sdk-go-dns/issues/new/choose](https://github.com/ionos-cloud/sdk-go-dns/issues/new/choose)

1. Can I contribute to the GO SDK?

Pure SDKs are automatically generated using OpenAPI Generator and don’t support manual changes. If you need changes please open an issue and we’ll try to take care of it.

## Debugging

If you want to see the API call request and response messages, you need to set the Debug field in the Configuration struct:

```golang
package main
import "github.com/ionos-cloud/sdk-go-dns"
func main() {
    // create your configuration. replace token and url with correct values, or use NewConfigurationFromEnv()
    // if you have set your env variables as explained above
    cfg := ionoscloud.NewConfiguration("token", "hostUrl")
    // enable request and response logging
    cfg.Debug = true
    // create you api client with the configuration
    apiClient := ionoscloud.NewAPIClient(cfg)
}
```

⚠️ **_Note: We recommend you only set this field for debugging purposes. Disable it in your production environments because it can log sensitive data. It logs the full request and response without encryption, even for an HTTPS call. Verbose request and response logging can also significantly impact your application’s performance._**
