// +build !linux,!windows,!plan9,!js,!solaris solaris,!cgo

// Copyright 2020 The Mangos Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ipc implements the IPC transport on top of UNIX domain sockets.
// To enable it simply import it.
package ipc

import (
	"go.nanomsg.org/mangos/v3/transport"
	"net"
)

func getPeer(c *net.UnixConn, pipe transport.ConnPipe) {
}
