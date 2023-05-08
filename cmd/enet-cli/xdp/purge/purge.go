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

package purge

import (
	"os"
	"path"

	"github.com/XiyouNiGo/eNet/pkg/xdp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	pinPath string
)

func NewPurgeCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "purge",
		Short:   "Purge the XDP program and map",
		Example: "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			if err := os.RemoveAll(pinPath); err != nil {
				logger.Fatalf("Failed to purge XDP program and map on %v: %v", pinPath, err)
			}
			logger.Infof("XDP program and map successfully purged from %v", pinPath)
		},
	}
	cmd.Flags().StringVarP(&pinPath, "pin-path", "p", path.Join(xdp.BpfFsPath,
		xdp.Namespace), "Path to pin up XDP program and map")
	cmd.MarkFlagFilename("pin-path")
	return cmd
}
