package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

func CreateToken(token interface{}) {
	GoRoot := runtime.GOROOT()
	HOME := os.Getenv("HOME")
	if len(HOME) == 0 {
		HOME = GoRoot
	}
	if runtime.GOOS == "windows" {
		HOME = HOME + "\\.cmdp"
		exist, err := PathExists(HOME)
		if err != nil {
			fmt.Println(err)
		}
		if !exist {
			err := os.Mkdir(HOME, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("create or update cmdp_token at :" + HOME + "\\.cmdp_token")
		err = ioutil.WriteFile(HOME+"\\.cmdp_token", []byte(token.(string)), 0644)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		HOME = HOME + "/.cmdp"
		exist, err := PathExists(HOME)
		if err != nil {
			fmt.Println(err)
		}
		if !exist {
			err := os.Mkdir(HOME, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("create or update cmdp_token at :" + HOME + "/.cmdp_token")
		err = ioutil.WriteFile(HOME+"/.cmdp_token", []byte(token.(string)), 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ReadToken() ([]byte, error) {
	GoRoot := runtime.GOROOT()
	HOME := os.Getenv("HOME")
	if len(HOME) == 0 {
		HOME = GoRoot
	}
	if runtime.GOOS == "windows" {
		token, err := ioutil.ReadFile(HOME + "\\.cmdp\\.cmdp_token")
		if err != nil {
			fmt.Println(err)
		}
		return token, err
	} else {
		token, err := ioutil.ReadFile(HOME + "/.cmdp/.cmdp_token")
		if err != nil {
			fmt.Println(err)
		}
		return token, err
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
