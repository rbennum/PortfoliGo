CSS_FN = style.css
CSS_SRC = ./static/
CSS_DEST = ./dist/css/

# JS_FN = script.js
# JS_SRC = ./src/js/script.js
# JS_DEST = ./dist/js/script.js

dev:
	@mkdir -p $(CSS_DEST)
	@cp $(CSS_SRC)$(CSS_FN) $(CSS_DEST)

debug-local: dev
	@dlv debug main.go --log

run-local: dev
	@go run main.go

minify-css:
	@cleancss -o $(CSS_DEST)$(CSS_FN) $(CSS_SRC)$(CSS_FN)

build-docker: apply-env minify-css
	@docker build -t rbennum2329/portfoligo:$(IMAGE_TAG) .
	@docker push rbennum2329/portfoligo:$(IMAGE_TAG)

run-docker: apply-env
	@docker pull rbennum2329/portfoligo:$(IMAGE_TAG)
	@docker container run --env-file prod.env -d \
	--name portfoligo -p $(APP_PORT):$(APP_PORT) \
	rbennum2329/portfoligo:$(IMAGE_TAG)

clean:
	@rm -rf ./dist/*
	@docker container stop portfoligo
	@docker container remove portfoligo

apply-env:
	@if [ -z "$(IMAGE_TAG)" ]; then \
		echo "ERROR: IMAGE_TAG is not set."; \
		exit 1; \
	fi