// Code generated by protoc-gen-go.
// source: protocol.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	PingRequest
	PingResponse
	HostAuthRequest
	Error
	HostAuth
	ReplyMetadata
	HostAuthResponse
	HostCertRequest
	HostCertResponse
	UserAuthRequest
	UserAuthResponse
	UserCertRequest
	OauthToken
	UserCertResponse
	HostCA
	UserCA
	PublicTrustedCARequest
	PublicTrustedCAResponse
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PingRequest struct {
	RequestTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=requestTime" json:"requestTime,omitempty"`
	Name        string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetRequestTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *PingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PingResponse struct {
	Metadata *ReplyMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Message  string         `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetMetadata() *ReplyMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HostAuthRequest struct {
	RequestTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=requestTime" json:"requestTime,omitempty"`
	// this will be using AES-GCM cipher
	// and include the ID to be the first 4 bytes after concatenating with _ the following
	// the <cloud_name>:<account_id>
	// region
	// environment
	// this is used to identify the preshared key with the cert server
	// in response, the cert server will return an ID that can be used for subsequent cert request
	AuthInfo []byte `protobuf:"bytes,2,opt,name=authInfo,proto3" json:"authInfo,omitempty"`
}

func (m *HostAuthRequest) Reset()                    { *m = HostAuthRequest{} }
func (m *HostAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*HostAuthRequest) ProtoMessage()               {}
func (*HostAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HostAuthRequest) GetRequestTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *HostAuthRequest) GetAuthInfo() []byte {
	if m != nil {
		return m.AuthInfo
	}
	return nil
}

type Error struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Error) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// this is the protobuf message to decrypt the HostAuthResponse.authResponse bytes to
type HostAuth struct {
	Id     []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Errors []*Error `protobuf:"bytes,2,rep,name=errors" json:"errors,omitempty"`
}

func (m *HostAuth) Reset()                    { *m = HostAuth{} }
func (m *HostAuth) String() string            { return proto.CompactTextString(m) }
func (*HostAuth) ProtoMessage()               {}
func (*HostAuth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HostAuth) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *HostAuth) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ReplyMetadata struct {
	// copies the request time from the client
	RequestTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=requestTime" json:"requestTime,omitempty"`
	// when the server initiated the response
	ResponseTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=responseTime" json:"responseTime,omitempty"`
}

func (m *ReplyMetadata) Reset()                    { *m = ReplyMetadata{} }
func (m *ReplyMetadata) String() string            { return proto.CompactTextString(m) }
func (*ReplyMetadata) ProtoMessage()               {}
func (*ReplyMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReplyMetadata) GetRequestTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *ReplyMetadata) GetResponseTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.ResponseTime
	}
	return nil
}

type HostAuthResponse struct {
	Metadata *ReplyMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// this should be the encrypted response based on the key sent
	// in the authInfo request for the AuthRequest
	// this is expected to contain the encrypted HostAuth message
	AuthResponse []byte `protobuf:"bytes,3,opt,name=authResponse,proto3" json:"authResponse,omitempty"`
}

func (m *HostAuthResponse) Reset()                    { *m = HostAuthResponse{} }
func (m *HostAuthResponse) String() string            { return proto.CompactTextString(m) }
func (*HostAuthResponse) ProtoMessage()               {}
func (*HostAuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *HostAuthResponse) GetMetadata() *ReplyMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *HostAuthResponse) GetAuthResponse() []byte {
	if m != nil {
		return m.AuthResponse
	}
	return nil
}

// this is only sent after the host has already authenticated with the server
// someone reasonably can read the ID from memory if an attacker is already root on the host
// but at that point all bets are off..
// we can keep the id and relevant tokens encrypted on disk if we need to persist it
// but since the AuthRequest should've taken care of identifying the machine
// id can be sent over the TLS connection to the server
type HostCertRequest struct {
	RequestTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=requestTime" json:"requestTime,omitempty"`
	ValidFrom   *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=validFrom" json:"validFrom,omitempty"`
	ValidUntil  *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=validUntil" json:"validUntil,omitempty"`
	Id          []byte                     `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	// these need to be FQDN of all the IPs the host needs to have
	Hostnames []string `protobuf:"bytes,5,rep,name=hostnames" json:"hostnames,omitempty"`
	PublicKey []byte   `protobuf:"bytes,6,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	// Send the HostMetadata after the Authentication step
	HostMetadata []byte `protobuf:"bytes,7,opt,name=hostMetadata,proto3" json:"hostMetadata,omitempty"`
}

