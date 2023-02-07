#!/bin/bash

GENESIS_FILE=~/.lottery/config/genesis.json
if [ -f $GENESIS_FILE ]
then
    echo "Genesis file exist, would you like to delete it? (y/n)"
    read delete_config
fi

if [[
	$delete_config == "Y" ||
	$delete_config == "y" ||
	! -f $GENESIS_FILE
   ]];
then
    rm -r ~/.lottery

    echo "Initialising chain"
    lotteryd init --chain-id=lottery-1 lottery
    echo "y" | lotteryd keys add validator
    echo "y" | lotteryd keys add player01
    echo "y" | lotteryd keys add player02
    echo "y" | lotteryd keys add player03
    echo "y" | lotteryd keys add player04
    echo "y" | lotteryd keys add player05
    echo "y" | lotteryd keys add player06
    echo "y" | lotteryd keys add player07
    echo "y" | lotteryd keys add player08
    echo "y" | lotteryd keys add player09
    echo "y" | lotteryd keys add player10
    echo "y" | lotteryd keys add player11
    echo "y" | lotteryd keys add player12
    echo "y" | lotteryd keys add player13
    echo "y" | lotteryd keys add player14
    echo "y" | lotteryd keys add player15
    echo "y" | lotteryd keys add player16
    echo "y" | lotteryd keys add player17
    echo "y" | lotteryd keys add player18
    echo "y" | lotteryd keys add player19
    echo "y" | lotteryd keys add player20
    echo "y" | lotteryd keys add player21

    echo "video adult rule exhaust tube crater lunch route clap pudding poet pencil razor pluck veteran hill stock thunder sense riot fox oppose glare bar" | lotteryd keys add player21 --recover --keyring-backend test

    echo "Adding genesis account"
    
    # this is to have the accounts on chain
    lotteryd add-genesis-account $(lotteryd keys show player01 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player02 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player03 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player04 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player05 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player06 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player07 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player08 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player09 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player10 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player11 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player12 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player13 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player14 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player15 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player16 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player17 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player18 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player19 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player20 -a) 500000000stake
    lotteryd add-genesis-account $(lotteryd keys show player21 -a) 500000000stake
    ## add the validator 
    lotteryd add-genesis-account $(lotteryd keys show validator -a) 500000000000000stake
    lotteryd gentx validator 100000000000000stake --chain-id lottery-1
    lotteryd collect-gentxs

    # the community tax in the distribution must be disabled, since the community tax is
    # already distributed by the mint module
    echo "$( jq '.app_state.distribution.params.community_tax = "0.1"' ~/.lottery/config/genesis.json )" > ~/.lottery/config/genesis.json

    echo
fi


echo "Starting lottery chain"
lotteryd start
