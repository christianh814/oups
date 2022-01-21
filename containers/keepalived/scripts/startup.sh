#!/bin/bash
#
## This is the startup script for the Keealived container

#
## Variables for keepalived contianer
kadConfig=/etc/keepalived/keepalived.conf
helperNodeYaml=/opt/helpernode/etc/helper.yaml
kadTemplate=/usr/local/src/keepalived.conf.j2
ansibleLog=/var/log/helper-ansiblerun.log
helpernodectl=/usr/local/bin/helpernodectl

#
## Make sure the HELPERPOD_CONFIG_YAML env var has size
[[ ${#HELPERPOD_CONFIG_YAML} -eq 0 ]] && echo "FATAL: HELPERPOD_CONFIG_YAML env var not set!!!" && exit 254

#
## Take the HELPERPOD_CONFIG_YAML env variable and write out the YAML file.
echo ${HELPERPOD_CONFIG_YAML} | base64 -d > ${helperNodeYaml}

#
## Set up the keepalived config
ansible localhost -c local -e @${helperNodeYaml} -m template -a "src=${kadTemplate} dest=${kadConfig}" > ${ansibleLog} 2>&1

#
## Make sure helpernodectl binary is available
[[ ! -f  ${helpernodectl} ]] && echo "FATAL: helpernodectl not found!!!" && exit 254

#
## Prepull the images to speedup the keepalived process
echo "Pulling helpernode images"
echo "========================="
${helpernodectl} pull

#
## Holder until we can figure out how to "check" the config
if false ; then
	echo "================================"
	echo "FATAL: Invalid Keepalived config"
	echo "================================"
	exit 254
else
	echo "=============================="
	echo "Starting Keepalived service..."
	echo "=============================="
	# Holder for now
	## /usr/sbin/keepalived -f ${kadConfig} --dont-fork -l
	sleep infinity
fi
##
##
