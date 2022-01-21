#!/bin/bash
#
##
TYPE=$1
NAME=$2
STATE=$3
HELPERYAML=/opt/helpernode/etc/helper.yaml
export PATH=/usr/local/bin:$PATH
case $STATE in
        "MASTER") sleep 5; helpernodectl start --config ${HELPERYAML} --skip-preflight
                  echo "$(date +%c) - Switching to MASTER, starting helpernode " | tee -a /var/log/ha-helper.log
                  ;;
        "BACKUP") helpernodectl stop
                  echo "$(date +%c) - Switching to BACKUP, stopping helpernode " | tee -a /var/log/ha-helper.log
                  exit 0
                  ;;
        "FAULT")  helpernodectl stop
                  echo "$(date +%c) - FAULTING, helpernode STOPPING " | tee -a /var/log/ha-helper.log
                  exit 0
                  ;;
        *)        echo "$(date +%c) - FATAL, Unknown helpernode state " | tee -a /var/log/ha-helper.log
                  exit 1
                  ;;
esac
##
##
