#!/bin/bash
set -e

# core configuration
INPUT_SUBDIRECTORY="${INPUT_SUBDIRECTORY:-"."}"
INPUT_PRUNE_COUNT="${INPUT_PRUNE_COUNT:-"0"}"
INPUT_BENCHMARKS_OUT="${INPUT_BENCHMARKS_OUT:-"benchmarks.json"}"
INPUT_GO_TEST_PKGS="${INPUT_GO_TEST_PKGS:-"./..."}"
INPUT_GO_BENCHMARKS="${INPUT_GO_BENCHMARKS:-"."}"
INPUT_GIT_COMMIT_MESSAGE="${INPUT_GIT_COMMIT_MESSAGE:-"add benchmark run for ${GITHUB_SHA}"}"

# publishing configuration
INPUT_PUBLISH_REPO="${INPUT_PUBLISH_REPO:-${GITHUB_REPOSITORY}}"
INPUT_PUBLISH_BRANCH="${INPUT_PUBLISH_BRANCH:-"gh-pages"}"

# pull request checks
INPUT_CHECKS="${INPUT_CHECKS:-"false"}"
INPUT_CHECKS_CONFIG="${INPUT_CHECKS_CONFIG:-"gobenchdata-checks.yml"}"

# output build data
echo '========================'
command -v gobenchdata
gobenchdata version
env | grep 'INPUT_'

# The default env.GITHUB_ACTOR support seems broken.  Try a different argument
if [[ "$OVERRIDE_GITHUB_ACTOR" ]]
then
  GITHUB_ACTOR="$OVERRIDE_GITHUB_ACTOR"
fi
if [[ "$OVERRIDE_GITHUB_TOKEN" ]]
then
  GITHUB_TOKEN="$OVERRIDE_GITHUB_TOKEN"
fi
if [[ -z "$GITHUB_ACTOR_EMAIL" ]]
then
  GITHUB_ACTOR_EMAIL="${GITHUB_ACTOR}@users.noreply.github.com"
fi
echo "GITHUB_ACTOR=${GITHUB_ACTOR}"
echo "GITHUB_TOKEN=${GITHUB_TOKEN}"
echo "GITHUB_WORKSPACE=${GITHUB_WORKSPACE}"
echo "GITHUB_REPOSITORY=${GITHUB_REPOSITORY}"
echo "GITHUB_SHA=${GITHUB_SHA}"
echo "GITHUB_REF=${GITHUB_REF}"
echo '========================'

# setup
mkdir -p /tmp/{gobenchdata,build}
git config --global user.email "${GITHUB_ACTOR_EMAIL}"
git config --global user.name "${GITHUB_ACTOR}"
git config --global url."https://${GITHUB_ACTOR}:${GITHUB_TOKEN}@github".insteadOf https://github
git remote set-url origin https://x-access-token:${GITHUB_TOKEN}@github.com/${INPUT_PUBLISH_REPO}

# Set ssh keys for git
mkdir -p ~/.ssh
ssh-keyscan github.com >> ~/.ssh/known_hosts
ssh-agent -a ${SSH_AUTH_SOCK} > /dev/null
ssh-add - <<< "${SSH_KEY}"

# run benchmarks from configured directory
echo
echo '📊 Running benchmarks...'
RUN_OUTPUT="/tmp/gobenchdata/benchmarks.json"
cd "${GITHUB_WORKSPACE}"
cd "${INPUT_SUBDIRECTORY}"
# https://github.community/t/environment-variables-are-overwritten-by-previous-actions-and-break-consecutive-docker-actions/122532/2
unset GOROOT

# For some reason, our modified version hits this issue.
# export PATH=$PATH:$(go env GOPATH)/bin
go test \
  -bench "${INPUT_GO_BENCHMARKS}" \
  -benchmem \
  ${INPUT_GO_TEST_FLAGS} \
  ${INPUT_GO_TEST_PKGS} \
  | gobenchdata --json "${RUN_OUTPUT}" -v "${GITHUB_SHA}" -t "ref=${GITHUB_REF}"
cd "${GITHUB_WORKSPACE}"

# fetch published data
echo
echo "📚 Checking out ${INPUT_PUBLISH_REPO}@${INPUT_PUBLISH_BRANCH}..."
cd /tmp/build
echo "Executing 'git clone https://${GITHUB_ACTOR}:hiddentoken@github.com/${INPUT_PUBLISH_REPO}.git .'"
git clone https://${GITHUB_ACTOR}:${GITHUB_TOKEN}@github.com/${INPUT_PUBLISH_REPO}.git .
echo "Executing 'git checkout ${INPUT_PUBLISH_BRANCH}'"
git checkout ${INPUT_PUBLISH_BRANCH}
echo

if [[ "${INPUT_CHECKS}" == "true" ]]; then

  # check results against published
  echo '🔎 Evaluating results against base runs...'
  CHECKS_OUTPUT="/tmp/gobenchdata/checks-results.json"
  gobenchdata checks eval "${INPUT_BENCHMARKS_OUT}" "${RUN_OUTPUT}" \
    --checks.config "${GITHUB_WORKSPACE}/${INPUT_CHECKS_CONFIG}" \
    --json ${CHECKS_OUTPUT} \
    --flat
  RESULTS=$(cat ${CHECKS_OUTPUT})
  echo "::set-output name=checks-results::$RESULTS"

  # output results
  echo
  echo '📝 Generating checks report...'
  gobenchdata checks report ${CHECKS_OUTPUT}

fi

if [[ "${INPUT_PUBLISH}" == "true" ]]; then

  # merge results with published
  echo '☝️ Updating results...'
  if [[ -f "${INPUT_BENCHMARKS_OUT}" ]]; then
    echo '📈 Existing report found - merging...'
    gobenchdata merge "${RUN_OUTPUT}" "${INPUT_BENCHMARKS_OUT}" \
      --prune "${INPUT_PRUNE_COUNT}" \
      --json "${INPUT_BENCHMARKS_OUT}" \
      --flat
  else
    cp "${RUN_OUTPUT}" "${INPUT_BENCHMARKS_OUT}"
  fi
  # Actually generate the webpage.
  [ ! -d "app" ] && mkdir app
  # I can't get the benchmarks.json to be read from the parent directory so I'm copying it to the app directory.
  cp "${INPUT_BENCHMARKS_OUT}" ./app/benchmarks.json
  gobenchdata web generate ./app

  # publish results
  echo
  echo '📷 Committing and pushing new benchmark data...'
  git add .
  git commit -m "${INPUT_GIT_COMMIT_MESSAGE}"
  git push -f origin ${INPUT_PUBLISH_BRANCH}

fi

echo
echo '🚀 Done!'
