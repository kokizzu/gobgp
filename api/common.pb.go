// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/common.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Family_Afi int32

const (
	Family_AFI_UNSPECIFIED Family_Afi = 0
	Family_AFI_IP          Family_Afi = 1
	Family_AFI_IP6         Family_Afi = 2
	Family_AFI_L2VPN       Family_Afi = 25
	Family_AFI_LS          Family_Afi = 16388
	Family_AFI_OPAQUE      Family_Afi = 16397
)

// Enum value maps for Family_Afi.
var (
	Family_Afi_name = map[int32]string{
		0:     "AFI_UNSPECIFIED",
		1:     "AFI_IP",
		2:     "AFI_IP6",
		25:    "AFI_L2VPN",
		16388: "AFI_LS",
		16397: "AFI_OPAQUE",
	}
	Family_Afi_value = map[string]int32{
		"AFI_UNSPECIFIED": 0,
		"AFI_IP":          1,
		"AFI_IP6":         2,
		"AFI_L2VPN":       25,
		"AFI_LS":          16388,
		"AFI_OPAQUE":      16397,
	}
)

func (x Family_Afi) Enum() *Family_Afi {
	p := new(Family_Afi)
	*p = x
	return p
}

func (x Family_Afi) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Family_Afi) Descriptor() protoreflect.EnumDescriptor {
	return file_api_common_proto_enumTypes[0].Descriptor()
}

func (Family_Afi) Type() protoreflect.EnumType {
	return &file_api_common_proto_enumTypes[0]
}

func (x Family_Afi) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Family_Afi.Descriptor instead.
func (Family_Afi) EnumDescriptor() ([]byte, []int) {
	return file_api_common_proto_rawDescGZIP(), []int{0, 0}
}

type Family_Safi int32

const (
	Family_SAFI_UNSPECIFIED              Family_Safi = 0
	Family_SAFI_UNICAST                  Family_Safi = 1
	Family_SAFI_MULTICAST                Family_Safi = 2
	Family_SAFI_MPLS_LABEL               Family_Safi = 4
	Family_SAFI_ENCAPSULATION            Family_Safi = 7
	Family_SAFI_VPLS                     Family_Safi = 65
	Family_SAFI_EVPN                     Family_Safi = 70
	Family_SAFI_LS                       Family_Safi = 71
	Family_SAFI_SR_POLICY                Family_Safi = 73
	Family_SAFI_MUP                      Family_Safi = 85
	Family_SAFI_MPLS_VPN                 Family_Safi = 128
	Family_SAFI_MPLS_VPN_MULTICAST       Family_Safi = 129
	Family_SAFI_ROUTE_TARGET_CONSTRAINTS Family_Safi = 132
	Family_SAFI_FLOW_SPEC_UNICAST        Family_Safi = 133
	Family_SAFI_FLOW_SPEC_VPN            Family_Safi = 134
	Family_SAFI_KEY_VALUE                Family_Safi = 241
)

// Enum value maps for Family_Safi.
var (
	Family_Safi_name = map[int32]string{
		0:   "SAFI_UNSPECIFIED",
		1:   "SAFI_UNICAST",
		2:   "SAFI_MULTICAST",
		4:   "SAFI_MPLS_LABEL",
		7:   "SAFI_ENCAPSULATION",
		65:  "SAFI_VPLS",
		70:  "SAFI_EVPN",
		71:  "SAFI_LS",
		73:  "SAFI_SR_POLICY",
		85:  "SAFI_MUP",
		128: "SAFI_MPLS_VPN",
		129: "SAFI_MPLS_VPN_MULTICAST",
		132: "SAFI_ROUTE_TARGET_CONSTRAINTS",
		133: "SAFI_FLOW_SPEC_UNICAST",
		134: "SAFI_FLOW_SPEC_VPN",
		241: "SAFI_KEY_VALUE",
	}
	Family_Safi_value = map[string]int32{
		"SAFI_UNSPECIFIED":              0,
		"SAFI_UNICAST":                  1,
		"SAFI_MULTICAST":                2,
		"SAFI_MPLS_LABEL":               4,
		"SAFI_ENCAPSULATION":            7,
		"SAFI_VPLS":                     65,
		"SAFI_EVPN":                     70,
		"SAFI_LS":                       71,
		"SAFI_SR_POLICY":                73,
		"SAFI_MUP":                      85,
		"SAFI_MPLS_VPN":                 128,
		"SAFI_MPLS_VPN_MULTICAST":       129,
		"SAFI_ROUTE_TARGET_CONSTRAINTS": 132,
		"SAFI_FLOW_SPEC_UNICAST":        133,
		"SAFI_FLOW_SPEC_VPN":            134,
		"SAFI_KEY_VALUE":                241,
	}
)

