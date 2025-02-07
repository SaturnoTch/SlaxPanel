function showSection(section) {
	switch(section) {
		case "home":

			//localStorage
			localStorage.setItem("section", "home");
			

			//sections
			document.getElementById("console_section").setAttribute("style", "display: none;");
			document.getElementById("home_section").setAttribute("style", "display: block;");
			document.getElementById("disk_section").setAttribute("style", "display: none;");
			document.getElementById("ram_section").setAttribute("style", "display: none;");
			document.getElementById("cpu_section").setAttribute("style", "display: none;");
			console.log("El usuario ha elegido ver el inicio");
			break;
		case "console":
			//localStorage
			localStorage.setItem("section", "console");
			

			//sections
			document.getElementById("console_section").setAttribute("style", "display: block;");
			document.getElementById("home_section").setAttribute("style", "display: none;");
			document.getElementById("disk_section").setAttribute("style", "display: none;");
			document.getElementById("ram_section").setAttribute("style", "display: none;");
			document.getElementById("cpu_section").setAttribute("style", "display: none;");
			console.log("El usuario ha elegido ver la consola");
			break;
		case "disk":
			//localStorage
			localStorage.setItem("section", "disk");
			

			//sections
			document.getElementById("console_section").setAttribute("style", "display: none;");
			document.getElementById("home_section").setAttribute("style", "display: none;");
			document.getElementById("disk_section").setAttribute("style", "display: block;");
			document.getElementById("ram_section").setAttribute("style", "display: none;");
			document.getElementById("cpu_section").setAttribute("style", "display: none;");
			console.log("El usuario ha elegido ver el almacenamiento");
			break;
		case "ram":
			//localStorage
			localStorage.setItem("section", "ram");
			

			//sections
			document.getElementById("console_section").setAttribute("style", "display: none;");
			document.getElementById("home_section").setAttribute("style", "display: none;");
			document.getElementById("disk_section").setAttribute("style", "display: none;");
			document.getElementById("ram_section").setAttribute("style", "display: block;");
			document.getElementById("cpu_section").setAttribute("style", "display: none;");
			console.log("El usuario ha elegido ver la RAM");
			break;
		case "cpu":
			//localStorage
			localStorage.setItem("section", "cpu");
			

			//sections
			document.getElementById("console_section").setAttribute("style", "display: none;");
			document.getElementById("home_section").setAttribute("style", "display: none;");
			document.getElementById("disk_section").setAttribute("style", "display: none;");
			document.getElementById("ram_section").setAttribute("style", "display: none;");
			document.getElementById("cpu_section").setAttribute("style", "display: block;");
			console.log("El usuario ha elegido ver el CPU");
	}
}

document.getElementById("home").addEventListener("click", (event) => {
	event.preventDefault();

	showSection("home");
})

document.getElementById("console").addEventListener("click", (event) => {
	event.preventDefault();

	showSection("console");
})

document.getElementById("disk").addEventListener("click", (event) => {
	event.preventDefault();

	showSection("disk");
})

document.getElementById("ram").addEventListener("click", (event) => {
	event.preventDefault();

	showSection("ram");
})

document.getElementById("cpu").addEventListener("click", (event) => {
	event.preventDefault();

	showSection("cpu");
})

const section = localStorage.getItem("section");

switch(section) {
case "home":
	showSection("home");
	break;
case "console":
	showSection("console");
	break;
case "disk":
	showSection("disk");
	break;
case "ram":
	showSection("ram");
	break;
case "cpu":
	showSection("cpu");
}