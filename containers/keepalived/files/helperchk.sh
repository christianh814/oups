#!/bin/bash
#
##
if [[ $(/usr/local/bin/helpernodectl status | wc -l) -eq 0 ]]; then
	exit 13
fi
#
exit 0
##
##
