<script lang="ts">
	import SimplePane from '../../components/SimplePane.svelte';
	import SearchBar from '../../components/SearchBar.svelte';
    import Card from './card.svelte'

	let searchQuery = $state('');

    type Option = {
        value: boolean;
        description: string;
    };

    import { writable } from 'svelte/store';

    let options = writable<Record<string, Option>>({
        hideSources: { value: false, description: 'Sources' },
        hideEditLogo: { value: true, description: 'Edit' },
    });

    function toggle(key: string): void {
        options.update(opts => {
            opts[key].value = !opts[key].value;
            return opts;
        });
    }

    let chips = ['TODO', 'People', 'Projects'];
</script>

<SimplePane>
    <div class="flex-shrink-0 flex items-center w-auto h-auto">
        <SearchBar bind:value={searchQuery} />
        <button
            type="button"
            onclick={() => console.log('Plus clicked')}
            aria-label="Open drawer"
            class="small-button bg-green-300 dark:bg-green-700"
            style="margin-left: 8px;"
        >
            <svg class="feather">
                <use href="/icons/feather-sprite.svg#file-plus" />
            </svg>
        </button>

        {#each Object.entries($options) as [key, option]}
            <button
                class={`small-button p-1 ${option.value ? 'bg-surface-600 dark:bg-surface-900' : 'bg-inherit'}`}
                onclick={() => { toggle(key); }}
                style="margin-left: 8px;"
            >
                {option.description}
            </button>
        {/each}

        <div style="margin-left: auto;">
            <button
                type="button"
                aria-label="Open drawer"
                class="button bg-secondary-300 dark:bg-secondary-600 flex items-center"
                style="box-shadow: 0 0 0 1px;"
            >
                Export
                <svg class="feather ml-2">
                    <use href="/icons/feather-sprite.svg#upload" />
                </svg>
            </button>
        </div>
    </div>

    <div class="h-[calc(100%-3rem)] overflow-auto">
        <Card {chips}/>
    </div>

</SimplePane>
