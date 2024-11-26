<script>
	 import {  onDestroy } from 'svelte'
	 import {getInterface, getActiveIps} from "$lib/API.svelte"
	 import {inetName} from "$lib/store"
	 import { Button } from "carbon-components-svelte";

	 onDestroy(() => {
		$inetName = ''
	});

	let interfaceInfo =getInterface($inetName);
	let ronaldo;

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
			<Button size="small" iconDescription="Info" on:click={() =>   ronaldo = getActiveIps(users.Name, test.IP, 10, 10)}/>
		{/each}
	</div> 
{/await}


{#await ronaldo}
	<h2>Loading....</h2>
{:then users}
 	yey
{/await}