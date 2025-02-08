
cs-i:
	helm upgrade --install currency-service  ./currency-service

cs-dry:
	helm upgrade --install currency-service  ./currency-service --dry-run

cs-rm:
	helm uninstall  currency-service

aps-i:
	helm upgrade --install asset-price-service  ./asset-price-service

aps-dry:
	helm upgrade --install asset-price-service  ./asset-price-service --dry-run

aps-rm:
	helm uninstall  asset-price-service

ts-i:
	helm upgrade --install transaction-service  ./transaction-service

ts-dry:
	helm upgrade --install transaction-service  ./transaction-service --dry-run

ts-rm:
	helm uninstall  transaction-service

.PHONY: cs-i cs-dry cs-rm  aps-i aps-dry aps-rm  ts-i ts-dry ts-rm

