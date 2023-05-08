<script>
	import 'carbon-components-svelte/css/white.css';
	import {
		Header,
		SkipToContent,
		HeaderUtilities,
		Content
	} from 'carbon-components-svelte/src/UIShell';

	let isSideNavOpen = false;
</script>

<Header company="MJ" platformName="Špajza" bind:isSideNavOpen>
	<svelte:fragment slot="skip-to-content">
		<SkipToContent />
	</svelte:fragment>
	<HeaderUtilities />
</Header>

<Content>
	<slot />
</Content>
