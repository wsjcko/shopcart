package handler

import (
	"context"
	"github.com/wsjcko/shopcart/common"
	"github.com/wsjcko/shopcart/domain/model"
	"github.com/wsjcko/shopcart/domain/service"
	pb "github.com/wsjcko/shopcart/protobuf/pb"
)

type ShopCart struct {
	CartService service.ICartService
}

// AddCart 添加购物车
func (h *ShopCart) AddCart(ctx context.Context, request *pb.CartInfo, response *pb.ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(request, cart)
	response.CartId, err = h.CartService.AddCart(cart)
	return err
}

// CleanCart 清空购物车
func (h *ShopCart) CleanCart(ctx context.Context, request *pb.Clean, response *pb.Response) error {
	if err := h.CartService.CleanCart(request.UserId); err != nil {
		return err
	}
	response.Meg = "购物车清空成功"
	return nil
}

// Incr 添加购物车数量
func (h *ShopCart) Incr(ctx context.Context, request *pb.Item, response *pb.Response) error {
	if err := h.CartService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "购物车添加成功"
	return nil
}

// Decr 购物车减少商品数量
func (h *ShopCart) Decr(ctx context.Context, request *pb.Item, response *pb.Response) error {
	if err := h.CartService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "购物程减少成功"
	return nil
}

// DeleteItemByID 删除购物车
func (h *ShopCart) DeleteItemByID(ctx context.Context, request *pb.CartID, response *pb.Response) error {
	if err := h.CartService.DeleteCart(request.Id); err != nil {
		return err
	}
	response.Meg = "购物车删除成功"
	return nil
}

// GetAll 查询用户所有的购物车信息
func (h *ShopCart) GetAll(ctx context.Context, request *pb.CartFindAll, response *pb.CartAll) error {
	cartAll, err := h.CartService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}

	for _, v := range cartAll {
		cart := &pb.CartInfo{}
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cart)
	}
	return nil
}
