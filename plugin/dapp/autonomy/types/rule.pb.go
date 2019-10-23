// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rule.proto

package types

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"
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

type AutonomyProposalRule struct {
	PropRule *ProposalRule `protobuf:"bytes,1,opt,name=propRule,proto3" json:"propRule,omitempty"`
	CurRule  *RuleConfig   `protobuf:"bytes,2,opt,name=curRule,proto3" json:"curRule,omitempty"`
	// 全体持票人投票结果
	VoteResult *VoteResult `protobuf:"bytes,3,opt,name=voteResult,proto3" json:"voteResult,omitempty"`
	// 状态
	Status               int32    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Height               int64    `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	Index                int32    `protobuf:"varint,7,opt,name=index,proto3" json:"index,omitempty"`
	ProposalID           string   `protobuf:"bytes,8,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutonomyProposalRule) Reset()         { *m = AutonomyProposalRule{} }
func (m *AutonomyProposalRule) String() string { return proto.CompactTextString(m) }
func (*AutonomyProposalRule) ProtoMessage()    {}
func (*AutonomyProposalRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{0}
}
func (m *AutonomyProposalRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutonomyProposalRule.Unmarshal(m, b)
}
func (m *AutonomyProposalRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutonomyProposalRule.Marshal(b, m, deterministic)
}
func (dst *AutonomyProposalRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutonomyProposalRule.Merge(dst, src)
}
func (m *AutonomyProposalRule) XXX_Size() int {
	return xxx_messageInfo_AutonomyProposalRule.Size(m)
}
func (m *AutonomyProposalRule) XXX_DiscardUnknown() {
	xxx_messageInfo_AutonomyProposalRule.DiscardUnknown(m)
}

var xxx_messageInfo_AutonomyProposalRule proto.InternalMessageInfo

func (m *AutonomyProposalRule) GetPropRule() *ProposalRule {
	if m != nil {
		return m.PropRule
	}
	return nil
}

func (m *AutonomyProposalRule) GetCurRule() *RuleConfig {
	if m != nil {
		return m.CurRule
	}
	return nil
}

func (m *AutonomyProposalRule) GetVoteResult() *VoteResult {
	if m != nil {
		return m.VoteResult
	}
	return nil
}

func (m *AutonomyProposalRule) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AutonomyProposalRule) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AutonomyProposalRule) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *AutonomyProposalRule) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AutonomyProposalRule) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

