package main

import (
	"fmt"
	"./wallet"
	// "math/big"
)

var (
	nonece uint64 = 0
)

//Just for test
func main() {
	number := 1
	err := GenerateWallets(uint32(number))
	if err != nil {
		println(err.Error())
		return
	}
}

func GenerateWallets(number uint32) (error) {
	var err error
	{
		ret,err:=wallet.CreateWalletByteRandAndPwd([]byte("sdfafdsaf"),"123456")
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(ret)
	}


	{
		// var secret, salt string
		// secret = "ShowSplashViewShowSplashViewShowSplashViewShowSplashView"
		// salt = "1234567890"
		// var byteSecret []byte = []byte(secret)
		// var byteSalt []byte = []byte(salt)
		// //wa, err := wallet.NewWalletAccount(wp.SecretBytes(), wp.SaltBytes())
		// // addr, privateKey, err := wallet.NewWalletAccount(byteSecret, byteSalt)
		// addrandprivateKey, err := wallet.TestbtcNewWalletAccount(byteSecret, byteSalt)
		// if err != nil {
		// 	return err
		// }
		// fmt.Println("The Private key is  " + addrandprivateKey)
	}
	// fmt.Println(" The Address is  " + addr)
	// testAddr := "0x4661dbc978fd123e2250a33c9eedcfeec3746ec5"
	// signedData, _ := wallet.SendETHRawTxByPrivateKey(privateKey, nonece+3, testAddr, big.NewInt(1000000000), big.NewInt(21000), big.NewInt(18000000000), nil)
	// fmt.Println("The real signed hex string is ", signedData)
	{
		// addrandprivateKey, err := wallet.TestbtcNewWalletAccount2()
		// if err != nil {
		// 	return  err
		// }
		// fmt.Println("The Private key is  " + addrandprivateKey)
	}
	return err
}
