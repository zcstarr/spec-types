from typing import TypedDict
from typing import Optional
from typing import NewType
from typing import Any
from typing import List
from typing import Mapping
from typing import Union
from enum import Enum
"""This string MUST be the [semantic version number](https://semver.org/spec/v2.0.0.html) of the [OpenRPC Specification version](#versions) that the OpenRPC document uses. The `openrpc` field SHOULD be used by tooling specifications and clients to interpret the OpenRPC document. This is *not* related to the API [`info.version`](#info-version) string.
"""
Openrpc = NewType("Openrpc", str)
"""The title of the application.
"""
InfoObjectTitle = NewType("InfoObjectTitle", str)
"""A verbose description of the application. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
"""
InfoObjectDescription = NewType("InfoObjectDescription", str)
"""A URL to the Terms of Service for the API. MUST be in the format of a URL.
"""
InfoObjectTermsOfService = NewType("InfoObjectTermsOfService", str)
"""The version of the OpenRPC document (which is distinct from the [OpenRPC Specification version](#openrpc-document-openrpc) or the API implementation version).
"""
InfoObjectVersion = NewType("InfoObjectVersion", str)
"""The identifying name of the contact person/organization.
"""
ContactObjectName = NewType("ContactObjectName", str)
"""The email address of the contact person/organization. MUST be in the format of an email address.
"""
ContactObjectEmail = NewType("ContactObjectEmail", str)
"""The URL pointing to the contact information. MUST be in the format of a URL.
"""
ContactObjectUrl = NewType("ContactObjectUrl", str)
"""This object MAY be extended with [Specification Extensions](#specification-extensions).
"""
SpecificationExtension = NewType("SpecificationExtension", Any)
"""Contact information for the exposed API.
"""
class ContactObject(TypedDict):
    name: Optional[ContactObjectName]
    email: Optional[ContactObjectEmail]
    url: Optional[ContactObjectUrl]
"""The license name used for the API.
"""
LicenseObjectName = NewType("LicenseObjectName", str)
"""A URL to the license used for the API. MUST be in the format of a URL.
"""
LicenseObjectUrl = NewType("LicenseObjectUrl", str)
"""License information for the exposed API.
"""
class LicenseObject(TypedDict):
    name: Optional[LicenseObjectName]
    url: Optional[LicenseObjectUrl]
"""The object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.
"""
InfoObject = NewType("InfoObject", Any)
"""A verbose explanation of the documentation. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
"""
ExternalDocumentationObjectDescription = NewType("ExternalDocumentationObjectDescription", str)
"""The URL for the target documentation. Value MUST be in the format of a URL.
"""
ExternalDocumentationObjectUrl = NewType("ExternalDocumentationObjectUrl", str)
"""Additional external documentation for this tag.
"""
class ExternalDocumentationObject(TypedDict):
    description: Optional[ExternalDocumentationObjectDescription]
    url: Optional[ExternalDocumentationObjectUrl]
"""A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenRPC document is being served. [Server Variables](#server-variables) are passed into the [Runtime Expression](#runtime-expression) to produce a server URL.
"""
ServerObjectUrl = NewType("ServerObjectUrl", str)
"""An optional string describing the name of the server. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
"""
ServerObjectName = NewType("ServerObjectName", str)
"""An optional string describing the host designated by the URL. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
"""
ServerObjectDescription = NewType("ServerObjectDescription", str)
"""A short summary of what the server is.
"""
ServerObjectSummary = NewType("ServerObjectSummary", str)
"""The default value to use for substitution, which SHALL be sent if an alternate value is _not_ supplied. Note this behavior is different than the [Schema Object's](#schema-object) treatment of default values, because in those cases parameter values are optional.
"""
ServerObjectVariableDefault = NewType("ServerObjectVariableDefault", str)
"""An optional description for the server variable. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
"""
ServerObjectVariableDescription = NewType("ServerObjectVariableDescription", str)
"""An enumeration of string values to be used if the substitution options are from a limited set.
"""
ServerObjectVariableEnumItem = NewType("ServerObjectVariableEnumItem", str)
"""An enumeration of string values to be used if the substitution options are from a limited set.
"""
ServerObjectVariableEnum = NewType("ServerObjectVariableEnum", List[ServerObjectVariableEnumItem])
"""An object representing a Server Variable for server URL template substitution.
"""
class ServerObjectVariable(TypedDict):
    default: Optional[ServerObjectVariableDefault]
    description: Optional[ServerObjectVariableDescription]
    enum: Optional[ServerObjectVariableEnum]
