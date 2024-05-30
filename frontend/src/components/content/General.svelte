<script>
  import { onMount } from "svelte";
  import { GlobalStats } from "../../../wailsjs/go/main/App.js";

  // Holds the general global leaderboard stats
  let stats = {};
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
  function getGlobalStats() {
    GlobalStats().then((result) => (stats = result));
  }
  onMount(() => {
    getGlobalStats();
  });
</script>

<main>
  {#each Object.entries(stats) as stat}
    <div>
      {generalStats[stat[0]] + ": " + stat[1]}
    </div>
  {/each}
</main>
