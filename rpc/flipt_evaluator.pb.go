// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flipt_evaluator.proto

package flipt

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("flipt_evaluator.proto", fileDescriptor_ef02beb77a9d4914)
}

var fileDescriptor_ef02beb77a9d4914 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x25, 0x85, 0xaf, 0x94, 0x2d, 0x5f, 0x29, 0x4b, 0x45, 0x0d, 0x58, 0x42, 0x7b, 0x13, 0xdb,
	0x85, 0x78, 0xeb, 0x2d, 0x45, 0x05, 0x0f, 0x42, 0xe9, 0x49, 0xbc, 0x94, 0x6d, 0x58, 0x93, 0xb5,
	0xe9, 0xce, 0xba, 0x3b, 0xd1, 0xff, 0x91, 0xb3, 0x77, 0x7f, 0x96, 0x7f, 0x45, 0xba, 0x9b, 0xb4,
	0x05, 0xf5, 0x36, 0xef, 0xbd, 0x99, 0xf7, 0x66, 0x18, 0x72, 0xf2, 0x5c, 0x48, 0x8d, 0x2b, 0xf1,
	0xc6, 0x8b, 0x92, 0x23, 0x98, 0xa9, 0x36, 0x80, 0x40, 0xff, 0x39, 0x3a, 0xbc, 0x72, 0x28, 0x9d,
	0x64, 0x42, 0x4d, 0xec, 0x3b, 0xcf, 0x32, 0x61, 0x18, 0x68, 0x94, 0xa0, 0x2c, 0xe3, 0x4a, 0x01,
	0x72, 0x57, 0xfb, 0xa1, 0xb0, 0xeb, 0x86, 0x3c, 0x88, 0xbf, 0x02, 0xd2, 0xbb, 0xdb, 0xe1, 0xdb,
	0xc6, 0x9a, 0x3e, 0x92, 0x4e, 0x0d, 0x04, 0x3d, 0x9b, 0xfa, 0xe6, 0x9a, 0x90, 0xa0, 0x96, 0xe2,
	0xb5, 0x14, 0x16, 0xc3, 0xf3, 0x5f, 0x14, 0xab, 0x41, 0x59, 0x31, 0x3a, 0xad, 0x92, 0x01, 0xe9,
	0xd4, 0xbb, 0x8a, 0xf0, 0xe0, 0x66, 0xc9, 0xff, 0x39, 0xc7, 0x34, 0xdf, 0x13, 0x17, 0xb5, 0xc9,
	0x31, 0x7b, 0x94, 0x31, 0xfc, 0x4b, 0xae, 0x83, 0xc6, 0x55, 0x12, 0x91, 0xde, 0x7a, 0xa7, 0xae,
	0xf6, 0x71, 0x3d, 0xd7, 0x1d, 0x35, 0x19, 0xf3, 0xcf, 0xa0, 0x4a, 0x3e, 0x02, 0xba, 0x24, 0x03,
	0x77, 0x67, 0x74, 0xf0, 0x89, 0x92, 0xc5, 0xfd, 0x28, 0x26, 0xdd, 0x07, 0x6e, 0x36, 0xd1, 0x22,
	0x17, 0x85, 0xb6, 0x74, 0x9c, 0x23, 0x6a, 0x3b, 0x63, 0x2c, 0x93, 0x98, 0x97, 0xeb, 0x69, 0x0a,
	0x5b, 0xb6, 0xe5, 0x66, 0xa3, 0x9d, 0xcc, 0xdc, 0x3e, 0x71, 0xbb, 0xe0, 0x28, 0x2c, 0x5e, 0xb6,
	0x82, 0x56, 0xdc, 0xe7, 0x5a, 0x17, 0x32, 0x75, 0x86, 0xec, 0xc5, 0x82, 0x9a, 0xfd, 0x60, 0xcc,
	0x90, 0x10, 0x9f, 0x7d, 0x03, 0xa9, 0xa5, 0xfd, 0x26, 0xc2, 0x1f, 0x27, 0xe1, 0xc9, 0xbf, 0x71,
	0xdd, 0x76, 0x2f, 0xb9, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x49, 0x17, 0x9b, 0xed, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FliptEvaluatorClient is the client API for FliptEvaluator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FliptEvaluatorClient interface {
	Evaluate(ctx context.Context, in *EvaluationRequest, opts ...grpc.CallOption) (*EvaluationResponse, error)
	BatchEvaluate(ctx context.Context, in *BatchEvaluationRequest, opts ...grpc.CallOption) (*BatchEvaluationResponse, error)
}

type fliptEvaluatorClient struct {
	cc grpc.ClientConnInterface
}

func NewFliptEvaluatorClient(cc grpc.ClientConnInterface) FliptEvaluatorClient {
	return &fliptEvaluatorClient{cc}
}

func (c *fliptEvaluatorClient) Evaluate(ctx context.Context, in *EvaluationRequest, opts ...grpc.CallOption) (*EvaluationResponse, error) {
	out := new(EvaluationResponse)
	err := c.cc.Invoke(ctx, "/flipt.FliptEvaluator/Evaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fliptEvaluatorClient) BatchEvaluate(ctx context.Context, in *BatchEvaluationRequest, opts ...grpc.CallOption) (*BatchEvaluationResponse, error) {
	out := new(BatchEvaluationResponse)
	err := c.cc.Invoke(ctx, "/flipt.FliptEvaluator/BatchEvaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FliptEvaluatorServer is the server API for FliptEvaluator service.
type FliptEvaluatorServer interface {
	Evaluate(context.Context, *EvaluationRequest) (*EvaluationResponse, error)
	BatchEvaluate(context.Context, *BatchEvaluationRequest) (*BatchEvaluationResponse, error)
}

// UnimplementedFliptEvaluatorServer can be embedded to have forward compatible implementations.
type UnimplementedFliptEvaluatorServer struct {
}

func (*UnimplementedFliptEvaluatorServer) Evaluate(ctx context.Context, req *EvaluationRequest) (*EvaluationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evaluate not implemented")
}
func (*UnimplementedFliptEvaluatorServer) BatchEvaluate(ctx context.Context, req *BatchEvaluationRequest) (*BatchEvaluationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchEvaluate not implemented")
}

func RegisterFliptEvaluatorServer(s *grpc.Server, srv FliptEvaluatorServer) {
	s.RegisterService(&_FliptEvaluator_serviceDesc, srv)
}

func _FliptEvaluator_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FliptEvaluatorServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flipt.FliptEvaluator/Evaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FliptEvaluatorServer).Evaluate(ctx, req.(*EvaluationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FliptEvaluator_BatchEvaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchEvaluationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FliptEvaluatorServer).BatchEvaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flipt.FliptEvaluator/BatchEvaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FliptEvaluatorServer).BatchEvaluate(ctx, req.(*BatchEvaluationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FliptEvaluator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flipt.FliptEvaluator",
	HandlerType: (*FliptEvaluatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Evaluate",
			Handler:    _FliptEvaluator_Evaluate_Handler,
		},
		{
			MethodName: "BatchEvaluate",
			Handler:    _FliptEvaluator_BatchEvaluate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flipt_evaluator.proto",
}
