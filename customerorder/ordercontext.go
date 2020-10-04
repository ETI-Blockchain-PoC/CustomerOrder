/*
 * SPDX-License-Identifier: Apache-2.0
 */

package customerorder

import (
	"contractapi"
)

// TransactionContextInterface an interface to
// describe the minimum required functions for
// a transaction context in the commercial
// paper
type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetOrderList() ListInterface
}

// TransactionContext implementation of
// TransactionContextInterface for use with
// commercial paper contract
type TransactionContext struct {
	contractapi.TransactionContext
	orderList *list
}

// GetPaperList return paper list
func (tc *TransactionContext) GetOrderList() ListInterface {
	if tc.orderList == nil {
		tc.orderList = newList(tc)
	}

	return tc.orderList
}
