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
SCHEMA_DIR=${PROJECT_ROOT}/data/schema
PKG_SCHEMA=${PROJECT_ROOT}/pkg/bindata/schema
TEMPLATE=${PROJECT_ROOT}/hack/template/generated.go.template

function gen::bindata::schema::format() {
  local schema=${1}
  (cat ${TEMPLATE} && cat ${schema}) > schema.go && mv schema.go ${schema}
  sed -i -e 's/Sql/SQL/g' ${schema}
  gofmt -s -w ${schema}
}
 
function gen::bindata::schema() {
  mkdir -p ${PKG_SCHEMA}
  go-bindata --pkg schema -o ${PKG_SCHEMA}/schema.go ${SCHEMA_DIR}
  gen::bindata::schema::format ${PKG_SCHEMA}/schema.go
  echo "[+] Generated schema bindata"
}

gen::bindata::schema