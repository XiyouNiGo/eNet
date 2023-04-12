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

package attach

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewAttachCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "detach",
		Short:   "Attaches the XDP program on the specified device",
		Example: "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Fatal("This command is being developing.")
		},
	}
	return cmd
}
