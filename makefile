SHELL:=$(if $(windir),cmd.exe,/bin/sh)

hrs/serve:
	go run -ldflags="-w -s -X 'main.Env=Dev' -X 'main.DataSourceName=$(GO_DSN_MYSQL)'" ./app/hrs

rpt/serve:
	go run -ldflags="-w -s -X 'main.Env=Dev' -X 'main.DataSourceName=$(GO_DSN_MYSQL)'" ./app/rpt

hrs/build:
	go build -ldflags "-w -s -X 'main.AppTitle=v1.2.3'" ./app/hrs

rpt/build:
	go build -ldflags "-w -s -X 'main.AppTitle=v1.2.3'" ./app/rpt

commit_id:
	@echo $(shell git rev-parse HEAD)