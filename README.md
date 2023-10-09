# Terraform Provider Code Specification

> _[Terraform Provider Code Generation](https://developer.hashicorp.com/terraform/plugin/code-generation) is currently in tech preview._

## Overview

The Terraform Provider Code Specification is a versioned interface which is defined with a JSON schema, and has associated Go language bindings. The specification defines a [Terraform Provider](https://developer.hashicorp.com/terraform/plugin) representation which can be used for the generation of Go code for use in a provider, for example, with the [Plugin Framework Code Generator](https://developer.hashicorp.com/terraform/plugin/code-generation/framework-generator).

A specification can be generated from a declarative API model, such as OpenAPI with the [OpenAPI Provider Spec Generator](https://developer.hashicorp.com/terraform/plugin/code-generation/openapi-generator).

## Documentation

Full details of each of the fields defined within the JSON schema are available on the HashiCorp Developer site: https://developer.hashicorp.com//terraform/plugin/code-generation/specification.

## Versioning

### Schema

Schema versioning (e.g., [v0.1](./spec/v0.1/schema.json)) follows the convention of using a _MAJOR.MINOR_ version, similar to the versioning used by [OpenAPI Specification](https://github.com/OAI/OpenAPI-Specification/blob/main/schemas/v3.1/schema.json) for example.

### Tags & Releases

Tags and releases follow the convention of [semantic versioning](https://semver.org/) adhering to _MAJOR.MINOR.PATCH_ versions. 

## Usage 

The JSON schema (e.g., [v0.1](./spec/v0.1/schema.json)) defines the structure of a specification. For example:

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
      "minLength": 1
    }
  },
  "required": [
    "provider",
    "version"
  ],
  ...
}
```

A corresponding specification (e.g., [example.json](./spec/v0.1/example.json)) could look as follows:

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