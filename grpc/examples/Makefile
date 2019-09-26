.PHONY: simple, keepalive, streaming, errorhandling, deadlinepropagation, cancellation, metadata

simple: 
	export EXAMPLENAME=01_simple;\
	docker-compose --file client_server.compose up --build --force-recreate 

keepalive: 
	export EXAMPLENAME=02_keepalive; \
	docker-compose --file client_server.compose up --build --force-recreate 

streaming: 
	export EXAMPLENAME=03_streamingclient; \
	docker-compose --file client_server.compose up --build --force-recreate

errorhandling: 
	export EXAMPLENAME=04_errorhandling; \
	docker-compose --file client_server.compose up --build --force-recreate

deadlinepropagation: 
	export EXAMPLENAME=06_deadlinepropagation; \
	docker-compose --file client_middle_server.compose up --build --force-recreate


cancellation: 
	export EXAMPLENAME=05_cancellation; \
	docker-compose --file client_middle_server.compose up --build --force-recreate

metadata:
	export EXAMPLENAME=07_metadata; \
	docker-compose --file client_server.compose up --build --force-recreate

retry_interceptor:
	export EXAMPLENAME=08_retry_interceptor; \
	docker-compose --file client_server.compose up --build --force-recreate

opentracing:
	export EXAMPLENAME=09_opentracing; \
	docker-compose --file opentracing.compose up --build --force-recreate

