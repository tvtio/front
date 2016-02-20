// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package catalog

import "os"

// CachePath is where the caché is stored on the system.
var CachePath = os.Getenv("CACHEPATH")
