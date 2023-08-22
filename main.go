package main

import (
	"io/fs"
	"os"
)

func main() {
	if err := os.WriteFile("/data/hello.txt", []byte("Hello World"), fs.ModePerm); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if err := os.Mkdir("/data/meow", fs.ModePerm); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if err := os.WriteFile("/data/meow/meow.txt", []byte("Meow Meow"), fs.ModePerm); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if err := os.Mkdir("/data/meow/woof", fs.ModePerm); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if err := os.WriteFile("/data/meow/woof/woof.txt", []byte("Woof Woof"), fs.ModePerm); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if err := os.Mkdir("/data/meow/woof/moo", fs.ModePerm); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if err := os.WriteFile("/data/meow/woof/moo/moo.txt", []byte("Moo Moo"), fs.ModePerm); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	os.Exit(1)
}
