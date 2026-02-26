SHELL:=$(if $(windir),cmd.exe,/bin/sh)

hrs/serve:
	go run -ldflags="-w -s -X 'main.Env=Dev' -X 'main.DataSourceName=$(GO_DSN_MYSQL)'" ./app/hrs

hrs/build:
	go build -ldflags "-w -s -X 'main.AppTitle=v1.2.3'" ./app/hrs

commit_id:
	@echo $(shell git rev-parse HEAD)