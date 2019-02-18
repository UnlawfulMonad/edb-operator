
CRDS := $(wildcard deploy/crds/*.yaml)

all: build

build: container
	@echo "Generating deploy script"
	@cat $(CRDS) > deploy.yaml
	for f in namespace service_account role role_binding operator; do \
		echo "\n---" >> deploy.yaml; \
		cat deploy/$${f}.yaml >> deploy.yaml; \
	done

container:
	@operator-sdk build unlawfulmonad/edb-operator

.PHONY: all build
