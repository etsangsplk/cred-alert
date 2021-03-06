// Code generated by counterfeiter. DO NOT EDIT.
package queuefakes

import (
	"cred-alert/queue"
	"sync"
)

type FakeEnqueuer struct {
	EnqueueStub        func(queue.Task) error
	enqueueMutex       sync.RWMutex
	enqueueArgsForCall []struct {
		arg1 queue.Task
	}
	enqueueReturns struct {
		result1 error
	}
	enqueueReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnqueuer) Enqueue(arg1 queue.Task) error {
	fake.enqueueMutex.Lock()
	ret, specificReturn := fake.enqueueReturnsOnCall[len(fake.enqueueArgsForCall)]
	fake.enqueueArgsForCall = append(fake.enqueueArgsForCall, struct {
		arg1 queue.Task
	}{arg1})
	fake.recordInvocation("Enqueue", []interface{}{arg1})
	fake.enqueueMutex.Unlock()
	if fake.EnqueueStub != nil {
		return fake.EnqueueStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.enqueueReturns.result1
}

func (fake *FakeEnqueuer) EnqueueCallCount() int {
	fake.enqueueMutex.RLock()
	defer fake.enqueueMutex.RUnlock()
	return len(fake.enqueueArgsForCall)
}

func (fake *FakeEnqueuer) EnqueueArgsForCall(i int) queue.Task {
	fake.enqueueMutex.RLock()
	defer fake.enqueueMutex.RUnlock()
	return fake.enqueueArgsForCall[i].arg1
}

func (fake *FakeEnqueuer) EnqueueReturns(result1 error) {
	fake.EnqueueStub = nil
	fake.enqueueReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEnqueuer) EnqueueReturnsOnCall(i int, result1 error) {
	fake.EnqueueStub = nil
	if fake.enqueueReturnsOnCall == nil {
		fake.enqueueReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.enqueueReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEnqueuer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.enqueueMutex.RLock()
	defer fake.enqueueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEnqueuer) recordInvocation(key string, args []interface{}) {
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

var _ queue.Enqueuer = new(FakeEnqueuer)
