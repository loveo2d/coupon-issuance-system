// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/campaign/create.proto

package campaign

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type CreateCampaignRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	CouponRemains int32                  `protobuf:"varint,2,opt,name=coupon_remains,json=couponRemains,proto3" json:"coupon_remains,omitempty"`
	BeginAt       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=begin_at,json=beginAt,proto3" json:"begin_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCampaignRequest) Reset() {
	*x = CreateCampaignRequest{}
	mi := &file_proto_campaign_create_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCampaignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCampaignRequest) ProtoMessage() {}

func (x *CreateCampaignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_campaign_create_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCampaignRequest.ProtoReflect.Descriptor instead.
func (*CreateCampaignRequest) Descriptor() ([]byte, []int) {
	return file_proto_campaign_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCampaignRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateCampaignRequest) GetCouponRemains() int32 {
	if x != nil {
		return x.CouponRemains
	}
	return 0
}

func (x *CreateCampaignRequest) GetBeginAt() *timestamppb.Timestamp {
	if x != nil {
		return x.BeginAt
	}
	return nil
}

type CreateCampaignResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CampaignId    int32                  `protobuf:"varint,1,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CouponRemains int32                  `protobuf:"varint,3,opt,name=coupon_remains,json=couponRemains,proto3" json:"coupon_remains,omitempty"`
	BeginAt       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=begin_at,json=beginAt,proto3" json:"begin_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCampaignResponse) Reset() {
	*x = CreateCampaignResponse{}
	mi := &file_proto_campaign_create_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCampaignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCampaignResponse) ProtoMessage() {}

func (x *CreateCampaignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_campaign_create_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCampaignResponse.ProtoReflect.Descriptor instead.
func (*CreateCampaignResponse) Descriptor() ([]byte, []int) {
	return file_proto_campaign_create_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCampaignResponse) GetCampaignId() int32 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

func (x *CreateCampaignResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateCampaignResponse) GetCouponRemains() int32 {
	if x != nil {
		return x.CouponRemains
	}
	return 0
}

func (x *CreateCampaignResponse) GetBeginAt() *timestamppb.Timestamp {
	if x != nil {
		return x.BeginAt
	}
	return nil
}

var File_proto_campaign_create_proto protoreflect.FileDescriptor

const file_proto_campaign_create_proto_rawDesc = "" +
	"\n" +
	"\x1bproto/campaign/create.proto\x12\frpc_campaign\x1a\x1fgoogle/protobuf/timestamp.proto\"\x8b\x01\n" +
	"\x15CreateCampaignRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12%\n" +
	"\x0ecoupon_remains\x18\x02 \x01(\x05R\rcouponRemains\x125\n" +
	"\bbegin_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\abeginAt\"\xad\x01\n" +
	"\x16CreateCampaignResponse\x12\x1f\n" +
	"\vcampaign_id\x18\x01 \x01(\x05R\n" +
	"campaignId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12%\n" +
	"\x0ecoupon_remains\x18\x03 \x01(\x05R\rcouponRemains\x125\n" +
	"\bbegin_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\abeginAtBEZCgithub.com/loveo2d/CouponIssuanceSystem/internal/api/proto/campaignb\x06proto3"

var (
	file_proto_campaign_create_proto_rawDescOnce sync.Once
	file_proto_campaign_create_proto_rawDescData []byte
)

func file_proto_campaign_create_proto_rawDescGZIP() []byte {
	file_proto_campaign_create_proto_rawDescOnce.Do(func() {
		file_proto_campaign_create_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_campaign_create_proto_rawDesc), len(file_proto_campaign_create_proto_rawDesc)))
	})
	return file_proto_campaign_create_proto_rawDescData
}

var file_proto_campaign_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_campaign_create_proto_goTypes = []any{
	(*CreateCampaignRequest)(nil),  // 0: rpc_campaign.CreateCampaignRequest
	(*CreateCampaignResponse)(nil), // 1: rpc_campaign.CreateCampaignResponse
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_proto_campaign_create_proto_depIdxs = []int32{
	2, // 0: rpc_campaign.CreateCampaignRequest.begin_at:type_name -> google.protobuf.Timestamp
	2, // 1: rpc_campaign.CreateCampaignResponse.begin_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_campaign_create_proto_init() }
func file_proto_campaign_create_proto_init() {
	if File_proto_campaign_create_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_campaign_create_proto_rawDesc), len(file_proto_campaign_create_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_campaign_create_proto_goTypes,
		DependencyIndexes: file_proto_campaign_create_proto_depIdxs,
		MessageInfos:      file_proto_campaign_create_proto_msgTypes,
	}.Build()
	File_proto_campaign_create_proto = out.File
	file_proto_campaign_create_proto_goTypes = nil
	file_proto_campaign_create_proto_depIdxs = nil
}
