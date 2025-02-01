<script lang="ts">
	import '../app.postcss';
	import { AppRail, AppRailTile, AppRailAnchor } from '@skeletonlabs/skeleton';
	import { fade, slide } from 'svelte/transition';
	import { showAppRail } from '../stores/appStore';
		
	let currentTile: number = 0;
</script>

<div transition:fade={{ duration: 300 }} class="app-rail-wrapper">
	{#if $showAppRail}
	<AppRail>
		<AppRailTile
			bind:group={currentTile}
			name="home"
			value={0}
			title="home"
		>
			<svelte:fragment slot="lead">
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
			</svelte:fragment>
		</AppRailTile>
		<AppRailTile bind:group={currentTile} name="calendar" value={1} title="calendar">
			<svelte:fragment slot="lead">
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
						></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="24" x2="21" y2="24"
						></line></svg
					>
				</div>
			</svelte:fragment>
		</AppRailTile>
		<AppRailTile bind:group={currentTile} name="open" value={2} title="open">
			<svelte:fragment slot="lead">
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
			</svelte:fragment>
		</AppRailTile>
		<AppRailTile bind:group={currentTile} name="close" value={3} title="close">
			<svelte:fragment slot="lead">
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
			</svelte:fragment>
		</AppRailTile>
		<svelte:fragment slot="trail">
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
		</svelte:fragment>
	</AppRail>
	{/if}
</div>
