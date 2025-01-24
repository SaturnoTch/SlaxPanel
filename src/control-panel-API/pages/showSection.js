function showSection(section) {
	switch(section) {
		case "home":

			//localStorage
			localStorage.setItem("section", "home");
			//etiquetas <a></a>
			document.getElementById("console").removeAttribute("class");
			document.getElementById("home").setAttribute("class", "is-active");
			document.getElementById("disk").removeAttribute("class");
			document.getElementById("ram").removeAttribute("class");
			document.getElementById("cpu").removeAttribute("class");

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
			//etiquetas <a></a>
			document.getElementById("console").setAttribute("class", "is-active");
			document.getElementById("home").removeAttribute("class");
			document.getElementById("disk").removeAttribute("class");
			document.getElementById("ram").removeAttribute("class");
			document.getElementById("cpu").removeAttribute("class");

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
			//etiquetas <a></a>
			document.getElementById("console").removeAttribute("class");
			document.getElementById("home").removeAttribute("class");
			document.getElementById("disk").setAttribute("class", "is-active");
			document.getElementById("ram").removeAttribute("class");
			document.getElementById("cpu").removeAttribute("class");

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
			//etiquetas <a></a>
			document.getElementById("console").removeAttribute("class");
			document.getElementById("home").removeAttribute("class");
			document.getElementById("disk").removeAttribute("class");
			document.getElementById("ram").setAttribute("class", "is-active");
			document.getElementById("cpu").removeAttribute("class");

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
			//etiquetas <a></a>
			document.getElementById("console").removeAttribute("class");
			document.getElementById("home").removeAttribute("class");
			document.getElementById("disk").removeAttribute("class");
			document.getElementById("ram").removeAttribute("class");
			document.getElementById("cpu").setAttribute("class", "is-active");

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