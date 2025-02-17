// Copyright 2021-2025 The phy-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fake

import "fmt"

type ErrorType int

const (
	ErrorTypeUnknown ErrorType = iota
	ErrorTypeInvalidRequest
	ErrorTypeNotFound
	ErrorTypeConflict
)

func (e ErrorType) String() string {
	switch e {
	case ErrorTypeInvalidRequest:
		return "invalid request"
	case ErrorTypeNotFound:
		return "not found"
	case ErrorTypeConflict:
		return "conflict"
	default:
		return "unknown error"
	}
}

type Error struct {
	Type          ErrorType
	Resource      string
	Id            interface{}
	msgFmtAndVars []interface{}
}

func NewError(errorType ErrorType, resource string, id interface{}, msgFmtAndVars ...interface{}) *Error {
	return &Error{
		Type:          errorType,
		Resource:      resource,
		Id:            id,
		msgFmtAndVars: msgFmtAndVars,
	}
}

func (e *Error) Error() string {
	return fmt.Errorf("%s: %s[%s]%s", e.Type, e.Resource, e.Id, e.message()).Error()
}

func (e *Error) message() string {
	switch len(e.msgFmtAndVars) {
	case 0:
		return ""
	case 1:
		return " " + e.msgFmtAndVars[0].(string)
	default:
		return " " + fmt.Sprintf(e.msgFmtAndVars[0].(string), e.msgFmtAndVars[1:]...)
	}
}
