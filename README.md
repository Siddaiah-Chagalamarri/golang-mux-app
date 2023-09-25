# golang-mux-app
1. Create a directory for golang-mux-app.
	mkdir golang-mux-app
	cd golang-mux-app
2. Run the following command to initialize Go modules.
	go mod init golang-mux-app
3. Run the following command to install the Gorilla Mux library.
	go get github.com/gorilla/mux
4. Create the main.go and Dockerfile.
5. To build your Go application, run the following command.
	go build -o main .
6. Open the command prompt and navigate to your project directory. Run the following command to build a Docker image.
	docker build -t golang-mux-app .
7. After the image is built, you can run a Docker container based on the image.
	docker run -p 8080:8080 golang-mux-app
8. You can use a tool like curl or Postman to test your /api/encrypt and /api/decrypt APIs.
	curl -X POST -H "Content-Type: application/json" -d '{"text":"Hello, World!"}' http://localhost:8080/api/encrypt





