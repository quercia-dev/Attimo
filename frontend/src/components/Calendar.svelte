<script>
    import Calendar from '@event-calendar/core';
    import TimeGrid from '@event-calendar/time-grid';

    let plugins = [TimeGrid];
    let options = {
        view: 'timeGridWeek',
        firstDay: 1,
        //slotMinTime: '08:00:00',
        //slotMaxTime: '20:00:00',
        events: [
            {
                title: 'Event 1',
                start: '2025-01-29T10:00:00',
                end: '2025-01-29T12:00:00',
            },
        ]
    };
</script>

<Calendar {plugins} {options} />