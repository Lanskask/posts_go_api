package errs

type ServiceError struct {
	Message string `json:"message"`
}

func NewServiceError(message string) ServiceError {
	return ServiceError{Message: message}
}

func (s ServiceError) Error() string {
	return s.Message
}
