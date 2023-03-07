package delivery

import "net/http"

func newErrorResponse(writer http.ResponseWriter, code int, err error) error {
	writer.WriteHeader(code)
	_, errWrite := writer.Write([]byte(err.Error()))
	if errWrite != nil {
		return errWrite
	}

	return nil
}
