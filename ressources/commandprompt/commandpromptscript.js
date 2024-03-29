var no_reset_cursor_keys = [
    "ArrowUp",
    "ArrowRight",
    "ArrowDown",
    "ArrowLeft",
    "Alt",
    "Control",
    "Shift",
]

var return_char = '\n'
var command_begin_char = '>'
var new_line_string = "\n>"

function main() {
    setupWindow()
    setupInteractivePrompt()
}

function setupWindow() {
    let command_prompt = getTextBox()
    window.addEventListener("keydown", function(event) {
        if (event.key === "Backspace") {
            if (beforeIsNewCommandLine(command_prompt)) {
                addString(command_prompt, command_begin_char)
            }
        }

        if (!no_reset_cursor_keys.includes(event.key)) {
            putCursorAtEnd(command_prompt)
        }
    })
}

function getTextBox() {
    return document.getElementById("command_prompt_input_area")
}

function getLastLine(command_prompt) {
    let command_prompt_input = command_prompt.value
    let last_line = ""
    for (var i = command_prompt_input.length - 2; i >= 0; i--) {
        let curr_char = command_prompt_input.charAt(i)
        if (curr_char === command_begin_char) {
            break
        }
        last_line = command_prompt_input.charAt(i) + last_line
    }
    if (last_line === "") {
        return "null"
    }
    return last_line
}

function putCursorAtEnd(command_prompt) {
    command_prompt.setSelectionRange(command_prompt.value.length, command_prompt.value.length)
}

function addString(command_prompt, string) {
    if (string != null) {
        let command_prompt_input = command_prompt.value
        command_prompt.value = command_prompt_input + string
    }
}

function getLastChar(string) {
    return string.charAt(string.length - 1)
}

function getStringBefore(string, len) {
    let substring = ""
    for (let i = string.length - 1; i > string.length - 1 - len; i--) {
        substring = string.charAt(i) + substring
    }
    return substring
}

function newCommandLine(command_prompt) {
    let last_char = getLastChar(command_prompt.value)
    let new_line_char = ""
    if (last_char != return_char) {
        new_line_char += "\n"
    }
    addString(command_prompt, new_line_char + command_begin_char)
    putCursorAtEnd(command_prompt)
}

function beforeIsNewCommandLine(command_prompt) {
    let command_prompt_input = command_prompt.value
    let possible_new_newline_string = getStringBefore(command_prompt_input, new_line_string.length)
    return possible_new_newline_string === new_line_string
}

function setupInteractivePrompt() {
    let command_prompt = getTextBox()
    newCommandLine(command_prompt)
    command_prompt.addEventListener("input", function(event) {
        let last_char = getLastChar(command_prompt.value)

        if (last_char == return_char) {
            let command_string = getLastLine(command_prompt)
            executePromptCommand(command_prompt, "/handlecommand/" + command_string)
        }
    })
}

function executePromptCommand(command_prompt, url) {
    command_prompt.readonly = true
    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(url + "-> Server response was not ok")
            }
            return response.text();
        })
        .then(text => {
            console.log(text)
            data = JSON.parse(text)
            console.log(url + "-> Response:", data)
            command_response = data.to_display
            addString(command_prompt, command_response)
            newCommandLine(command_prompt)
            command_prompt.readonly = false
        })
        .catch(error => console.error(url + "-> Error:", error))
}

window.onload = function(event) {
    main()
}