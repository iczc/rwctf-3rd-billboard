billboardcli tx billboard create-advertisement $ID $CONTENT --from $KEY --chain-id mainnet --fees 10ctc --node $RPC
billboardcli tx sign tx.json --from $KEY --chain-id mainnet --node $RPC > signtx.json
billboardcli tx broadcast signtx.json --node $RPC
billboardcli tx billboard withdraw $ID 100ctc --from $KEY --chain-id mainnet --fees 10ctc --node $RPC
billboardcli tx billboard ctf $ID --from $KEY --chain-id mainnet --fees 10ctc --node $RPC
