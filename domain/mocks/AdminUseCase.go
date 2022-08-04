// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "PROJECT-III/domain"

	mock "github.com/stretchr/testify/mock"
)

// AdminUseCase is an autogenerated mock type for the AdminUseCase type
type AdminUseCase struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: newProduct, adminid
func (_m *AdminUseCase) CreateProduct(newProduct domain.Product, adminid int) int {
	ret := _m.Called(newProduct, adminid)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.Product, int) int); ok {
		r0 = rf(newProduct, adminid)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// DeleteProduct provides a mock function with given fields: productid, adminid
func (_m *AdminUseCase) DeleteProduct(productid int, adminid int) int {
	ret := _m.Called(productid, adminid)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(productid, adminid)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ReadAllProduct provides a mock function with given fields:
func (_m *AdminUseCase) ReadAllProduct() ([]domain.Product, int) {
	ret := _m.Called()

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func() []domain.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: updatedData, productid, adminid
func (_m *AdminUseCase) UpdateProduct(updatedData domain.Product, productid int, adminid int) int {
	ret := _m.Called(updatedData, productid, adminid)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.Product, int, int) int); ok {
		r0 = rf(updatedData, productid, adminid)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewAdminUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAdminUseCase creates a new instance of AdminUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAdminUseCase(t mockConstructorTestingTNewAdminUseCase) *AdminUseCase {
	mock := &AdminUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
