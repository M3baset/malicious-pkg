package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("cat", "/root/flag.txt")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Flag:", string(output))
    }
}
