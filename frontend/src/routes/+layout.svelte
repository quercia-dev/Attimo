<script lang="ts">
	import '../app.postcss';
	import { LightSwitch, Avatar } from '@skeletonlabs/skeleton';


	import AppRail from './AppRail.svelte';
	import AppBar from './AppBar.svelte';
	import Home from './home/+page.svelte';

	let currentTile: number = 0;
</script>

<div style="display: flex; flex-direction: column; height: 100vh;">
	<AppBar />

	<div style="display: flex; flex: 1;">
		<AppRail />

		<div style="flex: 1; display: flex; flex-direction: column;">
			<Home />

			<footer class="dark:bg-surface-500" style="display: flex; justify-content: flex-end; align-items: center;">
				<span style="margin-right: 10px;">Quercia development team</span>
				<a href="https://github.com/quercia-dev/Attimo" target="_blank" style="display: flex; align-items: center;">
					<Avatar src="/icons/quercia-dev.png" width="w-10" rounded="rounded-l" />

				</a>
			</footer>
		</div>
	</div>
</div>
