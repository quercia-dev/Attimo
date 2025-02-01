<script lang="ts">
	import '../app.postcss';

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

		<footer class="bg-blue-500 p-1">footer</footer>
	</div>
</div>
