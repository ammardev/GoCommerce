package http

type ValidationErrors map[string]string

func (err ValidationErrors) Error() string {
	return ""
}

func (err ValidationErrors) Check() error {
	if len(err) > 0 {
		return &err
	}

	return nil
}
