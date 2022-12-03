users = ./internal/controllers/http/users,./internal/controllers/http/users/dto
transactions = ./internal/controllers/http/transactions,./internal/controllers/http/transactions/dto
models = "./internal/controllers/http/shared/views"
controllers = $(models),$(users),$(transactions)

generateDocs:
	swag i -md ./docs/md/** -d ./cmd/http,$(controllers)  --propertyStrategy camelcase --requiredByDefault