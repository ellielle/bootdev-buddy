<script>
  import { onMount } from "svelte";
  import {
    TopCommunity,
    TopDailyLearners,
    GlobalStats,
  } from "../../../wailsjs/go/main/App.js";
  import { User } from "../../stores/user.js";
  import { prevent_default } from "svelte/internal";

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

  function openBrowserLink(url) {
    //@ts-ignore
    window.runtime.BrowserOpenURL(url);
  }

  onMount(() => {
    getTopDailyStats();
    getTopCommunity();
    getGlobalStats();

    const refreshInterval = setInterval(() => {
      getTopDailyStats();
      getGlobalStats();
    }, 600);

    return () => {
      clearInterval(refreshInterval);
    };
  });

  // TODO: split arcanum stats and leaderboard stats into their own components with refresh timers
</script>

<main>
  <div class="grid grid-cols-3">
    <ul class="list grid">
      <a
        href="https://www.boot.dev/leaderboard"
        target="_blank"
        on:click={(e) => {
          e.preventDefault();
          openBrowserLink("https://www.boot.dev/leaderboard");
        }}
      >
        <h2 class="text-primary-500">Arcanum Stats</h2>
      </a>
      {#each Object.entries(stats) as stat}
        <li>
          <span>
            {generalStats[stat[0]] + ": " + stat[1]}
          </span>
        </li>
      {/each}
    </ul>

    <div>
      <a
        href="https://www.boot.dev/leaderboard"
        target="_blank"
        on:click={(e) => {
          e.preventDefault();
          openBrowserLink("https://www.boot.dev/leaderboard");
        }}
      >
        <h2 class="pb-5 text-primary-500">Archon Leaderboard</h2>
      </a>
      <ul class="list">
        {#each archons as sage, index}
          <div class="list">
            <a
              href="https://www.boot.dev/u/{sage?.Handle}"
              target="_blank"
              on:click={(e) => {
                e.preventDefault();
                openBrowserLink("https://www.boot.dev/u/{sage?.Handle}");
              }}>{sage?.Handle}</a
            >
          </div>
        {/each}
      </ul>
    </div>

    <div>
      <a
        href="https://www.boot.dev/leaderboard"
        target="_blank"
        on:click={(e) => {
          e.preventDefault();
          openBrowserLink("https://www.boot.dev/leaderboard");
        }}
      >
        <h2 class="pb-5 text-primary-500">Daily Leaderboard</h2>
      </a>
      <ul class="list">
        {#each leaders as lead}
          <div>
            <a
              href="https://www.boot.dev/u/{lead?.Handle}"
              target="_blank"
              on:click={(e) => {
                e.preventDefault();
                openBrowserLink("https://www.boot.dev/u/{lead?.Handle}");
              }}
            >
              <span
                style={lead.Handle === $User.userData.Handle
                  ? "color: text-warning-500"
                  : "color: text-primary-500"}>{lead?.Handle}</span
              >
            </a>
          </div>
        {/each}
      </ul>
    </div>
  </div>
</main>
