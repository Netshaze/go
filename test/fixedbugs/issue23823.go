// errorcheck

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p

type I1 = interface {
	I2
}

// BAD: type loop should mention I1; see also #41669
type I2 interface { // ERROR "invalid recursive type I2\n\tLINE: I2 refers to\n\tLINE: I2$|invalid recursive type"
	I1
}
