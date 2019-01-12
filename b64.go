// Copyright 2018 Ewout van Mansom. All rights reserved.
// Use of this source code is governed by a AGPLv3
// license that can be found in the LICENSE file.

// Package b64 provides an implementation of Sulake's FUSE
// tetrasexagesimal numeric system.
//
// It typically uses two ASCII characters between decimal
// indexes 64 (@) and 127 (DEL control character) to produce
// a two-character representation of a number between 0 and 4095.
package b64

// Encode transforms integer i into a tetrasexagesimal representation.
func Encode(i int) []byte {
	return []byte{byte(i/64 + 64), byte(i&63 + 64)}
}

// Decode transforms the tetrasexagesimal representation into an integer.
func Decode(b []byte) int {
	return int(64*(int(b[0])&63) + (int(b[1]) & 63))
}
