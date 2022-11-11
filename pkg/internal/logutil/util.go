/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package logutil

import (
	"fmt"

	"github.com/hyperledger/aries-framework-go/pkg/common/log"
	"go.uber.org/zap"
)

// LogError is a utility function to log error messages.
func LogError(logger *log.Log, command, action, errMsg string, data ...string) {
	logger.Errorf("command=[%s] action=[%s] %s errMsg=[%s]", command, action, data, errMsg)
}

// LogDebug is a utility function to log debug messages.
func LogDebug(logger *log.Log, command, action, msg string, data ...string) {
	logger.Debugf("command=[%s] action=[%s] %s msg=[%s]", command, action, data, msg)
}

// LogInfo is a utility function to log info messages.
func LogInfo(logger *log.Log, command, action, msg string, data ...string) {
	logger.Infof("command=[%s] action=[%s] %s msg=[%s]", command, action, data, msg)
}

// LogInfo is a utility function to log info messages.
func LogWarn(logger *log.Log, command, action, msg string, data ...string) {
	logger.Warnf("command=[%s] action=[%s] %s msg=[%s]", command, action, data, msg)
}

// LogErrorStructured is a utility function to log error messages with structured elements.
func LogErrorStructured(logger *log.StructuredLog, command, action, msg string, data ...zap.Field) {
	logger.Error(fmt.Sprintf("command=[%s] action=[%s] msg=[%s]", command, action, msg), data...)
}

// LogWarnStructured is a utility function to log warn messages with structured elements.
func LogWarnStructured(logger *log.StructuredLog, command, action, msg string, data ...zap.Field) {
	logger.Warn(fmt.Sprintf("command=[%s] action=[%s] msg=[%s]", command, action, msg), data...)
}

// LogDebugStructured is a utility function to log debug messages with structured elements.
func LogDebugStructured(logger *log.StructuredLog, command, action, msg string, data ...zap.Field) {
	logger.Debug(fmt.Sprintf("command=[%s] action=[%s] msg=[%s]", command, action, msg), data...)
}

// LogInfoStructured is a utility function to log info messages with structured elements.
func LogInfoStructured(logger *log.StructuredLog, command, action, msg string, data ...zap.Field) {
	logger.Info(fmt.Sprintf("command=[%s] action=[%s] msg=[%s]", command, action, msg), data...)
}

// CreateKeyValueString creates a concatenated string.
func CreateKeyValueString(key, val string) string {
	return fmt.Sprintf("%s=[%s]", key, val)
}
