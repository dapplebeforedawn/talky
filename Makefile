build:
	go build -o bin/talky talky.go

test:
	cd ngrammer && go test

.PHONEY : test