"""A map between a variable name and its value. The value is passed into the [Runtime Expression](#runtime-expression) to produce a server URL.
"""
ServerObjectVariables = NewType("ServerObjectVariables", Mapping[Any, Any])
"""A object representing a Server
"""
class ServerObject(TypedDict):
    url: Optional[ServerObjectUrl]
    name: Optional[ServerObjectName]
    description: Optional[ServerObjectDescription]
    summary: Optional[ServerObjectSummary]
    variables: Optional[ServerObjectVariables]
AlwaysFalse = NewType("AlwaysFalse", Any)
"""An array of Server Objects, which provide connectivity information to a target server. If the `servers` property is not provided, or is an empty array, the default value would be a [Server Object](#server-object) with a [url](#server-url) value of `localhost`. 
"""
Servers = NewType("Servers", List[ServerObject])
"""The cannonical name for the method. The name MUST be unique within the methods array.
"""
MethodObjectName = NewType("MethodObjectName", str)
"""A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.
"""
MethodObjectDescription = NewType("MethodObjectDescription", str)
"""A short summary of what the method does.
"""
MethodObjectSummary = NewType("MethodObjectSummary", str)
"""The name of the tag.
"""
TagObjectName = NewType("TagObjectName", str)
"""A verbose explanation for the tag. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
"""
TagObjectDescription = NewType("TagObjectDescription", str)
"""Adds metadata to a single tag that is used by the [Method Object](#method-object). It is not mandatory to have a Tag Object per tag defined in the Method Object instances.
"""
class TagObject(TypedDict):
    name: Optional[TagObjectName]
    description: Optional[TagObjectDescription]
    externalDocs: Optional[ExternalDocumentationObject]

Ref = NewType("Ref", str)

class ReferenceObject(TypedDict):
    $ref: undefined

TagOrReference = NewType("TagOrReference", Union[TagObject, ReferenceObject])
"""A list of tags for API documentation control. Tags can be used for logical grouping of methods by resources or any other qualifier.
"""
MethodObjectTags = NewType("MethodObjectTags", List[TagOrReference])
"""Format the server expects the params. Defaults to 'either'.
"""
class MethodObjectParamStructure(Enum):
    ByPosition = 0
    ByName = 1
    Either = 2
"""Name of the content that is being described. If the content described is a method parameter assignable [`by-name`](#method-param-structure), this field SHALL define the parameter's key (ie name).
"""
ContentDescriptorObjectName = NewType("ContentDescriptorObjectName", str)
"""A verbose explanation of the content descriptor behavior. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
"""
ContentDescriptorObjectDescription = NewType("ContentDescriptorObjectDescription", str)
"""A short summary of the content that is being described.
"""
ContentDescriptorObjectSummary = NewType("ContentDescriptorObjectSummary", str)

Id = NewType("Id", str)

Schema = NewType("Schema", str)

Comment = NewType("Comment", str)

Title = NewType("Title", str)

Description = NewType("Description", str)
AlwaysTrue = NewType("AlwaysTrue", Any)

ReadOnly = NewType("ReadOnly", bool)

Examples = NewType("Examples", List[AlwaysTrue])

MultipleOf = NewType("MultipleOf", float)

Maximum = NewType("Maximum", float)

ExclusiveMaximum = NewType("ExclusiveMaximum", float)

Minimum = NewType("Minimum", float)

ExclusiveMinimum = NewType("ExclusiveMinimum", float)

NonNegativeInteger = NewType("NonNegativeInteger", int)

NonNegativeIntegerDefaultZero = NewType("NonNegativeIntegerDefaultZero", int)

Pattern = NewType("Pattern", str)
"""Always valid if true. Never valid if false. Is constant.
"""
JSONSchemaBoolean = NewType("JSONSchemaBoolean", bool)

JSONSchema = NewType("JSONSchema", Union[JSONSchemaObject, JSONSchemaBoolean])

SchemaArray = NewType("SchemaArray", List[JSONSchema])

Items = NewType("Items", Union[JSONSchema, SchemaArray])

UniqueItems = NewType("UniqueItems", bool)

