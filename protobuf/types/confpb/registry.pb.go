// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: registry.proto

package confpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Etcd      *EtcdConfig      `protobuf:"bytes,1,opt,name=etcd,proto3,oneof" json:"etcd,omitempty"`
	Consul    *ConsulConfig    `protobuf:"bytes,2,opt,name=consul,proto3,oneof" json:"consul,omitempty"`
	Discovery *DiscoveryConfig `protobuf:"bytes,3,opt,name=discovery,proto3,oneof" json:"discovery,omitempty"`
	Nacos     *NacosConfig     `protobuf:"bytes,4,opt,name=nacos,proto3,oneof" json:"nacos,omitempty"`
	Polaris   *PolarisConfig   `protobuf:"bytes,5,opt,name=polaris,proto3,oneof" json:"polaris,omitempty"`
	Zookeeper *ZookeeperConfig `protobuf:"bytes,6,opt,name=zookeeper,proto3,oneof" json:"zookeeper,omitempty"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0}
}

func (x *Registry) GetEtcd() *EtcdConfig {
	if x != nil {
		return x.Etcd
	}
	return nil
}

func (x *Registry) GetConsul() *ConsulConfig {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *Registry) GetDiscovery() *DiscoveryConfig {
	if x != nil {
		return x.Discovery
	}
	return nil
}

func (x *Registry) GetNacos() *NacosConfig {
	if x != nil {
		return x.Nacos
	}
	return nil
}

func (x *Registry) GetPolaris() *PolarisConfig {
	if x != nil {
		return x.Polaris
	}
	return nil
}

func (x *Registry) GetZookeeper() *ZookeeperConfig {
	if x != nil {
		return x.Zookeeper
	}
	return nil
}

type Discovery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Etcd      *EtcdConfig      `protobuf:"bytes,1,opt,name=etcd,proto3,oneof" json:"etcd,omitempty"`
	Consul    *ConsulConfig    `protobuf:"bytes,2,opt,name=consul,proto3,oneof" json:"consul,omitempty"`
	Discovery *DiscoveryConfig `protobuf:"bytes,3,opt,name=discovery,proto3,oneof" json:"discovery,omitempty"`
	Nacos     *NacosConfig     `protobuf:"bytes,4,opt,name=nacos,proto3,oneof" json:"nacos,omitempty"`
	Polaris   *PolarisConfig   `protobuf:"bytes,5,opt,name=polaris,proto3,oneof" json:"polaris,omitempty"`
	Zookeeper *ZookeeperConfig `protobuf:"bytes,6,opt,name=zookeeper,proto3,oneof" json:"zookeeper,omitempty"`
}

func (x *Discovery) Reset() {
	*x = Discovery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Discovery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discovery) ProtoMessage() {}

func (x *Discovery) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discovery.ProtoReflect.Descriptor instead.
func (*Discovery) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{1}
}

func (x *Discovery) GetEtcd() *EtcdConfig {
	if x != nil {
		return x.Etcd
	}
	return nil
}

func (x *Discovery) GetConsul() *ConsulConfig {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *Discovery) GetDiscovery() *DiscoveryConfig {
	if x != nil {
		return x.Discovery
	}
	return nil
}

func (x *Discovery) GetNacos() *NacosConfig {
	if x != nil {
		return x.Nacos
	}
	return nil
}

func (x *Discovery) GetPolaris() *PolarisConfig {
	if x != nil {
		return x.Polaris
	}
	return nil
}

func (x *Discovery) GetZookeeper() *ZookeeperConfig {
	if x != nil {
		return x.Zookeeper
	}
	return nil
}

type EtcdConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Username  *string  `protobuf:"bytes,2,opt,name=username,proto3,oneof" json:"username,omitempty"`
	Password  *string  `protobuf:"bytes,3,opt,name=password,proto3,oneof" json:"password,omitempty"`
}

func (x *EtcdConfig) Reset() {
	*x = EtcdConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtcdConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtcdConfig) ProtoMessage() {}

func (x *EtcdConfig) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtcdConfig.ProtoReflect.Descriptor instead.
func (*EtcdConfig) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{2}
}

func (x *EtcdConfig) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *EtcdConfig) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *EtcdConfig) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type ConsulConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Username  *string  `protobuf:"bytes,2,opt,name=username,proto3,oneof" json:"username,omitempty"`
	Password  *string  `protobuf:"bytes,3,opt,name=password,proto3,oneof" json:"password,omitempty"`
}

func (x *ConsulConfig) Reset() {
	*x = ConsulConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsulConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsulConfig) ProtoMessage() {}

func (x *ConsulConfig) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsulConfig.ProtoReflect.Descriptor instead.
func (*ConsulConfig) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{3}
}

func (x *ConsulConfig) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ConsulConfig) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *ConsulConfig) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type DiscoveryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Region    string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	Env       string   `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	Zone      string   `protobuf:"bytes,4,opt,name=zone,proto3" json:"zone,omitempty"`
	Host      string   `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *DiscoveryConfig) Reset() {
	*x = DiscoveryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryConfig) ProtoMessage() {}

func (x *DiscoveryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryConfig.ProtoReflect.Descriptor instead.
func (*DiscoveryConfig) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{4}
}

func (x *DiscoveryConfig) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *DiscoveryConfig) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *DiscoveryConfig) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *DiscoveryConfig) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *DiscoveryConfig) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type NacosConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Username  *string  `protobuf:"bytes,2,opt,name=username,proto3,oneof" json:"username,omitempty"`
	Password  *string  `protobuf:"bytes,3,opt,name=password,proto3,oneof" json:"password,omitempty"`
}

func (x *NacosConfig) Reset() {
	*x = NacosConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NacosConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NacosConfig) ProtoMessage() {}

func (x *NacosConfig) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NacosConfig.ProtoReflect.Descriptor instead.
func (*NacosConfig) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{5}
}

func (x *NacosConfig) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *NacosConfig) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *NacosConfig) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type PolarisConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints    []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Namespace    *string  `protobuf:"bytes,2,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	ServiceToken *string  `protobuf:"bytes,3,opt,name=service_token,json=serviceToken,proto3,oneof" json:"service_token,omitempty"`
	Protocol     *string  `protobuf:"bytes,4,opt,name=protocol,proto3,oneof" json:"protocol,omitempty"`
	Weight       *int64   `protobuf:"varint,5,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
	//  optional int64 priority = 6;
	Healthy    *bool                `protobuf:"varint,6,opt,name=healthy,proto3,oneof" json:"healthy,omitempty"`
	Heartbeat  *bool                `protobuf:"varint,7,opt,name=heartbeat,proto3,oneof" json:"heartbeat,omitempty"`
	Isolate    *bool                `protobuf:"varint,8,opt,name=isolate,proto3,oneof" json:"isolate,omitempty"`
	Ttl        *int64               `protobuf:"varint,9,opt,name=ttl,proto3,oneof" json:"ttl,omitempty"`
	Timeout    *durationpb.Duration `protobuf:"bytes,10,opt,name=timeout,proto3,oneof" json:"timeout,omitempty"`
	RetryCount *int64               `protobuf:"varint,11,opt,name=retry_count,json=retryCount,proto3,oneof" json:"retry_count,omitempty"`
}

func (x *PolarisConfig) Reset() {
	*x = PolarisConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolarisConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolarisConfig) ProtoMessage() {}

func (x *PolarisConfig) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolarisConfig.ProtoReflect.Descriptor instead.
func (*PolarisConfig) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{6}
}

func (x *PolarisConfig) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *PolarisConfig) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *PolarisConfig) GetServiceToken() string {
	if x != nil && x.ServiceToken != nil {
		return *x.ServiceToken
	}
	return ""
}

func (x *PolarisConfig) GetProtocol() string {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return ""
}

func (x *PolarisConfig) GetWeight() int64 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

func (x *PolarisConfig) GetHealthy() bool {
	if x != nil && x.Healthy != nil {
		return *x.Healthy
	}
	return false
}

func (x *PolarisConfig) GetHeartbeat() bool {
	if x != nil && x.Heartbeat != nil {
		return *x.Heartbeat
	}
	return false
}

func (x *PolarisConfig) GetIsolate() bool {
	if x != nil && x.Isolate != nil {
		return *x.Isolate
	}
	return false
}

func (x *PolarisConfig) GetTtl() int64 {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return 0
}

func (x *PolarisConfig) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *PolarisConfig) GetRetryCount() int64 {
	if x != nil && x.RetryCount != nil {
		return *x.RetryCount
	}
	return 0
}

type ZookeeperConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string             `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Namespace *string              `protobuf:"bytes,2,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Username  *string              `protobuf:"bytes,3,opt,name=username,proto3,oneof" json:"username,omitempty"`
	Password  *string              `protobuf:"bytes,4,opt,name=password,proto3,oneof" json:"password,omitempty"`
	Timeout   *durationpb.Duration `protobuf:"bytes,5,opt,name=timeout,proto3,oneof" json:"timeout,omitempty"`
}

func (x *ZookeeperConfig) Reset() {
	*x = ZookeeperConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZookeeperConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZookeeperConfig) ProtoMessage() {}

func (x *ZookeeperConfig) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZookeeperConfig.ProtoReflect.Descriptor instead.
func (*ZookeeperConfig) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{7}
}

func (x *ZookeeperConfig) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ZookeeperConfig) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *ZookeeperConfig) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *ZookeeperConfig) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *ZookeeperConfig) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

var File_registry_proto protoreflect.FileDescriptor

var file_registry_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x03,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x04, 0x65, 0x74,
	0x63, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x48, 0x00, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x02, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x05, 0x6e, 0x61, 0x63, 0x6f, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x4e, 0x61, 0x63, 0x6f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x03,
	0x52, 0x05, 0x6e, 0x61, 0x63, 0x6f, 0x73, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x6f,
	0x6c, 0x61, 0x72, 0x69, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x50, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x04, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x61, 0x72,
	0x69, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x09, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x5a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x05, 0x52, 0x09, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6e, 0x61, 0x63,
	0x6f, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x22, 0xad, 0x03, 0x0a,
	0x09, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x04, 0x65, 0x74,
	0x63, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x48, 0x00, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x02, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x05, 0x6e, 0x61, 0x63, 0x6f, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x4e, 0x61, 0x63, 0x6f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x03,
	0x52, 0x05, 0x6e, 0x61, 0x63, 0x6f, 0x73, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x6f,
	0x6c, 0x61, 0x72, 0x69, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x50, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x04, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x61, 0x72,
	0x69, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x09, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x5a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x05, 0x52, 0x09, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6e, 0x61, 0x63,
	0x6f, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x22, 0x86, 0x01, 0x0a,
	0x0a, 0x45, 0x74, 0x63, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x81, 0x01, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e,
	0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x12, 0x0a, 0x04,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x4e, 0x61, 0x63, 0x6f, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x92,
	0x04, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x21,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x28, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x07, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x79, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05, 0x52, 0x09, 0x68,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x69,
	0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x07,
	0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x74, 0x74,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x48, 0x07, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x38, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x08, 0x52,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x72,
	0x65, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x74, 0x74, 0x6c, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x82, 0x02, 0x0a, 0x0f, 0x5a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x65, 0x73, 0x6f, 0x70, 0x65, 0x72, 0x31, 0x30,
	0x31, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x70, 0x62, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_registry_proto_rawDescOnce sync.Once
	file_registry_proto_rawDescData = file_registry_proto_rawDesc
)

func file_registry_proto_rawDescGZIP() []byte {
	file_registry_proto_rawDescOnce.Do(func() {
		file_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_proto_rawDescData)
	})
	return file_registry_proto_rawDescData
}

var file_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_registry_proto_goTypes = []interface{}{
	(*Registry)(nil),            // 0: kratos.conf.Registry
	(*Discovery)(nil),           // 1: kratos.conf.Discovery
	(*EtcdConfig)(nil),          // 2: kratos.conf.EtcdConfig
	(*ConsulConfig)(nil),        // 3: kratos.conf.ConsulConfig
	(*DiscoveryConfig)(nil),     // 4: kratos.conf.DiscoveryConfig
	(*NacosConfig)(nil),         // 5: kratos.conf.NacosConfig
	(*PolarisConfig)(nil),       // 6: kratos.conf.PolarisConfig
	(*ZookeeperConfig)(nil),     // 7: kratos.conf.ZookeeperConfig
	(*durationpb.Duration)(nil), // 8: google.protobuf.Duration
}
var file_registry_proto_depIdxs = []int32{
	2,  // 0: kratos.conf.Registry.etcd:type_name -> kratos.conf.EtcdConfig
	3,  // 1: kratos.conf.Registry.consul:type_name -> kratos.conf.ConsulConfig
	4,  // 2: kratos.conf.Registry.discovery:type_name -> kratos.conf.DiscoveryConfig
	5,  // 3: kratos.conf.Registry.nacos:type_name -> kratos.conf.NacosConfig
	6,  // 4: kratos.conf.Registry.polaris:type_name -> kratos.conf.PolarisConfig
	7,  // 5: kratos.conf.Registry.zookeeper:type_name -> kratos.conf.ZookeeperConfig
	2,  // 6: kratos.conf.Discovery.etcd:type_name -> kratos.conf.EtcdConfig
	3,  // 7: kratos.conf.Discovery.consul:type_name -> kratos.conf.ConsulConfig
	4,  // 8: kratos.conf.Discovery.discovery:type_name -> kratos.conf.DiscoveryConfig
	5,  // 9: kratos.conf.Discovery.nacos:type_name -> kratos.conf.NacosConfig
	6,  // 10: kratos.conf.Discovery.polaris:type_name -> kratos.conf.PolarisConfig
	7,  // 11: kratos.conf.Discovery.zookeeper:type_name -> kratos.conf.ZookeeperConfig
	8,  // 12: kratos.conf.PolarisConfig.timeout:type_name -> google.protobuf.Duration
	8,  // 13: kratos.conf.ZookeeperConfig.timeout:type_name -> google.protobuf.Duration
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_registry_proto_init() }
func file_registry_proto_init() {
	if File_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
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
		file_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Discovery); i {
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
		file_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EtcdConfig); i {
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
		file_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsulConfig); i {
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
		file_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveryConfig); i {
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
		file_registry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NacosConfig); i {
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
		file_registry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolarisConfig); i {
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
		file_registry_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZookeeperConfig); i {
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
	file_registry_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_registry_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_registry_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_registry_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_registry_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_registry_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_registry_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_registry_proto_goTypes,
		DependencyIndexes: file_registry_proto_depIdxs,
		MessageInfos:      file_registry_proto_msgTypes,
	}.Build()
	File_registry_proto = out.File
	file_registry_proto_rawDesc = nil
	file_registry_proto_goTypes = nil
	file_registry_proto_depIdxs = nil
}