type ProposalRule struct {
	// 提案时间
	Year  int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	// 规则可修改项,如果某项不修改则置为-1
	RuleCfg *RuleConfig `protobuf:"bytes,4,opt,name=ruleCfg,proto3" json:"ruleCfg,omitempty"`
	// 投票相关
	StartBlockHeight     int64    `protobuf:"varint,5,opt,name=startBlockHeight,proto3" json:"startBlockHeight,omitempty"`
	EndBlockHeight       int64    `protobuf:"varint,6,opt,name=endBlockHeight,proto3" json:"endBlockHeight,omitempty"`
	RealEndBlockHeight   int64    `protobuf:"varint,7,opt,name=realEndBlockHeight,proto3" json:"realEndBlockHeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProposalRule) Reset()         { *m = ProposalRule{} }
func (m *ProposalRule) String() string { return proto.CompactTextString(m) }
func (*ProposalRule) ProtoMessage()    {}
func (*ProposalRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{1}
}
func (m *ProposalRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalRule.Unmarshal(m, b)
}
func (m *ProposalRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalRule.Marshal(b, m, deterministic)
}
func (dst *ProposalRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalRule.Merge(dst, src)
}
func (m *ProposalRule) XXX_Size() int {
	return xxx_messageInfo_ProposalRule.Size(m)
}
func (m *ProposalRule) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalRule.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalRule proto.InternalMessageInfo

func (m *ProposalRule) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *ProposalRule) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *ProposalRule) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *ProposalRule) GetRuleCfg() *RuleConfig {
	if m != nil {
		return m.RuleCfg
	}
	return nil
}

func (m *ProposalRule) GetStartBlockHeight() int64 {
	if m != nil {
		return m.StartBlockHeight
	}
	return 0
}

func (m *ProposalRule) GetEndBlockHeight() int64 {
	if m != nil {
		return m.EndBlockHeight
	}
	return 0
}

func (m *ProposalRule) GetRealEndBlockHeight() int64 {
	if m != nil {
		return m.RealEndBlockHeight
	}
	return 0
}

type RevokeProposalRule struct {
	ProposalID           string   `protobuf:"bytes,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeProposalRule) Reset()         { *m = RevokeProposalRule{} }
func (m *RevokeProposalRule) String() string { return proto.CompactTextString(m) }
func (*RevokeProposalRule) ProtoMessage()    {}
func (*RevokeProposalRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{2}
}
func (m *RevokeProposalRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeProposalRule.Unmarshal(m, b)
}
func (m *RevokeProposalRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeProposalRule.Marshal(b, m, deterministic)
}
func (dst *RevokeProposalRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeProposalRule.Merge(dst, src)
}
func (m *RevokeProposalRule) XXX_Size() int {
	return xxx_messageInfo_RevokeProposalRule.Size(m)
}
func (m *RevokeProposalRule) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeProposalRule.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeProposalRule proto.InternalMessageInfo

func (m *RevokeProposalRule) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

type VoteProposalRule struct {
	ProposalID           string   `protobuf:"bytes,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	Approve              bool     `protobuf:"varint,2,opt,name=approve,proto3" json:"approve,omitempty"`
	OriginAddr           []string `protobuf:"bytes,3,rep,name=originAddr,proto3" json:"originAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteProposalRule) Reset()         { *m = VoteProposalRule{} }
func (m *VoteProposalRule) String() string { return proto.CompactTextString(m) }
func (*VoteProposalRule) ProtoMessage()    {}
func (*VoteProposalRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{3}
}
func (m *VoteProposalRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteProposalRule.Unmarshal(m, b)
}
func (m *VoteProposalRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteProposalRule.Marshal(b, m, deterministic)
}
func (dst *VoteProposalRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteProposalRule.Merge(dst, src)
}
func (m *VoteProposalRule) XXX_Size() int {
	return xxx_messageInfo_VoteProposalRule.Size(m)
}
func (m *VoteProposalRule) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteProposalRule.DiscardUnknown(m)
}

var xxx_messageInfo_VoteProposalRule proto.InternalMessageInfo

func (m *VoteProposalRule) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

func (m *VoteProposalRule) GetApprove() bool {
	if m != nil {
		return m.Approve
	}
	return false
}

func (m *VoteProposalRule) GetOriginAddr() []string {
	if m != nil {
		return m.OriginAddr
	}
	return nil
}

type TerminateProposalRule struct {
	ProposalID           string   `protobuf:"bytes,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminateProposalRule) Reset()         { *m = TerminateProposalRule{} }
func (m *TerminateProposalRule) String() string { return proto.CompactTextString(m) }
func (*TerminateProposalRule) ProtoMessage()    {}
func (*TerminateProposalRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{4}
}
func (m *TerminateProposalRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateProposalRule.Unmarshal(m, b)
}
func (m *TerminateProposalRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateProposalRule.Marshal(b, m, deterministic)
}
func (dst *TerminateProposalRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateProposalRule.Merge(dst, src)
}
func (m *TerminateProposalRule) XXX_Size() int {
	return xxx_messageInfo_TerminateProposalRule.Size(m)
}
func (m *TerminateProposalRule) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateProposalRule.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateProposalRule proto.InternalMessageInfo

func (m *TerminateProposalRule) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

// receipt
type ReceiptProposalRule struct {
	Prev                 *AutonomyProposalRule `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	Current              *AutonomyProposalRule `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ReceiptProposalRule) Reset()         { *m = ReceiptProposalRule{} }
func (m *ReceiptProposalRule) String() string { return proto.CompactTextString(m) }
func (*ReceiptProposalRule) ProtoMessage()    {}
func (*ReceiptProposalRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{5}
}
func (m *ReceiptProposalRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptProposalRule.Unmarshal(m, b)
}
func (m *ReceiptProposalRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptProposalRule.Marshal(b, m, deterministic)
}
func (dst *ReceiptProposalRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptProposalRule.Merge(dst, src)
}
func (m *ReceiptProposalRule) XXX_Size() int {
	return xxx_messageInfo_ReceiptProposalRule.Size(m)
}
func (m *ReceiptProposalRule) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptProposalRule.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptProposalRule proto.InternalMessageInfo

func (m *ReceiptProposalRule) GetPrev() *AutonomyProposalRule {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptProposalRule) GetCurrent() *AutonomyProposalRule {
	if m != nil {
		return m.Current
	}
	return nil
}

type LocalProposalRule struct {
	PropRule             *AutonomyProposalRule `protobuf:"bytes,1,opt,name=propRule,proto3" json:"propRule,omitempty"`
	Comments             []string              `protobuf:"bytes,2,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LocalProposalRule) Reset()         { *m = LocalProposalRule{} }
func (m *LocalProposalRule) String() string { return proto.CompactTextString(m) }
func (*LocalProposalRule) ProtoMessage()    {}
func (*LocalProposalRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{6}
}
func (m *LocalProposalRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalProposalRule.Unmarshal(m, b)
}
func (m *LocalProposalRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalProposalRule.Marshal(b, m, deterministic)
}
func (dst *LocalProposalRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalProposalRule.Merge(dst, src)
}
func (m *LocalProposalRule) XXX_Size() int {
	return xxx_messageInfo_LocalProposalRule.Size(m)
}
func (m *LocalProposalRule) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalProposalRule.DiscardUnknown(m)
}

var xxx_messageInfo_LocalProposalRule proto.InternalMessageInfo

func (m *LocalProposalRule) GetPropRule() *AutonomyProposalRule {
	if m != nil {
		return m.PropRule
	}
	return nil
}

func (m *LocalProposalRule) GetComments() []string {
	if m != nil {
		return m.Comments
	}
	return nil
}

// query
type ReqQueryProposalRule struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Direction            int32    `protobuf:"varint,4,opt,name=direction,proto3" json:"direction,omitempty"`
	Height               int64    `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	Index                int32    `protobuf:"varint,6,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqQueryProposalRule) Reset()         { *m = ReqQueryProposalRule{} }
func (m *ReqQueryProposalRule) String() string { return proto.CompactTextString(m) }
func (*ReqQueryProposalRule) ProtoMessage()    {}
func (*ReqQueryProposalRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{7}
}
func (m *ReqQueryProposalRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqQueryProposalRule.Unmarshal(m, b)
}
func (m *ReqQueryProposalRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqQueryProposalRule.Marshal(b, m, deterministic)
}
func (dst *ReqQueryProposalRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqQueryProposalRule.Merge(dst, src)
}
func (m *ReqQueryProposalRule) XXX_Size() int {
	return xxx_messageInfo_ReqQueryProposalRule.Size(m)
}
func (m *ReqQueryProposalRule) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqQueryProposalRule.DiscardUnknown(m)
}

