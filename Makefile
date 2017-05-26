NAME=echohead/pomo

all: build

build:
	go build -o pomo .

release: build
	$(eval VERSION := $(shell ./pomo --version))
	git tag $(VERSION)
	git push origin --tags
	hub release create -a pomo $(VERSION)

