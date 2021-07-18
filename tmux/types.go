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
	BehaveAttachSession bool
	DetachSession       bool
	DetachClients       bool
	IgnoreEnvironment   bool
	PrintInfo           bool
	DetachClient        bool
	StartDirectory      *string
	Format              *string
	WindowName          *string
	SessionName         *string
	GroupName           *string
	Width               *string
	Height              *string
	ShellCommand        *string
}

// Window represents a tmux window to create.
type Window struct {
	NextIndexUp     bool
	DontMakeCurrent bool
	DestroyTarget   bool
	PrintInfo       bool
	StartDirectory  *string
	Environment     []string
	Format          *string
	WindowName      *string
	TargetWindow    *string
	ShellCommand    *string
}
