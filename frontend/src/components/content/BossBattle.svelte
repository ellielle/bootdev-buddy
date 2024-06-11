<script>
  import { BossBattle } from "../../../wailsjs/go/main/App.js";

  let boss = {};

  let promise = BossBattle().then((result) => (boss = result));
</script>

<main class="flex flex-col mx-auto">
  {#await promise}
    <p>Loading...</p>
  {:then battle}
    <!-- FIXME: fix condition that checks for a boss defeat || new Date(battle.Event.DefeatedAt) < new Date(Date.now())-->
    {#if new Date(battle.Event.ExpiresAt) < new Date(Date.now())}
      <!-- boss is inactive, show stats -->
      <a
        href="https://www.boot.dev/lore/{battle.Event.Boss.LoreSlug}"
        class="text-primary-500"
        target="_blank"
      >
        {battle.Event.Boss.Name}
      </a>
      has been defeated!
      <p>You gained {battle.XPUser} XP during the fight.</p>
      <p>
        The final blow was dealt on {new Date(
          battle.Event.DefeatedAt,
        ).toLocaleDateString()}.
      </p>
    {:else}
      <!-- boss is active, show a slowed feed and stats -->
      <div>
        <img
          src={battle.Event.Boss.ImageURL}
          alt="Boss fight"
          height="auto"
          width="100%"
        />
      </div>
      <h1 class="text-xl my-4 text-center">{battle.Event.Boss.Name}</h1>
      <p>
        {battle.Event.Boss.Description}
      </p>
      <h2 class="mt-8 mb-4 text-center">
        The battle against {battle.Event.Boss.Name} has begun!
      </h2>
      <section class="boss-bar">
        <div class="boss-background-div">
          <div
            class="boss-foreground-div bg-error-500"
            style="width: {100 -
              (battle.XPTotal / battle.Event.Boss.HealthPoints) * 100}%;"
          >
            <div class="xp-text align-center justify-center mb-12">
              {(
                100 -
                (battle.XPTotal / battle.Event.Boss.HealthPoints) * 100
              ).toFixed(0)}%
            </div>
          </div>
        </div>
      </section>

      Reward 1: {battle.Event.Boss.Rewards[0].UnlockedAt}
      <section>
        <div>
          Damage dealt by the arcanum: {battle.XPTotal}
        </div>
        <div>
          Damage dealt by you: {battle.XPUser} / 5000 xp
        </div>
        {#each battle.Event.Boss.Rewards as reward (reward.UUID)}
          <div>
            Reward: {reward.UnlockedAt} : {battle.XPTotal > reward.UnlockedAt
              ? "🟢"
              : "❌"}
          </div>
        {/each}
      </section>
    {/if}
  {/await}
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
