<script>
  import "./app.css";
  import { onMount } from "svelte";
  import { UserData, LoginUserWithToken } from "../wailsjs/go/main/App.js";
  import Login from "./components/Login.svelte";
  import Tabs from "./components/UI/Tabs.svelte";
  import { User } from "./stores/user.js";

  let userPromise = UserData().then((result) => ($User.userData = result));

  onMount(() => {
    // Attempt to log the user on mount by refreshing their
    // access token
    LoginUserWithToken().then((result) => ($User.isLoggedIn = result));
  });
</script>

<main>
  <!-- insert Tab component. The rest of the content is rendered via Tabs -->
  {#await userPromise}
    <p>Loading...</p>
  {:then _}
    <Tabs />
    <div class="container-buddy">
      <div class="menu-container">
        <!-- show user login button if automatic sign in fails -->
        {#if !$User.isLoggedIn || typeof $User.isLoggedIn != "boolean"}
          <Login bind:loggedIn={$User.isLoggedIn} />
        {/if}
      </div>
    </div>
  {/await}
</main>

<style>
  .container-buddy {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-auto-flow: row;
  }

  main {
    padding: 1em 0 0 1em;
  }
</style>
