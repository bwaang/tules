package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"gopkg.in/pipe.v2"
	"os"
	"os/exec"
	"strings"
	"sync"
)

const (
	APP  = "gitb"
	DESC = "Utility script for common git functions"
)

func init() {

}

func main() {
	// testFlag
	var testflag string

	app := cli.NewApp()
	app.Name = APP
	app.Usage = DESC

	// To enable bash completion in your ~/.bash_profile add:
	// PROG=gitta source <path to this executable>
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "testflag, tf",
			Value:       "/home/path/",
			Usage:       "This is used for testing the flags",
			EnvVar:      "GOPATH",
			Destination: &testflag,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "add-upstream",
			Aliases: []string{"a"},
			Usage:   "adds the upstream remote to the stack",
			Action: func(c *cli.Context) {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Github user name: ")
				username, _ := reader.ReadString('\n')
				username = strings.TrimRight(username, "\n")

				reader = bufio.NewReader(os.Stdin)
				fmt.Print("Upstream user name: ")
				upstream, _ := reader.ReadString('\n')
				upstream = strings.TrimRight(upstream, "\n")

				p := pipe.Line(
					pipe.Exec("git", parseArgs("remote -v")...),
					pipe.Exec("grep", "origin"),
					pipe.Exec("awk", "{ print $2 }"),
					pipe.Exec("sed", "s/"+username+"/"+upstream+"/g"),
					pipe.Exec("head", "-1"),
				)

				_, url := runPipe(p)
				runCmd("git remote add upstream " + url)
			},
		},
		{
			Name:    "origin-diff",
			Aliases: []string{"o"},
			Usage:   "compare diff with origin",
			Action: func(c *cli.Context) {
				p := pipe.Line(
					pipe.Exec("git", parseArgs("branch -a")...),
					pipe.Exec("grep", "'*'"),
					pipe.Exec("awk", "{ print $NF }"),
				)

				_, currentBranch := runPipe(p)

				fmt.Println(currentBranch)
				runCmd("git diff --name-status origin.." + currentBranch)
			},
		},
		{
			Name:    "log-pretty",
			Aliases: []string{"l"},
			Usage:   "print out pretty log",
			Action: func(c *cli.Context) {
				runCmd("git log --graph --oneline --decorate --all")
			},
		},
		{
			Name:    "sync-origin",
			Aliases: []string{"o"},
			Usage:   "sync all local branches to remote",
			Action: func(c *cli.Context) {
				runCmd("git remote update")
				runCmd("git remote prune origin")
			},
		},
		{
			Name:    "sync-upstream",
			Aliases: []string{"u"},
			Usage:   "synch current branch with upstream",
			Action: func(c *cli.Context) {
				currentBranch := outputCmd("git rev-parse --abbrev-ref HEAD")

				fmt.Println(currentBranch)
				runCmd("git fetch upstream")
				runCmd("git reset upstream/" + currentBranch)
			},
		},
		{
			Name:    "quick-commit-squash",
			Aliases: []string{"q"},
			Usage:   "sync all remote branches",
			Action: func(c *cli.Context) {
				runCmd("git commit -am Squashable")
				runCmd("git rebase -i HEAD~2")
			},
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "set push url to your fork instead of upstream",
			Action: func(c *cli.Context) {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Github user name: ")
				username, _ := reader.ReadString('\n')
				username = strings.TrimRight(username, "\n")

				reader = bufio.NewReader(os.Stdin)
				fmt.Print("Upstream user name: ")
				upstream, _ := reader.ReadString('\n')
				upstream = strings.TrimRight(username, "\n")

				p := pipe.Line(
					pipe.Exec("git", parseArgs("remote -v")...),
					pipe.Exec("awk", "{ print $2 }"),
					pipe.Exec("sed", "s/"+upstream+"/"+username+"/g"),
					pipe.Exec("head", "-1"),
				)

				_, url := runPipe(p)
				runCmd("git remote set-url origin --push " + url)
			},
		},
	}

	app.Action = func(c *cli.Context) {
		fmt.Println(testflag)
	}

	app.Run(os.Args)
	post()

}

func post() {
	var wg sync.WaitGroup
	wg.Wait()
}

// ----------------
// UTILITY COMMANDS
// ----------------

// mkCmd creates command and returns it, useful for iopiping
func mkCmd(cmd string) *exec.Cmd {
	// fmt.Printf("%s-shell$ %s\n", APP, cmd)
	head, args := parseCmd(cmd)

	return exec.Command(head, args...)
}

// runCmd runs the command as if it was natively on the machine
func runCmd(command string) {
	cmd := mkCmd(command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func outputCmd(command string) string {
	cmd := mkCmd(command)
	out, err := cmd.Output()
	errCheck(err)
	return string(out)
}

// runPipe runs the created pipeline
func runPipe(p pipe.Pipe) ([]byte, string) {
	out, err := pipe.CombinedOutput(p)
	errCheck(err)
	return out, string(out)
}

func parseCmd(cmd string) (string, []string) {
	command := strings.Fields(cmd)
	return command[0], command[1:]
}

// parseArgs alias for converting argument string into an argument slice
func parseArgs(args string) []string {
	return strings.Fields(args)
}

// errCheck runs the standard error checking
func errCheck(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: ", err)
	}
}

// debugCmd just prints it out formatted
func debugCmd(head string, args ...string) {
	fmt.Printf("Command: %s; Args: %s\n", head, args)
}
