package src

import "time"

type User struct {
	Id        int
	Username  string
	Info      string
	Official  bool
	CreatedAt time.Time
}

type File struct {
	Id        int
	UserId    int
	Name      string // file name
	Content   string // you can save txt,json,markdown,sh... here
	Comment   string // comment for Cmd
	Keyword   string // keyword for Cmd
	Private   bool   // is private cmd, default false
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 所有命令,代码 都保存在这个表中
type Cmd struct {
	Id        int
	UserId    int
	Cmd       string
	Keyword   string
	Comment   string
	Private   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResult struct {
	Username  string
	Info      string
	CreatedAt time.Time
	CmdCount  int
	FileCount int
	StarCount int
}
