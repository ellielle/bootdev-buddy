<script>
  import "./app.css";
  import { onMount } from "svelte";
  import { UserData, LoginUserWithToken } from "../wailsjs/go/main/App.js";
  import Login from "./components/Login.svelte";
  import Tabs from "./components/UI/Tabs.svelte";
  import { User } from "./stores/user.js";

  function updateUserData() {
    UserData().then((result) => ($User.userData = result));
  }

  // FIXME: i'm bad at svelte reactivity
  onMount(async () => {
    // Attempt to log the user on mount by refreshing their
    // access token
    LoginUserWithToken().then((result) => ($User.isLoggedIn = result));
    UserData()
      .then((result) => ($User.userData = result))
      .catch((e) => console.log(e));
    // Update user data every 60 seconds
    const refreshInterval = setInterval(() => {
      if ($User.isLoggedIn) {
        updateUserData();
      }
    }, 30000);

    return () => {
      clearInterval(refreshInterval);
    };
  });
</script>

<main>
  {#if $User.isLoggedIn && $User.userData.Handle !== ""}
    <Tabs />
  {:else if $User.isLoggedIn}
    <p>Loading data...</p>
    <p>Data not loading? Please restart the app!</p>
  {:else}
    <div class="container-buddy">
      <div class="menu-container">
        <!-- show user login button if automatic sign in fails -->
        <Login />
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
