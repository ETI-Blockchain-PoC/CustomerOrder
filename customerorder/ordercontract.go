/*
 * SPDX-License-Identifier: Apache-2.0
 */

package customerorder

import (
	"fmt"

	"contractapi"
)

// Contract chaincode that defines
// the business logic for managing commercial
// paper
type Contract struct {
	contractapi.Contract
}

// Instantiate does nothing
func (c *Contract) Instantiate() {
	fmt.Println("Instantiated")
}

func (c *Contract) Create(ctx TransactionContextInterface, order CustomerOrder) (*CustomerOrder, error) {
	if order.Owner != order.CustomerNumber {
		return nil, fmt.Errorf("Order is not owned by this customer");
	}
	customerOrder := CustomerOrder{	OrderNumber: order.OrderNumber,
									CustomerNumber: order.CustomerNumber,
									RetailerNumber: order.RetailerNumber,
									ShippingNumber: order.ShippingNumber,
									CreatedDate: order.CreatedDate, 
									ReceivedDate: order.ReceivedDate,
									ToShippingDate: order.ToShippingDate
									InStockDate: order.InStockDate
									InDeliveryDate: order.InDeliveryDate
									CompletedDate:order.CompletedDate
									Owner: order.Owner
									DeliveryAddress: order.DeliveryAddress
									PaymentMethod: order.PaymentMethod
									Currency: order.Currency
									TotalOrderAmount: order.TotalOrderAmount
									TotalShippingCost: order.TotalShippingCost
									TotalDiscountAmount: order.TotalDiscountAmount
									TotalPaidAmount: order.TotalPaidAmount,
									OrderItems: order.OrderItems	
								}
	customerOrder.SetCreated();

	err := ctx.GetOrderList.AddOrder(&customerOrder)

	if err != nil {
		return nil, err
	}

	return &customerOrder, nil
}

func (c *Contract) Receive(ctx TransactionContextInterface, orderNumber string) (*CommercialPaper, error) {
	order, err := ctx.GetOrderList().GetOrder(orderNumber);

	if err != nil {
		return nil, err
	}


	if order.Owner != order.RetailerNumber {
		return nil, fmt.Errorf("Order is not owned by this retailer")
	}

	// if(order.GetState() == )

	order.SetReceived();

	order.Owner = order.RetailerNumber;

	err = ctx.GetOrderList().UpdateOrder(order);

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (c *Contract) DeliveryToShipping(ctx TransactionContextInterface, orderNumber string) (*CommercialPaper, error) {
	order, err := ctx.GetOrderList().GetOrder(orderNumber);

	if err != nil {
		return nil, err
	}

	if order.Owner != order.ShippingNumber {
		return nil, fmt.Errorf("Order is not owned by this shipping")
	}

	order.SetDelToShip();

	order.Owner = order.ShippingNumber;

	err = ctx.GetOrderList().UpdateOrder(order);

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (c *Contract) ReceiveToStock(ctx TransactionContextInterface, orderNumber string) (*CommercialPaper, error) {
	order, err := ctx.GetOrderList().GetOrder(orderNumber);

	if err != nil {
		return nil, err
	}

	if order.Owner != order.ShippingNumber {
		return nil, fmt.Errorf("Order is not owned by this shipping")
	}

	order.SetInstock();

	order.Owner = order.ShippingNumber;

	err = ctx.GetOrderList().UpdateOrder(order);

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (c *Contract) DeliverToCustomer(ctx TransactionContextInterface, orderNumber string) (*CommercialPaper, error) {
	order, err := ctx.GetOrderList().GetOrder(orderNumber);

	if err != nil {
		return nil, err
	}

	if order.Owner != order.ShippingNumber {
		return nil, fmt.Errorf("Order is not owned by this shipping")
	}

	order.SetDelivering();

	order.Owner = order.ShippingNumber;

	err = ctx.GetOrderList().UpdateOrder(order);

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (c *Contract) CompleteOrder(ctx TransactionContextInterface, orderNumber string) (*CommercialPaper, error) {
	order, err := ctx.GetOrderList().GetOrder(orderNumber);

	if err != nil {
		return nil, err
	}

	if order.Owner != order.ShippingNumber {
		return nil, fmt.Errorf("Order is not owned by this shipping")
	}

	order.SetCompleted();

	order.Owner = order.ShippingNumber;

	err = ctx.GetOrderList().UpdateOrder(order);

	if err != nil {
		return nil, err
	}

	return order, nil
}
