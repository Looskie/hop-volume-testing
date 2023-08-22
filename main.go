package main

import (
	"io/fs"
	"os"
)

func main() {
	os.WriteFile("hello.txt", []byte("Hello World"), fs.ModePerm)

	os.Mkdir("meow", fs.ModePerm)
	os.WriteFile("meow/meow.txt", []byte("Meow Meow"), fs.ModePerm)

	os.Mkdir("meow/woof", fs.ModePerm)
	os.WriteFile("meow/woof/woof.txt", []byte("Woof Woof"), fs.ModePerm)

	os.Mkdir("meow/woof/moo", fs.ModePerm)
	os.WriteFile("meow/woof/moo/moo.txt", []byte("Moo Moo"), fs.ModePerm)
}
