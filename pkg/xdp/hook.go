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

import (
	"log"
	"net"
	"os"
	"path"
	"syscall"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

// Hook provides a set of operations that allow for managing the execution of the XDP program
// including attaching it on the network interface, or removing the program from the interface.
type Hook struct {
	bpfObjs bpfObjects
	pinPath string
	bpfLink link.Link
}

// NewHook constructs a new instance of the XDP hook from provided XDP code.
func NewHook(pinPath string) (*Hook, error) {
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("Failed to remove memory lock: %v", err)
	}
	if err := os.MkdirAll(pinPath, os.ModePerm); err != nil {
		if os.IsNotExist(err) {
			if merr := syscall.Mount("bpf", BpfFsPath, "bpf", 0, "rw"); merr != nil {
				log.Fatalf("Failed to mount bps file system path: %v", merr)
			}
		} else {
			log.Fatalf("Failed to create bpf fs subpath: %v", err)
		}
	}
	hook := Hook{
		pinPath: pinPath,
	}
	if err := loadBpfObjects(&hook.bpfObjs, &ebpf.CollectionOptions{
		Maps: ebpf.MapOptions{
			PinPath: pinPath,
		},
	}); err != nil {
		log.Fatalf("Failed to load bpf objects: %v", err)
	}
	return &hook, nil
}

// Attach loads the XDP program to specified interface.
func (h *Hook) Attach(device string, mode link.XDPAttachFlags) error {
	iface, err := net.InterfaceByName(device)
	if err != nil {
		log.Fatalf("Failed to lookup network interface[%v]: %v", device, err)
	}
	l, err := link.AttachXDP(link.XDPOptions{
		Program:   h.bpfObjs.EnetFunc,
		Interface: iface.Index,
		Flags:     mode,
	})
	if err != nil {
		log.Fatalf("Failed to attach XDP program: %v", err)
	}
	h.bpfLink = l
	if err := l.Pin(path.Join(h.pinPath, EnetProgName)); err != nil {
		log.Fatalf("Failed to pin XDP program on %v: %v", h.pinPath, err)
	}
	return nil
}

// Remove unloads the XDP program from the interface.
func (h *Hook) Remove() error {
	l, err := link.LoadPinnedLink(path.Join(h.pinPath, EnetProgName), &ebpf.LoadPinOptions{})
	if err != nil {
		log.Fatalf("Failed to attach XDP program: %v", err)
	}
	h.bpfLink = l
	if err := l.Unpin(); err != nil {
		log.Fatalf("Failed to unpin XDP program on %v: %v", h.pinPath, err)
	}
	return nil
}

// Close closes the underlying eBPF module by disposing any allocated resources.
func (h *Hook) Close() error {
	if err := h.bpfLink.Close(); err != nil {
		log.Fatalf("Failed to close bpf link: %v", err)
	}
	if err := h.bpfObjs.Close(); err != nil {
		log.Fatalf("Failed to close bpf object: %v", err)
	}
	return nil
}

func (h *Hook) BPFObject() bpfObjects {
	return h.bpfObjs
}
