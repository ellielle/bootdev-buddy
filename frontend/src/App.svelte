<script>
  import "./app.css";
  import { onMount } from "svelte";
  import {
    TopDailyLearners,
    TopCommunity,
    UserData,
    LoginUserWithToken,
  } from "../wailsjs/go/main/App.js";
  import Login from "./components/Login.svelte";
  import Tabs from "./components/UI/Tabs.svelte";

  // Holds all Archmages and their publicly available data
  let archmages = [];
  // Holds the top 30 discord karma leaderboard users
  let archons = [];
  // Holds the top 30 leaderboard users
  let leaders = [];

  // Information about the user
  //TODO: archmage, loggedin, and possibly more need to be in a store probably(?)
  export let isArchmage = false;
  export let isLoggedIn = false;

  // getTopLearners returns the 30 highest XP earners in the last
  // 24 hour period, sliced down to 10 to fit
  function getTopLearners() {
    TopDailyLearners().then((result) => (leaders = result.slice(0, 10)));
  }

  // getTopCommunity returns the top 30 users with the highest karma,
  // a stat earned by being active in the Discord channel
  function getTopCommunity() {
    TopCommunity().then((result) => (archons = result.slice(0, 10)));
  }

  // getUserData returns the user's Boot.Dev data, including courses
  // completed, current xp and to next level xp, and much more.
  function getUserData() {
    UserData().then((result) => console.log(result));
  }

  // Attempt to log the user on mount by refreshing their
  // access token
  onMount(() => {
    LoginUserWithToken().then((result) => (isLoggedIn = result));
  });
</script>

<main>
  <!-- insert Tab component -->
  <Tabs />
  <div class="container-buddy">
    <div class="menu-container">
      <!-- show user login button if automatic sign in fails -->
      <Login loggedIn={isLoggedIn} />

      <div class="menu-item">
        <button class="btn" on:click={getTopLearners}
          >Show Learning Leaders!</button
        >
        {#each leaders as lead}
          <div>
            {lead?.Handle}
          </div>
        {/each}
      </div>

      <div class="menu-item">
        <button class="btn" on:click={getTopCommunity}
          >Show Discord Leaders!</button
        >
        {#each archons as sage}
          <div>
            {sage?.Handle}
          </div>
        {/each}
      </div>

      <div class="menu-item">
        <button class="btn" on:click={getUserData}>Get User Data</button>
        {#each archons as sage}
          <div>
            {sage?.Handle}
          </div>
        {/each}
      </div>
    </div>
  </div>
  <div class="container-data"></div>
</main>

<style>
  .container-buddy {
    display: grid;
    grid-template-columns: 1fr 2fr;
    grid-auto-flow: row;
    color: orange;
  }
</style>
