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
	"math/rand"
	"time"
)

type Metrics struct {
	LabelValues []string
	HitCount    float64
}

var (
	aclMetricsTestOut = Metrics{
		LabelValues: []string{
			"1",
			"native",
			"tcp",
			"",
			"",
			"39.156.66.10", // baidu.com
			"80",           // http
			"drop",
			time.Now().Format("2006-01-02 15:04:05"),
		},
		HitCount: 1000,
	}
	aclMetricsTestIn = Metrics{
		LabelValues: []string{
			"10",
			"generic",
			"tcp",
			"172.21.132.133", // WSL2
			"22",             // ssh
			"",
			"",
			"accept",
			time.Now().Format("2006-01-02 15:04:05"),
		},
		HitCount: 2000,
	}
	snatMetricsTest = Metrics{
		LabelValues: []string{
			"snat",
			"1",
			"tcp",
			"172.21.132.133", // WSL2
			"80",             // http
			"127.0.0.1",      // localhost
			"80",             // http
			time.Now().Format("2006-01-02 15:04:05"),
		},
		HitCount: 3000,
	}
	dnatMetricsTest = Metrics{
		LabelValues: []string{
			"dnat",
			"1",
			"tcp",
			"127.0.0.1",    // localhost
			"80",           // http
			"39.156.66.10", // baidu.com
			"80",           // http
			time.Now().Format("2006-01-02 15:04:05"),
		},
		HitCount: 4000,
	}
)

func autoIncreaseTestCase() {
	for {
		time.Sleep(time.Second)
		aclMetricsTestOut.HitCount += float64(rand.Intn(100))
		aclMetricsTestIn.HitCount += float64(rand.Intn(100))
		snatMetricsTest.HitCount += float64(rand.Intn(100))
		dnatMetricsTest.HitCount += float64(rand.Intn(100))
	}
}
