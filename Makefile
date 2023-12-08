# duplicates
# See LICENSE for copyright and license details.
.POSIX:

PREFIX ?= /usr/local
GO ?= go
GOFLAGS ?=
RM ?= rm -f

all: duplicates

duplicates:
	$(GO) build $(GOFLAGS)

clean:
	$(RM) duplicates

install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f duplicates $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/duplicates

uninstall:
	$(RM) $(DESTDIR)$(PREFIX)/bin/duplicates

.DEFAULT_GOAL := all

.PHONY: all duplicates clean install uninstall
