/**
 *
 * This string MUST be the [semantic version number](https://semver.org/spec/v2.0.0.html) of the [OpenRPC Specification version](#versions) that the OpenRPC document uses. The `openrpc` field SHOULD be used by tooling specifications and clients to interpret the OpenRPC document. This is *not* related to the API [`info.version`](#info-version) string.
 *
 */
export type Openrpc = string;
/**
 *
 * The title of the application.
 *
 */
export type InfoObjectTitle = string;
/**
 *
 * A verbose description of the application. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
 *
 */
export type InfoObjectDescription = string;
/**
 *
 * A URL to the Terms of Service for the API. MUST be in the format of a URL.
 *
 */
export type InfoObjectTermsOfService = string;
/**
 *
 * The version of the OpenRPC document (which is distinct from the [OpenRPC Specification version](#openrpc-version) or the API implementation version).
 *
 */
export type InfoObjectVersion = string;
/**
 *
 * The identifying name of the contact person/organization.
 *
 */
export type ContactObjectName = string;
/**
 *
 * The email address of the contact person/organization. MUST be in the format of an email address.
 *
 */
export type ContactObjectEmail = string;
/**
 *
 * The URL pointing to the contact information. MUST be in the format of a URL.
 *
 */
export type ContactObjectUrl = string;
/**
 *
 * This object MAY be extended with [Specification Extensions](#specification-extensions).
 *
 */
export type SpecificationExtension = any;
/**
 *
 * Contact information for the exposed API.
 *
 */
export interface ContactObject {
    name?: ContactObjectName;
    email?: ContactObjectEmail;
    url?: ContactObjectUrl;
    [regex: string]: SpecificationExtension | any;
}
/**
 *
 * The license name used for the API.
 *
 */
export type LicenseObjectName = string;
/**
 *
 * A URL to the license used for the API. MUST be in the format of a URL.
 *
 */
export type LicenseObjectUrl = string;
/**
 *
 * License information for the exposed API.
 *
 */
export interface LicenseObject {
    name?: LicenseObjectName;
    url?: LicenseObjectUrl;
    [regex: string]: SpecificationExtension | any;
}
/**
 *
 * The object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.
 *
 */
export type InfoObject = any;
/**
 *
 * A verbose explanation of the documentation. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
 *
 */
export type ExternalDocumentationObjectDescription = string;
/**
 *
 * The URL for the target documentation. Value MUST be in the format of a URL.
 *
 */
export type ExternalDocumentationObjectUrl = string;
/**
 *
 * Additional external documentation for this tag.
 *
 */
export interface ExternalDocumentationObject {
    description?: ExternalDocumentationObjectDescription;
    url: ExternalDocumentationObjectUrl;
    [regex: string]: SpecificationExtension | any;
}
/**
 *
 * A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenRPC document is being served. [Server Variables](#server-variables) are passed into the [Runtime Expression](#runtime-expression) to produce a server URL.
 *
 */
export type ServerObjectUrl = string;
/**
 *
 * An optional string describing the name of the server. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
 *
 */
export type ServerObjectName = string;
/**
 *
 * An optional string describing the host designated by the URL. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
 *
 */
export type ServerObjectDescription = string;
/**
 *
 * A short summary of what the server is.
 *
 */
export type ServerObjectSummary = string;
/**
 *
 * The default value to use for substitution, which SHALL be sent if an alternate value is _not_ supplied. Note this behavior is different than the [Schema Object's](#schema-object) treatment of default values, because in those cases parameter values are optional.
 *
 */
export type ServerObjectVariableDefault = string;
/**
 *
 * An optional description for the server variable. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
 *
 */
export type ServerObjectVariableDescription = string;
/**
 *
 * An enumeration of string values to be used if the substitution options are from a limited set.
 *
 */
export type ServerObjectVariableEnumItem = string;
/**
 *
 * An enumeration of string values to be used if the substitution options are from a limited set.
 *
 */
export type ServerObjectVariableEnum = ServerObjectVariableEnumItem[];
/**
 *
 * An object representing a Server Variable for server URL template substitution.
 *
 */
