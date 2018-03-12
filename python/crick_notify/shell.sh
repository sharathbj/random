#!/bin/sh
ttl=$1
msg=$2
#/Applications/terminal-notifier.app/Contents/MacOS/terminal-notifier -title $1 -message $2
terminal-notifier -title "$1" -message "$2" -open "https://cricbuzz.com"
sleep 5
terminal-notifier -title "commentory" -message "$3" -open "https://cricbuzz.com"
