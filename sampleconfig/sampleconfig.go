// Copyright (c) 2017-2022 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package sampleconfig

import (
	_ "embed"
)

// sampleVeraConf is a string containing the commented example config for vera.
//
//go:embed sample-vera.conf
var sampleVeraConf string

// sampleVeratConf is a string containing the commented example config for
// verat.
//
//go:embed sample-vera-client.conf
var sampleVeratConf string

// Vera returns a string containing the commented example config for vera.
func Vera() string {
	return sampleVeraConf
}

// FileContents returns a string containing the commented example config for
// vera.
//
// Deprecated: Use the [Vera] function instead.
func FileContents() string {
	return Vera()
}

// Verat returns a string containing the commented example config for verat.
func Verat() string {
	return sampleVeratConf
}

// VeratSampleConfig is a string containing the commented example config for
// verat.
//
// Deprecated: Use the [Verat] function instead.
var VeratSampleConfig = Verat()
