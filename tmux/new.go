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
func NewSession(n *Name, w *Window) *Session {
	if n == nil {
		return nil
	}

	if w == nil {
		return nil
	}

	return &Session{n, []*Window{w}}
}

// Create a new tmux Window.
func NewWindow(n *Name, p *Pane) *Window {
	if n == nil {
		return nil
	}

	if p == nil {
		return nil
	}

	return &Window{n, []*Pane{p}}
}

// Create a new tmux Pane.
func NewPane(s *StartDir, i InitShellCmd) *Pane {
	if s == nil {
		return nil
	}

	return &Pane{s, i}
}

// Create a new tmux Name.
func NewName(set bool, name string) *Name {
	return &Name{set, name}
}

// Create a new tmux StartDir.
func NewStartDir(set bool, dir string) *StartDir {
	return &StartDir{set, dir}
}
