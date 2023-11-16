// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/extensions/load_balancing_policies/subset/v3/subset.proto

package subsetv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/config/cluster/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// If NO_FALLBACK is selected, a result
// equivalent to no healthy hosts is reported. If ANY_ENDPOINT is selected,
// any cluster endpoint may be returned (subject to policy, health checks,
// etc). If DEFAULT_SUBSET is selected, load balancing is performed over the
// endpoints matching the values from the default_subset field.
type Subset_LbSubsetFallbackPolicy int32

const (
	Subset_NO_FALLBACK    Subset_LbSubsetFallbackPolicy = 0
	Subset_ANY_ENDPOINT   Subset_LbSubsetFallbackPolicy = 1
	Subset_DEFAULT_SUBSET Subset_LbSubsetFallbackPolicy = 2
)

// Enum value maps for Subset_LbSubsetFallbackPolicy.
var (
	Subset_LbSubsetFallbackPolicy_name = map[int32]string{
		0: "NO_FALLBACK",
		1: "ANY_ENDPOINT",
		2: "DEFAULT_SUBSET",
	}
	Subset_LbSubsetFallbackPolicy_value = map[string]int32{
		"NO_FALLBACK":    0,
		"ANY_ENDPOINT":   1,
		"DEFAULT_SUBSET": 2,
	}
)

func (x Subset_LbSubsetFallbackPolicy) Enum() *Subset_LbSubsetFallbackPolicy {
	p := new(Subset_LbSubsetFallbackPolicy)
	*p = x
	return p
}

func (x Subset_LbSubsetFallbackPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subset_LbSubsetFallbackPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_enumTypes[0].Descriptor()
}

func (Subset_LbSubsetFallbackPolicy) Type() protoreflect.EnumType {
	return &file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_enumTypes[0]
}

func (x Subset_LbSubsetFallbackPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subset_LbSubsetFallbackPolicy.Descriptor instead.
func (Subset_LbSubsetFallbackPolicy) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescGZIP(), []int{0, 0}
}

type Subset_LbSubsetMetadataFallbackPolicy int32

const (
	// No fallback. Route metadata will be used as-is.
	Subset_METADATA_NO_FALLBACK Subset_LbSubsetMetadataFallbackPolicy = 0
	// A special metadata key “fallback_list“ will be used to provide variants of metadata to try.
	// Value of “fallback_list“ key has to be a list. Every list element has to be a struct - it will
	// be merged with route metadata, overriding keys that appear in both places.
	// “fallback_list“ entries will be used in order until a host is found.
	//
	// “fallback_list“ key itself is removed from metadata before subset load balancing is performed.
	//
	// Example:
	//
	// for metadata:
	//
	// .. code-block:: yaml
	//
	//	version: 1.0
	//	fallback_list:
	//	  - version: 2.0
	//	    hardware: c64
	//	  - hardware: c32
	//	  - version: 3.0
	//
	// at first, metadata:
	//
	// .. code-block:: json
	//
	//	{"version": "2.0", "hardware": "c64"}
	//
	// will be used for load balancing. If no host is found, metadata:
	//
	// .. code-block:: json
	//
	//	{"version": "1.0", "hardware": "c32"}
	//
	// is next to try. If it still results in no host, finally metadata:
	//
	// .. code-block:: json
	//
	//	{"version": "3.0"}
	//
	// is used.
	Subset_FALLBACK_LIST Subset_LbSubsetMetadataFallbackPolicy = 1
)

// Enum value maps for Subset_LbSubsetMetadataFallbackPolicy.
var (
	Subset_LbSubsetMetadataFallbackPolicy_name = map[int32]string{
		0: "METADATA_NO_FALLBACK",
		1: "FALLBACK_LIST",
	}
	Subset_LbSubsetMetadataFallbackPolicy_value = map[string]int32{
		"METADATA_NO_FALLBACK": 0,
		"FALLBACK_LIST":        1,
	}
)

func (x Subset_LbSubsetMetadataFallbackPolicy) Enum() *Subset_LbSubsetMetadataFallbackPolicy {
	p := new(Subset_LbSubsetMetadataFallbackPolicy)
	*p = x
	return p
}

func (x Subset_LbSubsetMetadataFallbackPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subset_LbSubsetMetadataFallbackPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_enumTypes[1].Descriptor()
}

