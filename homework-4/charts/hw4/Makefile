.PHONY: install
install:
	helm upgrade --install hw4-release . \
      --dependency-update \
      --create-namespace \
      --namespace=hw4

.PHONY: upgrade
upgrade:
	helm upgrade hw4-release . \
      --dependency-update

.PHONY: uninstall
uninstall:
	helm uninstall hw4-release
