let socket = new WebSocket("ws://localhost:8080/ws")
console.log("Attempting WS connection")

socket.onopen = () => {
    console.log("Succesfully connected to websocket")
    socket.send("Hello from client")
}


socket.onmessage = (msg) => {
    console.log(msg.data)
    let array = JSON.parse(msg.data)

    array.forEach((item) => {
        let div = document.createElement("div")
        div.innerHTML = `
        <input type="checkbox" id=${item.id} name=${item.id}>
        <label for=${item.id}>${item.name}</label>
        `
        document.getElementById("tasks").appendChild(div)
    }
    )
}

socket.onclose = (ev) => {
    console.log("Closing socket", ev)
}