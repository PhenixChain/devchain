// Code generated by protoc-gen-go. DO NOT EDIT.
// source: retrieve.proto

package types

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	types "github.com/PhenixChain/devchain/types"

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

// message for retrieve start
type RetrievePara struct {
	DefaultAddress       string   `protobuf:"bytes,1,opt,name=defaultAddress,proto3" json:"defaultAddress,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime           int64    `protobuf:"varint,3,opt,name=createTime,proto3" json:"createTime,omitempty"`
	PrepareTime          int64    `protobuf:"varint,4,opt,name=prepareTime,proto3" json:"prepareTime,omitempty"`
	DelayPeriod          int64    `protobuf:"varint,5,opt,name=delayPeriod,proto3" json:"delayPeriod,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrievePara) Reset()         { *m = RetrievePara{} }
func (m *RetrievePara) String() string { return proto.CompactTextString(m) }
func (*RetrievePara) ProtoMessage()    {}
func (*RetrievePara) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieve_1666ef19b265e4fc, []int{0}
}
func (m *RetrievePara) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrievePara.Unmarshal(m, b)
}
func (m *RetrievePara) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrievePara.Marshal(b, m, deterministic)
}
func (dst *RetrievePara) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrievePara.Merge(dst, src)
}
func (m *RetrievePara) XXX_Size() int {
	return xxx_messageInfo_RetrievePara.Size(m)
}
func (m *RetrievePara) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrievePara.DiscardUnknown(m)
}

var xxx_messageInfo_RetrievePara proto.InternalMessageInfo

func (m *RetrievePara) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

func (m *RetrievePara) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RetrievePara) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *RetrievePara) GetPrepareTime() int64 {
	if m != nil {
		return m.PrepareTime
	}
	return 0
}

func (m *RetrievePara) GetDelayPeriod() int64 {
	if m != nil {
		return m.DelayPeriod
	}
	return 0
}

type Retrieve struct {
	// used as key
	BackupAddress        string          `protobuf:"bytes,1,opt,name=backupAddress,proto3" json:"backupAddress,omitempty"`
	RetPara              []*RetrievePara `protobuf:"bytes,2,rep,name=retPara,proto3" json:"retPara,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Retrieve) Reset()         { *m = Retrieve{} }
func (m *Retrieve) String() string { return proto.CompactTextString(m) }
func (*Retrieve) ProtoMessage()    {}
func (*Retrieve) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieve_1666ef19b265e4fc, []int{1}
}
func (m *Retrieve) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Retrieve.Unmarshal(m, b)
}
func (m *Retrieve) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Retrieve.Marshal(b, m, deterministic)
}
func (dst *Retrieve) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Retrieve.Merge(dst, src)
}
func (m *Retrieve) XXX_Size() int {
	return xxx_messageInfo_Retrieve.Size(m)
}
func (m *Retrieve) XXX_DiscardUnknown() {
	xxx_messageInfo_Retrieve.DiscardUnknown(m)
}

var xxx_messageInfo_Retrieve proto.InternalMessageInfo

func (m *Retrieve) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *Retrieve) GetRetPara() []*RetrievePara {
	if m != nil {
		return m.RetPara
	}
	return nil
}

type RetrieveAction struct {
	// Types that are valid to be assigned to Value:
	//	*RetrieveAction_Prepare
	//	*RetrieveAction_Perform
	//	*RetrieveAction_Backup
	//	*RetrieveAction_Cancel
	Value                isRetrieveAction_Value `protobuf_oneof:"value"`
	Ty                   int32                  `protobuf:"varint,5,opt,name=ty,proto3" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RetrieveAction) Reset()         { *m = RetrieveAction{} }
