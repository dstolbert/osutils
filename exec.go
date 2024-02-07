package osutils

import (
	"io"
	"os/exec"
	"syscall"
)

func (o *osutils) CommandStart(cmd *exec.Cmd) error {
	return cmd.Start()
}

func (o *osutils) CommandWait(cmd *exec.Cmd) error {
	return cmd.Wait()
}

func (o *osutils) CommandOutput(cmd *exec.Cmd) ([]byte, error) {
	return cmd.Output()
}

func (o *osutils) SyscallKill(pid int, sig syscall.Signal) (err error) {
	return syscall.Kill(pid, sig)
}

func (o *osutils) StdoutPipe(cmd *exec.Cmd) (io.ReadCloser, error) {
	return cmd.StdoutPipe()
}
