#!/usr/bin/env bash
BINDIR=$HOME/.tokopedia_grpc_config
PROTOC_VERSION=3.5.1
PROTOC_GO_VERSION=v1.2.0
OS=
function get_os() {
  # Detect the platform (similar to $OSTYPE)
  case $(uname) in
  'Linux')
    OS='Linux'
    ;;
  'Darwin')
    OS='OSX'
    ;;
  *)
    echo "Unsupported OS, only support Linux-64 & OSX-64"
    exit 1
    ;;
  esac
}
function install_protoc() {
  PROTOC_ZIP=protoc-${PROTOC_VERSION}-$(echo "$OS" | awk '{print tolower($0)}')-x86_64.zip
  echo "Installing protoc ${PROTOC_ZIP}"
  curl -OL "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}"
  unzip -o $PROTOC_ZIP -d $BINDIR/protoc_$PROTOC_VERSION
  rm -f $PROTOC_ZIP
}
if [[ ! -d $BINDIR ]]; then
  mkdir $BINDIR
fi
if [ ! -d "$BINDIR/protoc_$PROTOC_VERSION" ]; then
  mkdir $BINDIR/protoc_$PROTOC_VERSION
fi
get_os
if [ ! -f "$BINDIR/protoc_$PROTOC_VERSION/bin/protoc" ]; then
  install_protoc
fi
if [[ $($BINDIR/protoc_$PROTOC_VERSION/bin/protoc --version | sed 's/libprotoc //') != $PROTOC_VERSION ]]; then
  install_protoc
fi
# protoc-gen-go doesn't have flag version
# we check it by the binary directory
if [[ ! -f "${BINDIR}/protoc_gen_go_${PROTOC_GO_VERSION}/protoc-gen-go" ]]; then
  echo "installing protoc gen go"
  if [ ! -d "$GOPATH/src/github.com/golang/protobuf" ]; then
      echo "protobuf is not found, download it"
      git clone git@github.com:golang/protobuf.git $GOPATH/src/github.com/golang/protobuf
  fi
  git -C $GOPATH/src/github.com/golang/protobuf checkout $PROTOC_GO_VERSION
  CWD=$(pwd)
  # disable GO111MODULE
#  cd $GOPATH/src/github.com/golang/protobuf/protoc-gen-go && GO111MODULE= go build
#  mkdir ${BINDIR}/protoc_gen_go_${PROTOC_GO_VERSION}
#  mv protoc-gen-go ${BINDIR}/protoc_gen_go_${PROTOC_GO_VERSION}
#  cd $CWD
   GO111MODULE=off go install github.com/golang/protobuf/protoc-gen-go
   mkdir ${BINDIR}/protoc_gen_go_${PROTOC_GO_VERSION}
   if [ -f $GOPATH/bin/protoc-gen-go ]; then
      mv $GOPATH/bin/protoc-gen-go ${BINDIR}/protoc_gen_go_${PROTOC_GO_VERSION}
   elif [ -f $GOPATH/bin/go/protoc-gen-go ]; then
      mv $GOPATH/bin/go/protoc-gen-go ${BINDIR}/protoc_gen_go_${PROTOC_GO_VERSION}
   else
     echo "Failed to install protoc-gen-go"
     exit 1
   fi
fi
cd proto && $BINDIR/protoc_$PROTOC_VERSION/bin/protoc \
 --plugin=${BINDIR}/protoc_gen_go_${PROTOC_GO_VERSION}/protoc-gen-go \
 --go_out=plugins=grpc:. *.proto