StringDoaGddGA = NewType("StringDoaGddGA", str)

StringArray = NewType("StringArray", List[StringDoaGddGA])

Definitions = NewType("Definitions", Mapping[Any, Any])

Properties = NewType("Properties", Mapping[Any, Any])

PropertyNames = NewType("PropertyNames", Any)

PatternProperties = NewType("PatternProperties", Mapping[Any, Any])

DependenciesSet = NewType("DependenciesSet", Union[JSONSchema, StringArray])

Dependencies = NewType("Dependencies", Mapping[Any, Any])

Enum = NewType("Enum", List[AlwaysTrue])

class SimpleTypes(Enum):
    Array = 0
    Boolean = 1
    Integer = 2
    Null = 3
    Number = 4
    Object = 5
    String = 6

ArrayOfSimpleTypes = NewType("ArrayOfSimpleTypes", List[SimpleTypes])

Type = NewType("Type", Union[SimpleTypes, ArrayOfSimpleTypes])

Format = NewType("Format", str)

ContentMediaType = NewType("ContentMediaType", str)

ContentEncoding = NewType("ContentEncoding", str)

class JSONSchemaObject(TypedDict):
    $id: Optional[Id]
    $schema: Optional[Schema]
    $ref: Optional[Ref]
    $comment: Optional[Comment]
    title: Optional[Title]
    description: Optional[Description]
    default: Optional[AlwaysTrue]
    readOnly: Optional[ReadOnly]
    examples: Optional[Examples]
    multipleOf: Optional[MultipleOf]
    maximum: Optional[Maximum]
    exclusiveMaximum: Optional[ExclusiveMaximum]
    minimum: Optional[Minimum]
    exclusiveMinimum: Optional[ExclusiveMinimum]
    maxLength: Optional[NonNegativeInteger]
    minLength: Optional[NonNegativeIntegerDefaultZero]
    pattern: Optional[Pattern]
    additionalItems: Optional[JSONSchema]
    items: Optional[Items]
    maxItems: Optional[NonNegativeInteger]
    minItems: Optional[NonNegativeIntegerDefaultZero]
    uniqueItems: Optional[UniqueItems]
    contains: Optional[JSONSchema]
    maxProperties: Optional[NonNegativeInteger]
    minProperties: Optional[NonNegativeIntegerDefaultZero]
    required: Optional[StringArray]
    additionalProperties: Optional[JSONSchema]
    definitions: Optional[Definitions]
    properties: Optional[Properties]
    patternProperties: Optional[PatternProperties]
    dependencies: Optional[Dependencies]
    propertyNames: Optional[JSONSchema]
    const: Optional[AlwaysTrue]
    enum: Optional[Enum]
    type: Optional[Type]
    format: Optional[Format]
    contentMediaType: Optional[ContentMediaType]
    contentEncoding: Optional[ContentEncoding]
    if: Optional[JSONSchema]
    then: Optional[JSONSchema]
    else: Optional[JSONSchema]
    allOf: Optional[SchemaArray]
    anyOf: Optional[SchemaArray]
    oneOf: Optional[SchemaArray]
    not: Optional[JSONSchema]
"""Schema that describes the content.
"""
ContentDescriptorObjectSchema = NewType("ContentDescriptorObjectSchema", Union[JSONSchemaObject, JSONSchemaBoolean])
"""Determines if the content is a required field. Default value is `false`.
"""
ContentDescriptorObjectRequired = NewType("ContentDescriptorObjectRequired", bool)
"""Specifies that the content is deprecated and SHOULD be transitioned out of usage. Default value is `false`.
"""
ContentDescriptorObjectDeprecated = NewType("ContentDescriptorObjectDeprecated", bool)
"""Content Descriptors are objects that do just as they suggest - describe content. They are reusable ways of describing either parameters or result. They MUST have a schema.
"""
class ContentDescriptorObject(TypedDict):
    name: Optional[ContentDescriptorObjectName]
    description: Optional[ContentDescriptorObjectDescription]
    summary: Optional[ContentDescriptorObjectSummary]
    schema: Optional[ContentDescriptorObjectSchema]
    required: Optional[ContentDescriptorObjectRequired]
    deprecated: Optional[ContentDescriptorObjectDeprecated]

