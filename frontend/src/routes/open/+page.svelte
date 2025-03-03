<script lang="ts">
    import { categories, formFields } from "../../stores/formStore";
    import type { FormField } from '../../types/types';
    import SimplePane from "../../components/SimplePane.svelte";
    import SearchBar from '../../components/SearchBar.svelte';
    
    let searchQuery = '';
    let selectedCategory = "general"; //default
    let filteredFields: FormField[] = [];

    // Handle category selection
    $: {
        if ($formFields[selectedCategory]) {
            filteredFields = $formFields[selectedCategory];
        }
    }

    // Search/filter functionality
    $: {
        if (searchQuery) {
            filteredFields = $formFields[selectedCategory].filter(field => 
                field.label.toLowerCase().includes(searchQuery.toLowerCase())
            );
        } else {
            filteredFields = $formFields[selectedCategory];
        }
    }
</script>

<SimplePane>
    <div class="flex flex-col space-y-4">
        <div style="display: flex; align-items: center;">
            <SearchBar bind:value={searchQuery} />
            <button
                type="button"
                onclick={() => console.log('Plus clicked')}
                aria-label="Open drawer"
                class="small-button bg-green-300 dark:bg-green-700 ml-2"
            >
                <svg class="feather">
                    <use href="/icons/feather-sprite.svg#plus" />
                </svg>
            </button>
        </div>    
    </div>

    <div class="p-2">
        <p>This button is here for TESTING purposes- not part of final build</p>
    <!-- Category Selection -->
    <select 
        bind:value={selectedCategory}
        class="select w-full max-w-32"
        size="3"
    >
        {#each $categories as category}
            <option value={category.id}>{category.name}</option>
        {/each}
    </select>
    </div>
    <!-- Dynamic Form -->
    <div class="p-4 space-y-4">
        {#each filteredFields as field (field.id)}
            <div class="form-group">
                <label 
                    for={field.id}
                    class="block text-sm font-medium mb-1"
                >
                    {field.label}
                    {#if field.required}
                        <span class="text-red-500">*</span>
                    {/if}
                </label>
                
                {#if field.type === 'select'}
                    <select
                        id={field.id}
                        class="select"
                    >
                        <option value="">Select an option</option>
                        {#each field.options || [] as option}
                            <option value={option}>{option}</option>
                        {/each}
                    </select>
                {:else if field.type === 'textarea'}
                    <textarea
                        class="textarea"
                        rows="4"
                        id={field.id}
                        placeholder={field.placeholder}
                    ></textarea>
                {:else if field.type === 'file'}
                    <input
                        class="input"
                        type="file"
                        id={field.id}
                    />
                {:else}
                    <input
                        class="input"
                        id={field.id}
                        type={field.type}
                        placeholder={field.placeholder}
                    />
                {/if}
            </div>
        {/each}
    </div>

</SimplePane>
