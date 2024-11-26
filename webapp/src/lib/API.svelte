<script lang="ts" context="module">

	export const getAllInterfaces = async () => {
		var response = await fetch('http://localhost:8080/getSystemInfo');
		var result = await response.json();
		return result;
	}

	export const getInterface = async (name) => {
		var uriName = encodeURIComponent(name)
		var response = await fetch('http://localhost:8080/getInterfaceInfo/'+uriName);
		var result = await response.json();
		return result;  
	}

	export const getActiveIps = async (name, ip, timeout, timeBetweenPackets) => {
		var response = await fetch('http://localhost:8080/getAllActiveIps', {
			method: 'POST',
			body: JSON.stringify({
				name,
				ip,
				timeout,
				timeBetweenPackets
			})
		});
		console.log(JSON.stringify({
				name,
				ip,
				timeout,
				timeBetweenPackets
			}))
		var result = await response.json();
		return result;  
	}
</script>
