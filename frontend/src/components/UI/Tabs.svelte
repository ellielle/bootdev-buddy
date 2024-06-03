<script>
  import Archmages from "../content/Archmages.svelte";
  import General from "../content/General.svelte";
  import XPMeter from "../UI/XPMeter.svelte";
  import Stats from "../content/Stats.svelte";
  import { Tab, TabGroup } from "@skeletonlabs/skeleton";
  import { User } from "../../stores/user";

  /** @type number */
  let tabSet = 0;
  let tabs = ["General", "Courses", "Stats", "Boss Battle", "Archmages"];

  // Setting some convenience variables so it doesn't look ugly
  let level = $User.userData.Level;
  let currentXP = $User.userData.XPForLevel;
  let levelXP = $User.userData.XPTotalForLevel;
</script>

<main>
  <TabGroup>
    {#each tabs as option, index}
      <Tab
        bind:group={tabSet}
        name={option}
        value={index}
        active="border-b-2 border-r-2 border-primary-400-500-token"
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
        <!-- <Courses /> -->
      {/if}
      {#if tabSet === 2}
        <Stats />
      {/if}
      {#if tabSet === 3}
        <!-- <BossBattle /> -->
      {/if}
      {#if tabSet === 4}
        <Archmages />
      {/if}
    </svelte:fragment>
    {#if $User.isLoggedIn}
      <XPMeter {level} {currentXP} {levelXP} />
    {/if}
  </TabGroup>
</main>
