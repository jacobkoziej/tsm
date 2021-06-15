// new.go -- create new tmux types
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

// Create a new tmux Server.
func New(name string, socket string) *Server {
	return &Server{Name: name, Socket: socket}
}

// Create a new tmux Session.
func NewSession(name *Name, window *Window) *Session {
	if name == nil {
		return nil
	}

	if window == nil {
		return nil
	}

	return &Session{name, []*Window{window}}
}

// Create a new tmux Window.
func NewWindow(name *Name, pane *Pane) *Window {
	if name == nil {
		return nil
	}

	if pane == nil {
		return nil
	}

	return &Window{name, []*Pane{pane}}
}

// Create a new tmux Pane.
func NewPane(startDir StartDir, initShellCmd InitShellCmd) *Pane {
	return &Pane{startDir, initShellCmd}
}

// Create a new tmux Name.
func NewName(set bool, name string) *Name {
	return &Name{set, name}
}
