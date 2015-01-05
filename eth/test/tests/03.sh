#!/bin/bash

TIMEOUT=35

cat >> $JSFILE <<EOF
eth.addPeer("localhost:30311");
sleep(30000);
eth.export("$CHAIN_TEST");
EOF

peer 11 12k
sleep 2
test_node $NAME "" -loglevel 5 $JSFILE

