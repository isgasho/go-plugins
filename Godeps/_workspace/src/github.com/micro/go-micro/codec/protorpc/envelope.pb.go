// Code generated by protoc-gen-go.
// source: envelope.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	envelope.proto

It has these top-level messages:
	Request
	Response
*/
package protorpc

import proto "github.com/micro/go-plugins/Godeps/_workspace/src/github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	ServiceMethod    *string `protobuf:"bytes,1,opt,name=service_method" json:"service_method,omitempty"`
	Seq              *uint64 `protobuf:"fixed64,2,opt,name=seq" json:"seq,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetServiceMethod() string {
	if m != nil && m.ServiceMethod != nil {
		return *m.ServiceMethod
	}
	return ""
}

func (m *Request) GetSeq() uint64 {
	if m != nil && m.Seq != nil {
		return *m.Seq
	}
	return 0
}

type Response struct {
	ServiceMethod    *string `protobuf:"bytes,1,opt,name=service_method" json:"service_method,omitempty"`
	Seq              *uint64 `protobuf:"fixed64,2,opt,name=seq" json:"seq,omitempty"`
	Error            *string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetServiceMethod() string {
	if m != nil && m.ServiceMethod != nil {
		return *m.ServiceMethod
	}
	return ""
}

func (m *Response) GetSeq() uint64 {
	if m != nil && m.Seq != nil {
		return *m.Seq
	}
	return 0
}

func (m *Response) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func init() {
}
