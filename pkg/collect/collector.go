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

package collect

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var namespace = "enet_"

var fromLabel = prometheus.Labels{
	"from": "enet-exporter",
}

// Define a struct for you collector that contains pointers
// to prometheus descriptors for each metric you wish to expose.
// Note you can also include fields of other types if they provide utility
// but we just won't be exposing them as metrics.
type ENetCollector struct {
	aclMetrics typedDesc
	natMetrics typedDesc
}

type typedDesc struct {
	desc      *prometheus.Desc
	valueType prometheus.ValueType
}

// You must create a constructor for you collector that
// initializes every descriptor and returns a pointer to the collector
func NewENetCollector() *ENetCollector {
	go autoIncreaseTestCase()
	return &ENetCollector{
		aclMetrics: typedDesc{
			prometheus.NewDesc(
				namespace+"acl_rule_hit_counts",
				"eNet ACL rule hit counts",
				[]string{
					"prior",
					"strategy", // native generic offloaded
					"proto",
					"src",
					"sport",
					"dst",
					"dport",
					"action", // accept drop
					"ctime",
				},
				fromLabel,
			),
			prometheus.CounterValue,
		},
		natMetrics: typedDesc{
			prometheus.NewDesc(
				namespace+"nat_rule_hit_counts",
				"eNet NAT rule hit counts",
				[]string{
					"type", // dnat snat fullnat
					"prior",
					"proto",
					"src",
					"sport",
					"dst",
					"dport",
					"ctime",
				},
				fromLabel,
			),
			prometheus.CounterValue,
		},
	}
}

// Each and every collector must implement the Describe function.
// It essentially writes all descriptors to the prometheus desc channel.
func (collector *ENetCollector) Describe(ch chan<- *prometheus.Desc) {
	// Update this section with the each metric you create for a given collector
	ch <- collector.aclMetrics.desc
	ch <- collector.natMetrics.desc
}

// We haven't implement eNet CLI actually, so sample metrics will be used.
func (collector *ENetCollector) collectACLMetrics() ([]Metrics, error) {
	var metrics []Metrics
	metrics = append(metrics,
		aclMetricsTestIn,
		aclMetricsTestOut,
	)
	return metrics, nil
}

func (collector *ENetCollector) collectNATMetrics() ([]Metrics, error) {
	var metrics []Metrics
	metrics = append(metrics,
		snatMetricsTest,
		dnatMetricsTest,
	)
	return metrics, nil
}

// Collect implements required collect function for all promehteus collectors
func (collector *ENetCollector) Collect(ch chan<- prometheus.Metric) {

	// Implement logic here to determine proper metric value to return to prometheus
	// for each descriptor or call other functions that do so.
	aclMetrics, err := collector.collectACLMetrics()
	if err != nil {
		logrus.Warnf("Failed to collect ACL metrics from eBPF Map: %v.", err)
	}
	natMetrics, err := collector.collectNATMetrics()
	if err != nil {
		logrus.Warnf("Failed to collect NAT metrics from eBPF Map: %v.", err)
	}

	// Write latest value for each metric in the prometheus metric channel.
	// Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	for _, metrics := range aclMetrics {
		ch <- prometheus.MustNewConstMetric(collector.aclMetrics.desc,
			prometheus.CounterValue, metrics.HitCount, metrics.LabelValues...)
	}
	for _, metrics := range natMetrics {
		ch <- prometheus.MustNewConstMetric(collector.natMetrics.desc,
			prometheus.CounterValue, metrics.HitCount, metrics.LabelValues...)
	}
}
