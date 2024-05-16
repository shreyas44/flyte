// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/core/security.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Secret_MountType int32

const (
	// Default case, indicates the client can tolerate either mounting options.
	Secret_ANY Secret_MountType = 0
	// ENV_VAR indicates the secret needs to be mounted as an environment variable.
	Secret_ENV_VAR Secret_MountType = 1
	// FILE indicates the secret needs to be mounted as a file.
	Secret_FILE Secret_MountType = 2
)

// Enum value maps for Secret_MountType.
var (
	Secret_MountType_name = map[int32]string{
		0: "ANY",
		1: "ENV_VAR",
		2: "FILE",
	}
	Secret_MountType_value = map[string]int32{
		"ANY":     0,
		"ENV_VAR": 1,
		"FILE":    2,
	}
)

func (x Secret_MountType) Enum() *Secret_MountType {
	p := new(Secret_MountType)
	*p = x
	return p
}

func (x Secret_MountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Secret_MountType) Descriptor() protoreflect.EnumDescriptor {
	return file_flyteidl_core_security_proto_enumTypes[0].Descriptor()
}

func (Secret_MountType) Type() protoreflect.EnumType {
	return &file_flyteidl_core_security_proto_enumTypes[0]
}

func (x Secret_MountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Secret_MountType.Descriptor instead.
func (Secret_MountType) EnumDescriptor() ([]byte, []int) {
	return file_flyteidl_core_security_proto_rawDescGZIP(), []int{0, 0}
}

// Type of the token requested.
type OAuth2TokenRequest_Type int32

const (
	// CLIENT_CREDENTIALS indicates a 2-legged OAuth token requested using client credentials.
	OAuth2TokenRequest_CLIENT_CREDENTIALS OAuth2TokenRequest_Type = 0
)

// Enum value maps for OAuth2TokenRequest_Type.
var (
	OAuth2TokenRequest_Type_name = map[int32]string{
		0: "CLIENT_CREDENTIALS",
	}
	OAuth2TokenRequest_Type_value = map[string]int32{
		"CLIENT_CREDENTIALS": 0,
	}
)

func (x OAuth2TokenRequest_Type) Enum() *OAuth2TokenRequest_Type {
	p := new(OAuth2TokenRequest_Type)
	*p = x
	return p
}

func (x OAuth2TokenRequest_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OAuth2TokenRequest_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_flyteidl_core_security_proto_enumTypes[1].Descriptor()
}

func (OAuth2TokenRequest_Type) Type() protoreflect.EnumType {
	return &file_flyteidl_core_security_proto_enumTypes[1]
}

func (x OAuth2TokenRequest_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OAuth2TokenRequest_Type.Descriptor instead.
func (OAuth2TokenRequest_Type) EnumDescriptor() ([]byte, []int) {
	return file_flyteidl_core_security_proto_rawDescGZIP(), []int{4, 0}
}

// Secret encapsulates information about the secret a task needs to proceed. An environment variable
// FLYTE_SECRETS_ENV_PREFIX will be passed to indicate the prefix of the environment variables that will be present if
// secrets are passed through environment variables.
// FLYTE_SECRETS_DEFAULT_DIR will be passed to indicate the prefix of the path where secrets will be mounted if secrets
// are passed through file mounts.
type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the secret group where to find the key referenced below. For K8s secrets, this should be the name of
	// the v1/secret object. For Confidant, this should be the Credential name. For Vault, this should be the secret name.
	// For AWS Secret Manager, this should be the name of the secret.
	// +required
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// The group version to fetch. This is not supported in all secret management systems. It'll be ignored for the ones
	// that do not support it.
	// +optional
	GroupVersion string `protobuf:"bytes,2,opt,name=group_version,json=groupVersion,proto3" json:"group_version,omitempty"`
	// The name of the secret to mount. This has to match an existing secret in the system. It's up to the implementation
	// of the secret management system to require case sensitivity. For K8s secrets, Confidant and Vault, this should
	// match one of the keys inside the secret. For AWS Secret Manager, it's ignored.
	// +optional
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// mount_requirement is optional. Indicates where the secret has to be mounted. If provided, the execution will fail
	// if the underlying key management system cannot satisfy that requirement. If not provided, the default location
	// will depend on the key management system.
	// +optional
	MountRequirement Secret_MountType `protobuf:"varint,4,opt,name=mount_requirement,json=mountRequirement,proto3,enum=flyteidl.core.Secret_MountType" json:"mount_requirement,omitempty"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_security_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_security_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_security_proto_rawDescGZIP(), []int{0}
}

func (x *Secret) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Secret) GetGroupVersion() string {
	if x != nil {
		return x.GroupVersion
	}
	return ""
}

func (x *Secret) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Secret) GetMountRequirement() Secret_MountType {
	if x != nil {
		return x.MountRequirement
	}
	return Secret_ANY
}

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The credentials to use for the connection, such as API keys, OAuth2 tokens, etc.
	// The key is the name of the secret, and it's defined in the flytekit.
	// flytekit uses the key to locate the desired secret within the map.
	Secrets map[string]string `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The configuration to use for the connection, such as the endpoint, account name, etc.
	// The key is the name of the config, and it's defined in the flytekit.
	Configs map[string]string `protobuf:"bytes,2,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_security_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_security_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_security_proto_rawDescGZIP(), []int{1}
}

func (x *Connection) GetSecrets() map[string]string {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *Connection) GetConfigs() map[string]string {
	if x != nil {
		return x.Configs
	}
	return nil
}

// OAuth2Client encapsulates OAuth2 Client Credentials to be used when making calls on behalf of that task.
type OAuth2Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// client_id is the public id for the client to use. The system will not perform any pre-auth validation that the
	// secret requested matches the client_id indicated here.
	// +required
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// client_secret is a reference to the secret used to authenticate the OAuth2 client.
	// +required
	ClientSecret *Secret `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *OAuth2Client) Reset() {
	*x = OAuth2Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_security_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2Client) ProtoMessage() {}

func (x *OAuth2Client) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_security_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2Client.ProtoReflect.Descriptor instead.
func (*OAuth2Client) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_security_proto_rawDescGZIP(), []int{2}
}