var xxx_messageInfo_ReqQueryProposalRule proto.InternalMessageInfo

func (m *ReqQueryProposalRule) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReqQueryProposalRule) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqQueryProposalRule) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqQueryProposalRule) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqQueryProposalRule) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReqQueryProposalRule) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ReplyQueryProposalRule struct {
	PropRules            []*AutonomyProposalRule `protobuf:"bytes,1,rep,name=propRules,proto3" json:"propRules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ReplyQueryProposalRule) Reset()         { *m = ReplyQueryProposalRule{} }
func (m *ReplyQueryProposalRule) String() string { return proto.CompactTextString(m) }
func (*ReplyQueryProposalRule) ProtoMessage()    {}
func (*ReplyQueryProposalRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{8}
}
func (m *ReplyQueryProposalRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyQueryProposalRule.Unmarshal(m, b)
}
func (m *ReplyQueryProposalRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyQueryProposalRule.Marshal(b, m, deterministic)
}
func (dst *ReplyQueryProposalRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyQueryProposalRule.Merge(dst, src)
}
func (m *ReplyQueryProposalRule) XXX_Size() int {
	return xxx_messageInfo_ReplyQueryProposalRule.Size(m)
}
func (m *ReplyQueryProposalRule) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyQueryProposalRule.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyQueryProposalRule proto.InternalMessageInfo

func (m *ReplyQueryProposalRule) GetPropRules() []*AutonomyProposalRule {
	if m != nil {
		return m.PropRules
	}
	return nil
}

// TransferFund action
type TransferFund struct {
	Amount               int64    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Note                 string   `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferFund) Reset()         { *m = TransferFund{} }
func (m *TransferFund) String() string { return proto.CompactTextString(m) }
func (*TransferFund) ProtoMessage()    {}
func (*TransferFund) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{9}
}
func (m *TransferFund) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferFund.Unmarshal(m, b)
}
func (m *TransferFund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferFund.Marshal(b, m, deterministic)
}
func (dst *TransferFund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferFund.Merge(dst, src)
}
func (m *TransferFund) XXX_Size() int {
	return xxx_messageInfo_TransferFund.Size(m)
}
func (m *TransferFund) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferFund.DiscardUnknown(m)
}

