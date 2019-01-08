#!/bin/bash
kill -SIGINT `pgrep omxplayer` 2>/dev/null
DISPLAY=:0 omxplayer "$1" --no-keys >/dev/null

