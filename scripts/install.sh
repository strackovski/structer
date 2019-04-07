#!/bin/bash
# ===================================================================
#
# nv3-x/structer installation script
#
# Author:   Vladimir Strackovski <vladimir.strackovski@gmail.com>
# Year:     2019
#
# ===================================================================

source ${PWD}/scripts/.values.sh
source ${PWD}/scripts/.helper.sh

printf "$(tput setaf 69)\n⚡ Installing nv3-x/structer \n$(tput sgr0)"
printf "$(tput setaf 40)➜$(tput sgr0)  Copying executable to /usr/local/bin..."

mv build/"${PROJECT_VERSION}"/"${RELEASE_FILENAME}" /usr/local/bin/"${PROJECT_NAME}" && chmod 755 /usr/local/bin/"${PROJECT_NAME}"
if [[ -e /usr/local/bin/structer ]]; then
    # todo implement fmtSuccessBlock
    printf "$(tput setaf 40)✓\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 40)                                             $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 40)   Installation complete.                    $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 40)                                             $(tput sgr0)\n"
else
    # todo implement fmtErrorBlock
    printf "$(tput setaf 1) ✗\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)                                             $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)   Installation failed.                      $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)   Unable to move file to path directory.    $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)                                             $(tput sgr0)\n"
    exit
fi
