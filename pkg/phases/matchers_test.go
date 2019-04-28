package phases

import (
	"fmt"

	"github.com/onsi/gomega/types"
)

func MatchCommand(expected interface{}) types.GomegaMatcher {
	return &CommandMatcher{
		expected: expected,
	}
}

type CommandMatcher struct {
	expected interface{}
}

func (matcher *CommandMatcher) Match(actual interface{}) (success bool, err error) {
	sys, ok := actual.(*SystemConfig)
	if !ok {
		return false, fmt.Errorf("CommandMatcher matcher expects a SystemConfig")
	}
	for _, cmd := range sys.PreCommands {
		if cmd.Cmd == matcher.expected {
			return true, nil
		}
	}
	for _, cmd := range sys.Commands {
		if cmd.Cmd == matcher.expected {
			return true, nil
		}
	}
	return false, nil
}

func (matcher *CommandMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected \t%#v to contain command: \t%#v", actual, matcher.expected)
}

func (matcher *CommandMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected \t%#v to NOT contain command: \t%#v", actual, matcher.expected)
}
