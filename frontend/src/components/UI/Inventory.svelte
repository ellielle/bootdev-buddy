<script>
  import { onMount } from "svelte";
  import { User } from "../../stores/user";
  import {
    SeerStones,
    BakedSalmon,
    XPPotion,
  } from "../../../wailsjs/go/main/App.js";

  onMount(() => {
    SeerStones().then((result) => ($User.seerStones = result));
    BakedSalmon().then((result) => ($User.bakedSalmon = result));
    XPPotion().then((result) => {
      // @ts-ignore
      $User.xpPotion = result.NumUnusedXPPotion;
    });
  });
</script>

<main>
  <div class="flex items-end">
    <div class="mr-4">
      <img
        src="src/assets/images/seerstone.webp"
        alt="Seer Stone"
        style="width:25px;height:auto"
      />
      {$User.seerStones}
    </div>
    <div class="mr-4">
      <img
        src="src/assets/images/salmon.webp"
        alt="Baked Salmon"
        style="width:25px;height:auto"
      />
      {$User.bakedSalmon}
    </div>
    <div class="mr-4">
      <img
        src="src/assets/images/xppotion.webp"
        alt="XP Potions"
        style="width:25px;height:auto"
      />
      {$User.xpPotion}
    </div>
  </div>
</main>
