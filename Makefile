all: data/syukujitsu.csv server server-image
data/syukujitsu.csv:
	wget http://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv -O data/syukujitsu.csv
server: **/*.go
	GOOS=linux GOARCH=amd64 go build server.go
server-image: data/syukujitsu.csv config/config.toml server
	docker build -t suzukishunsuke/japanese-holiday-api .
