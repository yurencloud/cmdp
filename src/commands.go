package src

import (
	"fmt"
	"github.com/urfave/cli"
)

var (
	// cmd
	VersionCommand = cli.Command{
		Name:    "version",
		Usage:   "show version",
		Aliases: []string{"v"},
		Action: func(ctx *cli.Context) {
			fmt.Println("cmdp version 2.0.0")
		},
	}
	SearchCommand = cli.Command{
		Name:    "search",
		Aliases: []string{"s"},
		Usage:   "search command, code, account, text, etc.",
		Action:  SearchAction,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "page,p",
				Usage: "set page number",
				Value: 1,
			},
			cli.IntFlag{
				Name:  "size,s",
				Usage: "set page size",
				Value: 20,
			},
			cli.BoolFlag{
				Name:  "all,a",
				Usage: "get all commands list",
			},
		},
	}
	CreateCmdCommand = cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "create command to remote",
		Action:  CreateCmdAction,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "private,p",
				Usage: "set private",
			},
		},
	}
	DeleteCmdCommand = cli.Command{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "delete command by id",
		Action:  DeleteCmdAction,
	}
	ExecCommand = cli.Command{
		Name:    "exec",
		Aliases: []string{"e"},
		Usage:   "exec command",
		Action:  ExecAction,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "file,f",
				Usage: "exec command from file",
			},
			cli.BoolFlag{
				Name:  "print,p",
				Usage: "not exec, only print command or file content",
			},
			cli.BoolFlag{
				Name:  "force,F",
				Usage: "force exec other user's dangerous command!",
			},
		},
	}
	// user
	RegisterCommand = cli.Command{
		Name:   "register",
		Usage:  "user register",
		Action: RegisterAction,
	}
	LoginCommand = cli.Command{
		Name:   "login",
		Usage:  "login by username and password",
		Action: LoginAction,
	}
	LogoutCommand = cli.Command{
		Name:   "logout",
		Usage:  "logout",
		Action: LogoutAction,
	}
	ResetPasswordCommand = cli.Command{
		Name:   "reset",
		Usage:  "reset password",
		Action: ResetPasswordAction,
	}
	UpdateInfoCommand = cli.Command{
		Name:   "info",
		Usage:  "update user introduction",
		Action: UpdateInfoAction,
	}

	// file
	PushFileCommand = cli.Command{
		Name:    "push",
		Aliases: []string{"p"},
		Usage:   "push your file to remote",
		Action:  PushFileAction,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "private,p",
				Usage: "set private",
			},
		},
	}
	PullFileCommand = cli.Command{
		Name:    "pull",
		Aliases: []string{"l", "pl", "P"},
		Usage:   "pull your file from remote",
		Action:  PullFileAction,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "print,p",
				Usage: "not download, only print file content",
			},
		},
	}
	FindFileCommand = cli.Command{
		Name:    "find",
		Aliases: []string{"f"},
		Usage:   "find your files from remote by keyword",
		Action:  FindFileAction,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "page,p",
				Usage: "set page number",
				Value: 1,
			},
			cli.IntFlag{
				Name:  "size,s",
				Usage: "set page size",
				Value: 20,
			},
			cli.BoolFlag{
				Name:  "all,a",
				Usage: "get all files list",
			},
		},
	}
	RemoveFileCommand = cli.Command{
		Name:    "remove",
		Aliases: []string{"r"},
		Usage:   "remove your remote file by id",
		Action:  RemoveFileAction,
	}
	// star
	StarCommand = cli.Command{
		Name:   "star",
		Usage:  "star other user",
		Action: StarAction,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "page,p",
				Usage: "set page number",
				Value: 1,
			},
			cli.IntFlag{
				Name:  "size,s",
				Usage: "set page size",
				Value: 20,
			},
			cli.IntFlag{
				Name:  "delete,d",
				Usage: "delete a star",
				Value: 0,
			},
		},
	}
)
