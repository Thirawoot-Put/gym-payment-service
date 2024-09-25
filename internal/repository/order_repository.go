package repository

import (
	"github.com/Thirawoot-Put/event-ticketing/payment-service/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *domain.Orders) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) Update(order *domain.Orders) error {
	return r.db.Model(&order).Updates(order).Error
	// return r.db.Save(order).Error
}

func (r *OrderRepository) GetByUserId(userId uint) ([]domain.Orders, error) {
	var orders []domain.Orders
	err := r.db.Find(&orders, r.db.Where("userId = ?", userId)).Error
	return orders, err
}
