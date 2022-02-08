abigen = /Users/dnice/go/bin/abigen

main:
	GOOS=linux GOARCH=amd64 go build main.go

doc:
	swag init

abi:
	cd contract &&  solc  --abi *.sol -o ./abi && solc  --bin *.sol -o ./abi

token:
	cd contract && ${abigen} --bin=./abi/PixelToken.bin --abi=./abi/PixelToken.abi --pkg=contract --out=../contract/pixel_erc20.go && ${abigen} --bin=./abi/PixelItem.bin --abi=./abi/PixelItem.abi --pkg=contract --out=./pixel_erc721.go

md:
	cd docs && i5ting_toc -f doc.md -o
.PHONY: md