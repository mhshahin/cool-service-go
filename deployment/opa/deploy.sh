#!/usr/bin/env sh

echo "rolling out OPA ..."
helm upgrade ${SERVICE_NAME:-opa} . \
--install \
--history-max 5 \
--atomic \
--values ${VALUES_FILE:-values.yaml} \
--namespace ${SERVICE_NAMESPACE:-mhshahin} \
--set image.tag=${SERVICE_IMAGE_TAG:-edge-rootless} \