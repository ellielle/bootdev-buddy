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
    // FIXME: signing in isn't reactive and doesn't change pages
    // the issue seems to be frontend related. Data is fine on backend
    // but comes empty to frontend
    // * Sign in works fine from file though
    LoginUserWithToken().then((result) => ($User.isLoggedIn = result));
    UserData().then((result) => ($User.userData = result));
    console.log("userdata", $User.userData);
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
        <Login loggedIn={$User.isLoggedIn} />
      </div>
    </div>
  {/if}
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
