#!/bin/bash
docker build -t hyperdrive-oidc-authservice .
k3d image import --cluster hyperdrive-local hyperdrive-oidc-authservice:latest
kubectl rollout restart statefulset -n istio-system authservice