setup:
	go install golang.org/x/tools/cmd/godoc@latest
doc:
	if ! godoc -v godoc; then make setup; fi
	godoc -url=/pkg/github.com/NalbertLeal/go-helpers/slice/ > ./docs/slice.html