<script>
  import { onMount } from "svelte";
  import { BossBattle } from "../../../wailsjs/go/main/App.js";

  let boss = {};

  function fetchBossUpdate() {
    BossBattle().then((result) => (boss = result));
  }

  onMount(() => {
    BossBattle().then((result) => (boss = result));
    const refreshInterval = setInterval(() => {
      fetchBossUpdate();
    }, 5000);

    return () => {
      clearInterval(refreshInterval);
    };
  });
</script>

<!-- BUG: the issue is most likely that you're using await / then in svelte which have their own scope, that isn't being watched or used in a reactive way. So boss data doesn't get updated. use `onMount` like in other components -->

<main class="flex flex-col mx-auto">
  {#if Object.keys(boss).length === 0}
    <p>Loading...</p>
  {:else}
    <!-- FIXME: fix condition that checks for a boss defeat -->
    {#if new Date(boss.Event.ExpiresAt) < new Date(Date.now())}
      <!-- boss is inactive, show stats -->
      <div>
        <a
          href="https://www.boot.dev/lore/{boss.Event.Boss.LoreSlug}"
          class="text-primary-500"
          target="_blank"
        >
          {boss.Event.Boss.Name}
        </a>
        has been defeated!
      </div>
      <p>You gained {boss.XPUser} XP during the fight.</p>
      <p>
        The final blow was dealt on {new Date(
          boss.Event.DefeatedAt,
        ).toLocaleDateString()}.
      </p>
    {:else}
      <!-- boss is active, show a slowed feed and stats -->
      <div>
        <img
          src={boss.Event.Boss.ImageURL}
          alt="Boss fight"
          height="auto"
          width="100%"
        />
      </div>
      <h1 class="text-xl my-4 text-center">{boss.Event.Boss.Name}</h1>
      <p>
        {boss.Event.Boss.Description}
      </p>
      <h2 class="mt-8 mb-4 text-center">
        The battle against {boss.Event.Boss.Name} has begun!
      </h2>
      <section class="boss-bar">
        <div class="boss-background-div">
          <div
            class="boss-foreground-div bg-error-500"
            style="width: {100 -
              (boss.XPTotal / boss.Event.Boss.HealthPoints) * 100}%;"
          >
            <div class="xp-text align-center justify-center mb-12">
              {(
                100 -
                (boss.XPTotal / boss.Event.Boss.HealthPoints) * 100
              ).toFixed(0)}%
            </div>
          </div>
        </div>
      </section>

      <section>
        <div>
          Damage dealt by the arcanum: {boss.XPTotal}
        </div>
        <div>
          Damage dealt by you: {boss.XPUser} / 5000 required to qualify.
        </div>
        {#each boss.Event.Boss.Rewards as reward (reward.UUID)}
          <div>
            Reward at {reward.UnlockedAt} damage: {boss.XPTotal >
            reward.UnlockedAt
              ? "✓"
              : "✗"}
          </div>
        {/each}
      </section>
    {/if}
  {/if}
</main>

<style>
  .boss-bar {
    display: flex;
    text-align: center;
    flex-direction: column;
    font-weight: 600;
  }
  .boss-background-div {
    border-radius: 0.5rem;
  }
  .boss-foreground-div {
    border-radius: 0.5rem;
  }
</style>
