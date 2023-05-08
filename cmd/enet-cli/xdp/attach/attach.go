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
	"path"

	"github.com/XiyouNiGo/eNet/pkg/xdp"
	"github.com/cilium/ebpf/link"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	offloaded bool
	native    bool
	generic   bool
	pinPath   string
)

func toXDPModeString(flags link.XDPAttachFlags) string {
	switch flags {
	case link.XDPOffloadMode:
		return "XDPOffloadMode"
	case link.XDPDriverMode:
		return "XDPDriverMode"
	default:
		return "XDPGenericMode"
	}
}

func NewAttachCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "attach",
		Short:   "Attach the XDP program on the specified device",
		Example: "TODO",
		Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			device := args[0]
			mode := func() link.XDPAttachFlags {
				switch {
				case offloaded:
					return link.XDPOffloadMode
				case native:
					return link.XDPDriverMode
				default:
					return link.XDPGenericMode
				}
			}()
			hook, err := xdp.NewHook(pinPath, xdp.XDPProgTypeACL)
			if err != nil {
				logger.Fatalf("Failed to new hook: %v", err)
			}
			defer hook.Close()
			if err := hook.Attach(device, mode); err != nil {
				logger.Fatalf("Failed to attach hook in mode %v: %v",
					toXDPModeString(mode), err)
			}
			logger.Infof("XDP program successfully attached to %v device in mode %v",
				device, toXDPModeString(mode))
		},
	}
	cmd.Flags().BoolVarP(&offloaded, "offloaded", "", false, "XDP offloaded mode")
	cmd.Flags().BoolVarP(&native, "native", "", false, "XDP native mode")
	cmd.Flags().BoolVarP(&generic, "generic", "", true, "XDP generic mode (default)")
	cmd.MarkFlagsMutuallyExclusive("offloaded", "native", "generic")
	cmd.Flags().StringVarP(&pinPath, "pin-path", "p", path.Join(xdp.BpfFsPath,
		xdp.Namespace), "Path to pin up XDP program and map")
	cmd.MarkFlagFilename("pin-path")
	return cmd
}
