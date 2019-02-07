package main

import (
	"encoding/hex"
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	cmd := exec.Command("osascript", "-e", "set clipboardData to the clipboard as record")
	r, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	binData := []byte(r)
	hexData := string(binData)
	hexData = strings.Split(hexData, ",")[0]
	hexData = strings.Split(hexData, ":")[1]
	hexData = strings.NewReplacer("«data", "", "»", "", "\r\n", "", "\r", "", "\n", "").Replace(hexData)
	hexData = hexData[5:len(hexData)]

	strData, err := hex.DecodeString(hexData)
	if err != nil {
		panic(err)
	}

	// Can't set huge data such as a filemaker's tab controll to clipboard...
	strTxt := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "\"", "\\\"").Replace(string(strData))
	err = exec.Command("osascript", "-e", fmt.Sprintf("set the clipboard to \"%s\"", strTxt)).Run()
	if err != nil {
		panic(err)
	}

}
