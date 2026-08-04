build-chat:
	@go build -o ./bin/chat ./...

chat: build-chat
	@./bin/chat

clean:
	@rm -rf bin