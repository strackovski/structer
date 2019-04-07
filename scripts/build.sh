#!/bin/bash
# ===================================================================
#
# nv3-x/structer build script
#
# Author:   Vladimir Strackovski <vladimir.strackovski@gmail.com>
# Year:     2019
#
# ===================================================================

source ${PWD}/scripts/.values.sh
source ${PWD}/scripts/.helper.sh

fmtMagenta "⌛" "Preparing dependencies\n"

if which jq >/dev/null || [[ -x jq ]]; then
    printf "   jq "
    fmtGreen '✓\n'
else
    printf '   Downloading jq...'
    `curl http://github.com/stedolan/jq/releases/download/jq-1.5/jq-osx-amd64 -k -L -o jq`
    `chmod +x jq`
fi

declare -a AR=(`go tool dist list -json | ./jq -r '[(. | unique | .[] | [.GOOS, .GOARCH ])] | sort | .[] | (.[0]) + "_"+ ( .[1])'`)

fmtMagenta "\n⌛" "Uploading installer\n"
awsres=$(aws s3api put-object --acl public-read --profile fdigolan --body dist/installer.sh --bucket structer --key dist/structer_installer.sh)
printf "   Installer successfully uploaded "
fmtGreen '✓\n'

fmtMagenta "\n⌛" "Build started, version: ${PROJECT_VERSION}\n"
for ((idx=0; idx<${#AR[@]}; ++idx)); do
    IFS="_ " read -r -a array <<< "${PROJECT_NAME}_${AR[idx]}"
    ARR1="${array[1]}"
    ARR2="${array[2]}"
    BUILD_FNAME="${PROJECT_NAME}"_"${PROJECT_VERSION}"_"${AR[idx]}"
    CHECKSUM_FNAME="${BUILD_FNAME}".checksum
    BUILD_FPATH=build/"${PROJECT_VERSION}"/"${BUILD_FNAME}"
    CHECKSUM_FPATH=build/"${PROJECT_VERSION}"/"${CHECKSUM_FNAME}"

    if containsElement "$ARR1" "${EXCLUDE_OS[@]}" || containsElement "$ARR2" "${EXCLUDE_ARCH[@]}"; then
        printf "   Skipping '$ARR1/$ARR2'\n"
    else
        fmtYellowRev "⚡" "Building '$ARR1/$ARR2'\n"
        printf "   Building............."
        `GOOS="${array[1]}" GOARCH="${array[2]}" go build -o "${BUILD_FPATH}" app/main.go`
        fmtGreen '✓\n'
        fmtGreen "☘  Built '$ARR1/$ARR2'\n"

        if [[ "${OS}" == 'linux' ]]; then
            cd "$PWD" && md5sum "${BUILD_FPATH}" > "${CHECKSUM_FPATH}"
        else
            cd "$PWD" && md5 -q "${BUILD_FPATH}" > "${CHECKSUM_FPATH}"
        fi

        if [[ ! -z "$1" ]] && [[ $1 == publish ]]
          then
            printf "   Uploading checksum..."
            ARES=$(aws s3api put-object --acl public-read --profile fdigolan --body "${CHECKSUM_FPATH}" --bucket "${PROJECT_NAME}" --key "${PROJECT_VERSION}"/"${CHECKSUM_FNAME}")
            fmtGreen '✓\n'
            printf "   Publishing..........."
            ARES=$(aws s3api put-object --acl public-read --profile fdigolan --body "${BUILD_FPATH}" --bucket "${PROJECT_NAME}" --key "${PROJECT_VERSION}"/"${BUILD_FNAME}")
            fmtGreen '✓\n'
            fmtGreen "✈  Published '$ARR1/$ARR2'\n"
        fi
    fi
done

fmtMagenta "\n⌛" "Cleaning up\n"
`rm jq`
printf "   Removed jq "
fmtGreen '✓\n'