export interface ServerObjectVariable {
    default: ServerObjectVariableDefault;
    description?: ServerObjectVariableDescription;
    enum?: ServerObjectVariableEnum;
    [k: string]: any;
}
/**
 *
 * A map between a variable name and its value. The value is passed into the [Runtime Expression](#runtime-expression) to produce a server URL.
 *
 */
export interface ServerObjectVariables {
    [key: string]: any;
}
/**
 *
 * A object representing a Server
 *
 */
export interface ServerObject {
    url: ServerObjectUrl;
    name?: ServerObjectName;
    description?: ServerObjectDescription;
    summary?: ServerObjectSummary;
    variables?: ServerObjectVariables;
    [regex: string]: SpecificationExtension | any;
}
/**
 *
 * An array of Server Objects, which provide connectivity information to a target server. If the `servers` property is not provided, or is an empty array, the default value would be a [Server Object](#server-object) with a [url](#server-url) value of `localhost`.
 *
 */
export type Servers = ServerObject[];
/**
 *
 * The cannonical name for the method. The name MUST be unique within the methods array.
 *
 */
export type MethodObjectName = string;
/**
 *
 * A verbose explanation of the method behavior. GitHub Flavored Markdown syntax MAY be used for rich text representation.
 *
 */
export type MethodObjectDescription = string;
/**
 *
 * A short summary of what the method does.
 *
 */
export type MethodObjectSummary = string;
/**
 *
 * The name of the tag.
 *
 */
export type TagObjectName = string;
/**
 *
 * A verbose explanation for the tag. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
 *
 */
export type TagObjectDescription = string;
/**
 *
 * Adds metadata to a single tag that is used by the [Method Object](#method-object). It is not mandatory to have a Tag Object per tag defined in the Method Object instances.
 *
 */
export interface TagObject {
    name: TagObjectName;
    description?: TagObjectDescription;
    externalDocs?: ExternalDocumentationObject;
    [regex: string]: SpecificationExtension | any;
}
export type $Ref = string;
export interface ReferenceObject {
    $ref: $Ref;
}
export type TagOrReference = TagObject | ReferenceObject;
/**
 *
 * A list of tags for API documentation control. Tags can be used for logical grouping of methods by resources or any other qualifier.
 *
 */
export type MethodObjectTags = TagOrReference[];
/**
 *
 * Format the server expects the params. Defaults to 'either'.
 *
 * @default either
 *
 */
export type MethodObjectParamStructure = "by-position" | "by-name" | "either";
/**
 *
 * Name of the content that is being described. If the content described is a method parameter assignable [`by-name`](#method-param-structure), this field SHALL define the parameter's key (ie name).
 *
 */
export type ContentDescriptorObjectName = string;
/**
 *
 * A verbose explanation of the content descriptor behavior. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
 *
 */
export type ContentDescriptorObjectDescription = string;
/**
 *
 * A short summary of the content that is being described.
 *
 */
export type ContentDescriptorObjectSummary = string;
export type $Id = string;
export type $Schema = string;
export type $Comment = string;
export type Title = string;
export type Description = string;
type AlwaysTrue = any;
export type ReadOnly = boolean;
export type Examples = AlwaysTrue[];
export type MultipleOf = number;
export type Maximum = number;
export type ExclusiveMaximum = number;
export type Minimum = number;
export type ExclusiveMinimum = number;
export type NonNegativeInteger = number;
export type NonNegativeIntegerDefaultZero = number;
export type Pattern = string;
/**
 *
 * Always valid if true. Never valid if false. Is constant.
 *
 */
export type JSONSchemaBoolean = boolean;
/**
 *
 * @default {}
 *
 */
export type JSONSchema = JSONSchemaObject | JSONSchemaBoolean;
export type SchemaArray = JSONSchema[];
/**
 *
 * @default true
 *
 */
export type Items = JSONSchema | SchemaArray;
export type UniqueItems = boolean;
export type StringDoaGddGA = string;
/**
 *
 * @default []
 *
 */
export type StringArray = StringDoaGddGA[];
/**
 *
 * @default {}
 *
 */
