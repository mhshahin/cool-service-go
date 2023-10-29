#!/usr/bin/env sh

echo "rolling out coolservice ..."
helm upgrade ${SERVICE_NAME:-coolservice} . \
--install \
--history-max 5 \
--atomic \
--values ${VALUES_FILE:-values.yaml} \
--namespace ${SERVICE_NAMESPACE:-mhshahin} \
--set image.tag=${SERVICE_IMAGE_TAG:-master} \