func (Subset_LbSubsetMetadataFallbackPolicy) Type() protoreflect.EnumType {
	return &file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_enumTypes[1]
}

func (x Subset_LbSubsetMetadataFallbackPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subset_LbSubsetMetadataFallbackPolicy.Descriptor instead.
func (Subset_LbSubsetMetadataFallbackPolicy) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescGZIP(), []int{0, 1}
}

// Allows to override top level fallback policy per selector.
type Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy int32

const (
	// If NOT_DEFINED top level config fallback policy is used instead.
	Subset_LbSubsetSelector_NOT_DEFINED Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy = 0
	// If NO_FALLBACK is selected, a result equivalent to no healthy hosts is reported.
	Subset_LbSubsetSelector_NO_FALLBACK Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy = 1
	// If ANY_ENDPOINT is selected, any cluster endpoint may be returned
	// (subject to policy, health checks, etc).
	Subset_LbSubsetSelector_ANY_ENDPOINT Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy = 2
	// If DEFAULT_SUBSET is selected, load balancing is performed over the
	// endpoints matching the values from the default_subset field.
	Subset_LbSubsetSelector_DEFAULT_SUBSET Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy = 3
	// If KEYS_SUBSET is selected, subset selector matching is performed again with metadata
	// keys reduced to
	// :ref:`fallback_keys_subset<envoy_v3_api_field_extensions.load_balancing_policies.subset.v3.Subset.LbSubsetSelector.fallback_keys_subset>`.
	// It allows for a fallback to a different, less specific selector if some of the keys of
	// the selector are considered optional.
	Subset_LbSubsetSelector_KEYS_SUBSET Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy = 4
)

// Enum value maps for Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy.
var (
	Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy_name = map[int32]string{
		0: "NOT_DEFINED",
		1: "NO_FALLBACK",
		2: "ANY_ENDPOINT",
		3: "DEFAULT_SUBSET",
		4: "KEYS_SUBSET",
	}
	Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy_value = map[string]int32{
		"NOT_DEFINED":    0,
		"NO_FALLBACK":    1,
		"ANY_ENDPOINT":   2,
		"DEFAULT_SUBSET": 3,
		"KEYS_SUBSET":    4,
	}
)

func (x Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy) Enum() *Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy {
	p := new(Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy)
	*p = x
	return p
}

func (x Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_enumTypes[2].Descriptor()
}

func (Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy) Type() protoreflect.EnumType {
	return &file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_enumTypes[2]
}

