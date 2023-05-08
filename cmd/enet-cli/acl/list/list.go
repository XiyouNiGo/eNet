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

package list

import (
	"fmt"
	"log"
	"net"
	"path"
	"strings"
	"time"

	"github.com/XiyouNiGo/eNet/pkg/xdp"
	"github.com/cilium/ebpf"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	pinPath string
)

func NewListCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Show all rules registered in ACL",
		Example: "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			hook, err := xdp.NewHook(pinPath, xdp.XDPProgTypeACL)
			if err != nil {
				logger.Fatalf("Failed to new hook: %v", err)
			}
			defer hook.Close()
			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()
			for range ticker.C {
				s, err := formatMapContents(hook.BPFObject().XdpStatsMap)
				if err != nil {
					log.Printf("Error reading map: %s", err)
					continue
				}
				log.Printf("Map contents:\n%s", s)
			}
		},
	}
	cmd.Flags().StringVarP(&pinPath, "pin-path", "p", path.Join(xdp.BpfFsPath,
		xdp.Namespace), "Path to pin up XDP program and map")
	cmd.MarkFlagFilename("pin-path")
	return cmd
}

func formatMapContents(m *ebpf.Map) (string, error) {
	var (
		sb  strings.Builder
		key []byte
		val uint32
	)
	iter := m.Iterate()
	for iter.Next(&key, &val) {
		sourceIP := net.IP(key) // IPv4 source address in network byte order.
		packetCount := val
		sb.WriteString(fmt.Sprintf("\t%s => %d\n", sourceIP, packetCount))
	}
	return sb.String(), iter.Err()
}