func (x Family_Safi) Enum() *Family_Safi {
	p := new(Family_Safi)
	*p = x
	return p
}

func (x Family_Safi) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Family_Safi) Descriptor() protoreflect.EnumDescriptor {
	return file_api_common_proto_enumTypes[1].Descriptor()
}

func (Family_Safi) Type() protoreflect.EnumType {
	return &file_api_common_proto_enumTypes[1]
}

func (x Family_Safi) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Family_Safi.Descriptor instead.
func (Family_Safi) EnumDescriptor() ([]byte, []int) {
	return file_api_common_proto_rawDescGZIP(), []int{0, 1}
}

type Family struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Afi           Family_Afi             `protobuf:"varint,1,opt,name=afi,proto3,enum=api.Family_Afi" json:"afi,omitempty"`
	Safi          Family_Safi            `protobuf:"varint,2,opt,name=safi,proto3,enum=api.Family_Safi" json:"safi,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Family) Reset() {
	*x = Family{}
	mi := &file_api_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Family) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Family) ProtoMessage() {}

func (x *Family) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Family.ProtoReflect.Descriptor instead.
func (*Family) Descriptor() ([]byte, []int) {
	return file_api_common_proto_rawDescGZIP(), []int{0}
}

func (x *Family) GetAfi() Family_Afi {
	if x != nil {
		return x.Afi
	}
	return Family_AFI_UNSPECIFIED
}

func (x *Family) GetSafi() Family_Safi {
	if x != nil {
		return x.Safi
	}
	return Family_SAFI_UNSPECIFIED
}

type RouteDistinguisherTwoOctetASN struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Admin         uint32                 `protobuf:"varint,1,opt,name=admin,proto3" json:"admin,omitempty"`
	Assigned      uint32                 `protobuf:"varint,2,opt,name=assigned,proto3" json:"assigned,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RouteDistinguisherTwoOctetASN) Reset() {
	*x = RouteDistinguisherTwoOctetASN{}
	mi := &file_api_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteDistinguisherTwoOctetASN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteDistinguisherTwoOctetASN) ProtoMessage() {}

func (x *RouteDistinguisherTwoOctetASN) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteDistinguisherTwoOctetASN.ProtoReflect.Descriptor instead.
func (*RouteDistinguisherTwoOctetASN) Descriptor() ([]byte, []int) {
	return file_api_common_proto_rawDescGZIP(), []int{1}
}

func (x *RouteDistinguisherTwoOctetASN) GetAdmin() uint32 {
	if x != nil {
		return x.Admin
	}
	return 0
}

func (x *RouteDistinguisherTwoOctetASN) GetAssigned() uint32 {
	if x != nil {
		return x.Assigned
	}
	return 0
}

type RouteDistinguisherIPAddress struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Admin         string                 `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	Assigned      uint32                 `protobuf:"varint,2,opt,name=assigned,proto3" json:"assigned,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RouteDistinguisherIPAddress) Reset() {
	*x = RouteDistinguisherIPAddress{}
	mi := &file_api_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteDistinguisherIPAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteDistinguisherIPAddress) ProtoMessage() {}

func (x *RouteDistinguisherIPAddress) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteDistinguisherIPAddress.ProtoReflect.Descriptor instead.
func (*RouteDistinguisherIPAddress) Descriptor() ([]byte, []int) {
	return file_api_common_proto_rawDescGZIP(), []int{2}
}

func (x *RouteDistinguisherIPAddress) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *RouteDistinguisherIPAddress) GetAssigned() uint32 {
	if x != nil {
		return x.Assigned
	}
	return 0
}

type RouteDistinguisherFourOctetASN struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Admin         uint32                 `protobuf:"varint,1,opt,name=admin,proto3" json:"admin,omitempty"`
	Assigned      uint32                 `protobuf:"varint,2,opt,name=assigned,proto3" json:"assigned,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RouteDistinguisherFourOctetASN) Reset() {
	*x = RouteDistinguisherFourOctetASN{}
	mi := &file_api_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteDistinguisherFourOctetASN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteDistinguisherFourOctetASN) ProtoMessage() {}

func (x *RouteDistinguisherFourOctetASN) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteDistinguisherFourOctetASN.ProtoReflect.Descriptor instead.
func (*RouteDistinguisherFourOctetASN) Descriptor() ([]byte, []int) {
	return file_api_common_proto_rawDescGZIP(), []int{3}
}

func (x *RouteDistinguisherFourOctetASN) GetAdmin() uint32 {
	if x != nil {
		return x.Admin
	}
	return 0
}

func (x *RouteDistinguisherFourOctetASN) GetAssigned() uint32 {
	if x != nil {
		return x.Assigned
	}
	return 0
}

type RouteDistinguisher struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Rd:
	//
	//	*RouteDistinguisher_TwoOctetAsn
	//	*RouteDistinguisher_IpAddress
	//	*RouteDistinguisher_FourOctetAsn
	Rd            isRouteDistinguisher_Rd `protobuf_oneof:"rd"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RouteDistinguisher) Reset() {
	*x = RouteDistinguisher{}
	mi := &file_api_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteDistinguisher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteDistinguisher) ProtoMessage() {}

