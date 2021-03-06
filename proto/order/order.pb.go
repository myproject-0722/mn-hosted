// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package go_mnhosted_srv_order

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NewRequest struct {
	UserID               int64    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Coinname             string   `protobuf:"bytes,2,opt,name=coinname,proto3" json:"coinname,omitempty"`
	Timetype             int32    `protobuf:"varint,3,opt,name=timetype,proto3" json:"timetype,omitempty"`
	Price                int32    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	TxID                 string   `protobuf:"bytes,5,opt,name=txID,proto3" json:"txID,omitempty"`
	IsRenew              int32    `protobuf:"varint,6,opt,name=isRenew,proto3" json:"isRenew,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewRequest) Reset()         { *m = NewRequest{} }
func (m *NewRequest) String() string { return proto.CompactTextString(m) }
func (*NewRequest) ProtoMessage()    {}
func (*NewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *NewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewRequest.Unmarshal(m, b)
}
func (m *NewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewRequest.Marshal(b, m, deterministic)
}
func (m *NewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewRequest.Merge(m, src)
}
func (m *NewRequest) XXX_Size() int {
	return xxx_messageInfo_NewRequest.Size(m)
}
func (m *NewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewRequest proto.InternalMessageInfo

func (m *NewRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *NewRequest) GetCoinname() string {
	if m != nil {
		return m.Coinname
	}
	return ""
}

func (m *NewRequest) GetTimetype() int32 {
	if m != nil {
		return m.Timetype
	}
	return 0
}

func (m *NewRequest) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *NewRequest) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *NewRequest) GetIsRenew() int32 {
	if m != nil {
		return m.IsRenew
	}
	return 0
}

type NewResponse struct {
	Rescode              int32    `protobuf:"varint,1,opt,name=rescode,proto3" json:"rescode,omitempty"`
	ID                   int64    `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewResponse) Reset()         { *m = NewResponse{} }
