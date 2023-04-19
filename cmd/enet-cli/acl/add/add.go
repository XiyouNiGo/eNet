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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewAddCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add",
		Short:   "Appends a new rule to ACL",
		Example: "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Fatal("This command is being developing.")
		},
	}
	return cmd
}
