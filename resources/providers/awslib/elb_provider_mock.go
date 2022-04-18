// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
// Code generated by mockery v2.10.4. DO NOT EDIT.

package awslib

import (
	context "context"

	elasticloadbalancing "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	mock "github.com/stretchr/testify/mock"
)

// MockedELBLoadBalancerDescriber is an autogenerated mock type for the ELBLoadBalancerDescriber type
type MockedELBLoadBalancerDescriber struct {
	mock.Mock
}

type MockELBLoadBalancerDescriber_Expecter struct {
	mock *mock.Mock
}

func (_m *MockedELBLoadBalancerDescriber) EXPECT() *MockELBLoadBalancerDescriber_Expecter {
	return &MockELBLoadBalancerDescriber_Expecter{mock: &_m.Mock}
}

// DescribeLoadBalancer provides a mock function with given fields: ctx, balancersNames
func (_m *MockedELBLoadBalancerDescriber) DescribeLoadBalancer(ctx context.Context, balancersNames []string) ([]elasticloadbalancing.LoadBalancerDescription, error) {
	ret := _m.Called(ctx, balancersNames)

	var r0 []elasticloadbalancing.LoadBalancerDescription
	if rf, ok := ret.Get(0).(func(context.Context, []string) []elasticloadbalancing.LoadBalancerDescription); ok {
		r0 = rf(ctx, balancersNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]elasticloadbalancing.LoadBalancerDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, balancersNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockELBLoadBalancerDescriber_DescribeLoadBalancer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeLoadBalancer'
type MockELBLoadBalancerDescriber_DescribeLoadBalancer_Call struct {
	*mock.Call
}

// DescribeLoadBalancer is a helper method to define mock.On call
//  - ctx context.Context
//  - balancersNames []string
func (_e *MockELBLoadBalancerDescriber_Expecter) DescribeLoadBalancer(ctx interface{}, balancersNames interface{}) *MockELBLoadBalancerDescriber_DescribeLoadBalancer_Call {
	return &MockELBLoadBalancerDescriber_DescribeLoadBalancer_Call{Call: _e.mock.On("DescribeLoadBalancer", ctx, balancersNames)}
}

func (_c *MockELBLoadBalancerDescriber_DescribeLoadBalancer_Call) Run(run func(ctx context.Context, balancersNames []string)) *MockELBLoadBalancerDescriber_DescribeLoadBalancer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockELBLoadBalancerDescriber_DescribeLoadBalancer_Call) Return(_a0 []elasticloadbalancing.LoadBalancerDescription, _a1 error) *MockELBLoadBalancerDescriber_DescribeLoadBalancer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}
