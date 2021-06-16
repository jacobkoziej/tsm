# Makefile -- automate builds
# Copyright (C) 2021  Jacob Koziej <jacobkoziej@gmail.com>
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

GO      ?= go
GOFLAGS ?= -v

BIN_NAME  ?= tsm
BUILD_DIR ?= build
SOURCES    = cmd/tsm/*

TAG    = $(shell git describe  --abbrev=0        2> /dev/null)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null)
COMMIT = $(shell git rev-parse --short      HEAD 2> /dev/null)

LDFLAGS  = "-X main.tag=$(TAG) -X main.branch=$(BRANCH) -X main.commit=$(COMMIT)"
GOFLAGS += -ldflags $(LDFLAGS)


.PHONY: all
all: build

.PHONY: build
build:
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BIN_NAME) $(SOURCES)

.PHONY: clean
clean:
	@rm -rvf $(BUILD_DIR)
