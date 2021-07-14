package Week07

import (
	"net"
	"net/http"
)

func Server() error {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token, ok := r.Header["Token"]; ok {
			CheckToken(token[0])
		}
	})

	err = http.Serve(l, handler)
	if err != nil {
		return err
	}
	return nil
}

func CheckToken(token string) (bool, error) {
	return true, nil
}
