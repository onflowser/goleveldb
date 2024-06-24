// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//go:build js || wasm
// +build js wasm

package storage

import (
	"os"
)

type jsFileLock struct{}

var _ fileLock = &jsFileLock{}

func (fl *jsFileLock) release() error {
	panic("Unimplemented")
}

func newFileLock(path string, readOnly bool) (fl fileLock, err error) {
	panic("Unimplemented")
}

func setFileLock(f *os.File, readOnly, lock bool) error {
	panic("Unimplemented")
}

func rename(oldpath, newpath string) error {

	panic("Unimplemented")
}

func isErrInvalid(err error) bool {
	panic("Unimplemented")
}

func syncDir(name string) error {
	panic("Unimplemented")
}
