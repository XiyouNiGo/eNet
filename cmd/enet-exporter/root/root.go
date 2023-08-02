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

package root

import (
	"fmt"
	"net/http"

	"github.com/XiyouNiGo/eNet/pkg/collect"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	port int16
	url  string
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "enet-exporter",
		Short:   "eNet Exporter is a exporter for prometheus, collecting network message generated from eNet CLI.",
		Version: "0.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			prometheus.MustRegister(collect.NewENetCollector())
			http.Handle(url, promhttp.Handler())
			logrus.Infof("eNet Exporter serve on %v use port %v.", url, port)
			if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
				logrus.Fatalf(err.Error())
			}
		},
	}
	cmd.Flags().Int16VarP(&port, "port", "p", 9002, "Exporter serve port")
	cmd.Flags().StringVarP(&url, "url", "u", "/metrics", "Exporter serve url")
	return cmd
}