func (x *OAuth2Client) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OAuth2Client) GetClientSecret() *Secret {
	if x != nil {
		return x.ClientSecret
	}
	return nil
}

// Identity encapsulates the various security identities a task can run as. It's up to the underlying plugin to pick the
// right identity for the execution environment.
type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// iam_role references the fully qualified name of Identity & Access Management role to impersonate.
	IamRole string `protobuf:"bytes,1,opt,name=iam_role,json=iamRole,proto3" json:"iam_role,omitempty"`
	// k8s_service_account references a kubernetes service account to impersonate.
	K8SServiceAccount string `protobuf:"bytes,2,opt,name=k8s_service_account,json=k8sServiceAccount,proto3" json:"k8s_service_account,omitempty"`
	// oauth2_client references an oauth2 client. Backend plugins can use this information to impersonate the client when
	// making external calls.
	Oauth2Client *OAuth2Client `protobuf:"bytes,3,opt,name=oauth2_client,json=oauth2Client,proto3" json:"oauth2_client,omitempty"`
	// execution_identity references the subject who makes the execution
	ExecutionIdentity string `protobuf:"bytes,4,opt,name=execution_identity,json=executionIdentity,proto3" json:"execution_identity,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_security_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_security_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_security_proto_rawDescGZIP(), []int{3}
}

func (x *Identity) GetIamRole() string {
	if x != nil {
		return x.IamRole
	}
	return ""
}

func (x *Identity) GetK8SServiceAccount() string {
	if x != nil {
		return x.K8SServiceAccount
	}
	return ""
}

func (x *Identity) GetOauth2Client() *OAuth2Client {
	if x != nil {
		return x.Oauth2Client
	}
	return nil
}

func (x *Identity) GetExecutionIdentity() string {
	if x != nil {
		return x.ExecutionIdentity
	}
	return ""
}

// OAuth2TokenRequest encapsulates information needed to request an OAuth2 token.
// FLYTE_TOKENS_ENV_PREFIX will be passed to indicate the prefix of the environment variables that will be present if
// tokens are passed through environment variables.
// FLYTE_TOKENS_PATH_PREFIX will be passed to indicate the prefix of the path where secrets will be mounted if tokens
// are passed through file mounts.
type OAuth2TokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name indicates a unique id for the token request within this task token requests. It'll be used as a suffix for
	// environment variables and as a filename for mounting tokens as files.
	// +required
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// type indicates the type of the request to make. Defaults to CLIENT_CREDENTIALS.
	// +required
	Type OAuth2TokenRequest_Type `protobuf:"varint,2,opt,name=type,proto3,enum=flyteidl.core.OAuth2TokenRequest_Type" json:"type,omitempty"`
	// client references the client_id/secret to use to request the OAuth2 token.
	// +required
	Client *OAuth2Client `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	// idp_discovery_endpoint references the discovery endpoint used to retrieve token endpoint and other related
	// information.
	// +optional
	IdpDiscoveryEndpoint string `protobuf:"bytes,4,opt,name=idp_discovery_endpoint,json=idpDiscoveryEndpoint,proto3" json:"idp_discovery_endpoint,omitempty"`
	// token_endpoint references the token issuance endpoint. If idp_discovery_endpoint is not provided, this parameter is
	// mandatory.
	// +optional
	TokenEndpoint string `protobuf:"bytes,5,opt,name=token_endpoint,json=tokenEndpoint,proto3" json:"token_endpoint,omitempty"`
}

