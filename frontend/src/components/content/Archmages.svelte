<script>
  import { onMount } from "svelte";
  import { ArchmagesList } from "../../../wailsjs/go/main/App.js";

  let archmages = [];
  $: avgTimeToArchmage = 0;
  $: totalArchmages = archmages.length;
  $: if (archmages.length) {
    timeToArchmage();
  }

  // timeToArchmage counts
  function timeToArchmage() {
    if (totalArchmages == 0) {
      return 0;
    }

    // TimeToArchmage
    const TTA = archmages.reduce((total, archmage) => {
      const start = new Date(archmage.CreatedAt).getTime();
      const end = new Date(archmage.ArchmageUnlockedAt).getTime();
      const totalTime = end - start; // time in milliseconds
      return total + totalTime;
    }, 0);

    let millisecondsInADay = 1000 * 60 * 60 * 24;

    // @ts-ignore
    avgTimeToArchmage = (TTA / millisecondsInADay / totalArchmages).toFixed(2);

    return avgTimeToArchmage;
  }

  onMount(() => {
    ArchmagesList().then((result) => {
      archmages = result;
    });
    timeToArchmage();
  });
</script>

<main>
  <h2>Total Archmages: {totalArchmages}</h2>
  <p>Average Time to Archmage: {avgTimeToArchmage} Days!</p>
</main>
