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

#include <linux/stddef.h>
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
#include <sys/cdefs.h>
#include <stdlib.h>
#include "map.h"

char __license[] SEC("license") = "Dual MIT/GPL";

static __always_inline int parse_ethhdr(const struct xdp_md *ctx,
                                        struct ethhdr **eth_hdr) {
  if (NULL == ctx || NULL == eth_hdr) {
    return -1;
  }
  void *data = (void *)(long)ctx->data;
  void *data_end = (void *)(long)ctx->data_end;
  if ((void *)((struct ethhdr *)data + 1) > data_end) {
    return -1;
  }
  *eth_hdr = data;
  return (*eth_hdr)->h_proto;
}

// static __always_inline int parse_iphdr(const struct ethhdr *eth_hdr,
//                                        const void *data_end,
//                                        struct iphdr **ip_hdr) {
//   if (NULL == eth_hdr || NULL == data_end  || NULL == ip_hdr) {
//     return -1;
//   }
//   void *data = (void*)(eth_hdr + 1);
//   if ((void *)((struct iphdr *)data + 1) > data_end) {
//     return -1;
//   }
//   *ip_hdr = data;
//   return (*ip_hdr)->protocol;
// }

// static __always_inline int parse_tcphdr(const struct iphdr *ip_hdr,
//                                         const void *data_end,
//                                         struct tcphdr **tcp_hdr) {
//   if (NULL == ip_hdr || NULL == data_end || NULL == tcp_hdr) {
//     return -1;
//   }
//   void *data = (void*)(ip_hdr + 1);
//   if ((void *)((struct tcphdr *)data + 1) > data_end) {
//     return -1;
//   }
//   *tcp_hdr = data;
//   return (*tcp_hdr)->doff << 2;
// }

// static __always_inline int parse_udphdr(const struct iphdr *ip_hdr,
//                                         const void *data_end,
//                                         struct udphdr **udp_hdr) {
//   if (NULL == ip_hdr || NULL == data_end || NULL == udp_hdr) {
//     return -1;
//   }
//   void *data = (void*)(ip_hdr + 1);
//   if ((void *)((struct udphdr *)data + 1) > data_end) {
//     return -1;
//   }
//   *udp_hdr = data;
//   return bpf_ntohs((*udp_hdr)->len) - sizeof(struct udphdr);
// }

// static __always_inline int append_bitmap(struct bitmap_array *bitmap_arr,
//                                          bitmap *to_append) {
//   if (NULL == bitmap_arr || bitmap_arr->len >= ACL_MAP_NUM) {
//     return -1;
//   }
//   bitmap_arr->arr[bitmap_arr->len] = to_append;
//   bitmap_arr->len++;
//   return 0;
// }

// static __always_inline void get_bitmap_array(__u32 src, __u32 dst, __u16 sport,
//                                             __u16 dport, __u32 proto,
//                                             struct bitmap_array *bitmap_arr) {
//   bitmap *tmp_bitmap = NULL;
//   tmp_bitmap = bpf_map_lookup_elem(&enet_acl_src_map, &src);
//   if (NULL != tmp_bitmap) {
//     append_bitmap(bitmap_arr, tmp_bitmap);
//   }
//   tmp_bitmap = bpf_map_lookup_elem(&enet_acl_dst_map, &dst);
//   if (NULL != tmp_bitmap) {
//     append_bitmap(bitmap_arr, tmp_bitmap);
//   }
//   tmp_bitmap = bpf_map_lookup_elem(&enet_acl_sport_map, &sport);
//   if (NULL != tmp_bitmap) {
//     append_bitmap(bitmap_arr, tmp_bitmap);
//   }
//   tmp_bitmap = bpf_map_lookup_elem(&enet_acl_dport_map, &dport);
//   if (NULL != tmp_bitmap) {
//     append_bitmap(bitmap_arr, tmp_bitmap);
//   }
//   tmp_bitmap = bpf_map_lookup_elem(&enet_acl_proto_map, &proto);
//   if (NULL != tmp_bitmap) {
//     append_bitmap(bitmap_arr, tmp_bitmap);
//   }
// }

// static __always_inline void
// get_bitmap_array_for_icmp(__u32 src, __u32 dst, __u32 proto,
//                           struct bitmap_array *bitmap_arr) {
//   bitmap *tmp_bitmap = NULL;
//   tmp_bitmap = bpf_map_lookup_elem(&enet_acl_src_map, &src);
//   if (NULL != tmp_bitmap) {
//     append_bitmap(bitmap_arr, tmp_bitmap);
//   }
//   tmp_bitmap = bpf_map_lookup_elem(&enet_acl_dst_map, &dst);
//   if (NULL != tmp_bitmap) {
//     append_bitmap(bitmap_arr, tmp_bitmap);
//   }
//   tmp_bitmap = bpf_map_lookup_elem(&enet_acl_proto_map, &proto);
//   if (NULL != tmp_bitmap) {
//     append_bitmap(bitmap_arr, tmp_bitmap);
//   }
// }

