<script lang="ts">
	import '../app.postcss';
	import { AppShell, AppRail, AppRailTile, AppRailAnchor } from '@skeletonlabs/skeleton';

	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	let currentTile: number = 0;
</script>

<!-- App Shell -->
<AppShell>
	<!-- Sidebar -->
	 <AppRail>
		<AppRailTile bind:group={currentTile} name="home" value={0} title="home">
			<svelte:fragment slot="lead">
				<div style="display: flex; justify-content: center;"><img src="/home.svg" alt="Home" /></div>
			</svelte:fragment>
		</AppRailTile>
		<AppRailTile bind:group={currentTile} name="calendar" value={1} title="calendar">
			<svelte:fragment slot="lead">
				<div style="display: flex; justify-content: center;"> <img src="/calendar.svg" alt="Calendar" /> </div>
			</svelte:fragment>
		</AppRailTile>
		<AppRailTile bind:group={currentTile} name="open" value={2} title="open">
			<svelte:fragment slot="lead">
				<div style="display: flex; justify-content: center;"> <img src="/plus.svg" alt="Open" /> </div>
			</svelte:fragment>
		</AppRailTile>
		<AppRailTile bind:group={currentTile} name="close" value={3} title="close">
			<svelte:fragment slot="lead">
				<div style="display: flex; justify-content: center;"> <img src="/minus.svg" alt="Close" /> </div>
			</svelte:fragment>
		</AppRailTile>
		<!-- --- -->
		<svelte:fragment slot="trail">
			<AppRailAnchor href="/" target="_blank" title="SETTINGS">
				<div style="display: flex; justify-content: center;"> <img src="/settings.svg" alt="Settings" /> </div></AppRailAnchor>
		</svelte:fragment>
		</AppRail>
	<!-- Page Route Content -->
	<slot />
</AppShell>

