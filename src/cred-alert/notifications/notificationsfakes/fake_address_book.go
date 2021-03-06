// Code generated by counterfeiter. DO NOT EDIT.
package notificationsfakes

import (
	"context"
	"cred-alert/notifications"
	"sync"

	"code.cloudfoundry.org/lager"
)

type FakeAddressBook struct {
	AddressForRepoStub        func(ctx context.Context, logger lager.Logger, owner, name string) []notifications.Address
	addressForRepoMutex       sync.RWMutex
	addressForRepoArgsForCall []struct {
		ctx    context.Context
		logger lager.Logger
		owner  string
		name   string
	}
	addressForRepoReturns struct {
		result1 []notifications.Address
	}
	addressForRepoReturnsOnCall map[int]struct {
		result1 []notifications.Address
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAddressBook) AddressForRepo(ctx context.Context, logger lager.Logger, owner string, name string) []notifications.Address {
	fake.addressForRepoMutex.Lock()
	ret, specificReturn := fake.addressForRepoReturnsOnCall[len(fake.addressForRepoArgsForCall)]
	fake.addressForRepoArgsForCall = append(fake.addressForRepoArgsForCall, struct {
		ctx    context.Context
		logger lager.Logger
		owner  string
		name   string
	}{ctx, logger, owner, name})
	fake.recordInvocation("AddressForRepo", []interface{}{ctx, logger, owner, name})
	fake.addressForRepoMutex.Unlock()
	if fake.AddressForRepoStub != nil {
		return fake.AddressForRepoStub(ctx, logger, owner, name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addressForRepoReturns.result1
}

func (fake *FakeAddressBook) AddressForRepoCallCount() int {
	fake.addressForRepoMutex.RLock()
	defer fake.addressForRepoMutex.RUnlock()
	return len(fake.addressForRepoArgsForCall)
}

func (fake *FakeAddressBook) AddressForRepoArgsForCall(i int) (context.Context, lager.Logger, string, string) {
	fake.addressForRepoMutex.RLock()
	defer fake.addressForRepoMutex.RUnlock()
	return fake.addressForRepoArgsForCall[i].ctx, fake.addressForRepoArgsForCall[i].logger, fake.addressForRepoArgsForCall[i].owner, fake.addressForRepoArgsForCall[i].name
}

func (fake *FakeAddressBook) AddressForRepoReturns(result1 []notifications.Address) {
	fake.AddressForRepoStub = nil
	fake.addressForRepoReturns = struct {
		result1 []notifications.Address
	}{result1}
}

func (fake *FakeAddressBook) AddressForRepoReturnsOnCall(i int, result1 []notifications.Address) {
	fake.AddressForRepoStub = nil
	if fake.addressForRepoReturnsOnCall == nil {
		fake.addressForRepoReturnsOnCall = make(map[int]struct {
			result1 []notifications.Address
		})
	}
	fake.addressForRepoReturnsOnCall[i] = struct {
		result1 []notifications.Address
	}{result1}
}

func (fake *FakeAddressBook) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addressForRepoMutex.RLock()
	defer fake.addressForRepoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAddressBook) recordInvocation(key string, args []interface{}) {
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

var _ notifications.AddressBook = new(FakeAddressBook)