func (x *OAuth2TokenRequest) Reset() {
	*x = OAuth2TokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_security_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2TokenRequest) ProtoMessage() {}

func (x *OAuth2TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_security_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2TokenRequest.ProtoReflect.Descriptor instead.
func (*OAuth2TokenRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_security_proto_rawDescGZIP(), []int{4}
}

func (x *OAuth2TokenRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OAuth2TokenRequest) GetType() OAuth2TokenRequest_Type {
	if x != nil {
		return x.Type
	}
	return OAuth2TokenRequest_CLIENT_CREDENTIALS
}

func (x *OAuth2TokenRequest) GetClient() *OAuth2Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *OAuth2TokenRequest) GetIdpDiscoveryEndpoint() string {
	if x != nil {
		return x.IdpDiscoveryEndpoint
	}
	return ""
}

func (x *OAuth2TokenRequest) GetTokenEndpoint() string {
	if x != nil {
		return x.TokenEndpoint
	}
	return ""
}

// SecurityContext holds security attributes that apply to tasks.
type SecurityContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// run_as encapsulates the identity a pod should run as. If the task fills in multiple fields here, it'll be up to the
	// backend plugin to choose the appropriate identity for the execution engine the task will run on.
	RunAs *Identity `protobuf:"bytes,1,opt,name=run_as,json=runAs,proto3" json:"run_as,omitempty"`
	// secrets indicate the list of secrets the task needs in order to proceed. Secrets will be mounted/passed to the
	// pod as it starts. If the plugin responsible for kicking of the task will not run it on a flyte cluster (e.g. AWS
	// Batch), it's the responsibility of the plugin to fetch the secret (which means propeller identity will need access
	// to the secret) and to pass it to the remote execution engine.
	Secrets []*Secret `protobuf:"bytes,2,rep,name=secrets,proto3" json:"secrets,omitempty"`
	// tokens indicate the list of token requests the task needs in order to proceed. Tokens will be mounted/passed to the
	// pod as it starts. If the plugin responsible for kicking of the task will not run it on a flyte cluster (e.g. AWS
	// Batch), it's the responsibility of the plugin to fetch the secret (which means propeller identity will need access
	// to the secret) and to pass it to the remote execution engine.
	Tokens []*OAuth2TokenRequest `protobuf:"bytes,3,rep,name=tokens,proto3" json:"tokens,omitempty"`
	// The name of the connection.
	// The connection is defined in the externalResourceAttributes or flyteadmin configmap.
	// The connection config take precedence in the following order:
	// 1. connection in the externalResourceAttributes in the project-domain settings.
	// 2. connection in the externalResourceAttributes in the project settings.
	// 3. connection in the flyteadmin configmap.
	// +optional
	ConnectionRef string `protobuf:"bytes,4,opt,name=connection_ref,json=connectionRef,proto3" json:"connection_ref,omitempty"`
}

func (x *SecurityContext) Reset() {
	*x = SecurityContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_security_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityContext) ProtoMessage() {}

func (x *SecurityContext) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_security_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityContext.ProtoReflect.Descriptor instead.
func (*SecurityContext) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_security_proto_rawDescGZIP(), []int{5}
}

func (x *SecurityContext) GetRunAs() *Identity {
	if x != nil {
		return x.RunAs
	}
	return nil
}

func (x *SecurityContext) GetSecrets() []*Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *SecurityContext) GetTokens() []*OAuth2TokenRequest {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *SecurityContext) GetConnectionRef() string {
	if x != nil {
		return x.ConnectionRef
	}
	return ""
}

var File_flyteidl_core_security_proto protoreflect.FileDescriptor

