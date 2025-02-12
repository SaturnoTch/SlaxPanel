const button = document.getElementById("trash");

button.addEventListener("click", async (event) => {
	event.preventDefault();

	const request = await axios.delete("http://localhost:8080/home/dashboard/tools/cleanup");

	
	alert(`${request.data.cleanupmessage}`)
})