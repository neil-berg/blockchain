<img src="./assets/blockchain.svg">

[Source](https://www.amplework.com/blockchain-app-development/)


# Blockchain

This work-in-progress project is writing a simplified cryptocurrency (mostly
following Bitcoin) blockchain in Go.

The desire to do so is two-fold: 
1. To learn more about the nuts and bolts of crypto 
2. To learn more about Go!

While the result is not intended to be production-level quality nor used for 
actual transactions, the hope is to create enough realistic features on it to 
mimic the flow of real transactions. 

## Getting started

Clone the repository:

```sh
git clone git@github.com:neil-berg/blockchain.git
```

Build the binary:

```sh
go build main.go
```

Use the CLI (more to come!):

```sh
./main addblock --data "Paying Jane .0005 BTC for coffee"
./main printchain
```

## Database

This blockchain utilitizes the [NutsDB](https://xujiajun.cn/nutsdb/) as a fast
and simple key-value store. [/database](./database/database.go) contains the
custom wrappers around various NutsDB methods.

## Resources

This project largely follows an [excellent tutorial](https://jeiwan.net/posts/building-blockchain-in-go-part-1/)
on building a Blockchain in Go. The main differences are project organization
and the use of NutsDB instead of BoltDB.

## Contributing and feedback

I welcome any contributions or feedback to this project. Please feel free to
file issues or submit PRs. 