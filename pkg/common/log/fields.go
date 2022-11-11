/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package log

import (
	"go.uber.org/zap"
)

// Note: These fields are not currently used, but exist as example for how fields can be constructed for use in structured logs. 

// Log Fields.
const (
	FieldTotal                  = "total"
)

// WithError sets the error field.
func WithError(err error) zap.Field {
	return zap.Error(err)
}

// WithTotal sets the total field.
func WithTotal(value int) zap.Field {
	return zap.Int(FieldTotal, value)
}
