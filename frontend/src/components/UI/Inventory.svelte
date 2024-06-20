<script>
  import { onMount } from "svelte";
  import { User } from "../../stores/user";
  import {
    SeerStones,
    BakedSalmon,
    XPPotion,
    FrozenFlames,
  } from "../../../wailsjs/go/main/App.js";

  onMount(() => {
    // Call all inventory-related functions and load it up
    SeerStones().then((result) => ($User.seerStones = result));
    BakedSalmon().then((result) => ($User.bakedSalmon = result));
    XPPotion().then((result) => {
      // @ts-ignore
      $User.xpPotion = result.NumUnusedXPPotion;
    });
    FrozenFlames().then((result) => ($User.frozenFlame = result));
  });
</script>

<main>
  <div class="flex items-center">
    <div class="mr-4 flex">
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
