// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	middleware "github.com/go-kratos/kratos/v2/middleware"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
// context./http./middleware.
const _ = http1.SupportPackageIsVersion1

type CartHTTPServer interface {
	AddItem(context.Context, *AddItemReq) (*AddItemReply, error)

	DeleteCart(context.Context, *DeleteCartReq) (*DeleteCartReply, error)

	DeleteItem(context.Context, *DeleteItemReq) (*DeleteItemReply, error)

	GetCart(context.Context, *GetCartReq) (*GetCartReply, error)

	UpdateItem(context.Context, *UpdateItemReq) (*UpdateItemReply, error)
}

func RegisterCartHTTPServer(s http1.ServiceRegistrar, srv CartHTTPServer) {
	s.RegisterService(&_HTTP_Cart_serviceDesc, srv)
}

func _HTTP_Cart_GetCart_0(srv interface{}, ctx context.Context, req *http.Request, dec func(interface{}) error, m middleware.Middleware) (interface{}, error) {
	var in GetCartReq

	if err := http1.BindForm(req, &in); err != nil {
		return nil, err
	}

	h := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).GetCart(ctx, &in)
	}
	out, err := m(h)(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _HTTP_Cart_DeleteCart_0(srv interface{}, ctx context.Context, req *http.Request, dec func(interface{}) error, m middleware.Middleware) (interface{}, error) {
	var in DeleteCartReq

	if err := http1.BindForm(req, &in); err != nil {
		return nil, err
	}

	h := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).DeleteCart(ctx, &in)
	}
	out, err := m(h)(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _HTTP_Cart_AddItem_0(srv interface{}, ctx context.Context, req *http.Request, dec func(interface{}) error, m middleware.Middleware) (interface{}, error) {
	var in AddItemReq

	if err := dec(&in); err != nil {
		return nil, err
	}

	h := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).AddItem(ctx, &in)
	}
	out, err := m(h)(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _HTTP_Cart_UpdateItem_0(srv interface{}, ctx context.Context, req *http.Request, dec func(interface{}) error, m middleware.Middleware) (interface{}, error) {
	var in UpdateItemReq

	if err := http1.BindForm(req, &in); err != nil {
		return nil, err
	}

	h := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).UpdateItem(ctx, &in)
	}
	out, err := m(h)(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _HTTP_Cart_DeleteItem_0(srv interface{}, ctx context.Context, req *http.Request, dec func(interface{}) error, m middleware.Middleware) (interface{}, error) {
	var in DeleteItemReq

	if err := http1.BindForm(req, &in); err != nil {
		return nil, err
	}

	h := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).DeleteItem(ctx, &in)
	}
	out, err := m(h)(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _HTTP_Cart_serviceDesc = http1.ServiceDesc{
	ServiceName: "cart.v1.Cart",
	Methods: []http1.MethodDesc{

		{
			Path:    "/api/v1/cart/",
			Method:  "GET",
			Handler: _HTTP_Cart_GetCart_0,
		},

		{
			Path:    "/api/v1/cart/",
			Method:  "DELETE",
			Handler: _HTTP_Cart_DeleteCart_0,
		},

		{
			Path:    "/api/v1/cart/",
			Method:  "POST",
			Handler: _HTTP_Cart_AddItem_0,
		},

		{
			Path:    "/cart.v1.Cart/UpdateItem",
			Method:  "POST",
			Handler: _HTTP_Cart_UpdateItem_0,
		},

		{
			Path:    "/cart.v1.Cart/DeleteItem",
			Method:  "POST",
			Handler: _HTTP_Cart_DeleteItem_0,
		},
	},
	Metadata: "api/cart/v1/cart.proto",
}
