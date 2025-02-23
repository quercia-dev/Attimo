<script lang="ts">

export let tableArr = [];
</script>

<div class="table-container">
    <table class="table table-hover table-cell-fit bg-primary-200">
        <thead>
            <tr>
                <th>Close</th>
                <th>Table</th>
                <th>First Created</th>
                <th>Edited last</th>
                <th>Principal Value</th>
                <th>Discard</th>
            </tr>
        </thead>
        <tbody>
            {#each tableArr as row, i}
                <tr>
                    <td>
                        <button
                            type="button"
                            aria-label="Open drawer"
                            class="small-button bg-green-300 dark:bg-green-700"
                        >
                            <svg class="feather">
                                <use href="/icons/feather-sprite.svg#minus" />
                            </svg>
                        </button>
                    </td>
                    <td>
                        <button
                            type="button"
                            aria-label="Open drawer"
                            class="button"
                            style="white-space: nowrap;"
                        >
                            {row.tableName}
                        </button>
                    </td>
                    <td>
                        <button
                            type="button"
                            aria-label="Open drawer"
                            class="button"
                            style="white-space: nowrap;"
                        >
                            {row.openDate}
                        </button>
                    </td>
                    <td>
                        <button
                            type="button"
                            aria-label="Open drawer"
                            class="button"
                            style="white-space: nowrap;"
                        >
                            {row.lastEdit}
                        </button>
                    </td>
                    <td>
                        <button
                            type="button"
                            aria-label="Open drawer"
                            class="button"
                            style="white-space: nowrap;"
                        >
                            {row.valueSummary}
                        </button>
                    </td>
                    <td>
                        <button
                            type="button"
                            aria-label="Open drawer"
                            class="small-button bg-red-300 dark:bg-red-900"
                        >
                            <svg class="feather">
                                <use href="/icons/feather-sprite.svg#trash-2" />
                            </svg>
                        </button>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
