// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"
	"thy/errors"
	"thy/requests"
)

type FakeGraphClient struct {
	DoRequestStub        func(string, interface{}, map[string]interface{}) ([]byte, *errors.ApiError)
	doRequestMutex       sync.RWMutex
	doRequestArgsForCall []struct {
		arg1 string
		arg2 interface{}
		arg3 map[string]interface{}
	}
	doRequestReturns struct {
		result1 []byte
		result2 *errors.ApiError
	}
	doRequestReturnsOnCall map[int]struct {
		result1 []byte
		result2 *errors.ApiError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGraphClient) DoRequest(arg1 string, arg2 interface{}, arg3 map[string]interface{}) ([]byte, *errors.ApiError) {
	fake.doRequestMutex.Lock()
	ret, specificReturn := fake.doRequestReturnsOnCall[len(fake.doRequestArgsForCall)]
	fake.doRequestArgsForCall = append(fake.doRequestArgsForCall, struct {
		arg1 string
		arg2 interface{}
		arg3 map[string]interface{}
	}{arg1, arg2, arg3})
	stub := fake.DoRequestStub
	fakeReturns := fake.doRequestReturns
	fake.recordInvocation("DoRequest", []interface{}{arg1, arg2, arg3})
	fake.doRequestMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGraphClient) DoRequestCallCount() int {
	fake.doRequestMutex.RLock()
	defer fake.doRequestMutex.RUnlock()
	return len(fake.doRequestArgsForCall)
}

func (fake *FakeGraphClient) DoRequestCalls(stub func(string, interface{}, map[string]interface{}) ([]byte, *errors.ApiError)) {
	fake.doRequestMutex.Lock()
	defer fake.doRequestMutex.Unlock()
	fake.DoRequestStub = stub
}

func (fake *FakeGraphClient) DoRequestArgsForCall(i int) (string, interface{}, map[string]interface{}) {
	fake.doRequestMutex.RLock()
	defer fake.doRequestMutex.RUnlock()
	argsForCall := fake.doRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGraphClient) DoRequestReturns(result1 []byte, result2 *errors.ApiError) {
	fake.doRequestMutex.Lock()
	defer fake.doRequestMutex.Unlock()
	fake.DoRequestStub = nil
	fake.doRequestReturns = struct {
		result1 []byte
		result2 *errors.ApiError
	}{result1, result2}
}

func (fake *FakeGraphClient) DoRequestReturnsOnCall(i int, result1 []byte, result2 *errors.ApiError) {
	fake.doRequestMutex.Lock()
	defer fake.doRequestMutex.Unlock()
	fake.DoRequestStub = nil
	if fake.doRequestReturnsOnCall == nil {
		fake.doRequestReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 *errors.ApiError
		})
	}
	fake.doRequestReturnsOnCall[i] = struct {
		result1 []byte
		result2 *errors.ApiError
	}{result1, result2}
}

func (fake *FakeGraphClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doRequestMutex.RLock()
	defer fake.doRequestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGraphClient) recordInvocation(key string, args []interface{}) {
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

var _ requests.GraphClient = new(FakeGraphClient)