#!/bin/bash

# to get the balance of a player run 
# lotteryd query bank balances $(lotteryd keys show -a player20)

echo "place bets"

lotteryd tx lottery place-bet 10000000 --from player01 --chain-id lottery-1 -b block -y 
lotteryd tx lottery place-bet 11000000 --from player02 --chain-id lottery-1 -b block -y 
lotteryd tx lottery place-bet 12000000 --from player03 --chain-id lottery-1 -b block -y 
lotteryd tx lottery place-bet 13000000 --from player04 --chain-id lottery-1 -b block -y  
lotteryd tx lottery place-bet 14000000 --from player05 --chain-id lottery-1 -b block -y 
lotteryd tx lottery place-bet 15000000 --from player06 --chain-id lottery-1 -b block -y 
lotteryd tx lottery place-bet 16000000 --from player07 --chain-id lottery-1 -b block -y 
lotteryd tx lottery place-bet 17000000 --from player08 --chain-id lottery-1 -b block -y 
lotteryd tx lottery place-bet 18000000 --from player09 --chain-id lottery-1 -b block -y 
lotteryd tx lottery place-bet 19000000 --from player10 --chain-id lottery-1 -b block -y 
lotteryd tx lottery place-bet 20000000 --from player11 --chain-id lottery-1 -b block -y 
lotteryd tx lottery place-bet 21000000 --from player12 --chain-id lottery-1 -b block -y 
lotteryd tx lottery place-bet 22000000 --from player13 --chain-id lottery-1 -b block -y 
## invalid bet - bet too big
lotteryd tx lottery place-bet 222000000 --from player13 --chain-id lottery-1 -b block
## invalid bet - bet too small
lotteryd tx lottery place-bet 500000 --from player13 --chain-id lottery-1 -b block





