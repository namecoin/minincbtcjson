// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2019 The Namecoin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// NOTE: This file is intended to house the RPC commands that are supported by
// a name lookup client.

package minincbtcjson

// Encoding represents an encoding for name identifiers and values.
type Encoding string

const (
	ASCII Encoding = "ascii"
	UTF8  Encoding = "utf8"
	Hex   Encoding = "hex"
)

// HashAlgo represents a hash function for name identifiers.
type HashAlgo string

const (
	Direct  HashAlgo = "direct"
	SHA256d HashAlgo = "sha256d"
)

// NameShowOptions represents the optional options struct provided with a
// NameShowCmd command.
type NameShowOptions struct {
	NameEncoding  Encoding `json:"nameEncoding,omitempty"`
	ValueEncoding Encoding `json:"valueEncoding,omitempty"`
	ByHash        HashAlgo `json:"byHash,omitempty"`
	AllowExpired  bool     `json:"allowExpired"`
	StreamID      string   `json:"streamID,omitempty"`
}

// NameShowCmd defines the name_show JSON-RPC command.
//
//nolint:govet // fieldalignment: order of fields has semantic value.
type NameShowCmd struct {
	Name    string
	Options *NameShowOptions
}

// NameScanOptions represents the optional options struct provided with a
// NameScanCmd command.
type NameScanOptions struct {
	NameEncoding  Encoding `json:"nameEncoding,omitempty"`
	ValueEncoding Encoding `json:"valueEncoding,omitempty"`
	MinConf       *int32   `json:"minConf,omitempty"`
	MaxConf       *int32   `json:"maxConf,omitempty"`
	RegExp        *string  `json:"regexp,omitempty"`
	Prefix        string   `json:"prefix,omitempty"`
}

// NameScanCmd defines the name_scan JSON-RPC command.
//
//nolint:govet // fieldalignment: order of fields has semantic value.
type NameScanCmd struct {
	Start   string
	Count   *uint32
	Options *NameScanOptions
}
