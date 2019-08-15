package src

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func Exec(s string) (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", s)
	} else {
		cmd = exec.Command("/bin/bash", "-c", s)
	}
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	return out.String(), err
}

// TODO:: 增加执行sh文件的功能
