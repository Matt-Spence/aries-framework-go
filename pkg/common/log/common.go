/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package log

import "go.uber.org/zap"

type loggerFunc func(msg string, fields ...zap.Field)

// CloseResponseBodyError outputs a 'close response body' error log to the given logger.
func CloseResponseBodyError(log loggerFunc, err error) {
	log("Error closing response body", WithError(err))
}
