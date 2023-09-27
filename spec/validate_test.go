// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spec_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
)

func TestValidate_Version0_1(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		document []byte
		expected error
	}{
		"nil": {
			document: nil,
			expected: fmt.Errorf("empty document"),
		},
		"empty": {
			document: []byte{},
			expected: fmt.Errorf("empty document"),
		},
		"unsupported_version": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
		"attributes": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  },
  "version": "a.b"
}`),
			expected: fmt.Errorf(`version: "a.b" is unsupported`),
		},
		"datasource-attributes-only": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
		"attributes": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  },
  "version": "0.1"
}`),
		},
		"datasource-blocks-only": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
		"blocks": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  },
  "version": "0.1"
}`),
		},
		"datasource-attributes-and-blocks": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
		"attributes": [],
		"blocks": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  },
  "version": "0.1"
}`),
		},
		"datasource-no-attributes-or-blocks": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
      }
    }
  ],
  "provider": {
    "name": "provider"
  },
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema: Must have at least 1 properties"),
		},
		"resource-attributes-only": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
		"attributes": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  },
  "version": "0.1"
}`),
		},
		"resource-blocks-only": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
		"blocks": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  },
  "version": "0.1"
}`),
		},
		"resource-attributes-and-blocks": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
		"attributes": [],
		"blocks": []
      }
    }
  ],
  "provider": {
    "name": "provider"
  },
  "version": "0.1"
}`),
		},
		"resource-no-attributes-or-blocks": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
      }
    }
  ],
  "provider": {
    "name": "provider"
  },
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema: Must have at least 1 properties"),
		},
		"example": {
			document: testReadFile("./v0.1/example.json"),
		},
		"datasource_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "Example"
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_name_invalid": {
			document: []byte(`{
  "provider": {
    "name": "Example"
  },
  "version": "0.1"
}`),
			expected: fmt.Errorf("provider.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "Example"
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_bool_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Bool_attribute",
            "bool": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_float64_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Float64_attribute",
            "float64": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_int64_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Int64_attribute",
            "int64": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_list_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "List_attribute",
            "list": {
              "computed_optional_required": "optional",
              "element_type": {
                "bool": {}
              }
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_list_nested_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "List_nested_attribute",
            "list_nested": {
              "computed_optional_required": "optional",
              "nested_object": {}
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_list_nested_block_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "blocks": [
          {
            "name": "List_nested_block",
            "list_nested": {
              "nested_object": {}
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.blocks.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_map_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Map_attribute",
            "map": {
              "computed_optional_required": "optional",
              "element_type": {
                "bool": {}
              }
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_map_nested_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Map_nested_attribute",
            "map_nested": {
              "computed_optional_required": "optional",
              "nested_object": {}
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_number_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Number_attribute",
            "number": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_object_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Object_attribute",
            "object": {
              "attribute_types": [
                {
                  "name": "str",
                  "string": {}
                }
              ],
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_set_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Set_attribute",
            "set": {
              "computed_optional_required": "optional",
              "element_type": {
                "bool": {}
              }
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_set_nested_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Set_nested_attribute",
            "set_nested": {
              "computed_optional_required": "optional",
              "nested_object": {}
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_set_nested_block_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "blocks": [
          {
            "name": "Set_nested_block",
            "bool": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.blocks.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_single_nested_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Single_nested_attribute",
            "single_nested": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_single_nested_block_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "blocks": [
          {
            "name": "Single_nested_block",
            "single_nested": {
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.blocks.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"datasource_string_attribute_name_invalid": {
			document: []byte(`{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "String_attribute",
            "string": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("datasources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_bool_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Bool_attribute",
            "bool": {
              "optional_required": "optional"
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_float64_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Float64_attribute",
            "float64": {
              "optional_required": "optional"
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_int64_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Int64_attribute",
            "int64": {
              "optional_required": "optional"
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_list_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "List_attribute",
            "list": {
              "optional_required": "optional",
              "element_type": {
                "bool": {}
              }
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_list_nested_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "List_nested_attribute",
            "list_nested": {
              "optional_required": "optional",
              "nested_object": {}
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_list_nested_block_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "blocks": [
          {
            "name": "List_nested_block",
            "list_nested": {
              "nested_object": {}
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.blocks.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_map_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Map_attribute",
            "map": {
              "optional_required": "optional",
              "element_type": {
                "bool": {}
              }
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_map_nested_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Map_nested_attribute",
            "map_nested": {
              "optional_required": "optional",
              "nested_object": {}
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_number_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Number_attribute",
            "number": {
              "optional_required": "optional"
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_object_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Object_attribute",
            "object": {
              "attribute_types": [
                {
                  "name": "str",
                  "string": {}
                }
              ],
              "optional_required": "optional"
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_set_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Set_attribute",
            "set": {
              "optional_required": "optional",
              "element_type": {
                "bool": {}
              }
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_set_nested_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Set_nested_attribute",
            "set_nested": {
              "optional_required": "optional",
              "nested_object": {}
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_set_nested_block_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "blocks": [
          {
            "name": "Set_nested_block",
            "bool": {
              "optional_required": "optional"
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.blocks.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_single_nested_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Single_nested_attribute",
            "single_nested": {
              "optional_required": "optional"
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_single_nested_block_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "blocks": [
          {
            "name": "Single_nested_block",
            "single_nested": {
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.blocks.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"provider_string_attribute_name_invalid": {
			document: []byte(`{
  "provider": {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "String_attribute",
            "string": {
              "optional_required": "optional"
            }
          }
        ]
      }
    },
  "version": "0.1"
  }`),
			expected: fmt.Errorf("provider.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_bool_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Bool_attribute",
            "bool": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_float64_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Float64_attribute",
            "float64": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_int64_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Int64_attribute",
            "int64": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_list_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "List_attribute",
            "list": {
              "computed_optional_required": "optional",
              "element_type": {
                "bool": {}
              }
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_list_nested_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "List_nested_attribute",
            "list_nested": {
              "computed_optional_required": "optional",
              "nested_object": {}
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_list_nested_block_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "blocks": [
          {
            "name": "List_nested_block",
            "list_nested": {
              "nested_object": {}
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.blocks.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_map_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Map_attribute",
            "map": {
              "computed_optional_required": "optional",
              "element_type": {
                "bool": {}
              }
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_map_nested_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Map_nested_attribute",
            "map_nested": {
              "computed_optional_required": "optional",
              "nested_object": {}
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_number_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Number_attribute",
            "number": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_object_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Object_attribute",
            "object": {
              "attribute_types": [
                {
                  "name": "str",
                  "string": {}
                }
              ],
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_set_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Set_attribute",
            "set": {
              "computed_optional_required": "optional",
              "element_type": {
                "bool": {}
              }
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_set_nested_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Set_nested_attribute",
            "set_nested": {
              "computed_optional_required": "optional",
              "nested_object": {}
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_set_nested_block_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "blocks": [
          {
            "name": "Set_nested_block",
            "bool": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.blocks.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_single_nested_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "Single_nested_attribute",
            "single_nested": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_single_nested_block_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "blocks": [
          {
            "name": "Single_nested_block",
            "single_nested": {
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.blocks.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
		"resource_string_attribute_name_invalid": {
			document: []byte(`{
  "resources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "String_attribute",
            "string": {
              "computed_optional_required": "optional"
            }
          }
        ]
      }
    }
  ],
  "version": "0.1"
}`),
			expected: fmt.Errorf("resources.0.schema.attributes.0.name: Does not match pattern '^[a-z_][a-z0-9_]*$'"),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := spec.Validate(context.Background(), testCase.document)

			if err != nil {
				if testCase.expected == nil {
					t.Fatalf("expected no error, got: %s", err)
				}

				if !strings.Contains(err.Error(), testCase.expected.Error()) {
					t.Fatalf("expected error %q, got: %s", testCase.expected, err)
				}
			}

			if err == nil && testCase.expected != nil {
				t.Fatalf("got no error, expected: %s", testCase.expected)
			}
		})
	}
}
