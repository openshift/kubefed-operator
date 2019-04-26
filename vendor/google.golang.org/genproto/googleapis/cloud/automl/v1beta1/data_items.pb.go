// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/data_items.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/golang/protobuf/ptypes/duration"
import _struct "github.com/golang/protobuf/ptypes/struct"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A representation of an image.
// Only images up to 30MB in size are supported.
type Image struct {
	// Input only. The data representing the image.
	// For Predict calls [image_bytes][] must be set, as other options are not
	// currently supported by prediction API. You can read the contents of an
	// uploaded image by using the [content_uri][] field.
	//
	// Types that are valid to be assigned to Data:
	//	*Image_ImageBytes
	//	*Image_InputConfig
	Data isImage_Data `protobuf_oneof:"data"`
	// Output only. HTTP URI to the thumbnail image.
	ThumbnailUri         string   `protobuf:"bytes,4,opt,name=thumbnail_uri,json=thumbnailUri,proto3" json:"thumbnail_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_items_f94343a769df0351, []int{0}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

type isImage_Data interface {
	isImage_Data()
}

type Image_ImageBytes struct {
	ImageBytes []byte `protobuf:"bytes,1,opt,name=image_bytes,json=imageBytes,proto3,oneof"`
}

type Image_InputConfig struct {
	InputConfig *InputConfig `protobuf:"bytes,6,opt,name=input_config,json=inputConfig,proto3,oneof"`
}

func (*Image_ImageBytes) isImage_Data() {}

func (*Image_InputConfig) isImage_Data() {}

func (m *Image) GetData() isImage_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Image) GetImageBytes() []byte {
	if x, ok := m.GetData().(*Image_ImageBytes); ok {
		return x.ImageBytes
	}
	return nil
}

func (m *Image) GetInputConfig() *InputConfig {
	if x, ok := m.GetData().(*Image_InputConfig); ok {
		return x.InputConfig
	}
	return nil
}

func (m *Image) GetThumbnailUri() string {
	if m != nil {
		return m.ThumbnailUri
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Image) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Image_OneofMarshaler, _Image_OneofUnmarshaler, _Image_OneofSizer, []interface{}{
		(*Image_ImageBytes)(nil),
		(*Image_InputConfig)(nil),
	}
}

func _Image_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Image)
	// data
	switch x := m.Data.(type) {
	case *Image_ImageBytes:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.ImageBytes)
	case *Image_InputConfig:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InputConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Image.Data has unexpected type %T", x)
	}
	return nil
}

func _Image_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Image)
	switch tag {
	case 1: // data.image_bytes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Data = &Image_ImageBytes{x}
		return true, err
	case 6: // data.input_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InputConfig)
		err := b.DecodeMessage(msg)
		m.Data = &Image_InputConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Image_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Image)
	// data
	switch x := m.Data.(type) {
	case *Image_ImageBytes:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ImageBytes)))
		n += len(x.ImageBytes)
	case *Image_InputConfig:
		s := proto.Size(x.InputConfig)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A representation of a text snippet.
type TextSnippet struct {
	// Required. The content of the text snippet as a string. Up to 250000
	// characters long.
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// The format of the source text. Currently the only two allowed values are
	// "text/html" and "text/plain". If left blank the format is automatically
	// determined from the type of the uploaded content.
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Output only. HTTP URI where you can download the content.
	ContentUri           string   `protobuf:"bytes,4,opt,name=content_uri,json=contentUri,proto3" json:"content_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSnippet) Reset()         { *m = TextSnippet{} }
func (m *TextSnippet) String() string { return proto.CompactTextString(m) }
func (*TextSnippet) ProtoMessage()    {}
func (*TextSnippet) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_items_f94343a769df0351, []int{1}
}
func (m *TextSnippet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSnippet.Unmarshal(m, b)
}
func (m *TextSnippet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSnippet.Marshal(b, m, deterministic)
}
func (dst *TextSnippet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSnippet.Merge(dst, src)
}
func (m *TextSnippet) XXX_Size() int {
	return xxx_messageInfo_TextSnippet.Size(m)
}
func (m *TextSnippet) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSnippet.DiscardUnknown(m)
}

var xxx_messageInfo_TextSnippet proto.InternalMessageInfo

func (m *TextSnippet) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TextSnippet) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *TextSnippet) GetContentUri() string {
	if m != nil {
		return m.ContentUri
	}
	return ""
}

