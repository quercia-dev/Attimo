<script>
    let searchQuery = '';
    let topDivHeight = 150; 

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
            }
        ]
    };

    function startResizing(event) {
        const startY = event.clientY;
        const startHeight = topDivHeight;

        function onMouseMove(e) {
            const newHeight = startHeight + (e.clientY - startY);
            const maxHeight = window.innerHeight - 200; // Leave space for calendar
            topDivHeight = Math.max(50, Math.min(newHeight, maxHeight));
        }

        function onMouseUp() {
            window.removeEventListener('mousemove', onMouseMove);
            window.removeEventListener('mouseup', onMouseUp);
        }

        window.addEventListener('mousemove', onMouseMove);
        window.addEventListener('mouseup', onMouseUp);
    }
</script>

<main class="bg-surface-500 dark:bg-inherit p-4 flex flex-col h-[calc(100vh-8rem)] overflow-hidden no-select">
    <div class="bg-surface-200 dark:bg-surface-500 p-4 space-y-2 rounded-md shadow-xl flex-shrink-0" 
         style="height: {topDivHeight}px; min-height: 110px;">
        <h1 class="text-2xl font-bold">Search</h1>
        <input
            id="search"
            type="search"
            placeholder="dammi un Attimo..."
            bind:value={searchQuery}
            on:input={() => console.log(searchQuery)}
            class="mt-1 block w-4/5 rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-gray-800"
        />
    </div>

    <div class="slider bg-inherit flex-shrink-0" on:mousedown={startResizing}></div>
    
    <div class="bg-white text-gray-900 p-4 rounded-md shadow-xl flex-1 overflow-hidden">
        <div class="calendar-container h-full w-full">
            <Calendar {plugins} {options} />
        </div>
    </div>
</main>

<style>
    .calendar-container {
        width: 100%;
        height: 100%;
        overflow: auto;
    }
    .slider {
        height: 10px;
        background: transparent;
        cursor: row-resize;
    }
    .no-select {
        user-select: none;
    }
</style>