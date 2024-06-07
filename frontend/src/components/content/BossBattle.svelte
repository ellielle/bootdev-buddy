<script>
  import { BossBattle } from "../../../wailsjs/go/main/App.js";

  let boss = {};

  let promise = BossBattle().then((result) => (boss = result));
</script>

<main>
  {#await promise}
    <p>Loading...</p>
  {:then battle}
    {#if new Date(battle.Event.ExpiresAt) > new Date(Date.now()) || new Date(battle.Event.DefeatedAt) < new Date(Date.now())}
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
      <!-- TODO: Add active boss battle stats -->
    {/if}
  {/await}
</main>
