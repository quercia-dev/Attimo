<script>

    import homeIcon from "../assets/images/home.png";
    import calendarIcon from "../assets/images/calendar.png";
    import openIcon from "../assets/images/open.png";
    import closeIcon from "../assets/images/close.png";
    import dataIcon from "../assets/images/data.png";
    import settingsIcon from "../assets/images/settings.png";

    let menuItems = [
        { icon: homeIcon, label: "Home", route: "/" },
        { icon: calendarIcon, label: "Calendar", route: "/calendar" },
        { icon: openIcon, label: "Open", route: "/open" },
        { icon: closeIcon, label: "Close", route: "/close" },
        { icon: dataIcon, label: "Data", route: "/data" },
        { icon: settingsIcon, label: "Settings", route: "/settings" },
    ];


    function navigateTo(route) {
        window.location.href = route;
    }
</script>

<nav>
 {#each menuItems as item}
    <button on:click={() => navigateTo(item.route)}>
        <img src={item.icon} alt={item.label} />
        <span>{item.label}</span>    
    </button>
    {/each}
</nav>

<style> 
    nav {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        height: 100%;
        width: 100%;
        padding: 1rem;
        background-color: #f4f4f4;
    }

    button {
        display: flex;
        align-items: center;
        padding: 0.5rem;
        border: none;
        background-color: transparent;
        cursor: pointer;
    }

    img {
        width: 16px;
        height: 16px;
    }
    </style>