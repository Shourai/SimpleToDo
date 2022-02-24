let socket = new WebSocket("ws://localhost:8080/ws")
console.log("Attempting WS connection")

socket.onopen = () => {
    console.log("Succesfully connected to websocket")
    // socket.send("")
    // socket.send("Hello from client")
}

function checkExists(id) {
    let list = document.getElementById("ToDoList")
    return false
    // list.childNodes.forEach(node => {
    //     return id === node.attributes.taskid
    // })

}

socket.onmessage = (msg) => {
    console.log(msg.data)
    let taskList = JSON.parse(msg.data)
    let list = document.getElementById("ToDoList")

    taskList.forEach(el => {
        if (!checkExists(el.id)) {
            let li = document.createElement("li")
            li.textContent = el.name
            li.setAttribute("taskid", el.id)

            list.append(li)
        }
    });


    // array.forEach((item) => {
    //     let div = document.createElement("div")
    //     div.innerHTML = `
    //     <input type="checkbox" id=${item.id} name=${item.id}>
    //     <label for=${item.id}>${item.name}</label>
    //     `
    //     document.getElementById("tasks").appendChild(div)
    // }
    // )
}

socket.onclose = (ev) => {
    console.log("Closing socket", ev)
}

function addItem() {
    let li = document.createElement("li")
    let inputValue = document.getElementById("myInput").value
    let list = document.getElementById("ToDoList")

    if (inputValue === "") {
        console.log("Input field empty")
    } else {
        li.textContent = inputValue
        socket.send(inputValue)
        list.append(li)
        document.getElementById("myInput").value = ""
    }

}