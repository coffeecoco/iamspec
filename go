#!/bin/bash
set -e

SCRIPT_DIR=$(cd $(dirname $0) ; pwd -P)

tf() {
  pushd "${SCRIPT_DIR}/example_infra" > /dev/null
    terraform init
    terraform $1
  popd > /dev/null
}

goal_test() {
  pushd "${SCRIPT_DIR}" > /dev/null
    bundle exec rake spec
  popd > /dev/null
}

goal_test-unit() {
  pushd "${SCRIPT_DIR}" > /dev/null
    bundle exec rspec --exclude-pattern spec/integration_spec.rb
  popd > /dev/null
}


goal_example-infra-plan() {
  tf "plan"
}

goal_example-infra-apply() {
  tf "apply"
}

goal_setup() {
  pushd "${SCRIPT_DIR}" > /dev/null
    bundle install
  popd > /dev/null
}

if type -t "goal_$1" &>/dev/null; then
  goal_$1 ${@:2}
else
  echo "usage: $0 <goal>
goal:
    setup                   -- install all dependencies, make repo ready for development

    test                    -- run all tests
    test-unit               -- run only unit tests

    example-infra-plan      -- terraform plan on example infra
    example-infra-apply     -- terraform apply on example infra
"
  exit 1
fi
