/**
 * Created by chaoluo on 10/27/2016.
 */
package client

import (
	"net/http"
	"errors"
	"strconv"
	"go-http-client/entityMapper"
)

type CResponse struct {
	*http.Response
	Payload      string
	EntityMapper entityMapper.EntityMapper
}

func (cR *CResponse) ForEntity(v interface{})(error){
	statusCode := cR.StatusCode
	// status code 2** is ok
	if statusCode/http.StatusContinue != 2 {
		return errors.New("please check status, status is " + strconv.Itoa(statusCode))
	}
	if cR.EntityMapper == nil {
		return errors.New("entity mapper is empty!")
	}

	error := cR.EntityMapper.Unmarshal([]byte(cR.Payload), &v)
	if error != nil {
		return error
	}
	return nil
}
