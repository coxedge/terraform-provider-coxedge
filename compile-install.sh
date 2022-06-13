#!/bin/sh
go build -o build/terraform-provider-coxedge
mkdir -p ~/.terraform.d/plugins/coxedge.com/cox/coxedge/0.1/darwin_arm64/
rm ~/.terraform.d/plugins/coxedge.com/cox/coxedge/0.1/darwin_arm64/* || true
mv build/terraform-provider-coxedge ~/.terraform.d/plugins/coxedge.com/cox/coxedge/0.1/darwin_arm64/
