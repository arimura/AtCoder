pid = $$(ps aux| grep stub/main | grep -v grep | perl -anle 'print $$F[1]')
.PHONY: run stop
# pid = ""

run: 
	go run stub/main.go &
	go run main.go
stop:
	- ps | grep -v grep | grep stub/main && kill $(pid)
