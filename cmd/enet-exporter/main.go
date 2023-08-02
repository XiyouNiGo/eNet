/*
 * Copyright (c) 2023 NiGo.
 *
 * Licensed under the MIT license; you may not use this
 * file except in compliance with the License. You may
 * obtain a copy of the License at
 *
 * http://valums.com/mit-license/
 *
 */

package main

import (
	"github.com/XiyouNiGo/eNet/cmd/enet-exporter/root"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	cmd := root.NewCommand()
	if err := cmd.Execute(); err != nil {
		logrus.Errorf("Failed to execute enet: %v", err)
	}
}
