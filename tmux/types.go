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

// Session represents a tmux session to create.
type Session struct {
	BehaveAttachSession bool    // -A
	DetachSession       bool    // -d
	DetachClients       bool    // -D
	IgnoreEnvironment   bool    // -E
	PrintInfo           bool    // -P
	DetachClient        bool    // -X
	StartDirectory      *string // -c
	Format              *string // -F
	WindowName          *string // -n
	SessionName         *string // -s
	GroupName           *string // -t
	Width               *string // -x
	Height              *string // -y
	ShellCommand        *string
}

// Window represents a tmux window to create.
type Window struct {
	NextIndexUp     bool     // -a
	DontMakeCurrent bool     // -d
	DestroyTarget   bool     // -k
	PrintInfo       bool     // -P
	StartDirectory  *string  // -c
	Environment     []string // -e
	Format          *string  // -F
	WindowName      *string  // -n
	TargetWindow    *string  // -t
	ShellCommand    *string
}
