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

package main

import (
	"aws-observability.io/collector/pkg/config"
	"aws-observability.io/collector/pkg/defaultcomponents"
	"aws-observability.io/collector/pkg/logger"
	"aws-observability.io/collector/tools/version"
	"github.com/pkg/errors"
	"go.opentelemetry.io/collector/service"
	"go.uber.org/zap/zapcore"
	"log"
)

func main() {
	logger.SetupErrorLogger()
	handleErr := func(message string, err error) {
		if err != nil {
			log.Fatalf(
				"%s: %v", message, err)
		}
	}

	factories, err := defaultcomponents.Components()
	handleErr("Failed to build components", err)

	// init cfgFactory
	cfgFactory := config.GetCfgFactory()

	// init lumberFunc for zap logger
	lumberHook := logger.GetLumberHook()

	info := service.ApplicationStartInfo{
		ExeName:  "aws-observability-collector",
		LongName: "AWS Observability Collector",
		Version:  version.Version,
		GitHash:  version.GitHash,
	}

	if err := run(service.Parameters{
		Factories:            factories,
		ApplicationStartInfo: info,
		ConfigFactory:        cfgFactory,
		LoggingHooks:         []func(entry zapcore.Entry) error{lumberHook}}); err != nil {
		log.Fatal(err)
	}

}

func runInteractive(params service.Parameters) error {
	app, err := service.New(params)
	if err != nil {
		return errors.Wrap(err, "failed to construct the application")
	}

	err = app.Start()
	if err != nil {
		return errors.Wrap(err, "application run finished with error: %v")
	}

	return nil
}
