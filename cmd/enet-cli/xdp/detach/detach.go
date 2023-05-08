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

package detach

import (
	"path"

	"github.com/XiyouNiGo/eNet/pkg/xdp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	pinPath string
)

func NewDetachCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "detach",
		Short:   "Removes the XDP program from the specified device",
		Example: "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			hook, err := xdp.NewHook(pinPath, xdp.XDPProgTypeACL)
			if err != nil {
				logger.Fatalf("Failed to new hook: %v", err)
			}
			defer hook.Close()
			if err := hook.Remove(); err != nil {
				logger.Fatalf("Failed to detach hook: %v", err)
			}
			logger.Infof("XDP program successfully detached from %v", pinPath)
		},
	}
	cmd.Flags().StringVarP(&pinPath, "pin-path", "p", path.Join(xdp.BpfFsPath,
		xdp.Namespace), "Path to pin up XDP program and map")
	cmd.MarkFlagFilename("pin-path")
	return cmd
}
