<script>
  import Archmages from "../content/Archmages.svelte";
  import BossBattle from "../content/BossBattle.svelte";
  import General from "../content/General.svelte";
  import XPMeter from "../UI/XPMeter.svelte";
  import Avatar from "../UI/Avatar.svelte";
  import Stats from "../content/Stats.svelte";
  import Courses from "../content/Courses.svelte";
  import { Tab, TabGroup } from "@skeletonlabs/skeleton";
  import { User } from "../../stores/user";
  import { LogoutUser, CloseApp } from "../../../wailsjs/go/main/App.js";

  /** @type number */
  let tabSet = 0;
  let tabs = ["General", "Courses", "Stats", "Boss Battle", "Archmages"];
  let result;

  function handleSignout() {
    LogoutUser().then((res) => (result = res));
    $User.isLoggedIn = false;
    $User.isArchmage = false;
    $User.userData = {};
  }

  function handleClose() {
    CloseApp();
  }
</script>

<main>
  <TabGroup>
    {#each tabs as option, index}
      <Tab
        bind:group={tabSet}
        name={option}
        value={index}
        active="border-b-2 border-r-2 border-gray-500"
        rounded=""
      >
        <span>{option}</span>
      </Tab>
    {/each}

    <!-- tab panels -->
    <svelte:fragment slot="panel">
      {#if tabSet === 0}
        <General />
      {/if}
      {#if tabSet === 1}
        <Courses />
      {/if}
      {#if tabSet === 2}
        <Stats />
      {/if}
      {#if tabSet === 3}
        <BossBattle />
      {/if}
      {#if tabSet === 4}
        <Archmages />
      {/if}
    </svelte:fragment>

    <!-- TODO: shove xp bar, avatar and logout button to the right -->

    <!-- user profile and level -->
    <div style="display: flex;">
      {#if $User.isLoggedIn}
        <XPMeter
          bind:level={$User.userData.Level}
          bind:currentXP={$User.userData.XPForLevel}
          levelXP={$User.userData.XPTotalForLevel}
        />
        <Avatar
          bind:image={$User.userData.ProfileImageURL}
          bind:handle={$User.userData.Handle}
        />
      {/if}
    </div>

    <!-- sign out button -->
    <button on:click={handleSignout}>Sign out</button>

    <button on:click={handleClose} class="absolute top-5% right-5">✗</button>
  </TabGroup>
</main>
