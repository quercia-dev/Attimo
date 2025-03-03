<script lang="ts">
	import SimplePane from '../../components/SimplePane.svelte';
	import SearchBar from '../../components/SearchBar.svelte';
	import OpenedTable from './openedTable.svelte';
	import SecondaryButton from '../../components/SecondaryButton.svelte';

	const tableArr = [
        { id: 1, tableName: 'Todo', openDate: '2025-01-02 10:30', lastEdit: '2025-01-15 11:18', valueSummary: 'Personal' },
        { id: 2, tableName: 'Projects', openDate: '2025-02-10 09:00', lastEdit: '2025-02-20 14:45', valueSummary: 'Work' },
        { id: 3, tableName: 'People', openDate: '2025-03-05 08:30', lastEdit: '2025-03-15 10:20', valueSummary: 'Contacts' },
        { id: 4, tableName: 'Tasks', openDate: '2025-04-12 11:00', lastEdit: '2025-04-22 12:30', valueSummary: 'Work' },
        { id: 5, tableName: 'Events', openDate: '2025-05-18 13:45', lastEdit: '2025-05-28 15:50', valueSummary: 'Personal' },
        { id: 6, tableName: 'Notes', openDate: '2025-06-25 07:15', lastEdit: '2025-07-05 09:40', valueSummary: 'Personal' }
	];

	let searchQuery = $state('');
</script>

<SimplePane>
    <div style="display: flex; align-items: center;">
        <SearchBar bind:value={searchQuery} />

        <div style="margin-left: auto;">
            <SecondaryButton>
                History
            	<svg class="feather ml-2">
                    <use href="/icons/feather-sprite.svg#trash-2" />
                </svg>
            </SecondaryButton>
        </div>
    </div>

    <div style="height: calc(100% - 3rem); overflow: auto;">
        <OpenedTable {tableArr} />
    </div>
</SimplePane>
