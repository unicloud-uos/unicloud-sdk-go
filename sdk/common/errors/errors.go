package errors

import (
	"fmt"
)

type UnicloudCloudSDKError struct {
	Code      string
	RequestId string
	Msg string
}

func (e *UnicloudCloudSDKError) Error() string {
	return fmt.Sprintf("[UnicloudCloudSDKError] Code=%s, Message=%s, RequestId=%s", e.Code, e.Msg, e.RequestId)
}

func NewUnicloudCloudSDKError(code, message, requestId string) error {
	return &UnicloudCloudSDKError{
		Code:      code,
		Msg:   message,
		RequestId: requestId,
	}
}

func (e *UnicloudCloudSDKError) GetCode() string {
	return e.Code
}

func (e *UnicloudCloudSDKError) GetMessage() string {
	return e.Msg
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