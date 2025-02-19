<script lang="ts">
    import { onMount } from "svelte";
    import { categories, formFields } from "../../stores/formStore";
    import type { FormField } from '../../types/types';
    import SimplePane from "../../components/SimplePane.svelte";
    
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
<main class="p-4">
    <div class="flex flex-col space-y-4">
        <!-- Search and Add buttons -->
        <div class="flex items-center gap-4">
            <input
                id="search"
                type="search"
                placeholder="Dammi un Attimo..."
                bind:value={searchQuery}
                class="bg-surface-100 dark:bg-surface-700 text-surface-900 dark:text-surface-50 placeholder:text-surface-800 mt-1 block w-4/5 rounded-md shadow-sm"
            />
            <button class="p-2 rounded bg-blue-500 text-white">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
                </svg>
            </button>
        </div>
    </div>

    <!-- Category Selection -->
    <select 
        bind:value={selectedCategory}
        class="p-2 rounded border dark:bg-gray-700 dark:text-white"
    >
        {#each $categories as category}
            <option value={category.id}>{category.name}</option>
        {/each}
    </select>

    <!-- Dynamic Form -->
    <div class="space-y-4">
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
                        class="w-full p-2 rounded border"
                    >
                        <option value="">Select an option</option>
                        {#each field.options || [] as option}
                            <option value={option}>{option}</option>
                        {/each}
                    </select>
                {:else if field.type === 'textarea'}
                    <textarea
                        id={field.id}
                        placeholder={field.placeholder}
                        class="w-full p-2 rounded border"
                    ></textarea>
                {:else if field.type === 'file'}
                    <input
                        id={field.id}
                        type="file"
                        class="w-full p-2"
                    />
                {:else}
                    <input
                        id={field.id}
                        type={field.type}
                        placeholder={field.placeholder}
                        class="w-full p-2 rounded border"
                    />
                {/if}
            </div>
        {/each}
    </div>
</main>
</SimplePane>