func (m *NewResponse) String() string { return proto.CompactTextString(m) }
func (*NewResponse) ProtoMessage()    {}
func (*NewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *NewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewResponse.Unmarshal(m, b)
}
func (m *NewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewResponse.Marshal(b, m, deterministic)
}
func (m *NewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewResponse.Merge(m, src)
}
func (m *NewResponse) XXX_Size() int {
	return xxx_messageInfo_NewResponse.Size(m)
}
func (m *NewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewResponse proto.InternalMessageInfo

func (m *NewResponse) GetRescode() int32 {
	if m != nil {
		return m.Rescode
	}
	return 0
}

func (m *NewResponse) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type AlipayRequest struct {
	UserID               int64    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	CoinName             string   `protobuf:"bytes,2,opt,name=coinName,proto3" json:"coinName,omitempty"`
	MNKey                string   `protobuf:"bytes,3,opt,name=MNKey,proto3" json:"MNKey,omitempty"`
	MNName               string   `protobuf:"bytes,4,opt,name=MNName,proto3" json:"MNName,omitempty"`
	TxID                 string   `protobuf:"bytes,5,opt,name=txID,proto3" json:"txID,omitempty"`
	TxIndex              int32    `protobuf:"varint,6,opt,name=txIndex,proto3" json:"txIndex,omitempty"`
	TimeType             int32    `protobuf:"varint,7,opt,name=timeType,proto3" json:"timeType,omitempty"`
	IsRenew              int32    `protobuf:"varint,8,opt,name=isRenew,proto3" json:"isRenew,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlipayRequest) Reset()         { *m = AlipayRequest{} }
func (m *AlipayRequest) String() string { return proto.CompactTextString(m) }
func (*AlipayRequest) ProtoMessage()    {}
func (*AlipayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *AlipayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlipayRequest.Unmarshal(m, b)
}
func (m *AlipayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlipayRequest.Marshal(b, m, deterministic)
}
func (m *AlipayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlipayRequest.Merge(m, src)
}
func (m *AlipayRequest) XXX_Size() int {
	return xxx_messageInfo_AlipayRequest.Size(m)
}
func (m *AlipayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlipayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlipayRequest proto.InternalMessageInfo

func (m *AlipayRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *AlipayRequest) GetCoinName() string {
	if m != nil {
		return m.CoinName
	}
	return ""
}

func (m *AlipayRequest) GetMNKey() string {
	if m != nil {
		return m.MNKey
	}
	return ""
}

func (m *AlipayRequest) GetMNName() string {
	if m != nil {
		return m.MNName
	}
	return ""
}

func (m *AlipayRequest) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *AlipayRequest) GetTxIndex() int32 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

func (m *AlipayRequest) GetTimeType() int32 {
	if m != nil {
		return m.TimeType
	}
	return 0
}

func (m *AlipayRequest) GetIsRenew() int32 {
	if m != nil {
		return m.IsRenew
	}
	return 0
}

type AlipayResponse struct {
	Rescode              int32    `protobuf:"varint,1,opt,name=rescode,proto3" json:"rescode,omitempty"`
	PayUrl               string   `protobuf:"bytes,2,opt,name=payUrl,proto3" json:"payUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlipayResponse) Reset()         { *m = AlipayResponse{} }
func (m *AlipayResponse) String() string { return proto.CompactTextString(m) }
func (*AlipayResponse) ProtoMessage()    {}
func (*AlipayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}

func (m *AlipayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlipayResponse.Unmarshal(m, b)
}
func (m *AlipayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlipayResponse.Marshal(b, m, deterministic)
}
func (m *AlipayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlipayResponse.Merge(m, src)
}
func (m *AlipayResponse) XXX_Size() int {
	return xxx_messageInfo_AlipayResponse.Size(m)
}
func (m *AlipayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlipayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlipayResponse proto.InternalMessageInfo

func (m *AlipayResponse) GetRescode() int32 {
	if m != nil {
		return m.Rescode
	}
	return 0
}

func (m *AlipayResponse) GetPayUrl() string {
	if m != nil {
		return m.PayUrl
	}
	return ""
}

type ConfirmAlipayRequest struct {
	OrderID              int64    `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
	Price                int32    `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmAlipayRequest) Reset()         { *m = ConfirmAlipayRequest{} }
func (m *ConfirmAlipayRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmAlipayRequest) ProtoMessage()    {}
func (*ConfirmAlipayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{4}
}

func (m *ConfirmAlipayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmAlipayRequest.Unmarshal(m, b)
}
func (m *ConfirmAlipayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmAlipayRequest.Marshal(b, m, deterministic)
}
func (m *ConfirmAlipayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmAlipayRequest.Merge(m, src)
}
func (m *ConfirmAlipayRequest) XXX_Size() int {
	return xxx_messageInfo_ConfirmAlipayRequest.Size(m)
}
func (m *ConfirmAlipayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmAlipayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmAlipayRequest proto.InternalMessageInfo

func (m *ConfirmAlipayRequest) GetOrderID() int64 {
	if m != nil {
		return m.OrderID
	}
	return 0
}

func (m *ConfirmAlipayRequest) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type ConfirmAlipayResponse struct {
	Rescode              int32    `protobuf:"varint,1,opt,name=rescode,proto3" json:"rescode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmAlipayResponse) Reset()         { *m = ConfirmAlipayResponse{} }
func (m *ConfirmAlipayResponse) String() string { return proto.CompactTextString(m) }
func (*ConfirmAlipayResponse) ProtoMessage()    {}
func (*ConfirmAlipayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{5}
}

func (m *ConfirmAlipayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmAlipayResponse.Unmarshal(m, b)
}
func (m *ConfirmAlipayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmAlipayResponse.Marshal(b, m, deterministic)
}
func (m *ConfirmAlipayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmAlipayResponse.Merge(m, src)
}
func (m *ConfirmAlipayResponse) XXX_Size() int {
	return xxx_messageInfo_ConfirmAlipayResponse.Size(m)
}
func (m *ConfirmAlipayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmAlipayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmAlipayResponse proto.InternalMessageInfo

func (m *ConfirmAlipayResponse) GetRescode() int32 {
	if m != nil {
		return m.Rescode
	}
	return 0
}

type UpdateRequest struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	MNKey                string   `protobuf:"bytes,2,opt,name=MNKey,proto3" json:"MNKey,omitempty"`
	Status               int32    `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{6}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *UpdateRequest) GetMNKey() string {
	if m != nil {
		return m.MNKey
	}
	return ""
}

func (m *UpdateRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type UpdateResponse struct {
	Rescode              int32    `protobuf:"varint,1,opt,name=rescode,proto3" json:"rescode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{7}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetRescode() int32 {
	if m != nil {
		return m.Rescode
	}
	return 0
}

type GetInfoRequest struct {
	UserID               int64    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{8}
}

func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoRequest.Unmarshal(m, b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(m, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoRequest.Size(m)
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

func (m *GetInfoRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type GetInfoResponse struct {
	Rescode              int32    `protobuf:"varint,1,opt,name=rescode,proto3" json:"rescode,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Payout               float64  `protobuf:"fixed64,3,opt,name=payout,proto3" json:"payout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoResponse) Reset()         { *m = GetInfoResponse{} }
func (m *GetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetInfoResponse) ProtoMessage()    {}
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{9}
}

func (m *GetInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoResponse.Unmarshal(m, b)
}
func (m *GetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoResponse.Merge(m, src)
}
func (m *GetInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetInfoResponse.Size(m)
}
func (m *GetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoResponse proto.InternalMessageInfo

func (m *GetInfoResponse) GetRescode() int32 {
	if m != nil {
		return m.Rescode
	}
	return 0
}

func (m *GetInfoResponse) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *GetInfoResponse) GetPayout() float64 {
	if m != nil {
		return m.Payout
	}
	return 0
}

type GetOrderListRequest struct {
	UserID               int64    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderListRequest) Reset()         { *m = GetOrderListRequest{} }
func (m *GetOrderListRequest) String() string { return proto.CompactTextString(m) }
func (*GetOrderListRequest) ProtoMessage()    {}
func (*GetOrderListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{10}
}

func (m *GetOrderListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderListRequest.Unmarshal(m, b)
}
func (m *GetOrderListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderListRequest.Marshal(b, m, deterministic)
}
func (m *GetOrderListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderListRequest.Merge(m, src)
}
func (m *GetOrderListRequest) XXX_Size() int {
	return xxx_messageInfo_GetOrderListRequest.Size(m)
}
func (m *GetOrderListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderListRequest proto.InternalMessageInfo

func (m *GetOrderListRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type OrderItem struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CoinName             string   `protobuf:"bytes,2,opt,name=coinName,proto3" json:"coinName,omitempty"`
	MNKey                string   `protobuf:"bytes,3,opt,name=MNKey,proto3" json:"MNKey,omitempty"`
	TimeType             int32    `protobuf:"varint,4,opt,name=timeType,proto3" json:"timeType,omitempty"`
	Price                float64  `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Status               int32    `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`
	IsRenew              int32    `protobuf:"varint,7,opt,name=isRenew,proto3" json:"isRenew,omitempty"`
	CreateTime           string   `protobuf:"bytes,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderItem) Reset()         { *m = OrderItem{} }
func (m *OrderItem) String() string { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()    {}
func (*OrderItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{11}
}

func (m *OrderItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderItem.Unmarshal(m, b)
}
func (m *OrderItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderItem.Marshal(b, m, deterministic)
}
func (m *OrderItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderItem.Merge(m, src)
}
func (m *OrderItem) XXX_Size() int {
	return xxx_messageInfo_OrderItem.Size(m)
}
func (m *OrderItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderItem proto.InternalMessageInfo

func (m *OrderItem) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *OrderItem) GetCoinName() string {
	if m != nil {
		return m.CoinName
	}
	return ""
}

func (m *OrderItem) GetMNKey() string {
	if m != nil {
		return m.MNKey
	}
	return ""
}

func (m *OrderItem) GetTimeType() int32 {
	if m != nil {
		return m.TimeType
	}
	return 0
}

func (m *OrderItem) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderItem) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *OrderItem) GetIsRenew() int32 {
	if m != nil {
		return m.IsRenew
	}
	return 0
}

func (m *OrderItem) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

type GetOrderListResponse struct {
	Rescode              int32        `protobuf:"varint,1,opt,name=rescode,proto3" json:"rescode,omitempty"`
	OrderList            []*OrderItem `protobuf:"bytes,2,rep,name=orderList,proto3" json:"orderList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetOrderListResponse) Reset()         { *m = GetOrderListResponse{} }
func (m *GetOrderListResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrderListResponse) ProtoMessage()    {}
func (*GetOrderListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{12}
}

func (m *GetOrderListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderListResponse.Unmarshal(m, b)
}
func (m *GetOrderListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderListResponse.Marshal(b, m, deterministic)
}
func (m *GetOrderListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderListResponse.Merge(m, src)
}
func (m *GetOrderListResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrderListResponse.Size(m)
}
func (m *GetOrderListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderListResponse proto.InternalMessageInfo

func (m *GetOrderListResponse) GetRescode() int32 {
	if m != nil {
		return m.Rescode
	}
	return 0
}

func (m *GetOrderListResponse) GetOrderList() []*OrderItem {
	if m != nil {
		return m.OrderList
	}
	return nil
}

func init() {
	proto.RegisterType((*NewRequest)(nil), "go.mnhosted.srv.order.NewRequest")
	proto.RegisterType((*NewResponse)(nil), "go.mnhosted.srv.order.NewResponse")
	proto.RegisterType((*AlipayRequest)(nil), "go.mnhosted.srv.order.AlipayRequest")
	proto.RegisterType((*AlipayResponse)(nil), "go.mnhosted.srv.order.AlipayResponse")
	proto.RegisterType((*ConfirmAlipayRequest)(nil), "go.mnhosted.srv.order.ConfirmAlipayRequest")
	proto.RegisterType((*ConfirmAlipayResponse)(nil), "go.mnhosted.srv.order.ConfirmAlipayResponse")
	proto.RegisterType((*UpdateRequest)(nil), "go.mnhosted.srv.order.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "go.mnhosted.srv.order.UpdateResponse")
	proto.RegisterType((*GetInfoRequest)(nil), "go.mnhosted.srv.order.GetInfoRequest")
	proto.RegisterType((*GetInfoResponse)(nil), "go.mnhosted.srv.order.GetInfoResponse")
	proto.RegisterType((*GetOrderListRequest)(nil), "go.mnhosted.srv.order.GetOrderListRequest")
	proto.RegisterType((*OrderItem)(nil), "go.mnhosted.srv.order.OrderItem")
	proto.RegisterType((*GetOrderListResponse)(nil), "go.mnhosted.srv.order.GetOrderListResponse")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xd1, 0x8e, 0x93, 0x40,
	0x14, 0x5d, 0xa0, 0xd0, 0xed, 0x5d, 0xb7, 0x9a, 0x71, 0xb7, 0x21, 0x3c, 0x98, 0x3a, 0x51, 0xd3,
	0xec, 0x2a, 0x89, 0xeb, 0x83, 0x6f, 0x26, 0x6a, 0xe3, 0x86, 0x68, 0x31, 0xc1, 0x6d, 0x4c, 0x7c,
	0xc3, 0xf6, 0xae, 0x92, 0x14, 0x06, 0x61, 0x6a, 0xdb, 0xbf, 0xf1, 0x4b, 0xfc, 0x07, 0x1f, 0xfc,
	0x1f, 0xc3, 0x30, 0x50, 0xa8, 0x8b, 0xd4, 0x37, 0x0e, 0x73, 0xe6, 0xce, 0x3d, 0x67, 0xce, 0x05,
	0x38, 0x62, 0xc9, 0x1c, 0x13, 0x3b, 0x4e, 0x18, 0x67, 0xe4, 0xf4, 0x0b, 0xb3, 0xc3, 0xe8, 0x2b,
	0x4b, 0x39, 0xce, 0xed, 0x34, 0xf9, 0x6e, 0x8b, 0x45, 0xfa, 0x43, 0x01, 0x70, 0x71, 0xe5, 0xe1,
	0xb7, 0x25, 0xa6, 0x9c, 0x0c, 0xc0, 0x58, 0xa6, 0x98, 0x38, 0x63, 0x53, 0x19, 0x2a, 0x23, 0xcd,
	0x93, 0x88, 0x58, 0x70, 0x38, 0x63, 0x41, 0x14, 0xf9, 0x21, 0x9a, 0xea, 0x50, 0x19, 0xf5, 0xbc,
	0x12, 0x67, 0x6b, 0x3c, 0x08, 0x91, 0x6f, 0x62, 0x34, 0xb5, 0xa1, 0x32, 0xd2, 0xbd, 0x12, 0x93,
	0x13, 0xd0, 0xe3, 0x24, 0x98, 0xa1, 0xd9, 0x11, 0x0b, 0x39, 0x20, 0x04, 0x3a, 0x7c, 0xed, 0x8c,
	0x4d, 0x5d, 0x54, 0x12, 0xcf, 0xc4, 0x84, 0x6e, 0x90, 0x7a, 0x18, 0xe1, 0xca, 0x34, 0x04, 0xb7,
	0x80, 0xf4, 0x39, 0x1c, 0x89, 0x0e, 0xd3, 0x98, 0x45, 0x29, 0x66, 0xc4, 0x04, 0xd3, 0x19, 0x9b,
	0xa3, 0xe8, 0x51, 0xf7, 0x0a, 0x48, 0xfa, 0xa0, 0x3a, 0x63, 0xd1, 0x9e, 0xe6, 0xa9, 0xce, 0x98,
	0xfe, 0x56, 0xe0, 0xf8, 0xe5, 0x22, 0x88, 0xfd, 0xcd, 0x9e, 0xf2, 0xdc, 0x1d, 0x79, 0x19, 0xce,
	0x24, 0x4c, 0xdc, 0xb7, 0xb8, 0x11, 0xda, 0x7a, 0x5e, 0x0e, 0xb2, 0x4a, 0x13, 0x57, 0xf0, 0x3b,
	0xe2, 0xb5, 0x44, 0x4d, 0xd2, 0xf8, 0xda, 0x89, 0xe6, 0xb8, 0x2e, 0xa4, 0x49, 0x58, 0x58, 0x77,
	0x95, 0x59, 0xd7, 0xdd, 0x5a, 0x97, 0xe1, 0xaa, 0x21, 0x87, 0x75, 0x43, 0x5e, 0x41, 0xbf, 0x90,
	0xd5, 0xea, 0xc9, 0x00, 0x8c, 0xd8, 0xdf, 0x4c, 0x93, 0x85, 0xd4, 0x25, 0x11, 0x7d, 0x03, 0x27,
	0xaf, 0x59, 0x74, 0x1d, 0x24, 0x61, 0xdd, 0x21, 0x13, 0xba, 0x22, 0x18, 0xa5, 0x45, 0x05, 0xdc,
	0x5e, 0xa5, 0x5a, 0xb9, 0x4a, 0xfa, 0x14, 0x4e, 0x77, 0xea, 0xb4, 0xb5, 0x44, 0x27, 0x70, 0x3c,
	0x8d, 0xe7, 0x3e, 0xc7, 0xe2, 0xcc, 0xfc, 0xde, 0x94, 0xe2, 0xde, 0xb6, 0x8e, 0xab, 0x3b, 0x8e,
	0xa7, 0xdc, 0xe7, 0xcb, 0x54, 0x86, 0x4c, 0x22, 0x7a, 0x06, 0xfd, 0xa2, 0x5c, 0xeb, 0xd1, 0x23,
	0xe8, 0x5f, 0x22, 0x77, 0xa2, 0x6b, 0xd6, 0x92, 0x08, 0x3a, 0x85, 0xdb, 0x25, 0xb3, 0xd5, 0xe4,
	0x3b, 0xa0, 0x45, 0xcb, 0x50, 0x1a, 0x93, 0x3d, 0x4a, 0xdb, 0xd9, 0x92, 0x8b, 0x66, 0x15, 0x4f,
	0x22, 0xfa, 0x04, 0xee, 0x5e, 0x22, 0x7f, 0x9f, 0x59, 0xfa, 0x2e, 0x48, 0x79, 0x5b, 0x17, 0xbf,
	0x14, 0xe8, 0x09, 0xb2, 0xc3, 0x31, 0xfc, 0xcb, 0xa7, 0xff, 0x4f, 0x6d, 0x35, 0x6f, 0x9d, 0x9d,
	0xbc, 0x95, 0xf7, 0xab, 0x8b, 0x8e, 0xe5, 0xa8, 0x0e, 0xc0, 0xf8, 0x90, 0xbb, 0x9e, 0x47, 0x57,
	0xa2, 0x6a, 0x3a, 0xbb, 0xb5, 0x74, 0x92, 0x7b, 0x00, 0xb3, 0x04, 0x7d, 0x8e, 0x57, 0x41, 0x88,
	0x22, 0xba, 0x3d, 0xaf, 0xf2, 0x86, 0xc6, 0x70, 0x52, 0xb7, 0xa0, 0xd5, 0xde, 0x17, 0xd0, 0x63,
	0x05, 0xdd, 0x54, 0x87, 0xda, 0xe8, 0xe8, 0x62, 0x68, 0xdf, 0xf8, 0x39, 0xb3, 0x4b, 0xb3, 0xbc,
	0xed, 0x96, 0x8b, 0x9f, 0x1d, 0xd0, 0x05, 0x22, 0x2e, 0x68, 0x2e, 0xae, 0xc8, 0xfd, 0x86, 0xdd,
	0xdb, 0x0f, 0xa1, 0x45, 0xff, 0x45, 0xc9, 0x3b, 0xa6, 0x07, 0xe4, 0x23, 0x18, 0x79, 0xec, 0xc9,
	0x83, 0x06, 0x7e, 0x6d, 0xba, 0xac, 0x87, 0x2d, 0xac, 0xb2, 0xf0, 0x02, 0x8e, 0x6b, 0x63, 0x45,
	0xce, 0x1b, 0x76, 0xde, 0x34, 0xc4, 0xd6, 0xe3, 0xfd, 0xc8, 0x55, 0x19, 0xf9, 0x08, 0x35, 0xca,
	0xa8, 0x0d, 0x6c, 0xa3, 0x8c, 0xfa, 0x1c, 0xd2, 0x03, 0xf2, 0x09, 0xba, 0x72, 0x8a, 0x48, 0xd3,
	0x9e, 0xfa, 0x3c, 0x5a, 0x8f, 0xda, 0x68, 0x65, 0xed, 0x00, 0x6e, 0x55, 0x73, 0x44, 0xce, 0x9a,
	0x77, 0xee, 0xce, 0x9b, 0x75, 0xbe, 0x17, 0xb7, 0x38, 0xea, 0xb3, 0x21, 0x7e, 0xa1, 0xcf, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x6d, 0xb6, 0xab, 0x51, 0x07, 0x00, 0x00,
}