ContentDescriptorOrReference = NewType("ContentDescriptorOrReference", Union[ContentDescriptorObject, ReferenceObject])
""" A list of parameters that are applicable for this method. The list MUST NOT include duplicated parameters and therefore require [name](#content-descriptor-name) to be unique. The list can use the [Reference Object](#reference-object) to link to parameters that are defined by the [Content Descriptor Object](#content-descriptor-object). All optional params (content descriptor objects with "required": false) MUST be positioned after all required params in the list.
"""
MethodObjectParams = NewType("MethodObjectParams", List[ContentDescriptorOrReference])
"""The description of the result returned by the method. If defined, it MUST be a Content Descriptor or Reference Object. If undefined, the method MUST only be used as a [notification](https://www.jsonrpc.org/specification#notification)
"""
MethodObjectResult = NewType("MethodObjectResult", Union[ContentDescriptorObject, ReferenceObject])
"""A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.
"""
ErrorObjectCode = NewType("ErrorObjectCode", int)
"""A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.
"""
ErrorObjectMessage = NewType("ErrorObjectMessage", str)
"""A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).
"""
ErrorObjectData = NewType("ErrorObjectData", Any)
"""Defines an application level error.
"""
class ErrorObject(TypedDict):
    code: Optional[ErrorObjectCode]
    message: Optional[ErrorObjectMessage]
    data: Optional[ErrorObjectData]

ErrorOrReference = NewType("ErrorOrReference", Union[ErrorObject, ReferenceObject])
"""A list of custom application defined errors that MAY be returned. The Errors MUST have unique error codes.
"""
MethodObjectErrors = NewType("MethodObjectErrors", List[ErrorOrReference])
"""Cannonical name of the link.
"""
LinkObjectName = NewType("LinkObjectName", Any)
"""Short description for the link.
"""
LinkObjectSummary = NewType("LinkObjectSummary", str)
"""The name of an existing, resolvable OpenRPC method, as defined with a unique `method`. This field MUST resolve to a unique [Method Object](#method-object). As opposed to Open Api, Relative `method` values ARE NOT permitted.
"""
LinkObjectMethod = NewType("LinkObjectMethod", str)
"""A description of the link. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
"""
LinkObjectDescription = NewType("LinkObjectDescription", str)
"""A map representing parameters to pass to a method as specified with `method`. The key is the parameter name to be used, whereas the value can be a constant or a [runtime expression](#runtime-expression) to be evaluated and passed to the linked method.
"""
LinkObjectParams = NewType("LinkObjectParams", Any)
"""A server object to be used by the target method.
"""
class LinkObjectServer(TypedDict):
    url: Optional[ServerObjectUrl]
    name: Optional[ServerObjectName]
    description: Optional[ServerObjectDescription]
    summary: Optional[ServerObjectSummary]
    variables: Optional[ServerObjectVariables]
"""A object representing a Link
"""
LinkObject = NewType("LinkObject", Any)

LinkOrReference = NewType("LinkOrReference", Union[LinkObject, ReferenceObject])
"""A list of possible links from this method call.
"""
MethodObjectLinks = NewType("MethodObjectLinks", List[LinkOrReference])
"""Name for the example pairing.
"""
ExamplePairingObjectName = NewType("ExamplePairingObjectName", str)
"""A verbose explanation of the example pairing.
"""
ExamplePairingObjectDescription = NewType("ExamplePairingObjectDescription", str)
"""Short description for the example.
"""
ExampleObjectSummary = NewType("ExampleObjectSummary", str)
"""Embedded literal example. The `value` field and `externalValue` field are mutually exclusive. To represent examples of media types that cannot naturally represented in JSON, use a string value to contain the example, escaping where necessary.
"""
ExampleObjectValue = NewType("ExampleObjectValue", Any)
"""A verbose explanation of the example. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
"""
ExampleObjectDescription = NewType("ExampleObjectDescription", str)
"""Cannonical name of the example.
"""
ExampleObjectName = NewType("ExampleObjectName", str)
"""The Example object is an object that defines an example that is intended to match the `schema` of a given [Content Descriptor](#content-descriptor-object).
"""
class ExampleObject(TypedDict):
    summary: Optional[ExampleObjectSummary]
    value: Optional[ExampleObjectValue]
    description: Optional[ExampleObjectDescription]
    name: Optional[ExampleObjectName]

