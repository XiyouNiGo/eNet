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

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc $BPF_CLANG -cflags $BPF_CFLAGS bpf ./xdp.c -- -I./linux/
package xdp

import "C"
