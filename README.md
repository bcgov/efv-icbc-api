# go-sample-api

Simple Go "Hello, OpenShift" HTTP API and OpenShift manifests.

Build and push a Docker image (example using Docker):

```bash
# build locally
docker build -t <registry>/<namespace>/go-sample-api:latest .

# push to registry
docker push <registry>/<namespace>/go-sample-api:latest
```

Deploy to OpenShift (replace image if you pushed to registry):

```bash
# apply k8s Deployment and Service
oc apply -f openshift/deployment.yaml
oc apply -f openshift/service.yaml

# create a Route to expose the service
oc apply -f openshift/route.yaml

# or use oc set image to update the placeholder image in the Deployment
oc set image deployment/go-sample-api go-sample-api=<registry>/<namespace>/go-sample-api:latest
```

Quick notes:
- The Dockerfile creates a non-root user (UID 1001) so it is compatible with OpenShift's default SCC.
- The app listens on port `8080`. The Service maps external port 80 to container port 8080.
