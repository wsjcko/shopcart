// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: shopcart.proto

package go_micro_service_shop_cart

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ShopCart service

func NewShopCartEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ShopCart service

type ShopCartService interface {
	AddCart(ctx context.Context, in *CartInfo, opts ...client.CallOption) (*ResponseAdd, error)
	CleanCart(ctx context.Context, in *Clean, opts ...client.CallOption) (*Response, error)
	Incr(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error)
	Decr(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error)
	DeleteItemByID(ctx context.Context, in *CartID, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *CartFindAll, opts ...client.CallOption) (*CartAll, error)
}

type shopCartService struct {
	c    client.Client
	name string
}

func NewShopCartService(name string, c client.Client) ShopCartService {
	return &shopCartService{
		c:    c,
		name: name,
	}
}

func (c *shopCartService) AddCart(ctx context.Context, in *CartInfo, opts ...client.CallOption) (*ResponseAdd, error) {
	req := c.c.NewRequest(c.name, "ShopCart.AddCart", in)
	out := new(ResponseAdd)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopCartService) CleanCart(ctx context.Context, in *Clean, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ShopCart.CleanCart", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopCartService) Incr(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ShopCart.Incr", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopCartService) Decr(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ShopCart.Decr", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopCartService) DeleteItemByID(ctx context.Context, in *CartID, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ShopCart.DeleteItemByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopCartService) GetAll(ctx context.Context, in *CartFindAll, opts ...client.CallOption) (*CartAll, error) {
	req := c.c.NewRequest(c.name, "ShopCart.GetAll", in)
	out := new(CartAll)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShopCart service

type ShopCartHandler interface {
	AddCart(context.Context, *CartInfo, *ResponseAdd) error
	CleanCart(context.Context, *Clean, *Response) error
	Incr(context.Context, *Item, *Response) error
	Decr(context.Context, *Item, *Response) error
	DeleteItemByID(context.Context, *CartID, *Response) error
	GetAll(context.Context, *CartFindAll, *CartAll) error
}

func RegisterShopCartHandler(s server.Server, hdlr ShopCartHandler, opts ...server.HandlerOption) error {
	type shopCart interface {
		AddCart(ctx context.Context, in *CartInfo, out *ResponseAdd) error
		CleanCart(ctx context.Context, in *Clean, out *Response) error
		Incr(ctx context.Context, in *Item, out *Response) error
		Decr(ctx context.Context, in *Item, out *Response) error
		DeleteItemByID(ctx context.Context, in *CartID, out *Response) error
		GetAll(ctx context.Context, in *CartFindAll, out *CartAll) error
	}
	type ShopCart struct {
		shopCart
	}
	h := &shopCartHandler{hdlr}
	return s.Handle(s.NewHandler(&ShopCart{h}, opts...))
}

type shopCartHandler struct {
	ShopCartHandler
}

func (h *shopCartHandler) AddCart(ctx context.Context, in *CartInfo, out *ResponseAdd) error {
	return h.ShopCartHandler.AddCart(ctx, in, out)
}

func (h *shopCartHandler) CleanCart(ctx context.Context, in *Clean, out *Response) error {
	return h.ShopCartHandler.CleanCart(ctx, in, out)
}

func (h *shopCartHandler) Incr(ctx context.Context, in *Item, out *Response) error {
	return h.ShopCartHandler.Incr(ctx, in, out)
}

func (h *shopCartHandler) Decr(ctx context.Context, in *Item, out *Response) error {
	return h.ShopCartHandler.Decr(ctx, in, out)
}

func (h *shopCartHandler) DeleteItemByID(ctx context.Context, in *CartID, out *Response) error {
	return h.ShopCartHandler.DeleteItemByID(ctx, in, out)
}

func (h *shopCartHandler) GetAll(ctx context.Context, in *CartFindAll, out *CartAll) error {
	return h.ShopCartHandler.GetAll(ctx, in, out)
}
