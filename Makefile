image:
	docker build . -t buildso
	touch image

simple.so:image config.go filter.go
	docker run --rm -it -v ./:/source buildso

run:simple.so
	docker run -it --rm -e ENVOY_UID=0 --network host -v ./envoy.yaml:/envoy.yaml -v ./simple.so:/simple.so envoyproxy/envoy:contrib-v1.30.1 -c /envoy.yaml