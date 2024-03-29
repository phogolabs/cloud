// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/phogolabs/cloud"
)

type FakeEventArgsConverter struct {
	DataAsStub        func(interface{}) error
	dataAsMutex       sync.RWMutex
	dataAsArgsForCall []struct {
		arg1 interface{}
	}
	dataAsReturns struct {
		result1 error
	}
	dataAsReturnsOnCall map[int]struct {
		result1 error
	}
	DataContentTypeStub        func() string
	dataContentTypeMutex       sync.RWMutex
	dataContentTypeArgsForCall []struct {
	}
	dataContentTypeReturns struct {
		result1 string
	}
	dataContentTypeReturnsOnCall map[int]struct {
		result1 string
	}
	DataSchemaStub        func() string
	dataSchemaMutex       sync.RWMutex
	dataSchemaArgsForCall []struct {
	}
	dataSchemaReturns struct {
		result1 string
	}
	dataSchemaReturnsOnCall map[int]struct {
		result1 string
	}
	ExtensionsStub        func() map[string]interface{}
	extensionsMutex       sync.RWMutex
	extensionsArgsForCall []struct {
	}
	extensionsReturns struct {
		result1 map[string]interface{}
	}
	extensionsReturnsOnCall map[int]struct {
		result1 map[string]interface{}
	}
	SourceStub        func() string
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct {
	}
	sourceReturns struct {
		result1 string
	}
	sourceReturnsOnCall map[int]struct {
		result1 string
	}
	SubjectStub        func() string
	subjectMutex       sync.RWMutex
	subjectArgsForCall []struct {
	}
	subjectReturns struct {
		result1 string
	}
	subjectReturnsOnCall map[int]struct {
		result1 string
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct {
	}
	typeReturns struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventArgsConverter) DataAs(arg1 interface{}) error {
	fake.dataAsMutex.Lock()
	ret, specificReturn := fake.dataAsReturnsOnCall[len(fake.dataAsArgsForCall)]
	fake.dataAsArgsForCall = append(fake.dataAsArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("DataAs", []interface{}{arg1})
	fake.dataAsMutex.Unlock()
	if fake.DataAsStub != nil {
		return fake.DataAsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.dataAsReturns
	return fakeReturns.result1
}

func (fake *FakeEventArgsConverter) DataAsCallCount() int {
	fake.dataAsMutex.RLock()
	defer fake.dataAsMutex.RUnlock()
	return len(fake.dataAsArgsForCall)
}

func (fake *FakeEventArgsConverter) DataAsCalls(stub func(interface{}) error) {
	fake.dataAsMutex.Lock()
	defer fake.dataAsMutex.Unlock()
	fake.DataAsStub = stub
}

func (fake *FakeEventArgsConverter) DataAsArgsForCall(i int) interface{} {
	fake.dataAsMutex.RLock()
	defer fake.dataAsMutex.RUnlock()
	argsForCall := fake.dataAsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEventArgsConverter) DataAsReturns(result1 error) {
	fake.dataAsMutex.Lock()
	defer fake.dataAsMutex.Unlock()
	fake.DataAsStub = nil
	fake.dataAsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventArgsConverter) DataAsReturnsOnCall(i int, result1 error) {
	fake.dataAsMutex.Lock()
	defer fake.dataAsMutex.Unlock()
	fake.DataAsStub = nil
	if fake.dataAsReturnsOnCall == nil {
		fake.dataAsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.dataAsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventArgsConverter) DataContentType() string {
	fake.dataContentTypeMutex.Lock()
	ret, specificReturn := fake.dataContentTypeReturnsOnCall[len(fake.dataContentTypeArgsForCall)]
	fake.dataContentTypeArgsForCall = append(fake.dataContentTypeArgsForCall, struct {
	}{})
	fake.recordInvocation("DataContentType", []interface{}{})
	fake.dataContentTypeMutex.Unlock()
	if fake.DataContentTypeStub != nil {
		return fake.DataContentTypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.dataContentTypeReturns
	return fakeReturns.result1
}

func (fake *FakeEventArgsConverter) DataContentTypeCallCount() int {
	fake.dataContentTypeMutex.RLock()
	defer fake.dataContentTypeMutex.RUnlock()
	return len(fake.dataContentTypeArgsForCall)
}

func (fake *FakeEventArgsConverter) DataContentTypeCalls(stub func() string) {
	fake.dataContentTypeMutex.Lock()
	defer fake.dataContentTypeMutex.Unlock()
	fake.DataContentTypeStub = stub
}

func (fake *FakeEventArgsConverter) DataContentTypeReturns(result1 string) {
	fake.dataContentTypeMutex.Lock()
	defer fake.dataContentTypeMutex.Unlock()
	fake.DataContentTypeStub = nil
	fake.dataContentTypeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventArgsConverter) DataContentTypeReturnsOnCall(i int, result1 string) {
	fake.dataContentTypeMutex.Lock()
	defer fake.dataContentTypeMutex.Unlock()
	fake.DataContentTypeStub = nil
	if fake.dataContentTypeReturnsOnCall == nil {
		fake.dataContentTypeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.dataContentTypeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventArgsConverter) DataSchema() string {
	fake.dataSchemaMutex.Lock()
	ret, specificReturn := fake.dataSchemaReturnsOnCall[len(fake.dataSchemaArgsForCall)]
	fake.dataSchemaArgsForCall = append(fake.dataSchemaArgsForCall, struct {
	}{})
	fake.recordInvocation("DataSchema", []interface{}{})
	fake.dataSchemaMutex.Unlock()
	if fake.DataSchemaStub != nil {
		return fake.DataSchemaStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.dataSchemaReturns
	return fakeReturns.result1
}

func (fake *FakeEventArgsConverter) DataSchemaCallCount() int {
	fake.dataSchemaMutex.RLock()
	defer fake.dataSchemaMutex.RUnlock()
	return len(fake.dataSchemaArgsForCall)
}

func (fake *FakeEventArgsConverter) DataSchemaCalls(stub func() string) {
	fake.dataSchemaMutex.Lock()
	defer fake.dataSchemaMutex.Unlock()
	fake.DataSchemaStub = stub
}

func (fake *FakeEventArgsConverter) DataSchemaReturns(result1 string) {
	fake.dataSchemaMutex.Lock()
	defer fake.dataSchemaMutex.Unlock()
	fake.DataSchemaStub = nil
	fake.dataSchemaReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventArgsConverter) DataSchemaReturnsOnCall(i int, result1 string) {
	fake.dataSchemaMutex.Lock()
	defer fake.dataSchemaMutex.Unlock()
	fake.DataSchemaStub = nil
	if fake.dataSchemaReturnsOnCall == nil {
		fake.dataSchemaReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.dataSchemaReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventArgsConverter) Extensions() map[string]interface{} {
	fake.extensionsMutex.Lock()
	ret, specificReturn := fake.extensionsReturnsOnCall[len(fake.extensionsArgsForCall)]
	fake.extensionsArgsForCall = append(fake.extensionsArgsForCall, struct {
	}{})
	fake.recordInvocation("Extensions", []interface{}{})
	fake.extensionsMutex.Unlock()
	if fake.ExtensionsStub != nil {
		return fake.ExtensionsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.extensionsReturns
	return fakeReturns.result1
}

func (fake *FakeEventArgsConverter) ExtensionsCallCount() int {
	fake.extensionsMutex.RLock()
	defer fake.extensionsMutex.RUnlock()
	return len(fake.extensionsArgsForCall)
}

func (fake *FakeEventArgsConverter) ExtensionsCalls(stub func() map[string]interface{}) {
	fake.extensionsMutex.Lock()
	defer fake.extensionsMutex.Unlock()
	fake.ExtensionsStub = stub
}

func (fake *FakeEventArgsConverter) ExtensionsReturns(result1 map[string]interface{}) {
	fake.extensionsMutex.Lock()
	defer fake.extensionsMutex.Unlock()
	fake.ExtensionsStub = nil
	fake.extensionsReturns = struct {
		result1 map[string]interface{}
	}{result1}
}

func (fake *FakeEventArgsConverter) ExtensionsReturnsOnCall(i int, result1 map[string]interface{}) {
	fake.extensionsMutex.Lock()
	defer fake.extensionsMutex.Unlock()
	fake.ExtensionsStub = nil
	if fake.extensionsReturnsOnCall == nil {
		fake.extensionsReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
		})
	}
	fake.extensionsReturnsOnCall[i] = struct {
		result1 map[string]interface{}
	}{result1}
}

func (fake *FakeEventArgsConverter) Source() string {
	fake.sourceMutex.Lock()
	ret, specificReturn := fake.sourceReturnsOnCall[len(fake.sourceArgsForCall)]
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct {
	}{})
	fake.recordInvocation("Source", []interface{}{})
	fake.sourceMutex.Unlock()
	if fake.SourceStub != nil {
		return fake.SourceStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sourceReturns
	return fakeReturns.result1
}

func (fake *FakeEventArgsConverter) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeEventArgsConverter) SourceCalls(stub func() string) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = stub
}

func (fake *FakeEventArgsConverter) SourceReturns(result1 string) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventArgsConverter) SourceReturnsOnCall(i int, result1 string) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	if fake.sourceReturnsOnCall == nil {
		fake.sourceReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.sourceReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventArgsConverter) Subject() string {
	fake.subjectMutex.Lock()
	ret, specificReturn := fake.subjectReturnsOnCall[len(fake.subjectArgsForCall)]
	fake.subjectArgsForCall = append(fake.subjectArgsForCall, struct {
	}{})
	fake.recordInvocation("Subject", []interface{}{})
	fake.subjectMutex.Unlock()
	if fake.SubjectStub != nil {
		return fake.SubjectStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.subjectReturns
	return fakeReturns.result1
}

func (fake *FakeEventArgsConverter) SubjectCallCount() int {
	fake.subjectMutex.RLock()
	defer fake.subjectMutex.RUnlock()
	return len(fake.subjectArgsForCall)
}

func (fake *FakeEventArgsConverter) SubjectCalls(stub func() string) {
	fake.subjectMutex.Lock()
	defer fake.subjectMutex.Unlock()
	fake.SubjectStub = stub
}

func (fake *FakeEventArgsConverter) SubjectReturns(result1 string) {
	fake.subjectMutex.Lock()
	defer fake.subjectMutex.Unlock()
	fake.SubjectStub = nil
	fake.subjectReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventArgsConverter) SubjectReturnsOnCall(i int, result1 string) {
	fake.subjectMutex.Lock()
	defer fake.subjectMutex.Unlock()
	fake.SubjectStub = nil
	if fake.subjectReturnsOnCall == nil {
		fake.subjectReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.subjectReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventArgsConverter) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct {
	}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.typeReturns
	return fakeReturns.result1
}

func (fake *FakeEventArgsConverter) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeEventArgsConverter) TypeCalls(stub func() string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = stub
}

func (fake *FakeEventArgsConverter) TypeReturns(result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventArgsConverter) TypeReturnsOnCall(i int, result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventArgsConverter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dataAsMutex.RLock()
	defer fake.dataAsMutex.RUnlock()
	fake.dataContentTypeMutex.RLock()
	defer fake.dataContentTypeMutex.RUnlock()
	fake.dataSchemaMutex.RLock()
	defer fake.dataSchemaMutex.RUnlock()
	fake.extensionsMutex.RLock()
	defer fake.extensionsMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	fake.subjectMutex.RLock()
	defer fake.subjectMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventArgsConverter) recordInvocation(key string, args []interface{}) {
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

var _ cloud.EventArgsConverter = new(FakeEventArgsConverter)
