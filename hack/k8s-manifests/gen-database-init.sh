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

PROJECT_ROOT=$(dirname ${BASH_SOURCE})/../..
SCHEMA_DIR=${PROJECT_ROOT}/data/schema
DEPLOYMENT_DIR=${PROJECT_ROOT}/deployment
TEMPLATE=${PROJECT_ROOT}/hack/template/generated.yaml.template

function gen::database::init::format() {
  local schema=${1}
  sed -i '/--/d' ${schema}
  echo -e "$(cat ${TEMPLATE} ${schema})" > ${schema}
}

function gen::database::init() {
    in=${SCHEMA_DIR}/
    out=${DEPLOYMENT_DIR}/kubernetes/publr/database/configmap-initialization.yaml
    kubectl create configmap publr-database-initialization \
    --from-file=$in \
    --dry-run -o yaml > $out
    gen::database::init::format $out
    echo "Generated publr-database-initialization configmap"
}

gen::database::init