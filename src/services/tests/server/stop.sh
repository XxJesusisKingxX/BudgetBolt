#!/bin/bash
kill $(ps aux | grep './start.sh' | grep -v grep | awk '{print $2}')