func (m *RetrieveAction) String() string { return proto.CompactTextString(m) }
func (*RetrieveAction) ProtoMessage()    {}
func (*RetrieveAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieve_1666ef19b265e4fc, []int{2}
}
func (m *RetrieveAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveAction.Unmarshal(m, b)
}
func (m *RetrieveAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveAction.Marshal(b, m, deterministic)
}
func (dst *RetrieveAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveAction.Merge(dst, src)
}
func (m *RetrieveAction) XXX_Size() int {
	return xxx_messageInfo_RetrieveAction.Size(m)
}
func (m *RetrieveAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveAction.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveAction proto.InternalMessageInfo

type isRetrieveAction_Value interface {
	isRetrieveAction_Value()
}

type RetrieveAction_Prepare struct {
	Prepare *PrepareRetrieve `protobuf:"bytes,1,opt,name=prepare,proto3,oneof"`
}

type RetrieveAction_Perform struct {
	Perform *PerformRetrieve `protobuf:"bytes,2,opt,name=perform,proto3,oneof"`
}

type RetrieveAction_Backup struct {
	Backup *BackupRetrieve `protobuf:"bytes,3,opt,name=backup,proto3,oneof"`
}

type RetrieveAction_Cancel struct {
	Cancel *CancelRetrieve `protobuf:"bytes,4,opt,name=cancel,proto3,oneof"`
}

func (*RetrieveAction_Prepare) isRetrieveAction_Value() {}

func (*RetrieveAction_Perform) isRetrieveAction_Value() {}

func (*RetrieveAction_Backup) isRetrieveAction_Value() {}

func (*RetrieveAction_Cancel) isRetrieveAction_Value() {}

func (m *RetrieveAction) GetValue() isRetrieveAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RetrieveAction) GetPrepare() *PrepareRetrieve {
	if x, ok := m.GetValue().(*RetrieveAction_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *RetrieveAction) GetPerform() *PerformRetrieve {
	if x, ok := m.GetValue().(*RetrieveAction_Perform); ok {
		return x.Perform
	}
	return nil
}

func (m *RetrieveAction) GetBackup() *BackupRetrieve {
	if x, ok := m.GetValue().(*RetrieveAction_Backup); ok {
		return x.Backup
	}
	return nil
}

func (m *RetrieveAction) GetCancel() *CancelRetrieve {
	if x, ok := m.GetValue().(*RetrieveAction_Cancel); ok {
		return x.Cancel
	}
	return nil
}

func (m *RetrieveAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RetrieveAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RetrieveAction_OneofMarshaler, _RetrieveAction_OneofUnmarshaler, _RetrieveAction_OneofSizer, []interface{}{
		(*RetrieveAction_Prepare)(nil),
		(*RetrieveAction_Perform)(nil),
		(*RetrieveAction_Backup)(nil),
		(*RetrieveAction_Cancel)(nil),
	}
}

func _RetrieveAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RetrieveAction)
	// value
	switch x := m.Value.(type) {
	case *RetrieveAction_Prepare:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *RetrieveAction_Perform:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Perform); err != nil {
			return err
		}
	case *RetrieveAction_Backup:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Backup); err != nil {
			return err
		}
	case *RetrieveAction_Cancel:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Cancel); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RetrieveAction.Value has unexpected type %T", x)
	}
	return nil
}

