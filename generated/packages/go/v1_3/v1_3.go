package v1_3

import "encoding/json"
import "errors"
// This string MUST be the [semantic version number](https://semver.org/spec/v2.0.0.html) of the [OpenRPC Specification version](#versions) that the OpenRPC document uses. The `openrpc` field SHOULD be used by tooling specifications and clients to interpret the OpenRPC document. This is *not* related to the API [`info.version`](#info-version) string.
type Openrpc string
// The title of the application.
type InfoObjectTitle string
// A verbose description of the application. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
type InfoObjectDescription string
// A URL to the Terms of Service for the API. MUST be in the format of a URL.
type InfoObjectTermsOfService string
// The version of the OpenRPC document (which is distinct from the [OpenRPC Specification version](#openrpc-version) or the API implementation version).
type InfoObjectVersion string
// The identifying name of the contact person/organization.
type ContactObjectName string
// The email address of the contact person/organization. MUST be in the format of an email address.
type ContactObjectEmail string
// The URL pointing to the contact information. MUST be in the format of a URL.
type ContactObjectUrl string
// This object MAY be extended with [Specification Extensions](#specification-extensions).
type SpecificationExtension interface{}
// Contact information for the exposed API.
type ContactObject struct {
	Name  *ContactObjectName  `json:"name,omitempty"`
	Email *ContactObjectEmail `json:"email,omitempty"`
	Url   *ContactObjectUrl   `json:"url,omitempty"`
}
// The license name used for the API.
type LicenseObjectName string
// A URL to the license used for the API. MUST be in the format of a URL.
type LicenseObjectUrl string
// License information for the exposed API.
type LicenseObject struct {
	Name *LicenseObjectName `json:"name,omitempty"`
	Url  *LicenseObjectUrl  `json:"url,omitempty"`
}
// The object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.
type InfoObject interface{}
// A verbose explanation of the documentation. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
type ExternalDocumentationObjectDescription string
// The URL for the target documentation. Value MUST be in the format of a URL.
type ExternalDocumentationObjectUrl string
// Additional external documentation for this tag.
type ExternalDocumentationObject struct {
	Description *ExternalDocumentationObjectDescription `json:"description,omitempty"`
	Url         *ExternalDocumentationObjectUrl         `json:"url"`
}
// A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenRPC document is being served. [Server Variables](#server-variables) are passed into the [Runtime Expression](#runtime-expression) to produce a server URL.
type ServerObjectUrl string
// An optional string describing the name of the server. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
type ServerObjectName string
// An optional string describing the host designated by the URL. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
type ServerObjectDescription string
// A short summary of what the server is.
type ServerObjectSummary string
// The default value to use for substitution, which SHALL be sent if an alternate value is _not_ supplied. Note this behavior is different than the [Schema Object's](#schema-object) treatment of default values, because in those cases parameter values are optional.
type ServerObjectVariableDefault string
// An optional description for the server variable. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
type ServerObjectVariableDescription string
// An enumeration of string values to be used if the substitution options are from a limited set.
type ServerObjectVariableEnumItem string
// An enumeration of string values to be used if the substitution options are from a limited set.
type ServerObjectVariableEnum []ServerObjectVariableEnumItem
// An object representing a Server Variable for server URL template substitution.
type ServerObjectVariable struct {
	Default     *ServerObjectVariableDefault     `json:"default"`
	Description *ServerObjectVariableDescription `json:"description,omitempty"`
	Enum        *ServerObjectVariableEnum        `json:"enum,omitempty"`
}
// A map between a variable name and its value. The value is passed into the [Runtime Expression](#runtime-expression) to produce a server URL.
type ServerObjectVariables map[string]interface{}
// A object representing a Server
type ServerObject struct {
	Url         *ServerObjectUrl         `json:"url"`
	Name        *ServerObjectName        `json:"name,omitempty"`
	Description *ServerObjectDescription `json:"description,omitempty"`
	Summary     *ServerObjectSummary     `json:"summary,omitempty"`
	Variables   *ServerObjectVariables   `json:"variables,omitempty"`
}
type AlwaysFalse interface{}
// An array of Server Objects, which provide connectivity information to a target server. If the `servers` property is not provided, or is an empty array, the default value would be a [Server Object](#server-object) with a [url](#server-url) value of `localhost`. 
type Servers []ServerObject
// The cannonical name for the method. The name MUST be unique within the methods array.
type MethodObjectName string
// A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.
type MethodObjectDescription string
// A short summary of what the method does.
type MethodObjectSummary string
// The name of the tag.
type TagObjectName string
// A verbose explanation for the tag. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
type TagObjectDescription string
// Adds metadata to a single tag that is used by the [Method Object](#method-object). It is not mandatory to have a Tag Object per tag defined in the Method Object instances.
type TagObject struct {
	Name         *TagObjectName               `json:"name"`
	Description  *TagObjectDescription        `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentationObject `json:"externalDocs,omitempty"`
}
type Ref string
type ReferenceObject struct {
	Ref *Ref `json:"$ref"`
}
type TagOrReference struct {
	TagObject       *TagObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *TagOrReference) UnmarshalJSON(bytes []byte) error {
	var myTagObject TagObject
	if err := json.Unmarshal(bytes, &myTagObject); err == nil {
		o.TagObject = &myTagObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o TagOrReference) MarshalJSON() ([]byte, error) {
	if o.TagObject != nil {
		return json.Marshal(o.TagObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// A list of tags for API documentation control. Tags can be used for logical grouping of methods by resources or any other qualifier.
type MethodObjectTags []TagOrReference
// Format the server expects the params. Defaults to 'either'.
//
// --- Default ---
//
// either
type MethodObjectParamStructure string
const (
	MethodObjectParamStructureEnum0 MethodObjectParamStructure = "by-position"
	MethodObjectParamStructureEnum1 MethodObjectParamStructure = "by-name"
	MethodObjectParamStructureEnum2 MethodObjectParamStructure = "either"
)
// Name of the content that is being described. If the content described is a method parameter assignable [`by-name`](#method-param-structure), this field SHALL define the parameter's key (ie name).
type ContentDescriptorObjectName string
// A verbose explanation of the content descriptor behavior. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
type ContentDescriptorObjectDescription string
// A short summary of the content that is being described.
type ContentDescriptorObjectSummary string
type Id string
type Schema string
type Comment string
type Title string
type Description string
type AlwaysTrue interface{}
type ReadOnly bool
type Examples []AlwaysTrue
type MultipleOf float64
type Maximum float64
type ExclusiveMaximum float64
type Minimum float64
type ExclusiveMinimum float64
type NonNegativeInteger int64
type NonNegativeIntegerDefaultZero int64
type Pattern string
// Always valid if true. Never valid if false. Is constant.
type JSONSchemaBoolean bool
//
// --- Default ---
//
// {}
type JSONSchema struct {
	JSONSchemaObject  *JSONSchemaObject
	JSONSchemaBoolean *JSONSchemaBoolean
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *JSONSchema) UnmarshalJSON(bytes []byte) error {
	var myJSONSchemaObject JSONSchemaObject
	if err := json.Unmarshal(bytes, &myJSONSchemaObject); err == nil {
		o.JSONSchemaObject = &myJSONSchemaObject
		return nil
	}
	var myJSONSchemaBoolean JSONSchemaBoolean
	if err := json.Unmarshal(bytes, &myJSONSchemaBoolean); err == nil {
		o.JSONSchemaBoolean = &myJSONSchemaBoolean
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o JSONSchema) MarshalJSON() ([]byte, error) {
	if o.JSONSchemaObject != nil {
		return json.Marshal(o.JSONSchemaObject)
	}
	if o.JSONSchemaBoolean != nil {
		return json.Marshal(o.JSONSchemaBoolean)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type SchemaArray []JSONSchema
//
// --- Default ---
//
// true
type Items struct {
	JSONSchema  *JSONSchema
	SchemaArray *SchemaArray
}
func (a *Items) UnmarshalJSON(bytes []byte) error {
	var ok bool
	var myJSONSchema JSONSchema
	if err := json.Unmarshal(bytes, &myJSONSchema); err == nil {
		ok = true
		a.JSONSchema = &myJSONSchema
	}
	var mySchemaArray SchemaArray
	if err := json.Unmarshal(bytes, &mySchemaArray); err == nil {
		ok = true
		a.SchemaArray = &mySchemaArray
	}
	if ok {
		return nil
	}
	return errors.New("failed to unmarshal any of the object properties")
}
func (o Items) MarshalJSON() ([]byte, error) {
	out := []interface{}{}
	if o.JSONSchema != nil {
		out = append(out, o.JSONSchema)
	}
	if o.SchemaArray != nil {
		out = append(out, o.SchemaArray)
	}
	return json.Marshal(out)
}
type UniqueItems bool
type StringDoaGddGA string
//
// --- Default ---
//
// []
type StringArray []StringDoaGddGA
//
// --- Default ---
//
// {}
type Definitions map[string]interface{}
//
// --- Default ---
//
// {}
type Properties map[string]interface{}
type PropertyNames interface{}
//
// --- Default ---
//
// {}
type PatternProperties map[string]interface{}
type DependenciesSet struct {
	JSONSchema  *JSONSchema
	StringArray *StringArray
}
func (a *DependenciesSet) UnmarshalJSON(bytes []byte) error {
	var ok bool
	var myJSONSchema JSONSchema
	if err := json.Unmarshal(bytes, &myJSONSchema); err == nil {
		ok = true
		a.JSONSchema = &myJSONSchema
	}
	var myStringArray StringArray
	if err := json.Unmarshal(bytes, &myStringArray); err == nil {
		ok = true
		a.StringArray = &myStringArray
	}
	if ok {
		return nil
	}
	return errors.New("failed to unmarshal any of the object properties")
}
func (o DependenciesSet) MarshalJSON() ([]byte, error) {
	out := []interface{}{}
	if o.JSONSchema != nil {
		out = append(out, o.JSONSchema)
	}
	if o.StringArray != nil {
		out = append(out, o.StringArray)
	}
	return json.Marshal(out)
}
type Dependencies map[string]interface{}
type Enum []AlwaysTrue
type SimpleTypes string
const (
	SimpleTypesEnum0 SimpleTypes = "array"
	SimpleTypesEnum1 SimpleTypes = "boolean"
	SimpleTypesEnum2 SimpleTypes = "integer"
	SimpleTypesEnum3 SimpleTypes = "null"
	SimpleTypesEnum4 SimpleTypes = "number"
	SimpleTypesEnum5 SimpleTypes = "object"
	SimpleTypesEnum6 SimpleTypes = "string"
)
type ArrayOfSimpleTypes []SimpleTypes
type Type struct {
	SimpleTypes        *SimpleTypes
	ArrayOfSimpleTypes *ArrayOfSimpleTypes
}
func (a *Type) UnmarshalJSON(bytes []byte) error {
	var ok bool
	var mySimpleTypes SimpleTypes
	if err := json.Unmarshal(bytes, &mySimpleTypes); err == nil {
		ok = true
		a.SimpleTypes = &mySimpleTypes
	}
	var myArrayOfSimpleTypes ArrayOfSimpleTypes
	if err := json.Unmarshal(bytes, &myArrayOfSimpleTypes); err == nil {
		ok = true
		a.ArrayOfSimpleTypes = &myArrayOfSimpleTypes
	}
	if ok {
		return nil
	}
	return errors.New("failed to unmarshal any of the object properties")
}
func (o Type) MarshalJSON() ([]byte, error) {
	out := []interface{}{}
	if o.SimpleTypes != nil {
		out = append(out, o.SimpleTypes)
	}
	if o.ArrayOfSimpleTypes != nil {
		out = append(out, o.ArrayOfSimpleTypes)
	}
	return json.Marshal(out)
}
type Format string
type ContentMediaType string
type ContentEncoding string
type JSONSchemaObject struct {
	Id                   *Id                            `json:"$id,omitempty"`
	Schema               *Schema                        `json:"$schema,omitempty"`
	Ref                  *Ref                           `json:"$ref,omitempty"`
	Comment              *Comment                       `json:"$comment,omitempty"`
	Title                *Title                         `json:"title,omitempty"`
	Description          *Description                   `json:"description,omitempty"`
	Default              *AlwaysTrue                    `json:"default,omitempty"`
	ReadOnly             *ReadOnly                      `json:"readOnly,omitempty"`
	Examples             *Examples                      `json:"examples,omitempty"`
	MultipleOf           *MultipleOf                    `json:"multipleOf,omitempty"`
	Maximum              *Maximum                       `json:"maximum,omitempty"`
	ExclusiveMaximum     *ExclusiveMaximum              `json:"exclusiveMaximum,omitempty"`
	Minimum              *Minimum                       `json:"minimum,omitempty"`
	ExclusiveMinimum     *ExclusiveMinimum              `json:"exclusiveMinimum,omitempty"`
	MaxLength            *NonNegativeInteger            `json:"maxLength,omitempty"`
	MinLength            *NonNegativeIntegerDefaultZero `json:"minLength,omitempty"`
	Pattern              *Pattern                       `json:"pattern,omitempty"`
	AdditionalItems      *JSONSchema                    `json:"additionalItems,omitempty"`
	Items                *Items                         `json:"items,omitempty"`
	MaxItems             *NonNegativeInteger            `json:"maxItems,omitempty"`
	MinItems             *NonNegativeIntegerDefaultZero `json:"minItems,omitempty"`
	UniqueItems          *UniqueItems                   `json:"uniqueItems,omitempty"`
	Contains             *JSONSchema                    `json:"contains,omitempty"`
	MaxProperties        *NonNegativeInteger            `json:"maxProperties,omitempty"`
	MinProperties        *NonNegativeIntegerDefaultZero `json:"minProperties,omitempty"`
	Required             *StringArray                   `json:"required,omitempty"`
	AdditionalProperties *JSONSchema                    `json:"additionalProperties,omitempty"`
	Definitions          *Definitions                   `json:"definitions,omitempty"`
	Properties           *Properties                    `json:"properties,omitempty"`
	PatternProperties    *PatternProperties             `json:"patternProperties,omitempty"`
	Dependencies         *Dependencies                  `json:"dependencies,omitempty"`
	PropertyNames        *JSONSchema                    `json:"propertyNames,omitempty"`
	Const                *AlwaysTrue                    `json:"const,omitempty"`
	Enum                 *Enum                          `json:"enum,omitempty"`
	Type                 *Type                          `json:"type,omitempty"`
	Format               *Format                        `json:"format,omitempty"`
	ContentMediaType     *ContentMediaType              `json:"contentMediaType,omitempty"`
	ContentEncoding      *ContentEncoding               `json:"contentEncoding,omitempty"`
	If                   *JSONSchema                    `json:"if,omitempty"`
	Then                 *JSONSchema                    `json:"then,omitempty"`
	Else                 *JSONSchema                    `json:"else,omitempty"`
	AllOf                *SchemaArray                   `json:"allOf,omitempty"`
	AnyOf                *SchemaArray                   `json:"anyOf,omitempty"`
	OneOf                *SchemaArray                   `json:"oneOf,omitempty"`
	Not                  *JSONSchema                    `json:"not,omitempty"`
}
// Schema that describes the content.
//
// --- Default ---
//
// {}
type ContentDescriptorObjectSchema struct {
	JSONSchemaObject  *JSONSchemaObject
	JSONSchemaBoolean *JSONSchemaBoolean
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ContentDescriptorObjectSchema) UnmarshalJSON(bytes []byte) error {
	var myJSONSchemaObject JSONSchemaObject
	if err := json.Unmarshal(bytes, &myJSONSchemaObject); err == nil {
		o.JSONSchemaObject = &myJSONSchemaObject
		return nil
	}
	var myJSONSchemaBoolean JSONSchemaBoolean
	if err := json.Unmarshal(bytes, &myJSONSchemaBoolean); err == nil {
		o.JSONSchemaBoolean = &myJSONSchemaBoolean
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ContentDescriptorObjectSchema) MarshalJSON() ([]byte, error) {
	if o.JSONSchemaObject != nil {
		return json.Marshal(o.JSONSchemaObject)
	}
	if o.JSONSchemaBoolean != nil {
		return json.Marshal(o.JSONSchemaBoolean)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// Determines if the content is a required field. Default value is `false`.
type ContentDescriptorObjectRequired bool
// Specifies that the content is deprecated and SHOULD be transitioned out of usage. Default value is `false`.
type ContentDescriptorObjectDeprecated bool
// Content Descriptors are objects that do just as they suggest - describe content. They are reusable ways of describing either parameters or result. They MUST have a schema.
type ContentDescriptorObject struct {
	Name        *ContentDescriptorObjectName        `json:"name"`
	Description *ContentDescriptorObjectDescription `json:"description,omitempty"`
	Summary     *ContentDescriptorObjectSummary     `json:"summary,omitempty"`
	Schema      *ContentDescriptorObjectSchema      `json:"schema"`
	Required    *ContentDescriptorObjectRequired    `json:"required,omitempty"`
	Deprecated  *ContentDescriptorObjectDeprecated  `json:"deprecated,omitempty"`
}
type ContentDescriptorOrReference struct {
	ContentDescriptorObject *ContentDescriptorObject
	ReferenceObject         *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ContentDescriptorOrReference) UnmarshalJSON(bytes []byte) error {
	var myContentDescriptorObject ContentDescriptorObject
	if err := json.Unmarshal(bytes, &myContentDescriptorObject); err == nil {
		o.ContentDescriptorObject = &myContentDescriptorObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ContentDescriptorOrReference) MarshalJSON() ([]byte, error) {
	if o.ContentDescriptorObject != nil {
		return json.Marshal(o.ContentDescriptorObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
//  A list of parameters that are applicable for this method. The list MUST NOT include duplicated parameters and therefore require [name](#content-descriptor-name) to be unique. The list can use the [Reference Object](#reference-object) to link to parameters that are defined by the [Content Descriptor Object](#content-descriptor-object). All optional params (content descriptor objects with "required": false) MUST be positioned after all required params in the list.
type MethodObjectParams []ContentDescriptorOrReference
// The description of the result returned by the method. If defined, it MUST be a Content Descriptor or Reference Object. If undefined, the method MUST only be used as a [notification](https://www.jsonrpc.org/specification#notification)
type MethodObjectResult struct {
	ContentDescriptorObject *ContentDescriptorObject
	ReferenceObject         *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *MethodObjectResult) UnmarshalJSON(bytes []byte) error {
	var myContentDescriptorObject ContentDescriptorObject
	if err := json.Unmarshal(bytes, &myContentDescriptorObject); err == nil {
		o.ContentDescriptorObject = &myContentDescriptorObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o MethodObjectResult) MarshalJSON() ([]byte, error) {
	if o.ContentDescriptorObject != nil {
		return json.Marshal(o.ContentDescriptorObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.
type ErrorObjectCode int64
// A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.
type ErrorObjectMessage string
// A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).
type ErrorObjectData interface{}
// Defines an application level error.
type ErrorObject struct {
	Code    *ErrorObjectCode    `json:"code"`
	Message *ErrorObjectMessage `json:"message"`
	Data    *ErrorObjectData    `json:"data,omitempty"`
}
type ErrorOrReference struct {
	ErrorObject     *ErrorObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ErrorOrReference) UnmarshalJSON(bytes []byte) error {
	var myErrorObject ErrorObject
	if err := json.Unmarshal(bytes, &myErrorObject); err == nil {
		o.ErrorObject = &myErrorObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ErrorOrReference) MarshalJSON() ([]byte, error) {
	if o.ErrorObject != nil {
		return json.Marshal(o.ErrorObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// A list of custom application defined errors that MAY be returned. The Errors MUST have unique error codes.
type MethodObjectErrors []ErrorOrReference
// Cannonical name of the link.
type LinkObjectName interface{}
// Short description for the link.
type LinkObjectSummary string
// The name of an existing, resolvable OpenRPC method, as defined with a unique `method`. This field MUST resolve to a unique [Method Object](#method-object). As opposed to Open Api, Relative `method` values ARE NOT permitted.
type LinkObjectMethod string
// A description of the link. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
type LinkObjectDescription string
// A map representing parameters to pass to a method as specified with `method`. The key is the parameter name to be used, whereas the value can be a constant or a [runtime expression](#runtime-expression) to be evaluated and passed to the linked method.
type LinkObjectParams interface{}
// A server object to be used by the target method.
type LinkObjectServer struct {
	Url         *ServerObjectUrl         `json:"url"`
	Name        *ServerObjectName        `json:"name,omitempty"`
	Description *ServerObjectDescription `json:"description,omitempty"`
	Summary     *ServerObjectSummary     `json:"summary,omitempty"`
	Variables   *ServerObjectVariables   `json:"variables,omitempty"`
}
// A object representing a Link
type LinkObject interface{}
type LinkOrReference struct {
	LinkObject      *LinkObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *LinkOrReference) UnmarshalJSON(bytes []byte) error {
	var myLinkObject LinkObject
	if err := json.Unmarshal(bytes, &myLinkObject); err == nil {
		o.LinkObject = &myLinkObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o LinkOrReference) MarshalJSON() ([]byte, error) {
	if o.LinkObject != nil {
		return json.Marshal(o.LinkObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// A list of possible links from this method call.
type MethodObjectLinks []LinkOrReference
// Name for the example pairing.
type ExamplePairingObjectName string
// A verbose explanation of the example pairing.
type ExamplePairingObjectDescription string
// Short description for the example.
type ExampleObjectSummary string
// Embedded literal example. The `value` field and `externalValue` field are mutually exclusive. To represent examples of media types that cannot naturally represented in JSON, use a string value to contain the example, escaping where necessary.
type ExampleObjectValue interface{}
// A verbose explanation of the example. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
type ExampleObjectDescription string
// Cannonical name of the example.
type ExampleObjectName string
// The Example object is an object that defines an example that is intended to match the `schema` of a given [Content Descriptor](#content-descriptor-object).
type ExampleObject struct {
	Summary     *ExampleObjectSummary     `json:"summary,omitempty"`
	Value       *ExampleObjectValue       `json:"value"`
	Description *ExampleObjectDescription `json:"description,omitempty"`
	Name        *ExampleObjectName        `json:"name"`
}
type ExampleOrReference struct {
	ExampleObject   *ExampleObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ExampleOrReference) UnmarshalJSON(bytes []byte) error {
	var myExampleObject ExampleObject
	if err := json.Unmarshal(bytes, &myExampleObject); err == nil {
		o.ExampleObject = &myExampleObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ExampleOrReference) MarshalJSON() ([]byte, error) {
	if o.ExampleObject != nil {
		return json.Marshal(o.ExampleObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// Example parameters.
type ExamplePairingObjectParams []ExampleOrReference
// Example result. When not provided, the example pairing represents usage of the method as a notification.
type ExamplePairingObjectResult struct {
	ExampleObject   *ExampleObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ExamplePairingObjectResult) UnmarshalJSON(bytes []byte) error {
	var myExampleObject ExampleObject
	if err := json.Unmarshal(bytes, &myExampleObject); err == nil {
		o.ExampleObject = &myExampleObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ExamplePairingObjectResult) MarshalJSON() ([]byte, error) {
	if o.ExampleObject != nil {
		return json.Marshal(o.ExampleObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// The Example Pairing object consists of a set of example params and result. The result is what you can expect from the JSON-RPC service given the exact params.
type ExamplePairingObject struct {
	Name        *ExamplePairingObjectName        `json:"name"`
	Description *ExamplePairingObjectDescription `json:"description,omitempty"`
	Params      *ExamplePairingObjectParams      `json:"params"`
	Result      *ExamplePairingObjectResult      `json:"result,omitempty"`
}
type ExamplePairingOrReference struct {
	ExamplePairingObject *ExamplePairingObject
	ReferenceObject      *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ExamplePairingOrReference) UnmarshalJSON(bytes []byte) error {
	var myExamplePairingObject ExamplePairingObject
	if err := json.Unmarshal(bytes, &myExamplePairingObject); err == nil {
		o.ExamplePairingObject = &myExamplePairingObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ExamplePairingOrReference) MarshalJSON() ([]byte, error) {
	if o.ExamplePairingObject != nil {
		return json.Marshal(o.ExamplePairingObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// Array of [Example Pairing Objects](#example-pairing-object) where each example includes a valid params-to-result [Content Descriptor](#content-descriptor-object) pairing.
type MethodObjectExamples []ExamplePairingOrReference
// Declares this method to be deprecated. Consumers SHOULD refrain from usage of the declared method. Default value is `false`.
type MethodObjectDeprecated bool
// Describes the interface for the given method name. The method name is used as the `method` field of the JSON-RPC body. It therefore MUST be unique.
type MethodObject struct {
	Name           *MethodObjectName            `json:"name"`
	Description    *MethodObjectDescription     `json:"description,omitempty"`
	Summary        *MethodObjectSummary         `json:"summary,omitempty"`
	Servers        *Servers                     `json:"servers,omitempty"`
	Tags           *MethodObjectTags            `json:"tags,omitempty"`
	ParamStructure *MethodObjectParamStructure  `json:"paramStructure,omitempty"`
	Params         *MethodObjectParams          `json:"params"`
	Result         *MethodObjectResult          `json:"result,omitempty"`
	Errors         *MethodObjectErrors          `json:"errors,omitempty"`
	Links          *MethodObjectLinks           `json:"links,omitempty"`
	Examples       *MethodObjectExamples        `json:"examples,omitempty"`
	Deprecated     *MethodObjectDeprecated      `json:"deprecated,omitempty"`
	ExternalDocs   *ExternalDocumentationObject `json:"externalDocs,omitempty"`
}
type MethodOrReference struct {
	MethodObject    *MethodObject
	ReferenceObject *ReferenceObject
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *MethodOrReference) UnmarshalJSON(bytes []byte) error {
	var myMethodObject MethodObject
	if err := json.Unmarshal(bytes, &myMethodObject); err == nil {
		o.MethodObject = &myMethodObject
		return nil
	}
	var myReferenceObject ReferenceObject
	if err := json.Unmarshal(bytes, &myReferenceObject); err == nil {
		o.ReferenceObject = &myReferenceObject
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o MethodOrReference) MarshalJSON() ([]byte, error) {
	if o.MethodObject != nil {
		return json.Marshal(o.MethodObject)
	}
	if o.ReferenceObject != nil {
		return json.Marshal(o.ReferenceObject)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// The available methods for the API. While it is required, the array may be empty (to handle security filtering, for example).
type Methods []MethodOrReference
// An object to hold reusable [Schema Objects](#schema-object).
type SchemaComponents map[string]interface{}
// An object to hold reusable [Link Objects](#link-object).
type LinkComponents map[string]interface{}
// An object to hold reusable [Error Objects](#error-object).
type ErrorComponents map[string]interface{}
// An object to hold reusable [Example Objects](#example-object).
type ExampleComponents map[string]interface{}
// An object to hold reusable [Example Pairing Objects](#example-pairing-object).
type ExamplePairingComponents map[string]interface{}
// An object to hold reusable [Content Descriptor Objects](#content-descriptor-object).
type ContentDescriptorComponents map[string]interface{}
// An object to hold reusable [Tag Objects](#tag-object).
type TagComponents map[string]interface{}
// Holds a set of reusable objects for different aspects of the OpenRPC. All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.
type Components struct {
	Schemas            *SchemaComponents            `json:"schemas,omitempty"`
	Links              *LinkComponents              `json:"links,omitempty"`
	Errors             *ErrorComponents             `json:"errors,omitempty"`
	Examples           *ExampleComponents           `json:"examples,omitempty"`
	ExamplePairings    *ExamplePairingComponents    `json:"examplePairings,omitempty"`
	ContentDescriptors *ContentDescriptorComponents `json:"contentDescriptors,omitempty"`
	Tags               *TagComponents               `json:"tags,omitempty"`
}
// JSON Schema URI (used by some editors)
//
// --- Default ---
//
// https://meta.open-rpc.org/
type MetaSchema string
type OpenrpcDocument struct {
	Openrpc      *Openrpc                     `json:"openrpc"`
	Info         *InfoObject                  `json:"info"`
	ExternalDocs *ExternalDocumentationObject `json:"externalDocs,omitempty"`
	Servers      *Servers                     `json:"servers,omitempty"`
	Methods      *Methods                     `json:"methods"`
	Components   *Components                  `json:"components,omitempty"`
	Schema       *MetaSchema                  `json:"$schema,omitempty"`
}

const RawOpenrpcDocument = "{\"$schema\":\"https://meta.json-schema.tools/\",\"$id\":\"https://meta.open-rpc.org/\",\"title\":\"openrpcDocument\",\"type\":\"object\",\"required\":[\"info\",\"methods\",\"openrpc\"],\"additionalProperties\":false,\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}},\"properties\":{\"openrpc\":{\"description\":\"This string MUST be the [semantic version number](https://semver.org/spec/v2.0.0.html) of the [OpenRPC Specification version](#versions) that the OpenRPC document uses. The `openrpc` field SHOULD be used by tooling specifications and clients to interpret the OpenRPC document. This is *not* related to the API [`info.version`](#info-version) string.\",\"title\":\"openrpc\",\"type\":\"string\",\"regex\":\"^1\\\\.4\\\\.\\\\d+$\"},\"info\":{\"$ref\":\"#/definitions/infoObject\"},\"externalDocs\":{\"$ref\":\"#/definitions/externalDocumentationObject\"},\"servers\":{\"description\":\"An array of Server Objects, which provide connectivity information to a target server. If the `servers` property is not provided, or is an empty array, the default value would be a [Server Object](#server-object) with a [url](#server-url) value of `localhost`. \",\"title\":\"servers\",\"type\":\"array\",\"additionalItems\":false,\"items\":{\"$ref\":\"#/definitions/serverObject\"}},\"methods\":{\"title\":\"methods\",\"type\":\"array\",\"description\":\"The available methods for the API. While it is required, the array may be empty (to handle security filtering, for example).\",\"additionalItems\":false,\"items\":{\"title\":\"methodOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/methodObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"components\":{\"title\":\"components\",\"description\":\"Holds a set of reusable objects for different aspects of the OpenRPC. All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.\",\"type\":\"object\",\"properties\":{\"schemas\":{\"title\":\"schemaComponents\",\"description\":\"An object to hold reusable [Schema Objects](#schema-object).\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/JSONSchema\"}}},\"links\":{\"title\":\"linkComponents\",\"type\":\"object\",\"description\":\"An object to hold reusable [Link Objects](#link-object).\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/linkObject\"}}},\"errors\":{\"title\":\"errorComponents\",\"description\":\"An object to hold reusable [Error Objects](#error-object).\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/errorObject\"}}},\"examples\":{\"title\":\"exampleComponents\",\"description\":\"An object to hold reusable [Example Objects](#example-object).\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/exampleObject\"}}},\"examplePairings\":{\"title\":\"examplePairingComponents\",\"description\":\"An object to hold reusable [Example Pairing Objects](#example-pairing-object).\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/examplePairingObject\"}}},\"contentDescriptors\":{\"title\":\"contentDescriptorComponents\",\"description\":\"An object to hold reusable [Content Descriptor Objects](#content-descriptor-object).\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/contentDescriptorObject\"}}},\"tags\":{\"title\":\"tagComponents\",\"description\":\"An object to hold reusable [Tag Objects](#tag-object).\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"$ref\":\"#/definitions/tagObject\"}}}}},\"$schema\":{\"title\":\"metaSchema\",\"description\":\"JSON Schema URI (used by some editors)\",\"type\":\"string\",\"default\":\"https://meta.open-rpc.org/\"}},\"definitions\":{\"specificationExtension\":{\"title\":\"specificationExtension\",\"description\":\"This object MAY be extended with [Specification Extensions](#specification-extensions).\"},\"JSONSchema\":{\"$ref\":\"https://meta.json-schema.tools\"},\"referenceObject\":{\"title\":\"referenceObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"$ref\"],\"properties\":{\"$ref\":{\"description\":\"The reference string.\",\"$ref\":\"https://meta.json-schema.tools/#/definitions/JSONSchemaObject/properties/$ref\"}}},\"errorObject\":{\"title\":\"errorObject\",\"type\":\"object\",\"description\":\"Defines an application level error.\",\"additionalProperties\":false,\"required\":[\"code\",\"message\"],\"properties\":{\"code\":{\"title\":\"errorObjectCode\",\"description\":\"A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.\",\"type\":\"integer\"},\"message\":{\"title\":\"errorObjectMessage\",\"description\":\"A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.\",\"type\":\"string\"},\"data\":{\"title\":\"errorObjectData\",\"description\":\"A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).\"}}},\"licenseObject\":{\"title\":\"licenseObject\",\"description\":\"License information for the exposed API.\",\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"licenseObjectName\",\"description\":\"The license name used for the API.\",\"type\":\"string\"},\"url\":{\"title\":\"licenseObjectUrl\",\"description\":\"A URL to the license used for the API. MUST be in the format of a URL.\",\"type\":\"string\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"contactObject\":{\"description\":\"Contact information for the exposed API.\",\"title\":\"contactObject\",\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"contactObjectName\",\"description\":\"The identifying name of the contact person/organization.\",\"type\":\"string\"},\"email\":{\"title\":\"contactObjectEmail\",\"description\":\"The email address of the contact person/organization. MUST be in the format of an email address.\",\"type\":\"string\"},\"url\":{\"title\":\"contactObjectUrl\",\"description\":\"The URL pointing to the contact information. MUST be in the format of a URL.\",\"type\":\"string\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"infoObject\":{\"title\":\"infoObject\",\"description\":\"The object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.\",\"additionalProperties\":false,\"required\":[\"title\",\"version\"],\"properties\":{\"title\":{\"title\":\"infoObjectTitle\",\"description\":\"The title of the application.\",\"type\":\"string\"},\"description\":{\"title\":\"infoObjectDescription\",\"description\":\"A verbose description of the application. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.\",\"type\":\"string\"},\"termsOfService\":{\"title\":\"infoObjectTermsOfService\",\"description\":\"A URL to the Terms of Service for the API. MUST be in the format of a URL.\",\"type\":\"string\",\"format\":\"uri\"},\"version\":{\"title\":\"infoObjectVersion\",\"description\":\"The version of the OpenRPC document (which is distinct from the [OpenRPC Specification version](#openrpc-version) or the API implementation version).\",\"type\":\"string\"},\"contact\":{\"$ref\":\"#/definitions/contactObject\"},\"license\":{\"$ref\":\"#/definitions/licenseObject\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"serverObject\":{\"title\":\"serverObject\",\"description\":\"A object representing a Server\",\"type\":\"object\",\"required\":[\"url\"],\"additionalProperties\":false,\"properties\":{\"url\":{\"title\":\"serverObjectUrl\",\"description\":\"A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenRPC document is being served. [Server Variables](#server-variables) are passed into the [Runtime Expression](#runtime-expression) to produce a server URL.\",\"type\":\"string\",\"format\":\"uri\"},\"name\":{\"title\":\"serverObjectName\",\"description\":\"An optional string describing the name of the server. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.\",\"type\":\"string\"},\"description\":{\"title\":\"serverObjectDescription\",\"description\":\"An optional string describing the host designated by the URL. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.\",\"type\":\"string\"},\"summary\":{\"title\":\"serverObjectSummary\",\"description\":\"A short summary of what the server is.\",\"type\":\"string\"},\"variables\":{\"title\":\"serverObjectVariables\",\"description\":\"A map between a variable name and its value. The value is passed into the [Runtime Expression](#runtime-expression) to produce a server URL.\",\"type\":\"object\",\"patternProperties\":{\"[0-z]+\":{\"title\":\"serverObjectVariable\",\"description\":\"An object representing a Server Variable for server URL template substitution.\",\"type\":\"object\",\"required\":[\"default\"],\"properties\":{\"default\":{\"title\":\"serverObjectVariableDefault\",\"description\":\"The default value to use for substitution, which SHALL be sent if an alternate value is _not_ supplied. Note this behavior is different than the [Schema Object's](#schema-object) treatment of default values, because in those cases parameter values are optional.\",\"type\":\"string\"},\"description\":{\"title\":\"serverObjectVariableDescription\",\"description\":\"An optional description for the server variable. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.\",\"type\":\"string\"},\"enum\":{\"title\":\"serverObjectVariableEnum\",\"description\":\"An enumeration of string values to be used if the substitution options are from a limited set.\",\"type\":\"array\",\"items\":{\"title\":\"serverObjectVariableEnumItem\",\"description\":\"An enumeration of string values to be used if the substitution options are from a limited set.\",\"type\":\"string\"}}}}}}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"linkObject\":{\"title\":\"linkObject\",\"description\":\"A object representing a Link\",\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"linkObjectName\",\"description\":\"Cannonical name of the link.\",\"minLength\":1},\"summary\":{\"title\":\"linkObjectSummary\",\"description\":\"Short description for the link.\",\"type\":\"string\"},\"method\":{\"title\":\"linkObjectMethod\",\"description\":\"The name of an existing, resolvable OpenRPC method, as defined with a unique `method`. This field MUST resolve to a unique [Method Object](#method-object). As opposed to Open Api, Relative `method` values ARE NOT permitted.\",\"type\":\"string\"},\"description\":{\"title\":\"linkObjectDescription\",\"description\":\"A description of the link. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.\",\"type\":\"string\"},\"params\":{\"title\":\"linkObjectParams\",\"description\":\"A map representing parameters to pass to a method as specified with `method`. The key is the parameter name to be used, whereas the value can be a constant or a [runtime expression](#runtime-expression) to be evaluated and passed to the linked method.\"},\"server\":{\"title\":\"linkObjectServer\",\"description\":\"A server object to be used by the target method.\",\"$ref\":\"#/definitions/serverObject\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"externalDocumentationObject\":{\"description\":\"Additional external documentation.\",\"title\":\"externalDocumentationObject\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"url\"],\"properties\":{\"description\":{\"description\":\"A verbose explanation of the documentation. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.\",\"title\":\"externalDocumentationObjectDescription\",\"type\":\"string\"},\"url\":{\"description\":\"The URL for the target documentation. Value MUST be in the format of a URL.\",\"title\":\"externalDocumentationObjectUrl\",\"type\":\"string\",\"format\":\"uri\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"methodObject\":{\"title\":\"methodObject\",\"description\":\"Describes the interface for the given method name. The method name is used as the `method` field of the JSON-RPC body. It therefore MUST be unique.\",\"type\":\"object\",\"required\":[\"name\",\"params\"],\"additionalProperties\":false,\"properties\":{\"name\":{\"title\":\"methodObjectName\",\"description\":\"The cannonical name for the method. The name MUST be unique within the methods array.\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"methodObjectDescription\",\"description\":\"A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.\",\"type\":\"string\"},\"summary\":{\"title\":\"methodObjectSummary\",\"description\":\"A short summary of what the method does.\",\"type\":\"string\"},\"servers\":{\"title\":\"servers\",\"type\":\"array\",\"description\":\"An array of Server Objects, which provide connectivity information to a target server. If the `servers` property is not provided, or is an empty array, the default value would be a [Server Object](#server-object) with a [url](#server-url) value of `localhost`. \",\"additionalItems\":false,\"items\":{\"$ref\":\"#/definitions/serverObject\"}},\"tags\":{\"title\":\"methodObjectTags\",\"description\":\"A list of tags for API documentation control. Tags can be used for logical grouping of methods by resources or any other qualifier.\",\"type\":\"array\",\"items\":{\"title\":\"tagOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/tagObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"paramStructure\":{\"title\":\"methodObjectParamStructure\",\"type\":\"string\",\"description\":\"Format the server expects the params. Defaults to 'either'.\",\"enum\":[\"by-position\",\"by-name\",\"either\"],\"default\":\"either\"},\"params\":{\"title\":\"methodObjectParams\",\"description\":\" A list of parameters that are applicable for this method. The list MUST NOT include duplicated parameters and therefore require [name](#content-descriptor-name) to be unique. The list can use the [Reference Object](#reference-object) to link to parameters that are defined by the [Content Descriptor Object](#content-descriptor-object). All optional params (content descriptor objects with \\\"required\\\": false) MUST be positioned after all required params in the list.\",\"type\":\"array\",\"items\":{\"title\":\"contentDescriptorOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/contentDescriptorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"result\":{\"title\":\"methodObjectResult\",\"description\":\"The description of the result returned by the method. If defined, it MUST be a Content Descriptor or Reference Object. If undefined, the method MUST only be used as a [notification](https://www.jsonrpc.org/specification#notification)\",\"oneOf\":[{\"$ref\":\"#/definitions/contentDescriptorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]},\"errors\":{\"title\":\"methodObjectErrors\",\"description\":\"A list of custom application defined errors that MAY be returned. The Errors MUST have unique error codes.\",\"type\":\"array\",\"items\":{\"title\":\"errorOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/errorObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"links\":{\"title\":\"methodObjectLinks\",\"description\":\"A list of possible links from this method call.\",\"type\":\"array\",\"items\":{\"title\":\"linkOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/linkObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"examples\":{\"title\":\"methodObjectExamples\",\"description\":\"Array of [Example Pairing Objects](#example-pairing-object) where each example includes a valid params-to-result [Content Descriptor](#content-descriptor-object) pairing.\",\"type\":\"array\",\"items\":{\"title\":\"examplePairingOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/examplePairingObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"deprecated\":{\"title\":\"methodObjectDeprecated\",\"description\":\"Declares this method to be deprecated. Consumers SHOULD refrain from usage of the declared method. Default value is `false`.\",\"type\":\"boolean\",\"default\":false},\"externalDocs\":{\"$ref\":\"#/definitions/externalDocumentationObject\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"tagObject\":{\"title\":\"tagObject\",\"description\":\"Adds metadata to a single tag that is used by the [Method Object](#method-object). It is not mandatory to have a Tag Object per tag defined in the Method Object instances.\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"name\"],\"properties\":{\"name\":{\"title\":\"tagObjectName\",\"description\":\"The name of the tag.\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"tagObjectDescription\",\"description\":\"A verbose explanation for the tag. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.\",\"type\":\"string\"},\"externalDocs\":{\"description\":\"Additional external documentation for this tag.\",\"$ref\":\"#/definitions/externalDocumentationObject\"}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"exampleObject\":{\"title\":\"exampleObject\",\"description\":\"The Example object is an object that defines an example that is intended to match the `schema` of a given [Content Descriptor](#content-descriptor-object).\",\"type\":\"object\",\"required\":[\"name\",\"value\"],\"properties\":{\"summary\":{\"title\":\"exampleObjectSummary\",\"description\":\"Short description for the example.\",\"type\":\"string\"},\"value\":{\"title\":\"exampleObjectValue\",\"description\":\"Embedded literal example. The `value` field and `externalValue` field are mutually exclusive. To represent examples of media types that cannot naturally represented in JSON, use a string value to contain the example, escaping where necessary.\"},\"description\":{\"title\":\"exampleObjectDescription\",\"description\":\"A verbose explanation of the example. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.\",\"type\":\"string\"},\"name\":{\"title\":\"exampleObjectName\",\"description\":\"Cannonical name of the example.\",\"type\":\"string\",\"minLength\":1}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}},\"examplePairingObject\":{\"title\":\"examplePairingObject\",\"description\":\"The Example Pairing object consists of a set of example params and result. The result is what you can expect from the JSON-RPC service given the exact params.\",\"type\":\"object\",\"required\":[\"name\",\"params\"],\"properties\":{\"name\":{\"title\":\"examplePairingObjectName\",\"description\":\"Name for the example pairing.\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"examplePairingObjectDescription\",\"description\":\"A verbose explanation of the example pairing.\",\"type\":\"string\"},\"params\":{\"title\":\"examplePairingObjectParams\",\"description\":\"Example parameters.\",\"type\":\"array\",\"items\":{\"title\":\"exampleOrReference\",\"oneOf\":[{\"$ref\":\"#/definitions/exampleObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}},\"result\":{\"title\":\"examplePairingObjectResult\",\"description\":\"Example result. When not provided, the example pairing represents usage of the method as a notification.\",\"oneOf\":[{\"$ref\":\"#/definitions/exampleObject\"},{\"$ref\":\"#/definitions/referenceObject\"}]}}},\"contentDescriptorObject\":{\"title\":\"contentDescriptorObject\",\"description\":\"Content Descriptors are objects that do just as they suggest - describe content. They are reusable ways of describing either parameters or result. They MUST have a schema.\",\"type\":\"object\",\"additionalProperties\":false,\"required\":[\"name\",\"schema\"],\"properties\":{\"name\":{\"title\":\"contentDescriptorObjectName\",\"description\":\"Name of the content that is being described. If the content described is a method parameter assignable [`by-name`](#method-param-structure), this field SHALL define the parameter's key (ie name).\",\"type\":\"string\",\"minLength\":1},\"description\":{\"title\":\"contentDescriptorObjectDescription\",\"description\":\"A verbose explanation of the content descriptor behavior. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.\",\"type\":\"string\"},\"summary\":{\"title\":\"contentDescriptorObjectSummary\",\"description\":\"A short summary of the content that is being described.\",\"type\":\"string\"},\"schema\":{\"title\":\"contentDescriptorObjectSchema\",\"description\":\"Schema that describes the content.\",\"$ref\":\"#/definitions/JSONSchema\"},\"required\":{\"title\":\"contentDescriptorObjectRequired\",\"description\":\"Determines if the content is a required field. Default value is `false`.\",\"type\":\"boolean\",\"default\":false},\"deprecated\":{\"title\":\"contentDescriptorObjectDeprecated\",\"description\":\"Specifies that the content is deprecated and SHOULD be transitioned out of usage. Default value is `false`.\",\"type\":\"boolean\",\"default\":false}},\"patternProperties\":{\"^x-\":{\"$ref\":\"#/definitions/specificationExtension\"}}}}}"
