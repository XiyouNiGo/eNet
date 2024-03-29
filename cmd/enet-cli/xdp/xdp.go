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

package xdp

import (
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/xdp/attach"
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/xdp/detach"
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/xdp/purge"
	"github.com/spf13/cobra"
)

func NewXDPCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "xdp",
		Short: "Command about eXpress Data Path",
	}
	cmd.AddCommand(attach.NewAttachCommand())
	cmd.AddCommand(detach.NewDetachCommand())
	cmd.AddCommand(purge.NewPurgeCommand())
	return cmd
}