func (x Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy.Descriptor instead.
func (Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Optionally divide the endpoints in this cluster into subsets defined by
// endpoint metadata and selected by route and weighted cluster metadata.
// [#next-free-field: 11]
type Subset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The behavior used when no endpoint subset matches the selected route's
	// metadata. The value defaults to
	// :ref:`NO_FALLBACK<envoy_v3_api_enum_value_extensions.load_balancing_policies.subset.v3.Subset.LbSubsetFallbackPolicy.NO_FALLBACK>`.
	FallbackPolicy Subset_LbSubsetFallbackPolicy `protobuf:"varint,1,opt,name=fallback_policy,json=fallbackPolicy,proto3,enum=envoy.extensions.load_balancing_policies.subset.v3.Subset_LbSubsetFallbackPolicy" json:"fallback_policy,omitempty"`
	// Specifies the default subset of endpoints used during fallback if
	// fallback_policy is
	// :ref:`DEFAULT_SUBSET<envoy_v3_api_enum_value_extensions.load_balancing_policies.subset.v3.Subset.LbSubsetFallbackPolicy.DEFAULT_SUBSET>`.
	// Each field in default_subset is
	// compared to the matching LbEndpoint.Metadata under the “envoy.lb“
	// namespace. It is valid for no hosts to match, in which case the behavior
	// is the same as a fallback_policy of
	// :ref:`NO_FALLBACK<envoy_v3_api_enum_value_extensions.load_balancing_policies.subset.v3.Subset.LbSubsetFallbackPolicy.NO_FALLBACK>`.
	DefaultSubset *_struct.Struct `protobuf:"bytes,2,opt,name=default_subset,json=defaultSubset,proto3" json:"default_subset,omitempty"`
	// For each entry, LbEndpoint.Metadata's
	// “envoy.lb“ namespace is traversed and a subset is created for each unique
	// combination of key and value. For example:
	//
	// .. code-block:: json
	//
	//	{ "subset_selectors": [
	//	    { "keys": [ "version" ] },
	//	    { "keys": [ "stage", "hardware_type" ] }
	//	]}
	//
	// A subset is matched when the metadata from the selected route and
	// weighted cluster contains the same keys and values as the subset's
	// metadata. The same host may appear in multiple subsets.
	SubsetSelectors []*Subset_LbSubsetSelector `protobuf:"bytes,3,rep,name=subset_selectors,json=subsetSelectors,proto3" json:"subset_selectors,omitempty"`
	// By default, only when the request metadata has exactly the **same** keys as one of subset selectors and
	// the values of the related keys are matched, the load balancer will have a valid subset for the request.
	// For example, given the following subset selectors:
	//
	// .. code-block:: json
	//
	//	{ "subset_selectors": [
	//	    { "keys": [ "version" ] },
	//	    { "keys": [ "stage", "version" ] }
	//	]}
	//
	// A request with metadata “{"redundant-key": "redundant-value", "stage": "prod", "version": "v1"}“ or
	// “{"redundant-key": "redundant-value", "version": "v1"}“ will not have a valid subset even if the values
	// of keys “stage“ and “version“ are matched because of the redundant key/value pair in the request
	// metadata.
	//
	// By setting this field to true, the most appropriate keys will be filtered out from the request metadata
	// according to the subset selectors. And then the filtered keys and related values will be matched to find
	// the valid host subset. By this way, redundant key/value pairs are allowed in the request metadata. The keys
	// of a request metadata could be superset of the keys of the subset selectors and need not to be exactly the
	// same as the keys of the subset selectors.
	//
	// More specifically, if the keys of a request metadata is a superset of one of the subset selectors, then only
	// the values of the keys that in the selector keys will be matched. Take the above example, if the request
	// metadata is “{"redundant-key": "redundant-value", "stage": "prod", "version": "v1"}“, the load balancer
	// will only match the values of “stage“ and “version“ to find an appropriate subset because “stage“
	// “version“ are contained by the second subset selector and the redundant “redundant-key“ will be
	// ignored.
	//
	// .. note::
	//
	//	If the keys of request metadata is superset of multiple different subset selectors keys, the subset
	//	selector with most keys to win. For example, given subset selectors
	//	``{"subset_selectors": ["keys": ["A", "B", "C"], ["A", "B"]]}`` and request metadata ``{"A": "-",
	//	"B": "-", "C": "-", "D": "-"}``, keys ``A``, ``B``, ``C`` will be evaluated.
	//	If the keys of request metadata is superset of multiple different subset selectors keys and the number
	//	of selector keys are same, then the one placed in front to win. For example, given subset selectors
	//	``{"subset_selectors": ["keys": ["A", "B"], ["C", "D"]]}`` and request metadata ``{"A": "-", "B": "-",
	//	"C": "-", "D": "-"}``, keys ``A``, ``B`` will be evaluated.
	AllowRedundantKeys bool `protobuf:"varint,10,opt,name=allow_redundant_keys,json=allowRedundantKeys,proto3" json:"allow_redundant_keys,omitempty"`
	// If true, routing to subsets will take into account the localities and locality weights of the
	// endpoints when making the routing decision.
	//
	// There are some potential pitfalls associated with enabling this feature, as the resulting
	// traffic split after applying both a subset match and locality weights might be undesirable.
	//
	// Consider for example a situation in which you have 50/50 split across two localities X/Y
	// which have 100 hosts each without subsetting. If the subset LB results in X having only 1
	// host selected but Y having 100, then a lot more load is being dumped on the single host in X
	// than originally anticipated in the load balancing assignment delivered via EDS.
	LocalityWeightAware bool `protobuf:"varint,4,opt,name=locality_weight_aware,json=localityWeightAware,proto3" json:"locality_weight_aware,omitempty"`
	// When used with locality_weight_aware, scales the weight of each locality by the ratio
	// of hosts in the subset vs hosts in the original subset. This aims to even out the load
	// going to an individual locality if said locality is disproportionately affected by the
	// subset predicate.
	ScaleLocalityWeight bool `protobuf:"varint,5,opt,name=scale_locality_weight,json=scaleLocalityWeight,proto3" json:"scale_locality_weight,omitempty"`
	// If true, when a fallback policy is configured and its corresponding subset fails to find
	// a host this will cause any host to be selected instead.
	//
	// This is useful when using the default subset as the fallback policy, given the default
	// subset might become empty. With this option enabled, if that happens the LB will attempt
	// to select a host from the entire cluster.
	PanicModeAny bool `protobuf:"varint,6,opt,name=panic_mode_any,json=panicModeAny,proto3" json:"panic_mode_any,omitempty"`
	// If true, metadata specified for a metadata key will be matched against the corresponding
	// endpoint metadata if the endpoint metadata matches the value exactly OR it is a list value
	// and any of the elements in the list matches the criteria.
	ListAsAny bool `protobuf:"varint,7,opt,name=list_as_any,json=listAsAny,proto3" json:"list_as_any,omitempty"`
	// Fallback mechanism that allows to try different route metadata until a host is found.
	// If load balancing process, including all its mechanisms (like
	// :ref:`fallback_policy<envoy_v3_api_field_extensions.load_balancing_policies.subset.v3.subset.fallback_policy>`)
	// fails to select a host, this policy decides if and how the process is repeated using another metadata.
	//
	// The value defaults to
	// :ref:`METADATA_NO_FALLBACK
	// <envoy_v3_api_enum_value_extensions.load_balancing_policies.subset.v3.subset.LbSubsetMetadataFallbackPolicy.METADATA_NO_FALLBACK>`.
	MetadataFallbackPolicy Subset_LbSubsetMetadataFallbackPolicy `protobuf:"varint,8,opt,name=metadata_fallback_policy,json=metadataFallbackPolicy,proto3,enum=envoy.extensions.load_balancing_policies.subset.v3.Subset_LbSubsetMetadataFallbackPolicy" json:"metadata_fallback_policy,omitempty"`
	// The child LB policy to create for endpoint-picking within the chosen subset.
	SubsetLbPolicy *v3.LoadBalancingPolicy `protobuf:"bytes,9,opt,name=subset_lb_policy,json=subsetLbPolicy,proto3" json:"subset_lb_policy,omitempty"`
}

func (x *Subset) Reset() {
	*x = Subset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subset) ProtoMessage() {}

func (x *Subset) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subset.ProtoReflect.Descriptor instead.
func (*Subset) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescGZIP(), []int{0}
}

