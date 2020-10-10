#!/bin/bash

set -ue

# Expect the following envvars to be set:
# - APP
# - VERSION
# - COMMIT
# - TARGET_OS
# - LEDGER_ENABLED
# - DEBUG

# Source builder's functions library
. /usr/local/share/cosmos-sdk/buildlib.sh

# These variables are now available
# - BASEDIR
# - OUTDIR

# Build for each os-architecture pair
for platform in ${TARGET_PLATFORMS} ; do
    # This function sets GOOS, GOARCH, and OS_FILE_EXT environment variables
    # according to the build target platform. OS_FILE_EXT is empty in all
    # cases except when the target platform is 'windows'.
    setup_build_env_for_platform "${platform}"

    make clean
    echo Building for $(go env GOOS)/$(go env GOARCH) >&2
    GOROOT_FINAL="$(go env GOROOT)" \
    make ${APP} \
        LDFLAGS="-w -s -buildid=${VERSION}" \
        VERSION=${VERSION} \
        COMMIT=${COMMIT} \
        LEDGER_ENABLED=${LEDGER_ENABLED}
    mv ${APP}${OS_FILE_EXT} ${OUTDIR}/${APP}-${VERSION}-$(go env GOOS)-$(go env GOARCH)${OS_FILE_EXT}

    # This function restore the build environment variables to their
    # original state.
    restore_build_env
done

# Generate and display build report.
generate_build_report
cat ${OUTDIR}/build_report
