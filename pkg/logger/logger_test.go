/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package logger

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"testing"
)

func TestGetLumberHook(t *testing.T) {
	entry := zapcore.Entry{
		Message: "test",
	}
	funcCall := GetLumberHook()
	err := funcCall(entry)
	require.NoError(t, err)
}

func TestSetupErrorLogger(t *testing.T) {
	SetupErrorLogger()
	_, ok := log.Writer().(*lumberjack.Logger)
	assert.True(t, ok)
}
