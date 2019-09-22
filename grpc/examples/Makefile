.PHONY: simple, keepalive, streaming, errorhandling

simple: 
	export EXAMPLENAME=simple;\
	docker-compose --file client_server.compose up

keepalive: 
	export EXAMPLENAME=keepalive; \
	docker-compose --file client_server.compose up

streaming: 
	export EXAMPLENAME=streamingclient; \
	docker-compose --file client_server.compose up --build --force-recreate

errorhandling: 
	export EXAMPLENAME=errorhandling; \
	docker-compose --file client_server.compose up --build --force-recreate