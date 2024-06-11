<script>
  import "./app.css";
  import { onMount } from "svelte";
  import { UserData, LoginUserWithToken } from "../wailsjs/go/main/App.js";
  import Login from "./components/Login.svelte";
  import Tabs from "./components/UI/Tabs.svelte";
  import { User } from "./stores/user.js";

  $: loggedIn = $User.isLoggedIn;

  onMount(() => {
    // Attempt to log the user on mount by refreshing their
    // access token
    LoginUserWithToken().then((result) => ($User.isLoggedIn = result));
    UserData().then((result) => ($User.userData = result));
  });
</script>

<main>
  {#if loggedIn}
    <Tabs />
  {/if}
  {#if !$User.isLoggedIn || typeof $User.isLoggedIn != "boolean"}
    <div class="container-buddy">
      <div class="menu-container">
        <!-- show user login button if automatic sign in fails -->
        <Login bind:loggedIn={$User.isLoggedIn} />
      </div>
    </div>
  {/if}
</main>

<style>
  .container-buddy {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-auto-flow: row;
  }

  main {
    padding: 1em;
  }
</style>
