build:
	protoc pb/*.proto --go_out=plugins=grpc:. --js_out=client/js --python_out=client/py
