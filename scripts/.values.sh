#!/bin/bash
# ===================================================================
#
# nv3-x/structer installer
#
# Author:   Vladimir Strackovski <vladimir.strackovski@gmail.com>
# Year:     2019
#
# ===================================================================

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
INSTALLER_URL="${REMOTE_ROOT}/${PROJECT_NAME}/scripts/${PROJECT_NAME}_installer.sh";

