// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: protos/movieapp.proto

package __

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_movieapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_movieapp_proto_msgTypes[0]
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
	return file_protos_movieapp_proto_rawDescGZIP(), []int{0}
}

type MovieInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Isbn     string    `protobuf:"bytes,2,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Title    string    `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Director *Director `protobuf:"bytes,4,opt,name=director,proto3" json:"director,omitempty"`
}

func (x *MovieInfo) Reset() {
	*x = MovieInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_movieapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieInfo) ProtoMessage() {}

func (x *MovieInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_movieapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieInfo.ProtoReflect.Descriptor instead.
func (*MovieInfo) Descriptor() ([]byte, []int) {
	return file_protos_movieapp_proto_rawDescGZIP(), []int{1}
}

func (x *MovieInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MovieInfo) GetIsbn() string {
	if x != nil {
		return x.Isbn
	}
	return ""
}

func (x *MovieInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieInfo) GetDirector() *Director {
	if x != nil {
		return x.Director
	}
	return nil
}

type Director struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
}

func (x *Director) Reset() {
	*x = Director{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_movieapp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Director) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Director) ProtoMessage() {}

func (x *Director) ProtoReflect() protoreflect.Message {
	mi := &file_protos_movieapp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Director.ProtoReflect.Descriptor instead.
func (*Director) Descriptor() ([]byte, []int) {
	return file_protos_movieapp_proto_rawDescGZIP(), []int{2}
}

func (x *Director) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Director) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password      string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	WatchedMovies string `protobuf:"bytes,5,opt,name=watched_movies,json=watchedMovies,proto3" json:"watched_movies,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_movieapp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_movieapp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_protos_movieapp_proto_rawDescGZIP(), []int{3}
}

func (x *UserInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserInfo) GetWatchedMovies() string {
	if x != nil {
		return x.WatchedMovies
	}
	return ""
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_movieapp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_protos_movieapp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_protos_movieapp_proto_rawDescGZIP(), []int{4}
}

func (x *Id) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_movieapp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_protos_movieapp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_protos_movieapp_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_protos_movieapp_proto protoreflect.FileDescriptor

var file_protos_movieapp_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x61, 0x70,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x71, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2a, 0x0a,
	0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x44, 0x0a, 0x08, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x87, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x22, 0x1a, 0x0a, 0x02, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xad, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x30, 0x01, 0x12, 0x23, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x08,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x08, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12,
	0x0b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x30, 0x01, 0x12,
	0x25, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x08, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x08, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x08, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_movieapp_proto_rawDescOnce sync.Once
	file_protos_movieapp_proto_rawDescData = file_protos_movieapp_proto_rawDesc
)

func file_protos_movieapp_proto_rawDescGZIP() []byte {
	file_protos_movieapp_proto_rawDescOnce.Do(func() {
		file_protos_movieapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_movieapp_proto_rawDescData)
	})
	return file_protos_movieapp_proto_rawDescData
}

var file_protos_movieapp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_movieapp_proto_goTypes = []interface{}{
	(*Empty)(nil),     // 0: test.Empty
	(*MovieInfo)(nil), // 1: test.MovieInfo
	(*Director)(nil),  // 2: test.Director
	(*UserInfo)(nil),  // 3: test.UserInfo
	(*Id)(nil),        // 4: test.Id
	(*Status)(nil),    // 5: test.Status
}
var file_protos_movieapp_proto_depIdxs = []int32{
	2,  // 0: test.MovieInfo.director:type_name -> test.Director
	0,  // 1: test.ComputeService.GetUsers:input_type -> test.Empty
	4,  // 2: test.ComputeService.GetUser:input_type -> test.Id
	3,  // 3: test.ComputeService.CreateUser:input_type -> test.UserInfo
	3,  // 4: test.ComputeService.UpdateUser:input_type -> test.UserInfo
	4,  // 5: test.ComputeService.DeleteUser:input_type -> test.Id
	0,  // 6: test.ComputeService.GetMovies:input_type -> test.Empty
	4,  // 7: test.ComputeService.GetMovie:input_type -> test.Id
	1,  // 8: test.ComputeService.CreateMovie:input_type -> test.MovieInfo
	1,  // 9: test.ComputeService.UpdateMovie:input_type -> test.MovieInfo
	4,  // 10: test.ComputeService.DeleteMovie:input_type -> test.Id
	3,  // 11: test.ComputeService.GetUsers:output_type -> test.UserInfo
	3,  // 12: test.ComputeService.GetUser:output_type -> test.UserInfo
	4,  // 13: test.ComputeService.CreateUser:output_type -> test.Id
	5,  // 14: test.ComputeService.UpdateUser:output_type -> test.Status
	5,  // 15: test.ComputeService.DeleteUser:output_type -> test.Status
	1,  // 16: test.ComputeService.GetMovies:output_type -> test.MovieInfo
	1,  // 17: test.ComputeService.GetMovie:output_type -> test.MovieInfo
	4,  // 18: test.ComputeService.CreateMovie:output_type -> test.Id
	5,  // 19: test.ComputeService.UpdateMovie:output_type -> test.Status
	5,  // 20: test.ComputeService.DeleteMovie:output_type -> test.Status
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_protos_movieapp_proto_init() }
func file_protos_movieapp_proto_init() {
	if File_protos_movieapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_movieapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_movieapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieInfo); i {
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
		file_protos_movieapp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Director); i {
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
		file_protos_movieapp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_protos_movieapp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_protos_movieapp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_protos_movieapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_movieapp_proto_goTypes,
		DependencyIndexes: file_protos_movieapp_proto_depIdxs,
		MessageInfos:      file_protos_movieapp_proto_msgTypes,
	}.Build()
	File_protos_movieapp_proto = out.File
	file_protos_movieapp_proto_rawDesc = nil
	file_protos_movieapp_proto_goTypes = nil
	file_protos_movieapp_proto_depIdxs = nil
}