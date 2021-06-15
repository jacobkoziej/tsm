// types.go -- tmux types
// Copyright (C) 2021  Jacob Koziej <jacobkoziej@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package tmux

// start directory properties
type StartDir struct {
	// Specify whether to set the start directory.
	Set bool

	// The start directory.
	Dir string
}

// initial shell command properties
type InitShellCmd struct {
	// Specify whether to execute the shell command.
	Exec bool

	// The shell command.
	Cmd string
}

// name properties
type Name struct {
	set  bool
	name string
}

// tmux pane properties
type Pane struct {
	startDir     StartDir
	initShellCmd InitShellCmd
}

// tmux window properties
type Window struct {
	name *Name
	pane []*Pane
}

// tmux session properties
type Session struct {
	name   *Name
	window []*Window
}

// tmux server properties
type Server struct {
	// The name of the tmux server.
	Name string

	// The path of the tmux server socket.
	Socket string

	// Sessions contained in the tmux server.
	Sessions []Session

	// Count of sessions containted in the tmux server.
	SessionCnt int
}
