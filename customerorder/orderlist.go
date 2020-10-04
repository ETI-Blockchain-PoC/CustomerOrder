/*
 * SPDX-License-Identifier: Apache-2.0
 */

package customerorder

import ledgerapi "ledger-api"

// ListInterface defines functionality needed
// to interact with the world state on behalf
// of a customer order
// KhanhTNG
type ListInterface interface {
	AddOrder(*CustomerOrder) error
	GetOrder(string, string) (*CustomerOrder, error)
	UpdateOrder(*CustomerOrder) error
}

// KhanhTNg
type list struct {
	stateList ledgerapi.StateListInterface
}

// KhanhTNG
func (cpl *list) AddOrder(order *CustomerOrder) error {
	return cpl.stateList.AddState(order)
}

// KhanhTNG
func (cpl *list) GetOrder(orderNumber string) (*CustomerOrder, error) {
	cp := new(CustomerOrder)

	err := cpl.stateList.GetState(CreateCustomerOrderKey(orderNumber), cp)

	if err != nil {
		return nil, err
	}

	return cp, nil
}

// KhanhTNG
func (cpl *list) UpdateOrder(order *CustomerOrder) error {
	return cpl.stateList.UpdateState(order)
}

// KhanhTNG - NewList create a new list from context
func newList(ctx TransactionContextInterface) *list {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.shouldbeconsider"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return Deserialize(bytes, state.(*CustomerOrder))
	}

	list := new(list)
	list.stateList = stateList

	return list
}
