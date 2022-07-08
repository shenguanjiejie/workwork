#!/bin/bash

RELEASE_VERSION="v0.0.1"
# First check OS.
OS="$(uname)"
if [[ "${OS}" == "Linux" ]]
then
  WW_ON_LINUX=1
elif [[ "${OS}" != "Darwin" ]]
then
  abort "workwork is only supported on macOS and Linux."
fi

if [[ -z "${HOMEBREW_ON_LINUX-}" ]]
then
# TODO
else
# TODO
fi