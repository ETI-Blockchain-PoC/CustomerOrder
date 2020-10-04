/*
 * SPDX-License-Identifier: Apache-2.0
 */

package customerorder

import (
	"encoding/json"
	"fmt"

	// KhanhTNG - Should be consider again.
	ledgerapi "ledgerapi"

	// KhanhTNG - For time/date lib
	"fmt"
	"time"
)

// State enum for Customer Order state property
type State uint

const (
	// ISSUED state for when a paper has been issued - KhanhTNG - iota begin = 0, should be start with 1 to avoid issues in initial structure 
	// KhanhTNG - IF don't want to start with 1, we could start with UNKNOW State and increasing 1 for other state
	// iota = 0, CREATE = 1.
	CREATED State = iota + 1
	// iota = 1, RECEIVED = 2.
	RECEIVED
	// iota = 2, DELTOSHIP = 3.
	DELTOSHIP
	// iota = 3, INSTOCK = 4.
	INSTOCK
	// iota = 4, DELIVERING = 5.
	DELIVERING
	// iota = 5, COMPLETED = 6.
	COMPLETED 
)

func (state State) String() string {
	names := []string{"CREATED", "RECEIVED", "DELTOSHIP", "INSTOCK", "DELIVERING", "COMPLETED"}

	// KhanhTNG - Check if state over of enum range, return unknow
	if state < CREATED || state > COMPLETED {
		return "UNKNOWN STATE"
	}

	return names[state-1]
}

// KhanhTNG - CreateCustomerOrderKey creates a key for customer order
func CreateCustomerOrderKey(orderNumber string) string {
	return ledgerapi.MakeKey(orderNumber)
}

// KhanhTNG - Used for managing the fact status is private but want it in world state
// KhanhTNG - Definition for json customer order contained customer order as well. Json is a tag that allow to parse data from JSON to structure
type customerOrderAlias CustomerOrder
type jsonCustomerOrder struct {
	*customerOrderAlias
	State State  `json:"currentState"`
	// Class
	Class string `json:"class"`
	// Get key from State ( From ledge API)
	Key   string `json:"key"`
}

// KhanhTNG - Structure for OrderItem
type CustomerOrderItems struct {
	ProductNumber		string `json:"ProductNumber"`
	ProductName			string `json:"ProductName"`
	RetailerNumber		string `json:"RetailerNumber"`
	RetailerName		string `json:"RetailerName"`
	ShippingNumber		string `json:"ShippingNumber"`
	ShippingName		string `json:"ShippingName"`
	ShippingCondition	string `json:"ShippingCondition"`
	ShippingCost		string `json:"ShippingCost"`
	ProductQuantity		string `json:"ProductQuantity"`
	ProductPrice		string `json:"ProductPrice"`
	ProductAmount		string `json:"ProductAmount"`
}

// KhanhTNG - CustomerOrder defines a customer order
type CustomerOrder struct {
	// Header
	OrderNumber			string `json:"OrderNumber"`
	CustomerNumber		string `json:"CustomerNumber"`
	RetailerNumber		string `json:"RetailerNumber"`
	ShippingNumber		string `json:"ShippingNumber"`
	CreatedDate			string `json:"CreatedDate"`
	ReceivedDate		string `json:"ReceivedDate"`
	ToShippingDate		string `json:"ToShippingDate"`
	InStockDate			string `json:"InStockDate"`
	InDeliveryDate		string `json:"InDeliveryDate"`
	CompletedDate		string `json:"CompletedDate"`
	Owner				string `json:"Owner"`
	DeliveryAddress		string `json:"DeliveryAddress"`
	PaymentMethod		string `json:"PaymentMethod"`
	Currency			string `json:"Currency"`
	TotalOrderAmount	string `json:"TotalOrderAmount"`
	TotalShippingCost	string `json:"TotalShippingCost"`
	TotalDiscountAmount	string `json:"TotalDiscountAmount"`
	TotalPaidAmount		string `json:"TotalPaidAmount"`
	state            	State  `metadata:"currentState"`
	class            	string `metadata:"class"`
	key              	string `metadata:"key"`
	//Items
	OrderItems 			[]CustomerOrderItems
}

