<script lang="ts">
	import { TableHandler, Datatable, ThSort, ThFilter } from '@vincjo/datatables';
	import { writable } from 'svelte/store';
	import { showAppRail } from '../stores/appStore';

	const tableArr = [
		{ id: 1, name: 'Todo', total: 100, continuative: false, ongoing: 70, lastWeek: 50, },
		{ id: 2, name: 'Projects', total: 80, continuative: true, ongoing: 80, lastWeek: 60, },
		{ id: 3, name: 'People', total: 65, continuative: true, ongoing: 65, lastWeek: 40, },
		{ id: 4, name: 'People', total: 65, continuative: true, ongoing: 65, lastWeek: 40, },
		{ id: 5, name: 'People', total: 65, continuative: false, ongoing: 65, lastWeek: 40, },
		{ id: 6, name: 'Tasks', total: 75, continuative: true, ongoing: 75, lastWeek: 55, }
	];
</script>


<div class="table-container">
	
	<table class="table table-hover table-cell-fit">
		<thead>
			<tr>
				<th>Name</th>
				<th>Number</th>
				<th></th>
				<th>Open</th>
				<th></th>
				<th>Last week</th>
				<th>Activity Graph</th>
			</tr>
		</thead>
		<tbody>
			{#each tableArr as row, i}
            <tr>
                <td>{row.name}</td>
				<td style="text-align: center;">{row.total}</td>

                <td>
                    <button
                        type="button"
                        onclick={() => showAppRail.update(value => !value)}
                        aria-label="Open drawer"
                        class="hover:bg-primary-300 dark:hover:bg-primary-900"
                    >
                        <svg class="feather">
                            <use href="/icons/feather-sprite.svg#plus" />
                        </svg>
                    </button>
                </td>
				{#if row.continuative}
				<td style="text-align: center;">{row.ongoing}</td>

                <td>
                    <button
                        type="button"
                        onclick={() => showAppRail.update(value => !value)}
                        aria-label="Open drawer"
                        class="hover:bg-primary-300 dark:hover:bg-primary-900"
                    >
                        <svg class="feather">
                            <use href="/icons/feather-sprite.svg#minus" />
                        </svg>
                    </button>
                </td>
				{:else}
				<td></td>
				<td></td>
				{/if}
				<td style="text-align: center;">{row.lastWeek}</td>
				
				<td style="text-align: center; width: 100%;">|___/\___/\_____</td>
			</tr>
			{/each}
		</tbody>
	</table>
</div>


<style>

</style>
