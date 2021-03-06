// Code generated by counterfeiter. DO NOT EDIT.
package revokfakes

import (
	"cred-alert/revok"
	"cred-alert/revokpb"
	"sync"

	"golang.org/x/net/context"
)

type FakeServer struct {
	GetCredentialCountsStub        func(context.Context, *revokpb.CredentialCountRequest) (*revokpb.CredentialCountResponse, error)
	getCredentialCountsMutex       sync.RWMutex
	getCredentialCountsArgsForCall []struct {
		arg1 context.Context
		arg2 *revokpb.CredentialCountRequest
	}
	getCredentialCountsReturns struct {
		result1 *revokpb.CredentialCountResponse
		result2 error
	}
	getCredentialCountsReturnsOnCall map[int]struct {
		result1 *revokpb.CredentialCountResponse
		result2 error
	}
	GetOrganizationCredentialCountsStub        func(context.Context, *revokpb.OrganizationCredentialCountRequest) (*revokpb.OrganizationCredentialCountResponse, error)
	getOrganizationCredentialCountsMutex       sync.RWMutex
	getOrganizationCredentialCountsArgsForCall []struct {
		arg1 context.Context
		arg2 *revokpb.OrganizationCredentialCountRequest
	}
	getOrganizationCredentialCountsReturns struct {
		result1 *revokpb.OrganizationCredentialCountResponse
		result2 error
	}
	getOrganizationCredentialCountsReturnsOnCall map[int]struct {
		result1 *revokpb.OrganizationCredentialCountResponse
		result2 error
	}
	GetRepositoryCredentialCountsStub        func(context.Context, *revokpb.RepositoryCredentialCountRequest) (*revokpb.RepositoryCredentialCountResponse, error)
	getRepositoryCredentialCountsMutex       sync.RWMutex
	getRepositoryCredentialCountsArgsForCall []struct {
		arg1 context.Context
		arg2 *revokpb.RepositoryCredentialCountRequest
	}
	getRepositoryCredentialCountsReturns struct {
		result1 *revokpb.RepositoryCredentialCountResponse
		result2 error
	}
	getRepositoryCredentialCountsReturnsOnCall map[int]struct {
		result1 *revokpb.RepositoryCredentialCountResponse
		result2 error
	}
	SearchStub        func(*revokpb.SearchQuery, revokpb.Revok_SearchServer) error
	searchMutex       sync.RWMutex
	searchArgsForCall []struct {
		arg1 *revokpb.SearchQuery
		arg2 revokpb.Revok_SearchServer
	}
	searchReturns struct {
		result1 error
	}
	searchReturnsOnCall map[int]struct {
		result1 error
	}
	BoshBlobsStub        func(context.Context, *revokpb.BoshBlobsRequest) (*revokpb.BoshBlobsResponse, error)
	boshBlobsMutex       sync.RWMutex
	boshBlobsArgsForCall []struct {
		arg1 context.Context
		arg2 *revokpb.BoshBlobsRequest
	}
	boshBlobsReturns struct {
		result1 *revokpb.BoshBlobsResponse
		result2 error
	}
	boshBlobsReturnsOnCall map[int]struct {
		result1 *revokpb.BoshBlobsResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServer) GetCredentialCounts(arg1 context.Context, arg2 *revokpb.CredentialCountRequest) (*revokpb.CredentialCountResponse, error) {
	fake.getCredentialCountsMutex.Lock()
	ret, specificReturn := fake.getCredentialCountsReturnsOnCall[len(fake.getCredentialCountsArgsForCall)]
	fake.getCredentialCountsArgsForCall = append(fake.getCredentialCountsArgsForCall, struct {
		arg1 context.Context
		arg2 *revokpb.CredentialCountRequest
	}{arg1, arg2})
	fake.recordInvocation("GetCredentialCounts", []interface{}{arg1, arg2})
	fake.getCredentialCountsMutex.Unlock()
	if fake.GetCredentialCountsStub != nil {
		return fake.GetCredentialCountsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getCredentialCountsReturns.result1, fake.getCredentialCountsReturns.result2
}

func (fake *FakeServer) GetCredentialCountsCallCount() int {
	fake.getCredentialCountsMutex.RLock()
	defer fake.getCredentialCountsMutex.RUnlock()
	return len(fake.getCredentialCountsArgsForCall)
}

func (fake *FakeServer) GetCredentialCountsArgsForCall(i int) (context.Context, *revokpb.CredentialCountRequest) {
	fake.getCredentialCountsMutex.RLock()
	defer fake.getCredentialCountsMutex.RUnlock()
	return fake.getCredentialCountsArgsForCall[i].arg1, fake.getCredentialCountsArgsForCall[i].arg2
}

func (fake *FakeServer) GetCredentialCountsReturns(result1 *revokpb.CredentialCountResponse, result2 error) {
	fake.GetCredentialCountsStub = nil
	fake.getCredentialCountsReturns = struct {
		result1 *revokpb.CredentialCountResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeServer) GetCredentialCountsReturnsOnCall(i int, result1 *revokpb.CredentialCountResponse, result2 error) {
	fake.GetCredentialCountsStub = nil
	if fake.getCredentialCountsReturnsOnCall == nil {
		fake.getCredentialCountsReturnsOnCall = make(map[int]struct {
			result1 *revokpb.CredentialCountResponse
			result2 error
		})
	}
	fake.getCredentialCountsReturnsOnCall[i] = struct {
		result1 *revokpb.CredentialCountResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeServer) GetOrganizationCredentialCounts(arg1 context.Context, arg2 *revokpb.OrganizationCredentialCountRequest) (*revokpb.OrganizationCredentialCountResponse, error) {
	fake.getOrganizationCredentialCountsMutex.Lock()
	ret, specificReturn := fake.getOrganizationCredentialCountsReturnsOnCall[len(fake.getOrganizationCredentialCountsArgsForCall)]
	fake.getOrganizationCredentialCountsArgsForCall = append(fake.getOrganizationCredentialCountsArgsForCall, struct {
		arg1 context.Context
		arg2 *revokpb.OrganizationCredentialCountRequest
	}{arg1, arg2})
	fake.recordInvocation("GetOrganizationCredentialCounts", []interface{}{arg1, arg2})
	fake.getOrganizationCredentialCountsMutex.Unlock()
	if fake.GetOrganizationCredentialCountsStub != nil {
		return fake.GetOrganizationCredentialCountsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getOrganizationCredentialCountsReturns.result1, fake.getOrganizationCredentialCountsReturns.result2
}

func (fake *FakeServer) GetOrganizationCredentialCountsCallCount() int {
	fake.getOrganizationCredentialCountsMutex.RLock()
	defer fake.getOrganizationCredentialCountsMutex.RUnlock()
	return len(fake.getOrganizationCredentialCountsArgsForCall)
}

func (fake *FakeServer) GetOrganizationCredentialCountsArgsForCall(i int) (context.Context, *revokpb.OrganizationCredentialCountRequest) {
	fake.getOrganizationCredentialCountsMutex.RLock()
	defer fake.getOrganizationCredentialCountsMutex.RUnlock()
	return fake.getOrganizationCredentialCountsArgsForCall[i].arg1, fake.getOrganizationCredentialCountsArgsForCall[i].arg2
}

func (fake *FakeServer) GetOrganizationCredentialCountsReturns(result1 *revokpb.OrganizationCredentialCountResponse, result2 error) {
	fake.GetOrganizationCredentialCountsStub = nil
	fake.getOrganizationCredentialCountsReturns = struct {
		result1 *revokpb.OrganizationCredentialCountResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeServer) GetOrganizationCredentialCountsReturnsOnCall(i int, result1 *revokpb.OrganizationCredentialCountResponse, result2 error) {
	fake.GetOrganizationCredentialCountsStub = nil
	if fake.getOrganizationCredentialCountsReturnsOnCall == nil {
		fake.getOrganizationCredentialCountsReturnsOnCall = make(map[int]struct {
			result1 *revokpb.OrganizationCredentialCountResponse
			result2 error
		})
	}
	fake.getOrganizationCredentialCountsReturnsOnCall[i] = struct {
		result1 *revokpb.OrganizationCredentialCountResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeServer) GetRepositoryCredentialCounts(arg1 context.Context, arg2 *revokpb.RepositoryCredentialCountRequest) (*revokpb.RepositoryCredentialCountResponse, error) {
	fake.getRepositoryCredentialCountsMutex.Lock()
	ret, specificReturn := fake.getRepositoryCredentialCountsReturnsOnCall[len(fake.getRepositoryCredentialCountsArgsForCall)]
	fake.getRepositoryCredentialCountsArgsForCall = append(fake.getRepositoryCredentialCountsArgsForCall, struct {
		arg1 context.Context
		arg2 *revokpb.RepositoryCredentialCountRequest
	}{arg1, arg2})
	fake.recordInvocation("GetRepositoryCredentialCounts", []interface{}{arg1, arg2})
	fake.getRepositoryCredentialCountsMutex.Unlock()
	if fake.GetRepositoryCredentialCountsStub != nil {
		return fake.GetRepositoryCredentialCountsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getRepositoryCredentialCountsReturns.result1, fake.getRepositoryCredentialCountsReturns.result2
}

func (fake *FakeServer) GetRepositoryCredentialCountsCallCount() int {
	fake.getRepositoryCredentialCountsMutex.RLock()
	defer fake.getRepositoryCredentialCountsMutex.RUnlock()
	return len(fake.getRepositoryCredentialCountsArgsForCall)
}

func (fake *FakeServer) GetRepositoryCredentialCountsArgsForCall(i int) (context.Context, *revokpb.RepositoryCredentialCountRequest) {
	fake.getRepositoryCredentialCountsMutex.RLock()
	defer fake.getRepositoryCredentialCountsMutex.RUnlock()
	return fake.getRepositoryCredentialCountsArgsForCall[i].arg1, fake.getRepositoryCredentialCountsArgsForCall[i].arg2
}

func (fake *FakeServer) GetRepositoryCredentialCountsReturns(result1 *revokpb.RepositoryCredentialCountResponse, result2 error) {
	fake.GetRepositoryCredentialCountsStub = nil
	fake.getRepositoryCredentialCountsReturns = struct {
		result1 *revokpb.RepositoryCredentialCountResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeServer) GetRepositoryCredentialCountsReturnsOnCall(i int, result1 *revokpb.RepositoryCredentialCountResponse, result2 error) {
	fake.GetRepositoryCredentialCountsStub = nil
	if fake.getRepositoryCredentialCountsReturnsOnCall == nil {
		fake.getRepositoryCredentialCountsReturnsOnCall = make(map[int]struct {
			result1 *revokpb.RepositoryCredentialCountResponse
			result2 error
		})
	}
	fake.getRepositoryCredentialCountsReturnsOnCall[i] = struct {
		result1 *revokpb.RepositoryCredentialCountResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeServer) Search(arg1 *revokpb.SearchQuery, arg2 revokpb.Revok_SearchServer) error {
	fake.searchMutex.Lock()
	ret, specificReturn := fake.searchReturnsOnCall[len(fake.searchArgsForCall)]
	fake.searchArgsForCall = append(fake.searchArgsForCall, struct {
		arg1 *revokpb.SearchQuery
		arg2 revokpb.Revok_SearchServer
	}{arg1, arg2})
	fake.recordInvocation("Search", []interface{}{arg1, arg2})
	fake.searchMutex.Unlock()
	if fake.SearchStub != nil {
		return fake.SearchStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.searchReturns.result1
}

func (fake *FakeServer) SearchCallCount() int {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	return len(fake.searchArgsForCall)
}

func (fake *FakeServer) SearchArgsForCall(i int) (*revokpb.SearchQuery, revokpb.Revok_SearchServer) {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	return fake.searchArgsForCall[i].arg1, fake.searchArgsForCall[i].arg2
}

func (fake *FakeServer) SearchReturns(result1 error) {
	fake.SearchStub = nil
	fake.searchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServer) SearchReturnsOnCall(i int, result1 error) {
	fake.SearchStub = nil
	if fake.searchReturnsOnCall == nil {
		fake.searchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.searchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServer) BoshBlobs(arg1 context.Context, arg2 *revokpb.BoshBlobsRequest) (*revokpb.BoshBlobsResponse, error) {
	fake.boshBlobsMutex.Lock()
	ret, specificReturn := fake.boshBlobsReturnsOnCall[len(fake.boshBlobsArgsForCall)]
	fake.boshBlobsArgsForCall = append(fake.boshBlobsArgsForCall, struct {
		arg1 context.Context
		arg2 *revokpb.BoshBlobsRequest
	}{arg1, arg2})
	fake.recordInvocation("BoshBlobs", []interface{}{arg1, arg2})
	fake.boshBlobsMutex.Unlock()
	if fake.BoshBlobsStub != nil {
		return fake.BoshBlobsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.boshBlobsReturns.result1, fake.boshBlobsReturns.result2
}

func (fake *FakeServer) BoshBlobsCallCount() int {
	fake.boshBlobsMutex.RLock()
	defer fake.boshBlobsMutex.RUnlock()
	return len(fake.boshBlobsArgsForCall)
}

func (fake *FakeServer) BoshBlobsArgsForCall(i int) (context.Context, *revokpb.BoshBlobsRequest) {
	fake.boshBlobsMutex.RLock()
	defer fake.boshBlobsMutex.RUnlock()
	return fake.boshBlobsArgsForCall[i].arg1, fake.boshBlobsArgsForCall[i].arg2
}

func (fake *FakeServer) BoshBlobsReturns(result1 *revokpb.BoshBlobsResponse, result2 error) {
	fake.BoshBlobsStub = nil
	fake.boshBlobsReturns = struct {
		result1 *revokpb.BoshBlobsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeServer) BoshBlobsReturnsOnCall(i int, result1 *revokpb.BoshBlobsResponse, result2 error) {
	fake.BoshBlobsStub = nil
	if fake.boshBlobsReturnsOnCall == nil {
		fake.boshBlobsReturnsOnCall = make(map[int]struct {
			result1 *revokpb.BoshBlobsResponse
			result2 error
		})
	}
	fake.boshBlobsReturnsOnCall[i] = struct {
		result1 *revokpb.BoshBlobsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCredentialCountsMutex.RLock()
	defer fake.getCredentialCountsMutex.RUnlock()
	fake.getOrganizationCredentialCountsMutex.RLock()
	defer fake.getOrganizationCredentialCountsMutex.RUnlock()
	fake.getRepositoryCredentialCountsMutex.RLock()
	defer fake.getRepositoryCredentialCountsMutex.RUnlock()
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	fake.boshBlobsMutex.RLock()
	defer fake.boshBlobsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServer) recordInvocation(key string, args []interface{}) {
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

var _ revok.Server = new(FakeServer)
