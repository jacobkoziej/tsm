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

	if s.BehaveAttachSession {
		arg = append(arg, "-A")
	}

	if s.DetachSession {
		arg = append(arg, "-d")
	}

	if s.DetachClients {
		arg = append(arg, "-D")
	}

	if s.IgnoreEnvironment {
		arg = append(arg, "-E")
	}

	if s.PrintInfo {
		arg = append(arg, "-P")
	}

	if s.DetachClient {
		arg = append(arg, "-X")
	}

	if s.StartDirectory != nil {
		arg = append(arg, "-c", *s.StartDirectory)
	}

	if s.Format != nil {
		arg = append(arg, "-F", *s.Format)
	}

	if s.WindowName != nil {
		arg = append(arg, "-n", *s.WindowName)
	}

	if s.GroupName != nil {
		arg = append(arg, "-t", *s.GroupName)
	}

	if s.Width != nil {
		arg = append(arg, "-x", *s.Width)
	}

	if s.Height != nil {
		arg = append(arg, "-y", *s.Height)
	}

	if s.ShellCommand != nil {
		arg = append(arg, *s.ShellCommand)
	}

	return
}

// New generates arguments to create a new window from the Window receiver.
func (w *Window) New() (arg []string) {
	if w == nil {
		return
	}

	arg = append(arg, "new-window")

	if w.NextIndexUp {
		arg = append(arg, "-a")
	}

	if w.DontMakeCurrent {
		arg = append(arg, "-d")
	}

	if w.DestroyTarget {
		arg = append(arg, "-k")
	}

	if w.PrintInfo {
		arg = append(arg, "-P")
	}

	if w.StartDirectory != nil {
		arg = append(arg, "-c", *w.StartDirectory)
	}

	if w.Environment != nil {
		for _, env := range w.Environment {
			arg = append(arg, "-e", env)
		}
	}

	if w.Format != nil {
		arg = append(arg, "-F", *w.Format)
	}

	if w.WindowName != nil {
		arg = append(arg, "-n", *w.WindowName)
	}

	if w.TargetWindow != nil {
		arg = append(arg, "-t", *w.TargetWindow)
	}

	if w.shellCommand != nil {
		arg = append(arg, *w.shellCommand)
	}

	return
}
