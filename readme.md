# RealWorld CTF 3rd Finals - Billboard

**billboard** is a blockchain application built using Cosmos SDK and Tendermint.

## Challenge

### Description
Welcome to the billboard, you can post an advertisement on the billboard chain, and the more coin you deposit the higher your advertisement ranking will be.
* Attachment: https://github.com/iczc/billboard/releases/tag/v1.0.0
* Playground: http://ip:80
* RPC: http://ip:26657

### Goal
Send a successful `CaptureTheFlag` type transaction.

### Instruction
1. Add player private key

```
$ echo "your mnemonic here" | billboardcli keys add $KEY --recover
```
mnemonic: chief control turn hurt dance system focus enjoy nasty draw cash rose boring example cause venture neither bind rack seven misery until exhibit hood

>PS: During the competition, we provided players preset built-in accounts rather than the genesis account, then transferred enough coins to those addresses. 
>The preset accounts was derived from player's team tokens in the following algorithm.

>$ billboardcli keys mnemonic --unsafe-entropy  
> \> your team token md5 + 32 * "0"

2. Check balance, should be none-zero

```
$ billboardcli query account $ADDRESS --node $RPC
```

3. Post an advertisement

```
$ billboardcli tx billboard create-advertisement $ID $CONTENT --from $KEY --chain-id mainnet --fees 10ctc --node $RPC
```

### Hint
* The playground website is only used for AD display and TX hash submission (Not a web challenge !!!)
* For fairness, we have banned some query RPC https://github.com/iczc/tendermint/commit/42111f2a780d7910fcc7adac61d65628d0fa4ea7

## Deployment
```
$ git clone https://github.com/iczc/billboard.git
$ cd billboard/deploy
$ docker-compose up -d
```

## Writeup
* http://www.iczc.me/post/rwctf-3rd-billboard-writeup
* https://github.com/iczc/billboard/tree/main/writeup
