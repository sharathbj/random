#!/bin/sh
ttl=$1
msg=$2
echo 'printing title and message'
echo $ttl
echo $msg
/Applications/terminal-notifier.app/Contents/MacOS/terminal-notifier -title $ttl -message $msg
