/*
 * Copyright 2020 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"path/filepath"
	"testing"

	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/v2/openw"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func TestWalletManager_GetTransactions(t *testing.T) {
	tm := testInitWalletManager()
	list, err := tm.GetTransactions(testApp, 0, -1, "Received", false)
	if err != nil {
		log.Error("GetTransactions failed, unexpected error:", err)
		return
	}
	for i, tx := range list {
		log.Info("trx[", i, "] :", tx)
	}
	log.Info("trx count:", len(list))
}

func TestWalletManager_GetTransactionByWxID(t *testing.T) {
	tm := testInitWalletManager()
	wxID := openwallet.GenTransactionWxID(&openwallet.Transaction{
		TxID: "ab41a606f0ef0ce4c000514e8f4dd3a7dd0ecee8",
		Coin: openwallet.Coin{
			Symbol:     "UFC",
			IsContract: true,
			ContractID: "1.3.0",
		},
	})
	log.Info("wxID:", wxID)
	//"D0+rxcKSqEsFMfGesVzBdf6RloM="
	tx, err := tm.GetTransactionByWxID(testApp, wxID)
	if err != nil {
		log.Error("GetTransactionByTxID failed, unexpected error:", err)
		return
	}
	log.Info("tx:", tx)
}

// esQiqx2xKUfGK2pneR9NvCg2GH/vVq5PGMgz2CXKRmw= ab41a606f0ef0ce4c000514e8f4dd3a7dd0ecee8

func TestWalletManager_GetAssetsAccountBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W4pPxx4E4tM1TJWiAyMcGSFMbE8c68URbW"
	accountID := "7z1CyywoB1vC1xnvV4kbvVk1mMjQumV4Vmgp3BbHsHJo"

	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func TestWalletManager_GetAssetsAccountTokenBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W4pPxx4E4tM1TJWiAyMcGSFMbE8c68URbW"
	accountID := "7z1CyywoB1vC1xnvV4kbvVk1mMjQumV4Vmgp3BbHsHJo"

	contract := openwallet.SmartContract{
		Address:  "1.3.0",
		Symbol:   "UFC",
		Name:     "UFC",
		Token:    "UFC",
		Decimals: 4,
	}

	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance.Balance)
}

func TestWalletManager_GetEstimateFeeRate(t *testing.T) {
	tm := testInitWalletManager()
	coin := openwallet.Coin{
		Symbol: "UFC",
	}
	feeRate, unit, err := tm.GetEstimateFeeRate(coin)
	if err != nil {
		log.Error("GetEstimateFeeRate failed, unexpected error:", err)
		return
	}
	log.Std.Info("feeRate: %s %s/%s", feeRate, coin.Symbol, unit)
}

func TestGetAddressVerify(t *testing.T) {
	symbol := "UFC"
	assetsMgr, err := openw.GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}
	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)
	addrDec := assetsMgr.GetAddressDecoderV2()

	flag := addrDec.AddressVerify("zbtest333")
	log.Infof("flag: %v", flag)

}
