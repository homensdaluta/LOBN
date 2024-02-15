<script>
	import {
	  Header,
	  HeaderNav,
	  HeaderNavItem,
	  HeaderNavMenu,
	  SideNav,
	  SideNavItems,
	  SideNavMenu,
	  SideNavMenuItem,
	  SideNavLink,
	  SideNavDivider,
	  SkipToContent,
	  Content,
	  Grid,
	  Row,
	  Column,
	} from "carbon-components-svelte";
	let isSideNavOpen = false;
</script>

<svelte:head><link rel="stylesheet" href="https://unpkg.com/carbon-components-svelte@0.30.0/css/g10.css" /></svelte:head>
<div>
	<Header company="LOBN" platformName="Lack of better name" bind:isSideNavOpen>
		<svelte:fragment slot="skip-to-content">
			<SkipToContent />
		</svelte:fragment>
		<HeaderNav>
			<HeaderNavItem href="/" text="Link 1" />
			<HeaderNavItem href="/" text="Link 2" />
			<HeaderNavItem href="/" text="Link 3" />
		</HeaderNav>
	</Header>
</div>
  