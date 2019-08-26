package list

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/urfave/cli"
)

func listFiles(path *string, showHidden *bool) func(c *cli.Context) error {

	return func(c *cli.Context) error {

		// fmt.Printf("[DEBUG] Path: %s\n", *path)
		// fmt.Printf("[DEBUG] Show hidden: %v\n", *showHidden)

		files, err := ioutil.ReadDir(*path)
		if err != nil {
			return err
		}

		for _, file := range files {
			if !file.IsDir() {
				// fmt.Printf("[DEBUG] Path %s is hidden: %v\n", file.Name(), isHidden(file.Name()))
				if *showHidden {
					fmt.Println(file.Name())
				} else {
					hidden, hiddenErr := isHidden(file.Name())
					if hiddenErr != nil {
						return hiddenErr
					}
					if !hidden {
						fmt.Println(file.Name())
					}
				}
			}
		}

		return nil
	}
}

func listFolders(path *string, showHidden *bool) func(c *cli.Context) error {

	return func(c *cli.Context) error {

		// fmt.Printf("[DEBUG] Path: %s\n", *path)
		// fmt.Printf("[DEBUG] Show hidden: %v\n", *showHidden)

		files, err := ioutil.ReadDir(*path)
		if err != nil {
			return err
		}

		for _, file := range files {
			if file.IsDir() {
				// fmt.Printf("[DEBUG] Path %s is hidden: %v\n", file.Name(), isHidden(file.Name()))
				if *showHidden {
					fmt.Println(file.Name())
				} else {
					hidden, hiddenErr := isHidden(file.Name())
					if hiddenErr != nil {
						return hiddenErr
					}
					if !hidden {
						fmt.Println(file.Name())
					}
				}
			}
		}

		return nil
	}
}

func listAll(path *string) func(c *cli.Context) error {

	return func(c *cli.Context) error {

		err := filepath.Walk(*path, func(filepath string, info os.FileInfo, err error) error {
			fmt.Println(filepath)
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	}
}

func isHidden(filepath string) (bool, error) {

	if runtime.GOOS != "windows" {
		return isHiddenUnderUnixLinux(filepath), nil
		// } else {
		// 	return isHiddenUnderWindows(filepath)
	}

	// fmt.Println("[ERROR] Unable to check if file is hidden under Windows")
	return false, cli.NewExitError("[ERROR] Unable to check if file is hidden under Windows", 86)
}

// unix/linux file or directory that starts with . is hidden
func isHiddenUnderUnixLinux(filepath string) bool {
	return filepath[0:1] == "."
}

// windows file or directory has a specific attribute
// func isHiddenUnderWindows(filepath string) bool {
// 	pointer, err := syscall.UTF16PtrFromString(filepath)
// 	if err != nil {
// 		return false
// 	}
// 	attributes, err := syscall.GetFileAttributes(pointer)
// 	if err != nil {
// 		return false
// 	}
// 	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0
// }
