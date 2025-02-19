<script lang="ts">
	import '../app.postcss';
	import { LightSwitch, Avatar } from '@skeletonlabs/skeleton';

	import AppRail from './AppRail.svelte';
	import AppBar from './AppBar.svelte';
	import Table from './data/+page.svelte';
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
				class="bg-surface-600 dark:bg-surface-900 p-4 flex flex-col h-full overflow-hidden no-select"
			>
				<slot />
			</main>

			<footer
				class="bg-primary-400 dark:bg-primary-900"
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
