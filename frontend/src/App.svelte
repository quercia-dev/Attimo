<script>
  import {Greet} from '../wailsjs/go/main/App.js'
  import Calendar from './components/Calendar.svelte';  
  import Lists from './components/Lists.svelte';
  import Overview from './components/Overview.svelte';
  import SearchBar from './components/SearchBar.svelte';
  import Logo from './components/Logo.svelte';
  import Sidebar from './components/Sidebar.svelte';

  let resultText = "Filter the tables by name"
  let name
  function greet() {
    Greet(name).then(result => resultText = result)
  }

</script>

<main>
  <Sidebar />
  
  <div class="container">
    <div class="result" id="result">{resultText}</div>
    <div class="input-box" id="input">
      <input autocomplete="off" bind:value={name} class="input" id="name" type="text" placeholder="Table name..."/>
      <button class="btn" id="tot-button" on:click={greet} title="Total Entries">####</button>
      <button class="btn" id="thirty-day-button" on:click={greet} title="Entries in the last 30 days">#####</button>     
    </div>
    <Overview />
    <Logo />
  </div>  

</main>

<style>
  main {
    display: flex;
    height: 100vh;
  }

  .container {
   flex-grow: 1;
   position: relative;
   display: flex;
   flex-direction: column;
   padding: 10px;
   align-self: stretch;
 }
 

 
 #result {
  height: 20px;
  line-height: 20px;
  margin: 1rem 0;
  text-align: left;
}
 
 .input-box {
   display: flex;
   width: 100%;
   gap: 20px;
 }
 
 .input-box .btn {
   flex-shrink: 0;
   width: 60px;
   height: 30px;
   line-height: 30px;
   border-radius: 3px;
   border: none;
   padding: 0 8px;
   cursor: pointer;
 }
 
 .input-box .btn:hover {
   background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
   color: #4444;
 }
 
 .input-box .input {
   flex-grow: 1;
   border: none;
   border-radius: 3px;
   outline: none;
   height: 30px;
   line-height: 30px;
   padding: 0 10px;
   background-color: rgba(240, 240, 240, 1);
   -webkit-font-smoothing: antialiased
 }
 
 .input-box .input:hover,
 .input-box .input:focus {
   background-color: rgba(255, 255, 255, 1);
 }
 </style>
