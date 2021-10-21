package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"strings"
	"syscall/js"

	"github.com/o1egl/govatar"
	uuid "github.com/satori/go.uuid"
)

var g = govatar.MALE

func generateAvatar(n string) (string, error) {
	img, err := govatar.GenerateForUsername(g, n)
	if err != nil {
		return "", err
	}

	var buf = new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func generateRandomAvatar(name js.Value) (string, error) {
	n := uuid.NewV4().String()
	name.Set("value", n)
	return generateAvatar(n)
}

func changeGender(gender js.Value) {
	if strings.ToLower(gender.Get("value").String()) == "male" {
		g = govatar.MALE
	} else {
		g = govatar.FEMALE
	}
}

func main() {
	doc := js.Global().Get("document")
	img := doc.Call("getElementById", "img")
	btn := doc.Call("getElementById", "generate")
	name := doc.Call("getElementById", "name")
	btn.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		i, err := generateAvatar(name.Get("value").String())
		fmt.Println(name.Get("value").String())
		if err != nil {
			fmt.Println(err)
			return nil
		}
		img.Set("src", fmt.Sprintf("data:image/png;base64,%s", i))
		return nil
	}))
	btn.Set("disabled", false)

	btnR := doc.Call("getElementById", "generateRandom")
	btnR.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		i, err := generateRandomAvatar(name)
		fmt.Println(name.Get("value").String())
		if err != nil {
			fmt.Println(err)
			return nil
		}
		img.Set("src", fmt.Sprintf("data:image/png;base64,%s", i))
		return nil
	}))
	btnR.Set("disabled", false)

	gender := doc.Call("getElementById", "gender")
	gender.Call("addEventListener", "change", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		changeGender(gender)
		return nil
	}))
	gender.Set("disabled", false)
	changeGender(gender)

	js.Global().Get("document").Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	}))

	select {}
}
