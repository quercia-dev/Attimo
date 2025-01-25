<script>

    import homeIcon from "../assets/images/home.svg";
    import calendarIcon from "../assets/images/calendar.svg";
    import openIcon from "../assets/images/open.svg";
    import closeIcon from "../assets/images/close.svg";
    import dataIcon from "../assets/images/data.svg";
    import settingsIcon from "../assets/images/settings.svg";

    let menuItems = [
        { icon: homeIcon, label: "", route: "/" },
        { icon: calendarIcon, label: "", route: "/calendar" },
        { icon: openIcon, label: "", route: "/open" },
        { icon: closeIcon, label: "", route: "/close" },
        { icon: dataIcon, label: "", route: "/data" },
        { icon: settingsIcon, label: "", route: "/settings" },
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
        height: 300px;
        width: 200%;
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