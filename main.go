package main

import (
	"github.com/urfave/cli"
	"github.com/yurencloud/cmdp/src"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "cmdp"
	app.Version = "2.1.0"

	app.Commands = []cli.Command{
		// cmd
		src.SearchCommand,
		src.VersionCommand,
		src.CreateCmdCommand,
		src.DeleteCmdCommand,
		src.ExecCommand,
		src.ForkCmdCommand,
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
		src.ForkFileCommand,
		// star
		src.StarCommand,
		src.UpdateCommand,
		src.UserCommand,
		src.StarListCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