export interface Definitions {
    [key: string]: any;
}
/**
 *
 * @default {}
 *
 */
export interface Properties {
    [key: string]: any;
}
export type PropertyNames = any;
/**
 *
 * @default {}
 *
 */
export interface PatternProperties {
    [key: string]: any;
}
export type DependenciesSet = JSONSchema | StringArray;
export interface Dependencies {
    [key: string]: any;
}
export type Enum = AlwaysTrue[];
export type SimpleTypes = "array" | "boolean" | "integer" | "null" | "number" | "object" | "string";
export type ArrayOfSimpleTypes = SimpleTypes[];
export type Type = SimpleTypes | ArrayOfSimpleTypes;
export type Format = string;
export type ContentMediaType = string;
export type ContentEncoding = string;
export interface JSONSchemaObject {
    $id?: $Id;
    $schema?: $Schema;
    $ref?: $Ref;
    $comment?: $Comment;
    title?: Title;
    description?: Description;
    default?: AlwaysTrue;
    readOnly?: ReadOnly;
    examples?: Examples;
    multipleOf?: MultipleOf;
    maximum?: Maximum;
    exclusiveMaximum?: ExclusiveMaximum;
    minimum?: Minimum;
    exclusiveMinimum?: ExclusiveMinimum;
    maxLength?: NonNegativeInteger;
    minLength?: NonNegativeIntegerDefaultZero;
    pattern?: Pattern;
    additionalItems?: JSONSchema;
    items?: Items;
    maxItems?: NonNegativeInteger;
    minItems?: NonNegativeIntegerDefaultZero;
    uniqueItems?: UniqueItems;
    contains?: JSONSchema;
    maxProperties?: NonNegativeInteger;
    minProperties?: NonNegativeIntegerDefaultZero;
    required?: StringArray;
    additionalProperties?: JSONSchema;
    definitions?: Definitions;
    properties?: Properties;
    patternProperties?: PatternProperties;
    dependencies?: Dependencies;
    propertyNames?: JSONSchema;
    const?: AlwaysTrue;
    enum?: Enum;
    type?: Type;
    format?: Format;
    contentMediaType?: ContentMediaType;
    contentEncoding?: ContentEncoding;
    if?: JSONSchema;
    then?: JSONSchema;
    else?: JSONSchema;
    allOf?: SchemaArray;
    anyOf?: SchemaArray;
    oneOf?: SchemaArray;
    not?: JSONSchema;
    [k: string]: any;
}
/**
 *
 * Schema that describes the content.
 *
 * @default {}
 *
 */
export type ContentDescriptorObjectSchema = JSONSchemaObject | JSONSchemaBoolean;
/**
 *
 * Determines if the content is a required field. Default value is `false`.
 *
 */
export type ContentDescriptorObjectRequired = boolean;
/**
 *
 * Specifies that the content is deprecated and SHOULD be transitioned out of usage. Default value is `false`.
 *
 */
export type ContentDescriptorObjectDeprecated = boolean;
/**
 *
 * Content Descriptors are objects that do just as they suggest - describe content. They are reusable ways of describing either parameters or result. They MUST have a schema.
 *
 */
export interface ContentDescriptorObject {
    name: ContentDescriptorObjectName;
    description?: ContentDescriptorObjectDescription;
    summary?: ContentDescriptorObjectSummary;
    schema: ContentDescriptorObjectSchema;
    required?: ContentDescriptorObjectRequired;
    deprecated?: ContentDescriptorObjectDeprecated;
    [regex: string]: SpecificationExtension | any;
}
export type ContentDescriptorOrReference = ContentDescriptorObject | ReferenceObject;
/**
 *
 *  A list of parameters that are applicable for this method. The list MUST NOT include duplicated parameters and therefore require [name](#content-descriptor-name) to be unique. The list can use the [Reference Object](#reference-object) to link to parameters that are defined by the [Content Descriptor Object](#content-descriptor-object). All optional params (content descriptor objects with "required": false) MUST be positioned after all required params in the list.
 *
 */
