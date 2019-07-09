/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

// Package wgcfg has types and a parser for representing WireGuard config.
package wgcfg

import (
	"fmt"
	"net"
	"strings"
)

// Config is a WireGuard configuration.
type Config struct {
	Name      string
	Interface Interface
	Peers     []Peer
}

type Interface struct {
	PrivateKey PrivateKey
	Addresses  []CIDR
	ListenPort uint16
	MTU        uint16
	DNS        []net.IP // TODO(crawshaw): use wgcfg.IP type
}

type Peer struct {
	PublicKey           Key
	PresharedKey        SymmetricKey
	AllowedIPs          []CIDR
	Endpoint            Endpoint
	PersistentKeepalive uint16
}

type Endpoint struct {
	Host string
	Port uint16
}

func (e Endpoint) String() string {
	if strings.IndexByte(e.Host, ':') > 0 {
		return fmt.Sprintf("[%s]:%d", e.Host, e.Port)
	}
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

func (e Endpoint) IsEmpty() bool {
	return len(e.Host) == 0
}