// A structured text document e.g. a PDF.
type Document struct {
	// An input config specifying the content of the document.
	InputConfig          *DocumentInputConfig `protobuf:"bytes,1,opt,name=input_config,json=inputConfig,proto3" json:"input_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_items_f94343a769df0351, []int{2}
}
func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (dst *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(dst, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetInputConfig() *DocumentInputConfig {
	if m != nil {
		return m.InputConfig
	}
	return nil
}

// A representation of a row in a relational table.
type Row struct {
	// The resource IDs of the column specs describing the columns of the row.
	// If set must contain, but possibly in a different order, all input feature
	//
	// [column_spec_ids][google.cloud.automl.v1beta1.TablesModelMetadata.input_feature_column_specs]
	// of the Model this row is being passed to.
	// Note: The below `values` field must match order of this field, if this
	// field is set.
	ColumnSpecIds []string `protobuf:"bytes,2,rep,name=column_spec_ids,json=columnSpecIds,proto3" json:"column_spec_ids,omitempty"`
	// Required. The values of the row cells, given in the same order as the
	// column_spec_ids, or, if not set, then in the same order as input feature
	//
	// [column_specs][google.cloud.automl.v1beta1.TablesModelMetadata.input_feature_column_specs]
	// of the Model this row is being passed to.
	Values               []*_struct.Value `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_items_f94343a769df0351, []int{3}
}
func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (dst *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(dst, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetColumnSpecIds() []string {
	if m != nil {
		return m.ColumnSpecIds
	}
	return nil
}

func (m *Row) GetValues() []*_struct.Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// Example data used for training or prediction.
type ExamplePayload struct {
	// Required. Input only. The example data.
	//
	// Types that are valid to be assigned to Payload:
	//	*ExamplePayload_Image
	//	*ExamplePayload_TextSnippet
	//	*ExamplePayload_Document
	//	*ExamplePayload_Row
	Payload              isExamplePayload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ExamplePayload) Reset()         { *m = ExamplePayload{} }
func (m *ExamplePayload) String() string { return proto.CompactTextString(m) }
func (*ExamplePayload) ProtoMessage()    {}
func (*ExamplePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_items_f94343a769df0351, []int{4}
}
func (m *ExamplePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExamplePayload.Unmarshal(m, b)
}
func (m *ExamplePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExamplePayload.Marshal(b, m, deterministic)
}
func (dst *ExamplePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExamplePayload.Merge(dst, src)
}
func (m *ExamplePayload) XXX_Size() int {
	return xxx_messageInfo_ExamplePayload.Size(m)
}
func (m *ExamplePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ExamplePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ExamplePayload proto.InternalMessageInfo

type isExamplePayload_Payload interface {
	isExamplePayload_Payload()
}

type ExamplePayload_Image struct {
	Image *Image `protobuf:"bytes,1,opt,name=image,proto3,oneof"`
}

type ExamplePayload_TextSnippet struct {
	TextSnippet *TextSnippet `protobuf:"bytes,2,opt,name=text_snippet,json=textSnippet,proto3,oneof"`
}

type ExamplePayload_Document struct {
	Document *Document `protobuf:"bytes,4,opt,name=document,proto3,oneof"`
}

type ExamplePayload_Row struct {
	Row *Row `protobuf:"bytes,3,opt,name=row,proto3,oneof"`
}

func (*ExamplePayload_Image) isExamplePayload_Payload() {}

func (*ExamplePayload_TextSnippet) isExamplePayload_Payload() {}

func (*ExamplePayload_Document) isExamplePayload_Payload() {}

func (*ExamplePayload_Row) isExamplePayload_Payload() {}

func (m *ExamplePayload) GetPayload() isExamplePayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ExamplePayload) GetImage() *Image {
	if x, ok := m.GetPayload().(*ExamplePayload_Image); ok {
		return x.Image
	}
	return nil
}

func (m *ExamplePayload) GetTextSnippet() *TextSnippet {
	if x, ok := m.GetPayload().(*ExamplePayload_TextSnippet); ok {
		return x.TextSnippet
	}
	return nil
}

func (m *ExamplePayload) GetDocument() *Document {
	if x, ok := m.GetPayload().(*ExamplePayload_Document); ok {
		return x.Document
	}
	return nil
}

func (m *ExamplePayload) GetRow() *Row {
	if x, ok := m.GetPayload().(*ExamplePayload_Row); ok {
		return x.Row
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExamplePayload) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExamplePayload_OneofMarshaler, _ExamplePayload_OneofUnmarshaler, _ExamplePayload_OneofSizer, []interface{}{
		(*ExamplePayload_Image)(nil),
		(*ExamplePayload_TextSnippet)(nil),
		(*ExamplePayload_Document)(nil),
		(*ExamplePayload_Row)(nil),
	}
}

func _ExamplePayload_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExamplePayload)
	// payload
	switch x := m.Payload.(type) {
	case *ExamplePayload_Image:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Image); err != nil {
			return err
		}
	case *ExamplePayload_TextSnippet:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextSnippet); err != nil {
			return err
		}
	case *ExamplePayload_Document:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Document); err != nil {
			return err
		}
	case *ExamplePayload_Row:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Row); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ExamplePayload.Payload has unexpected type %T", x)
	}
	return nil
}

