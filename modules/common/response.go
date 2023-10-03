package common

type SuccessResponse struct {
	Message  string       `json:"message"`
	Metadata *interface{} `json:"metadata,omitempty"`
	Data     *interface{} `json:"data,omitempty"`
}

type FailResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func GetResponseSuccess(data, metadata any) SuccessResponse {
	res := SuccessResponse{
		Message:  "success",
		Metadata: nil,
		Data:     nil,
	}

	if data != nil {
		res.Data = &data
	}

	if metadata != nil {
		res.Metadata = &metadata
	}

	return res
}

func GetResponseFail(err error) FailResponse {
	return FailResponse{
		Message: "fail",
		Error:   err.Error(),
	}
}
