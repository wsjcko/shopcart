package service

import (
	"github.com/wsjcko/shopcart/domain/model"
	"github.com/wsjcko/shopcart/domain/repository"
)

type ICartService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	DecrNum(int64, int64) error
	IncrNum(int64, int64) error
}

// NewCartDataService 创建
func NewCartService(cartRepository repository.ICartRepository) ICartService {
	return &CartService{cartRepository}
}

type CartService struct {
	CartRepository repository.ICartRepository
}

// AddCart 插入
func (u *CartService) AddCart(cart *model.Cart) (int64, error) {
	return u.CartRepository.CreateCart(cart)
}

// DeleteCart 删除
func (u *CartService) DeleteCart(cartID int64) error {
	return u.CartRepository.DeleteCartByID(cartID)
}

// UpdateCart 更新
func (u *CartService) UpdateCart(cart *model.Cart) error {
	return u.CartRepository.UpdateCart(cart)
}

// FindCartByID 查找
func (u *CartService) FindCartByID(cartID int64) (*model.Cart, error) {
	return u.CartRepository.FindCartByID(cartID)
}

// FindAllCart 查找
func (u *CartService) FindAllCart(userID int64) ([]model.Cart, error) {
	return u.CartRepository.FindAll(userID)
}

func (u *CartService) CleanCart(userID int64) error {
	return u.CartRepository.CleanCart(userID)
}

func (u *CartService) DecrNum(cartID int64, num int64) error {
	return u.CartRepository.DecrNum(cartID, num)
}

func (u *CartService) IncrNum(cartID int64, num int64) error {
	return u.CartRepository.IncrNum(cartID, num)
}
