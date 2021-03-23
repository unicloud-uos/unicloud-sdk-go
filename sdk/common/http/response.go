package common

import (
	"encoding/json"
	"fmt"
	"github.com/unicloud-uos/unicloud-sdk-go/sdk/common/errors"
	"io/ioutil"
	//"log"
	"net/http"

)

type Response interface {
}

type BaseResponse struct {
}

type ErrorResponse struct {
	Response struct {
		Error struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
		RequestId string `json:"RequestId"`
	} `json:"Response"`
}



func ParseFromHttpResponse(hr *http.Response, response Response) (err error) {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body because %s", err)
		return errors.NewUnicloudCloudSDKError("ClientError.IOError", msg, "")
	}
	if hr.StatusCode != 200 {
		msg := fmt.Sprintf("Request fail with http status code: %s, with body: %s", hr.Status, body)
		return errors.NewUnicloudCloudSDKError("ClientError.HttpStatusCodeError", msg, "")
	}
	//log.Printf("[DEBUG] Response Body=%s", body)
	/*err = response.ParseErrorFromHTTPResponse(body)
	if err != nil {
		return
	}*/
	err = json.Unmarshal(body, &response)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", string(body), err)
		return errors.NewUnicloudCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	return
}
