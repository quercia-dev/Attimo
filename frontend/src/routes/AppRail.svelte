<script lang="ts">
	import '../app.postcss';
	import { AppRail, AppRailAnchor } from '@skeletonlabs/skeleton';
	import { page } from '$app/stores';
	import { fade, slide } from 'svelte/transition';
	import { showAppRail } from '$lib/stores/appStore';

	$: pane = $page.url.pathname;
</script>

<!--
	When adding a new element, ensure that the value attribute is set to +1 of the largest number so far,
	irrespective of where it is placed in the list. This ensures that each element has a unique and sequential value.
-->
<div class="app-rail-wrapper no-select">
	{#if $showAppRail}
		<AppRail background="bg-primary-300 dark:bg-primary-800">
			<AppRailAnchor href="/" title="home" selected={pane == '/'}>
				<div style="display: flex; justify-content: center;">
					<svg class="feather">
						<use href="/icons/feather-sprite.svg#home" />
					</svg>
				</div>

				<span>Home</span>
			</AppRailAnchor>
			<AppRailAnchor href="/calendar" title="calendar" selected={pane == '/calendar'}>
				<div style="display: flex; justify-content: center;">
					<svg class="feather">
						<use href="/icons/feather-sprite.svg#calendar" />
					</svg>
				</div>

				<span>Calendar</span>
			</AppRailAnchor>
			<AppRailAnchor href="/data" title="data" selected={pane == '/data'}>
				<div style="display: flex; justify-content: center;">
					<svg class="feather">
						<use href="/icons/feather-sprite.svg#list" />
					</svg>
				</div>

				<span>Data</span>
			</AppRailAnchor>
			<AppRailAnchor href="/dashboard" title="data" selected={pane == '/dashboard'}>
				<div style="display: flex; justify-content: center;">
					<svg class="feather">
						<use href="/icons/feather-sprite.svg#bar-chart-2" />
					</svg>
				</div>

				<span>Dashboard</span>
			</AppRailAnchor>
			<AppRailAnchor href="/open" title="open" selected={pane == '/open'}>
				<div style="display: flex; justify-content: center;">
					<svg class="feather">
						<use href="/icons/feather-sprite.svg#plus" />
					</svg>
				</div>

				<span>Open</span>
			</AppRailAnchor>
			<AppRailAnchor href="/close" title="close" selected={pane == '/close'}>
				<div style="display: flex; justify-content: center;">
					<svg class="feather">
						<use href="/icons/feather-sprite.svg#minus" />
					</svg>
				</div>
				<span>Close</span>
			</AppRailAnchor>

			<svelte:fragment slot="trail">
				<AppRailAnchor
					href="/settings"
					title="settings"
					selected={$page.url.pathname == '/settings'}
				>
					<div style="display: flex; justify-content: center;">
						<svg class="feather">
							<use href="/icons/feather-sprite.svg#settings" />
						</svg>
					</div>
					<span>Settings</span>
				</AppRailAnchor>
			</svelte:fragment>
		</AppRail>
	{/if}
</div>
