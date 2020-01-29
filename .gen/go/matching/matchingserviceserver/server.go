// The MIT License (MIT)
// 
// Copyright (c) 2020 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by thriftrw-plugin-yarpc
// @generated

package matchingserviceserver

import (
	context "context"
	matching "github.com/temporalio/temporal/.gen/go/matching"
	shared "github.com/temporalio/temporal/.gen/go/shared"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
)

// Interface is the server-side interface for the MatchingService service.
type Interface interface {
	AddActivityTask(
		ctx context.Context,
		AddRequest *matching.AddActivityTaskRequest,
	) error

	AddDecisionTask(
		ctx context.Context,
		AddRequest *matching.AddDecisionTaskRequest,
	) error

	CancelOutstandingPoll(
		ctx context.Context,
		Request *matching.CancelOutstandingPollRequest,
	) error

	DescribeTaskList(
		ctx context.Context,
		Request *matching.DescribeTaskListRequest,
	) (*shared.DescribeTaskListResponse, error)

	ListTaskListPartitions(
		ctx context.Context,
		Request *matching.ListTaskListPartitionsRequest,
	) (*shared.ListTaskListPartitionsResponse, error)

	PollForActivityTask(
		ctx context.Context,
		PollRequest *matching.PollForActivityTaskRequest,
	) (*shared.PollForActivityTaskResponse, error)

	PollForDecisionTask(
		ctx context.Context,
		PollRequest *matching.PollForDecisionTaskRequest,
	) (*matching.PollForDecisionTaskResponse, error)

	QueryWorkflow(
		ctx context.Context,
		QueryRequest *matching.QueryWorkflowRequest,
	) (*shared.QueryWorkflowResponse, error)

	RespondQueryTaskCompleted(
		ctx context.Context,
		Request *matching.RespondQueryTaskCompletedRequest,
	) error
}

// New prepares an implementation of the MatchingService service for
// registration.
//
// 	handler := MatchingServiceHandler{}
// 	dispatcher.Register(matchingserviceserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "MatchingService",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "AddActivityTask",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.AddActivityTask),
				},
				Signature:    "AddActivityTask(AddRequest *matching.AddActivityTaskRequest)",
				ThriftModule: matching.ThriftModule,
			},

			thrift.Method{
				Name: "AddDecisionTask",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.AddDecisionTask),
				},
				Signature:    "AddDecisionTask(AddRequest *matching.AddDecisionTaskRequest)",
				ThriftModule: matching.ThriftModule,
			},

			thrift.Method{
				Name: "CancelOutstandingPoll",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.CancelOutstandingPoll),
				},
				Signature:    "CancelOutstandingPoll(Request *matching.CancelOutstandingPollRequest)",
				ThriftModule: matching.ThriftModule,
			},

			thrift.Method{
				Name: "DescribeTaskList",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.DescribeTaskList),
				},
				Signature:    "DescribeTaskList(Request *matching.DescribeTaskListRequest) (*shared.DescribeTaskListResponse)",
				ThriftModule: matching.ThriftModule,
			},

			thrift.Method{
				Name: "ListTaskListPartitions",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ListTaskListPartitions),
				},
				Signature:    "ListTaskListPartitions(Request *matching.ListTaskListPartitionsRequest) (*shared.ListTaskListPartitionsResponse)",
				ThriftModule: matching.ThriftModule,
			},

			thrift.Method{
				Name: "PollForActivityTask",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.PollForActivityTask),
				},
				Signature:    "PollForActivityTask(PollRequest *matching.PollForActivityTaskRequest) (*shared.PollForActivityTaskResponse)",
				ThriftModule: matching.ThriftModule,
			},

			thrift.Method{
				Name: "PollForDecisionTask",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.PollForDecisionTask),
				},
				Signature:    "PollForDecisionTask(PollRequest *matching.PollForDecisionTaskRequest) (*matching.PollForDecisionTaskResponse)",
				ThriftModule: matching.ThriftModule,
			},

			thrift.Method{
				Name: "QueryWorkflow",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.QueryWorkflow),
				},
				Signature:    "QueryWorkflow(QueryRequest *matching.QueryWorkflowRequest) (*shared.QueryWorkflowResponse)",
				ThriftModule: matching.ThriftModule,
			},

			thrift.Method{
				Name: "RespondQueryTaskCompleted",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RespondQueryTaskCompleted),
				},
				Signature:    "RespondQueryTaskCompleted(Request *matching.RespondQueryTaskCompletedRequest)",
				ThriftModule: matching.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 9)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

func (h handler) AddActivityTask(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args matching.MatchingService_AddActivityTask_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.AddActivityTask(ctx, args.AddRequest)

	hadError := err != nil
	result, err := matching.MatchingService_AddActivityTask_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) AddDecisionTask(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args matching.MatchingService_AddDecisionTask_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.AddDecisionTask(ctx, args.AddRequest)

	hadError := err != nil
	result, err := matching.MatchingService_AddDecisionTask_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) CancelOutstandingPoll(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args matching.MatchingService_CancelOutstandingPoll_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.CancelOutstandingPoll(ctx, args.Request)

	hadError := err != nil
	result, err := matching.MatchingService_CancelOutstandingPoll_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) DescribeTaskList(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args matching.MatchingService_DescribeTaskList_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.DescribeTaskList(ctx, args.Request)

	hadError := err != nil
	result, err := matching.MatchingService_DescribeTaskList_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ListTaskListPartitions(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args matching.MatchingService_ListTaskListPartitions_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.ListTaskListPartitions(ctx, args.Request)

	hadError := err != nil
	result, err := matching.MatchingService_ListTaskListPartitions_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) PollForActivityTask(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args matching.MatchingService_PollForActivityTask_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.PollForActivityTask(ctx, args.PollRequest)

	hadError := err != nil
	result, err := matching.MatchingService_PollForActivityTask_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) PollForDecisionTask(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args matching.MatchingService_PollForDecisionTask_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.PollForDecisionTask(ctx, args.PollRequest)

	hadError := err != nil
	result, err := matching.MatchingService_PollForDecisionTask_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) QueryWorkflow(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args matching.MatchingService_QueryWorkflow_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.QueryWorkflow(ctx, args.QueryRequest)

	hadError := err != nil
	result, err := matching.MatchingService_QueryWorkflow_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RespondQueryTaskCompleted(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args matching.MatchingService_RespondQueryTaskCompleted_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.RespondQueryTaskCompleted(ctx, args.Request)

	hadError := err != nil
	result, err := matching.MatchingService_RespondQueryTaskCompleted_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}
