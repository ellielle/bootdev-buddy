<script>
  import "./app.css";
  import { onMount } from "svelte";
  import { UserData, LoginUserWithToken } from "../wailsjs/go/main/App.js";
  import Login from "./components/Login.svelte";
  import Tabs from "./components/UI/Tabs.svelte";
  import { User } from "./stores/user.js";

  // Holds all Archmages and their publicly available data
  let archmages = [];
  // Holds the top 30 leaderboard users
  let leaders = [];

  // Information about the user

  // Attempt to log the user on mount by refreshing their
  // access token
  onMount(() => {
    LoginUserWithToken().then((result) => ($User.isLoggedIn = result));
    UserData().then((result) => ($User.userData = result));
  });
</script>

<main>
  <!-- insert Tab component. The rest of the content is rendered via Tabs -->
  <Tabs />
  <div class="container-buddy">
    <div class="menu-container">
      <!-- show user login button if automatic sign in fails -->
      {#if !$User.isLoggedIn || typeof $User.isLoggedIn != "boolean"}
        <Login loggedIn={$User.isLoggedIn} />
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

  main {
    padding: 1em 0 0 1em;
  }
</style>
