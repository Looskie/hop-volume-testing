package main

import (
	"os"
)

func main() {
	println("running")
	if err := copy("./mock", "/data"); err != nil {
		println(err.Error())
	}

	os.Exit(0)
}

func copy(src string, dest string) error {
	if err := os.Mkdir(dest, 0755); err != nil {
		return err
	}

	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			if err := copy(src+"/"+f.Name(), dest+"/"+f.Name()); err != nil {
				return err
			}
		}

		if !f.IsDir() {
			content, err := os.ReadFile(src + "/" + f.Name())
			if err != nil {
				return err
			}

			if err := os.WriteFile(dest+"/"+f.Name(), content, 0755); err != nil {
				return err
			}
		}
	}

	return nil
}
