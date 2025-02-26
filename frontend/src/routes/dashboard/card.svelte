<script lang="ts">
	export let showChips: boolean = true;
	export let showEdit: boolean = true;
	export let chips: string[] = [];
</script>

<div class="card p-2 bg-surface-300 dark:bg-surface-600" style="border: 1px solid;">
	<header class="flex justify-between items-center">
		<div class="flex-1">
			{#if showChips}
				<div>
					{#each chips as chip, i}
						<span class="chip variant-filled ml-1">{chip}</span>
					{/each}
				</div>
			{/if}
		</div>

		{#if showEdit}
			<button
				onclick={() => console.log('Plus clicked')}
				aria-label="Open drawer"
				class="small-button"
			>
				<svg class="feather">
					<use href="/icons/feather-sprite.svg#edit" />
				</svg>
			</button>
		{/if}
	</header>

	<section class="p-2">(content)</section>
</div>
