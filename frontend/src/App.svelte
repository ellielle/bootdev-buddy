<script>
  import { GetArchmagesList } from "../wailsjs/go/main/App.js";

  let archmages = [];

  function getArchmages() {
    GetArchmagesList().then((result) => (archmages = result));
  }
</script>

<main>
  <div class="result" id="result">{archmages}</div>
  <div>
    <button class="btn" on:click={getArchmages}>Get Last 5 Archmages</button>
  </div>
</main>

<style>
</style>
