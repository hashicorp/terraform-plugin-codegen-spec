# Terraform Provider Code Specification

> _[Terraform Provider Code Generation](https://developer.hashicorp.com/terraform/plugin/code-generation) is currently in tech preview._

## Overview

The Terraform Provider Code Specification is a versioned interface which is defined with a JSON schema, and has associated Go language bindings. The specification defines a [Terraform Provider](https://developer.hashicorp.com/terraform/plugin) representation which can be used for the generation of Go code for use in a provider, for example, with the [Plugin Framework Code Generator](https://developer.hashicorp.com/terraform/plugin/code-generation/framework-generator).

A specification can be generated from a declarative API model, such as OpenAPI with the [OpenAPI Provider Spec Generator](https://developer.hashicorp.com/terraform/plugin/code-generation/openapi-generator).

## Documentation

Full details of each of the fields defined within the JSON schema are available on the HashiCorp Developer site: https://developer.hashicorp.com//terraform/plugin/code-generation/specification.

## Versioning

### Schema

Schema versioning (e.g., [0.1](./spec/v0.1/schema.json)) follows the convention of using a _MAJOR.MINOR_ version, similar to the versioning used by [OpenAPI Specification](https://github.com/OAI/OpenAPI-Specification/blob/main/schemas/v3.1/schema.json) for example.

### Tags & Releases

Tags and releases follow the convention of [semantic versioning](https://semver.org/) adhering to _MAJOR.MINOR.PATCH_ versions. 

## Usage 

A JSON schema (e.g., [0.1](./spec/v0.1/schema.json)) defines the structure of a specification implementation by provider developers. For example:

```json
{
  "$id": "https://github.com/hashicorp/terraform-plugin-codegen-spec/spec/v0.1/schema.json",
  "$schema": "https://json-schema.org/draft-07/schema",
  "type": "object",
  "properties": {
    "datasources": {
      "type": "array",
      "items": {
        "$ref": "#/$defs/datasource"
      }
    },
    "provider": {
      "$ref": "#/$defs/provider"
    },
    "resources": {
      "type": "array",
      "items": {
        "$ref": "#/$defs/resource"
      }
    },
    "version": {
      "type": "string",
      "minLength": 3
    }
  },
  "required": [
    "provider",
    "version"
  ],
  ...
}
```

A corresponding specification implementation (e.g., [example.json](./spec/v0.1/example.json)) could look as follows:

```json
{
  "datasources": [
    { 
      ...
    }
  ],
  "provider": {
    ...
  },
  "resources": [
    {
      ...
    }
  ]
}
```

Refer to [example.json](./spec/v0.1/example.json) for an example specification.

## License

Refer to [Mozilla Public License v2.0](./LICENSE).

## Experimental Status

By using the software in this repository (the "Software"), you acknowledge that: (1) the Software is still in development, may change, and has not been released as a commercial product by HashiCorp and is not currently supported in any way by HashiCorp; (2) the Software is provided on an "as-is" basis, and may include bugs, errors, or other issues; (3) the Software is NOT INTENDED FOR PRODUCTION USE, use of the Software may result in unexpected results, loss of data, or other unexpected results, and HashiCorp disclaims any and all liability resulting from use of the Software; and (4) HashiCorp reserves all rights to make all decisions about the features, functionality and commercial release (or non-release) of the Software, at any time and without any obligation or liability whatsoever.
