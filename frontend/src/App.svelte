<script>
  import logo from './assets/images/logo-universal.png'
  import {Greet} from '../wailsjs/go/main/App.js'
  import Calendar from './components/Calendar.svelte';  
  import Lists from './components/Lists.svelte';
  import Overview from './components/Overview.svelte';
  import SearchBar from './components/SearchBar.svelte';
  import Sidebar from './components/Sidebar.svelte';

</script>

<main>
  <Sidebar />
</main>

<style>
  main {
    display: flex;
    height: 100vh;
  }
</style>