export type MethodObjectParams = ContentDescriptorOrReference[];
/**
 *
 * The description of the result returned by the method. If defined, it MUST be a Content Descriptor or Reference Object. If undefined, the method MUST only be used as a [notification](https://www.jsonrpc.org/specification#notification)
 *
 */
export type MethodObjectResult = ContentDescriptorObject | ReferenceObject;
/**
 *
 * A Number that indicates the error type that occurred. This MUST be an integer. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. These pre-defined errors SHOULD be assumed to be returned from any JSON-RPC api.
 *
 */
export type ErrorObjectCode = number;
/**
 *
 * A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.
 *
 */
export type ErrorObjectMessage = string;
/**
 *
 * A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).
 *
 */
export type ErrorObjectData = any;
/**
 *
 * Defines an application level error.
 *
 */
export interface ErrorObject {
    code: ErrorObjectCode;
    message: ErrorObjectMessage;
    data?: ErrorObjectData;
}
export type ErrorOrReference = ErrorObject | ReferenceObject;
/**
 *
 * A list of custom application defined errors that MAY be returned. The Errors MUST have unique error codes.
 *
 */
export type MethodObjectErrors = ErrorOrReference[];
/**
 *
 * Cannonical name of the link.
 *
 */
export type LinkObjectName = any;
/**
 *
 * Short description for the link.
 *
 */
export type LinkObjectSummary = string;
/**
 *
 * The name of an existing, resolvable OpenRPC method, as defined with a unique `method`. This field MUST resolve to a unique [Method Object](#method-object). As opposed to Open Api, Relative `method` values ARE NOT permitted.
 *
 */
export type LinkObjectMethod = string;
/**
 *
 * A description of the link. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
 *
 */
export type LinkObjectDescription = string;
/**
 *
 * A map representing parameters to pass to a method as specified with `method`. The key is the parameter name to be used, whereas the value can be a constant or a [runtime expression](#runtime-expression) to be evaluated and passed to the linked method.
 *
 */
export type LinkObjectParams = any;
/**
 *
 * A server object to be used by the target method.
 *
 */
export interface LinkObjectServer {
    url: ServerObjectUrl;
    name?: ServerObjectName;
    description?: ServerObjectDescription;
    summary?: ServerObjectSummary;
    variables?: ServerObjectVariables;
    [regex: string]: SpecificationExtension | any;
}
/**
 *
 * A object representing a Link
 *
 */
export type LinkObject = any;
export type LinkOrReference = LinkObject | ReferenceObject;
/**
 *
 * A list of possible links from this method call.
 *
 */
export type MethodObjectLinks = LinkOrReference[];
/**
 *
 * Name for the example pairing.
 *
 */
export type ExamplePairingObjectName = string;
/**
 *
 * A verbose explanation of the example pairing.
 *
 */
export type ExamplePairingObjectDescription = string;
/**
 *
 * Short description for the example.
 *
 */
export type ExampleObjectSummary = string;
/**
 *
 * Embedded literal example. The `value` field and `externalValue` field are mutually exclusive. To represent examples of media types that cannot naturally represented in JSON, use a string value to contain the example, escaping where necessary.
 *
 */
export type ExampleObjectValue = any;
/**
 *
 * A verbose explanation of the example. [GitHub Flavored Markdown syntax](https://github.github.com/gfm/) MAY be used for rich text representation.
 *
 */
export type ExampleObjectDescription = string;
/**
 *
 * Cannonical name of the example.
 *
 */
export type ExampleObjectName = string;
/**
 *
 * The Example object is an object that defines an example that is intended to match the `schema` of a given [Content Descriptor](#content-descriptor-object).
 *
 */
export interface ExampleObject {
    summary?: ExampleObjectSummary;
    value: ExampleObjectValue;
    description?: ExampleObjectDescription;
    name: ExampleObjectName;
    [regex: string]: SpecificationExtension | any;
}
export type ExampleOrReference = ExampleObject | ReferenceObject;
/**
 *
 * Example parameters.
 *
 */
