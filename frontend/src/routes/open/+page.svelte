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
        <div class="flex items-center gap-2">
            <input
                id="search"
                type="search"
                placeholder="Dammi un Attimo..."
                bind:value={searchQuery}
                class="bg-surface-100 dark:bg-surface-700 text-surface-900 dark:text-surface-50 placeholder:text-surface-800 mt-1 block w-4/5 rounded-md shadow-sm"
            />
            <button
            type="button"
            onclick={() => console.log('search dashboard')}
            aria-label="Open drawer"
            class="icon-button hover:bg-surface-400 dark:hover:bg-primary-600"
        >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="32"
                    height="32"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="3"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    class="feather feather-plus"
                    >
                    <line x1="12" y1="5" x2="12" y2="19"></line>
                    <line x1="5" y1="12" x2="19" y2="12"></line>
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
</main>
</SimplePane>
