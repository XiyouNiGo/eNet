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
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/acl"
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/nat"
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/xdp"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "enet",
		Short:   "eNet is a eBPF-based net tool, which supporting ACL and NAT now.",
		Version: "0.0.0",
	}
	cmd.AddCommand(acl.NewACLCommand())
	cmd.AddCommand(nat.NewNATCommand())
	cmd.AddCommand(xdp.NewXDPCommand())
	return cmd
}
