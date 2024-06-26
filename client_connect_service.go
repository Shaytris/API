// © chlove 2024 / This code is licenced under the MIT License for Shaytris
// This code's GUI version will be published once bamboo is added to Shaytris Nexus.
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter your name:")

    for {
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if input == "sys" {
            fmt.Println("Backend account can't be logged in to via the frontend")
            break
        } else {
            fmt.Println("Invalid name. Try again:")
        }
    }
}
