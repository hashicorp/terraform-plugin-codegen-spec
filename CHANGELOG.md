## 0.2.0 (September 09, 2024)

NOTES:

* all: This Go module has been updated to Go 1.22 per the [Go support policy](https://go.dev/doc/devel/release#policy). It is recommended to review the [Go 1.22 release notes](https://go.dev/doc/go1.22) before upgrading. Any consumers building on earlier Go versions may experience errors. ([#111](https://github.com/hashicorp/terraform-plugin-codegen-spec/issues/111))

FEATURES:

* Initial support of dynamic attributes in data source, provider, and resource schemas ([#98](https://github.com/hashicorp/terraform-plugin-codegen-spec/issues/98))

## 0.1.1 (April 04, 2024)

ENHANCEMENTS:

* Add associated external types for Bool, Float64, Int64, Number, and String Attributes for data sources, provider, and resources ([#62](https://github.com/hashicorp/terraform-plugin-codegen-spec/issues/62))
* Add associated external type for list, map, object, and set attributes ([#65](https://github.com/hashicorp/terraform-plugin-codegen-spec/issues/65))
* Add `CustomPlanModifiers` type and `CustomPlanModifiers()` function to bool, float64, int64, list, map, number, object, set, and string plan modifier types ([#77](https://github.com/hashicorp/terraform-plugin-codegen-spec/issues/77))
* Add `CustomValidators` type and `CustomValidators()` function to  bool, float64, int64, list, map, number, object, set, and string validator types ([#77](https://github.com/hashicorp/terraform-plugin-codegen-spec/issues/77))
* Add `CustomDefault()` function to list, map, number, object, and set default types ([#77](https://github.com/hashicorp/terraform-plugin-codegen-spec/issues/77))
* Added `Description`, `MarkdownDescription` and `DeprecationMessage` string fields to all schemas ([#81](https://github.com/hashicorp/terraform-plugin-codegen-spec/issues/81))

BUG FIXES:

* schema: Remove Framework-specific `requires_replace` and `use_state_for_unknown` plan modifiers from specification which are unused during code generation, use custom plan modifiers instead ([#84](https://github.com/hashicorp/terraform-plugin-codegen-spec/issues/84))

## 0.1.0 (October 16, 2023)

NOTES:

* Initial release of Provider Code Specification schema and Go bindings for Terraform Provider Code Generation tech preview ([#56](https://github.com/hashicorp/terraform-plugin-codegen-spec/issues/56))

