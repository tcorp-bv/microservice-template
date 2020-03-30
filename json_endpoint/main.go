/*
 * MIT License
 *
 * Copyright (c) 2020 TCorp BV
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

// If you wish to expose a json interface manually (eg. for testing), this is possible as well.
// This class provides an example for a proxy that translates json to protocolbuffers.

// It requires an extra output when you compile your protocol buffer:
// add `--grpc-gateway_out=logtostderr=true: $GOPATH/src` to protoc to generate the proxy

// Please do not deploy the .gw.go file to source control.

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	pb "github.com/tcorp-bv/microservice-template/pb" // Update
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcAddress = "localhost:8080"
	address     = ":8081"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(address, mux)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
