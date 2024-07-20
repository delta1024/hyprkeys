package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func main() {
	hyprctl := exec.Command("hyprctl", "binds", "-j")
	stdout, err := hyprctl.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := hyprctl.Start(); err != nil {
		panic(err)
	}
	decode := json.NewDecoder(stdout)
	var vals []hyprKey
	err = decode.Decode(&vals)
	if err != nil {
		panic(err)
	}
	if err := hyprctl.Wait(); err != nil {
		panic(err)
	}
	for _, key := range vals {
		fmt.Printf("%+v\n", key)
	}
}
