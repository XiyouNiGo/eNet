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

package model

import (
	"os"
	"sort"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	RuleStrategyNative    = "native"
	RuleStrategyGeneric   = "generic"
	RuleStrategyOffloaded = "offloaded"
)

const (
	RuleProtoTCP  = "TCP"
	RuleProtoUDP  = "UDP"
	RuleProtoICMP = "ICMP"
)

const (
	RuleActionPass = "PASS"
	RuleActionDrop = "DROP"
)

type ACLRule struct {
	Priority uint32 `yaml:"priority"`
	Strategy string `yaml:"strategy"`
	Proto    string `yaml:"proto"`
	Ctime    int64  `yaml:"ctime"`
	Src      string `yaml:"src"`
	Sport    uint16 `yaml:"sport"`
	Dst      string `yaml:"dst"`
	Dport    uint16 `yaml:"dport"`
	Action   string `yaml:"action"`
	HitCount string `yaml:"hit_count"`
}

type ACLRules []ACLRule

func NewACLRules(configFile string) *ACLRules {
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		logrus.Fatalf("Failed to read file %v: %v", configFile, err)
	}
	var rules ACLRules
	if err := yaml.Unmarshal(bytes, &rules); err != nil {
		logrus.Fatalf("Failed to unmarshal acl rules: %v", err)
	}
	sort.Slice(rules, func(i, j int) bool {
		return (rules)[i].Ctime > (rules)[j].Ctime
	})
	return &rules
}

func (rules *ACLRules) Validate() {

}
