PLUGIN_NAME=db
BINARY_NAME=trivy-db
PLUGIN_DIR=$(HOME)/.trivy/plugins/$(PLUGIN_NAME)

build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) main.go

clean:
	rm -rf $(PLUGIN_DIR)

install: build clean
	mkdir -p $(PLUGIN_DIR)
	cp $(BINARY_NAME) $(PLUGIN_DIR)/
	cp plugin.yaml $(PLUGIN_DIR)/
	chmod +x $(PLUGIN_DIR)/$(BINARY_NAME)
	@echo "✅ Plugin installé dans $(PLUGIN_DIR)"
	@ls -l $(PLUGIN_DIR)
	@if [ ! -x $(PLUGIN_DIR)/$(BINARY_NAME) ]; then \
	echo "❌ Le binaire n'est pas exécutable. Vérifie les permissions."; \
	exit 1; \
	fi

run:
	@if [ "$(word 2,$(MAKECMDGOALS))" = "" ]; then \
		echo "❌ Erreur : tu dois fournir une CVE en argument."; \
		echo "👉 Exemple : make run CVE-2025-27789"; \
		exit 1; \
	fi
	trivy $(PLUGIN_NAME) $(word 2,$(MAKECMDGOALS))

.PHONY: build clean install run