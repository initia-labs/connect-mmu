// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	coinmarketcap "github.com/skip-mev/connect-mmu/market-indexer/coinmarketcap"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// CryptoIDMap provides a mock function with given fields: ctx
func (_m *Client) CryptoIDMap(ctx context.Context) (coinmarketcap.CryptoIDMapResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CryptoIDMap")
	}

	var r0 coinmarketcap.CryptoIDMapResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (coinmarketcap.CryptoIDMapResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) coinmarketcap.CryptoIDMapResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(coinmarketcap.CryptoIDMapResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_CryptoIDMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CryptoIDMap'
type Client_CryptoIDMap_Call struct {
	*mock.Call
}

// CryptoIDMap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Client_Expecter) CryptoIDMap(ctx interface{}) *Client_CryptoIDMap_Call {
	return &Client_CryptoIDMap_Call{Call: _e.mock.On("CryptoIDMap", ctx)}
}

func (_c *Client_CryptoIDMap_Call) Run(run func(ctx context.Context)) *Client_CryptoIDMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_CryptoIDMap_Call) Return(_a0 coinmarketcap.CryptoIDMapResponse, _a1 error) *Client_CryptoIDMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_CryptoIDMap_Call) RunAndReturn(run func(context.Context) (coinmarketcap.CryptoIDMapResponse, error)) *Client_CryptoIDMap_Call {
	_c.Call.Return(run)
	return _c
}

// ExchangeAssets provides a mock function with given fields: ctx, exchange
func (_m *Client) ExchangeAssets(ctx context.Context, exchange int) (coinmarketcap.ExchangeAssetsResponse, error) {
	ret := _m.Called(ctx, exchange)

	if len(ret) == 0 {
		panic("no return value specified for ExchangeAssets")
	}

	var r0 coinmarketcap.ExchangeAssetsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (coinmarketcap.ExchangeAssetsResponse, error)); ok {
		return rf(ctx, exchange)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) coinmarketcap.ExchangeAssetsResponse); ok {
		r0 = rf(ctx, exchange)
	} else {
		r0 = ret.Get(0).(coinmarketcap.ExchangeAssetsResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, exchange)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_ExchangeAssets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExchangeAssets'
type Client_ExchangeAssets_Call struct {
	*mock.Call
}

// ExchangeAssets is a helper method to define mock.On call
//   - ctx context.Context
//   - exchange int
func (_e *Client_Expecter) ExchangeAssets(ctx interface{}, exchange interface{}) *Client_ExchangeAssets_Call {
	return &Client_ExchangeAssets_Call{Call: _e.mock.On("ExchangeAssets", ctx, exchange)}
}

func (_c *Client_ExchangeAssets_Call) Run(run func(ctx context.Context, exchange int)) *Client_ExchangeAssets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *Client_ExchangeAssets_Call) Return(_a0 coinmarketcap.ExchangeAssetsResponse, _a1 error) *Client_ExchangeAssets_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_ExchangeAssets_Call) RunAndReturn(run func(context.Context, int) (coinmarketcap.ExchangeAssetsResponse, error)) *Client_ExchangeAssets_Call {
	_c.Call.Return(run)
	return _c
}

// ExchangeIDMap provides a mock function with given fields: ctx
func (_m *Client) ExchangeIDMap(ctx context.Context) (coinmarketcap.ExchangeIDMapResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ExchangeIDMap")
	}

	var r0 coinmarketcap.ExchangeIDMapResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (coinmarketcap.ExchangeIDMapResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) coinmarketcap.ExchangeIDMapResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(coinmarketcap.ExchangeIDMapResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_ExchangeIDMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExchangeIDMap'
type Client_ExchangeIDMap_Call struct {
	*mock.Call
}

// ExchangeIDMap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Client_Expecter) ExchangeIDMap(ctx interface{}) *Client_ExchangeIDMap_Call {
	return &Client_ExchangeIDMap_Call{Call: _e.mock.On("ExchangeIDMap", ctx)}
}

func (_c *Client_ExchangeIDMap_Call) Run(run func(ctx context.Context)) *Client_ExchangeIDMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_ExchangeIDMap_Call) Return(_a0 coinmarketcap.ExchangeIDMapResponse, _a1 error) *Client_ExchangeIDMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_ExchangeIDMap_Call) RunAndReturn(run func(context.Context) (coinmarketcap.ExchangeIDMapResponse, error)) *Client_ExchangeIDMap_Call {
	_c.Call.Return(run)
	return _c
}

// ExchangeMarkets provides a mock function with given fields: ctx, exchange
func (_m *Client) ExchangeMarkets(ctx context.Context, exchange int) (coinmarketcap.ExchangeMarketsResponse, error) {
	ret := _m.Called(ctx, exchange)

	if len(ret) == 0 {
		panic("no return value specified for ExchangeMarkets")
	}

	var r0 coinmarketcap.ExchangeMarketsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (coinmarketcap.ExchangeMarketsResponse, error)); ok {
		return rf(ctx, exchange)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) coinmarketcap.ExchangeMarketsResponse); ok {
		r0 = rf(ctx, exchange)
	} else {
		r0 = ret.Get(0).(coinmarketcap.ExchangeMarketsResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, exchange)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_ExchangeMarkets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExchangeMarkets'
type Client_ExchangeMarkets_Call struct {
	*mock.Call
}

// ExchangeMarkets is a helper method to define mock.On call
//   - ctx context.Context
//   - exchange int
func (_e *Client_Expecter) ExchangeMarkets(ctx interface{}, exchange interface{}) *Client_ExchangeMarkets_Call {
	return &Client_ExchangeMarkets_Call{Call: _e.mock.On("ExchangeMarkets", ctx, exchange)}
}

func (_c *Client_ExchangeMarkets_Call) Run(run func(ctx context.Context, exchange int)) *Client_ExchangeMarkets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *Client_ExchangeMarkets_Call) Return(_a0 coinmarketcap.ExchangeMarketsResponse, _a1 error) *Client_ExchangeMarkets_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_ExchangeMarkets_Call) RunAndReturn(run func(context.Context, int) (coinmarketcap.ExchangeMarketsResponse, error)) *Client_ExchangeMarkets_Call {
	_c.Call.Return(run)
	return _c
}

