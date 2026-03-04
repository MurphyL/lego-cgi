SHELL:=$(if $(windir),cmd.exe,/bin/sh)

hrs/serve:
	go run -ldflags="-w -s -X 'main.Env=Dev' -X 'main.DataSourceName=$(GO_DSN_MYSQL)'" ./app/hrs

rpt/serve:
	go run -ldflags="-w -s -X 'main.Env=Dev' -X 'main.DataSourceName=$(GO_DSN_MYSQL)'" ./app/rpt

hrs/build:
	go build -ldflags "-w -s -X 'main.AppTitle=v1.2.3'" ./app/hrs

rpt/build:
	go build -ldflags "-w -s -X 'main.AppTitle=v1.2.3'" ./app/rpt

mysql/run:
	docker run \
		--name mysql-lego -d \
		-p 3306:3306 \
		--restart unless-stopped \
		-v $(res)/res/mysql/log:/var/log/mysql \
		-v $(res)/res/mysql/data:/var/lib/mysql \
		-e MYSQL_ROOT_PASSWORD=123456 \
		mysql

commit_id:
	@echo $(shell git rev-parse HEAD)