<script>
  import "./app.css";
  import { onMount } from "svelte";
  import { UserData, LoginUserWithToken } from "../wailsjs/go/main/App.js";
  import Login from "./components/Login.svelte";
  import Tabs from "./components/UI/Tabs.svelte";

  // Holds all Archmages and their publicly available data
  let archmages = [];
  // Holds the top 30 leaderboard users
  let leaders = [];

  // Information about the user
  //TODO: archmage, loggedin, and possibly more need to be in a store probably(?)
  export let isArchmage = false;
  export let isLoggedIn = false;
  export let userData;

  // getUserData returns the user's Boot.Dev data, including courses
  // completed, current xp and to next level xp, and much more.
  function getUserData() {}

  // Attempt to log the user on mount by refreshing their
  // access token
  onMount(() => {
    LoginUserWithToken().then((result) => (isLoggedIn = result));
    UserData().then((result) => (userData = result));
  });
</script>

<main>
  <!-- insert Tab component -->
  <Tabs />
  <div class="container-buddy">
    <div class="menu-container">
      <!-- show user login button if automatic sign in fails -->
      {#if !isLoggedIn || typeof isLoggedIn != "boolean"}
        <Login loggedIn={isLoggedIn} />
      {/if}
    </div>
  </div>
  <div class="container-data"></div>
</main>

<style>
  .container-buddy {
    display: grid;
    grid-template-columns: 1fr 2fr;
    grid-auto-flow: row;
  }
</style>
