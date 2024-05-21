<script>
  import {
    ArchmagesList,
    GlobalStats,
    TopDailyLearners,
    TopCommunity,
  } from "../wailsjs/go/main/App.js";

  // Holds all Archmages and their publicly available data
  let archmages = [];
  // Holds the top 30 discord karma leaderboard users
  let archsages = [];
  // Holds the top 30 leaderboard users
  let leaders = [];

  // Holds the general global leaderboard stats
  let stats = {};
  // is the current user an Archmage? Use for a few features
  // such as name highlighting

  let isArchmage = false;

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

  function getTopLearners() {
    TopDailyLearners().then((result) => (leaders = result.slice(0, 10)));
  }

  function getTopCommunity() {
    TopCommunity().then((result) => (archsages = result.slice(0, 10)));
  }
</script>

<main>
  <div class="menu-container">
    <div class="menu-item">
      <button class="btn" on:click={getArchmages}>Get Last 10 Archmages</button>
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
      {#each archsages as sage}
        <div>
          {sage?.Handle}
        </div>
      {/each}
    </div>
  </div>
</main>

<style>
</style>
