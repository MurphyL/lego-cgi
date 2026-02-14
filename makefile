SHELL:=$(if $(windir),cmd.exe,/bin/sh)

run/cgi:
	go run ./cgi/


shell:
	echo $(shell git rev-parse HEAD)