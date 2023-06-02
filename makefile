run:
	go run cmd/tour/main.go -http=0.0.0.0:3000 -openbrowser=true
setup:
	go install github.com/cosmtrek/air@latest
run-with-reload:
	air
