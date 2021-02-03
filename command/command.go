package command

import (
	"errors"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

func Exec(name string, duration time.Duration, args ...string) (string, error) {
	var (
		timeoutError error
		timer        = time.NewTicker(duration)
		timerClose   = make(chan bool, 1)
	)
	cmd := exec.Command(name, args...)
	cmd.Dir = filepath.Dir(name)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	defer cmd.Wait()
	go func(cmd *exec.Cmd) {
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); err != nil {
					timeoutError = errors.New("timeout kill error: " + err.Error())
				} else {
					timeoutError = errors.New(`timeout`)
				}
				return
			case <-timerClose:
				return
			}
		}
	}(cmd)
	outputBytes, _ := cmd.CombinedOutput()
	timerClose <- true
	if timeoutError != nil {
		return "", timeoutError
	}
	return string(outputBytes), nil
}
