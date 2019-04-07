#!/bin/bash

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