func _ExamplePayload_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExamplePayload)
	switch tag {
	case 1: // payload.image
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Image)
		err := b.DecodeMessage(msg)
		m.Payload = &ExamplePayload_Image{msg}
		return true, err
	case 2: // payload.text_snippet
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TextSnippet)
		err := b.DecodeMessage(msg)
		m.Payload = &ExamplePayload_TextSnippet{msg}
		return true, err
	case 4: // payload.document
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Document)
		err := b.DecodeMessage(msg)
		m.Payload = &ExamplePayload_Document{msg}
		return true, err
	case 3: // payload.row
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Row)
		err := b.DecodeMessage(msg)
		m.Payload = &ExamplePayload_Row{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ExamplePayload_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExamplePayload)
	// payload
	switch x := m.Payload.(type) {
	case *ExamplePayload_Image:
		s := proto.Size(x.Image)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExamplePayload_TextSnippet:
		s := proto.Size(x.TextSnippet)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExamplePayload_Document:
		s := proto.Size(x.Document)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExamplePayload_Row:
		s := proto.Size(x.Row)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Image)(nil), "google.cloud.automl.v1beta1.Image")
	proto.RegisterType((*TextSnippet)(nil), "google.cloud.automl.v1beta1.TextSnippet")
	proto.RegisterType((*Document)(nil), "google.cloud.automl.v1beta1.Document")
	proto.RegisterType((*Row)(nil), "google.cloud.automl.v1beta1.Row")
	proto.RegisterType((*ExamplePayload)(nil), "google.cloud.automl.v1beta1.ExamplePayload")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/data_items.proto", fileDescriptor_data_items_f94343a769df0351)
}

