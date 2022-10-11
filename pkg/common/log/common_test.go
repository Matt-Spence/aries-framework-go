/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package log

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommonLogs(t *testing.T) {
	const module = "test_module"

	t.Run("CloseResponseBodyError", func(t *testing.T) {
		stdOut := newMockWriter()

		logger := NewStructured(module,
			WithStdOut(stdOut),
		)

		CloseResponseBodyError(logger.Info, errors.New("response body error"))

		require.Contains(t, stdOut.Buffer.String(), `Error closing response body`)
		require.Contains(t, stdOut.Buffer.String(), `"service": "myservice"`)
		require.Contains(t, stdOut.Buffer.String(), `"error": "response body error"`)
	})
}
