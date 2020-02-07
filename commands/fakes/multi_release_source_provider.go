// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/kiln/commands"
	"github.com/pivotal-cf/kiln/fetcher"
	"github.com/pivotal-cf/kiln/internal/cargo"
)

type MultiReleaseSourceProvider struct {
	Stub        func(cargo.Kilnfile, bool) fetcher.MultiReleaseSource
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 cargo.Kilnfile
		arg2 bool
	}
	returns struct {
		result1 fetcher.MultiReleaseSource
	}
	returnsOnCall map[int]struct {
		result1 fetcher.MultiReleaseSource
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MultiReleaseSourceProvider) Spy(arg1 cargo.Kilnfile, arg2 bool) fetcher.MultiReleaseSource {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 cargo.Kilnfile
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("MultiReleaseSourceProvider", []interface{}{arg1, arg2})
	fake.mutex.Unlock()
	if fake.Stub != nil {
		return fake.Stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.returns.result1
}

func (fake *MultiReleaseSourceProvider) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *MultiReleaseSourceProvider) Calls(stub func(cargo.Kilnfile, bool) fetcher.MultiReleaseSource) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *MultiReleaseSourceProvider) ArgsForCall(i int) (cargo.Kilnfile, bool) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1, fake.argsForCall[i].arg2
}

func (fake *MultiReleaseSourceProvider) Returns(result1 fetcher.MultiReleaseSource) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 fetcher.MultiReleaseSource
	}{result1}
}

func (fake *MultiReleaseSourceProvider) ReturnsOnCall(i int, result1 fetcher.MultiReleaseSource) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 fetcher.MultiReleaseSource
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 fetcher.MultiReleaseSource
	}{result1}
}

func (fake *MultiReleaseSourceProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MultiReleaseSourceProvider) recordInvocation(key string, args []interface{}) {
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

var _ commands.MultiReleaseSourceProvider = new(MultiReleaseSourceProvider).Spy
