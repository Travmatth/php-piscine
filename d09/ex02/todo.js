const removeHandler = e => {
    node = e.target || e.srcElement
    node.remove()
}

const addTodo = anchor => todo => {
    next = document.createElement("div")
    next.textContent = todo
    next.addEventListener("click", removeHandler)
    anchor.prepend(next)
}

const todoHandler = anchor => () => {
    todo = window.prompt("Enter your todo:")
    if (todo && todo !== "") {
        console.log("todo ", todo)
        addTodo(anchor)(todo)
    }
}

const saveHandler = () => {
    console.log("saving ", document.cookie)
    anchor = document.getElementById("ft_list")
    todos = anchor.children
    cookie = []
    for (let i = 0; i < todos.length; i++) {
        cookie.unshift(todos[i].textContent)
    }
    document.cookie = JSON.stringify(cookie)
}

const registerHandler = () => {
    anchor = document.getElementById("ft_list")
    newButton = document.getElementById("new")
    newButton.addEventListener("click", todoHandler(anchor))
    serializedCookie = document.cookie
    if (serializedCookie) {
        cookie = JSON.parse(serializedCookie)
        console.log("loaded ", cookie)
        cookie.forEach(addTodo(anchor))
    }
}

if (typeof document !== "undefined") {
    document.addEventListener("DOMContentLoaded", registerHandler)
    document.addEventListener("onunload", saveHandler)
}