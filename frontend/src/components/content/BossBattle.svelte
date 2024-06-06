<script>
  import { BossBattle } from "../../../wailsjs/go/main/App.js";

  let boss = {};

  let promise = BossBattle().then((result) => (boss = result));
</script>

<main>
  <p>Boss Battles!</p>
  {#await promise}
    <p>Loading...</p>
  {:then b}
    <p>{b.Event}</p>
  {/await}
</main>