var xxx_messageInfo_TransferFund proto.InternalMessageInfo

func (m *TransferFund) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransferFund) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

// Comment action
type Comment struct {
	ProposalID           string   `protobuf:"bytes,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	RepHash              string   `protobuf:"bytes,2,opt,name=repHash,proto3" json:"repHash,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{10}
}
func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (dst *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(dst, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

func (m *Comment) GetRepHash() string {
	if m != nil {
		return m.RepHash
	}
	return ""
}

func (m *Comment) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type ReceiptProposalComment struct {
	Cmt                  *Comment `protobuf:"bytes,1,opt,name=cmt,proto3" json:"cmt,omitempty"`
	Height               int64    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Index                int32    `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Hash                 string   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiptProposalComment) Reset()         { *m = ReceiptProposalComment{} }
func (m *ReceiptProposalComment) String() string { return proto.CompactTextString(m) }
func (*ReceiptProposalComment) ProtoMessage()    {}
func (*ReceiptProposalComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{11}
}
func (m *ReceiptProposalComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptProposalComment.Unmarshal(m, b)
}
func (m *ReceiptProposalComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptProposalComment.Marshal(b, m, deterministic)
}
func (dst *ReceiptProposalComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptProposalComment.Merge(dst, src)
}
func (m *ReceiptProposalComment) XXX_Size() int {
	return xxx_messageInfo_ReceiptProposalComment.Size(m)
}
func (m *ReceiptProposalComment) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptProposalComment.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptProposalComment proto.InternalMessageInfo

func (m *ReceiptProposalComment) GetCmt() *Comment {
	if m != nil {
		return m.Cmt
	}
	return nil
}

func (m *ReceiptProposalComment) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReceiptProposalComment) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ReceiptProposalComment) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// query
type ReqQueryProposalComment struct {
	ProposalID           string   `protobuf:"bytes,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Direction            int32    `protobuf:"varint,3,opt,name=direction,proto3" json:"direction,omitempty"`
	Height               int64    `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Index                int32    `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqQueryProposalComment) Reset()         { *m = ReqQueryProposalComment{} }
func (m *ReqQueryProposalComment) String() string { return proto.CompactTextString(m) }
func (*ReqQueryProposalComment) ProtoMessage()    {}
func (*ReqQueryProposalComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{12}
}
func (m *ReqQueryProposalComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqQueryProposalComment.Unmarshal(m, b)
}
func (m *ReqQueryProposalComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqQueryProposalComment.Marshal(b, m, deterministic)
}
func (dst *ReqQueryProposalComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqQueryProposalComment.Merge(dst, src)
}
func (m *ReqQueryProposalComment) XXX_Size() int {
	return xxx_messageInfo_ReqQueryProposalComment.Size(m)
}
func (m *ReqQueryProposalComment) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqQueryProposalComment.DiscardUnknown(m)
}

var xxx_messageInfo_ReqQueryProposalComment proto.InternalMessageInfo

func (m *ReqQueryProposalComment) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

func (m *ReqQueryProposalComment) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqQueryProposalComment) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqQueryProposalComment) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReqQueryProposalComment) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type RelationCmt struct {
	RepHash              string   `protobuf:"bytes,1,opt,name=repHash,proto3" json:"repHash,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Height               int64    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Index                int32    `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	Hash                 string   `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationCmt) Reset()         { *m = RelationCmt{} }
func (m *RelationCmt) String() string { return proto.CompactTextString(m) }
func (*RelationCmt) ProtoMessage()    {}
func (*RelationCmt) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{13}
}
func (m *RelationCmt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationCmt.Unmarshal(m, b)
}
func (m *RelationCmt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationCmt.Marshal(b, m, deterministic)
}
func (dst *RelationCmt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationCmt.Merge(dst, src)
}
func (m *RelationCmt) XXX_Size() int {
	return xxx_messageInfo_RelationCmt.Size(m)
}
func (m *RelationCmt) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationCmt.DiscardUnknown(m)
}

var xxx_messageInfo_RelationCmt proto.InternalMessageInfo

func (m *RelationCmt) GetRepHash() string {
	if m != nil {
		return m.RepHash
	}
	return ""
}

func (m *RelationCmt) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *RelationCmt) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *RelationCmt) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RelationCmt) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type ReplyQueryProposalComment struct {
	RltCmt               []*RelationCmt `protobuf:"bytes,1,rep,name=rltCmt,proto3" json:"rltCmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyQueryProposalComment) Reset()         { *m = ReplyQueryProposalComment{} }
func (m *ReplyQueryProposalComment) String() string { return proto.CompactTextString(m) }
func (*ReplyQueryProposalComment) ProtoMessage()    {}
func (*ReplyQueryProposalComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_rule_2ed9cd03418a949e, []int{14}
}
func (m *ReplyQueryProposalComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyQueryProposalComment.Unmarshal(m, b)
}
func (m *ReplyQueryProposalComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyQueryProposalComment.Marshal(b, m, deterministic)
}
func (dst *ReplyQueryProposalComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyQueryProposalComment.Merge(dst, src)
}
func (m *ReplyQueryProposalComment) XXX_Size() int {
	return xxx_messageInfo_ReplyQueryProposalComment.Size(m)
}
func (m *ReplyQueryProposalComment) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyQueryProposalComment.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyQueryProposalComment proto.InternalMessageInfo

func (m *ReplyQueryProposalComment) GetRltCmt() []*RelationCmt {
	if m != nil {
		return m.RltCmt
	}
	return nil
}

func init() {
	proto.RegisterType((*AutonomyProposalRule)(nil), "types.AutonomyProposalRule")
	proto.RegisterType((*ProposalRule)(nil), "types.ProposalRule")
	proto.RegisterType((*RevokeProposalRule)(nil), "types.RevokeProposalRule")
	proto.RegisterType((*VoteProposalRule)(nil), "types.VoteProposalRule")
	proto.RegisterType((*TerminateProposalRule)(nil), "types.TerminateProposalRule")
	proto.RegisterType((*ReceiptProposalRule)(nil), "types.ReceiptProposalRule")
	proto.RegisterType((*LocalProposalRule)(nil), "types.LocalProposalRule")
	proto.RegisterType((*ReqQueryProposalRule)(nil), "types.ReqQueryProposalRule")
	proto.RegisterType((*ReplyQueryProposalRule)(nil), "types.ReplyQueryProposalRule")
	proto.RegisterType((*TransferFund)(nil), "types.TransferFund")
	proto.RegisterType((*Comment)(nil), "types.Comment")
	proto.RegisterType((*ReceiptProposalComment)(nil), "types.ReceiptProposalComment")
	proto.RegisterType((*ReqQueryProposalComment)(nil), "types.ReqQueryProposalComment")
	proto.RegisterType((*RelationCmt)(nil), "types.RelationCmt")
	proto.RegisterType((*ReplyQueryProposalComment)(nil), "types.ReplyQueryProposalComment")
}

func init() { proto.RegisterFile("rule.proto", fileDescriptor_rule_2ed9cd03418a949e) }

var fileDescriptor_rule_2ed9cd03418a949e = []byte{
	// 733 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4d, 0x6f, 0xd4, 0x3a,
	0x14, 0x55, 0x26, 0xc9, 0x7c, 0xdc, 0xf6, 0x55, 0xad, 0xdb, 0xd7, 0x97, 0xd7, 0xf7, 0x84, 0x46,
	0x59, 0xa0, 0x51, 0x91, 0x06, 0xf1, 0xa5, 0x0a, 0x76, 0x65, 0xf8, 0x28, 0x12, 0x0b, 0x30, 0x15,
	0x3b, 0x16, 0x21, 0xb9, 0x9d, 0x89, 0x9a, 0xd8, 0xc1, 0x71, 0x46, 0x8c, 0x04, 0x2b, 0x7e, 0x06,
	0x5b, 0x24, 0x7e, 0x24, 0x1b, 0x64, 0xc7, 0x99, 0x49, 0x3a, 0x69, 0x69, 0x77, 0xbe, 0xf6, 0xb1,
	0x7d, 0xcf, 0xf1, 0x39, 0x09, 0x80, 0x28, 0x12, 0x1c, 0x67, 0x82, 0x4b, 0x4e, 0x5c, 0xb9, 0xc8,
	0x30, 0x3f, 0xf8, 0x2b, 0x09, 0x79, 0x9a, 0x72, 0x56, 0xce, 0xfa, 0x3f, 0x3b, 0xb0, 0x77, 0x5c,
	0x48, 0xce, 0x78, 0xba, 0x78, 0x23, 0x78, 0xc6, 0xf3, 0x20, 0xa1, 0x45, 0x82, 0xe4, 0x2e, 0xf4,
	0x33, 0xc1, 0x33, 0x35, 0xf6, 0xac, 0xa1, 0x35, 0xda, 0xb8, 0xbf, 0x3b, 0xd6, 0x27, 0x8c, 0xeb,
	0x30, 0xba, 0x04, 0x91, 0x3b, 0xd0, 0x0b, 0x0b, 0xa1, 0xf1, 0x1d, 0x8d, 0xdf, 0x31, 0x78, 0x35,
	0x35, 0xe1, 0xec, 0x2c, 0x9e, 0xd2, 0x0a, 0x41, 0xee, 0x01, 0xcc, 0xb9, 0x44, 0x8a, 0x79, 0x91,
	0x48, 0xcf, 0x6e, 0xe0, 0xdf, 0x2f, 0x17, 0x68, 0x0d, 0x44, 0xf6, 0xa1, 0x9b, 0xcb, 0x40, 0x16,
	0xb9, 0xe7, 0x0c, 0xad, 0x91, 0x4b, 0x4d, 0x45, 0x3c, 0xe8, 0x05, 0x51, 0x24, 0x30, 0xcf, 0x3d,
	0x77, 0x68, 0x8d, 0x06, 0xb4, 0x2a, 0xd5, 0x8e, 0x19, 0xc6, 0xd3, 0x99, 0xf4, 0xba, 0x43, 0x6b,
	0x64, 0x53, 0x53, 0x91, 0x3d, 0x70, 0x63, 0x16, 0xe1, 0x67, 0xaf, 0xa7, 0x0f, 0x2a, 0x0b, 0x72,
	0x0b, 0x20, 0x33, 0xcc, 0x5e, 0x3d, 0xf3, 0xfa, 0xfa, 0xa8, 0xda, 0x8c, 0xff, 0xcb, 0x82, 0xcd,
	0x86, 0x42, 0x04, 0x9c, 0x05, 0x06, 0x42, 0xab, 0xe3, 0x52, 0x3d, 0x56, 0x47, 0xa7, 0x9c, 0xc9,
	0x99, 0x96, 0xc0, 0xa5, 0x65, 0x41, 0xb6, 0xc1, 0x8e, 0x82, 0x85, 0xa6, 0xe9, 0x52, 0x35, 0x54,
	0x62, 0xa9, 0xa7, 0x99, 0x9c, 0x4d, 0x35, 0x9b, 0x76, 0xb1, 0x0c, 0x82, 0x1c, 0xc2, 0x76, 0x2e,
	0x03, 0x21, 0x9f, 0x26, 0x3c, 0x3c, 0x3f, 0x29, 0x19, 0xb9, 0x9a, 0xd1, 0xda, 0x3c, 0xb9, 0x0d,
	0x5b, 0xc8, 0xa2, 0x3a, 0xb2, 0xe4, 0x7e, 0x61, 0x96, 0x8c, 0x81, 0x08, 0x0c, 0x92, 0xe7, 0x4d,
	0x6c, 0x4f, 0x63, 0x5b, 0x56, 0xfc, 0x87, 0x40, 0x28, 0xce, 0xf9, 0x39, 0x36, 0x24, 0x68, 0x6a,
	0x66, 0xad, 0x69, 0x96, 0xc0, 0xb6, 0x7a, 0xcd, 0x9b, 0xec, 0xd1, 0xef, 0x99, 0x65, 0x82, 0xcf,
	0x4b, 0x1f, 0xf5, 0x69, 0x55, 0xaa, 0x9d, 0x5c, 0xc4, 0xd3, 0x98, 0x1d, 0x47, 0x91, 0xf0, 0xec,
	0xa1, 0xad, 0x76, 0xae, 0x66, 0xfc, 0x23, 0xf8, 0xfb, 0x14, 0x45, 0x1a, 0xb3, 0xe0, 0x66, 0x57,
	0xfa, 0x5f, 0x61, 0x97, 0x62, 0x88, 0x71, 0x26, 0x2f, 0x44, 0xc0, 0xc9, 0x04, 0xce, 0x8d, 0xfd,
	0xff, 0x33, 0x2f, 0xd4, 0x96, 0x16, 0xaa, 0x81, 0xe4, 0x91, 0x8e, 0x80, 0x40, 0x26, 0x4d, 0x04,
	0xae, 0xdc, 0x53, 0x61, 0xfd, 0x19, 0xec, 0xbc, 0xe6, 0x61, 0x90, 0x34, 0x2e, 0x3f, 0x5a, 0xcb,
	0xdf, 0x95, 0x87, 0xad, 0x72, 0x78, 0x00, 0x7d, 0x95, 0x70, 0x64, 0x32, 0xf7, 0x3a, 0x5a, 0xa3,
	0x65, 0xed, 0xff, 0xb0, 0x60, 0x8f, 0xe2, 0xa7, 0xb7, 0x05, 0x8a, 0x66, 0xda, 0x57, 0xe1, 0xb2,
	0x1a, 0xe1, 0x22, 0xe0, 0xa8, 0x34, 0x69, 0x3a, 0x03, 0xaa, 0xc7, 0xca, 0xe3, 0x21, 0x2f, 0x98,
	0x34, 0x7e, 0x2e, 0x0b, 0xf2, 0x3f, 0x0c, 0xa2, 0x58, 0x60, 0x28, 0x63, 0xce, 0x4c, 0x42, 0x57,
	0x13, 0xb5, 0x28, 0xba, 0xed, 0x51, 0xec, 0xd6, 0xa2, 0xe8, 0xbf, 0x83, 0x7d, 0x8a, 0x59, 0xb2,
	0x58, 0xef, 0xf3, 0x31, 0x0c, 0x2a, 0xa2, 0xaa, 0x55, 0xfb, 0x4f, 0xb2, 0xac, 0xd0, 0xfe, 0x13,
	0xd8, 0x3c, 0x15, 0x01, 0xcb, 0xcf, 0x50, 0xbc, 0x28, 0x58, 0xa4, 0x5a, 0x0a, 0x52, 0xcd, 0xc3,
	0x2a, 0x5b, 0x2a, 0x2b, 0x45, 0x99, 0x71, 0x89, 0x15, 0x65, 0x35, 0xf6, 0x3f, 0x40, 0x6f, 0x52,
	0x6a, 0x78, 0x1d, 0xfb, 0x0a, 0xcc, 0x4e, 0x82, 0x7c, 0x66, 0x4e, 0xa8, 0x4a, 0xb5, 0x62, 0x1e,
	0x42, 0x2b, 0x37, 0xa0, 0x55, 0xe9, 0x7f, 0x51, 0x7c, 0x1b, 0xfe, 0xab, 0x6e, 0x1b, 0x82, 0x1d,
	0xa6, 0xd2, 0x18, 0x60, 0xcb, 0x30, 0x35, 0x8b, 0x54, 0x2d, 0xd5, 0x94, 0xed, 0xb4, 0x2b, 0x6b,
	0xd7, 0x3f, 0x72, 0x04, 0x9c, 0x99, 0x6a, 0xcd, 0x29, 0xc9, 0xa9, 0xb1, 0xff, 0xdd, 0x82, 0x7f,
	0x2e, 0x9a, 0xe2, 0xba, 0x6c, 0x97, 0x5e, 0xe8, 0x5c, 0xea, 0x05, 0xfb, 0x72, 0x2f, 0x38, 0xed,
	0x1d, 0xbb, 0x75, 0x2f, 0x7c, 0xb3, 0x60, 0x83, 0x62, 0x12, 0xa8, 0xad, 0x93, 0x54, 0xd6, 0xf5,
	0xb5, 0x2e, 0xd5, 0xb7, 0xd3, 0xd0, 0xb7, 0x76, 0xa3, 0xdd, 0x7e, 0xa3, 0xd3, 0xa6, 0x91, 0x5b,
	0xd3, 0xe8, 0x25, 0xfc, 0xbb, 0xee, 0xc8, 0x4a, 0xa4, 0x43, 0xe8, 0x8a, 0x44, 0x4e, 0xf4, 0x3b,
	0x29, 0x47, 0x92, 0xea, 0x5b, 0xbe, 0x6a, 0x9b, 0x1a, 0xc4, 0xc7, 0xae, 0xfe, 0xed, 0x3e, 0xf8,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x04, 0xe9, 0xf7, 0xf7, 0x9a, 0x07, 0x00, 0x00,
}