ExampleOrReference = NewType("ExampleOrReference", Union[ExampleObject, ReferenceObject])
"""Example parameters.
"""
ExamplePairingObjectParams = NewType("ExamplePairingObjectParams", List[ExampleOrReference])
"""Example result. When not provided, the example pairing represents usage of the method as a notification.
"""
ExamplePairingObjectResult = NewType("ExamplePairingObjectResult", Union[ExampleObject, ReferenceObject])
"""The Example Pairing object consists of a set of example params and result. The result is what you can expect from the JSON-RPC service given the exact params.
"""
class ExamplePairingObject(TypedDict):
    name: Optional[ExamplePairingObjectName]
    description: Optional[ExamplePairingObjectDescription]
    params: Optional[ExamplePairingObjectParams]
    result: Optional[ExamplePairingObjectResult]

ExamplePairingOrReference = NewType("ExamplePairingOrReference", Union[ExamplePairingObject, ReferenceObject])
"""Array of [Example Pairing Objects](#example-pairing-object) where each example includes a valid params-to-result [Content Descriptor](#content-descriptor-object) pairing.
"""
MethodObjectExamples = NewType("MethodObjectExamples", List[ExamplePairingOrReference])
"""Declares this method to be deprecated. Consumers SHOULD refrain from usage of the declared method. Default value is `false`.
"""
MethodObjectDeprecated = NewType("MethodObjectDeprecated", bool)
"""Describes the interface for the given method name. The method name is used as the `method` field of the JSON-RPC body. It therefore MUST be unique.
"""
class MethodObject(TypedDict):
    name: Optional[MethodObjectName]
    description: Optional[MethodObjectDescription]
    summary: Optional[MethodObjectSummary]
    servers: Optional[Servers]
    tags: Optional[MethodObjectTags]
    paramStructure: Optional[MethodObjectParamStructure]
    params: Optional[MethodObjectParams]
    result: Optional[MethodObjectResult]
    errors: Optional[MethodObjectErrors]
    links: Optional[MethodObjectLinks]
    examples: Optional[MethodObjectExamples]
    deprecated: Optional[MethodObjectDeprecated]
    externalDocs: Optional[ExternalDocumentationObject]

MethodOrReference = NewType("MethodOrReference", Union[MethodObject, ReferenceObject])
"""The available methods for the API. While it is required, the array may be empty (to handle security filtering, for example).
"""
Methods = NewType("Methods", List[MethodOrReference])
"""An object to hold reusable [Schema Objects](#schema-object).
"""
SchemaComponents = NewType("SchemaComponents", Mapping[Any, Any])
"""An object to hold reusable [Link Objects](#link-object).
"""
LinkComponents = NewType("LinkComponents", Mapping[Any, Any])
"""An object to hold reusable [Error Objects](#error-object).
"""
ErrorComponents = NewType("ErrorComponents", Mapping[Any, Any])
"""An object to hold reusable [Example Objects](#example-object).
"""
ExampleComponents = NewType("ExampleComponents", Mapping[Any, Any])
"""An object to hold reusable [Example Pairing Objects](#example-pairing-object).
"""
ExamplePairingComponents = NewType("ExamplePairingComponents", Mapping[Any, Any])
"""An object to hold reusable [Content Descriptor Objects](#content-descriptor-object).
"""
ContentDescriptorComponents = NewType("ContentDescriptorComponents", Mapping[Any, Any])
"""An object to hold reusable [Tag Objects](#tag-object).
"""
TagComponents = NewType("TagComponents", Mapping[Any, Any])
"""Holds a set of reusable objects for different aspects of the OpenRPC. All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.
"""
class Components(TypedDict):
    schemas: Optional[SchemaComponents]
    links: Optional[LinkComponents]
    errors: Optional[ErrorComponents]
    examples: Optional[ExampleComponents]
    examplePairings: Optional[ExamplePairingComponents]
    contentDescriptors: Optional[ContentDescriptorComponents]
    tags: Optional[TagComponents]
"""JSON Schema URI (used by some editors)
"""
MetaSchema = NewType("MetaSchema", str)

class OpenrpcDocument(TypedDict):
    openrpc: undefined
    info: Optional[InfoObject]
    externalDocs: Optional[ExternalDocumentationObject]
    servers: Optional[Servers]
    methods: undefined
    components: Optional[Components]
    $schema: Optional[MetaSchema]