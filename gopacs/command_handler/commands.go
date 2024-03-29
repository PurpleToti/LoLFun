package command_handler

func helloCommand() jsonCommandResponse {
	return jsonCommandResponse{
		to_display: "Hello World!",
		details:    "Just a classic hello world :)",
	}
}
