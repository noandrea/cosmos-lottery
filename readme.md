> Disclaimer: this project was developed as a coding challenge

# Lottery
**Lottery** is a protocol that implements a simple lottery mechanism. 


## Code structure

The implementation of the lottery protocol is contained in the `x/lottery` folder, and the main logic for the implementation is contained in the following sources:

- [`x/lottery/keeper/msg_server.go`](x/lottery/keeper/msg_server.go) - responsible for processing incoming bet transactions
-  [`x/lottery/keeper/keeper.go`](x/lottery/keeper/msg_server.go) - responsible for processing incoming bet transactions
- [`x/lottery/module/abci.go`](x/lottery/module/abci.go) - logic to execute the lottery at the end of each block. 

The message for participating in a lottery is defined in the [`proto/lottery/tx.proto`](proto/lottery/tx.proto) file.

## Quickstart

To start the chain locally run the command 

```
make start-dev
```

this will setup a node with 22 accounts:

- 1 validator account: validator
- 21 accounts to use for betting: player01 - player21

(For more information about the chain setup check the [`scripts/seeds/00_start_chain.sh`](scripts/seeds/00_start_chain.sh) script)

To run a round of lottery bets run the seed script

```
scripts/seeds/01_bets.sh
```

> Note: logs entries for the lottery can be filtered by the `module=x/lottery` string


## Caveats

- Amounts are expressed in micro tokens.
- Automated testing has not been implemented, specifically the `msg_server_test.go` and `abci_test.go` are the critical sections that, in a production roadmap, should be implemented first.
- Protocol queries are missing, in particular, it would be useful to have a query that returns the current lottery balance.
- Governance support to change the parameters of the chain, specifically for the lottery fee, and the min-max bet amounts, has not been implemented.

## Challenge issues

#### Section: Enter Lottery transaction

> Lottery fee is 5token , minimal bet is 1token

since the fee is a set amount, it is not necessary to have it as a transaction parameter.

#### Section: Lottery blocks

> the chosen block proposer can't have any lottery transactions with itself as a sender, if this is the case, then the lottery won’t fire this block, and continue on the next one

With this clause, there is a risk that the lottery will never end.

> if the same user has new lottery transactions, then only the last one counts, counter doesn’t increase on substitution.

In the implementation, an account can only make a single bet. 

#### Section: Choosing a winner

> At the end of the lottery, on the block end, append the data of the transactions (retaining their order) , then hash the data to get the result.

To simplify the logic for calculating the hash for the lottery is different in the implementation and it is based on updating and storing a hash with the data of new transactions when they are received.