func (x *RouteDistinguisher) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteDistinguisher.ProtoReflect.Descriptor instead.
func (*RouteDistinguisher) Descriptor() ([]byte, []int) {
	return file_api_common_proto_rawDescGZIP(), []int{4}
}

func (x *RouteDistinguisher) GetRd() isRouteDistinguisher_Rd {
	if x != nil {
		return x.Rd
	}
	return nil
}

func (x *RouteDistinguisher) GetTwoOctetAsn() *RouteDistinguisherTwoOctetASN {
	if x != nil {
		if x, ok := x.Rd.(*RouteDistinguisher_TwoOctetAsn); ok {
			return x.TwoOctetAsn
		}
	}
	return nil
}

func (x *RouteDistinguisher) GetIpAddress() *RouteDistinguisherIPAddress {
	if x != nil {
		if x, ok := x.Rd.(*RouteDistinguisher_IpAddress); ok {
			return x.IpAddress
		}
	}
	return nil
}

func (x *RouteDistinguisher) GetFourOctetAsn() *RouteDistinguisherFourOctetASN {
	if x != nil {
		if x, ok := x.Rd.(*RouteDistinguisher_FourOctetAsn); ok {
			return x.FourOctetAsn
		}
	}
	return nil
}

type isRouteDistinguisher_Rd interface {
	isRouteDistinguisher_Rd()
}

type RouteDistinguisher_TwoOctetAsn struct {
	TwoOctetAsn *RouteDistinguisherTwoOctetASN `protobuf:"bytes,1,opt,name=two_octet_asn,json=twoOctetAsn,proto3,oneof"`
}

type RouteDistinguisher_IpAddress struct {
	IpAddress *RouteDistinguisherIPAddress `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3,oneof"`
}

type RouteDistinguisher_FourOctetAsn struct {
	FourOctetAsn *RouteDistinguisherFourOctetASN `protobuf:"bytes,3,opt,name=four_octet_asn,json=fourOctetAsn,proto3,oneof"`
}

func (*RouteDistinguisher_TwoOctetAsn) isRouteDistinguisher_Rd() {}

func (*RouteDistinguisher_IpAddress) isRouteDistinguisher_Rd() {}

func (*RouteDistinguisher_FourOctetAsn) isRouteDistinguisher_Rd() {}

var File_api_common_proto protoreflect.FileDescriptor

