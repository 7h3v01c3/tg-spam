// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"sync"
)

// OpenAIClientMock is a mock implementation of lib.openAIClient.
//
//	func TestSomethingThatUsesopenAIClient(t *testing.T) {
//
//		// make and configure a mocked lib.openAIClient
//		mockedopenAIClient := &OpenAIClientMock{
//			CreateChatCompletionFunc: func(contextMoqParam context.Context, chatCompletionRequest openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
//				panic("mock out the CreateChatCompletion method")
//			},
//		}
//
//		// use mockedopenAIClient in code that requires lib.openAIClient
//		// and then make assertions.
//
//	}
type OpenAIClientMock struct {
	// CreateChatCompletionFunc mocks the CreateChatCompletion method.
	CreateChatCompletionFunc func(contextMoqParam context.Context, chatCompletionRequest openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateChatCompletion holds details about calls to the CreateChatCompletion method.
		CreateChatCompletion []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// ChatCompletionRequest is the chatCompletionRequest argument value.
			ChatCompletionRequest openai.ChatCompletionRequest
		}
	}
	lockCreateChatCompletion sync.RWMutex
}

// CreateChatCompletion calls CreateChatCompletionFunc.
func (mock *OpenAIClientMock) CreateChatCompletion(contextMoqParam context.Context, chatCompletionRequest openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
	if mock.CreateChatCompletionFunc == nil {
		panic("OpenAIClientMock.CreateChatCompletionFunc: method is nil but openAIClient.CreateChatCompletion was just called")
	}
	callInfo := struct {
		ContextMoqParam       context.Context
		ChatCompletionRequest openai.ChatCompletionRequest
	}{
		ContextMoqParam:       contextMoqParam,
		ChatCompletionRequest: chatCompletionRequest,
	}
	mock.lockCreateChatCompletion.Lock()
	mock.calls.CreateChatCompletion = append(mock.calls.CreateChatCompletion, callInfo)
	mock.lockCreateChatCompletion.Unlock()
	return mock.CreateChatCompletionFunc(contextMoqParam, chatCompletionRequest)
}

// CreateChatCompletionCalls gets all the calls that were made to CreateChatCompletion.
// check the length with:
//
//	len(mockedopenAIClient.CreateChatCompletionCalls())
func (mock *OpenAIClientMock) CreateChatCompletionCalls() []struct {
	ContextMoqParam       context.Context
	ChatCompletionRequest openai.ChatCompletionRequest
} {
	var calls []struct {
		ContextMoqParam       context.Context
		ChatCompletionRequest openai.ChatCompletionRequest
	}
	mock.lockCreateChatCompletion.RLock()
	calls = mock.calls.CreateChatCompletion
	mock.lockCreateChatCompletion.RUnlock()
	return calls
}