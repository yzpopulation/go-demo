package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"go-demo/log"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	res, _, err := reader.ReadLine()
	if nil != err {
		fmt.Println("reader.ReadLine() error:", err)
	}
	log.Debug(string(res))
	log.Info(string(base64.StdEncoding.EncodeToString(res)))

}
