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

package add

import (
	"github.com/XiyouNiGo/eNet/pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	configFile string
)

func NewAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add",
		Short:   "Append a new rule to ACL",
		Example: "TODO",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			rules := model.NewACLRules(configFile)
			logrus.Fatal(rules)
		},
	}
	cmd.Flags().StringVarP(&configFile, "config-file", "f", "", "Initial ACL rule config file")
	cmd.MarkFlagRequired("config-file")
	cmd.MarkFlagFilename("config-file")
	return cmd
}