// FiatMap provides a mock function with given fields: ctx
func (_m *Client) FiatMap(ctx context.Context) (coinmarketcap.FiatResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FiatMap")
	}

	var r0 coinmarketcap.FiatResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (coinmarketcap.FiatResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) coinmarketcap.FiatResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(coinmarketcap.FiatResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_FiatMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FiatMap'
type Client_FiatMap_Call struct {
	*mock.Call
}

// FiatMap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Client_Expecter) FiatMap(ctx interface{}) *Client_FiatMap_Call {
	return &Client_FiatMap_Call{Call: _e.mock.On("FiatMap", ctx)}
}

func (_c *Client_FiatMap_Call) Run(run func(ctx context.Context)) *Client_FiatMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_FiatMap_Call) Return(_a0 coinmarketcap.FiatResponse, _a1 error) *Client_FiatMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_FiatMap_Call) RunAndReturn(run func(context.Context) (coinmarketcap.FiatResponse, error)) *Client_FiatMap_Call {
	_c.Call.Return(run)
	return _c
}

// Info provides a mock function with given fields: ctx, ids
func (_m *Client) Info(ctx context.Context, ids []int64) (coinmarketcap.InfoResponse, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for Info")
	}

	var r0 coinmarketcap.InfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64) (coinmarketcap.InfoResponse, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []int64) coinmarketcap.InfoResponse); ok {
		r0 = rf(ctx, ids)
	} else {
		r0 = ret.Get(0).(coinmarketcap.InfoResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []int64) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type Client_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []int64
func (_e *Client_Expecter) Info(ctx interface{}, ids interface{}) *Client_Info_Call {
	return &Client_Info_Call{Call: _e.mock.On("Info", ctx, ids)}
}

func (_c *Client_Info_Call) Run(run func(ctx context.Context, ids []int64)) *Client_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]int64))
	})
	return _c
}

func (_c *Client_Info_Call) Return(_a0 coinmarketcap.InfoResponse, _a1 error) *Client_Info_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_Info_Call) RunAndReturn(run func(context.Context, []int64) (coinmarketcap.InfoResponse, error)) *Client_Info_Call {
	_c.Call.Return(run)
	return _c
}

// Quote provides a mock function with given fields: ctx, id
func (_m *Client) Quote(ctx context.Context, id int64) (coinmarketcap.QuoteResponse, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Quote")
	}

	var r0 coinmarketcap.QuoteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (coinmarketcap.QuoteResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) coinmarketcap.QuoteResponse); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(coinmarketcap.QuoteResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Quote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Quote'
type Client_Quote_Call struct {
	*mock.Call
}

// Quote is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *Client_Expecter) Quote(ctx interface{}, id interface{}) *Client_Quote_Call {
	return &Client_Quote_Call{Call: _e.mock.On("Quote", ctx, id)}
}

func (_c *Client_Quote_Call) Run(run func(ctx context.Context, id int64)) *Client_Quote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *Client_Quote_Call) Return(_a0 coinmarketcap.QuoteResponse, _a1 error) *Client_Quote_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_Quote_Call) RunAndReturn(run func(context.Context, int64) (coinmarketcap.QuoteResponse, error)) *Client_Quote_Call {
	_c.Call.Return(run)
	return _c
}

// Quotes provides a mock function with given fields: ctx, ids
func (_m *Client) Quotes(ctx context.Context, ids []int64) (coinmarketcap.QuoteResponse, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for Quotes")
	}

	var r0 coinmarketcap.QuoteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64) (coinmarketcap.QuoteResponse, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []int64) coinmarketcap.QuoteResponse); ok {
		r0 = rf(ctx, ids)
	} else {
		r0 = ret.Get(0).(coinmarketcap.QuoteResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []int64) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Quotes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Quotes'
type Client_Quotes_Call struct {
	*mock.Call
}

// Quotes is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []int64
func (_e *Client_Expecter) Quotes(ctx interface{}, ids interface{}) *Client_Quotes_Call {
	return &Client_Quotes_Call{Call: _e.mock.On("Quotes", ctx, ids)}
}

func (_c *Client_Quotes_Call) Run(run func(ctx context.Context, ids []int64)) *Client_Quotes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]int64))
	})
	return _c
}

func (_c *Client_Quotes_Call) Return(_a0 coinmarketcap.QuoteResponse, _a1 error) *Client_Quotes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_Quotes_Call) RunAndReturn(run func(context.Context, []int64) (coinmarketcap.QuoteResponse, error)) *Client_Quotes_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
