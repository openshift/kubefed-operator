// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datalabeling/v1beta1/instruction.proto

package datalabeling // import "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Instruction of how to perform the labeling task for human operators.
// Currently two types of instruction are supported - CSV file and PDF.
// One of the two types instruction must be provided.
// CSV file is only supported for image classification task. Instructions for
// other task should be provided as PDF.
// For image classification, CSV and PDF can be provided at the same time.
type Instruction struct {
	// Output only. Instruction resource name, format:
	// projects/{project_id}/instructions/{instruction_id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the instruction. Maximum of 64 characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. User-provided description of the instruction.
	// The description can be up to 10000 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Creation time of instruction.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Last update time of instruction.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Required. The data type of this instruction.
	DataType DataType `protobuf:"varint,6,opt,name=data_type,json=dataType,proto3,enum=google.cloud.datalabeling.v1beta1.DataType" json:"data_type,omitempty"`
	// One of CSV and PDF instruction is required.
	// Instruction from a csv file, such as for classification task.
	// Csv file should have exact two columns, in the format of:
	// The first column is labeled data, such as image reference, text.
	// The second column is comma separated labels associated with data.
	CsvInstruction *CsvInstruction `protobuf:"bytes,7,opt,name=csv_instruction,json=csvInstruction,proto3" json:"csv_instruction,omitempty"`
	// One of CSV and PDF instruction is required.
	// Instruction from a PDF doc. The PDF doc should be in GCS bucket.
	PdfInstruction       *PdfInstruction `protobuf:"bytes,9,opt,name=pdf_instruction,json=pdfInstruction,proto3" json:"pdf_instruction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Instruction) Reset()         { *m = Instruction{} }
func (m *Instruction) String() string { return proto.CompactTextString(m) }
func (*Instruction) ProtoMessage()    {}
func (*Instruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_instruction_1a000bce9896ac69, []int{0}
}
func (m *Instruction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instruction.Unmarshal(m, b)
}
func (m *Instruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instruction.Marshal(b, m, deterministic)
}
func (dst *Instruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instruction.Merge(dst, src)
}
func (m *Instruction) XXX_Size() int {
	return xxx_messageInfo_Instruction.Size(m)
}
func (m *Instruction) XXX_DiscardUnknown() {
	xxx_messageInfo_Instruction.DiscardUnknown(m)
}

var xxx_messageInfo_Instruction proto.InternalMessageInfo

func (m *Instruction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instruction) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Instruction) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Instruction) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Instruction) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Instruction) GetDataType() DataType {
	if m != nil {
		return m.DataType
	}
	return DataType_DATA_TYPE_UNSPECIFIED
}

func (m *Instruction) GetCsvInstruction() *CsvInstruction {
	if m != nil {
		return m.CsvInstruction
	}
	return nil
}

func (m *Instruction) GetPdfInstruction() *PdfInstruction {
	if m != nil {
		return m.PdfInstruction
	}
	return nil
}

// Instruction from a CSV file.
type CsvInstruction struct {
	// CSV file for the instruction. Only gcs path is allowed.
	GcsFileUri           string   `protobuf:"bytes,1,opt,name=gcs_file_uri,json=gcsFileUri,proto3" json:"gcs_file_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsvInstruction) Reset()         { *m = CsvInstruction{} }
func (m *CsvInstruction) String() string { return proto.CompactTextString(m) }
func (*CsvInstruction) ProtoMessage()    {}
func (*CsvInstruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_instruction_1a000bce9896ac69, []int{1}
}
func (m *CsvInstruction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsvInstruction.Unmarshal(m, b)
}
func (m *CsvInstruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsvInstruction.Marshal(b, m, deterministic)
}
func (dst *CsvInstruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsvInstruction.Merge(dst, src)
}
func (m *CsvInstruction) XXX_Size() int {
	return xxx_messageInfo_CsvInstruction.Size(m)
}
func (m *CsvInstruction) XXX_DiscardUnknown() {
	xxx_messageInfo_CsvInstruction.DiscardUnknown(m)
}

var xxx_messageInfo_CsvInstruction proto.InternalMessageInfo

func (m *CsvInstruction) GetGcsFileUri() string {
	if m != nil {
		return m.GcsFileUri
	}
	return ""
}

// Instruction from a PDF file.
type PdfInstruction struct {
	// PDF file for the instruction. Only gcs path is allowed.
	GcsFileUri           string   `protobuf:"bytes,1,opt,name=gcs_file_uri,json=gcsFileUri,proto3" json:"gcs_file_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PdfInstruction) Reset()         { *m = PdfInstruction{} }
func (m *PdfInstruction) String() string { return proto.CompactTextString(m) }
func (*PdfInstruction) ProtoMessage()    {}
func (*PdfInstruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_instruction_1a000bce9896ac69, []int{2}
}
func (m *PdfInstruction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PdfInstruction.Unmarshal(m, b)
}
func (m *PdfInstruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PdfInstruction.Marshal(b, m, deterministic)
}
func (dst *PdfInstruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PdfInstruction.Merge(dst, src)
}
func (m *PdfInstruction) XXX_Size() int {
	return xxx_messageInfo_PdfInstruction.Size(m)
}
func (m *PdfInstruction) XXX_DiscardUnknown() {
	xxx_messageInfo_PdfInstruction.DiscardUnknown(m)
}

var xxx_messageInfo_PdfInstruction proto.InternalMessageInfo

func (m *PdfInstruction) GetGcsFileUri() string {
	if m != nil {
		return m.GcsFileUri
	}
	return ""
}

func init() {
	proto.RegisterType((*Instruction)(nil), "google.cloud.datalabeling.v1beta1.Instruction")
	proto.RegisterType((*CsvInstruction)(nil), "google.cloud.datalabeling.v1beta1.CsvInstruction")
	proto.RegisterType((*PdfInstruction)(nil), "google.cloud.datalabeling.v1beta1.PdfInstruction")
}

func init() {
	proto.RegisterFile("google/cloud/datalabeling/v1beta1/instruction.proto", fileDescriptor_instruction_1a000bce9896ac69)
}

var fileDescriptor_instruction_1a000bce9896ac69 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8b, 0xdb, 0x30,
	0x10, 0x85, 0x71, 0x77, 0xbb, 0x6d, 0xe4, 0x25, 0x05, 0x9f, 0x4c, 0x28, 0xd4, 0xbb, 0x50, 0x08,
	0x14, 0x24, 0x92, 0x3d, 0xee, 0xad, 0x2d, 0xa5, 0x3d, 0xb4, 0x04, 0x93, 0x5e, 0x72, 0x31, 0xb2,
	0x24, 0x0b, 0x81, 0x2c, 0x09, 0x4b, 0x0e, 0xcd, 0xdf, 0xed, 0x2f, 0x29, 0x92, 0x95, 0xc6, 0xbe,
	0xac, 0x73, 0xf3, 0xcc, 0xbc, 0xef, 0xf9, 0x31, 0x83, 0xc0, 0x13, 0xd7, 0x9a, 0x4b, 0x86, 0x88,
	0xd4, 0x3d, 0x45, 0x14, 0x3b, 0x2c, 0x71, 0xcd, 0xa4, 0x50, 0x1c, 0x1d, 0x37, 0x35, 0x73, 0x78,
	0x83, 0x84, 0xb2, 0xae, 0xeb, 0x89, 0x13, 0x5a, 0x41, 0xd3, 0x69, 0xa7, 0xb3, 0x87, 0x01, 0x82,
	0x01, 0x82, 0x63, 0x08, 0x46, 0x68, 0x85, 0xe6, 0x7d, 0x7d, 0xd3, 0x32, 0x37, 0x78, 0xae, 0x3e,
	0x44, 0x20, 0x54, 0x75, 0xdf, 0x20, 0x27, 0x5a, 0x66, 0x1d, 0x6e, 0x4d, 0x14, 0xbc, 0x8f, 0x02,
	0x6c, 0x04, 0xc2, 0x4a, 0x69, 0x87, 0x7d, 0x22, 0x3b, 0x4c, 0x1f, 0xff, 0xde, 0x80, 0xf4, 0xc7,
	0x25, 0x68, 0x96, 0x81, 0x5b, 0x85, 0x5b, 0x96, 0x27, 0x45, 0xb2, 0x5e, 0x94, 0xe1, 0x3b, 0x7b,
	0x00, 0xf7, 0x54, 0x58, 0x23, 0xf1, 0xa9, 0x0a, 0xb3, 0x57, 0x61, 0x96, 0xc6, 0xde, 0x2f, 0x2f,
	0x29, 0x40, 0x4a, 0x99, 0x25, 0x9d, 0x30, 0xde, 0x25, 0xbf, 0x89, 0x8a, 0x4b, 0x2b, 0x7b, 0x06,
	0x29, 0xe9, 0x18, 0x76, 0xac, 0xf2, 0x01, 0xf3, 0xdb, 0x22, 0x59, 0xa7, 0xdb, 0x15, 0x8c, 0x1b,
	0x39, 0xa7, 0x87, 0xfb, 0x73, 0xfa, 0x12, 0x0c, 0x72, 0xdf, 0xf0, 0x70, 0x6f, 0xe8, 0x7f, 0xf8,
	0xf5, 0x3c, 0x3c, 0xc8, 0x03, 0xfc, 0x1d, 0x2c, 0xfc, 0xca, 0x2a, 0x77, 0x32, 0x2c, 0xbf, 0x2b,
	0x92, 0xf5, 0x72, 0xfb, 0x09, 0xce, 0x5e, 0x02, 0x7e, 0xc5, 0x0e, 0xef, 0x4f, 0x86, 0x95, 0x6f,
	0x69, 0xfc, 0xca, 0x0e, 0xe0, 0x1d, 0xb1, 0xc7, 0x6a, 0x74, 0xd8, 0xfc, 0x4d, 0x88, 0xb2, 0xb9,
	0xc2, 0xef, 0x8b, 0x3d, 0x8e, 0x16, 0x5d, 0x2e, 0xc9, 0xa4, 0xf6, 0xde, 0x86, 0x36, 0x13, 0xef,
	0xc5, 0xd5, 0xde, 0x3b, 0xda, 0x4c, 0xbc, 0xcd, 0xa4, 0x7e, 0xdc, 0x82, 0xe5, 0xf4, 0xef, 0x59,
	0x01, 0xee, 0x39, 0xb1, 0x55, 0x23, 0x24, 0xab, 0xfa, 0x4e, 0xc4, 0x73, 0x03, 0x4e, 0xec, 0x37,
	0x21, 0xd9, 0xef, 0x4e, 0x78, 0x66, 0xea, 0x3a, 0xcf, 0x7c, 0xfe, 0x03, 0x3e, 0x12, 0xdd, 0xce,
	0xe7, 0xdd, 0x25, 0x87, 0x9f, 0x51, 0xc4, 0xb5, 0xc4, 0x8a, 0x43, 0xdd, 0x71, 0xc4, 0x99, 0x0a,
	0x97, 0x8c, 0x4f, 0x00, 0x1b, 0x61, 0x5f, 0x78, 0x06, 0xcf, 0xe3, 0x66, 0x7d, 0x17, 0xc8, 0xa7,
	0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x89, 0x16, 0xfd, 0x97, 0x03, 0x00, 0x00,
}
