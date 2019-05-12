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
TESTDATA_DIR=${PROJECT_ROOT}/data/testdata
PKG_TESTDATA=${PROJECT_ROOT}/pkg/bindata/testdata
TEMPLATE=${PROJECT_ROOT}/hack/template/generated.go.template

function gen::bindata::testdata::format() {
  local testdata=${1}
  (cat ${TEMPLATE} && cat ${testdata}) > testdata.go && mv testdata.go ${testdata}
  sed -i -e 's/Sql/SQL/g' ${testdata}
}
 
function gen::bindata::testdata() {
  mkdir -p ${PKG_TESTDATA}
  go-bindata --pkg testdata -o ${PKG_TESTDATA}/testdata.go ${TESTDATA_DIR}
  gen::bindata::testdata::format ${PKG_TESTDATA}/testdata.go
  echo "[+] Generated testdata bindata"
}

gen::bindata::testdata