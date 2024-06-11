<script>
  import { User } from "../../stores/user";
</script>

<main>
  {console.log($User.userData)}
  Welcome to Boot.dev Buddy, {$User.userData.DiscordUserHandle}!
</main>
