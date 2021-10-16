package input

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func ReadDirectly() string {
	msg, err := Read()
	if err != nil {
		panic(err)
	}
	return msg
}
func Read() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	var input string
	if runtime.GOOS == "windows" {
		input = strings.TrimRight(string(password[:len(password)-1]), "\r\n")
	} else {
		input = strings.TrimRight(string(password[:len(password)-1]), "\n")
	}
	return input, nil
}

func ShowMsgAndReadDirect(msg string) string {
	m, err := ShowMsgAndRead(msg)
	if err != nil {
		panic(err)
	}
	return m
}

func ShowMsgAndRead(msg string) (string, error) {
	print(msg)
	return Read()
}

func WaitForPressAnykey(msg string){
	fmt.Println(msg)
	fmt.Scanln() // wait for Enter Key
}

