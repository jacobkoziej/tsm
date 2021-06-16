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
	set bool
	dir string
}

// initial shell command properties
type InitShellCmd struct {
	exec bool
	cmd  string
}

// name properties
type Name struct {
	set  bool
	name string
}

// socket properties
type Socket struct {
	exists bool
	path   string
}

// tmux pane properties
type Pane struct {
	startDir     *StartDir
	initShellCmd *InitShellCmd
	done         bool
}

// tmux window properties
type Window struct {
	name *Name
	pane []*Pane
	done bool
}

// tmux session properties
type Session struct {
	name   *Name
	window []*Window
	done   bool
}

// tmux server properties
type Server struct {
	name    *Name
	socket  *Socket
	session []*Session
	done    bool
}
