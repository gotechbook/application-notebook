.PHONY: run
run:
	go run .

.PHONY: main
main:
	GOOS=linux GOARCH=amd64 go build main.go

.PHONY: doc
doc:
	swag init

.PHONY: abi
abi:
	cd contract && solc --abi *.sol -o ./abi && solc --bin *.sol -o ./abi

.PHONY: game
game:
	cd contract && abigen --bin=./abi/GameItem.bin --abi=./abi/GameItem.abi --pkg=erc721 --out=../contract/erc721/erc721.go

.PHONY: token
token:
	cd contract && abigen --bin=./abi/GLDToken.bin --abi=./abi/GLDToken.abi --pkg=erc20 --out=../contract/erc20/erc20.go

.PHONY: md
md:
	cd docs && i5ting_toc -f doc.md -o

