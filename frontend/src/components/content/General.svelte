<script>
  import { User } from "../../stores/user";
  $: username = $User.userData.DiscordUserHandle;
</script>

<main>
  {console.log($User)}
  Welcome to Boot.dev Buddy, {username}!
</main>