var fileDescriptor_data_items_f94343a769df0351 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0x9a, 0xad, 0x5b, 0x9d, 0x0d, 0xa4, 0x1c, 0x50, 0xd8, 0x26, 0x16, 0xc2, 0x0f, 0xe5,
	0x80, 0x12, 0x56, 0x38, 0xc1, 0x69, 0x2d, 0x88, 0xf6, 0x30, 0x69, 0x72, 0xb7, 0x1d, 0x50, 0x51,
	0xe4, 0x26, 0x5e, 0xb0, 0x94, 0xd8, 0x56, 0x62, 0xaf, 0xed, 0x9d, 0x7f, 0x84, 0x0b, 0x7f, 0x0c,
	0x7f, 0x15, 0xb2, 0xe3, 0x8e, 0x52, 0xa1, 0xc0, 0x2d, 0xef, 0x7d, 0xdf, 0xf7, 0x9e, 0xdf, 0xf7,
	0x29, 0xe0, 0x55, 0xce, 0x58, 0x5e, 0xe0, 0x38, 0x2d, 0x98, 0xcc, 0x62, 0x24, 0x05, 0x2b, 0x8b,
	0xf8, 0xee, 0x6c, 0x8e, 0x05, 0x3a, 0x8b, 0x33, 0x24, 0x50, 0x42, 0x04, 0x2e, 0xeb, 0x88, 0x57,
	0x4c, 0x30, 0xf7, 0xb8, 0x61, 0x47, 0x9a, 0x1d, 0x35, 0xec, 0xc8, 0xb0, 0x8f, 0x4e, 0xcc, 0x28,
	0xc4, 0x49, 0x8c, 0x28, 0x65, 0x02, 0x09, 0xc2, 0xa8, 0x91, 0x1e, 0x3d, 0x6f, 0x5b, 0x44, 0x98,
	0x61, 0x3d, 0x36, 0x2c, 0x5d, 0xcd, 0xe5, 0x6d, 0x8c, 0xe8, 0xca, 0x40, 0x4f, 0xb6, 0xa1, 0x4c,
	0x56, 0x7a, 0x83, 0xc1, 0x4f, 0xb6, 0xf1, 0x5a, 0x54, 0x32, 0x15, 0x0d, 0x1a, 0xfc, 0xb0, 0xc0,
	0xee, 0xa4, 0x44, 0x39, 0x76, 0x9f, 0x02, 0x87, 0xa8, 0x8f, 0x64, 0xbe, 0x12, 0xb8, 0xf6, 0x2c,
	0xdf, 0x0a, 0x0f, 0xc6, 0x1d, 0x08, 0x74, 0x73, 0xa8, 0x7a, 0xee, 0x05, 0x38, 0x20, 0x94, 0x4b,
	0x91, 0xa4, 0x8c, 0xde, 0x92, 0xdc, 0xeb, 0xf9, 0x56, 0xe8, 0x0c, 0xc2, 0xa8, 0xe5, 0xfa, 0x68,
	0xa2, 0x04, 0x23, 0xcd, 0x1f, 0x77, 0xa0, 0x43, 0x7e, 0x97, 0xee, 0x33, 0x70, 0x28, 0xbe, 0xca,
	0x72, 0x4e, 0x11, 0x29, 0x12, 0x59, 0x11, 0x6f, 0xc7, 0xb7, 0xc2, 0x3e, 0x3c, 0xb8, 0x6f, 0x5e,
	0x57, 0x64, 0xd8, 0x03, 0x3b, 0xca, 0xee, 0x00, 0x03, 0xe7, 0x0a, 0x2f, 0xc5, 0x94, 0x12, 0xce,
	0xb1, 0x70, 0x3d, 0xb0, 0x97, 0x32, 0x2a, 0x30, 0x15, 0xfa, 0xa5, 0x7d, 0xb8, 0x2e, 0xdd, 0x63,
	0xd0, 0x2f, 0x49, 0x89, 0x13, 0xb1, 0xe2, 0xd8, 0xeb, 0x6a, 0x6c, 0x5f, 0x35, 0xae, 0x56, 0x1c,
	0xbb, 0xa7, 0xc0, 0x31, 0xbc, 0x8d, 0x85, 0xc0, 0xb4, 0xae, 0x2b, 0x12, 0x24, 0x60, 0xff, 0x03,
	0x4b, 0x65, 0xa9, 0x26, 0x4d, 0xb7, 0xce, 0xb5, 0xf4, 0xb9, 0xaf, 0x5b, 0xcf, 0x5d, 0x8b, 0x37,
	0xce, 0xfe, 0xe3, 0xe8, 0xe0, 0x0b, 0xb0, 0x21, 0x5b, 0xb8, 0x2f, 0xc1, 0xc3, 0x94, 0x15, 0xb2,
	0xa4, 0x49, 0xcd, 0x71, 0x9a, 0x90, 0xac, 0xf6, 0xba, 0xbe, 0x1d, 0xf6, 0xe1, 0x61, 0xd3, 0x9e,
	0x72, 0x9c, 0x4e, 0xb2, 0xda, 0x8d, 0x40, 0xef, 0x0e, 0x15, 0x12, 0xd7, 0x9e, 0xed, 0xdb, 0xa1,
	0x33, 0x78, 0xb4, 0xde, 0xbe, 0x8e, 0x33, 0xba, 0x51, 0x30, 0x34, 0xac, 0xe0, 0x7b, 0x17, 0x3c,
	0xf8, 0xb8, 0x44, 0x25, 0x2f, 0xf0, 0x25, 0x5a, 0x15, 0x0c, 0x65, 0xee, 0x3b, 0xb0, 0xab, 0x33,
	0x34, 0xef, 0x0f, 0xda, 0xe3, 0x52, 0xcc, 0x71, 0x07, 0x36, 0x12, 0x95, 0xb8, 0xc0, 0x4b, 0x91,
	0xd4, 0x8d, 0xed, 0xda, 0xcf, 0x7f, 0x25, 0xbe, 0x11, 0x93, 0x4a, 0x5c, 0x6c, 0xa4, 0x36, 0x02,
	0xfb, 0x99, 0x31, 0x48, 0x7b, 0xef, 0x0c, 0x5e, 0xfc, 0x97, 0x9b, 0xe3, 0x0e, 0xbc, 0x17, 0xba,
	0x6f, 0x81, 0x5d, 0xb1, 0x85, 0x67, 0x6b, 0xbd, 0xdf, 0xaa, 0x87, 0x6c, 0x31, 0xee, 0x40, 0x45,
	0x1f, 0xf6, 0xc1, 0x1e, 0x6f, 0x0c, 0x19, 0x7e, 0xb3, 0xc0, 0x69, 0xca, 0xca, 0x36, 0xe5, 0xa5,
	0xf5, 0xf9, 0xdc, 0xc0, 0x39, 0x2b, 0x10, 0xcd, 0x23, 0x56, 0xe5, 0x71, 0x8e, 0xa9, 0xb6, 0x3d,
	0x6e, 0x20, 0xc4, 0x49, 0xfd, 0xd7, 0xff, 0xf6, 0x7d, 0x53, 0xfe, 0xec, 0x1e, 0x7f, 0xd2, 0xc4,
	0xd9, 0x48, 0x91, 0x66, 0xe7, 0x52, 0xb0, 0x8b, 0x62, 0x76, 0xd3, 0x90, 0xe6, 0x3d, 0x3d, 0xeb,
	0xcd, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0xc9, 0x61, 0x25, 0x6b, 0x04, 0x00, 0x00,
}
