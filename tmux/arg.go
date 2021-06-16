// arg.go -- argument generation for tmux
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

// Return appropriate server arguments.
func (s *Server) ArgServer() (arg []string, err error) {
	err = s.nilCheck()
	if err != nil {
		return
	}

	if s.socket.exists {
		arg = append(
			arg,
			Commands.Socket.PathFlag,
			s.socket.path,
		)
		return
	}

	arg = append(arg, Commands.Socket.NameFlag)

	if s.name.set {
		arg = append(arg, s.name.name)
	} else {
		arg = append(arg, Commands.Socket.DefaultName)
	}

	return
}

// Return appropriate new-session arguments.
func (s *Session) ArgNewSession() (arg []string, err error) {
	err = s.nilCheck()
	if err != nil {
		return
	}

	arg = append(
		arg,
		Commands.NewSession.NewSession,
		Commands.NewSession.DetachedFlag,
	)

	if s.name.set {
		arg = append(
			arg,
			Commands.NewSession.SessionNameFlag,
			s.name.name,
		)
	}

	w, err := s.GetWindow(0)
	if err != nil {
		arg = nil
		return
	}

	if w.name.set {
		arg = append(
			arg,
			Commands.NewSession.WindowNameFlag,
			w.name.name,
		)
	}

	p, err := w.GetPane(0)
	if err != nil {
		arg = nil
		return
	}

	if p.startDir.set {
		arg = append(
			arg,
			Commands.NewSession.StartDirFlag,
			p.startDir.dir,
		)
	}

	if p.initShellCmd.exec {
		arg = append(
			arg,
			p.initShellCmd.cmd,
		)
	}

	return
}
