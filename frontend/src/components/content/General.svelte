<script>
  import { User } from "../../stores/user";
  import Inventory from "../UI/Inventory.svelte";
</script>

<main>
  Welcome to Boot.dev Buddy, {$User.userData.DiscordUserHandle}!
  <Inventory />
</main>
