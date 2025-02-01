<script>
    let searchQuery = '';

    import Calendar from '@event-calendar/core';
    import TimeGrid from '@event-calendar/time-grid';

    let plugins = [TimeGrid];
    let options = {
        view: 'timeGridWeek',
        slotMinTime: '08:00:00',
        slotMaxTime: '20:00:00',
        slotWidth: 10,
        events: [
            {
                title: 'Event 1',
                start: '2025-01-29T10:00:00',
                end: '2025-01-29T12:00:00'
            },
        ]
    };
</script>

<main class="bg-green-500 p-4 space-y-4 flex flex-col flex-1">
    <div class="bg-white text-gray-900 p-4 space-y-2 rounded-md shadow-md">
        <h1 class="text-2xl font-bold">Search</h1>
        <input 
            id="search" 
            type="search" 
            placeholder="dammi un attimo..." 
            bind:value={searchQuery} 
            on:input={() => console.log(searchQuery)}
            class="mt-1 block w-4/5 rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-gray-800"
        />
    </div>
    <div class="bg-white text-gray-900 p-4 rounded-md shadow-md">
        <h1 class="text-2xl font-bold">Lists</h1>
    </div>
<div class="bg-white text-gray-900 p-4 rounded-md shadow-md overflow-auto">
            <h1 class="text-2xl font-bold">Calendar</h1>
            <Calendar {plugins} {options} />
    </div>
    <p class="min-h-[500px] min-w-[500px] bg-purple-500 p-4 flex-1">Home Pane</p>
</main>

<style>
    .ec-time-grid .ec-slot{
        min-width: 10px !important; /* Match slotWidth */
        max-width: 10px !important;
    }
  </style>