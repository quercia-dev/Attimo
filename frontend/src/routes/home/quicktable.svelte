<script lang="ts">
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
				<th>Table</th>
				<th>Entries</th>
				<th>Open</th>
				<th>Ongoing</th>
				<th>Close</th>
				<th>Activity Graph</th>
				<th>Week</th>
				<th>Month</th>
			</tr>
		</thead>
		<tbody>
			{#each tableArr as row, i}
            <tr>
				<td>
					<button
						type="button"
						aria-label="Open drawer"
						class="button"
						style="white-space: nowrap;"
					>
						{row.name}
					</button>
				</td>
				<td>
					<button
						type="button"
						aria-label="Open drawer"
						class="button"
						style="white-space: nowrap;"
					>
						{row.total}
					</button>
				</td>

                <td>
					<button
						type="button"
						onclick={() => showAppRail.update(value => !value)}
						aria-label="Open drawer"
						class="small-button bg-green-300 dark:bg-green-700"
					>
						<svg class="feather">
							<use href="/icons/feather-sprite.svg#plus" />
						</svg>
					</button>
                </td>
				{#if row.continuative}
				<td>
					<button
						type="button"
						aria-label="Open drawer"
						class="button"
						style="white-space: nowrap;"
					>
						{row.ongoing}
					</button>
				</td>

                <td>
                    <button
                        type="button"
                        onclick={() => showAppRail.update(value => !value)}
                        aria-label="Open drawer"
                        class="small-button bg-red-300 dark:bg-red-900"
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

				<td>
					<button
						type="button"
						aria-label="Open drawer"
						class="button"
						style="white-space: nowrap;"
					>
						{row.lastWeek}
					</button>
				</td>
				<td>
					<button
						type="button"
						aria-label="Open drawer"
						class="button"
						style="white-space: nowrap;"
					>
						{row.lastMonth}
					</button>
				</td>
			</tr>
			{/each}
		</tbody>
	</table>
</div>