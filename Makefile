all: syukujitsu.csv server server-image
syukujitsu.csv:
	wget http://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv
server: server.go lib/*.go
	GOOS=linux GOARCH=amd64 go build server.go
server-image: syukujitsu.csv config.toml server
	docker build -t suzukishunsuke/japanese-holiday-api .
