#!/bin/bash
docker build -t hyperdrive-oidc-authservice -t hyperdrive-oidc-authservice:latest -f docker/oidc-authservice-dev.Dockerfile .
k3d image import --cluster hyperdrive-local hyperdrive-oidc-authservice:latest
kubectl rollout restart statefulset -n istio-system authservice
