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

package acl

import (
	"github.com/XiyouNiGo/eNet/cmd/acl/add"
	"github.com/XiyouNiGo/eNet/cmd/acl/list"
	"github.com/XiyouNiGo/eNet/cmd/acl/remove"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewACLCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acl",
		Short: "Command about Access Control List",
	}
	cmd.AddCommand(add.NewAddCommand(logger))
	cmd.AddCommand(list.NewListCommand(logger))
	cmd.AddCommand(remove.NewRemoveCommand(logger))
	return cmd
}
