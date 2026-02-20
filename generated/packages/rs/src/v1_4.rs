extern crate serde;
extern crate serde_json;
extern crate derive_builder;

use serde::{Serialize, Deserialize};
use derive_builder::Builder;
use std::collections::HashMap;
/// Openrpc
///
/// This string MUST be the [semantic version number](https://semver.org/spec/v2.0.0.html) of the [OpenRPC Specification version](#versions) that the OpenRPC document uses. The `openrpc` field SHOULD be used by tooling specifications and clients to interpret the OpenRPC document. This is *not* related to the API [`info.version`](#info-version) string.
///
pub type Openrpc = String;
/// InfoObjectTitle
///
/// The title of the application.
///
pub type InfoObjectTitle = String;
/// InfoObjectDescription
///
/// A verbose description of the application. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
///
pub type InfoObjectDescription = String;
/// InfoObjectTermsOfService
///
/// A URL to the Terms of Service for the API. MUST be in the format of a URL.
///
pub type InfoObjectTermsOfService = String;
/// InfoObjectVersion
///
/// The version of the OpenRPC document (which is distinct from the [OpenRPC Specification version](#openrpc-version) or the API implementation version).
///
pub type InfoObjectVersion = String;
/// ContactObjectName
///
/// The identifying name of the contact person/organization.
///
pub type ContactObjectName = String;
/// ContactObjectEmail
///
/// The email address of the contact person/organization. MUST be in the format of an email address.
///
pub type ContactObjectEmail = String;
/// ContactObjectUrl
///
/// The URL pointing to the contact information. MUST be in the format of a URL.
///
pub type ContactObjectUrl = String;
/// SpecificationExtension
///
/// This object MAY be extended with [Specification Extensions](#specification-extensions).
///
pub type SpecificationExtension = serde_json::Value;
/// ContactObject
///
/// Contact information for the exposed API.
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct ContactObject {
    #[serde(skip_serializing_if = "Option::is_none")]
    pub name: Option<ContactObjectName>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub email: Option<ContactObjectEmail>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub url: Option<ContactObjectUrl>,
}
/// LicenseObjectName
///
/// The license name used for the API.
///
pub type LicenseObjectName = String;
/// LicenseObjectUrl
///
/// A URL to the license used for the API. MUST be in the format of a URL.
///
pub type LicenseObjectUrl = String;
/// LicenseObject
///
/// License information for the exposed API.
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct LicenseObject {
    #[serde(skip_serializing_if = "Option::is_none")]
    pub name: Option<LicenseObjectName>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub url: Option<LicenseObjectUrl>,
}
/// InfoObject
///
/// The object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.
///
pub type InfoObject = serde_json::Value;
/// ExternalDocumentationObjectDescription
///
/// A verbose explanation of the documentation. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
///
pub type ExternalDocumentationObjectDescription = String;
/// ExternalDocumentationObjectUrl
///
/// The URL for the target documentation. Value MUST be in the format of a URL.
///
pub type ExternalDocumentationObjectUrl = String;
/// ExternalDocumentationObject
///
/// Additional external documentation for this tag.
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct ExternalDocumentationObject {
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<ExternalDocumentationObjectDescription>,
    pub url: ExternalDocumentationObjectUrl,
}
/// ServerObjectUrl
///
/// A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenRPC document is being served. [Server Variables](#server-variables) are passed into the [Runtime Expression](#runtime-expression) to produce a server URL.
///
pub type ServerObjectUrl = String;
/// ServerObjectName
///
/// An optional string describing the name of the server. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
///
pub type ServerObjectName = String;
/// ServerObjectDescription
///
/// An optional string describing the host designated by the URL. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
///
pub type ServerObjectDescription = String;
/// ServerObjectSummary
///
/// A short summary of what the server is.
///
pub type ServerObjectSummary = String;
/// ServerObjectVariableDefault
///
/// The default value to use for substitution, which SHALL be sent if an alternate value is _not_ supplied. Note this behavior is different than the [Schema Object's](#schema-object) treatment of default values, because in those cases parameter values are optional.
///
pub type ServerObjectVariableDefault = String;
/// ServerObjectVariableDescription
///
/// An optional description for the server variable. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
///
pub type ServerObjectVariableDescription = String;
/// ServerObjectVariableEnumItem
///
/// An enumeration of string values to be used if the substitution options are from a limited set.
///
pub type ServerObjectVariableEnumItem = String;
/// ServerObjectVariableEnum
///
/// An enumeration of string values to be used if the substitution options are from a limited set.
///
pub type ServerObjectVariableEnum = Vec<ServerObjectVariableEnumItem>;
/// ServerObjectVariable
///
/// An object representing a Server Variable for server URL template substitution.
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct ServerObjectVariable {
    #[serde(rename = "default")]
    pub _default: ServerObjectVariableDefault,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<ServerObjectVariableDescription>,
    #[serde(rename = "enum", skip_serializing_if = "Option::is_none")]
    pub _enum: Option<ServerObjectVariableEnum>,
}
/// ServerObjectVariables
///
/// A map between a variable name and its value. The value is passed into the [Runtime Expression](#runtime-expression) to produce a server URL.
///
pub type ServerObjectVariables = HashMap<String, serde_json::Value>;
/// ServerObject
///
/// A object representing a Server
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct ServerObject {
    pub url: ServerObjectUrl,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub name: Option<ServerObjectName>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<ServerObjectDescription>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub summary: Option<ServerObjectSummary>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub variables: Option<ServerObjectVariables>,
}
type AlwaysFalse = serde_json::Value;
/// Servers
///
/// An array of Server Objects, which provide connectivity information to a target server. If the `servers` property is not provided, or is an empty array, the default value would be a [Server Object](#server-object) with a [url](#server-url) value of `localhost`. 
///
pub type Servers = Vec<ServerObject>;
/// MethodObjectName
///
/// The cannonical name for the method. The name MUST be unique within the methods array.
///
pub type MethodObjectName = String;
/// MethodObjectDescription
///
/// A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.
///
pub type MethodObjectDescription = String;
/// MethodObjectSummary
///
/// A short summary of what the method does.
///
pub type MethodObjectSummary = String;
/// TagObjectName
///
/// The name of the tag.
///
pub type TagObjectName = String;
/// TagObjectDescription
///
/// A verbose explanation for the tag. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
///
pub type TagObjectDescription = String;
/// TagObject
///
/// Adds metadata to a single tag that is used by the [Method Object](#method-object). It is not mandatory to have a Tag Object per tag defined in the Method Object instances.
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct TagObject {
    pub name: TagObjectName,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<TagObjectDescription>,
    #[serde(rename = "externalDocs", skip_serializing_if = "Option::is_none")]
    pub external_docs: Option<ExternalDocumentationObject>,
}
pub type Ref = String;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct ReferenceObject {
    #[serde(rename = "$ref")]
    pub _ref: Ref,
}
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum TagOrReference {
    TagObject(TagObject),
    ReferenceObject(ReferenceObject),
}
/// MethodObjectTags
///
/// A list of tags for API documentation control. Tags can be used for logical grouping of methods by resources or any other qualifier.
///
pub type MethodObjectTags = Vec<TagOrReference>;
/// MethodObjectParamStructure
///
/// Format the server expects the params. Defaults to 'either'.
///
/// # Default
///
/// either
///
#[derive(Serialize, Deserialize, Clone, Debug, Eq, PartialEq)]
pub enum MethodObjectParamStructure {
    #[serde(rename = "by-position")]
    ByPosition,
    #[serde(rename = "by-name")]
    ByName,
    #[serde(rename = "either")]
    Either,
}
/// ContentDescriptorObjectName
///
/// Name of the content that is being described. If the content described is a method parameter assignable [`by-name`](#method-param-structure), this field SHALL define the parameter's key (ie name).
///
pub type ContentDescriptorObjectName = String;
/// ContentDescriptorObjectDescription
///
/// A verbose explanation of the content descriptor behavior. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
///
pub type ContentDescriptorObjectDescription = String;
/// ContentDescriptorObjectSummary
///
/// A short summary of the content that is being described.
///
pub type ContentDescriptorObjectSummary = String;
pub type Id = String;
pub type Schema = String;
pub type Comment = String;
pub type Title = String;
pub type Description = String;
type AlwaysTrue = serde_json::Value;
pub type ReadOnly = bool;
pub type Examples = Vec<AlwaysTrue>;
pub type MultipleOf = f64;
pub type Maximum = f64;
pub type ExclusiveMaximum = f64;
pub type Minimum = f64;
pub type ExclusiveMinimum = f64;
pub type NonNegativeInteger = i64;
pub type NonNegativeIntegerDefaultZero = i64;
pub type Pattern = String;
/// JSONSchemaBoolean
///
/// Always valid if true. Never valid if false. Is constant.
///
pub type JSONSchemaBoolean = bool;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum JSONSchema {
    JSONSchemaObject(Box<JSONSchemaObject>),
    JSONSchemaBoolean(JSONSchemaBoolean),
}
pub type SchemaArray = Vec<JSONSchema>;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum Items {
    JSONSchema(JSONSchema),
    SchemaArray(SchemaArray),
}
pub type UniqueItems = bool;
pub type StringDoaGddGA = String;
/// StringArray
///
/// # Default
///
/// []
///
pub type StringArray = Vec<StringDoaGddGA>;
/// Definitions
///
/// # Default
///
/// {}
///
pub type Definitions = HashMap<String, JSONSchema>;
/// Properties
///
/// # Default
///
/// {}
///
pub type Properties = HashMap<String, JSONSchema>;
pub type PropertyNames = serde_json::Value;
/// PatternProperties
///
/// # Default
///
/// {}
///
pub type PatternProperties = HashMap<String, JSONSchema>;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum DependenciesSet {
    JSONSchema(JSONSchema),
    StringArray(StringArray),
}
pub type Dependencies = HashMap<String, DependenciesSet>;
pub type Enum = Vec<AlwaysTrue>;
#[derive(Serialize, Deserialize, Clone, Debug, Eq, PartialEq)]
pub enum SimpleTypes {
    #[serde(rename = "array")]
    Array,
    #[serde(rename = "boolean")]
    Boolean,
    #[serde(rename = "integer")]
    Integer,
    #[serde(rename = "null")]
    Null,
    #[serde(rename = "number")]
    Number,
    #[serde(rename = "object")]
    Object,
    #[serde(rename = "string")]
    String,
}
pub type ArrayOfSimpleTypes = Vec<SimpleTypes>;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum Type {
    SimpleTypes(SimpleTypes),
    ArrayOfSimpleTypes(ArrayOfSimpleTypes),
}
pub type Format = String;
pub type ContentMediaType = String;
pub type ContentEncoding = String;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct JSONSchemaObject {
    #[serde(rename = "$id", skip_serializing_if = "Option::is_none")]
    pub id: Option<Id>,
    #[serde(rename = "$schema", skip_serializing_if = "Option::is_none")]
    pub schema: Option<Schema>,
    #[serde(rename = "$ref", skip_serializing_if = "Option::is_none")]
    pub _ref: Option<Ref>,
    #[serde(rename = "$comment", skip_serializing_if = "Option::is_none")]
    pub comment: Option<Comment>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub title: Option<Title>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<Description>,
    #[serde(rename = "default", skip_serializing_if = "Option::is_none")]
    pub _default: Option<AlwaysTrue>,
    #[serde(rename = "readOnly", skip_serializing_if = "Option::is_none")]
    pub read_only: Option<ReadOnly>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub examples: Option<Examples>,
    #[serde(rename = "multipleOf", skip_serializing_if = "Option::is_none")]
    pub multiple_of: Option<MultipleOf>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub maximum: Option<Maximum>,
    #[serde(rename = "exclusiveMaximum", skip_serializing_if = "Option::is_none")]
    pub exclusive_maximum: Option<ExclusiveMaximum>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub minimum: Option<Minimum>,
    #[serde(rename = "exclusiveMinimum", skip_serializing_if = "Option::is_none")]
    pub exclusive_minimum: Option<ExclusiveMinimum>,
    #[serde(rename = "maxLength", skip_serializing_if = "Option::is_none")]
    pub max_length: Option<NonNegativeInteger>,
    #[serde(rename = "minLength", skip_serializing_if = "Option::is_none")]
    pub min_length: Option<NonNegativeIntegerDefaultZero>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub pattern: Option<Pattern>,
    #[serde(rename = "additionalItems", skip_serializing_if = "Option::is_none")]
    pub additional_items: Option<JSONSchema>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub items: Option<Items>,
    #[serde(rename = "maxItems", skip_serializing_if = "Option::is_none")]
    pub max_items: Option<NonNegativeInteger>,
    #[serde(rename = "minItems", skip_serializing_if = "Option::is_none")]
    pub min_items: Option<NonNegativeIntegerDefaultZero>,
    #[serde(rename = "uniqueItems", skip_serializing_if = "Option::is_none")]
    pub unique_items: Option<UniqueItems>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub contains: Option<JSONSchema>,
    #[serde(rename = "maxProperties", skip_serializing_if = "Option::is_none")]
    pub max_properties: Option<NonNegativeInteger>,
    #[serde(rename = "minProperties", skip_serializing_if = "Option::is_none")]
    pub min_properties: Option<NonNegativeIntegerDefaultZero>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub required: Option<StringArray>,
    #[serde(rename = "additionalProperties", skip_serializing_if = "Option::is_none")]
    pub additional_properties: Option<JSONSchema>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub definitions: Option<Definitions>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub properties: Option<Properties>,
    #[serde(rename = "patternProperties", skip_serializing_if = "Option::is_none")]
    pub pattern_properties: Option<PatternProperties>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub dependencies: Option<Dependencies>,
    #[serde(rename = "propertyNames", skip_serializing_if = "Option::is_none")]
    pub property_names: Option<JSONSchema>,
    #[serde(rename = "const", skip_serializing_if = "Option::is_none")]
    pub _const: Option<AlwaysTrue>,
    #[serde(rename = "enum", skip_serializing_if = "Option::is_none")]
    pub _enum: Option<Enum>,
    #[serde(rename = "type", skip_serializing_if = "Option::is_none")]
    pub _type: Option<Type>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub format: Option<Format>,
    #[serde(rename = "contentMediaType", skip_serializing_if = "Option::is_none")]
    pub content_media_type: Option<ContentMediaType>,
    #[serde(rename = "contentEncoding", skip_serializing_if = "Option::is_none")]
    pub content_encoding: Option<ContentEncoding>,
    #[serde(rename = "if", skip_serializing_if = "Option::is_none")]
    pub _if: Option<JSONSchema>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub then: Option<JSONSchema>,
    #[serde(rename = "else", skip_serializing_if = "Option::is_none")]
    pub _else: Option<JSONSchema>,
    #[serde(rename = "allOf", skip_serializing_if = "Option::is_none")]
    pub all_of: Option<SchemaArray>,
    #[serde(rename = "anyOf", skip_serializing_if = "Option::is_none")]
    pub any_of: Option<SchemaArray>,
    #[serde(rename = "oneOf", skip_serializing_if = "Option::is_none")]
    pub one_of: Option<SchemaArray>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub not: Option<JSONSchema>,
}
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum ContentDescriptorObjectSchema {
    JSONSchemaObject(Box<JSONSchemaObject>),
    JSONSchemaBoolean(JSONSchemaBoolean),
}
/// ContentDescriptorObjectRequired
///
/// Determines if the content is a required field. Default value is `false`.
///
pub type ContentDescriptorObjectRequired = bool;
/// ContentDescriptorObjectDeprecated
///
/// Specifies that the content is deprecated and SHOULD be transitioned out of usage. Default value is `false`.
///
pub type ContentDescriptorObjectDeprecated = bool;
/// ContentDescriptorObject
///
/// Content Descriptors are objects that do just as they suggest - describe content. They are reusable ways of describing either parameters or result. They MUST have a schema.
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct ContentDescriptorObject {
    pub name: ContentDescriptorObjectName,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<ContentDescriptorObjectDescription>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub summary: Option<ContentDescriptorObjectSummary>,
    pub schema: ContentDescriptorObjectSchema,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub required: Option<ContentDescriptorObjectRequired>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub deprecated: Option<ContentDescriptorObjectDeprecated>,
}
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum ContentDescriptorOrReference {
    ContentDescriptorObject(ContentDescriptorObject),
    ReferenceObject(ReferenceObject),
}
/// MethodObjectParams
///
///  A list of parameters that are applicable for this method. The list MUST NOT include duplicated parameters and therefore require [name](#content-descriptor-name) to be unique. The list can use the [Reference Object](#reference-object) to link to parameters that are defined by the [Content Descriptor Object](#content-descriptor-object). All optional params (content descriptor objects with "required": false) MUST be positioned after all required params in the list.
///
pub type MethodObjectParams = Vec<ContentDescriptorOrReference>;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum MethodObjectResult {
    ContentDescriptorObject(ContentDescriptorObject),
    ReferenceObject(ReferenceObject),
}
/// ErrorObjectCode
///
/// A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.
///
pub type ErrorObjectCode = i64;
/// ErrorObjectMessage
///
/// A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.
///
pub type ErrorObjectMessage = String;
/// ErrorObjectData
///
/// A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).
///
pub type ErrorObjectData = serde_json::Value;
/// ErrorObject
///
/// Defines an application level error.
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct ErrorObject {
    pub code: ErrorObjectCode,
    pub message: ErrorObjectMessage,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub data: Option<ErrorObjectData>,
}
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum ErrorOrReference {
    ErrorObject(ErrorObject),
    ReferenceObject(ReferenceObject),
}
/// MethodObjectErrors
///
/// A list of custom application defined errors that MAY be returned. The Errors MUST have unique error codes.
///
pub type MethodObjectErrors = Vec<ErrorOrReference>;
/// LinkObjectName
///
/// Cannonical name of the link.
///
pub type LinkObjectName = serde_json::Value;
/// LinkObjectSummary
///
/// Short description for the link.
///
pub type LinkObjectSummary = String;
/// LinkObjectMethod
///
/// The name of an existing, resolvable OpenRPC method, as defined with a unique `method`. This field MUST resolve to a unique [Method Object](#method-object). As opposed to Open Api, Relative `method` values ARE NOT permitted.
///
pub type LinkObjectMethod = String;
/// LinkObjectDescription
///
/// A description of the link. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
///
pub type LinkObjectDescription = String;
/// LinkObjectParams
///
/// A map representing parameters to pass to a method as specified with `method`. The key is the parameter name to be used, whereas the value can be a constant or a [runtime expression](#runtime-expression) to be evaluated and passed to the linked method.
///
pub type LinkObjectParams = serde_json::Value;
/// LinkObjectServer
///
/// A server object to be used by the target method.
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct LinkObjectServer {
    pub url: ServerObjectUrl,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub name: Option<ServerObjectName>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<ServerObjectDescription>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub summary: Option<ServerObjectSummary>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub variables: Option<ServerObjectVariables>,
}
/// LinkObject
///
/// A object representing a Link
///
pub type LinkObject = serde_json::Value;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum LinkOrReference {
    LinkObject(LinkObject),
    ReferenceObject(ReferenceObject),
}
/// MethodObjectLinks
///
/// A list of possible links from this method call.
///
pub type MethodObjectLinks = Vec<LinkOrReference>;
/// ExamplePairingObjectName
///
/// Name for the example pairing.
///
pub type ExamplePairingObjectName = String;
/// ExamplePairingObjectDescription
///
/// A verbose explanation of the example pairing.
///
pub type ExamplePairingObjectDescription = String;
/// ExampleObjectSummary
///
/// Short description for the example.
///
pub type ExampleObjectSummary = String;
/// ExampleObjectValue
///
/// Embedded literal example. The `value` field and `externalValue` field are mutually exclusive. To represent examples of media types that cannot naturally represented in JSON, use a string value to contain the example, escaping where necessary.
///
pub type ExampleObjectValue = serde_json::Value;
/// ExampleObjectDescription
///
/// A verbose explanation of the example. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
///
pub type ExampleObjectDescription = String;
/// ExampleObjectName
///
/// Cannonical name of the example.
///
pub type ExampleObjectName = String;
/// ExampleObject
///
/// The Example object is an object that defines an example that is intended to match the `schema` of a given [Content Descriptor](#content-descriptor-object).
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct ExampleObject {
    #[serde(skip_serializing_if = "Option::is_none")]
    pub summary: Option<ExampleObjectSummary>,
    pub value: ExampleObjectValue,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<ExampleObjectDescription>,
    pub name: ExampleObjectName,
}
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum ExampleOrReference {
    ExampleObject(ExampleObject),
    ReferenceObject(ReferenceObject),
}
/// ExamplePairingObjectParams
///
/// Example parameters.
///
pub type ExamplePairingObjectParams = Vec<ExampleOrReference>;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum ExamplePairingObjectResult {
    ExampleObject(ExampleObject),
    ReferenceObject(ReferenceObject),
}
/// ExamplePairingObject
///
/// The Example Pairing object consists of a set of example params and result. The result is what you can expect from the JSON-RPC service given the exact params.
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct ExamplePairingObject {
    pub name: ExamplePairingObjectName,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<ExamplePairingObjectDescription>,
    pub params: ExamplePairingObjectParams,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub result: Option<ExamplePairingObjectResult>,
}
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum ExamplePairingOrReference {
    ExamplePairingObject(ExamplePairingObject),
    ReferenceObject(ReferenceObject),
}
/// MethodObjectExamples
///
/// Array of [Example Pairing Objects](#example-pairing-object) where each example includes a valid params-to-result [Content Descriptor](#content-descriptor-object) pairing.
///
pub type MethodObjectExamples = Vec<ExamplePairingOrReference>;
/// MethodObjectDeprecated
///
/// Declares this method to be deprecated. Consumers SHOULD refrain from usage of the declared method. Default value is `false`.
///
pub type MethodObjectDeprecated = bool;
/// MethodObject
///
/// Describes the interface for the given method name. The method name is used as the `method` field of the JSON-RPC body. It therefore MUST be unique.
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct MethodObject {
    pub name: MethodObjectName,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<MethodObjectDescription>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub summary: Option<MethodObjectSummary>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub servers: Option<Servers>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub tags: Option<MethodObjectTags>,
    #[serde(rename = "paramStructure", skip_serializing_if = "Option::is_none")]
    pub param_structure: Option<MethodObjectParamStructure>,
    pub params: MethodObjectParams,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub result: Option<MethodObjectResult>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub errors: Option<MethodObjectErrors>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub links: Option<MethodObjectLinks>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub examples: Option<MethodObjectExamples>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub deprecated: Option<MethodObjectDeprecated>,
    #[serde(rename = "externalDocs", skip_serializing_if = "Option::is_none")]
    pub external_docs: Option<ExternalDocumentationObject>,
}
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum MethodOrReference {
    MethodObject(MethodObject),
    ReferenceObject(ReferenceObject),
}
/// Methods
///
/// The available methods for the API. While it is required, the array may be empty (to handle security filtering, for example).
///
pub type Methods = Vec<MethodOrReference>;
/// SchemaComponents
///
/// An object to hold reusable [Schema Objects](#schema-object).
///
pub type SchemaComponents = HashMap<String, serde_json::Value>;
/// LinkComponents
///
/// An object to hold reusable [Link Objects](#link-object).
///
pub type LinkComponents = HashMap<String, serde_json::Value>;
/// ErrorComponents
///
/// An object to hold reusable [Error Objects](#error-object).
///
pub type ErrorComponents = HashMap<String, serde_json::Value>;
/// ExampleComponents
///
/// An object to hold reusable [Example Objects](#example-object).
///
pub type ExampleComponents = HashMap<String, serde_json::Value>;
/// ExamplePairingComponents
///
/// An object to hold reusable [Example Pairing Objects](#example-pairing-object).
///
pub type ExamplePairingComponents = HashMap<String, serde_json::Value>;
/// ContentDescriptorComponents
///
/// An object to hold reusable [Content Descriptor Objects](#content-descriptor-object).
///
pub type ContentDescriptorComponents = HashMap<String, serde_json::Value>;
/// TagComponents
///
/// An object to hold reusable [Tag Objects](#tag-object).
///
pub type TagComponents = HashMap<String, serde_json::Value>;
/// Components
///
/// Holds a set of reusable objects for different aspects of the OpenRPC. All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct Components {
    #[serde(skip_serializing_if = "Option::is_none")]
    pub schemas: Option<SchemaComponents>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub links: Option<LinkComponents>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub errors: Option<ErrorComponents>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub examples: Option<ExampleComponents>,
    #[serde(rename = "examplePairings", skip_serializing_if = "Option::is_none")]
    pub example_pairings: Option<ExamplePairingComponents>,
    #[serde(rename = "contentDescriptors", skip_serializing_if = "Option::is_none")]
    pub content_descriptors: Option<ContentDescriptorComponents>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub tags: Option<TagComponents>,
}
/// MetaSchema
///
/// JSON Schema URI (used by some editors)
///
/// # Default
///
/// https://meta.open-rpc.org/
///
pub type MetaSchema = String;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct OpenrpcDocument {
    pub openrpc: Openrpc,
    pub info: InfoObject,
    #[serde(rename = "externalDocs", skip_serializing_if = "Option::is_none")]
    pub external_docs: Option<ExternalDocumentationObject>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub servers: Option<Servers>,
    pub methods: Methods,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub components: Option<Components>,
    #[serde(rename = "$schema", skip_serializing_if = "Option::is_none")]
    pub schema: Option<MetaSchema>,
}