export type ExamplePairingObjectParams = ExampleOrReference[];
/**
 *
 * Example result. When not provided, the example pairing represents usage of the method as a notification.
 *
 */
export type ExamplePairingObjectResult = ExampleObject | ReferenceObject;
/**
 *
 * The Example Pairing object consists of a set of example params and result. The result is what you can expect from the JSON-RPC service given the exact params.
 *
 */
export interface ExamplePairingObject {
    name: ExamplePairingObjectName;
    description?: ExamplePairingObjectDescription;
    params: ExamplePairingObjectParams;
    result?: ExamplePairingObjectResult;
    [k: string]: any;
}
export type ExamplePairingOrReference = ExamplePairingObject | ReferenceObject;
/**
 *
 * Array of [Example Pairing Objects](#example-pairing-object) where each example includes a valid params-to-result [Content Descriptor](#content-descriptor-object) pairing.
 *
 */
export type MethodObjectExamples = ExamplePairingOrReference[];
/**
 *
 * Declares this method to be deprecated. Consumers SHOULD refrain from usage of the declared method. Default value is `false`.
 *
 */
export type MethodObjectDeprecated = boolean;
/**
 *
 * Describes the interface for the given method name. The method name is used as the `method` field of the JSON-RPC body. It therefore MUST be unique.
 *
 */
export interface MethodObject {
    name: MethodObjectName;
    description?: MethodObjectDescription;
    summary?: MethodObjectSummary;
    servers?: Servers;
    tags?: MethodObjectTags;
    paramStructure?: MethodObjectParamStructure;
    params: MethodObjectParams;
    result?: MethodObjectResult;
    errors?: MethodObjectErrors;
    links?: MethodObjectLinks;
    examples?: MethodObjectExamples;
    deprecated?: MethodObjectDeprecated;
    externalDocs?: ExternalDocumentationObject;
    [regex: string]: SpecificationExtension | any;
}
export type MethodOrReference = MethodObject | ReferenceObject;
/**
 *
 * The available methods for the API. While it is required, the array may be empty (to handle security filtering, for example).
 *
 */
export type Methods = MethodOrReference[];
/**
 *
 * An object to hold reusable [Schema Objects](#schema-object).
 *
 */
export interface SchemaComponents {
    [key: string]: any;
}
/**
 *
 * An object to hold reusable [Link Objects](#link-object).
 *
 */
export interface LinkComponents {
    [key: string]: any;
}
/**
 *
 * An object to hold reusable [Error Objects](#error-object).
 *
 */
export interface ErrorComponents {
    [key: string]: any;
}
/**
 *
 * An object to hold reusable [Example Objects](#example-object).
 *
 */
export interface ExampleComponents {
    [key: string]: any;
}
/**
 *
 * An object to hold reusable [Example Pairing Objects](#example-pairing-object).
 *
 */
export interface ExamplePairingComponents {
    [key: string]: any;
}
/**
 *
 * An object to hold reusable [Content Descriptor Objects](#content-descriptor-object).
 *
 */
export interface ContentDescriptorComponents {
    [key: string]: any;
}
/**
 *
 * An object to hold reusable [Tag Objects](#tag-object).
 *
 */
export interface TagComponents {
    [key: string]: any;
}
/**
 *
 * Holds a set of reusable objects for different aspects of the OpenRPC. All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.
 *
 */
export interface Components {
    schemas?: SchemaComponents;
    links?: LinkComponents;
    errors?: ErrorComponents;
    examples?: ExampleComponents;
    examplePairings?: ExamplePairingComponents;
    contentDescriptors?: ContentDescriptorComponents;
    tags?: TagComponents;
    [k: string]: any;
}
/**
 *
 * JSON Schema URI (used by some editors)
 *
 * @default https://meta.open-rpc.org/
 *
 */
export type MetaSchema = string;
export interface OpenrpcDocument {
    openrpc: Openrpc;
    info: InfoObject;
    externalDocs?: ExternalDocumentationObject;
    servers?: Servers;
    methods: Methods;
    components?: Components;
    $schema?: MetaSchema;
    [regex: string]: SpecificationExtension | any;
}
export {};
