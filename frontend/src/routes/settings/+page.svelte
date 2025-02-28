<script lang="ts">
    import SimplePane from '../../components/SimplePane.svelte';
    import { SlideToggle } from '@skeletonlabs/skeleton';
    import { browser } from '$app/environment';
    import { onMount } from "svelte";

    let value: boolean = false;
    let theme = 'earth';

    if (browser) {
        const savedTheme = localStorage.getItem('theme');
        if (savedTheme) {
            theme = savedTheme;
        }
    }

    function updateTheme(event: Event) {
        if (!browser) return;
        const newTheme = (event.target as HTMLSelectElement).value;
        theme = newTheme;
        localStorage.setItem('theme', newTheme);
        document.body.setAttribute('data-theme', newTheme);
    }

    // ensure there is a theme applied
    onMount(() => {
        if (browser) {
            document.body.setAttribute('data-theme', theme);
        }
    });

</script>


<SimplePane>
    <h1>Settings</h1>

    <div class="space-y-4">
        <!-- Theme Selection -->
        <label class="label">
            <span>Theme</span>
            <select class="select" bind:value={theme} on:change={updateTheme}>
                <option value="earth">Earth</option>
                <option value="quercus">Quercus</option>
                <option value="vintage">Halloween</option>
                <option value="rocket">Blue Lagoon</option>
            </select>
        </label>

        <!-- Autocomplete Toggle -->
        <label class="label">
            <span>Autocomplete</span>
            <SlideToggle name="slide" bind:checked={value} />
        </label>

        <!-- Contact Us -->
        <div class="card p-4 variant-filled-surface">
            <h3 class="h3">Contact Us</h3>
            <p class="py-2">Need help or want to contribute? Check out our github.</p>
            <a href="https://github.com/quercia-dev/Attimo" class="btn variant-filled-primary">Github</a>
        </div>
    </div>
</SimplePane>
