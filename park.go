// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gopark

import (
	"unsafe"

	"github.com/lemon-mint/libuseful"
)

// A waitReason explains why a goroutine has been stopped.
// See gopark. Do not re-use waitReasons, add new ones.
type waitReason uint8

const (
	waitReasonZero waitReason = iota // ""
)

// Puts the current goroutine into a waiting state and calls unlockf on the
// system stack.
//
// If unlockf returns false, the goroutine is resumed.
//
// unlockf must not access this G's stack, as it may be moved between
// the call to gopark and the call to unlockf.
//
// Note that because unlockf is called after putting the G into a waiting
// state, the G may have already been readied by the time unlockf is called
// unless there is external synchronization preventing the G from being
// readied. If unlockf returns false, it must guarantee that the G cannot be
// externally readied.
//
// Reason explains why the goroutine has been parked. It is displayed in stack
// traces and heap dumps. Reasons should be unique and descriptive. Do not
// re-use reasons, add new ones.
//go:linkname gopark runtime.gopark
func gopark(unlockf func(gp unsafe.Pointer, lock unsafe.Pointer) bool, lock unsafe.Pointer, reason waitReason, traceEv byte, traceskip int)

//go:linkname goready runtime.goready
func goready(gp unsafe.Pointer, traceskip int)

func unlockf(gp unsafe.Pointer, lock unsafe.Pointer) bool {
	return true
}

func Freeze() {
	gopark(unlockf, nil, waitReasonZero, 0, 0)
}

func Melt(gp unsafe.Pointer) {
	goready(gp, 0)
}

func GetG() (gp unsafe.Pointer) {
	return libuseful.GetG()
}