var file_flyteidl_core_security_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xd0, 0x01,
	0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x23,
	0x0a, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4c, 0x0a, 0x11, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1f, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x10, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x09, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x59, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x56,
	0x5f, 0x56, 0x41, 0x52, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x02,
	0x22, 0x88, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x40, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x67, 0x0a, 0x0c, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x61, 0x6d, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x13,
	0x6b, 0x38, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6b, 0x38, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x0d,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x0c, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2d,
	0x0a, 0x12, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x96, 0x02,
	0x0a, 0x12, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x64, 0x70,
	0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x69, 0x64, 0x70, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x1e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x12, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54,
	0x49, 0x41, 0x4c, 0x53, 0x10, 0x00, 0x22, 0xd4, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x75,
	0x6e, 0x5f, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x41, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x42, 0xb3, 0x01,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x42, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d,
	0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0xa2, 0x02, 0x03, 0x46, 0x43, 0x58, 0xaa, 0x02, 0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0xca, 0x02, 0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0xe2, 0x02, 0x19, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x43,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_core_security_proto_rawDescOnce sync.Once
	file_flyteidl_core_security_proto_rawDescData = file_flyteidl_core_security_proto_rawDesc
)

func file_flyteidl_core_security_proto_rawDescGZIP() []byte {
	file_flyteidl_core_security_proto_rawDescOnce.Do(func() {
		file_flyteidl_core_security_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_core_security_proto_rawDescData)
	})
	return file_flyteidl_core_security_proto_rawDescData
}

var file_flyteidl_core_security_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_flyteidl_core_security_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_flyteidl_core_security_proto_goTypes = []interface{}{
	(Secret_MountType)(0),        // 0: flyteidl.core.Secret.MountType
	(OAuth2TokenRequest_Type)(0), // 1: flyteidl.core.OAuth2TokenRequest.Type
	(*Secret)(nil),               // 2: flyteidl.core.Secret
	(*Connection)(nil),           // 3: flyteidl.core.Connection
	(*OAuth2Client)(nil),         // 4: flyteidl.core.OAuth2Client
	(*Identity)(nil),             // 5: flyteidl.core.Identity
	(*OAuth2TokenRequest)(nil),   // 6: flyteidl.core.OAuth2TokenRequest
	(*SecurityContext)(nil),      // 7: flyteidl.core.SecurityContext
	nil,                          // 8: flyteidl.core.Connection.SecretsEntry
	nil,                          // 9: flyteidl.core.Connection.ConfigsEntry
}
var file_flyteidl_core_security_proto_depIdxs = []int32{
	0,  // 0: flyteidl.core.Secret.mount_requirement:type_name -> flyteidl.core.Secret.MountType
	8,  // 1: flyteidl.core.Connection.secrets:type_name -> flyteidl.core.Connection.SecretsEntry
	9,  // 2: flyteidl.core.Connection.configs:type_name -> flyteidl.core.Connection.ConfigsEntry
	2,  // 3: flyteidl.core.OAuth2Client.client_secret:type_name -> flyteidl.core.Secret
	4,  // 4: flyteidl.core.Identity.oauth2_client:type_name -> flyteidl.core.OAuth2Client
	1,  // 5: flyteidl.core.OAuth2TokenRequest.type:type_name -> flyteidl.core.OAuth2TokenRequest.Type
	4,  // 6: flyteidl.core.OAuth2TokenRequest.client:type_name -> flyteidl.core.OAuth2Client
	5,  // 7: flyteidl.core.SecurityContext.run_as:type_name -> flyteidl.core.Identity
	2,  // 8: flyteidl.core.SecurityContext.secrets:type_name -> flyteidl.core.Secret
	6,  // 9: flyteidl.core.SecurityContext.tokens:type_name -> flyteidl.core.OAuth2TokenRequest
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_flyteidl_core_security_proto_init() }
func file_flyteidl_core_security_proto_init() {
	if File_flyteidl_core_security_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_core_security_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_core_security_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_core_security_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2Client); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_core_security_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_core_security_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2TokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_core_security_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityContext); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_core_security_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_core_security_proto_goTypes,
		DependencyIndexes: file_flyteidl_core_security_proto_depIdxs,
		EnumInfos:         file_flyteidl_core_security_proto_enumTypes,
		MessageInfos:      file_flyteidl_core_security_proto_msgTypes,
	}.Build()
	File_flyteidl_core_security_proto = out.File
	file_flyteidl_core_security_proto_rawDesc = nil
	file_flyteidl_core_security_proto_goTypes = nil
	file_flyteidl_core_security_proto_depIdxs = nil
}
