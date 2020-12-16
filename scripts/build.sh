#!/usr/bin/env bash

set -e

SRCDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
PROJECT=$(realpath $SRCDIR/..)

bold=$(tput bold)
normal=$(tput sgr0)

usage () {
  cat << HELP
${bold}SYNOPSIS${normal}
  ./build.sh [BUILD VERSION] -E <extra-build> -P

${bold}OPTIONS${normal}
  ${bold}-P${normal} - skip proto compilation

  ${bold}-p${normal} - compile proto files and exit

  ${bold}-E${normal} - build an extra utility
    psu - PSU control utility
HELP
}

mkdir -p build
mkdir -p build/{lnx,win}
mkdir -p build/proto

if [[ ! -z $1 ]] && [[ ! $1 =~ ^- ]]; then
  VERSION=$1
  shift 1
fi
TARGET=tagged
COMPILE_PROTO_ONLY=
SKIP_PROTO=
EXTRA=()

options=':PphE:t:'
while getopts $options option
do
    case "$option" in
        P  ) SKIP_PROTO=true;;
        p  ) COMPILE_PROTO_ONLY=true;;
        t  ) TARGET=$OPTARG;;
        E  ) EXTRA+=( "$OPTARG" );;
        h  ) usage; exit;;
        \? ) echo "Unknown option: -$OPTARG" >&2; exit 1;;
        :  ) echo "Missing option argument for -$OPTARG" >&2; exit 1;;
        *  ) echo "Unimplemented option: -$OPTARG" >&2; exit 1;;
    esac
done

Error() {
	tput setaf 1
	echo $'\t'Error: $1
	tput sgr0
	exit 1
}

GitLatestTag() {
  echo $(git tag --sort=-creatordate | head -n 1)
}

GitCurrentCommitId() {
  echo $(git rev-parse HEAD)
}

CompileProto() {
  # Copy proto files to temp dir
  cp -ar $PROTO_SRC $PROTO_DEST/
  cp -ar $PROTO_ENG_SRC $PROTO_DEST/
  cp -ar $PROTO_FRMW_SRC $PROTO_DEST/
  rm -rf $PROTO_OUT/beamtrail/pcb-debug-util/protobuf/*

  # Remove any usage of nanopb library, for the command-line utility we don't need its optimizations
  find $PROTO_DEST -name '*.proto' -exec sed -i '/import ".*nanopb.proto";/d' {} \;
  find $PROTO_DEST -name '*.proto' -exec sed -r -i 's/\[\(nanopb.*/;/g' {} \;

  for file in $(find $PROTO_DEST -name '*.proto' -type f);
  do
    if [ $(basename $file) == "nanopb.proto" ]; then
      continue
    fi
    echo Compiling ...$(basename $file)
    protoc -I=$PROTO_DEST --go_out=$PROTO_OUT $file
  done
}

go build -o $PROJECT/assets/gserv/middleware/http/static/static.so --buildmode=plugin middleware/static/main.go
go build -o build/gserv $PROJECT/cmd/gserver/*.go

exit

##
##
##

# Check for some build pre-conditions
#if [[ $(git diff --stat) != '' ]]; then
#  Error "repository must be clean"
#fi

if [[ $COMPILE_PROTO_ONLY == "true" ]]; then
  CompileProto
  exit 0
fi

if [[ $SKIP_PROTO != "true" ]]; then
  CompileProto
fi

buildPsuUtil() {
  env LD_LIBRARY_PATH=/usr/local/lib GOOS=linux go build -o build/lnx/$BIN_PSU $GO_PSU
}

buildAppUtil() {
  echo -e "\n"
  echo -n Building app utility...
  env GOOS=windows GOARCH=386 go build -o build/win/"$APPUTIL_BIN.exe" $APPUTIL_MAIN
  env GOOS=linux GOARCH=386 go build -o build/lnx/$APPUTIL_BIN $APPUTIL_MAIN
  echo -e " done\n"
}

buildMainUtil() {
  echo -e "\n"
  echo -n Building ${bold}$VERSION${normal}...
  env GOOS=windows GOARCH=386 go build -o build/win/"$BIN.exe" $GO_MAIN
  env GOOS=linux GOARCH=386 go build -o build/lnx/$BIN $GO_MAIN
  echo -e " done\n"
}

if [[ -z $VERSION ]]; then
  VERSION=$(GitLatestTag)
else
  echo -e "Latest version $(GitLatestTag)\n"
  read -p "Continue? yes or no..." -n 1 -r
  echo    # (optional) move to a new line
  if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    exit 1
  fi
  git tag -a -m "${VERSION}v" ${VERSION} $(git rev-parse master)
fi

cat <<  VER > $PROJECT/cmd/pcb-debug/version.go
package main
// AUTO-GENERATED on $(date +%Y-%m-%d_%H:%M:%S)
var VerStr string = "${VERSION}v"
var BuildStr string = "$(date)"
VER

buildMainUtil

if [[ " ${EXTRA[@]} " =~ " psu " ]]; then
  echo Building PSU utility...
  buildPsuUtil
fi

if [[ " ${EXTRA[@]} " =~ " app " ]]; then
  buildAppUtil
fi
