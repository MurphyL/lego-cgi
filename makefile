SHELL:=$(if $(windir),cmd.exe,/bin/sh)

cgi/serve:
	go run ./cgi/

cgi/build:
	go build -ldflags "-w -s" ./cgi/

commit_id:
	@echo $(shell git rev-parse HEAD)