// static __always_inline int get_lowest_bit_n(__u64 n) {
//   return n - (n & (n - 1)); // Or: n&(-n)
// }

// static __always_inline enum xdp_action
// get_xdp_action(struct bitmap_array *bitmap_arr) {
//   struct action_map_key action_key;
//   struct action_map_val *action_val;

//   // Unknown rule
//   if (unlikely(bitmap_arr->len != 3 || bitmap_arr->len != 5)) {
//     return XDP_DROP;
//   }

//   // Reverse(little endian)
//   __u64 tmp_res = 0;
//   __u64 bitmap_array_index = BITMAP_ARRAY_SIZE - 1;
//   for (; bitmap_array_index >= 0; bitmap_array_index--) {
//     if (5 == bitmap_arr->len) {
//       tmp_res = ((*bitmap_arr->arr[0])[bitmap_array_index]) &
//       ((*bitmap_arr->arr[1])[bitmap_array_index]) &
//       ((*bitmap_arr->arr[2])[bitmap_array_index]) &
//       ((*bitmap_arr->arr[3])[bitmap_array_index]) &
//       ((*bitmap_arr->arr[4])[bitmap_array_index]);
//     } else { // 3 == bitmap_arr->len
//       tmp_res = ((*bitmap_arr->arr[0])[bitmap_array_index]) &
//       ((*bitmap_arr->arr[1])[bitmap_array_index]) &
//       ((*bitmap_arr->arr[2])[bitmap_array_index]);
//     }
//     if (tmp_res > 0) {
//       break;
//     }
//   }
//   // Dont't match any rule
//   if (tmp_res == 0) {
//     return XDP_PASS;
//   }
//   action_key.bitmap_array_index = bitmap_array_index;
//   action_key.bitmap_lowest_bit_n = get_lowest_bit_n(tmp_res);
//   action_val = bpf_map_lookup_elem(&enet_acl_action_map, &action_key);
//   if (NULL != action_val) {
//     __sync_fetch_and_add(&action_val->hit_count, 1);
//     return action_val->action;
//   }
//   return XDP_PASS;
// }

SEC("xdp")
int enet_func(struct xdp_md *ctx) {
  // void *data_end = (void *)(long)ctx->data_end;
  struct ethhdr *eth_hdr = NULL;
  // struct iphdr *ip_hdr = NULL;
  // struct bitmap_array bitmap_arr;

  int eth_proto = parse_ethhdr(ctx, &eth_hdr);
  bpf_printk("nigo: %d", eth_proto);
  // IP Protocol
  // if (bpf_htons(ETH_P_IP) == eth_proto) {
  //   int ip_proto = parse_iphdr(eth_hdr, data_end, &ip_hdr);
  //   // TCP Protocol
  //   if (likely(IPPROTO_TCP == ip_proto)) {
  //     struct tcphdr *tcp_hdr = NULL;
  //     if (parse_tcphdr(ip_hdr, data_end, &tcp_hdr) < 0) {
  //       return XDP_DROP;
  //     }
  //     get_bitmap_array(ip_hdr->saddr, ip_hdr->daddr, tcp_hdr->source,
  //                      tcp_hdr->dest, ip_proto, &bitmap_arr);
  //   } else if (IPPROTO_UDP == ip_proto) { // UDP Protocol
  //     struct udphdr *udp_hdr = NULL;
  //     if (parse_udphdr(ip_hdr, data_end, &udp_hdr) < 0) {
  //       return XDP_DROP;
  //     }
  //     get_bitmap_array(ip_hdr->saddr, ip_hdr->daddr, udp_hdr->source,
  //                      udp_hdr->dest, ip_proto, &bitmap_arr);
  //   } else if (IPPROTO_ICMP == ip_proto) { // ICMP Protocol
  //     get_bitmap_array_for_icmp(ip_hdr->saddr, ip_hdr->daddr, ip_proto,
  //                               &bitmap_arr);
  //   } else { // Other L4 Protocol
  //     return XDP_PASS;
  //   }
  //   return get_xdp_action(&bitmap_arr);
  // }
  // // Other L3 Protocol
  // if (eth_proto > 0) {
  //   return XDP_PASS;
  // }
  // // Invalid Protocol
  // return XDP_DROP;
  return XDP_PASS;
}
