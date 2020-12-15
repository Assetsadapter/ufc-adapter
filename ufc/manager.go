/*
 * Copyright 2020 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package ufc

import (
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/whitecoin-adapter/whitecoin"
	"github.com/assetsadapterstore/ufc-adapter/addrdec"
)

type WalletManager struct {
	*whitecoin.WalletManager
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = whitecoin.NewWalletManager()
	wm.Config = whitecoin.NewConfig(Symbol)
	wm.Decoder = NewAddressDecoder(&wm)
	wm.DecoderV2 = addrdec.NewAddressDecoderV2()
	wm.Log = log.NewOWLogger(wm.Symbol())
	return &wm
}
