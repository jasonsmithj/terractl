package main

import (
	"regexp"
	"testing"
)

type WorkSpaceEnvMock struct {
	envPrefix string
}

func TestExists(t *testing.T) {
	if true != Exists("/usr/local/bin/terraform") {
		t.Fatal("failed test")
	}
}

func TestWorkspace(t *testing.T) {
	r := regexp.MustCompile(`current workspace:*`)
	if true != r.MatchString(CurrentWorkSpace()) {
		t.Fatal("failed test")
	}
}

func TestCommandExec(t *testing.T) {
	r := regexp.MustCompile(`Terraform v*`)
	cmd  := []string{"version"}
	if true != r.MatchString(CommandExec(cmd)) {
		t.Fatal("failed test")
	}
}

func TestChangeWorkSpace(t *testing.T) {
	workSpaceEnv := WorkSpaceEnvMock{}
	workSpaceEnv.envPrefix = "default"
	if true != ChangeWorkSpace(WorkSpaceEnv(workSpaceEnv)) {
		t.Fatal("failed test")
	}
}

func TestCurrentWorkSpace(t *testing.T) {
	if "current workspace : * default" != CurrentWorkSpace() {
		t.Fatal("faild test")
	}
}

func TestContainsTrue(t *testing.T) {
	arrTrue := []string{"plan", "apply", "test"}
	if true != Contains(arrTrue, "test") {
		t.Fatal("faild test")
	}

	arrFalse := []string{"plan", "apply", "test"}
	if false != Contains(arrFalse, "workspace") {
		t.Fatal("faild test")
	}
}

