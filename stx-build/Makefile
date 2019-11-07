#
# SPDX-License-Identifier: Apache-2.0
#
# Copyright (C) 2019 Intel Corporation
#

.PHONY:	clean cce cce-image

cce:
	go build -o dist/cce ../cmd/cce

cce-image:
	docker build -f Dockerfile_cce -t stx-cce:1.0 .

clean:
	rm -rf dist/*
