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
printf "$(tput setaf 40)➜$(tput sgr0)  Downloading distributable package...  "
curl -s ${DIST_URL} > "${PWD}/${RELEASE_FILENAME}";
curl -s ${CHECKSUM_URL} > "${PWD}/${RELEASE_FILENAME}".checksum;

if [[ -e bin/"${RELEASE_FILENAME}" ]]; then
    printf "$(tput setaf 40) ✓\n$(tput sgr0)"
else
    # todo implement fmtErrorBlock
    printf "$(tput setaf 9) ✗\n$(tput sgr0)"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)                                             $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)   Installation failed.                      $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)   Failed downloading distributable file.    $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)                                             $(tput sgr0)\n"
    exit
fi

if [[ -e bin/"${RELEASE_FILENAME}".checksum ]]; then
    MD5_SUM=`cat bin/"${RELEASE_FILENAME}".checksum`
else
    # todo implement fmtErrorBlock
    printf "$(tput setaf 9) ✗\n$(tput sgr0)"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)                                             $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)   Installation failed.                      $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)   Failed downloading checksum.              $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)                                             $(tput sgr0)\n"
    exit
fi

printf "$(tput setaf 40)➜$(tput sgr0)  Verifying downloaded file integrity...";

if [[ "${OS_NAME}" == 'Linux' ]]; then
    FILE_CHECKSUM_MD5=$(md5sum bin/"${RELEASE_FILENAME}");
else
    FILE_CHECKSUM_MD5=$(md5 -q bin/"${RELEASE_FILENAME}");
fi

if [[ "$FILE_CHECKSUM_MD5" == "$MD5_SUM" ]]; then
    printf "$(tput setaf 40) ✓\n$(tput sgr0)"
    rm bin/"${RELEASE_FILENAME}".checksum;
else
    # todo implement fmtErrorBlock
    printf "$(tput setaf 9) ✗\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)                                             $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)   Installation failed.                      $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)   File integrity check error.               $(tput sgr0)\n"
    printf "$(tput setaf 7)$(tput bold)$(tput setab 1)                                             $(tput sgr0)\n"
    exit
fi

mv bin/"${RELEASE_FILENAME}" /usr/local/bin/"${PROJECT_NAME}" && chmod 755 /usr/local/bin/"${PROJECT_NAME}"
printf "$(tput setaf 40)➜$(tput sgr0)  Installing to path directory...        ";
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
