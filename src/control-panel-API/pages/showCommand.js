document.getElementById("form_console").addEventListener("click", (event) => {
	event.preventDefault()

	const command = document.getElementById("command");
	const show_command = document.getElementById("show_command");

	if(command.value === "") {
		show_command.innerHTML = "No hay ningun comando para enviar";
	} else {
		show_command.innerHTML = command.value;
	}
})