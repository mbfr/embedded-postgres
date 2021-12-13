//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package embeddedpostgres

import (
	"syscall"
)

// ProcAttr sets what to run the user as
func (c Config) ProcAttr(procAttr *syscall.SysProcAttr) Config {
	c.postgresProcAttr = procAttr
	return c
}
