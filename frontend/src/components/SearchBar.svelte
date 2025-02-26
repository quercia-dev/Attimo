<script lang="ts">
	export let value: string;
	
	function handleInput(event: Event) {
		const target = event.target as HTMLInputElement;
		value = target.value;
		console.log(value);
	}
</script>

<input
	id="search"
	type="search"
	placeholder="Dammi un Attimo..."
	bind:value={value}
	oninput={handleInput}
	class="search-bar bg-surface-100 dark:bg-surface-700"
/>

<style>
	.search-bar {
		margin-top: 0.25rem;
		display: block;
		width: 50%;
		border-radius: 0.375rem;
		box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
	}
</style>
