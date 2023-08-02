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
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/nat/add"
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/nat/list"
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/nat/remove"
	"github.com/spf13/cobra"
)

func NewNATCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nat",
		Short: "Command about Network Address Translation",
	}
	cmd.AddCommand(add.NewAddCommand())
	cmd.AddCommand(list.NewListCommand())
	cmd.AddCommand(remove.NewRemoveCommand())
	return cmd
}
