// +build V7

// Code generated by counterfeiter. DO NOT EDIT.
package rpcfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/commandregistry"
	"code.cloudfoundry.org/cli/plugin/v7/rpc"
)

type FakeCommandRunner struct {
	CommandStub        func([]string, commandregistry.Dependency, bool) error
	commandMutex       sync.RWMutex
	commandArgsForCall []struct {
		arg1 []string
		arg2 commandregistry.Dependency
		arg3 bool
	}
	commandReturns struct {
		result1 error
	}
	commandReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCommandRunner) Command(arg1 []string, arg2 commandregistry.Dependency, arg3 bool) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.commandMutex.Lock()
	ret, specificReturn := fake.commandReturnsOnCall[len(fake.commandArgsForCall)]
	fake.commandArgsForCall = append(fake.commandArgsForCall, struct {
		arg1 []string
		arg2 commandregistry.Dependency
		arg3 bool
	}{arg1Copy, arg2, arg3})
	fake.recordInvocation("Command", []interface{}{arg1Copy, arg2, arg3})
	fake.commandMutex.Unlock()
	if fake.CommandStub != nil {
		return fake.CommandStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.commandReturns
	return fakeReturns.result1
}

func (fake *FakeCommandRunner) CommandCallCount() int {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	return len(fake.commandArgsForCall)
}

func (fake *FakeCommandRunner) CommandCalls(stub func([]string, commandregistry.Dependency, bool) error) {
	fake.commandMutex.Lock()
	defer fake.commandMutex.Unlock()
	fake.CommandStub = stub
}

func (fake *FakeCommandRunner) CommandArgsForCall(i int) ([]string, commandregistry.Dependency, bool) {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	argsForCall := fake.commandArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCommandRunner) CommandReturns(result1 error) {
	fake.commandMutex.Lock()
	defer fake.commandMutex.Unlock()
	fake.CommandStub = nil
	fake.commandReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandRunner) CommandReturnsOnCall(i int, result1 error) {
	fake.commandMutex.Lock()
	defer fake.commandMutex.Unlock()
	fake.CommandStub = nil
	if fake.commandReturnsOnCall == nil {
		fake.commandReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.commandReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCommandRunner) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ rpc.CommandRunner = new(FakeCommandRunner)
