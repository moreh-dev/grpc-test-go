PROTO_PACKAGE_PREFIX := moreh/sdamanager/internal
PROTO_FILES := sdamanager_common.proto sdamanager_core.proto sdamanager_api.proto sdamanager_job.proto 

.PHONY: all $(PROTO_FILES) clean

all: clean $(PROTO_FILES)

$(PROTO_FILES):
	protoc \
		-I $(PWD) \
		--go_out=. \
		--go_opt=paths=import \
		--go_opt=module=moreh.io/go-grpc-test/internal/proto \
		--go-grpc_out=. \
		--go-grpc_opt=paths=import \
		--go-grpc_opt=module=moreh.io/go-grpc-test/internal/proto \
		--go-grpc_opt=require_unimplemented_servers=false \
		$(PWD)/$@

clean:
	@for file in $(shell find . -type f -name "*.pb.go"); do \
		echo "rm: $$file"; \
		rm -f $$file; \
	done
