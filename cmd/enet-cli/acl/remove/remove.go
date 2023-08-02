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

package remove

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove",
		Short:   "Remove a rule from ACL",
		Example: "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Fatal("This command is being developing.")
		},
	}
	return cmd
}
