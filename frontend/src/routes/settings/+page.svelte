<script lang="ts">
    import SimplePane from '../../components/SimplePane.svelte';
    import { SlideToggle } from '@skeletonlabs/skeleton';
    import { theme, setTheme } from '$lib/stores/themeStore';

    let value: boolean = false;
    function updateTheme(event: Event) {
        const newTheme = (event.target as HTMLSelectElement).value;
        setTheme(newTheme);
    }

</script>


<SimplePane>
    <h1>Settings</h1>

    <div class="space-y-4">
        <!-- Theme Selection -->
        <label class="label">
            <span>Theme</span>
            <select class="select" bind:value={$theme} on:change={updateTheme}>
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
            <a href="https://github.com/quercia-dev/Attimo" target="_blank" class="btn variant-filled-primary">Github</a>
        </div>
    </div>
</SimplePane>
