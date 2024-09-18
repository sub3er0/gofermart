package service

import (
	"gofermart/internal/repository"
)

type OrderService struct {
	OrderRepository *repository.OrderRepository
}

func (or *OrderService) IsOrderExist(orderNumber string) bool {
	isExist, err := or.OrderRepository.IsOrderExist(orderNumber)

	if err != nil {
		return false
	}

	return isExist
}

func (or *OrderService) SaveOrder(orderNumber string, userID int, accrual int) error {
	err := or.OrderRepository.SaveOrder(orderNumber, userID, accrual)

	if err != nil {
		return err
	}

	return nil
}

func (or *OrderService) GetUserOrders(userID int) ([]repository.OrderData, error) {
	orderData, err := or.OrderRepository.GetUserOrders(userID)

	if err != nil {
		return orderData, err
	}

	return orderData, nil
}

func (or *OrderService) GetUserBalance(userID int) (repository.UserBalance, error) {
	userBalance, err := or.OrderRepository.GetUserBalance(userID)

	if err != nil {
		return userBalance, err
	}

	return userBalance, nil
}
