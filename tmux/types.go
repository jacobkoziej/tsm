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

// tmux pane properties
type Pane struct {
	// The working directory of the tmux pane.
	Dir string

	// Commands to execute in the tmux pane.
	Cmd []string
}

// tmux window properties
type Window struct {
	// The name of the tmux window.
	Name string

	// The index of the tmux window.
	Index uint

	// Panes contained in the tmux window.
	Panes []Pane
}