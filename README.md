# Golang_Templates_Parsing
This code demonstrates a basic way to serve an HTML file in Go using a template
## maingo
This code defines a simple HTTP server that serves a single HTML file (index.html) using the html/template package. The indexHandler function is responsible for rendering the template and sending the HTML response to the client. If there's an error during template execution, it sends a 500 Internal Server Error response.

## module
Initialize Go modules by Run the following command to initialize Go modules for your project:

bash
Copy code
go mod init github.com/Boukabouya/Golang_Templates_Parsing

Add dependencies:

bash
Copy code
go get github.com/gorilla/mux