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

<!--
	When adding a new element, ensure that the value attribute is set to +1 of the largest number so far,
	irrespective of where it is placed in the list. This ensures that each element has a unique and sequential value.
-->
<div class="app-rail-wrapper">
	{#if $showAppRail}
		<AppRail background="bg-primary-300 dark:bg-primary-800">
			<AppRailTile
				bind:group={currentTile}
				name="home"
				value={0}
				title="home"
				on:click={() => handleTileClick(0)}
			>
				{#snippet lead()}
					<div style="display: flex; justify-content: center;">
						<svg class="feather">
							<use href="/icons/feather-sprite.svg#home" />
						</svg>
					</div>
				{/snippet}
			</AppRailTile>
			<AppRailTile
				bind:group={currentTile}
				name="calendar"
				value={1}
				title="calendar"
				on:click={() => handleTileClick(1)}
			>
				{#snippet lead()}
					<div style="display: flex; justify-content: center;">
						<svg class="feather">
							<use href="/icons/feather-sprite.svg#calendar" />
						</svg>
					</div>
				{/snippet}
			</AppRailTile>
			<AppRailTile
				bind:group={currentTile}
				name="data"
				value={4}
				title="data"
				on:click={() => handleTileClick(4)}
			>
				<svelte:fragment slot="lead">
					<div style="display: flex; justify-content: center;">
						<svg class="feather">
							<use href="/icons/feather-sprite.svg#list" />
						</svg>
					</div>
				</svelte:fragment>
			</AppRailTile>
			<AppRailTile
				bind:group={currentTile}
				name="open"
				value={2}
				title="open"
				on:click={() => handleTileClick(2)}
			>
				{#snippet lead()}
					<div style="display: flex; justify-content: center;">
						<svg class="feather">
							<use href="/icons/feather-sprite.svg#plus" />
						</svg>
					</div>
				{/snippet}
			</AppRailTile>
			<AppRailTile
				bind:group={currentTile}
				name="close"
				value={3}
				title="close"
				on:click={() => handleTileClick(3)}
			>
				{#snippet lead()}
					<div style="display: flex; justify-content: center;">
						<svg class="feather">
							<use href="/icons/feather-sprite.svg#minus" />
						</svg>
					</div>
				{/snippet}
			</AppRailTile>
			{#snippet trail()}
				<AppRailTile
					bind:group={currentTile}
					name="settings"
					value={5}
					title="settings"
					on:click={() => handleTileClick(5)}
				>
					<div style="display: flex; justify-content: center;">
						<svg class="feather">
							<use href="/icons/feather-sprite.svg#settings" />
						</svg>
					</div>
				</AppRailTile>
			{/snippet}
		</AppRail>
	{/if}
</div>
