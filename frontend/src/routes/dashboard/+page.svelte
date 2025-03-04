<script lang="ts">
	import SimplePane from '../../components/SimplePane.svelte';
	import SearchBar from '../../components/SearchBar.svelte';
    import SecondaryButton from '../../components/SecondaryButton.svelte';
	import Card from './card.svelte';
	import { writable } from 'svelte/store';

	let searchQuery = $state('');

	type Option = {
		value: boolean;
		description: string;
	};

	let options = writable<Record<string, Option>>({
		showChips: { value: true, description: 'Sources' },
		showEdit: { value: true, description: 'Edit' }
	});

	function toggle(key: string): void {
		options.update((opts) => {
			opts[key].value = !opts[key].value;
			return opts;
		});
	}

	let chips = ['TODO', 'People', 'Projects'];
</script>

<SimplePane>
	<div class="flex items-center">
		<SearchBar bind:value={searchQuery} />
        <button
            type="button"
            onclick={() => console.log('Plus clicked')}
            aria-label="Open drawer"
            class="small-button bg-green-300 dark:bg-green-700 ml-2"
		>
			<svg class="feather">
				<use href="/icons/feather-sprite.svg#file-plus" />
			</svg>
		</button>

		{#each Object.entries($options) as [key, option]}
			<button
				class={`small-button p-1 ml-2 ${option.value ? 'bg-surface-600 dark:bg-surface-900' : 'bg-inherit'}`}
				onclick={() => {
					toggle(key);
				}}
			>
				{option.description}
			</button>
		{/each}

		<div style="margin-left: auto;">
			<SecondaryButton>
				Export
				<svg class="feather ml-2">
					<use href="/icons/feather-sprite.svg#upload" />
				</svg>
			</SecondaryButton>
		</div>
	</div>

	<div class="h-[calc(100%-3rem)] overflow-auto">
		<Card showChips={$options.showChips.value} {chips} showEdit={$options.showEdit.value} />
	</div>
</SimplePane>
