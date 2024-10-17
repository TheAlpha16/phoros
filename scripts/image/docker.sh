#!/bin/sh

set +x

cd ../..

image="phoros-test"
container="phoros-test"
SESSION_SECRET="heckingisfun"
ADMIN_SECRET="heckingisfun"
EVENT_START=""
EVENT_END=""
POST_EVENT="false"
OBJECT_STORE="native"
APP_PORT="9049"

docker rm -f ${container}
docker rmi -f ${image}

docker build -t ${image} .
docker run -d --name ${container} \
    -p ${APP_PORT}:${APP_PORT} \
    -e SESSION_SECRET=${SESSION_SECRET} \
    -e ADMIN_SECRET=${ADMIN_SECRET} \
    -e OBJECT_STORE=${OBJECT_STORE} \
    -e APP_PORT=${APP_PORT} \
    -e EVENT_START=${EVENT_START} \
    -e EVENT_END=${EVENT_END} \
    -e POST_EVENT=${POST_EVENT} \
    ${image}