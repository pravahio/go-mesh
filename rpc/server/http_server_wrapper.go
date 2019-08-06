package server

import (
	"net"
	"net/http"
)

type WrapperHTTP struct {
	ser *http.Server
}

func (wh *WrapperHTTP) Serve(l net.Listener) error {
	return wh.ser.Serve(l)
}

func (wh *WrapperHTTP) Stop() {
	wh.ser.Close()
}