func (m *HostCertRequest) Reset()                    { *m = HostCertRequest{} }
func (m *HostCertRequest) String() string            { return proto.CompactTextString(m) }
func (*HostCertRequest) ProtoMessage()               {}
func (*HostCertRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *HostCertRequest) GetRequestTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *HostCertRequest) GetValidFrom() *google_protobuf.Timestamp {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *HostCertRequest) GetValidUntil() *google_protobuf.Timestamp {
	if m != nil {
		return m.ValidUntil
	}
	return nil
}

func (m *HostCertRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *HostCertRequest) GetHostnames() []string {
	if m != nil {
		return m.Hostnames
	}
	return nil
}

func (m *HostCertRequest) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *HostCertRequest) GetHostMetadata() []byte {
	if m != nil {
		return m.HostMetadata
	}
	return nil
}

type HostCertResponse struct {
	Metadata *ReplyMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// this is the cert for the host
	HostCert []byte `protobuf:"bytes,3,opt,name=hostCert,proto3" json:"hostCert,omitempty"`
	// this is the cert the host should "trust" for users
	// logging into the machines
	TrustedUsersCACert []byte `protobuf:"bytes,4,opt,name=trustedUsersCACert,proto3" json:"trustedUsersCACert,omitempty"`
}

func (m *HostCertResponse) Reset()                    { *m = HostCertResponse{} }
func (m *HostCertResponse) String() string            { return proto.CompactTextString(m) }
func (*HostCertResponse) ProtoMessage()               {}
func (*HostCertResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *HostCertResponse) GetMetadata() *ReplyMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *HostCertResponse) GetHostCert() []byte {
	if m != nil {
		return m.HostCert
	}
	return nil
}

func (m *HostCertResponse) GetTrustedUsersCACert() []byte {
	if m != nil {
		return m.TrustedUsersCACert
	}
	return nil
}

type UserAuthRequest struct {
	RequestTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=requestTime" json:"requestTime,omitempty"`
	Username    string                     `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// send the access token
	Token *OauthToken `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
}

func (m *UserAuthRequest) Reset()                    { *m = UserAuthRequest{} }
func (m *UserAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*UserAuthRequest) ProtoMessage()               {}
func (*UserAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UserAuthRequest) GetRequestTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *UserAuthRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserAuthRequest) GetToken() *OauthToken {
	if m != nil {
		return m.Token
	}
	return nil
}

type UserAuthResponse struct {
	Metadata *ReplyMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Username string         `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	UserId   string         `protobuf:"bytes,3,opt,name=userId" json:"userId,omitempty"`
	Valid    bool           `protobuf:"varint,4,opt,name=valid" json:"valid,omitempty"`
	// something cryptographic -- haven't implemented yet
	AuthResponse []byte `protobuf:"bytes,5,opt,name=authResponse,proto3" json:"authResponse,omitempty"`
}

func (m *UserAuthResponse) Reset()                    { *m = UserAuthResponse{} }
func (m *UserAuthResponse) String() string            { return proto.CompactTextString(m) }
func (*UserAuthResponse) ProtoMessage()               {}
func (*UserAuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UserAuthResponse) GetMetadata() *ReplyMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *UserAuthResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserAuthResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserAuthResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *UserAuthResponse) GetAuthResponse() []byte {
	if m != nil {
		return m.AuthResponse
	}
	return nil
}

type UserCertRequest struct {
	RequestTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=requestTime" json:"requestTime,omitempty"`
	UserId      string                     `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
	Username    string                     `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	// this should go in ID of the cert
	RemoteUsername string `protobuf:"bytes,4,opt,name=remoteUsername" json:"remoteUsername,omitempty"`
	// the raw public key to sign with user cert
	PublicKey []byte `protobuf:"bytes,5,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	// this is a hack to just respond back with the
	// new principals added to the current cert
	// also it's useful to log before/after for audit
	// once again, never saved by server, just read
	// and discarded...
	// this will be disabled with a future version
	// once the server has a database
	CurrentUserCert      []byte                     `protobuf:"bytes,6,opt,name=currentUserCert,proto3" json:"currentUserCert,omitempty"`
	ValidFrom            *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=validFrom" json:"validFrom,omitempty"`
	ValidUntil           *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=validUntil" json:"validUntil,omitempty"`
	AuthorizedPrincipals []string                   `protobuf:"bytes,9,rep,name=authorizedPrincipals" json:"authorizedPrincipals,omitempty"`
	// this should be used for scripts to limit access
	ForceCommands []string `protobuf:"bytes,10,rep,name=forceCommands" json:"forceCommands,omitempty"`
}

