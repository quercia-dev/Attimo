<script lang="ts">
	import { TableHandler, Datatable, ThSort, ThFilter } from '@vincjo/datatables';
	import { writable } from 'svelte/store';
	import { showAppRail } from '../../stores/appStore';

	const tableArr = [
		{ id: 1, name: 'Todo', total: 100, continuative: false, ongoing: 70, lastWeek: 50, lastMonth: 80 },
		{ id: 2, name: 'Projects', total: 80, continuative: true, ongoing: 80, lastWeek: 60, lastMonth: 70 },
		{ id: 3, name: 'People', total: 65, continuative: true, ongoing: 65, lastWeek: 40, lastMonth: 50 },
		{ id: 4, name: 'People', total: 65, continuative: true, ongoing: 65, lastWeek: 40, lastMonth: 50 },
		{ id: 5, name: 'People', total: 65, continuative: false, ongoing: 65, lastWeek: 40, lastMonth: 50 },
		{ id: 6, name: 'Tasks', total: 75, continuative: true, ongoing: 75, lastWeek: 55, lastMonth: 65 }
	];
</script>


<div class="table-container">
	
	<table class="table table-hover table-cell-fit bg-primary-200">
		<thead>
			<tr>
				<th>Name</th>
				<th>Number</th>
				<th></th>
				<th>Open</th>
				<th></th>
				<th>Activity Graph</th>
				<th>Last week</th>
				<th>Last month</th>
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
						class="small-icon-button bg-green-300 dark:bg-green-700"
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
                        class="small-icon-button bg-red-300 dark:bg-red-900"
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
				<td style="text-align: center; width: 100%;">|___/\___/\_____</td>

				<td style="text-align: center;">{row.lastWeek}</td>
				<td style="text-align: center;">{row.lastMonth}</td>
			</tr>
			{/each}
		</tbody>
	</table>
</div>


<style>

</style>
