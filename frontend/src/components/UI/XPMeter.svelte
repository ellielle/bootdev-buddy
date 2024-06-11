<script>
  import { User } from "../../stores/user";
  export let level;
  export let currentXP;
  export let levelXP;

  // Set stroke-dashoffset to the circumference x percentage of level completed
</script>

<main>
  <section class="exp-bar">
    <div id="level" class="text-primary-500">
      {level}
    </div>
    <div class="xp-background-div flex bg-gray-700">
      <div
        class="xp-foreground-div bg-primary-500"
        style="width: {(currentXP / levelXP) * 100}%;"
      ></div>
    </div>
  </section>
</main>

<style>
  .exp-bar {
    display: flex;
    text-align: center;
    flex-direction: column;
  }
  .xp-background-div {
    width: 60px;
    height: 0.3rem;
    border-radius: 0.5rem;
  }
  .xp-foreground-div {
    border-radius: 0.5rem;
  }
</style>
