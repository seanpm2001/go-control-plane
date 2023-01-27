// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: envoy/extensions/http/cache/file_system_http_cache/v3/file_system_http_cache.proto

package file_system_http_cachev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/common/async_files/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Configuration for a cache implementation that caches in the local file system.
//
// By default this cache uses a least-recently-used eviction strategy.
// [#next-free-field: 7]
type FileSystemHttpCacheConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration of a manager for how the file system is used asynchronously.
	ManagerConfig *v3.AsyncFileManagerConfig `protobuf:"bytes,1,opt,name=manager_config,json=managerConfig,proto3" json:"manager_config,omitempty"`
	// Path at which the cache files will be stored.
	//
	// This also doubles as the unique identifier for a cache, so a cache can be shared
	// between different routes, or separate paths can be used to specify separate caches.
	//
	// If the same ``cache_path`` is used in more than one ``CacheConfig``, the rest of the
	// ``FileSystemHttpCacheConfig`` must also match, and will refer to the same cache
	// instance.
	CachePath string `protobuf:"bytes,2,opt,name=cache_path,json=cachePath,proto3" json:"cache_path,omitempty"`
	// The maximum size of the cache in bytes - when reached, another entry is removed.
	//
	// This is measured as the sum of file sizes, such that it includes headers, trailers,
	// and metadata, but does not include e.g. file system overhead and block size padding.
	//
	// If unset there is no limit except file system failure.
	//
	// [#not-implemented-hide:]
	MaxCacheSizeBytes *wrappers.UInt64Value `protobuf:"bytes,3,opt,name=max_cache_size_bytes,json=maxCacheSizeBytes,proto3" json:"max_cache_size_bytes,omitempty"`
	// The maximum size of a cache entry in bytes - larger responses will not be cached.
	//
	// This is measured as the file size for the cache entry, such that it includes
	// headers, trailers, and metadata.
	//
	// If unset there is no limit.
	//
	// [#not-implemented-hide:]
	MaxCacheEntrySizeBytes *wrappers.UInt64Value `protobuf:"bytes,4,opt,name=max_cache_entry_size_bytes,json=maxCacheEntrySizeBytes,proto3" json:"max_cache_entry_size_bytes,omitempty"`
	// The maximum number of cache entries - when reached, another entry is removed.
	//
	// If unset there is no limit.
	//
	// [#not-implemented-hide:]
	MaxCacheEntryCount *wrappers.UInt64Value `protobuf:"bytes,5,opt,name=max_cache_entry_count,json=maxCacheEntryCount,proto3" json:"max_cache_entry_count,omitempty"`
	// A number of folders into which to subdivide the cache.
	//
	// Setting this can help with performance in file systems where a large number of inodes
	// in a single branch degrades performance. The optimal value in that case would be
	// ``sqrt(expected_cache_entry_count)``.
	//
	// On file systems that perform well with many inodes, the default value of 1 should be used.
	//
	// [#not-implemented-hide:]
	CacheSubdivisions uint32 `protobuf:"varint,6,opt,name=cache_subdivisions,json=cacheSubdivisions,proto3" json:"cache_subdivisions,omitempty"`
}

func (x *FileSystemHttpCacheConfig) Reset() {
	*x = FileSystemHttpCacheConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSystemHttpCacheConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSystemHttpCacheConfig) ProtoMessage() {}

func (x *FileSystemHttpCacheConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSystemHttpCacheConfig.ProtoReflect.Descriptor instead.
func (*FileSystemHttpCacheConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescGZIP(), []int{0}
}

func (x *FileSystemHttpCacheConfig) GetManagerConfig() *v3.AsyncFileManagerConfig {
	if x != nil {
		return x.ManagerConfig
	}
	return nil
}

func (x *FileSystemHttpCacheConfig) GetCachePath() string {
	if x != nil {
		return x.CachePath
	}
	return ""
}

func (x *FileSystemHttpCacheConfig) GetMaxCacheSizeBytes() *wrappers.UInt64Value {
	if x != nil {
		return x.MaxCacheSizeBytes
	}
	return nil
}

func (x *FileSystemHttpCacheConfig) GetMaxCacheEntrySizeBytes() *wrappers.UInt64Value {
	if x != nil {
		return x.MaxCacheEntrySizeBytes
	}
	return nil
}

func (x *FileSystemHttpCacheConfig) GetMaxCacheEntryCount() *wrappers.UInt64Value {
	if x != nil {
		return x.MaxCacheEntryCount
	}
	return nil
}

func (x *FileSystemHttpCacheConfig) GetCacheSubdivisions() uint32 {
	if x != nil {
		return x.CacheSubdivisions
	}
	return 0
}

var File_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto protoreflect.FileDescriptor

var file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDesc = []byte{
	0x0a, 0x52, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x3f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x03, 0x0a, 0x19, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x48, 0x74, 0x74, 0x70, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x6f, 0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x0a, 0x0a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x4d, 0x0a, 0x14,
	0x6d, 0x61, 0x78, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x1a, 0x6d,
	0x61, 0x78, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x16, 0x6d,
	0x61, 0x78, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x15, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f,
	0x73, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x11, 0x63, 0x61, 0x63, 0x68, 0x65, 0x53, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0xe8, 0x01, 0x0a, 0x43, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x18, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x48, 0x74, 0x74, 0x70, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x75, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x76, 0x33,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescOnce sync.Once
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescData = file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDesc
)

func file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescGZIP() []byte {
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescData)
	})
	return file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescData
}

var file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_goTypes = []interface{}{
	(*FileSystemHttpCacheConfig)(nil), // 0: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig
	(*v3.AsyncFileManagerConfig)(nil), // 1: envoy.extensions.common.async_files.v3.AsyncFileManagerConfig
	(*wrappers.UInt64Value)(nil),      // 2: google.protobuf.UInt64Value
}
var file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig.manager_config:type_name -> envoy.extensions.common.async_files.v3.AsyncFileManagerConfig
	2, // 1: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig.max_cache_size_bytes:type_name -> google.protobuf.UInt64Value
	2, // 2: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig.max_cache_entry_size_bytes:type_name -> google.protobuf.UInt64Value
	2, // 3: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig.max_cache_entry_count:type_name -> google.protobuf.UInt64Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_init()
}
func file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_init() {
	if File_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSystemHttpCacheConfig); i {
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
			RawDescriptor: file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_msgTypes,
	}.Build()
	File_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto = out.File
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDesc = nil
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_goTypes = nil
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_depIdxs = nil
}