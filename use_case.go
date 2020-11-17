package calc

func DoUseCase(req Request) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	operation := Operations[req.Operation]
	result := Do(operation.lambda, req.Arguments)

	res := Response{
		Operation: req.Operation,
		Arguments: req.Arguments,
		Result:    result,
	}

	return &res, nil
}
