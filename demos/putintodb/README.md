

prerequisites:

kind installed locally

ollama installed

nomic-embed-text:latest model is pulled

kind create cluster --config tests/integration/deploys/kind-config.yaml

install ToolHive

helm upgrade -i toolhive-operator-crds oci://ghcr.io/stacklok/toolhive/toolhive-operator-crds

helm upgrade -i toolhive-operator oci://ghcr.io/stacklok/toolhive/toolhive-operator -n toolhive-system --create-namespace

install milvus

kubectl apply -f https://raw.githubusercontent.com/zilliztech/milvus-operator/main/deploy/manifests/deployment.yaml

kubectl apply -f https://raw.githubusercontent.com/zilliztech/milvus-operator/main/config/samples/milvus_cluster_woodpecker.yaml

create MCP ACCESS TOKEN secret

kubectl create secret generic githubtoken --from-literal=MCP_ACCESS_TOKEN="github token"

install RemoteMCPTool CRD

kubectl apply -f k8s/config/crd/bases

move to demos/putintodb directory

cd demos/putintodb

deploy maestro-k

kubectl apply -f maestro-k-mcp-deployment.yaml

install remote-github MCP and remote-maestro-k tools

uv run maestro create  tools.yaml

run workflow

uv run maestro run agents.yaml workflow.yaml
