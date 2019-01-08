#!/bin/bash
kill -SIGINT `pgrep omxplayer` 2>/dev/null
echo $1 > arg.txt
omxplayer "$1" --no-keys

