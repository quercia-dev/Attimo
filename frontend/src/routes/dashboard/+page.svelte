<script lang="ts">
	import SimplePane from '../../components/SimplePane.svelte';
	import SearchBar from '../../components/SearchBar.svelte';

	let searchQuery = $state('');
</script>

<SimplePane>
    <div style="display: flex; align-items: center;">
        <SearchBar bind:value={searchQuery} />
        <button
            type="button"
            onclick={() => console.log('Plus clicked')}
            aria-label="Open drawer"
            class="small-button bg-green-300 dark:bg-green-700"
            style="margin-left: 8px;"
        >
            <svg class="feather">
                <use href="/icons/feather-sprite.svg#plus" />
            </svg>
        </button>
    </div>
</SimplePane>
