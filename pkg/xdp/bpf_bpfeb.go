// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64be || armbe || mips || mips64 || mips64p32 || ppc64 || s390 || s390x || sparc || sparc64
// +build arm64be armbe mips mips64 mips64p32 ppc64 s390 s390x sparc sparc64

package xdp

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfActionMapKey struct {
	BitmapLowestBitN uint64
	BitmapArrayIndex uint64
}

type bpfActionMapVal struct {
	Action   uint32
	_        [4]byte
	HitCount uint64
}

type bpfBitmap [160]uint64

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	EnetFunc *ebpf.ProgramSpec `ebpf:"enet_func"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	EnetAclActionMap *ebpf.MapSpec `ebpf:"enet_acl_action_map"`
	EnetAclDportMap  *ebpf.MapSpec `ebpf:"enet_acl_dport_map"`
	EnetAclDstMap    *ebpf.MapSpec `ebpf:"enet_acl_dst_map"`
	EnetAclProtoMap  *ebpf.MapSpec `ebpf:"enet_acl_proto_map"`
	EnetAclSportMap  *ebpf.MapSpec `ebpf:"enet_acl_sport_map"`
	EnetAclSrcMap    *ebpf.MapSpec `ebpf:"enet_acl_src_map"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	EnetAclActionMap *ebpf.Map `ebpf:"enet_acl_action_map"`
	EnetAclDportMap  *ebpf.Map `ebpf:"enet_acl_dport_map"`
	EnetAclDstMap    *ebpf.Map `ebpf:"enet_acl_dst_map"`
	EnetAclProtoMap  *ebpf.Map `ebpf:"enet_acl_proto_map"`
	EnetAclSportMap  *ebpf.Map `ebpf:"enet_acl_sport_map"`
	EnetAclSrcMap    *ebpf.Map `ebpf:"enet_acl_src_map"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.EnetAclActionMap,
		m.EnetAclDportMap,
		m.EnetAclDstMap,
		m.EnetAclProtoMap,
		m.EnetAclSportMap,
		m.EnetAclSrcMap,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	EnetFunc *ebpf.Program `ebpf:"enet_func"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.EnetFunc,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfeb.o
var _BpfBytes []byte