func (x *Subset) GetFallbackPolicy() Subset_LbSubsetFallbackPolicy {
	if x != nil {
		return x.FallbackPolicy
	}
	return Subset_NO_FALLBACK
}

func (x *Subset) GetDefaultSubset() *_struct.Struct {
	if x != nil {
		return x.DefaultSubset
	}
	return nil
}

func (x *Subset) GetSubsetSelectors() []*Subset_LbSubsetSelector {
	if x != nil {
		return x.SubsetSelectors
	}
	return nil
}

func (x *Subset) GetAllowRedundantKeys() bool {
	if x != nil {
		return x.AllowRedundantKeys
	}
	return false
}

func (x *Subset) GetLocalityWeightAware() bool {
	if x != nil {
		return x.LocalityWeightAware
	}
	return false
}

func (x *Subset) GetScaleLocalityWeight() bool {
	if x != nil {
		return x.ScaleLocalityWeight
	}
	return false
}

func (x *Subset) GetPanicModeAny() bool {
	if x != nil {
		return x.PanicModeAny
	}
	return false
}

func (x *Subset) GetListAsAny() bool {
	if x != nil {
		return x.ListAsAny
	}
	return false
}

func (x *Subset) GetMetadataFallbackPolicy() Subset_LbSubsetMetadataFallbackPolicy {
	if x != nil {
		return x.MetadataFallbackPolicy
	}
	return Subset_METADATA_NO_FALLBACK
}

func (x *Subset) GetSubsetLbPolicy() *v3.LoadBalancingPolicy {
	if x != nil {
		return x.SubsetLbPolicy
	}
	return nil
}