func _RetrieveAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RetrieveAction)
	switch tag {
	case 1: // value.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrepareRetrieve)
		err := b.DecodeMessage(msg)
		m.Value = &RetrieveAction_Prepare{msg}
		return true, err
	case 2: // value.perform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PerformRetrieve)
		err := b.DecodeMessage(msg)
		m.Value = &RetrieveAction_Perform{msg}
		return true, err
	case 3: // value.backup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BackupRetrieve)
		err := b.DecodeMessage(msg)
		m.Value = &RetrieveAction_Backup{msg}
		return true, err
	case 4: // value.cancel
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CancelRetrieve)
		err := b.DecodeMessage(msg)
		m.Value = &RetrieveAction_Cancel{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RetrieveAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RetrieveAction)
	// value
	switch x := m.Value.(type) {
	case *RetrieveAction_Prepare:
		s := proto.Size(x.Prepare)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RetrieveAction_Perform:
		s := proto.Size(x.Perform)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RetrieveAction_Backup:
		s := proto.Size(x.Backup)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RetrieveAction_Cancel:
		s := proto.Size(x.Cancel)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BackupRetrieve struct {
	BackupAddress        string   `protobuf:"bytes,1,opt,name=backupAddress,proto3" json:"backupAddress,omitempty"`
	DefaultAddress       string   `protobuf:"bytes,2,opt,name=defaultAddress,proto3" json:"defaultAddress,omitempty"`
	DelayPeriod          int64    `protobuf:"varint,3,opt,name=delayPeriod,proto3" json:"delayPeriod,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackupRetrieve) Reset()         { *m = BackupRetrieve{} }
func (m *BackupRetrieve) String() string { return proto.CompactTextString(m) }
func (*BackupRetrieve) ProtoMessage()    {}
func (*BackupRetrieve) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieve_1666ef19b265e4fc, []int{3}
}
func (m *BackupRetrieve) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackupRetrieve.Unmarshal(m, b)
}
func (m *BackupRetrieve) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackupRetrieve.Marshal(b, m, deterministic)
}
func (dst *BackupRetrieve) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupRetrieve.Merge(dst, src)
}
func (m *BackupRetrieve) XXX_Size() int {
	return xxx_messageInfo_BackupRetrieve.Size(m)
}
func (m *BackupRetrieve) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupRetrieve.DiscardUnknown(m)
}

var xxx_messageInfo_BackupRetrieve proto.InternalMessageInfo

func (m *BackupRetrieve) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *BackupRetrieve) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

func (m *BackupRetrieve) GetDelayPeriod() int64 {
	if m != nil {
		return m.DelayPeriod
	}
	return 0
}

type PrepareRetrieve struct {
	BackupAddress        string   `protobuf:"bytes,1,opt,name=backupAddress,proto3" json:"backupAddress,omitempty"`
	DefaultAddress       string   `protobuf:"bytes,2,opt,name=defaultAddress,proto3" json:"defaultAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareRetrieve) Reset()         { *m = PrepareRetrieve{} }
func (m *PrepareRetrieve) String() string { return proto.CompactTextString(m) }
func (*PrepareRetrieve) ProtoMessage()    {}
func (*PrepareRetrieve) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieve_1666ef19b265e4fc, []int{4}
}
func (m *PrepareRetrieve) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareRetrieve.Unmarshal(m, b)
}
func (m *PrepareRetrieve) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareRetrieve.Marshal(b, m, deterministic)
}
func (dst *PrepareRetrieve) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareRetrieve.Merge(dst, src)
}
func (m *PrepareRetrieve) XXX_Size() int {
	return xxx_messageInfo_PrepareRetrieve.Size(m)
}
func (m *PrepareRetrieve) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareRetrieve.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareRetrieve proto.InternalMessageInfo

func (m *PrepareRetrieve) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *PrepareRetrieve) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

type AssetSymbol struct {
	Exec                 string   `protobuf:"bytes,1,opt,name=exec,proto3" json:"exec,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetSymbol) Reset()         { *m = AssetSymbol{} }
func (m *AssetSymbol) String() string { return proto.CompactTextString(m) }
func (*AssetSymbol) ProtoMessage()    {}
func (*AssetSymbol) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieve_1666ef19b265e4fc, []int{5}
}
func (m *AssetSymbol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetSymbol.Unmarshal(m, b)
}
func (m *AssetSymbol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetSymbol.Marshal(b, m, deterministic)
}
func (dst *AssetSymbol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetSymbol.Merge(dst, src)
}
func (m *AssetSymbol) XXX_Size() int {
	return xxx_messageInfo_AssetSymbol.Size(m)
}
func (m *AssetSymbol) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetSymbol.DiscardUnknown(m)
}

var xxx_messageInfo_AssetSymbol proto.InternalMessageInfo

func (m *AssetSymbol) GetExec() string {
	if m != nil {
		return m.Exec
	}
	return ""
}

func (m *AssetSymbol) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type PerformRetrieve struct {
	BackupAddress        string         `protobuf:"bytes,1,opt,name=backupAddress,proto3" json:"backupAddress,omitempty"`
	DefaultAddress       string         `protobuf:"bytes,2,opt,name=defaultAddress,proto3" json:"defaultAddress,omitempty"`
	Assets               []*AssetSymbol `protobuf:"bytes,3,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PerformRetrieve) Reset()         { *m = PerformRetrieve{} }
func (m *PerformRetrieve) String() string { return proto.CompactTextString(m) }
func (*PerformRetrieve) ProtoMessage()    {}
func (*PerformRetrieve) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieve_1666ef19b265e4fc, []int{6}
}
func (m *PerformRetrieve) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerformRetrieve.Unmarshal(m, b)
}
func (m *PerformRetrieve) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerformRetrieve.Marshal(b, m, deterministic)
}
func (dst *PerformRetrieve) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerformRetrieve.Merge(dst, src)
}
func (m *PerformRetrieve) XXX_Size() int {
	return xxx_messageInfo_PerformRetrieve.Size(m)
}
func (m *PerformRetrieve) XXX_DiscardUnknown() {
	xxx_messageInfo_PerformRetrieve.DiscardUnknown(m)
}

