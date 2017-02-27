// Create websocket connection.
var ws = new WebSocket("ws://localhost:9000/ws");

ws.onopen = function() {
	ws.send(JSON.stringify({
		message: "hello server!"
	}));
};

ws.onmessage = function(event) {
	var m = JSON.parse(event.data);
	console.debug("Received message", m.message);
};

ws.onerror = function(event) {
	console.debug(event);
};