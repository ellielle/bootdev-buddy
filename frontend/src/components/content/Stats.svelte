<script>
  import { onMount } from "svelte";
  import {
    TopCommunity,
    TopDailyLearners,
    GlobalStats,
  } from "../../../wailsjs/go/main/App.js";

  // Holds the top 30 discord karma leaderboard users
  let archons = [];
  // Holds the general global leaderboard stats
  let stats = {};
  // Holds the top 30 leaderboard users
  let leaders = [];

  // Because separating the keys with regex is worse
  const generalStats = {
    LessonCompletions: "Lessons Completed",
    XPEarned: "XP Earned",
    BootsChats: "Chats with Boots",
    SolutionsPeeked: "Solutions Peeked",
    DiscordMessagesSent: "Discord Messages",
    RegisteredUsersAlltime: "Registered Users",
  };

  // getGlobalStats returns the current global leaderboard stats
  // The stats are shaped like [name,value]
  function getTopDailyStats() {
    TopDailyLearners().then((result) => (leaders = result));
  }

  // getTopCommunity returns the top 30 users with the highest karma,
  // a stat earned by being active in the Discord channel
  function getTopCommunity() {
    TopCommunity().then((result) => (archons = result));
  }

  // getGlobalStats returns the current global leaderboard stats
  // The stats are shaped like [name,value]
  function getGlobalStats() {
    GlobalStats().then((result) => (stats = result));
  }

  onMount(() => {
    getTopDailyStats();
    getTopCommunity();
    getGlobalStats();
  });
</script>

<main>
  {#if false}
    {#each Object.entries(stats) as stat}
      <div>
        {generalStats[stat[0]] + ": " + stat[1]}
      </div>
    {/each}
    {#each archons as sage}
      <div>
        {sage?.Handle}
      </div>
    {/each}
    {#each leaders as lead}
      <div>
        {lead?.Handle}
      </div>
    {/each}
  {:else}
    Under work
  {/if}
</main>
