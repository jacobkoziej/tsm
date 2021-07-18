// arg.go -- tmux argument generation
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

// New generates arguments to create a new session from the Session receiver.
func (s *Session) New() (arg []string) {
	if s == nil {
		return
	}

	arg = append(arg, "new-session")

	if s.behaveAttachSession {
		arg = append(arg, "-A")
	}

	// session isn't attached by default
	if !s.attachSession {
		arg = append(arg, "-d")
	}

	if s.detachClients {
		arg = append(arg, "-D")
	}

	if s.ignoreEnvironment {
		arg = append(arg, "-E")
	}

	if s.printInfo {
		arg = append(arg, "-P")
	}

	if s.detachClient {
		arg = append(arg, "-X")
	}

	if s.startDirectory != nil {
		arg = append(arg, "-c", *s.startDirectory)
	}

	if s.format != nil {
		arg = append(arg, "-F", *s.format)
	}

	if s.windowName != nil {
		arg = append(arg, "-n", *s.windowName)
	}

	if s.groupName != nil {
		arg = append(arg, "-t", *s.groupName)
	}

	if s.width != nil {
		arg = append(arg, "-x", *s.width)
	}

	if s.height != nil {
		arg = append(arg, "-y", *s.height)
	}

	if s.shellCommand != nil {
		arg = append(arg, *s.shellCommand)
	}

	return
}

// New generates arguments to create a new window from the Window receiver.
func (w *Window) New() (arg []string) {
	if w == nil {
		return
	}

	arg = append(arg, "new-window")

	if w.nextIndexUp {
		arg = append(arg, "-a")
	}

	// window isn't made current by default
	if !w.makeCurrent {
		arg = append(arg, "-d")
	}

	if w.destroyTarget {
		arg = append(arg, "-k")
	}

	if w.printInfo {
		arg = append(arg, "-P")
	}

	if w.startDirectory != nil {
		arg = append(arg, "-c", *w.startDirectory)
	}

	if w.environment != nil {
		for _, env := range w.environment {
			arg = append(arg, "-e", env)
		}
	}

	if w.format != nil {
		arg = append(arg, "-F", *w.format)
	}

	if w.windowName != nil {
		arg = append(arg, "-n", *w.windowName)
	}

	if w.targetWindow != nil {
		arg = append(arg, "-t", *w.targetWindow)
	}

	if w.shellCommand != nil {
		arg = append(arg, *w.shellCommand)
	}

	return
}
