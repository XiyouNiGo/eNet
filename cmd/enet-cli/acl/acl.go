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
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/acl/add"
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/acl/list"
	"github.com/XiyouNiGo/eNet/cmd/enet-cli/acl/remove"
	"github.com/spf13/cobra"
)

func NewACLCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acl",
		Short: "Command about Access Control List",
	}
	cmd.AddCommand(add.NewAddCommand())
	cmd.AddCommand(list.NewListCommand())
	cmd.AddCommand(remove.NewRemoveCommand())
	return cmd
}
