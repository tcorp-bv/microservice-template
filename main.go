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

import (
	"context"
	"fmt"

	"log"
	"net"

	pb "github.com/tcorp-bv/microservice-template/pb"
	"google.golang.org/grpc"
)

const (
	msgFmt = "Hello, %s!"
	port   = "localhost:8080"
)

// Implements GreeterServer from pb/greeter.pb.go
type server struct {
	pb.UnimplementedGreeterServer
}

// Implements SayHello from the GreeterServer
func (s *server) SayHello(_ context.Context, p *pb.Person) (*pb.HelloReply, error) {
	r := pb.HelloReply{Message: fmt.Sprintf(msgFmt, p.Name)}
	return &r, nil
}

func run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

}
