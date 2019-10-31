FROM ubuntu:16.04
#ENV MNKEY=""
#ENV EXTIP=""
#COPY dashcore-0.14.0 .
COPY dashcore-0.14.0/bin/dashd /usr/local/bin/
COPY dashcore-0.14.0/conf/dash.conf /root/dash.conf
COPY docker-entrypoint.sh /usr/local/bin/
#COPY dashcore-0.14.0/conf/dash.conf /root/.dashcore/dash.conf
WORKDIR .
#sed '/masternodeprivkey/c masternodeprivkey=$MNKEY' dashcore-0.14.0/conf/dash.conf
#sed '/externalip/c externalip=$EXTIP' dashcore-0.14.0/conf/dash.conf
#COPY docker-entrypoint.sh /usr/local/bin/
ENTRYPOINT ["docker-entrypoint.sh"]
#RUN dashcore-0.14.0/bin/dashd -testnet -daemon
EXPOSE 19999 19999
EXPOSE 19998 19998
#CMD ["dashcore-0.14.0/bin/dashd"]