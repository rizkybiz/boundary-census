# Boundary Census

## Configuration

The configuration file for the server is specified as HCL, below is an example config file that contains all the possible
fields.

```hcl
config "controller" {
  nomad {
    address = "http://localhost:4646"
    token = "abc123" 
    region = "myregion"
    namespace = "mynamespace"
  }

  boundary {
    username = "nic"
    password = "password"
    address = "http://myaddress.com"

    org_id = "myorg"
    auth_method_id = "123"
    default_project = "hashicorp"
    default_groups = ["developers"]
  }
}
```

Environment variables can also be used to configure Boundary Census

|             ENV_VAR             | Description                                                                          |
|:-------------------------------:|--------------------------------------------------------------------------------------|
| NOMAD_ADDRESS                   | Address for Nomad server                                                             |
| NOMAD_TOKEN                     | Token for accessing Nomad (optional)                                                 |
| NOMAD_REGION                    | Nomad region (optional)                                                              |
| NOMAD_NAMESPACE                 | Nomad namespace (optional)                                                           |
| BOUNDARY_ENTERPRISE             | Boolean for Boundary enterprise (optional)                                           |
| BOUNDARY_ORG_ID                 | Boundary org ID                                                                      |
| BOUNDARY_DEFAULT_PROJECT        | The default Boundary project to place Nomad targets                                  |
| BOUNDARY_DEFAULT_GROUPS         | UNIMPLEMENTED                                                                        |
| BOUNDARY_AUTH_METHOD_ID         | ID of the Boundary auth method for the following username and password               |
| BOUNDARY_USERNAME               | Username of the Boundary Admin                                                       |
| BOUNDARY_PASSWORD               | Password of the Boundary Admin                                                       |
| BOUNDARY_ADDRESS                | Address of the Boundary control plane                                                |
| BOUNDARY_DEFAULT_INGRESS_FILTER | Ingress filter to apply when creating targets in boundary from Nomad Jobs (optional) |
| BOUNDARY_DEFAULT_EGRESS_FILTER  | Egress filter to apply when creating targets in boundary from Nomad Jobs (optional)  |

## Mocks

To re-generate the mock used for testing run the following command:

```shell
make generate_mocks
```


## Setup Local Nomad, Boundary, Consul

To setup and configure a local Nomad, Boundary, and Consul server use the following command:

```shell
shipyard run ./shipyard
```

You can determine the local addresses for the Boundary, Consul, and Nomad clusters by running:

```
shipyard output
```

These can also be set as environment variables with the following command:

```
eval $(shipyard env)
```

## Running

To run the server you can use the following command:

```shell
go run main.go -config=./example_config.hcl
```
