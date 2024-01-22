# Read variables from config.json
config := $(shell type config.json | jq -r .)
protoc := $(shell echo $(config) | jq -r .protoc)
project_dir := $(shell echo $(config) | jq -r .project_dir)
proto_dir := $(shell echo $(config) | jq -r .proto_dir)
output_dir := $(shell echo $(config) | jq -r .output_dir)

generate:
	${protoc} --proto_path=${project_dir} --go_out=${output_dir} --go_opt=paths=source_relative ${proto_dir}/*.proto
	${protoc} --proto_path=${project_dir} --go-grpc_out=${output_dir} --go-grpc_opt=paths=source_relative ${proto_dir}/*.proto

clean:
	del /Q /S ${output_dir}

target: 
	echo "this is file 1 ${output_dir}" > file1.txt