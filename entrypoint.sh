#!/bin/sh

if [ -f "/run/secrets/api_key" ]; then
    export API_KEY=$(cat /run/secrets/api_key)
fi
exec "$@"
