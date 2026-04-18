# efv-icbc-api

Simple Go HTTP API with OpenAPI docs and OpenShift manifests.

Build and push a Docker image (example using Docker):

```bash
# build locally
docker build -t <registry>/<namespace>/efv-icbc-api:latest .

# push to registry
docker push <registry>/<namespace>/efv-icbc-api:latest
```

Deploy to OpenShift (replace image if you pushed to registry):

```bash
# apply k8s Deployment and Service
oc apply -f openshift/deployment.yaml
oc apply -f openshift/service.yaml

# create a Route to expose the service
oc apply -f openshift/route.yaml

# or use oc set image to update the placeholder image in the Deployment
oc set image deployment/efv-icbc-api efv-icbc-api=<registry>/<namespace>/efv-icbc-api:latest
```

Run locally

Prerequisites

- Go 1.20+ installed and on your PATH

Windows (PowerShell)

```powershell
# optional: set a different port
$env:PORT="8080"

# run the service
go run .
```

macOS / Linux (bash)

```bash
# optional: set a different port
export PORT=8080

# run the service
go run .
```

Sanity checks

```powershell
# health
Invoke-WebRequest -UseBasicParsing 'http://localhost:8080/health' | Select-Object -ExpandProperty Content

# openapi spec
Invoke-WebRequest -UseBasicParsing 'http://localhost:8080/openapi.yaml' | Select-Object -ExpandProperty Content

# open Swagger UI in the browser
Start-Process 'http://localhost:8080/swagger/'
```

Troubleshooting

- If `go run .` fails with a bind error (port in use), find and kill the process using port 8080.

	PowerShell:

	```powershell
	$p = (Get-NetTCPConnection -LocalPort 8080 -ErrorAction SilentlyContinue | Select-Object -Expand OwningProcess -Unique)
	if ($p) { Stop-Process -Id $p -Force }
	```

	macOS / Linux:

	```bash
	lsof -i :8080
	kill <PID>
	```

- To run on a different port without killing processes, set `PORT` before running as shown above.


API endpoints (mock implementations)
- `GET /api/v1/driver/validate?license_number=...&province=...` — validate a driver (province must be `BC` for success)
- `GET /api/v1/insurance/status?vin=...` — returns mock insurance status
- `GET /api/v1/vehicle/verify?vin=...&plate=...&province=...` — verify vehicle (province must be `BC` for success)

Notes
- The server serves the OpenAPI spec at `/openapi.yaml` and a small Swagger UI at `/swagger/`.
- The Dockerfile uses a non-root user (UID 1001) to be OpenShift-compatible.
- The app listens on port `8080` by default.

Deployed Swagger UI

- A deployed instance of the Swagger UI is available at:

- https://efv-icbc-api-17db4f-dev.apps.silver.devops.gov.bc.ca/swagger/

	You can also load that deployed OpenAPI spec into the local Swagger UI via the selector in the top-left.
