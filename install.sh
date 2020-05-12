#!/usr/bin/env bash

tk=`date "+%Y-%m-%d %H:%M:%Sclient" | sha256sum |cut -c 1-60`
sed -i "/^jwt_secret:/cjwt_secret: $tk" conf/config.yaml

