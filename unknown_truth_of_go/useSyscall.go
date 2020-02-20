package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println("My pid is", pid)
	uid, _, _ := syscall.Syscall(24, 0, 0, 0)
	fmt.Println("User ID:", uid)

	message := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}
	fd := 1
	syscall.Write(fd, message)

	fmt.Println("Using syscall.Exec()")
	command := "/bin/ls"
	env := os.Environ()
	for _, v := range env {
		kv := strings.Split(v, "=")
		if kv[0] == "USER" {
			fmt.Println("PATH env exists...")
		}
		fmt.Printf("%25s : %s\n", kv[0], kv[1])
	}
	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)
}
