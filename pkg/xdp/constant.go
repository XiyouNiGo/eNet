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

const (
	Namespace = "enet"
	BpfFsPath = "/sys/fs/bpf"
)

const (
	XDPACLProgName     = "xdp_acl_prog"
	XDPNATProgName     = "xdp_nat_prog"
	XDPUnknownProgName = "unknown"
)

type XDPProgType uint8

const (
	XDPProgTypeACL XDPProgType = 1 << iota
	XDPProgTypeNAT
)

func ToXDPProgName(prog_type XDPProgType) string {
	switch prog_type {
	case XDPProgTypeACL:
		return XDPACLProgName
	case XDPProgTypeNAT:
		return XDPNATProgName
	default:
		return XDPUnknownProgName
	}
}
