// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: spec.proto

/*
Package pubsub is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pubsub

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_Receiver_Receive_0(ctx context.Context, marshaler runtime.Marshaler, client ReceiverClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ReceivedMessage
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Receive(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Receiver_Receive_0(ctx context.Context, marshaler runtime.Marshaler, server ReceiverServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ReceivedMessage
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Receive(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterReceiverHandlerServer registers the http handlers for service Receiver to "mux".
// UnaryRPC     :call ReceiverServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterReceiverHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ReceiverServer) error {

	mux.Handle("POST", pattern_Receiver_Receive_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Receiver_Receive_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Receiver_Receive_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterReceiverHandlerFromEndpoint is same as RegisterReceiverHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterReceiverHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterReceiverHandler(ctx, mux, conn)
}

// RegisterReceiverHandler registers the http handlers for service Receiver to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterReceiverHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterReceiverHandlerClient(ctx, mux, NewReceiverClient(conn))
}

// RegisterReceiverHandlerClient registers the http handlers for service Receiver
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ReceiverClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ReceiverClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ReceiverClient" to call the correct interceptors.
func RegisterReceiverHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ReceiverClient) error {

	mux.Handle("POST", pattern_Receiver_Receive_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Receiver_Receive_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Receiver_Receive_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Receiver_Receive_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0}, []string{"topics"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_Receiver_Receive_0 = runtime.ForwardResponseMessage
)
