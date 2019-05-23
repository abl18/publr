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
PROJECT_PACKAGE=github.com/prksu/publr
API_DIR=${PROJECT_ROOT}/api
PKG_DIR=${PROJECT_ROOT}/pkg
MOCK_DIR=${PKG_DIR}/service/mock
TEMPLATE=${PROJECT_ROOT}/hack/template/generated.go.template

function gen::mock::format() {
  local mock=${1}
  echo -e "$(cat ${TEMPLATE} ${mock})" > ${mock}
  sed -i -e "s#${PROJECT_ROOT}/##g" ${mock}
  gofmt -s -w ${mock}
}

function gen::mock::grpc() {
  for api in $(find ${PKG_DIR}/api -type f -name '*.pb.go'); do
    local package_dir=${PROJECT_PACKAGE}$(dirname ${api#${PROJECT_ROOT}})
    local package_name=$(basename ${api%.pb.go})
    local output=${MOCK_DIR}/${package_name}/mock.go
    mockgen -package=${package_name} ${package_dir} "$(echo "${package_name%s}" | sed 's/.*/\u&/')"ServiceClient > ${output}
    gen::mock::format ${output}
    echo "[+] Generated $(basename ${package_name}) grpc mock"
  done
}

gen::mock::grpc