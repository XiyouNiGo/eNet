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

package cmd

import (
	"github.com/XiyouNiGo/eNet/cmd/acl"
	"github.com/XiyouNiGo/eNet/cmd/nat"
	"github.com/XiyouNiGo/eNet/cmd/xdp"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "enet",
		Short:   "eNet is a eBPF-based net tool, which supporting ACL and NAT now.",
		Version: "0.0.0",
	}
	cmd.AddCommand(acl.NewACLCommand(logger))
	cmd.AddCommand(nat.NewNATCommand(logger))
	cmd.AddCommand(xdp.NewXDPCommand(logger))
	return cmd
}
