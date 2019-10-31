#!/bin/bash
if [[ $mnkey ]]; then
  sed -i '/masternodeblsprivkey/c masternodeblsprivkey='$mnkey'' /root/dash.conf
fi
if [[ $externalip ]]; then
  sed -i '/externalip/c externalip='$externalip'' /root/dash.conf
fi
sleep 10000
#dashd -testnet -conf=/root/dash.conf
exec "$@"