// Specifications for subsets.
type Subset_LbSubsetSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of keys to match with the weighted cluster metadata.
	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	// Selects a mode of operation in which each subset has only one host. This mode uses the same rules for
	// choosing a host, but updating hosts is faster, especially for large numbers of hosts.
	//
	// If a match is found to a host, that host will be used regardless of priority levels.
	//
	// When this mode is enabled, configurations that contain more than one host with the same metadata value for the single key in “keys“
	// will use only one of the hosts with the given key; no requests will be routed to the others. The cluster gauge
	// :ref:`lb_subsets_single_host_per_subset_duplicate<config_cluster_manager_cluster_stats_subset_lb>` indicates how many duplicates are
	// present in the current configuration.
	SingleHostPerSubset bool `protobuf:"varint,4,opt,name=single_host_per_subset,json=singleHostPerSubset,proto3" json:"single_host_per_subset,omitempty"`
	// The behavior used when no endpoint subset matches the selected route's
	// metadata.
	FallbackPolicy Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy `protobuf:"varint,2,opt,name=fallback_policy,json=fallbackPolicy,proto3,enum=envoy.extensions.load_balancing_policies.subset.v3.Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy" json:"fallback_policy,omitempty"`
	// Subset of
	// :ref:`keys<envoy_v3_api_field_extensions.load_balancing_policies.subset.v3.Subset.LbSubsetSelector.keys>` used by
	// :ref:`KEYS_SUBSET<envoy_v3_api_enum_value_extensions.load_balancing_policies.subset.v3.Subset.LbSubsetSelector.LbSubsetSelectorFallbackPolicy.KEYS_SUBSET>`
	// fallback policy.
	// It has to be a non empty list if KEYS_SUBSET fallback policy is selected.
	// For any other fallback policy the parameter is not used and should not be set.
	// Only values also present in
	// :ref:`keys<envoy_v3_api_field_extensions.load_balancing_policies.subset.v3.Subset.LbSubsetSelector.keys>` are allowed, but
	// “fallback_keys_subset“ cannot be equal to “keys“.
	FallbackKeysSubset []string `protobuf:"bytes,3,rep,name=fallback_keys_subset,json=fallbackKeysSubset,proto3" json:"fallback_keys_subset,omitempty"`
}

func (x *Subset_LbSubsetSelector) Reset() {
	*x = Subset_LbSubsetSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subset_LbSubsetSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subset_LbSubsetSelector) ProtoMessage() {}

func (x *Subset_LbSubsetSelector) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subset_LbSubsetSelector.ProtoReflect.Descriptor instead.
func (*Subset_LbSubsetSelector) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Subset_LbSubsetSelector) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Subset_LbSubsetSelector) GetSingleHostPerSubset() bool {
	if x != nil {
		return x.SingleHostPerSubset
	}
	return false
}

func (x *Subset_LbSubsetSelector) GetFallbackPolicy() Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy {
	if x != nil {
		return x.FallbackPolicy
	}
	return Subset_LbSubsetSelector_NOT_DEFINED
}

func (x *Subset_LbSubsetSelector) GetFallbackKeysSubset() []string {
	if x != nil {
		return x.FallbackKeysSubset
	}
	return nil
}

var File_envoy_extensions_load_balancing_policies_subset_v3_subset_proto protoreflect.FileDescriptor

var file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x65,
	0x74, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x0b, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74,
	0x12, 0x84, 0x01, 0x0a, 0x0f, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x51, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x33, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x2e, 0x4c, 0x62, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x46,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3e, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x12, 0x76, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x73, 0x65,
	0x74, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x4b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x65, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x2e, 0x4c, 0x62,
	0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0f,
	0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x30, 0x0a, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x6e, 0x64, 0x61,
	0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x64, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x4b, 0x65, 0x79,
	0x73, 0x12, 0x32, 0x0a, 0x15, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x61, 0x77, 0x61, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x41, 0x77, 0x61, 0x72, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x6e,
	0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x6e, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x41, 0x6e, 0x79, 0x12,
	0x1e, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x61, 0x73, 0x5f, 0x61, 0x6e, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x41, 0x6e, 0x79, 0x12,
	0x9d, 0x01, 0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x66, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x59, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x73, 0x75,
	0x62, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x2e, 0x4c,
	0x62, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x16, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x60, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x5f, 0x6c, 0x62, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x4c, 0x62, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x1a, 0xa8, 0x03, 0x0a, 0x10, 0x4c, 0x62, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x73, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x75,
	0x62, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x12,
	0x9d, 0x01, 0x0a, 0x0f, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x6a, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x65, 0x74, 0x2e, 0x4c, 0x62, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x62, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x0e, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x30, 0x0a, 0x14, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x6b, 0x65, 0x79, 0x73,
	0x5f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x66,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x53, 0x75, 0x62, 0x73, 0x65,
	0x74, 0x22, 0x79, 0x0a, 0x1e, 0x4c, 0x62, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x5f, 0x46, 0x41, 0x4c, 0x4c, 0x42,
	0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4e, 0x59, 0x5f, 0x45, 0x4e, 0x44,
	0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x45, 0x54, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4b,
	0x45, 0x59, 0x53, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x45, 0x54, 0x10, 0x04, 0x22, 0x4f, 0x0a, 0x16,
	0x4c, 0x62, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x5f, 0x46, 0x41, 0x4c,
	0x4c, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4e, 0x59, 0x5f, 0x45,
	0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45, 0x46,
	0x41, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x45, 0x54, 0x10, 0x02, 0x22, 0x4d, 0x0a,
	0x1e, 0x4c, 0x62, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x18, 0x0a, 0x14, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4e, 0x4f, 0x5f, 0x46,
	0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x41, 0x4c,
	0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x01, 0x3a, 0x2d, 0x9a, 0xc5,
	0x88, 0x1e, 0x28, 0x0a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x62, 0x53,
	0x75, 0x62, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xbd, 0x01, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x40, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x2f,
	0x76, 0x33, 0x3b, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescOnce sync.Once
	file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescData = file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDesc
)

