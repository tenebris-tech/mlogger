/******************************************************************************
 * Copyright (c) 2025-2026 Tenebris Technologies Inc.                         *
 * Please see LICENSE file for details.                                       *
 ******************************************************************************/

package mlogger

import (
	"fmt"
	"strings"
)

// FormatFields converts alternating key-value pairs into a logfmt string.
// Keys should be strings. Values are formatted with %v.
// Values containing spaces, '=', or '"' are quoted.
// An odd number of args results in the last key being assigned the value "MISSING".
func FormatFields(args ...any) string {
	if len(args) == 0 {
		return ""
	}

	var sb strings.Builder
	for i := 0; i < len(args); i += 2 {
		if i > 0 {
			sb.WriteByte(' ')
		}

		// Key
		key := fmt.Sprintf("%v", args[i])
		sb.WriteString(key)
		sb.WriteByte('=')

		// Value
		if i+1 >= len(args) {
			sb.WriteString("MISSING")
			break
		}
		val := fmt.Sprintf("%v", args[i+1])
		if strings.ContainsAny(val, " \t=\"") {
			sb.WriteByte('"')
			sb.WriteString(strings.ReplaceAll(val, `"`, `\"`))
			sb.WriteByte('"')
		} else {
			sb.WriteString(val)
		}
	}
	return sb.String()
}

// FormatMessage formats a message string followed by logfmt key=value pairs.
func FormatMessage(msg string, args ...any) string {
	if len(args) == 0 {
		return msg
	}
	return msg + " " + FormatFields(args...)
}

// DebugFields logs a structured debug message with logfmt key=value pairs.
func (m *MLogger) DebugFields(msg string, args ...any) {
	if m.debug {
		m.writeLog("DEBUG", FormatMessage(msg, args...))
	}
}

// InfoFields logs a structured informational message with logfmt key=value pairs.
func (m *MLogger) InfoFields(msg string, args ...any) {
	m.writeLog("INFO", FormatMessage(msg, args...))
}

// NoticeFields logs a structured notice message with logfmt key=value pairs.
func (m *MLogger) NoticeFields(msg string, args ...any) {
	m.writeLog("NOTICE", FormatMessage(msg, args...))
}

// WarningFields logs a structured warning message with logfmt key=value pairs.
func (m *MLogger) WarningFields(msg string, args ...any) {
	m.writeLog("WARNING", FormatMessage(msg, args...))
}

// ErrorFields logs a structured error message with logfmt key=value pairs.
func (m *MLogger) ErrorFields(msg string, args ...any) {
	m.writeLog("ERROR", FormatMessage(msg, args...))
}

// FatalFields logs a structured fatal message with logfmt key=value pairs, then exits.
func (m *MLogger) FatalFields(msg string, args ...any) {
	m.writeLog("FATAL", FormatMessage(msg, args...))
	m.FatalExit()
}
