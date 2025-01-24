document.getElementById("form_system").addEventListener("click", (event) => {
	event.preventDefault()

	const command = document.getElementById("timer");
	const show_command = document.getElementById("show_timer");

	if(command.value === "") {
		show_command.innerHTML = "No hay ninguna programacion para el reinicio de la maquina";
	} else {
		show_command.innerHTML = command.value;
	}
})