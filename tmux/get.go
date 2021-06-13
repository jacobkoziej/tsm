// get.go -- get individual tmux types
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

import "errors"

// Return a pointer to a specified Session in a receiver Server.
func (s *Server) GetSession(i int) (*Session, error) {
	if s == nil {
		return nil, errors.New("nil Server")
	}

	l := len(s.Sessions)

	if l == 0 {
		return nil, errors.New("no Sessions")
	}

	if i >= l || i < 0 {
		return nil, errors.New("invalid index")
	}

	return &s.Sessions[i], nil
}

// Return a pointer to a specified Window in a receiver Session.
func (s *Session) GetWindow(i int) (*Window, error) {
	if s == nil {
		return nil, errors.New("nil Session")
	}

	l := len(s.Windows)

	if l == 0 {
		return nil, errors.New("no Windows")
	}

	if i >= l || i < 0 {
		return nil, errors.New("invalid index")
	}

	return &s.Windows[i], nil
}

// Return a pointer to a specified Pane in a receiver Window.
func (w *Window) GetPane(i int) (*Pane, error) {
	if w == nil {
		return nil, errors.New("nil Window")
	}

	l := len(w.Panes)

	if l == 0 {
		return nil, errors.New("no Panes")
	}

	if i >= l || i < 0 {
		return nil, errors.New("invalid index")
	}

	return &w.Panes[i], nil
}
