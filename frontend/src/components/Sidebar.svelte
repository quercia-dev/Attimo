<script>
    import homeIcon from "../assets/images/home.svg";
    import calendarIcon from "../assets/images/calendar.svg";
    import openIcon from "../assets/images/open.svg";
    import closeIcon from "../assets/images/close.svg";
    import dataIcon from "../assets/images/data.svg";
    import settingsIcon from "../assets/images/settings.svg";

    let menuItems = [
        { icon: homeIcon, label: "", route: "/", backgroundColor: "#f4f4f4" },
        { icon: calendarIcon, label: "", route: "/calendar", backgroundColor: "#9adfff" },
        { icon: openIcon, label: "", route: "/open", backgroundColor: "#73e600" },
        { icon: closeIcon, label: "", route: "/close", backgroundColor: "#ff9a9a" },
        { icon: dataIcon, label: "", route: "/data", backgroundColor: "#ffdf81" },
        { icon: settingsIcon, label: "", route: "/settings", backgroundColor: "#4444" },
    ];

    function navigateTo(route) {
        window.history.pushState({}, "", route);
    }
</script>

<nav>
  {#each menuItems as item}
    <button 
      on:click={() => navigateTo(item.route)} 
      class="menu-btn" 
      style="--button-bg-color: {item.backgroundColor};">
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
    height: 24rem;
    width: 4rem;
    background-color: #f4f4f4;
  }

  button {
    align-items: center;
    width: 4rem;
    height: 4rem;
    border: none;
    background-color: var(--button-bg-color, transparent); /* Use the CSS variable */
    cursor: pointer;
    transition: background-color 0.3s ease;
  }

  button:hover {
    background-color: #1b63d8;
  }

  img {
    width: 1.5rem;
    height: 1.5rem;
  }
</style>