// KhanhTNG - UnmarshalJSON special handler for managing JSON marshalling - Unmarshall to convert from byte to jcp.
func (cp *CustomerOrder) UnmarshalJSON(data []byte) error {
	jcp := jsonCustomerOrder{customerOrderAlias: (*customerOrderAlias)(cp)}

	err := json.Unmarshal(data, &jcp)

	if err != nil {
		return err
	}

	cp.state = jcp.State
	// Added for party number
	cp.CustomerNumber = jcp.CustomerNumber 
	cp.RetailerNumber = jcp.RetailerNumber
	cp.ShippingNumber = jcp.ShippingNumber

	return nil
}

// KhanhTNG - MarshalJSON special handler for managing JSON marshalling - Return from jcp
func (cp CustomerOrder) MarshalJSON() ([]byte, error) {
	jcp := jsonCustomerOrder{customerOrderAlias: (*customerOrderAlias)(&cp), State: cp.state, Class: "org.shouldbeconsider", Key: ledgerapi.MakeKey(cp.OrderNumber)}

	return json.Marshal(&jcp)
}

// KhanhTNG - GetState returns the state
func (cp *CustomerOrder) GetState() State {
	return cp.state
} 

// KhanhTNG - SetCreated returns the state to issued
func (cp *CustomerOrder) SetCreated() {
	cp.state = CREATED
	// Add some new value
	//cp.Owner = ''
	cp.CreatedDate = time.Now()
}

// KhanhTNG - SetReceived returns the state to issued
func (cp *CustomerOrder) SetReceived() {
	cp.state = RECEIVED
	// Add some new value
	//cp.Owner = ''
	cp.ReceivedDate = time.Now()
}

// KhanhTNG - SetDelToShip returns the state to issued
func (cp *CustomerOrder) SetDelToShip() {
	cp.state = DELTOSHIP
	// Add some new value
	//cp.Owner = ''
	cp.ToShippingDate = time.Now()
}

// KhanhTNG - SetInstock returns the state to issued
func (cp *CustomerOrder) SetInstock() {
	cp.state = INSTOCK
	// Add some new value
	//cp.Owner = ''
	cp.InStockDate = time.Now()
}

// KhanhTNG - SetDelivering returns the state to issued
func (cp *CustomerOrder) SetDelivering() {
	cp.state = DELIVERING
	// Add some new value
	//cp.Owner = ''
	cp.InDeliveryDate = time.Now()
}

// KhanhTNG - SetCompleted returns the state to issued
func (cp *CustomerOrder) SetCompleted() {
	cp.state = COMPLETED
	// Add some new value
	//cp.Owner = ''
	cp.CompletedDate = time.Now()
}

// KhanhTNG - IsIssued returns true if state is issued
func (cp *CustomerOrder) IsCreated() bool {
	return cp.state == CREATED
}

// KhanhTNG - IsIssued returns true if state is issued
func (cp *CustomerOrder) IsReceived() bool {
	return cp.state == RECEIVED
}

// KhanhTNG - IsIssued returns true if state is issued
func (cp *CustomerOrder) IsDelToShip() bool {
	return cp.state == DELTOSHIP
}

// KhanhTNG - IsIssued returns true if state is issued
func (cp *CustomerOrder) IsInstock() bool {
	return cp.state == INSTOCK
}

// KhanhTNG - IsIssued returns true if state is issued
func (cp *CustomerOrder) IsDelivering() bool {
	return cp.state == DELIVERING
}

// KhanhTNG - IsIssued returns true if state is issued
func (cp *CustomerOrder) IsCompleted() bool {
	return cp.state == COMPLETED
}

// KhanhTNG - GetSplitKey returns values which should be used to form key
func (cp *CustomerOrder) GetSplitKey() []string {
	return []string{cp.orderNumber}
}

// KhanhTNG - Serialize formats the commercial paper as JSON bytes
func (cp *CustomerOrder) Serialize() ([]byte, error) {
	return json.Marshal(cp)
}

// KhanhTNG - Deserialize formats the commercial paper from JSON bytes
func Deserialize(bytes []byte, cp *CustomerOrder) error {
	err := json.Unmarshal(bytes, cp)

	if err != nil {
		return fmt.Errorf("Error deserializing customer order. %s", err.Error())
	}

	return nil
}