func file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescGZIP() []byte {
	file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescData)
	})
	return file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDescData
}

var file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_goTypes = []interface{}{
	(Subset_LbSubsetFallbackPolicy)(0),                          // 0: envoy.extensions.load_balancing_policies.subset.v3.Subset.LbSubsetFallbackPolicy
	(Subset_LbSubsetMetadataFallbackPolicy)(0),                  // 1: envoy.extensions.load_balancing_policies.subset.v3.Subset.LbSubsetMetadataFallbackPolicy
	(Subset_LbSubsetSelector_LbSubsetSelectorFallbackPolicy)(0), // 2: envoy.extensions.load_balancing_policies.subset.v3.Subset.LbSubsetSelector.LbSubsetSelectorFallbackPolicy
	(*Subset)(nil),                  // 3: envoy.extensions.load_balancing_policies.subset.v3.Subset
	(*Subset_LbSubsetSelector)(nil), // 4: envoy.extensions.load_balancing_policies.subset.v3.Subset.LbSubsetSelector
	(*_struct.Struct)(nil),          // 5: google.protobuf.Struct
	(*v3.LoadBalancingPolicy)(nil),  // 6: envoy.config.cluster.v3.LoadBalancingPolicy
}
var file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_depIdxs = []int32{
	0, // 0: envoy.extensions.load_balancing_policies.subset.v3.Subset.fallback_policy:type_name -> envoy.extensions.load_balancing_policies.subset.v3.Subset.LbSubsetFallbackPolicy
	5, // 1: envoy.extensions.load_balancing_policies.subset.v3.Subset.default_subset:type_name -> google.protobuf.Struct
	4, // 2: envoy.extensions.load_balancing_policies.subset.v3.Subset.subset_selectors:type_name -> envoy.extensions.load_balancing_policies.subset.v3.Subset.LbSubsetSelector
	1, // 3: envoy.extensions.load_balancing_policies.subset.v3.Subset.metadata_fallback_policy:type_name -> envoy.extensions.load_balancing_policies.subset.v3.Subset.LbSubsetMetadataFallbackPolicy
	6, // 4: envoy.extensions.load_balancing_policies.subset.v3.Subset.subset_lb_policy:type_name -> envoy.config.cluster.v3.LoadBalancingPolicy
	2, // 5: envoy.extensions.load_balancing_policies.subset.v3.Subset.LbSubsetSelector.fallback_policy:type_name -> envoy.extensions.load_balancing_policies.subset.v3.Subset.LbSubsetSelector.LbSubsetSelectorFallbackPolicy
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_init() }
func file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_init() {
	if File_envoy_extensions_load_balancing_policies_subset_v3_subset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subset); i {
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
		file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subset_LbSubsetSelector); i {
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
			RawDescriptor: file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_msgTypes,
	}.Build()
	File_envoy_extensions_load_balancing_policies_subset_v3_subset_proto = out.File
	file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_rawDesc = nil
	file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_goTypes = nil
	file_envoy_extensions_load_balancing_policies_subset_v3_subset_proto_depIdxs = nil
}
