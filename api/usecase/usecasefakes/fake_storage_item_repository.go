// Code generated by counterfeiter. DO NOT EDIT.
package usecasefakes

import (
	"sync"

	"github.com/martinjirku/zasobar/entity"
	"github.com/martinjirku/zasobar/usecase"
)

type FakeStorageItemRepository struct {
	ByIdStub        func(uint) (entity.StorageItem, error)
	byIdMutex       sync.RWMutex
	byIdArgsForCall []struct {
		arg1 uint
	}
	byIdReturns struct {
		result1 entity.StorageItem
		result2 error
	}
	byIdReturnsOnCall map[int]struct {
		result1 entity.StorageItem
		result2 error
	}
	CreateStub        func(entity.StorageItem) (entity.StorageItem, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 entity.StorageItem
	}
	createReturns struct {
		result1 entity.StorageItem
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 entity.StorageItem
		result2 error
	}
	ListStub        func() ([]entity.StorageItem, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
	}
	listReturns struct {
		result1 []entity.StorageItem
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []entity.StorageItem
		result2 error
	}
	UpdateStub        func(entity.StorageItem) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 entity.StorageItem
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStorageItemRepository) ById(arg1 uint) (entity.StorageItem, error) {
	fake.byIdMutex.Lock()
	ret, specificReturn := fake.byIdReturnsOnCall[len(fake.byIdArgsForCall)]
	fake.byIdArgsForCall = append(fake.byIdArgsForCall, struct {
		arg1 uint
	}{arg1})
	stub := fake.ByIdStub
	fakeReturns := fake.byIdReturns
	fake.recordInvocation("ById", []interface{}{arg1})
	fake.byIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStorageItemRepository) ByIdCallCount() int {
	fake.byIdMutex.RLock()
	defer fake.byIdMutex.RUnlock()
	return len(fake.byIdArgsForCall)
}

func (fake *FakeStorageItemRepository) ByIdCalls(stub func(uint) (entity.StorageItem, error)) {
	fake.byIdMutex.Lock()
	defer fake.byIdMutex.Unlock()
	fake.ByIdStub = stub
}

func (fake *FakeStorageItemRepository) ByIdArgsForCall(i int) uint {
	fake.byIdMutex.RLock()
	defer fake.byIdMutex.RUnlock()
	argsForCall := fake.byIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStorageItemRepository) ByIdReturns(result1 entity.StorageItem, result2 error) {
	fake.byIdMutex.Lock()
	defer fake.byIdMutex.Unlock()
	fake.ByIdStub = nil
	fake.byIdReturns = struct {
		result1 entity.StorageItem
		result2 error
	}{result1, result2}
}

func (fake *FakeStorageItemRepository) ByIdReturnsOnCall(i int, result1 entity.StorageItem, result2 error) {
	fake.byIdMutex.Lock()
	defer fake.byIdMutex.Unlock()
	fake.ByIdStub = nil
	if fake.byIdReturnsOnCall == nil {
		fake.byIdReturnsOnCall = make(map[int]struct {
			result1 entity.StorageItem
			result2 error
		})
	}
	fake.byIdReturnsOnCall[i] = struct {
		result1 entity.StorageItem
		result2 error
	}{result1, result2}
}

func (fake *FakeStorageItemRepository) Create(arg1 entity.StorageItem) (entity.StorageItem, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 entity.StorageItem
	}{arg1})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStorageItemRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeStorageItemRepository) CreateCalls(stub func(entity.StorageItem) (entity.StorageItem, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeStorageItemRepository) CreateArgsForCall(i int) entity.StorageItem {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStorageItemRepository) CreateReturns(result1 entity.StorageItem, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 entity.StorageItem
		result2 error
	}{result1, result2}
}

func (fake *FakeStorageItemRepository) CreateReturnsOnCall(i int, result1 entity.StorageItem, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 entity.StorageItem
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 entity.StorageItem
		result2 error
	}{result1, result2}
}

func (fake *FakeStorageItemRepository) List() ([]entity.StorageItem, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
	}{})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStorageItemRepository) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeStorageItemRepository) ListCalls(stub func() ([]entity.StorageItem, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeStorageItemRepository) ListReturns(result1 []entity.StorageItem, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []entity.StorageItem
		result2 error
	}{result1, result2}
}

func (fake *FakeStorageItemRepository) ListReturnsOnCall(i int, result1 []entity.StorageItem, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []entity.StorageItem
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []entity.StorageItem
		result2 error
	}{result1, result2}
}

func (fake *FakeStorageItemRepository) Update(arg1 entity.StorageItem) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 entity.StorageItem
	}{arg1})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStorageItemRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeStorageItemRepository) UpdateCalls(stub func(entity.StorageItem) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeStorageItemRepository) UpdateArgsForCall(i int) entity.StorageItem {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStorageItemRepository) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorageItemRepository) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorageItemRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.byIdMutex.RLock()
	defer fake.byIdMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStorageItemRepository) recordInvocation(key string, args []interface{}) {
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

var _ usecase.StorageItemRepository = new(FakeStorageItemRepository)
