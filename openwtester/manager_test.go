package openwtester

import (
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
	"github.com/blocktree/openwallet/v2/openwallet"
)

var (
	testApp        = "assets-adapter"
	configFilePath = filepath.Join("conf")
)

func testInitWalletManager() *openw.WalletManager {
	log.SetLogFuncCall(true)
	tc := openw.NewConfig()

	tc.ConfigDir = configFilePath
	tc.EnableBlockScan = false
	tc.SupportAssets = []string{
		"UFC",
	}

	return openw.NewWalletManager(tc)
}

func TestWalletManager_CreateWallet(t *testing.T) {
	tm := testInitWalletManager()
	w := &openwallet.Wallet{Alias: "HELLO UFC 3", IsTrust: true, Password: "12345678"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("wallet:", nw)
	log.Info("key:", key)

}

func TestWalletManager_GetWalletInfo(t *testing.T) {

	tm := testInitWalletManager()

	wallet, err := tm.GetWalletInfo(testApp, "WC7xUtsRRjfpbMHsv1XEPyyZt9zBqh1MrQ")
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	log.Info("wallet:", wallet)
}

func TestWalletManager_GetWalletList(t *testing.T) {

	tm := testInitWalletManager()

	list, err := tm.GetWalletList(testApp, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("wallet[", i, "] :", w)
	}
	log.Info("wallet count:", len(list))

	tm.CloseDB(testApp)
}

func TestWalletManager_CreateAssetsAccount(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WMUNKe9munq8jCKV8M9BMZ3GgmsmXupSVQ"
	account := &openwallet.AssetsAccount{Alias: "zbcat999", WalletID: walletID, Required: 1, Symbol: "UFC", IsTrust: true}
	account, address, err := tm.CreateAssetsAccount(testApp, walletID, "12345678", account, nil)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("account:", account)
	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAssetsAccountList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WC7xUtsRRjfpbMHsv1XEPyyZt9zBqh1MrQ"
	list, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("account[", i, "] :", w)
	}
	log.Info("account count:", len(list))

	tm.CloseDB(testApp)

}

func TestWalletManager_CreateAddress(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WMUNKe9munq8jCKV8M9BMZ3GgmsmXupSVQ"
	accountID := "H21A7sN6LFJisHoVFhoakMWGbmZkucsd4v3avCAWiSY"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 1)
	if err != nil {
		log.Error(err)
		return
	}

	for i, w := range address {
		log.Info("address[", i, "] :", w.Address)
	}

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAddressList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WC7xUtsRRjfpbMHsv1XEPyyZt9zBqh1MrQ"
	accountID := "3UNnJBcgYbwXNtZLHJzKBQFSz4FXC3rKuwb3UoyWaz7r" //zbalice999 UFCNdyTKXkLQbNCCDvi1WCiSTYh2pikL85Kwg

	// walletID := "WNThpn66zHFXYAHkBuUMTpgzndfcuY68hw"
	// accountID := "Hh3pAc3NuCTLcYFXnfJJqnYWykk5tLfnLbrn3MzM6bRu" //zbbob111 UFCNbbkUWSthHrf8wBGkaSzsHkrMrCKo97gfb

	// walletID := "WMUNKe9munq8jCKV8M9BMZ3GgmsmXupSVQ"
	// accountID := "H21A7sN6LFJisHoVFhoakMWGbmZkucsd4v3avCAWiSY" //zbcat999 UFCNVrqLPWJQVFbxY4mEuKsJcBYEgxbCMPStD

	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("address[", i, "] :", w.Address)
	}
	log.Info("address count:", len(list))

	tm.CloseDB(testApp)
}
