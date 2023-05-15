// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: driver.proto

package pb

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

type DriverStatus int32

const (
	DriverStatus_UNKNOWN DriverStatus = 0
	DriverStatus_FREE    DriverStatus = 1
	DriverStatus_BUSY    DriverStatus = 2
)

// Enum value maps for DriverStatus.
var (
	DriverStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "FREE",
		2: "BUSY",
	}
	DriverStatus_value = map[string]int32{
		"UNKNOWN": 0,
		"FREE":    1,
		"BUSY":    2,
	}
)

func (x DriverStatus) Enum() *DriverStatus {
	p := new(DriverStatus)
	*p = x
	return p
}

func (x DriverStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DriverStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_driver_proto_enumTypes[0].Descriptor()
}

func (DriverStatus) Type() protoreflect.EnumType {
	return &file_driver_proto_enumTypes[0]
}

func (x DriverStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DriverStatus.Descriptor instead.
func (DriverStatus) EnumDescriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{0}
}

type LocationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Radius    float64 `protobuf:"fixed64,3,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *LocationMetadata) Reset() {
	*x = LocationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationMetadata) ProtoMessage() {}

func (x *LocationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationMetadata.ProtoReflect.Descriptor instead.
func (*LocationMetadata) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{0}
}

func (x *LocationMetadata) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *LocationMetadata) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *LocationMetadata) GetRadius() float64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

type DriverLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Distance  float64 `protobuf:"fixed64,4,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *DriverLocation) Reset() {
	*x = DriverLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverLocation) ProtoMessage() {}

func (x *DriverLocation) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverLocation.ProtoReflect.Descriptor instead.
func (*DriverLocation) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{1}
}

func (x *DriverLocation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DriverLocation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *DriverLocation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *DriverLocation) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type DriverLocationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*DriverLocation `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *DriverLocationList) Reset() {
	*x = DriverLocationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverLocationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverLocationList) ProtoMessage() {}

func (x *DriverLocationList) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverLocationList.ProtoReflect.Descriptor instead.
func (*DriverLocationList) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{2}
}

func (x *DriverLocationList) GetLocations() []*DriverLocation {
	if x != nil {
		return x.Locations
	}
	return nil
}

type DriverStatusMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status DriverStatus `protobuf:"varint,2,opt,name=status,proto3,enum=DriverStatus" json:"status,omitempty"`
}

func (x *DriverStatusMetadata) Reset() {
	*x = DriverStatusMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverStatusMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverStatusMetadata) ProtoMessage() {}

func (x *DriverStatusMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverStatusMetadata.ProtoReflect.Descriptor instead.
func (*DriverStatusMetadata) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{3}
}

func (x *DriverStatusMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DriverStatusMetadata) GetStatus() DriverStatus {
	if x != nil {
		return x.Status
	}
	return DriverStatus_UNKNOWN
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{4}
}

var File_driver_proto protoreflect.FileDescriptor

var file_driver_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64,
	0x0a, 0x10, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61,
	0x64, 0x69, 0x75, 0x73, 0x22, 0x7a, 0x0a, 0x0e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x43, 0x0a, 0x12, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x51, 0x0a, 0x14, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0d, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x2a, 0x2f, 0x0a, 0x0c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x46, 0x52, 0x45, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x55, 0x53, 0x59,
	0x10, 0x02, 0x32, 0xa5, 0x01, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x34, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x13,
	0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x15, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x15, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a,
	0x0a, 0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x2e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_driver_proto_rawDescOnce sync.Once
	file_driver_proto_rawDescData = file_driver_proto_rawDesc
)

func file_driver_proto_rawDescGZIP() []byte {
	file_driver_proto_rawDescOnce.Do(func() {
		file_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_driver_proto_rawDescData)
	})
	return file_driver_proto_rawDescData
}

var file_driver_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_driver_proto_goTypes = []interface{}{
	(DriverStatus)(0),            // 0: DriverStatus
	(*LocationMetadata)(nil),     // 1: LocationMetadata
	(*DriverLocation)(nil),       // 2: DriverLocation
	(*DriverLocationList)(nil),   // 3: DriverLocationList
	(*DriverStatusMetadata)(nil), // 4: DriverStatusMetadata
	(*Empty)(nil),                // 5: Empty
}
var file_driver_proto_depIdxs = []int32{
	2, // 0: DriverLocationList.locations:type_name -> DriverLocation
	0, // 1: DriverStatusMetadata.status:type_name -> DriverStatus
	1, // 2: Driver.GetClosest:input_type -> LocationMetadata
	4, // 3: Driver.GetStatus:input_type -> DriverStatusMetadata
	4, // 4: Driver.SetStatus:input_type -> DriverStatusMetadata
	3, // 5: Driver.GetClosest:output_type -> DriverLocationList
	4, // 6: Driver.GetStatus:output_type -> DriverStatusMetadata
	5, // 7: Driver.SetStatus:output_type -> Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_driver_proto_init() }
func file_driver_proto_init() {
	if File_driver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_driver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationMetadata); i {
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
		file_driver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverLocation); i {
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
		file_driver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverLocationList); i {
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
		file_driver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverStatusMetadata); i {
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
		file_driver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_driver_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_driver_proto_goTypes,
		DependencyIndexes: file_driver_proto_depIdxs,
		EnumInfos:         file_driver_proto_enumTypes,
		MessageInfos:      file_driver_proto_msgTypes,
	}.Build()
	File_driver_proto = out.File
	file_driver_proto_rawDesc = nil
	file_driver_proto_goTypes = nil
	file_driver_proto_depIdxs = nil
}
