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

package nat

import (
	"github.com/XiyouNiGo/eNet/cmd/nat/add"
	"github.com/XiyouNiGo/eNet/cmd/nat/list"
	"github.com/XiyouNiGo/eNet/cmd/nat/remove"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewNATCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nat",
		Short: "Command about Network Address Translation",
	}
	cmd.AddCommand(add.NewAddCommand(logger))
	cmd.AddCommand(list.NewListCommand(logger))
	cmd.AddCommand(remove.NewRemoveCommand(logger))
	return cmd
}
