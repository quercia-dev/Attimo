<script lang="ts">
	import '../app.postcss';
	import { AppRail, AppRailTile, AppRailAnchor } from '@skeletonlabs/skeleton';
	import { fade, slide } from 'svelte/transition';
	import { tileState, showAppRail } from '../stores/appStore';

	let currentTile: number = $state(0);

	function handleTileClick(value: number) {
		tileState.set(value);
	}
</script>

<div class="app-rail-wrapper">
	{#if $showAppRail}
		<AppRail>
			<AppRailTile bind:group={currentTile} name="home" value={0} title="home" on:click={() => handleTileClick(0)}>
				{#snippet lead()}
							
						<div style="display: flex; justify-content: center;">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								class="feather feather-home"
								><path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path><polyline
									points="9 22 9 12 15 12 15 22"
								></polyline></svg
							>
						</div>
					
							{/snippet}
			</AppRailTile>
			<AppRailTile bind:group={currentTile} name="calendar" value={1} title="calendar" on:click={() => handleTileClick(1)}>
				{#snippet lead()}
							
						<div style="display: flex; justify-content: center;">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								class="feather feather-calendar"
								><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line
									x1="16"
									y1="2"
									x2="16"
									y2="6"
								></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"
								></line></svg
							>
						</div>
					
							{/snippet}
			</AppRailTile>
			<AppRailTile bind:group={currentTile} name="open" value={2} title="open" on:click={() => handleTileClick(2)}>
				{#snippet lead()}
							
						<div style="display: flex; justify-content: center;">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								class="feather feather-plus"
								><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"
								></line></svg
							>
						</div>
					
							{/snippet}
			</AppRailTile>
			<AppRailTile bind:group={currentTile} name="close" value={3} title="close" on:click={() => handleTileClick(3)}>
				{#snippet lead()}
							
						<div style="display: flex; justify-content: center;">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								class="feather feather-minus"><line x1="5" y1="12" x2="19" y2="12"></line></svg
							>
						</div>
					
							{/snippet}
			</AppRailTile>
			{#snippet trail()}
					
					<AppRailAnchor href="/" target="_blank" title="SETTINGS">
						<div style="display: flex; justify-content: center;">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								class="feather feather-settings"
								><circle cx="12" cy="12" r="3"></circle><path
									d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"
								></path></svg
							>
						</div>
					</AppRailAnchor>
				
					{/snippet}
		</AppRail>
	{/if}
</div>
