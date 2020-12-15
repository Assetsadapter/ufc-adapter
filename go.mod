module github.com/Assetsadapter/ufc-adapter

go 1.13

require (
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/go-owaddress v1.1.13
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.1
	github.com/blocktree/openwallet/v2 v2.0.10
	github.com/Assetsadapter/whitecoin-adapter v1.2.1
)

replace (
	github.com/imroc/req => github.com/blocktree/req v0.2.5
	github.com/Assetsadapter/whitecoin-adapter => /Users/leor/Dev/Code/adapters/whitecoin-adapter
)