const file_api_common_proto_rawDesc = "" +
	"\n" +
	"\x10api/common.proto\x12\x03api\"\x95\x04\n" +
	"\x06Family\x12!\n" +
	"\x03afi\x18\x01 \x01(\x0e2\x0f.api.Family.AfiR\x03afi\x12$\n" +
	"\x04safi\x18\x02 \x01(\x0e2\x10.api.Family.SafiR\x04safi\"b\n" +
	"\x03Afi\x12\x13\n" +
	"\x0fAFI_UNSPECIFIED\x10\x00\x12\n" +
	"\n" +
	"\x06AFI_IP\x10\x01\x12\v\n" +
	"\aAFI_IP6\x10\x02\x12\r\n" +
	"\tAFI_L2VPN\x10\x19\x12\f\n" +
	"\x06AFI_LS\x10\x84\x80\x01\x12\x10\n" +
	"\n" +
	"AFI_OPAQUE\x10\x8d\x80\x01\"\xdd\x02\n" +
	"\x04Safi\x12\x14\n" +
	"\x10SAFI_UNSPECIFIED\x10\x00\x12\x10\n" +
	"\fSAFI_UNICAST\x10\x01\x12\x12\n" +
	"\x0eSAFI_MULTICAST\x10\x02\x12\x13\n" +
	"\x0fSAFI_MPLS_LABEL\x10\x04\x12\x16\n" +
	"\x12SAFI_ENCAPSULATION\x10\a\x12\r\n" +
	"\tSAFI_VPLS\x10A\x12\r\n" +
	"\tSAFI_EVPN\x10F\x12\v\n" +
	"\aSAFI_LS\x10G\x12\x12\n" +
	"\x0eSAFI_SR_POLICY\x10I\x12\f\n" +
	"\bSAFI_MUP\x10U\x12\x12\n" +
	"\rSAFI_MPLS_VPN\x10\x80\x01\x12\x1c\n" +
	"\x17SAFI_MPLS_VPN_MULTICAST\x10\x81\x01\x12\"\n" +
	"\x1dSAFI_ROUTE_TARGET_CONSTRAINTS\x10\x84\x01\x12\x1b\n" +
	"\x16SAFI_FLOW_SPEC_UNICAST\x10\x85\x01\x12\x17\n" +
	"\x12SAFI_FLOW_SPEC_VPN\x10\x86\x01\x12\x13\n" +
	"\x0eSAFI_KEY_VALUE\x10\xf1\x01\"Q\n" +
	"\x1dRouteDistinguisherTwoOctetASN\x12\x14\n" +
	"\x05admin\x18\x01 \x01(\rR\x05admin\x12\x1a\n" +
	"\bassigned\x18\x02 \x01(\rR\bassigned\"O\n" +
	"\x1bRouteDistinguisherIPAddress\x12\x14\n" +
	"\x05admin\x18\x01 \x01(\tR\x05admin\x12\x1a\n" +
	"\bassigned\x18\x02 \x01(\rR\bassigned\"R\n" +
	"\x1eRouteDistinguisherFourOctetASN\x12\x14\n" +
	"\x05admin\x18\x01 \x01(\rR\x05admin\x12\x1a\n" +
	"\bassigned\x18\x02 \x01(\rR\bassigned\"\xf4\x01\n" +
	"\x12RouteDistinguisher\x12H\n" +
	"\rtwo_octet_asn\x18\x01 \x01(\v2\".api.RouteDistinguisherTwoOctetASNH\x00R\vtwoOctetAsn\x12A\n" +
	"\n" +
	"ip_address\x18\x02 \x01(\v2 .api.RouteDistinguisherIPAddressH\x00R\tipAddress\x12K\n" +
	"\x0efour_octet_asn\x18\x03 \x01(\v2#.api.RouteDistinguisherFourOctetASNH\x00R\ffourOctetAsnB\x04\n" +
	"\x02rdB\"Z github.com/osrg/gobgp/v4/api;apib\x06proto3"

var (
	file_api_common_proto_rawDescOnce sync.Once
	file_api_common_proto_rawDescData []byte
)

func file_api_common_proto_rawDescGZIP() []byte {
	file_api_common_proto_rawDescOnce.Do(func() {
		file_api_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_common_proto_rawDesc), len(file_api_common_proto_rawDesc)))
	})
	return file_api_common_proto_rawDescData
}

var file_api_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_common_proto_goTypes = []any{
	(Family_Afi)(0),                        // 0: api.Family.Afi
	(Family_Safi)(0),                       // 1: api.Family.Safi
	(*Family)(nil),                         // 2: api.Family
	(*RouteDistinguisherTwoOctetASN)(nil),  // 3: api.RouteDistinguisherTwoOctetASN
	(*RouteDistinguisherIPAddress)(nil),    // 4: api.RouteDistinguisherIPAddress
	(*RouteDistinguisherFourOctetASN)(nil), // 5: api.RouteDistinguisherFourOctetASN
	(*RouteDistinguisher)(nil),             // 6: api.RouteDistinguisher
}
var file_api_common_proto_depIdxs = []int32{
	0, // 0: api.Family.afi:type_name -> api.Family.Afi
	1, // 1: api.Family.safi:type_name -> api.Family.Safi
	3, // 2: api.RouteDistinguisher.two_octet_asn:type_name -> api.RouteDistinguisherTwoOctetASN
	4, // 3: api.RouteDistinguisher.ip_address:type_name -> api.RouteDistinguisherIPAddress
	5, // 4: api.RouteDistinguisher.four_octet_asn:type_name -> api.RouteDistinguisherFourOctetASN
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_common_proto_init() }
func file_api_common_proto_init() {
	if File_api_common_proto != nil {
		return
	}
	file_api_common_proto_msgTypes[4].OneofWrappers = []any{
		(*RouteDistinguisher_TwoOctetAsn)(nil),
		(*RouteDistinguisher_IpAddress)(nil),
		(*RouteDistinguisher_FourOctetAsn)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_common_proto_rawDesc), len(file_api_common_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_proto_goTypes,
		DependencyIndexes: file_api_common_proto_depIdxs,
		EnumInfos:         file_api_common_proto_enumTypes,
		MessageInfos:      file_api_common_proto_msgTypes,
	}.Build()
	File_api_common_proto = out.File
	file_api_common_proto_goTypes = nil
	file_api_common_proto_depIdxs = nil
}
