<script>
  import {
    ArchmagesList,
    GlobalStats,
    TopDailyLearners,
    TopCommunity,
    LoginUser,
    UserData,
  } from "../wailsjs/go/main/App.js";

  // Holds all Archmages and their publicly available data
  let archmages = [];
  // Holds the top 30 discord karma leaderboard users
  let archons = [];
  // Holds the top 30 leaderboard users
  let leaders = [];

  // Holds the general global leaderboard stats
  let stats = {};
  // is the current user an Archmage? Use for a few features
  // such as name highlighting

  // Empty field initially for the one-time password

  let otpField = "";

  // Information about the user

  let isArchmage = false;
  let isLoggedIn;

  // Because separating the keys with regex is worse
  const generalStats = {
    LessonCompletions: "Lessons Completed",
    XPEarned: "XP Earned",
    BootsChats: "Chats with Boots",
    SolutionsPeeked: "Solutions Peeked",
    DiscordMessagesSent: "Discord Messages",
    RegisteredUsersAlltime: "Registered Users",
  };

  // getArchmages returns an Array of the most recent 10 Archmages
  function getArchmages() {
    ArchmagesList().then((result) => (archmages = result.slice(0, 10)));
  }

  // getStats returns the current global leaderboard stats
  // The stats are shaped like [name,value]
  function getGlobalStats() {
    GlobalStats().then((result) => (stats = result));
  }

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

  // loginUser takes a user's OTP, trades it for an access token which
  // is saved for futher use, and marks the user as logged in
  function loginUser() {
    LoginUser(otpField).then((result) => (isLoggedIn = result));
  }

  // getUserData returns the user's Boot.Dev data, including courses
  // completed, current xp and to next level xp, and much more.
  function getUserData() {
    UserData().then((result) => console.log(result));
  }
</script>

<main>
  <div class="container-buddy">
    <div class="menu-container">
      {#if !isLoggedIn}
        <div class="menu-item btn-login">
          <div>
            You aren't currently logged in! You will only have limited
            functionality.
          </div>
          <div>
            Please
            <a
              href="https://www.boot.dev/cli/login?redirect=/cli/login"
              target="_blank">click here</a
            >
            to login.
          </div>
          The login instructions can&nbsp;<a
            href="https://github.com/ellielle/bootdev-buddy#logging-in"
            target="_blank"
          >
            be found here</a
          >.
        </div>
        <div class="menu-item">
          <input
            type="text"
            placeholder="Boot.Dev CLI Code"
            bind:value={otpField}
          />
          <button on:click={loginUser}>Sign in</button>
        </div>
      {/if}
      <div class="menu-item">
        <button class="btn" on:click={getArchmages}
          >Get Last 10 Archmages</button
        >
        {#each archmages as mage}
          <div>
            {mage?.Handle}
          </div>
        {/each}
      </div>
      <div class="menu-item">
        <button class="btn" on:click={getGlobalStats}>Show Stats!</button>
        {#each Object.entries(stats) as stat}
          <div>
            {generalStats[stat[0]] + ": " + stat[1]}
          </div>
        {/each}
      </div>
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
  }
</style>
