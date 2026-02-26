SHELL:=$(if $(windir),cmd.exe,/bin/sh)

cgi/serve:
	go run ./app/hrs

cgi/build:
	go build -ldflags "-w -s" ./app/hrs

commit_id:
	@echo $(shell git rev-parse HEAD)