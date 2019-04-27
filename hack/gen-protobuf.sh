#!/bin/bash
# Copyright 2019 Publr Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

PROJECT_ROOT=$(dirname ${BASH_SOURCE})/..
API_DIR=${PROJECT_ROOT}/api
PKG_DIR=${PROJECT_ROOT}/pkg
INTERNAL_DIR=${PROJECT_ROOT}/internal
GOOGLEAPIS=${PROJECT_ROOT}/third_party/googleapis/api-common-protos
TEMPLATE=${PROJECT_ROOT}/hack/template/generated.go.template

function gen::protobuf::format() {
  local protobuf=${1}
  echo -e "$(cat ${TEMPLATE} ${protobuf}.pb.go)" > ${protobuf}.pb.go
}

function gen::protobuf() {
  for apis in $(find ${API_DIR} -type f -name '*.proto'); do
    local api_proto_path=${apis#${PROJECT_ROOT}/}
    local proto_path=${api_proto_path#api/}
    local descriptor_path=${INTERNAL_DIR}/proto_descriptor/${proto_path%.proto}
    mkdir -p $(dirname ${descriptor_path})
    protoc ${api_proto_path} \
    -I ${GOPATH}/src:${PROJECT_ROOT} \
    -I ${GOOGLEAPIS} \
    --include_imports \
    --include_source_info \
    --descriptor_set_out=${descriptor_path}.pb \
    --go_out=plugins=grpc:${GOPATH}/src
    gen::protobuf::format ${PKG_DIR}/${api_proto_path%.proto}
    echo "[+] Generated $(basename ${api_proto_path%.proto}) proto"
  done
}

gen::protobuf