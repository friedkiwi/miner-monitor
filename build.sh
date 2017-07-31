#!/bin/bash


echo "Building monitor server..."


echo "Building for Linux..."

GOOS=linux GOARCH=arm go build miner-monitor.go 

RETURN=$?

if [ $RETURN -eq 0 ]
then
	echo "Build success."
else 
	echo "Build failed - exitting..."
	exit 1
fi

upx miner-monitor

echo "Copying binary to S7 /tmp"

scp miner-monitor root@10.90.113.122:/tmp/monitor.elf

echo "Transferring control to S7 binary..."
echo ""
echo "-"

ssh root@10.90.113.122 /tmp/monitor.elf
