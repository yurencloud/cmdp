package src

import "github.com/fatih/color"

type Respond struct {
	Status  string
	Message string
	Data    interface{}
}

type CmdsRespond struct {
	Status  string
	Message string
	Data    []Cmd
}
type StarsRespond struct {
	Status  string
	Message string
	Data    []StarUser
}
type CmdRespond struct {
	Status  string
	Message string
	Data    Cmd
}
type FilesRespond struct {
	Status  string
	Message string
	Data    []File
}
type UsersRespond struct {
	Status  string
	Message string
	Data    []UserResult
}
type FilesDownloadRespond struct {
	Status  string
	Message string
	Data    File
}

func printRespond(result Respond) {
	if result.Status == SUCCESS {
		color.Green(result.Message)
	} else {
		color.Red(result.Message)
	}
}
