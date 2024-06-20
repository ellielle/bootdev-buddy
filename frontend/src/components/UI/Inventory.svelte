<script>
  import { onMount } from "svelte";
  import { User } from "../../stores/user";
  import {
    SeerStones,
    BakedSalmon,
    XPPotion,
    FrozenFlames,
  } from "../../../wailsjs/go/main/App.js";

  function calculateXPTimer() {
    let endsAt = new Date(endTime);
    // @ts-ignore
    let timeLeft = endsAt - new Date(Date.now());

    let minutes = Math.floor((timeLeft % (1000 * 60 * 60)) / (1000 * 60));
    let seconds = Math.floor((timeLeft % (1000 * 60)) / 1000);

    timer = `${minutes}:${seconds}`;
  }

  let timer = "";
  let endTime;

  onMount(() => {
    // Call all inventory-related functions and load it up
    SeerStones().then((result) => ($User.seerStones = result));
    BakedSalmon().then((result) => ($User.bakedSalmon = result));
    XPPotion().then((result) => {
      // Using ts-ignore here as my IDE doesn't like the properties on result
      // @ts-ignore
      $User.xpPotion = result.NumUnusedXPPotion;
      // @ts-ignore
      // last entry holds the "newest" xp potion, which has
      // the ExpiresAt for when the buff drops off
      const timerLen = result.ActiveXPPotions.length - 1;
      endTime = result.ActiveXPPotions[timerLen].ExpiresAt;
      // @ts-ignore
      $User.xpTimer = calculateXPTimer();
    });
    FrozenFlames().then((result) => ($User.frozenFlame = result));

    // If there is a buff timer, display the countdown
    const interval = setInterval(calculateXPTimer, 1000);

    return () => {
      clearInterval(interval);
    };
  });
</script>

<main>
  <div class="flex items-center">
    <div id="timer">
      {timer !== "" && !timer.includes("NaN") ? timer : "00:00"}
    </div>
    <div class="mx-4 flex">
      <img
        src="src/assets/images/xppotion.webp"
        alt="XP Potions"
        style="width:25px;height:auto"
      />
      {$User.xpPotion}
    </div>
    <div class="mr-4 flex">
      <img
        src="src/assets/images/seerstone.webp"
        alt="Seer Stone"
        style="width:25px;height:auto"
      />
      {$User.seerStones}
    </div>
    <div class="mr-4 flex">
      <img
        src="src/assets/images/salmon.webp"
        alt="Baked Salmon"
        style="width:25px;height:auto"
      />
      {$User.bakedSalmon}
    </div>
    <div class="mr-4 flex">
      <img
        src="src/assets/images/flame.png"
        alt="XP Potions"
        style="width:25px;height:auto"
      />
      {$User.frozenFlame}
    </div>
  </div>
</main>
