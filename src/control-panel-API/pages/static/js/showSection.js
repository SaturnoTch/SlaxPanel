function showSection(section) {
	switch(section) {
		case "home":

			//localStorage and title
			localStorage.setItem("section", "home");
			document.querySelector("title").innerHTML = `Inicio - Dashboard SlaxPanel`;
			

			//sections
			document.getElementById("tools_section").setAttribute("style", "display: none;");
			document.getElementById("console_section").setAttribute("style", "display: none;");
			document.getElementById("home_section").setAttribute("style", "display: block;");
			document.getElementById("disk_section").setAttribute("style", "display: none;");
			document.getElementById("ram_section").setAttribute("style", "display: none;");
			document.getElementById("cpu_section").setAttribute("style", "display: none;");
			console.log("El usuario ha elegido ver el inicio");
			break;

		case "tools":
			//localStorage and title
			localStorage.setItem("section", "tools");
			document.querySelector("title").innerHTML = `Herramientas - Dashboard SlaxPanel`;
			

			//sections
			document.getElementById("tools_section").setAttribute("style", "display: block;");
			document.getElementById("console_section").setAttribute("style", "display: none;");
			document.getElementById("home_section").setAttribute("style", "display: none;");
			document.getElementById("disk_section").setAttribute("style", "display: none;");
			document.getElementById("ram_section").setAttribute("style", "display: none;");
			document.getElementById("cpu_section").setAttribute("style", "display: none;");
			console.log("El usuario ha elegido ver las herramientas");

			break;

		case "console":
			//localStorage and title
			localStorage.setItem("section", "console");
			document.querySelector("title").innerHTML = `Consola - Dashboard SlaxPanel`;

			//sections
			document.getElementById("tools_section").setAttribute("style", "display: none;");
			document.getElementById("console_section").setAttribute("style", "display: block;");
			document.getElementById("home_section").setAttribute("style", "display: none;");
			document.getElementById("disk_section").setAttribute("style", "display: none;");
			document.getElementById("ram_section").setAttribute("style", "display: none;");
			document.getElementById("cpu_section").setAttribute("style", "display: none;");
			console.log("El usuario ha elegido ver la consola");
			break;
		case "disk":
			//localStorage and title
			localStorage.setItem("section", "disk");
			document.querySelector("title").innerHTML = `Almacenamiento - Dashboard SlaxPanel`;
			

			//sections
			document.getElementById("tools_section").setAttribute("style", "display: none;");
			document.getElementById("console_section").setAttribute("style", "display: none;");
			document.getElementById("home_section").setAttribute("style", "display: none;");
			document.getElementById("disk_section").setAttribute("style", "display: block;");
			document.getElementById("ram_section").setAttribute("style", "display: none;");
			document.getElementById("cpu_section").setAttribute("style", "display: none;");
			console.log("El usuario ha elegido ver el almacenamiento");
			break;
		case "ram":
			//localStorage and title
			localStorage.setItem("section", "ram");
			document.querySelector("title").innerHTML = `RAM - Dashboard SlaxPanel`;
			

			//sections
			document.getElementById("tools_section").setAttribute("style", "display: none;");
			document.getElementById("console_section").setAttribute("style", "display: none;");
			document.getElementById("home_section").setAttribute("style", "display: none;");
			document.getElementById("disk_section").setAttribute("style", "display: none;");
			document.getElementById("ram_section").setAttribute("style", "display: block;");
			document.getElementById("cpu_section").setAttribute("style", "display: none;");
			console.log("El usuario ha elegido ver la RAM");
			break;
		case "cpu":
			//localStorage and title
			localStorage.setItem("section", "cpu");
			document.querySelector("title").innerHTML = `CPU - Dashboard SlaxPanel`;
			

			//sections
			document.getElementById("tools_section").setAttribute("style", "display: none;");
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

document.getElementById("tools").addEventListener("click", (event) => {
	event.preventDefault();

	showSection("tools");
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
case "tools":
	showSection("tools");
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