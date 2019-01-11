clean:
	go clean

all:
	go clean && go build

s:
	./beegoAppAdmin

start:
	nohup ./beegoAppAdmin &

stop:
	cat `pwd`/gateway.pid | xargs kill

restart: stop start

db:
	mysql -uroot -hlocalhost  --default-character-set=utf8 -e "source scripts/db_beego.sql"

rel:
	go clean
	rm -f gateway
	go build
	tar -czf beegoAppAdmin.tar beegoAppAdmin conf views  static makefile README.md
