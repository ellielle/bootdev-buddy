<script>
  import { User } from "../../stores/user";
  export let level;
  export let currentXP;
  export let levelXP;

  // Set stroke-dashoffset to the circumference x percentage of level completed
</script>

<main>
  <section class="exp-bar">
    {level}
    <div class="xp-background-div flex">
      <div
        class="xp-foreground-div"
        style="width: {levelXP / currentXP}%;"
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
    background-color: lime;
    border-radius: 0.5rem;
  }
  .xp-foreground-div {
    background-color: blue;
  }
</style>
