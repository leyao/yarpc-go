// Code generated by protoc-gen-yarpc-go
// source: internal/examples/protobuf/examplepb/example.proto
// DO NOT EDIT!

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package examplepb

import (
	"context"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/yarpcproto"
)

// KeyValueYARPCClient is the YARPC client-side interface for the KeyValue service.
type KeyValueYARPCClient interface {
	GetValue(context.Context, *GetValueRequest, ...yarpc.CallOption) (*GetValueResponse, error)
	SetValue(context.Context, *SetValueRequest, ...yarpc.CallOption) (*SetValueResponse, error)
}

// NewKeyValueYARPCClient builds a new YARPC client for the KeyValue service.
func NewKeyValueYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) KeyValueYARPCClient {
	return &_KeyValueYARPCCaller{protobuf.NewClient(
		protobuf.ClientParams{
			ServiceName:  "uber.yarpc.internal.examples.protobuf.example.KeyValue",
			ClientConfig: clientConfig,
			Options:      options,
		},
	)}
}

// KeyValueYARPCServer is the YARPC server-side interface for the KeyValue service.
type KeyValueYARPCServer interface {
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)
}

// BuildKeyValueYARPCProcedures prepares an implementation of the KeyValue service for YARPC registration.
func BuildKeyValueYARPCProcedures(server KeyValueYARPCServer) []transport.Procedure {
	handler := &_KeyValueYARPCHandler{server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "uber.yarpc.internal.examples.protobuf.example.KeyValue",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "GetValue",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:     handler.GetValue,
							NewRequest: newKeyValue_GetValueYARPCRequest,
						},
					),
				},
				{
					MethodName: "SetValue",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:     handler.SetValue,
							NewRequest: newKeyValue_SetValueYARPCRequest,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
		},
	)
}

type _KeyValueYARPCCaller struct {
	client protobuf.Client
}

func (c *_KeyValueYARPCCaller) GetValue(ctx context.Context, request *GetValueRequest, options ...yarpc.CallOption) (*GetValueResponse, error) {
	responseMessage, err := c.client.Call(ctx, "GetValue", request, newKeyValue_GetValueYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*GetValueResponse)
	if !ok {
		return nil, protobuf.CastError(emptyKeyValue_GetValueYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_KeyValueYARPCCaller) SetValue(ctx context.Context, request *SetValueRequest, options ...yarpc.CallOption) (*SetValueResponse, error) {
	responseMessage, err := c.client.Call(ctx, "SetValue", request, newKeyValue_SetValueYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*SetValueResponse)
	if !ok {
		return nil, protobuf.CastError(emptyKeyValue_SetValueYARPCResponse, responseMessage)
	}
	return response, err
}

type _KeyValueYARPCHandler struct {
	server KeyValueYARPCServer
}

func (h *_KeyValueYARPCHandler) GetValue(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *GetValueRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*GetValueRequest)
		if !ok {
			return nil, protobuf.CastError(emptyKeyValue_GetValueYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.GetValue(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func (h *_KeyValueYARPCHandler) SetValue(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *SetValueRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*SetValueRequest)
		if !ok {
			return nil, protobuf.CastError(emptyKeyValue_SetValueYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.SetValue(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newKeyValue_GetValueYARPCRequest() proto.Message {
	return &GetValueRequest{}
}

func newKeyValue_GetValueYARPCResponse() proto.Message {
	return &GetValueResponse{}
}

func newKeyValue_SetValueYARPCRequest() proto.Message {
	return &SetValueRequest{}
}

func newKeyValue_SetValueYARPCResponse() proto.Message {
	return &SetValueResponse{}
}

var (
	emptyKeyValue_GetValueYARPCRequest  = &GetValueRequest{}
	emptyKeyValue_GetValueYARPCResponse = &GetValueResponse{}
	emptyKeyValue_SetValueYARPCRequest  = &SetValueRequest{}
	emptyKeyValue_SetValueYARPCResponse = &SetValueResponse{}
)

// SinkYARPCClient is the YARPC client-side interface for the Sink service.
type SinkYARPCClient interface {
	Fire(context.Context, *FireRequest, ...yarpc.CallOption) (yarpc.Ack, error)
}

// NewSinkYARPCClient builds a new YARPC client for the Sink service.
func NewSinkYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) SinkYARPCClient {
	return &_SinkYARPCCaller{protobuf.NewClient(
		protobuf.ClientParams{
			ServiceName:  "uber.yarpc.internal.examples.protobuf.example.Sink",
			ClientConfig: clientConfig,
			Options:      options,
		},
	)}
}

// SinkYARPCServer is the YARPC server-side interface for the Sink service.
type SinkYARPCServer interface {
	Fire(context.Context, *FireRequest) error
}

// BuildSinkYARPCProcedures prepares an implementation of the Sink service for YARPC registration.
func BuildSinkYARPCProcedures(server SinkYARPCServer) []transport.Procedure {
	handler := &_SinkYARPCHandler{server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName:        "uber.yarpc.internal.examples.protobuf.example.Sink",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{
				{
					MethodName: "Fire",
					Handler: protobuf.NewOnewayHandler(
						protobuf.OnewayHandlerParams{
							Handle:     handler.Fire,
							NewRequest: newSink_FireYARPCRequest,
						},
					),
				},
			},
		},
	)
}

type _SinkYARPCCaller struct {
	client protobuf.Client
}

func (c *_SinkYARPCCaller) Fire(ctx context.Context, request *FireRequest, options ...yarpc.CallOption) (yarpc.Ack, error) {
	return c.client.CallOneway(ctx, "Fire", request, options...)
}

type _SinkYARPCHandler struct {
	server SinkYARPCServer
}

func (h *_SinkYARPCHandler) Fire(ctx context.Context, requestMessage proto.Message) error {
	var request *FireRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*FireRequest)
		if !ok {
			return protobuf.CastError(emptySink_FireYARPCRequest, requestMessage)
		}
	}
	return h.server.Fire(ctx, request)
}

func newSink_FireYARPCRequest() proto.Message {
	return &FireRequest{}
}

func newSink_FireYARPCResponse() proto.Message {
	return &yarpcproto.Oneway{}
}

var (
	emptySink_FireYARPCRequest  = &FireRequest{}
	emptySink_FireYARPCResponse = &yarpcproto.Oneway{}
)

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) KeyValueYARPCClient {
			return NewKeyValueYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) SinkYARPCClient {
			return NewSinkYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}
