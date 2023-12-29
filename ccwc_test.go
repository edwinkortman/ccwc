package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"os/exec"
	"strings"
	"testing"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run tests")
	}
}

var result string

func iRunCCWC(command string) error {
	cmdSliced := strings.Fields(command)

	cmd := exec.Command(cmdSliced[0], cmdSliced[1:]...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	result = strings.TrimSpace(string(out))

	return nil
}

func theResultShouldContain(expectedResult string) error {
	if expectedResult == result {
		return nil
	}

	return fmt.Errorf("Expected: %s, but got: %s", expectedResult, result)
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.When(`^I run "([^"]*)"$`, iRunCCWC)
	sc.Then(`^the result should contain "([^"]*)"$`, theResultShouldContain)
}