func (m *UserCertRequest) Reset()                    { *m = UserCertRequest{} }
func (m *UserCertRequest) String() string            { return proto.CompactTextString(m) }
func (*UserCertRequest) ProtoMessage()               {}
func (*UserCertRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UserCertRequest) GetRequestTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *UserCertRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserCertRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserCertRequest) GetRemoteUsername() string {
	if m != nil {
		return m.RemoteUsername
	}
	return ""
}

func (m *UserCertRequest) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *UserCertRequest) GetCurrentUserCert() []byte {
	if m != nil {
		return m.CurrentUserCert
	}
	return nil
}

func (m *UserCertRequest) GetValidFrom() *google_protobuf.Timestamp {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *UserCertRequest) GetValidUntil() *google_protobuf.Timestamp {
	if m != nil {
		return m.ValidUntil
	}
	return nil
}

func (m *UserCertRequest) GetAuthorizedPrincipals() []string {
	if m != nil {
		return m.AuthorizedPrincipals
	}
	return nil
}

func (m *UserCertRequest) GetForceCommands() []string {
	if m != nil {
		return m.ForceCommands
	}
	return nil
}

type OauthToken struct {
	AccessToken  string                     `protobuf:"bytes,1,opt,name=accessToken" json:"accessToken,omitempty"`
	TokenType    string                     `protobuf:"bytes,2,opt,name=tokenType" json:"tokenType,omitempty"`
	RefreshToken string                     `protobuf:"bytes,3,opt,name=refreshToken" json:"refreshToken,omitempty"`
	Expiry       *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=expiry" json:"expiry,omitempty"`
}

func (m *OauthToken) Reset()                    { *m = OauthToken{} }
func (m *OauthToken) String() string            { return proto.CompactTextString(m) }
func (*OauthToken) ProtoMessage()               {}
func (*OauthToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *OauthToken) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *OauthToken) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *OauthToken) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *OauthToken) GetExpiry() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

type UserCertResponse struct {
	Metadata       *ReplyMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Username       string         `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	UserCert       []byte         `protobuf:"bytes,3,opt,name=userCert,proto3" json:"userCert,omitempty"`
	TrustedHostCAs []*HostCA      `protobuf:"bytes,4,rep,name=trustedHostCAs" json:"trustedHostCAs,omitempty"`
}

func (m *UserCertResponse) Reset()                    { *m = UserCertResponse{} }
func (m *UserCertResponse) String() string            { return proto.CompactTextString(m) }
func (*UserCertResponse) ProtoMessage()               {}
func (*UserCertResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UserCertResponse) GetMetadata() *ReplyMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *UserCertResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserCertResponse) GetUserCert() []byte {
	if m != nil {
		return m.UserCert
	}
	return nil
}

func (m *UserCertResponse) GetTrustedHostCAs() []*HostCA {
	if m != nil {
		return m.TrustedHostCAs
	}
	return nil
}

// Public Host Certificate Authority's Public Key
// and additional information
// at any given time more than one public CA might
// be trusted, at most 2, so that servers can be
// rotated to new certificates on time
type HostCA struct {
	ValidFrom  *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=validFrom" json:"validFrom,omitempty"`
	ValidUntil *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=validUntil" json:"validUntil,omitempty"`
	PublicKey  []byte                     `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Id         uint64                     `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
}

func (m *HostCA) Reset()                    { *m = HostCA{} }
func (m *HostCA) String() string            { return proto.CompactTextString(m) }
func (*HostCA) ProtoMessage()               {}
func (*HostCA) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *HostCA) GetValidFrom() *google_protobuf.Timestamp {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *HostCA) GetValidUntil() *google_protobuf.Timestamp {
	if m != nil {
		return m.ValidUntil
	}
	return nil
}

func (m *HostCA) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *HostCA) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserCA struct {
	ValidFrom  *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=validFrom" json:"validFrom,omitempty"`
	ValidUntil *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=validUntil" json:"validUntil,omitempty"`
	PublicKey  []byte                     `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Id         uint64                     `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
}

