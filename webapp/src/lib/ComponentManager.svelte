<script>
    import url from './url.svelte'
	import Nav from "$lib/Nav.svelte";
	import Table from "$lib/Table.svelte";
    import {page} from '$app/stores'
</script>


<Nav></Nav>

{#if $page.url.pathname === '' || $page.url.pathname === '/'}
    <Table></Table>
{:else}
<Table></Table>
<p>{ $page.url.pathname}</p>
{/if}