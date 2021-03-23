package errors

import (
	"fmt"
)

type UnicloudCloudSDKError struct {
	Code      string
	Message   string
	RequestId string
}

func (e *UnicloudCloudSDKError) Error() string {
	return fmt.Sprintf("[UnicloudCloudSDKError] Code=%s, Message=%s, RequestId=%s", e.Code, e.Message, e.RequestId)
}

func NewUnicloudCloudSDKError(code, message, requestId string) error {
	return &UnicloudCloudSDKError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

func (e *UnicloudCloudSDKError) GetCode() string {
	return e.Code
}

func (e *UnicloudCloudSDKError) GetMessage() string {
	return e.Message
}

func (e *UnicloudCloudSDKError) GetRequestId() string {
	return e.RequestId
}

type ClientError struct {
	Message string
}

func MakeClientError(err error) ClientError {
	return ClientError{
		Message: err.Error(),
	}
}