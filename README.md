This script has written for airchain [switchyard testnet campaign](https://paragraph.xyz/@sarox.eth/airchain-rollup?referrer=0xbefEf0FE13B9bD398A88DAB74CCd62099C51333C)

## Install git

```
sudo apt-get install git
```

## Install go

```
VERSION="1.21.6"
ARCH="amd64"
curl -O -L "https://golang.org/dl/go${VERSION}.linux-${ARCH}.tar.gz"
tar -xf "go${VERSION}.linux-${ARCH}.tar.gz"
sudo chown -R root:root ./go
sudo mv -v go /usr/local
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
source ~/.bash_profile
go version
```

## Clone Repository

```
git clone https://github.com/sarox0987/evmos-farmer.git
cd evmos-farmer
```

## Run the script
```
go mod tidy
go run main.go
```

Enter Evmos wallet Hex Private Key, RPC URL, and let it execute!

<img width="503" alt="Screenshot 2024-06-25 at 4 30 34 PM" src="https://github.com/sarox0987/evmos-farmer/assets/153465797/a858ebd2-4045-4e28-bdb5-e5c1fe6ab788">




