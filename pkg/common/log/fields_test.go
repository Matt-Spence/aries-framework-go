/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package log

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStandardFields(t *testing.T) {
	const module = "test_module"

	t.Run("console error", func(t *testing.T) {
		stdErr := newMockWriter()

		logger := NewStructured(module,
			WithStdErr(stdErr),
		)

		logger.Error("Sample error", WithError(errors.New("some error")))

		require.Contains(t, stdErr.Buffer.String(), `Sample error	{"service": "myservice", "error": "some error"}`)
	})

	t.Run("json error", func(t *testing.T) {
		stdErr := newMockWriter()

		logger := NewStructured(module,
			WithStdErr(stdErr), WithEncoding(JSON),
		)

		logger.Error("Sample error", WithError(errors.New("some error")))

		l := unmarshalLogData(t, stdErr.Bytes())

		require.Equal(t, "some error", l.Error)
	})

	t.Run("json fields 1", func(t *testing.T) {
		stdOut := newMockWriter()

		logger := NewStructured(module, WithStdOut(stdOut), WithEncoding(JSON))

		logger.Info("Some message",
			WithTotal(12),
		)

		l := unmarshalLogData(t, stdOut.Bytes())

		require.Equal(t, 12, l.Total)
	})
}

type logData struct {
	Error  string `json:"error"`
	Total                  int                 `json:"total"`
}

func unmarshalLogData(t *testing.T, b []byte) *logData {
	t.Helper()

	l := &logData{}

	require.NoError(t, json.Unmarshal(b, l))

	return l
}
