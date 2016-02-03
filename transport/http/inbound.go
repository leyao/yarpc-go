// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package http

import (
	"io"
	"net"
	"net/http"

	"github.com/yarpc/yarpc-go/transport"

	"golang.org/x/net/context"
)

// Inbound builds a new HTTP inbound that listens on the given address.
func Inbound(addr string) transport.Inbound {
	return &httpInbound{addr: addr}
}

type httpInbound struct {
	addr     string
	listener net.Listener
}

func (i *httpInbound) Serve(h transport.Handler) error {
	var err error
	i.listener, err = net.Listen("tcp", i.addr)
	if err != nil {
		return err
	}

	server := &http.Server{Addr: i.addr, Handler: httpHandler{h}}
	return server.Serve(i.listener)
	// TODO Handle connection close errors
}

func (i *httpInbound) Close() error {
	if i.listener == nil {
		return nil
	}
	return i.listener.Close()
}

// httpHandler adapts a transport.Handler into a handler for net/http.
type httpHandler struct {
	Handler transport.Handler
}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.NotFound(w, req)
	}

	defer req.Body.Close()

	procedure := req.Header.Get(ProcedureHeader)
	if len(procedure) == 0 {
		http.Error(w, "procedure name is required", http.StatusBadRequest)
		return
	}
	req.Header.Del(ProcedureHeader)

	treq := &transport.Request{
		Procedure: procedure,
		Headers:   fromHTTPHeader(req.Header, nil),
		Body:      req.Body,
	}

	tres, err := h.Handler.Handle(
		context.TODO(), // TODO
		treq,
	)
	if err != nil {
		// TODO structured responses?
		err = internalError{Reason: err}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toHTTPHeader(tres.Headers, w.Header())

	defer tres.Body.Close()
	if _, err := io.Copy(w, tres.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}