var xxx_messageInfo_PerformRetrieve proto.InternalMessageInfo

func (m *PerformRetrieve) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *PerformRetrieve) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

func (m *PerformRetrieve) GetAssets() []*AssetSymbol {
	if m != nil {
		return m.Assets
	}
	return nil
}

type CancelRetrieve struct {
	BackupAddress        string   `protobuf:"bytes,1,opt,name=backupAddress,proto3" json:"backupAddress,omitempty"`
	DefaultAddress       string   `protobuf:"bytes,2,opt,name=defaultAddress,proto3" json:"defaultAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelRetrieve) Reset()         { *m = CancelRetrieve{} }
func (m *CancelRetrieve) String() string { return proto.CompactTextString(m) }
func (*CancelRetrieve) ProtoMessage()    {}
func (*CancelRetrieve) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieve_1666ef19b265e4fc, []int{7}
}
func (m *CancelRetrieve) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelRetrieve.Unmarshal(m, b)
}
func (m *CancelRetrieve) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelRetrieve.Marshal(b, m, deterministic)
}
func (dst *CancelRetrieve) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelRetrieve.Merge(dst, src)
}
func (m *CancelRetrieve) XXX_Size() int {
	return xxx_messageInfo_CancelRetrieve.Size(m)
}
func (m *CancelRetrieve) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelRetrieve.DiscardUnknown(m)
}

var xxx_messageInfo_CancelRetrieve proto.InternalMessageInfo

func (m *CancelRetrieve) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *CancelRetrieve) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

type ReqRetrieveInfo struct {
	BackupAddress        string   `protobuf:"bytes,1,opt,name=backupAddress,proto3" json:"backupAddress,omitempty"`
	DefaultAddress       string   `protobuf:"bytes,2,opt,name=defaultAddress,proto3" json:"defaultAddress,omitempty"`
	AssetExec            string   `protobuf:"bytes,3,opt,name=assetExec,proto3" json:"assetExec,omitempty"`
	AssetSymbol          string   `protobuf:"bytes,4,opt,name=assetSymbol,proto3" json:"assetSymbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRetrieveInfo) Reset()         { *m = ReqRetrieveInfo{} }
func (m *ReqRetrieveInfo) String() string { return proto.CompactTextString(m) }
func (*ReqRetrieveInfo) ProtoMessage()    {}
func (*ReqRetrieveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieve_1666ef19b265e4fc, []int{8}
}
func (m *ReqRetrieveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRetrieveInfo.Unmarshal(m, b)
}
func (m *ReqRetrieveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRetrieveInfo.Marshal(b, m, deterministic)
}
func (dst *ReqRetrieveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRetrieveInfo.Merge(dst, src)
}
func (m *ReqRetrieveInfo) XXX_Size() int {
	return xxx_messageInfo_ReqRetrieveInfo.Size(m)
}
func (m *ReqRetrieveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRetrieveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRetrieveInfo proto.InternalMessageInfo

func (m *ReqRetrieveInfo) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *ReqRetrieveInfo) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

func (m *ReqRetrieveInfo) GetAssetExec() string {
	if m != nil {
		return m.AssetExec
	}
	return ""
}

func (m *ReqRetrieveInfo) GetAssetSymbol() string {
	if m != nil {
		return m.AssetSymbol
	}
	return ""
}

type RetrieveQuery struct {
	BackupAddress        string   `protobuf:"bytes,1,opt,name=backupAddress,proto3" json:"backupAddress,omitempty"`
	DefaultAddress       string   `protobuf:"bytes,2,opt,name=defaultAddress,proto3" json:"defaultAddress,omitempty"`
	DelayPeriod          int64    `protobuf:"varint,3,opt,name=delayPeriod,proto3" json:"delayPeriod,omitempty"`
	PrepareTime          int64    `protobuf:"varint,4,opt,name=prepareTime,proto3" json:"prepareTime,omitempty"`
	RemainTime           int64    `protobuf:"varint,5,opt,name=remainTime,proto3" json:"remainTime,omitempty"`
	Status               int32    `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveQuery) Reset()         { *m = RetrieveQuery{} }
func (m *RetrieveQuery) String() string { return proto.CompactTextString(m) }
func (*RetrieveQuery) ProtoMessage()    {}
func (*RetrieveQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieve_1666ef19b265e4fc, []int{9}
}
func (m *RetrieveQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveQuery.Unmarshal(m, b)
}
func (m *RetrieveQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveQuery.Marshal(b, m, deterministic)
}
func (dst *RetrieveQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveQuery.Merge(dst, src)
}
func (m *RetrieveQuery) XXX_Size() int {
	return xxx_messageInfo_RetrieveQuery.Size(m)
}
func (m *RetrieveQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveQuery proto.InternalMessageInfo

func (m *RetrieveQuery) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *RetrieveQuery) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

func (m *RetrieveQuery) GetDelayPeriod() int64 {
	if m != nil {
		return m.DelayPeriod
	}
	return 0
}

func (m *RetrieveQuery) GetPrepareTime() int64 {
	if m != nil {
		return m.PrepareTime
	}
	return 0
}

func (m *RetrieveQuery) GetRemainTime() int64 {
	if m != nil {
		return m.RemainTime
	}
	return 0
}

func (m *RetrieveQuery) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*RetrievePara)(nil), "types.RetrievePara")
	proto.RegisterType((*Retrieve)(nil), "types.Retrieve")
	proto.RegisterType((*RetrieveAction)(nil), "types.RetrieveAction")
	proto.RegisterType((*BackupRetrieve)(nil), "types.BackupRetrieve")
	proto.RegisterType((*PrepareRetrieve)(nil), "types.PrepareRetrieve")
	proto.RegisterType((*AssetSymbol)(nil), "types.AssetSymbol")
	proto.RegisterType((*PerformRetrieve)(nil), "types.PerformRetrieve")
	proto.RegisterType((*CancelRetrieve)(nil), "types.CancelRetrieve")
	proto.RegisterType((*ReqRetrieveInfo)(nil), "types.ReqRetrieveInfo")
	proto.RegisterType((*RetrieveQuery)(nil), "types.RetrieveQuery")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RetrieveClient is the client API for Retrieve service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RetrieveClient interface {
	Prepare(ctx context.Context, in *PrepareRetrieve, opts ...grpc.CallOption) (*types.UnsignTx, error)
	Perform(ctx context.Context, in *PerformRetrieve, opts ...grpc.CallOption) (*types.UnsignTx, error)
	Backup(ctx context.Context, in *BackupRetrieve, opts ...grpc.CallOption) (*types.UnsignTx, error)
	Cancel(ctx context.Context, in *CancelRetrieve, opts ...grpc.CallOption) (*types.UnsignTx, error)
}

type retrieveClient struct {
	cc *grpc.ClientConn
}

func NewRetrieveClient(cc *grpc.ClientConn) RetrieveClient {
	return &retrieveClient{cc}
}

func (c *retrieveClient) Prepare(ctx context.Context, in *PrepareRetrieve, opts ...grpc.CallOption) (*types.UnsignTx, error) {
	out := new(types.UnsignTx)
	err := c.cc.Invoke(ctx, "/types.retrieve/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrieveClient) Perform(ctx context.Context, in *PerformRetrieve, opts ...grpc.CallOption) (*types.UnsignTx, error) {
	out := new(types.UnsignTx)
	err := c.cc.Invoke(ctx, "/types.retrieve/Perform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrieveClient) Backup(ctx context.Context, in *BackupRetrieve, opts ...grpc.CallOption) (*types.UnsignTx, error) {
	out := new(types.UnsignTx)
	err := c.cc.Invoke(ctx, "/types.retrieve/Backup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrieveClient) Cancel(ctx context.Context, in *CancelRetrieve, opts ...grpc.CallOption) (*types.UnsignTx, error) {
	out := new(types.UnsignTx)
	err := c.cc.Invoke(ctx, "/types.retrieve/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetrieveServer is the server API for Retrieve service.
type RetrieveServer interface {
	Prepare(context.Context, *PrepareRetrieve) (*types.UnsignTx, error)
	Perform(context.Context, *PerformRetrieve) (*types.UnsignTx, error)
	Backup(context.Context, *BackupRetrieve) (*types.UnsignTx, error)
	Cancel(context.Context, *CancelRetrieve) (*types.UnsignTx, error)
}

func RegisterRetrieveServer(s *grpc.Server, srv RetrieveServer) {
	s.RegisterService(&_Retrieve_serviceDesc, srv)
}

func _Retrieve_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareRetrieve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieveServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.retrieve/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieveServer).Prepare(ctx, req.(*PrepareRetrieve))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieve_Perform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformRetrieve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieveServer).Perform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.retrieve/Perform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieveServer).Perform(ctx, req.(*PerformRetrieve))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieve_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupRetrieve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieveServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.retrieve/Backup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieveServer).Backup(ctx, req.(*BackupRetrieve))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieve_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRetrieve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieveServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.retrieve/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieveServer).Cancel(ctx, req.(*CancelRetrieve))
	}
	return interceptor(ctx, in, info, handler)
}

var _Retrieve_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.retrieve",
	HandlerType: (*RetrieveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _Retrieve_Prepare_Handler,
		},
		{
			MethodName: "Perform",
			Handler:    _Retrieve_Perform_Handler,
		},
		{
			MethodName: "Backup",
			Handler:    _Retrieve_Backup_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Retrieve_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "retrieve.proto",
}

func init() { proto.RegisterFile("retrieve.proto", fileDescriptor_retrieve_1666ef19b265e4fc) }

var fileDescriptor_retrieve_1666ef19b265e4fc = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xed, 0xc6, 0x8d, 0xd3, 0x4c, 0xa8, 0x23, 0x16, 0x51, 0x59, 0x15, 0xaa, 0x22, 0x0b, 0xa1,
	0x08, 0x89, 0x20, 0x19, 0x2e, 0x1c, 0x53, 0x84, 0x04, 0xb7, 0x60, 0xca, 0x15, 0xb4, 0x71, 0x26,
	0xc8, 0x22, 0xb1, 0xc3, 0xee, 0xa6, 0xaa, 0x6f, 0xdc, 0xf9, 0x08, 0xfe, 0x81, 0xbf, 0xe1, 0xc2,
	0x4f, 0xf0, 0x03, 0xc8, 0xb3, 0x6b, 0x65, 0x13, 0x1c, 0xa9, 0x87, 0xa8, 0xb7, 0xf8, 0xed, 0x7b,
	0xeb, 0x99, 0x37, 0x93, 0x67, 0x08, 0x24, 0x6a, 0x99, 0xe1, 0x35, 0x8e, 0x56, 0xb2, 0xd0, 0x05,
	0x6f, 0xeb, 0x72, 0x85, 0xea, 0xfc, 0xbe, 0x96, 0x22, 0x57, 0x22, 0xd5, 0x59, 0x91, 0x9b, 0x93,
	0xe8, 0x17, 0x83, 0x7b, 0x89, 0x25, 0x4f, 0x84, 0x14, 0xfc, 0x09, 0x04, 0x33, 0x9c, 0x8b, 0xf5,
	0x42, 0x8f, 0x67, 0x33, 0x89, 0x4a, 0x85, 0x6c, 0xc0, 0x86, 0xdd, 0x64, 0x07, 0xe5, 0x67, 0xe0,
	0x2b, 0x2d, 0xf4, 0x5a, 0x85, 0xad, 0x01, 0x1b, 0xb6, 0x13, 0xfb, 0xc4, 0x2f, 0x00, 0x52, 0x89,
	0x42, 0xe3, 0x55, 0xb6, 0xc4, 0xd0, 0x1b, 0xb0, 0xa1, 0x97, 0x38, 0x08, 0x1f, 0x40, 0x6f, 0x25,
	0x71, 0x25, 0xa4, 0x21, 0x1c, 0x13, 0xc1, 0x85, 0x2a, 0xc6, 0x0c, 0x17, 0xa2, 0x9c, 0xa0, 0xcc,
	0x8a, 0x59, 0xd8, 0x36, 0x0c, 0x07, 0x8a, 0x3e, 0xc3, 0x49, 0x5d, 0x33, 0x7f, 0x0c, 0xa7, 0x53,
	0x91, 0x7e, 0x5d, 0xaf, 0xb6, 0xcb, 0xdd, 0x06, 0xf9, 0x33, 0xe8, 0x48, 0xd4, 0x55, 0x83, 0x61,
	0x6b, 0xe0, 0x0d, 0x7b, 0xf1, 0x83, 0x11, 0x59, 0x32, 0x72, 0x7b, 0x4f, 0x6a, 0x4e, 0xf4, 0x97,
	0x41, 0x50, 0x9f, 0x8c, 0xc9, 0x2e, 0x1e, 0x43, 0xc7, 0x16, 0x49, 0x6f, 0xe8, 0xc5, 0x67, 0xf6,
	0x86, 0x89, 0x41, 0x6b, 0xfa, 0xdb, 0xa3, 0xa4, 0x26, 0x92, 0x06, 0xe5, 0xbc, 0x90, 0x4b, 0x32,
	0xc9, 0xd1, 0x18, 0x74, 0x4b, 0x63, 0x20, 0xfe, 0x1c, 0x7c, 0x53, 0x3a, 0x79, 0xd7, 0x8b, 0x1f,
	0x5a, 0xc9, 0x25, 0x81, 0x8e, 0xc2, 0xd2, 0x2a, 0x41, 0x2a, 0xf2, 0x14, 0x17, 0xe4, 0xe5, 0x46,
	0xf0, 0x9a, 0x40, 0x57, 0x60, 0x68, 0x3c, 0x80, 0x96, 0x2e, 0xc9, 0xd6, 0x76, 0xd2, 0xd2, 0xe5,
	0x65, 0x07, 0xda, 0xd7, 0x62, 0xb1, 0xc6, 0xe8, 0x3b, 0x83, 0x60, 0xfb, 0x35, 0xb7, 0x74, 0xf7,
	0xff, 0x9d, 0x69, 0x35, 0xee, 0xcc, 0xce, 0x64, 0xbd, 0xa6, 0xc9, 0xf6, 0x77, 0xfc, 0x3c, 0x6c,
	0x09, 0xd1, 0x2b, 0xe8, 0x8d, 0x95, 0x42, 0xfd, 0xa1, 0x5c, 0x4e, 0x8b, 0x05, 0xe7, 0x70, 0x8c,
	0x37, 0x98, 0xda, 0x3b, 0xe9, 0x37, 0x6d, 0x36, 0x9d, 0xda, 0x2b, 0xec, 0x53, 0xf4, 0x83, 0x41,
	0x7f, 0x67, 0x70, 0x07, 0xf6, 0xe7, 0x29, 0xf8, 0xa2, 0x2a, 0x4e, 0x85, 0x1e, 0x2d, 0x29, 0xb7,
	0xa3, 0x74, 0x2a, 0x4e, 0x2c, 0x23, 0xfa, 0x04, 0xc1, 0xf6, 0x84, 0x0f, 0x6c, 0xd4, 0x4f, 0x06,
	0xfd, 0x04, 0xbf, 0xd5, 0xb7, 0xbf, 0xcb, 0xe7, 0xc5, 0x81, 0xbb, 0x7d, 0x04, 0x5d, 0xea, 0xe5,
	0x4d, 0x35, 0x00, 0x8f, 0x28, 0x1b, 0xa0, 0xda, 0x15, 0xb1, 0x69, 0x9b, 0x76, 0xbb, 0x9b, 0xb8,
	0x50, 0xf4, 0x9b, 0xc1, 0x69, 0x5d, 0xde, 0xfb, 0x35, 0xca, 0xf2, 0xae, 0xb7, 0xf5, 0x16, 0x59,
	0x76, 0x01, 0x20, 0x71, 0x29, 0xb2, 0x9c, 0x08, 0x26, 0xca, 0x1c, 0xc4, 0x49, 0x51, 0xdf, 0x4d,
	0xd1, 0xf8, 0x0f, 0x83, 0x93, 0x3a, 0xc3, 0xf9, 0x4b, 0xe8, 0xd8, 0x3f, 0x05, 0xdf, 0x13, 0x3a,
	0xe7, 0x7d, 0x8b, 0x7f, 0xcc, 0x55, 0xf6, 0x25, 0xbf, 0xba, 0x89, 0x8e, 0x48, 0x65, 0x33, 0x65,
	0x4f, 0xec, 0x34, 0xa9, 0x62, 0xf0, 0x4d, 0x04, 0xf0, 0xe6, 0xe0, 0xd9, 0xa3, 0x31, 0xab, 0xc8,
	0x9b, 0xb3, 0xa7, 0x41, 0x33, 0xf5, 0xe9, 0xf3, 0xf3, 0xe2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x01, 0xad, 0xc9, 0x3d, 0xaa, 0x06, 0x00, 0x00,
}
