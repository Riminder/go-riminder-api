package main

import (
	"fmt"

	"github.com/Xalrandion/go-riminder-api/riminder"
)

func main() {
	api := riminder.New("ask_874138568ebde822652c3ddf2218333a")
	resp, err := api.Source().Get(map[string]interface{}{"source_id": "5710d9eae87ef13b26eeaae8e0b20ea98b29774e"})
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("resp: ", resp.SourceID)
}
