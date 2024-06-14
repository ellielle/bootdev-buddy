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

  // skeletonUI color palette numbers, reversed so I don't have to do anything
  // funky in the HTML
  const skeletonUINum = [
    "900",
    "800",
    "700",
    "600",
    "500",
    "400",
    "300",
    "200",
    "100",
    "50",
  ];

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

  // TODO: split arcanum stats and leaderboard stats into their own components with refresh timers
</script>

<main>
  <div class="grid grid-cols-3">
    <ul class="list grid">
      <h2 class="text-primary-500">Arcanum Stats</h2>
      {#each Object.entries(stats) as stat}
        <li>
          <span>
            {generalStats[stat[0]] + ": " + stat[1]}
          </span>
        </li>
      {/each}
    </ul>

    <div>
      <h2 class="pb-5 text-primary-500">Archon Leaderboard</h2>
      <ul class="list">
        {#each archons as sage, index}
          <div class="list text-warning-{skeletonUINum[index]}">
            {sage?.Handle}
          </div>
        {/each}
      </ul>
    </div>

    <div>
      <h2 class="pb-5 text-primary-500">Daily Leaderboard</h2>
      <ul class="list">
        {#each leaders as lead}
          <div>
            {lead?.Handle}
          </div>
        {/each}
      </ul>
    </div>
  </div>
</main>
