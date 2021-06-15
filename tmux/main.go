// tsm/tmux -- tmux command execution
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

// tmux command execution
package tmux

import "errors"

// Create a new tmux Server.
func New(name string, socket string) *Server {
	return &Server{Name: name, Socket: socket}
}

// Create a new tmux Session.
func (s *Server) NewSession(name string) error {
	if s == nil {
		return errors.New("nil Server")
	}

	s.Sessions = append(s.Sessions, Session{Name: name})
	s.SessionCnt++

	return nil
}

// Create a new tmux Window.
func (s *Session) NewWindow(name string) error {
	if s == nil {
		return errors.New("nil Session")
	}

	s.Windows = append(s.Windows, Window{Name: name})
	s.WindowCnt++

	return nil
}

// Create a new tmux Pane.
func (w *Window) NewPane(dir string, cmd []string, keys []string) error {
	if w == nil {
		return errors.New("nil Window")
	}

	w.Panes = append(w.Panes, Pane{dir, cmd, keys})
	w.PaneCnt++

	return nil
}
