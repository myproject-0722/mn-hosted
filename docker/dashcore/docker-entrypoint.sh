#!/bin/bash
if [[ $mnkey ]]; then
  sed -i '/masternodeblsprivkey/c masternodeblsprivkey='$mnkey'' /root/dash.conf
fi
if [[ $externalip ]]; then
  sed -i '/externalip/c externalip='$externalip'' /root/dash.conf
fi
#dashd -testnet -conf=/root/dash.conf -daemon
dashd -testnet
#dashd -testnet -conf=/root/dash.conf
exec "$@"
