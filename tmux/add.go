// add.go -- add to tmux types
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

// Add an additional session to an existing server.
func (s *Server) AddSession(sn *Session) (err error) {
	if err = s.nilCheck(); err != nil {
		return
	}

	if err = sn.nilCheck(); err != nil {
		return
	}

	s.Sessions = append(s.Sessions, *sn)
	s.SessionCnt++

	return
}

// Add an additional window to an existing session.
func (s *Session) AddWindow(w *Window) (err error) {
	if err = s.nilCheck(); err != nil {
		return
	}

	if err = w.nilCheck(); err != nil {
		return
	}

	s.Windows = append(s.Windows, *w)
	s.WindowCnt++

	return
}
