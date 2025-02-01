<script lang="ts">
	import '../app.postcss';
	import { LightSwitch } from '@skeletonlabs/skeleton';

	import AppRail from './AppRail.svelte';
	import AppBar from './AppBar.svelte';
	import Home from './Home.svelte';

	let currentTile: number = 0;
</script>

<div style="display: flex; height: 100vh;">
	<AppRail />

	<div style="flex: 1; display: flex; flex-direction: column;">
		<AppBar />

		<Home />

		<footer class="p-1" style="display: flex; justify-content: flex-end; align-items: center;">
			<span style="margin-right: 10px;">Quercia development team</span>
			<a href="https://github.com/quercia-dev/Attimo" target="_blank" style="display: flex; align-items: center;">
				<img src="/icons/quercia-dev.png" alt="Quercia Dev Logo" style="max-width: 100%; height: auto;" width="30" height="30" />
			</a>
		</footer>
	</div>
</div>
