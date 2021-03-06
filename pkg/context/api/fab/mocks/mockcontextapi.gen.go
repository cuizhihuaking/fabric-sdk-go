// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/fabric-sdk-go/pkg/context/api/fab (interfaces: ProposalProcessor)

// Package mock_fab is a generated GoMock package.
package mock_fab

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	fab "github.com/hyperledger/fabric-sdk-go/pkg/context/api/fab"
)

// MockProposalProcessor is a mock of ProposalProcessor interface
type MockProposalProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProposalProcessorMockRecorder
}

// MockProposalProcessorMockRecorder is the mock recorder for MockProposalProcessor
type MockProposalProcessorMockRecorder struct {
	mock *MockProposalProcessor
}

// NewMockProposalProcessor creates a new mock instance
func NewMockProposalProcessor(ctrl *gomock.Controller) *MockProposalProcessor {
	mock := &MockProposalProcessor{ctrl: ctrl}
	mock.recorder = &MockProposalProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProposalProcessor) EXPECT() *MockProposalProcessorMockRecorder {
	return m.recorder
}

// ProcessTransactionProposal mocks base method
func (m *MockProposalProcessor) ProcessTransactionProposal(arg0 fab.ProcessProposalRequest) (*fab.TransactionProposalResponse, error) {
	ret := m.ctrl.Call(m, "ProcessTransactionProposal", arg0)
	ret0, _ := ret[0].(*fab.TransactionProposalResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessTransactionProposal indicates an expected call of ProcessTransactionProposal
func (mr *MockProposalProcessorMockRecorder) ProcessTransactionProposal(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTransactionProposal", reflect.TypeOf((*MockProposalProcessor)(nil).ProcessTransactionProposal), arg0)
}
