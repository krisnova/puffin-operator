
default: container


container:
	operator-sdk build krisnova/puffinoperator:latest
	docker push krisnova/puffinoperator:latest

api:
	operator-sdk generate k8s

rbac:
	kubectl create -f deploy/rbac.yaml

crd:
	kubectl create -f deploy/crd.yaml

operator:
	-kubectl delete deploy puffin-operator
	kubectl create -f deploy/operator.yaml

