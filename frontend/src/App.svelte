<script>
  import { GetArchmagesList } from "../wailsjs/go/main/App.js";

  let archmages = [];
  let isArchmage = false;

  function getArchmages() {
    archmages = [];
    GetArchmagesList().then((result) => (archmages = result.slice(0, 10)));
  }
</script>

<main>
  <div class="result" id="result"></div>
  <div>
    <button class="btn" on:click={getArchmages}>Get Last 5 Archmages</button>
  </div>
  <div>
    {#each archmages as mage}
      <div>
        {mage?.Handle || "hi"}
      </div>
    {/each}
  </div>
</main>

<style>
</style>
