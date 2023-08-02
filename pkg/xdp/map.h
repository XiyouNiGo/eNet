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

//go:build ignore_vet

#include <linux/types.h>
#include <bpf/bpf_endian.h>
#include <bpf/bpf_helpers.h>
#include <linux/bpf.h>
#include <linux/icmp.h>
#include <linux/if_ether.h>
#include <linux/in.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <linux/udp.h>
#include <stddef.h>

#define likely(x) __builtin_expect(!!(x), 1)
#define unlikely(x) __builtin_expect(!!(x), 0)

#define BITMAP_ARRAY_SIZE 160
#define RULE_MAX_NUM 64 * BITMAP_ARRAY_SIZE // __u64
#define IP_MAX_ENTRIES RULE_MAX_NUM
#define PORT_MAX_ENTRIES 65536 // 2^16
#define PROTO_MAX_ENTRIES 4    // TCP UDP ICMP unknown
#define ACTION_MAX_ENTRIES RULE_MAX_NUM

enum acl_map_type {
  ACL_SRC_MAP = 0,
  ACL_DST_MAP,
  ACL_SPORT_MAP,
  ACL_DPORT_MAP,
  ACL_PROTO_MAP,
  ACL_MAP_NUM,
};

typedef  __u64 bitmap[BITMAP_ARRAY_SIZE];

struct bitmap_array {
  int len;
  bitmap *arr[ACL_MAP_NUM];
};

struct action_map_key {
  __u64 bitmap_lowest_bit_n;
  __u64 bitmap_array_index;
};

struct action_map_val {
  enum xdp_action action;
  __u64 hit_count;
};

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, IP_MAX_ENTRIES);
	__type(key, __u32);
  __type(value, bitmap);
  __uint(pinning, LIBBPF_PIN_BY_NAME);
} enet_acl_src_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, IP_MAX_ENTRIES);
	__type(key, __u32);
  __type(value, bitmap);
  __uint(pinning, LIBBPF_PIN_BY_NAME);
} enet_acl_dst_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, PORT_MAX_ENTRIES);
	__type(key, __u16);
  __type(value, bitmap);
  __uint(pinning, LIBBPF_PIN_BY_NAME);
} enet_acl_sport_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, PORT_MAX_ENTRIES);
	__type(key, __u16);
  __type(value, bitmap);
  __uint(pinning, LIBBPF_PIN_BY_NAME);
} enet_acl_dport_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, PROTO_MAX_ENTRIES);
	__type(key, __u32);
  __type(value, bitmap);
  __uint(pinning, LIBBPF_PIN_BY_NAME);
} enet_acl_proto_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, ACTION_MAX_ENTRIES);
	__type(key, struct action_map_key);
  __type(value, struct action_map_val);
  __uint(pinning, LIBBPF_PIN_BY_NAME);
} enet_acl_action_map SEC(".maps");