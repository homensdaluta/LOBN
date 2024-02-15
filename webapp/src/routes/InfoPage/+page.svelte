<script>
	 import {  onDestroy } from 'svelte'
	 import {getInterface} from "$lib/API.svelte"
	 import {inetName} from "$lib/store"

	 onDestroy(() => {
		$inetName = ''
	});

	let interfaceInfo =getInterface($inetName);


</script>

{#await interfaceInfo}
    <h2>Loading....</h2>
{:then users}
	<div>
		<p>{users.Name}</p>
		<p>{users.Description}</p>
		<p>{users.Flags}</p>
		<p>Addresses:</p>
		{#each users.Addresses as test}
			<br>
			<p>IP: {test.IP}</p>
			<p>NetMask: {test.NetMask}</p>
			<p>Broadcaster IP: {test.Broadaddr}</p>
			<p>P2P: {test.P2P}</p>
		{/each}
	</div> 
{/await}