func (m *UserCA) Reset()                    { *m = UserCA{} }
func (m *UserCA) String() string            { return proto.CompactTextString(m) }
func (*UserCA) ProtoMessage()               {}
func (*UserCA) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *UserCA) GetValidFrom() *google_protobuf.Timestamp {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *UserCA) GetValidUntil() *google_protobuf.Timestamp {
	if m != nil {
		return m.ValidUntil
	}
	return nil
}

func (m *UserCA) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *UserCA) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// I may add more parameters for logging in future
type PublicTrustedCARequest struct {
	RequestTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=requestTime" json:"requestTime,omitempty"`
}

func (m *PublicTrustedCARequest) Reset()                    { *m = PublicTrustedCARequest{} }
func (m *PublicTrustedCARequest) String() string            { return proto.CompactTextString(m) }
func (*PublicTrustedCARequest) ProtoMessage()               {}
func (*PublicTrustedCARequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *PublicTrustedCARequest) GetRequestTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

type PublicTrustedCAResponse struct {
	Metadata *ReplyMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	HostCAs  []*HostCA      `protobuf:"bytes,2,rep,name=hostCAs" json:"hostCAs,omitempty"`
	UserCAs  []*UserCA      `protobuf:"bytes,3,rep,name=UserCAs" json:"UserCAs,omitempty"`
	// the clients should update their corresponding
	// revoked CA file with the contents
	RevokedHostCAs []*HostCA `protobuf:"bytes,4,rep,name=revokedHostCAs" json:"revokedHostCAs,omitempty"`
	RevokedUserCAs []*UserCA `protobuf:"bytes,5,rep,name=revokedUserCAs" json:"revokedUserCAs,omitempty"`
}

func (m *PublicTrustedCAResponse) Reset()                    { *m = PublicTrustedCAResponse{} }
func (m *PublicTrustedCAResponse) String() string            { return proto.CompactTextString(m) }
func (*PublicTrustedCAResponse) ProtoMessage()               {}
func (*PublicTrustedCAResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *PublicTrustedCAResponse) GetMetadata() *ReplyMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PublicTrustedCAResponse) GetHostCAs() []*HostCA {
	if m != nil {
		return m.HostCAs
	}
	return nil
}

func (m *PublicTrustedCAResponse) GetUserCAs() []*UserCA {
	if m != nil {
		return m.UserCAs
	}
	return nil
}

func (m *PublicTrustedCAResponse) GetRevokedHostCAs() []*HostCA {
	if m != nil {
		return m.RevokedHostCAs
	}
	return nil
}

