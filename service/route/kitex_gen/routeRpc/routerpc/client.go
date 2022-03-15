// Code generated by Kitex v0.2.0. DO NOT EDIT.

package routerpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"route/kitex_gen/routeRpc"
	"route/kitex_gen/rpc"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Login(ctx context.Context, Req *routeRpc.LoginReq, callOptions ...callopt.Option) (r *routeRpc.LoginReply, err error)
	GatePushMsg(ctx context.Context, Req *rpc.MessagePushReq, callOptions ...callopt.Option) (r *rpc.MessagePushReply, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRouteRpcClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRouteRpcClient struct {
	*kClient
}

func (p *kRouteRpcClient) Login(ctx context.Context, Req *routeRpc.LoginReq, callOptions ...callopt.Option) (r *routeRpc.LoginReply, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, Req)
}

func (p *kRouteRpcClient) GatePushMsg(ctx context.Context, Req *rpc.MessagePushReq, callOptions ...callopt.Option) (r *rpc.MessagePushReply, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GatePushMsg(ctx, Req)
}