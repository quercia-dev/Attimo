<script lang="ts">
	import '../app.postcss';
	import { LightSwitch, Avatar } from '@skeletonlabs/skeleton';

	import AppRail from './AppRail.svelte';
	import AppBar from './AppBar.svelte';
	import Home from './home/+page.svelte';
	import Table from './table/+page.svelte';
	import Calendar from './calendar/+page.svelte';
	import SimplePane from '../components/SimplePane.svelte';
	import { tileState } from '../stores/appStore';
</script>

<div style="display: flex; flex-direction: column; height: 100vh; no-select">
	<AppBar />

	<div style="display: flex; flex: 1;">
		<AppRail />

		<div style="flex: 1; display: flex; flex-direction: column;">
			<main
				class="bg-surface-500 dark:bg-inherit p-4 flex flex-col h-[calc(100vh-8rem)] overflow-hidden no-select"
			>
				{#if $tileState === 0}
					<Home />
				{:else if $tileState === 1}
					<Calendar />
				{:else if $tileState == 4}
					<Table />
				{:else}
					<SimplePane>
						<p>Page not found</p>
					</SimplePane>
				{/if}
			</main>

			<footer
				class="dark:bg-surface-500"
				style="display: flex; justify-content: flex-end; align-items: center;"
			>
				<span style="margin-right: 10px;">Quercia development team</span>
				<a
					href="https://github.com/quercia-dev/Attimo"
					target="_blank"
					style="display: flex; align-items: center;"
				>
					<Avatar src="/icons/quercia-dev.png" width="w-10" rounded="rounded-l" />
				</a>
			</footer>
		</div>
	</div>
</div>

<style>
	.no-select {
		user-select: none;
	}
</style>
