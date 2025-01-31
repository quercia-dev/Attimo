<script lang="ts">
	import '../app.postcss';

	import { AppBar } from '@skeletonlabs/skeleton';

	import AppRail from './AppRail.svelte';

	import Home from './Home.svelte';

	let currentTile: number = 0;
</script>

<div style="display: flex; height: 100vh;">
	<AppRail />

	<div style="flex: 1; display: flex; flex-direction: column;">
		<!-- AppBar -->
		<AppBar style="grid-template-columns: 1fr 3fr;">
			<svelte:fragment slot="lead">
				<button
					type="button"
					on:click={() => console.log('click hamburger')}
					aria-label="Open drawer"
					class="icon-button"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						width="32"
						height="32"
						viewBox="0 0 24 24"
						fill="none"
						stroke="currentColor"
						stroke-width="3"
						stroke-linecap="round"
						stroke-linejoin="round"
						class="feather feather-menu"
					>
						<line x1="3" y1="12" x2="21" y2="12"></line>
						<line x1="3" y1="6" x2="21" y2="6"></line>
						<line x1="3" y1="18" x2="21" y2="18"></line>
					</svg>
				</button>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<button
					type="button"
					on:click={() => console.log('click hamburger')}
					aria-label="Open drawer"
					class="icon-button"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						width="32"
						height="32"
						viewBox="0 0 24 24"
						fill="none"
						stroke="currentColor"
						stroke-width="3"
						stroke-linecap="round"
						stroke-linejoin="round"
						class="feather feather-message-square"
					>
						<path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path>
					</svg>
				</button>
			</svelte:fragment>
		</AppBar>

		<!-- Home Pane -->
		<Home />

		<footer class="bg-blue-500 p-1">footer</footer>
	</div>
</div>
