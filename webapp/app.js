const http = require('http')
const os = require('os')

console.log("Nodejs server starting...")

var handler = function(request, response) {
    console.log("Received request from " + request.connection.remoteAdress);
    response.writeHead(200);
    response.end("You've hit " + os.hostname() + "\n");
};

var www = http.createServer(handler);
www.listen(8080);