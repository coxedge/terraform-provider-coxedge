# Cox Edge Terraform Provider

## Build and Use
The source code repository contains a script to build and install
the Cox Edge Terraform Provider during testing. It is located at the
repository and is named `compile-install.sh`.

Terraform Providers are simply standalone go executables. Therefore,
the build, management, versioning, and testing of these providers
follow standard go conventions. The compile and install script 
simply calls `go build` and copies the binary to the required
location as specified in the Terraform documentation 
[here]().

## Using Terraform Plugin SDKv2
Hashicorp has a robust getting started documentation set on 
developing plugins/providers 
[here](https://learn.hashicorp.com/tutorials/terraform/provider-use?in=terraform/providers).

## Source Code Structure
The Cox Edge Terraform Provider has two main components. 
First, the Go API Client for the Cox Edge API. Second, the
Terraform Provider wrapper for the Cox Edge API Client.

The Hashicorp Terraform Plugin SDK is used as the basis for 
interacting with Terraform. Its source can be found 
[here](https://github.com/hashicorp/terraform-plugin-sdk).
The documentation from Hashicorp for using this SDK can
be found [here](https://www.terraform.io/plugin). Please note
that this provider is built using the Plugin SDKv2 not the 
Terraform Plugin Framework.

### Cox Edge Go Client
The code for this component is found within `coxedge/apiclient`.
The `models.go` file contains all Go struct representations of 
the structures within the Cox Edge API. The `client.go` file contains
the client setup including URLs. 

The Go API Client has code-native test cases written for it to validate
functionality. Some variables within the test may be changed to suite
the runtime environment.

### Cox Edge Terraform Provider
The code for this component is within `coxedge`. The `provider.go` 
file serves as the entrypoint for the provider and list the 
resources that the Terraform Provider can manage. The `tf_schemas.go`
file outlines the schema that users of the provider will interact with
as they write their code.

#### Notes on Schemas
- All schema objects must have either `optional`, `required`, 
or `computed` set.
- Terraform will validate the `tf_schema.go` at runtime. 
Issues in this file will prevent the provider from running.

## Additional Notes
### Configuring Log Output
Terraform's log output can be configured with the `TF_LOG` environment
variable. The different values and how the interact with the SDK can 
be found [here](https://www.terraform.io/plugin/log/managing).

### Terraform Provider Debugging Guide
Hashicorp has a specific section on debugging providers 
[here](https://www.terraform.io/plugin/sdkv2/debugging).

### Example Terraform Code
The following code shows an example use case and the required syntax.
```terraform
terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "[INSERT API KEY HERE]"
}

data "coxedge_environments" "test" {

}

# Workloads
resource "coxedge_workload" "test" {
  name = "test"
  environment_name = data.coxedge_environments.test.environments[0].name
  type = "CONTAINER"
  image = "ubuntu:latest"
  specs = "SP-1"
  deployment {
    name = "test"
    enable_autoscaling = false
    pops = ["MIA"]
    instances_per_pop = 1
  }
}
```

### Data Marshaling
The most common area of concern when tracing issues is the translation
of data types across API boundaries. Within the resources in this provider
(the `resource_*.go` files), this is usually handled by 
`convertResourceDataTo[OBJECTYPE]APIObject` and 
`convert[OBJECTYPE]APIObjectToResourceData`. These are good places to
start when attempt to hunt down a bug.

### API Data Flow
The flow below outlines how data is transform from terraform file
to Cox Edge API call.
1. User creates terraform .tf file.
2. Terraform parses .tf file and converts to defined schema.
3. Terraform validates data entered against schema.
4. Terraform passes data to `resource_*.go` or `data_source_*.go`.
5. Schema is converted to internal Cox Edge API struct.
6. Cox Edge Go struct is serialized as JSON and sent to API.