func (m *PublicTrustedCAResponse) GetRevokedUserCAs() []*UserCA {
	if m != nil {
		return m.RevokedUserCAs
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "protocol.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "protocol.PingResponse")
	proto.RegisterType((*HostAuthRequest)(nil), "protocol.HostAuthRequest")
	proto.RegisterType((*Error)(nil), "protocol.Error")
	proto.RegisterType((*HostAuth)(nil), "protocol.HostAuth")
	proto.RegisterType((*ReplyMetadata)(nil), "protocol.ReplyMetadata")
	proto.RegisterType((*HostAuthResponse)(nil), "protocol.HostAuthResponse")
	proto.RegisterType((*HostCertRequest)(nil), "protocol.HostCertRequest")
	proto.RegisterType((*HostCertResponse)(nil), "protocol.HostCertResponse")
	proto.RegisterType((*UserAuthRequest)(nil), "protocol.UserAuthRequest")
	proto.RegisterType((*UserAuthResponse)(nil), "protocol.UserAuthResponse")
	proto.RegisterType((*UserCertRequest)(nil), "protocol.UserCertRequest")
	proto.RegisterType((*OauthToken)(nil), "protocol.OauthToken")
	proto.RegisterType((*UserCertResponse)(nil), "protocol.UserCertResponse")
	proto.RegisterType((*HostCA)(nil), "protocol.HostCA")
	proto.RegisterType((*UserCA)(nil), "protocol.UserCA")
	proto.RegisterType((*PublicTrustedCARequest)(nil), "protocol.PublicTrustedCARequest")
	proto.RegisterType((*PublicTrustedCAResponse)(nil), "protocol.PublicTrustedCAResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cert service

type CertClient interface {
	HostAuth(ctx context.Context, in *HostAuthRequest, opts ...grpc.CallOption) (*HostAuthResponse, error)
	HostCert(ctx context.Context, in *HostCertRequest, opts ...grpc.CallOption) (*HostCertResponse, error)
	UserAuth(ctx context.Context, in *UserAuthRequest, opts ...grpc.CallOption) (*UserAuthResponse, error)
	UserCert(ctx context.Context, in *UserCertRequest, opts ...grpc.CallOption) (*UserCertResponse, error)
	// This responds back with both host CAs that the users should trust
	// and the user CA the servers should trust
	PublicTrustedCA(ctx context.Context, in *PublicTrustedCARequest, opts ...grpc.CallOption) (*PublicTrustedCAResponse, error)
	// this is just for test/sanity
	// We may report he metric to get a sense of how the latency between
	// environments is faring
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type certClient struct {
	cc *grpc.ClientConn
}

func NewCertClient(cc *grpc.ClientConn) CertClient {
	return &certClient{cc}
}

func (c *certClient) HostAuth(ctx context.Context, in *HostAuthRequest, opts ...grpc.CallOption) (*HostAuthResponse, error) {
	out := new(HostAuthResponse)
	err := grpc.Invoke(ctx, "/protocol.Cert/HostAuth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certClient) HostCert(ctx context.Context, in *HostCertRequest, opts ...grpc.CallOption) (*HostCertResponse, error) {
	out := new(HostCertResponse)
	err := grpc.Invoke(ctx, "/protocol.Cert/HostCert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certClient) UserAuth(ctx context.Context, in *UserAuthRequest, opts ...grpc.CallOption) (*UserAuthResponse, error) {
	out := new(UserAuthResponse)
	err := grpc.Invoke(ctx, "/protocol.Cert/UserAuth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certClient) UserCert(ctx context.Context, in *UserCertRequest, opts ...grpc.CallOption) (*UserCertResponse, error) {
	out := new(UserCertResponse)
	err := grpc.Invoke(ctx, "/protocol.Cert/UserCert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certClient) PublicTrustedCA(ctx context.Context, in *PublicTrustedCARequest, opts ...grpc.CallOption) (*PublicTrustedCAResponse, error) {
	out := new(PublicTrustedCAResponse)
	err := grpc.Invoke(ctx, "/protocol.Cert/PublicTrustedCA", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/protocol.Cert/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cert service

type CertServer interface {
	HostAuth(context.Context, *HostAuthRequest) (*HostAuthResponse, error)
	HostCert(context.Context, *HostCertRequest) (*HostCertResponse, error)
	UserAuth(context.Context, *UserAuthRequest) (*UserAuthResponse, error)
	UserCert(context.Context, *UserCertRequest) (*UserCertResponse, error)
	// This responds back with both host CAs that the users should trust
	// and the user CA the servers should trust
	PublicTrustedCA(context.Context, *PublicTrustedCARequest) (*PublicTrustedCAResponse, error)
	// this is just for test/sanity
	// We may report he metric to get a sense of how the latency between
	// environments is faring
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterCertServer(s *grpc.Server, srv CertServer) {
	s.RegisterService(&_Cert_serviceDesc, srv)
}

func _Cert_HostAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertServer).HostAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Cert/HostAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertServer).HostAuth(ctx, req.(*HostAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cert_HostCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertServer).HostCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Cert/HostCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertServer).HostCert(ctx, req.(*HostCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cert_UserAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertServer).UserAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Cert/UserAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertServer).UserAuth(ctx, req.(*UserAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cert_UserCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertServer).UserCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Cert/UserCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertServer).UserCert(ctx, req.(*UserCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cert_PublicTrustedCA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicTrustedCARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertServer).PublicTrustedCA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Cert/PublicTrustedCA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertServer).PublicTrustedCA(ctx, req.(*PublicTrustedCARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cert_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Cert/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cert_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Cert",
	HandlerType: (*CertServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HostAuth",
			Handler:    _Cert_HostAuth_Handler,
		},
		{
			MethodName: "HostCert",
			Handler:    _Cert_HostCert_Handler,
		},
		{
			MethodName: "UserAuth",
			Handler:    _Cert_UserAuth_Handler,
		},
		{
			MethodName: "UserCert",
			Handler:    _Cert_UserCert_Handler,
		},
		{
			MethodName: "PublicTrustedCA",
			Handler:    _Cert_PublicTrustedCA_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Cert_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol.proto",
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 958 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0x4b, 0x8f, 0xe3, 0x44,
	0x10, 0xc6, 0x76, 0x9e, 0x95, 0xec, 0x64, 0xd4, 0x1a, 0x66, 0x4d, 0x84, 0x44, 0xb0, 0x78, 0x44,
	0x2b, 0x91, 0x95, 0xb2, 0x07, 0x10, 0x42, 0x48, 0xd9, 0x08, 0xc4, 0x0a, 0x21, 0x46, 0x56, 0x06,
	0x71, 0x41, 0xc8, 0x63, 0xf7, 0x24, 0xd6, 0xc4, 0xee, 0xd0, 0xdd, 0x5e, 0x31, 0xfc, 0x07, 0x4e,
	0x1c, 0x90, 0x38, 0x73, 0x84, 0xe3, 0x1e, 0x38, 0xf3, 0xc7, 0x50, 0x3f, 0x1c, 0xb7, 0x7b, 0x93,
	0x61, 0x77, 0x93, 0xcb, 0xde, 0xdc, 0x55, 0x5f, 0x7f, 0x5d, 0x5d, 0xf5, 0x55, 0x57, 0x02, 0x27,
	0x1b, 0x4a, 0x38, 0x89, 0xc9, 0x7a, 0x22, 0x3f, 0x50, 0xa7, 0x5c, 0x0f, 0xdf, 0x59, 0x12, 0xb2,
	0x5c, 0xe3, 0x87, 0xd2, 0x70, 0x55, 0x5c, 0x3f, 0xe4, 0x69, 0x86, 0x19, 0x8f, 0xb2, 0x8d, 0x82,
	0x06, 0x3f, 0x42, 0xef, 0x22, 0xcd, 0x97, 0x21, 0xfe, 0xa9, 0xc0, 0x8c, 0xa3, 0xcf, 0xa0, 0x47,
	0xd5, 0xe7, 0x22, 0xcd, 0xb0, 0xef, 0x8c, 0x9c, 0x71, 0x6f, 0x3a, 0x9c, 0x28, 0x96, 0x49, 0xc9,
	0x32, 0x59, 0x94, 0x2c, 0xa1, 0x09, 0x47, 0x08, 0x1a, 0x79, 0x94, 0x61, 0xdf, 0x1d, 0x39, 0xe3,
	0x6e, 0x28, 0xbf, 0x83, 0x1f, 0xa0, 0xaf, 0x0e, 0x60, 0x1b, 0x92, 0x33, 0x8c, 0x1e, 0x41, 0x27,
	0xc3, 0x3c, 0x4a, 0x22, 0x1e, 0x69, 0xfa, 0xfb, 0x93, 0x6d, 0xf8, 0x21, 0xde, 0xac, 0x6f, 0xbf,
	0xd1, 0xee, 0x70, 0x0b, 0x44, 0x3e, 0xb4, 0x33, 0xcc, 0x58, 0xb4, 0x2c, 0xb9, 0xcb, 0x65, 0x70,
	0x03, 0x83, 0xaf, 0x08, 0xe3, 0xb3, 0x82, 0xaf, 0x8e, 0x73, 0x87, 0x21, 0x74, 0xa2, 0x82, 0xaf,
	0x9e, 0xe4, 0xd7, 0x44, 0x9e, 0xd5, 0x0f, 0xb7, 0xeb, 0xe0, 0x23, 0x68, 0x7e, 0x41, 0x29, 0xa1,
	0xe2, 0xa2, 0xfc, 0x76, 0xa3, 0xb8, 0xbb, 0xa1, 0xfc, 0x46, 0xa7, 0xe0, 0x65, 0x6c, 0xa9, 0xe3,
	0x13, 0x9f, 0xc1, 0x1c, 0x3a, 0x65, 0x6c, 0xe8, 0x04, 0xdc, 0x34, 0x91, 0xf8, 0x7e, 0xe8, 0xa6,
	0x09, 0xfa, 0x10, 0x5a, 0x58, 0x50, 0x31, 0xdf, 0x1d, 0x79, 0xe3, 0xde, 0x74, 0x50, 0x25, 0x41,
	0x1e, 0x11, 0x6a, 0x77, 0xf0, 0xab, 0x03, 0xf7, 0x6a, 0x69, 0x39, 0xf0, 0x7e, 0x9f, 0x43, 0x9f,
	0xea, 0x5a, 0xc8, 0xed, 0xee, 0xff, 0x6e, 0xaf, 0xe1, 0x83, 0x1b, 0x38, 0xad, 0x12, 0x7e, 0x48,
	0x4d, 0x03, 0xe8, 0x47, 0x06, 0x89, 0xef, 0xc9, 0xdc, 0xd4, 0x6c, 0xc1, 0x33, 0x57, 0x95, 0x77,
	0x8e, 0x29, 0x3f, 0x4e, 0x79, 0x3f, 0x81, 0xee, 0xd3, 0x68, 0x9d, 0x26, 0x5f, 0x52, 0x92, 0xbd,
	0xc0, 0xdd, 0x2b, 0x30, 0xfa, 0x14, 0x40, 0x2e, 0x2e, 0x73, 0x9e, 0xae, 0x65, 0xb4, 0x77, 0x6f,
	0x35, 0xd0, 0xba, 0xfa, 0x8d, 0x6d, 0xf5, 0xdf, 0x86, 0xee, 0x8a, 0x30, 0x2e, 0x1a, 0x84, 0xf9,
	0xcd, 0x91, 0x37, 0xee, 0x86, 0x95, 0x41, 0x78, 0x37, 0xc5, 0xd5, 0x3a, 0x8d, 0xbf, 0xc6, 0xb7,
	0x7e, 0x4b, 0x6e, 0xaa, 0x0c, 0x22, 0x6f, 0x02, 0x5a, 0x66, 0xd4, 0x6f, 0xab, 0xbc, 0x99, 0xb6,
	0xe0, 0x37, 0x47, 0x55, 0x49, 0xe5, 0xed, 0x90, 0x2a, 0x0d, 0xa1, 0xb3, 0xd2, 0x44, 0xba, 0x42,
	0xdb, 0x35, 0x9a, 0x00, 0xe2, 0xb4, 0x60, 0x1c, 0x27, 0x97, 0x0c, 0x53, 0x36, 0x9f, 0x49, 0x94,
	0xba, 0xe5, 0x0e, 0x4f, 0xf0, 0xbb, 0x03, 0x03, 0xb1, 0x3e, 0x6a, 0xb3, 0x16, 0x0c, 0x53, 0xe3,
	0xd1, 0xd9, 0xae, 0xd1, 0x03, 0x68, 0x72, 0x72, 0x83, 0x73, 0x19, 0x50, 0x6f, 0x7a, 0x56, 0xdd,
	0xf5, 0x5b, 0xa1, 0xb1, 0x85, 0xf0, 0x85, 0x0a, 0x12, 0x3c, 0x73, 0xe0, 0xb4, 0x8a, 0xec, 0xc0,
	0x7c, 0xed, 0x8d, 0xe8, 0x1c, 0x5a, 0xe2, 0xfb, 0x49, 0x22, 0x33, 0xd9, 0x0d, 0xf5, 0x0a, 0x9d,
	0x41, 0x53, 0x6a, 0x45, 0x46, 0xda, 0x09, 0xd5, 0xe2, 0xb9, 0xfe, 0x68, 0xee, 0xe8, 0x8f, 0x7f,
	0x3d, 0x95, 0xd1, 0xe3, 0xf5, 0x47, 0x15, 0xa3, 0x5b, 0x8b, 0xd1, 0xbc, 0x97, 0x67, 0xdd, 0xeb,
	0x03, 0x38, 0xa1, 0x38, 0x23, 0x1c, 0x5f, 0x96, 0x88, 0x86, 0x44, 0x58, 0xd6, 0xba, 0xae, 0x9b,
	0xb6, 0xae, 0xc7, 0x30, 0x88, 0x0b, 0x4a, 0x71, 0xce, 0xcb, 0x1b, 0x69, 0xed, 0xdb, 0xe6, 0x7a,
	0x0f, 0xb7, 0x5f, 0xbd, 0x87, 0x3b, 0x2f, 0xd5, 0xc3, 0x53, 0x38, 0x13, 0xb9, 0x27, 0x34, 0xfd,
	0x05, 0x27, 0x17, 0x34, 0xcd, 0xe3, 0x74, 0x13, 0xad, 0x99, 0xdf, 0x95, 0xed, 0xbb, 0xd3, 0x87,
	0xde, 0x83, 0x7b, 0xd7, 0x84, 0xc6, 0x78, 0x4e, 0xb2, 0x2c, 0xca, 0x13, 0xe6, 0x83, 0x04, 0xd7,
	0x8d, 0xc1, 0x9f, 0x0e, 0x40, 0xa5, 0x49, 0x34, 0x82, 0x5e, 0x14, 0xc7, 0x98, 0x31, 0xb9, 0xd4,
	0x33, 0xc6, 0x34, 0x89, 0x44, 0x4a, 0xdd, 0x2e, 0xc4, 0x0c, 0x52, 0x75, 0xaa, 0x0c, 0x42, 0x38,
	0x14, 0x5f, 0x53, 0xcc, 0x14, 0x9f, 0x2e, 0x57, 0xcd, 0x86, 0xa6, 0xd0, 0xc2, 0x3f, 0x6f, 0x52,
	0x7a, 0xab, 0xbb, 0xe3, 0xae, 0x24, 0x68, 0x64, 0xf0, 0x8f, 0x6e, 0x92, 0xa3, 0x3c, 0x2a, 0x7b,
	0x9b, 0x44, 0xfb, 0xcc, 0x07, 0xa7, 0xa8, 0x0a, 0x7f, 0xa2, 0x9f, 0x15, 0xf9, 0xb8, 0xcd, 0x98,
	0xdf, 0x90, 0xc3, 0xf3, 0xb4, 0x3a, 0x52, 0x39, 0x42, 0x0b, 0x17, 0xfc, 0xe5, 0x40, 0x4b, 0x7d,
	0xd7, 0xd5, 0xe3, 0xbc, 0xba, 0x7a, 0xdc, 0x97, 0x52, 0x4f, 0x4d, 0xfb, 0x9e, 0xad, 0xfd, 0x6a,
	0x3e, 0x34, 0xc4, 0x7c, 0x90, 0xe1, 0xca, 0x54, 0xbf, 0x1e, 0xe1, 0x7e, 0x07, 0xe7, 0x17, 0xd2,
	0xb9, 0x50, 0x59, 0x9f, 0xcf, 0x8e, 0xf2, 0x18, 0x05, 0x7f, 0xb8, 0x70, 0xff, 0x39, 0xe2, 0x43,
	0x84, 0xf7, 0x00, 0xda, 0x2b, 0xad, 0x1c, 0x77, 0x8f, 0x72, 0x4a, 0x80, 0xc0, 0xaa, 0x12, 0x30,
	0xdf, 0xb3, 0xb1, 0xca, 0x11, 0x96, 0x00, 0x21, 0x4c, 0x8a, 0x9f, 0x92, 0x9b, 0x17, 0x10, 0x66,
	0x1d, 0x67, 0xec, 0x2c, 0x0f, 0x6b, 0xee, 0x39, 0xcc, 0xc2, 0x4d, 0xff, 0xf6, 0xa0, 0x21, 0xbb,
	0xc2, 0xfc, 0x99, 0xf9, 0x56, 0xfd, 0x40, 0x63, 0xd2, 0x0e, 0x87, 0xbb, 0x5c, 0x7a, 0x8e, 0xbc,
	0x51, 0x92, 0x48, 0x42, 0x8b, 0xc4, 0x18, 0x2e, 0x36, 0x89, 0xf9, 0x14, 0x28, 0x92, 0x72, 0x8a,
	0x9a, 0x24, 0xd6, 0xcc, 0x37, 0x49, 0xec, 0xa1, 0x5b, 0x91, 0xd8, 0x91, 0x58, 0x63, 0xce, 0x26,
	0xb1, 0x22, 0xf9, 0x1e, 0x06, 0x96, 0x70, 0xd0, 0xa8, 0xda, 0xb0, 0x5b, 0xac, 0xc3, 0x77, 0xef,
	0x40, 0x6c, 0x99, 0x3f, 0x86, 0x86, 0xf8, 0x3f, 0x83, 0xde, 0x34, 0xc0, 0xd5, 0x1f, 0xa8, 0xe1,
	0xb9, 0x6d, 0x2e, 0x37, 0x3e, 0x7e, 0x1f, 0xfc, 0x98, 0x64, 0x93, 0x2c, 0x65, 0x7c, 0x12, 0xc5,
	0x31, 0xa1, 0xc9, 0x16, 0xfa, 0xb8, 0x3d, 0x8b, 0xa5, 0xe5, 0xc2, 0xb9, 0x6a, 0x49, 0xe3, 0xa3,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x54, 0x30, 0x8a, 0xd0, 0xd4, 0x0d, 0x00, 0x00,
}
