<script lang="ts">
	import '../app.postcss';
	import { LightSwitch, Avatar } from '@skeletonlabs/skeleton';

	import AppRail from './AppRail.svelte';
	import AppBar from './AppBar.svelte';
</script>

<div style="no-select display: flex; flex-direction: column; height: 100vh;">
	<AppBar />

	<div style="display: flex; flex: 1;">
		<AppRail />

		<div style="flex: 1; display: flex; flex-direction: column;">
			<main class="bg-primary-500 p-4 flex flex-col h-[calc(100vh-6rem)] overflow-hidden">
				<slot />
			</main>

			<footer
				class="bg-primary-400 dark:bg-primary-900"
				style="display: flex; justify-content: flex-end; align-items: center;"
			>
				<span style="margin-right: 1rem;">Quercia development team</span>
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
