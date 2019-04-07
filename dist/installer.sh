#!/bin/bash
# ===================================================================
#
# nv3-x/structer installation script
#
# Author:   Vladimir Strackovski <vladimir.strackovski@gmail.com>
# Year:     2019
#
# ===================================================================


containsElement () {
  local e match="$1"
  shift
  for e; do [[ "$e" == "$match" ]] && return 0; done
  return 1
}

fmtYellowRev () {
    local e icon="$1"
    local e text="$2"
    shift
    printf "${icon} $(tput setaf 3)$(tput bold)$(tput rev)${text}$(tput sgr0)"

    return 0
}

fmtGreenRev () {
    local e icon="$1"
    local e text="$2"
    shift
    printf "${icon} $(tput setaf 2)$(tput bold)$(tput rev)${text}$(tput sgr0)"

    return 0
}

fmtMagenta () {
    local e icon="$1"
    local e text="$2"
    shift
    printf "${icon} $(tput setaf 5)$(tput bold)${text}$(tput sgr0)"

    return 0
}

fmtGreen () {
    local e text="$1"
    shift
    printf "$(tput setaf 2)$(tput bold)${text}$(tput sgr0)"

    return 0
}


ARCH_NAME=`uname -m`
ARCH=""
OS_NAME=`uname -s`
OS=$(echo "${OS_NAME}" | tr '[:upper:]' '[:lower:]')

if [[ "${ARCH_NAME}" == 'x86_64' ]]; then
     ARCH='amd64'
 elif [[ "${ARCH_NAME}" == 'x86_64' ]]; then
    ARCH='amd64'
fi

EXCLUDE_OS=(android aix nacl plan9)
EXCLUDE_ARCH=(arm arm64 wasm)
PLATFORM="${OS}"_"${ARCH}"
PROJECT_NAME="structer"
CURRENT_VERSION='0.1.0'
PROJECT_VERSION="${CURRENT_VERSION}"
REMOTE_ROOT='https://s3.eu-central-1.amazonaws.com'
RELEASE_FILENAME="${PROJECT_NAME}"_"${PROJECT_VERSION}"_"${PLATFORM}"
CHECKSUM_FILENAME="${RELEASE_FILENAME}".checksum
INSTALLER_PATH="${PWD}"/dist/installer.sh
DIST_URL="${REMOTE_ROOT}/${PROJECT_NAME}/${PROJECT_VERSION}/${RELEASE_FILENAME}";
CHECKSUM_URL="${REMOTE_ROOT}/${PROJECT_NAME}/${PROJECT_VERSION}/${CHECKSUM_FILENAME}";
INSTALLER_URL="${REMOTE_ROOT}/${PROJECT_NAME}/dist/${PROJECT_NAME}_installer.sh";

source ${PWD}/scripts/.helper.sh

printf "$(tput setaf 69)\n⚡ Installing nv3-x/structer \n$(tput sgr0)"
printf "$(tput setaf 40)➜$(tput sgr0)  Downloading distributable package...  "
curl -s ${DIST_URL} > "${PWD}/${RELEASE_FILENAME}";
curl -s ${CHECKSUM_URL} > "${PWD}/${RELEASE_FILENAME}".checksum;

if [[ -e "${RELEASE_FILENAME}" ]]; then
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

if [[ -e "${RELEASE_FILENAME}".checksum ]]; then
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

mv "${RELEASE_FILENAME}" /usr/local/bin/"${PROJECT_NAME}" && chmod 755 /usr/local/bin/"${PROJECT_NAME}"
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
