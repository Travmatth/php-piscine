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
        addTodo(anchor)(todo)
    }
}

const saveHandler = () => {
    anchor = document.getElementById("ft_list")
    elements = Array.prototype.slice.call(anchor.children)
    todos = elements.reduce((arr, element) => {
        arr.unshift(element.textContent)
        return arr
    }, [])
    value = JSON.stringify(todos)
    cookie = "ft_list=" + value + ";"
    document.cookie = cookie

}

const registerHandler = () => {
    anchor = document.getElementById("ft_list")
    newButton = document.getElementById("new")
    newButton.addEventListener("click", todoHandler(anchor))
    if (document.cookie === undefined) {
        return
    }
    matches = document.cookie.match(new RegExp("^ft_list=(.*)$"));
    cookie = matches ? JSON.parse(matches[1]) : undefined;
    if (cookie === undefined) {
        return
    }
    addFunc = addTodo(anchor)
    cookie.map(todo => addFunc(todo))
}

if (typeof document !== "undefined") {
    document.addEventListener("DOMContentLoaded", registerHandler)
    window.addEventListener("beforeunload", saveHandler)
}