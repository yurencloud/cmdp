package main

import (
	"github.com/yurencloud/cmdp/src"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "cmdp"
	app.Version = "2.0.0"

	app.Commands = []cli.Command{
		// cmd
		src.SearchCommand,
		src.VersionCommand,
		src.CreateCmdCommand,
		src.DeleteCmdCommand,
		src.ExecCommand,
		// user
		src.RegisterCommand,
		src.LoginCommand,
		src.LogoutCommand,
		src.ResetPasswordCommand,
		src.UpdateInfoCommand,
		// file
		src.PushFileCommand,
		src.PullFileCommand,
		src.FindFileCommand,
		src.RemoveFileCommand,
		// star
		src.StarCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
