package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const terraform string = "/usr/local/bin/terraform"

type WorkSpaceEnv struct {
	envPrefix string `required:"true"`
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func CurrentWorkSpace() string{
	currentString := regexp.MustCompile("\\*.*")

	out, err := exec.Command(terraform, "workspace", "list").Output()
	if err != nil {
		fmt.Println(string(out))
		fmt.Println(err)
		os.Exit(2)
	}

	return "current workspace : " + currentString.FindString(strings.TrimSpace(string(out)))
}

func ChangeWorkSpace(w WorkSpaceEnv) bool{
	_, err := exec.Command(terraform, "workspace", "select", w.envPrefix).Output()
	if err != nil {
		fmt.Println("Environ Not Found (", w.envPrefix, ")")
		fmt.Println(err)
		os.Exit(2)
	}
	return true
}

func Contains(args []string, option string) bool {
	for _, v := range args {
		if option == v {
			return true
		}
	}
	return false
}

func DryRunExec(args []string) string{
	plan := []string{}

	for _, v := range args{
		if v == "apply" {
			plan = append(plan,"plan")
		} else {
			plan = append(plan, v)
		}
	}

	out, err := exec.Command(terraform, plan...).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return string(out)
}

func CommandExec(args []string) string {
	// If 'apply' is included in the argument, add '-auto-approve' option
	if Contains(args, "apply") {
		fmt.Println(DryRunExec(args))
		// Exec Check
		fmt.Print("Do you want to run it?[y/N]: ")
		var execFlag string
		fmt.Scan(&execFlag)
		if execFlag != "y" {
			os.Exit(0)
		} else {
			// silent exec
			args = append(args, "-auto-approve")
		}
	}

	// Exec Terraform Command
	out, err := exec.Command(terraform, args...).Output()
	if err != nil {
		fmt.Println(string(out))
		fmt.Println(err)
		os.Exit(2)
	}
	return strings.TrimSpace(string(out))
}

func main()  {
	// Terraform Command Exist Check
	if Exists(terraform) == false {
		fmt.Println("Error: Terraform Not Found (/usr/local/bin/terraform).")
		os.Exit(4)
	}
	// Get environment
	optEnvPrefix  := flag.String("env", "", "Environment")

	flag.Parse()
	workSpaceEnv := WorkSpaceEnv{}
	// Set Environment on Struct
	workSpaceEnv.envPrefix = *optEnvPrefix

	// Get All Arguments
	args := flag.Args()

	// Change WorkSpace to Designated Value
	ChangeWorkSpace(workSpaceEnv)
	fmt.Println(CurrentWorkSpace())


	// Exec Terraform Command And Output STDIO
	fmt.Println(CommandExec(args))
}
