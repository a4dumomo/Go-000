package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		log.Printf("connect fail:%v", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err: %v\n", err)
			break
		}
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}
		_, err = conn.Write([]byte(trimmedInput))

		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}
