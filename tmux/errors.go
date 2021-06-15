// errors.go -- handle tmux errors
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

type tmuxError interface {
	nilCheck() error
}

func (s *Server) nilCheck() error {
	if s == nil {
		return errors.New("nil Server")
	} else {
		return nil
	}
}

func (s *Session) nilCheck() error {
	if s == nil {
		return errors.New("nil Session")
	} else {
		return nil
	}
}

func (w *Window) nilCheck() error {
	if w == nil {
		return errors.New("nil Window")
	} else {
		return nil
	}
}

func (p *Pane) nilCheck() error {
	if p == nil {
		return errors.New("nil Pane")
	} else {
		return nil
	}
}

func (s *Socket) nilCheck() error {
	if s == nil {
		return errors.New("nil Socket")
	} else {
		return nil
	}
}

func (n *Name) nilCheck() error {
	if n == nil {
		return errors.New("nil Name")
	} else {
		return nil
	}
}

func (i *InitShellCmd) nilCheck() error {
	if i == nil {
		return errors.New("nil InitShellCmd")
	} else {
		return nil
	}
}

func (s *StartDir) nilCheck() error {
	if s == nil {
		return errors.New("nil StartDir")
	} else {
		return nil
	}
}
