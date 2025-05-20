package errors

/*
func (e *Error) ToErrorJson(method string) *ErrorJson {
	detail := e.Message
	if e.Inner != nil {
		detail += ": " + e.Inner.Error()
	}
	return &ErrorJson{
		Code:       int(e.Kind),
		Message:    e.Kind.Error(),
		Detail:     detail,
		Method:     method,
		Attributes: e.Metadata,
	}
}

func ErrorFromJson(jsonErr string) error {
	unmarshaled := &ErrorJson{}
	// this is a bit of a weird case but oh well - need to think of a better interface
	err := json.Unmarshal([]byte(jsonErr), unmarshaled)
	if err != nil {
		return err
	}
}

type ErrorJson struct {
	Code       int               `json:"code"`
	Message    string            `json:"message"`
	Detail     string            `json:"detail"`
	Method     string            `json:"method"`
	Attributes map[string]string `json:"attributes"`
}
*/
