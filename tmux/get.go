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
func (s *Server) GetSession(i int) (sn *Session, err error) {
	err = s.nilCheck()
	if err != nil {
		return
	}

	l := len(s.session)

	if l == 0 {
		err = errors.New("no Session(s)")
		return
	}

	if i >= l || i < 0 {
		err = errors.New("invalid index")
		return
	}

	sn = s.session[i]
	return
}

// Return a pointer to a specified Window in a receiver Session.
func (s *Session) GetWindow(i int) (w *Window, err error) {
	err = s.nilCheck()
	if err != nil {
		return
	}

	l := len(s.window)

	if l == 0 {
		err = errors.New("no Window(s)")
	}

	if i >= l || i < 0 {
		return nil, errors.New("invalid index")
	}

	w = s.window[i]
	return
}

// Return a pointer to a specified Pane in a receiver Window.
func (w *Window) GetPane(i int) (p *Pane, err error) {
	err = w.nilCheck()
	if err != nil {
		return
	}

	l := len(w.pane)

	if l == 0 {
		err = errors.New("no Pane(s)")
		return
	}

	if i >= l || i < 0 {
		err = errors.New("invalid index")
		return
	}

	p = w.pane[i]
	return
}
