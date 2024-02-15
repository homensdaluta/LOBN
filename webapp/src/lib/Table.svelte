<script>
    import {
      StructuredList,
      StructuredListHead,
      StructuredListRow,
      StructuredListCell,
      StructuredListBody,
      Button,
      ButtonSet
    } from "carbon-components-svelte";
    import Play from "carbon-icons-svelte/lib/Play.svelte";
    import Information from "carbon-icons-svelte/lib/Information.svelte";
    import {getAllInterfaces} from "./API.svelte"
    import {inetName} from "./store"
	let userPromise =getAllInterfaces();

  function navigateToInfo(nsda){
    $inetName = nsda;
  }

  </script>

<svelte:head><link rel="stylesheet" href="https://unpkg.com/carbon-components-svelte@0.30.0/css/g10.css" /></svelte:head>
<div>
  <StructuredList>
    <StructuredListHead>
      <StructuredListRow head>
        <StructuredListCell head>Name</StructuredListCell>
        <StructuredListCell head>Description</StructuredListCell>
        <StructuredListCell head>Flags</StructuredListCell>
        <StructuredListCell head>Addresses</StructuredListCell>
      </StructuredListRow>
    </StructuredListHead>
    <StructuredListBody>
        {#await userPromise}
        <h2>Loading....</h2>
        {:then users}
            {#each users as network}
                <StructuredListRow>
                    <StructuredListCell noWrap>{network.Name}</StructuredListCell>
                    <StructuredListCell >{network.Description}</StructuredListCell>
                    <StructuredListCell>{network.Flags}</StructuredListCell>
                    <StructuredListCell>
                        {#each network.Addresses as test}
                            <p>IP: {test.IP}</p>
                        {/each}                
                    </StructuredListCell>
                    <StructuredListCell>
                        <Button size="small" iconDescription="Info" icon={Information} href="/InfoPage" on:click={() =>   navigateToInfo(network.Name)}/>
                    </StructuredListCell>
                </StructuredListRow>
            {/each}
      {:catch err}
      <h2>Error while loading the data</h2>
    {/await}
    </StructuredListBody>
  </StructuredList>
  <p>{$inetName}</p>
</div>