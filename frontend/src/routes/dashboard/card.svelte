<script lang="ts">
    export let showChips: bool = true;
    export let chips: string[] = [];
</script>

<div class="card p-2 bg-surface-300 dark:bg-surface-600" style="border: 1px solid;">
    <header>
        {#if showChips}
            {#each chips as chip, i}
                <span class="chip variant-filled ml-1">{chip}</span>
            {/each}
        {/if}
    </header>
    
    <section class="p-2">(content)</section>
</div>
