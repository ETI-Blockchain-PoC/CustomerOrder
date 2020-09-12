/*
 * SPDX-License-Identifier: Apache-2.0
 */

package ledgerapi

import (
	"strings"
)

// KhanhTNG - SplitKey splits a key on colon
// KhanhTNG - Still keep this logic for refer in future if needed - Current model not need to split and join anymore
func SplitKey(key string) []string {
	return strings.Split(key, ":")
}

// KhanhTNG - MakeKey joins key parts using colon
func MakeKey(keyParts ...string) string {
	return strings.Join(keyParts, ":")
}

// StateInterface interface states must implement
// for use in a list
type StateInterface interface {
	// GetSplitKey return components that combine to form the key
	GetSplitKey() []string
	Serialize() ([]byte, error)
}
