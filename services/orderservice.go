package services

import (
	"errors"
	"time"

	"github.com/cuffymate1/pos-api/models"
	"gorm.io/gorm"
)

func ListOrders(db *gorm.DB) ([]models.OrderResponse, error) {
	var orders []models.Order
	result := db.Preload("User").
		Preload("Payment").
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		Preload("Items.Toppings").
		Preload("Items.Toppings.Topping").
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	// Convert to response format
	var response []models.OrderResponse
	for _, order := range orders {
		response = append(response, convertToOrderResponse(order))
	}

	return response, nil
}

func GetOrders(db *gorm.DB, id uint) (*models.OrderResponse, error) {
	var order models.Order
	result := db.Preload("User").
		Preload("Payment").
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		Preload("Items.Toppings").
		Preload("Items.Toppings.Topping").
		First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}

	response := convertToOrderResponse(order)
	return &response, nil
}

func CreateOrder(db *gorm.DB, input *models.Order) (string, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		// Calculate total from items
		var total float32
		for _, item := range input.Items {
			total += item.Price * float32(item.Quantity)
			// Add topping prices
			for _, topping := range item.Toppings {
				var toppingModel models.Topping
				if err := tx.First(&toppingModel, topping.ToppingID).Error; err != nil {
					return err
				}
				total += toppingModel.Price * float32(item.Quantity)
			}
		}
		input.Total = total

		// Create Order without items first
		orderItems := input.Items
		input.Items = nil
		if err := tx.Create(&input).Error; err != nil {
			return err
		}

		// Create OrderItems and OrderItemToppings
		for _, item := range orderItems {
			orderItem := models.OrderItem{
				OrderID:   input.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
			}
			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}

			// Create toppings for this item
			for _, topping := range item.Toppings {
				orderItemTopping := models.OrderItemTopping{
					OrderItemID: orderItem.ID,
					ToppingID:   topping.ToppingID,
				}
				if err := tx.Create(&orderItemTopping).Error; err != nil {
					return err
				}
			}
		}

		// Create Payment if exists
		if input.Payment != nil {
			payment := &models.Payment{
				OrderID:    input.ID,
				Method:     input.Payment.Method,
				AmountPaid: input.Payment.AmountPaid,
				Change:     input.Payment.AmountPaid - input.Total,
				PaidAt:     time.Now(),
			}
			if err := tx.Create(payment).Error; err != nil {
				return err
			}

			// Update IsPaid
			if err := tx.Model(&models.Order{}).Where("id = ?", input.ID).Update("is_paid", true).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return "Something Went Wrong!", err
	}

	return "Create Order Successful", nil
}

func UpdateOrder(db *gorm.DB, order *models.Order) (string, error) {
	result := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order).Error
	if result != nil {
		return "Something Went Wrong!", result
	}

	return "Update Order Successful", nil
}

func DeleteOrder(db *gorm.DB, id uint) (string, error) {
	var order *models.Order
	if err := db.Preload("Items.Toppings").First(&order).Error; err != nil {
		return "Order Not Found!", err
	}

	result := db.Delete(&order, id)
	if result.Error != nil {
		return "Something Went Wrong!", result.Error
	}
	return "Delete Order Successful", nil
}

func PayOrder(db *gorm.DB, orderID uint, method string, amount float32) error {
	var order models.Order
	if err := db.First(&order, orderID).Error; err != nil {
		return err
	}

	if order.IsPaid {
		return errors.New("order already paid")
	}

	change := amount - order.Total
	if change < 0 {
		return errors.New("insufficient amount paid")
	}

	payment := models.Payment{
		OrderID:    order.ID,
		Method:     method,
		AmountPaid: amount,
		Change:     change,
		PaidAt:     time.Now(),
	}

	if err := db.Create(&payment).Error; err != nil {
		return err
	}

	// mark order as paid
	order.IsPaid = true
	return db.Save(&order).Error
}

func convertToOrderResponse(order models.Order) models.OrderResponse {
	// Convert items
	var items []models.OrderItemResponse
	for _, item := range order.Items {
		// Convert toppings
		var toppings []models.ToppingBrief
		for _, t := range item.Toppings {
			toppings = append(toppings, models.ToppingBrief{
				ID:    t.Topping.ID,
				Name:  t.Topping.Name,
				Price: t.Topping.Price,
			})
		}

		// Convert product
		product := models.ProductBrief{
			ID:          item.Product.ID,
			Name:        item.Product.Name,
			Description: item.Product.Description,
			Price:       item.Product.Price,
			Category: models.CategoryBrief{
				ID:   item.Product.Category.ID,
				Name: item.Product.Category.Name,
			},
		}

		items = append(items, models.OrderItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Product:   product,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Toppings:  toppings,
		})
	}

	// Convert payment if exists
	var payment *models.PaymentBrief
	if order.Payment != nil {
		payment = &models.PaymentBrief{
			Method:     order.Payment.Method,
			AmountPaid: order.Payment.AmountPaid,
			Change:     order.Payment.Change,
			PaidAt:     order.Payment.PaidAt.Format(time.RFC3339),
		}
	}

	return models.OrderResponse{
		ID:        order.ID,
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		UserID:    order.UserID,
		User: models.UserBrief{
			ID:       order.User.ID,
			Username: order.User.Username,
			Fullname: order.User.Fullname,
			Role:     order.User.Role,
		},
		Total:   order.Total,
		IsPaid:  order.IsPaid,
		Payment: payment,
		Items:   items,
	}
}
