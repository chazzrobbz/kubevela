package workloaddefinition

var ContainerizedWorkload = `apiVersion: core.oam.dev/v1alpha2
kind: WorkloadDefinition
metadata:
  name: containerizedworkloads.core.oam.dev
  annotations:
    oam.appengine.info/apiVersion: "core.oam.dev/v1alpha2"
    oam.appengine.info/kind: "ContainerizedWorkload"
spec:
  definitionRef:
    name: containerizedworkloads.core.oam.dev
  childResourceKinds:
    - apiVersion: apps/v1
      kind: Deployment
    - apiVersion: v1
      kind: Service
  extension:
    template: |
      #Template: {
      	apiVersion: "core.oam.dev/v1alpha2"
      	kind:       "ContainerizedWorkload"
      	metadata: name: containerized.name
      	spec: {
      		containers: [{
      			image: containerized.image
      			name:  containerized.name
      			ports: [{
      				containerPort: containerized.port
      				protocol:      "TCP"
      				name:          "default"
      			}]
      		}]
      	}
      }
      containerized: {
      	name: string
      	// +usage=specify app image
      	// +short=i
      	image: string
      	// +usage=specify port for container
      	// +short=p
      	port: *6379 | int
      }
`