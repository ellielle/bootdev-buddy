<script>
  import { User } from "../../stores/user";
  $: username = $User.userData.DiscordUserHandle;
</script>

<main>
  Welcome to Boot.dev Buddy, {username}!
</main>
