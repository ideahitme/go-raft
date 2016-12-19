// define cluster related constants
package cluster

import (
	"time"
)

const (
	BootTime         = time.Millisecond * 1000 // three seconds on the boot up
	HeartbeatTimeout = time.Millisecond * 300
)
