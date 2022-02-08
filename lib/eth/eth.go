package eth

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func Client(url string) (cli *ethclient.Client, err error) {
	return ethclient.Dial(url)
}

func GetAuth(ctx context.Context, cli *ethclient.Client, pk string) (auth *bind.TransactOpts, err error) {
	var (
		privateKey  *ecdsa.PrivateKey
		publicKey   *ecdsa.PublicKey
		fromAddress common.Address
		nonce       uint64
		gasPrice    *big.Int
	)
	privateKey, err = crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}
	publicKey = privateKey.Public().(*ecdsa.PublicKey)
	fromAddress = crypto.PubkeyToAddress(*publicKey)
	nonce, err = cli.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err = cli.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	auth = bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice
	return auth, nil
}

