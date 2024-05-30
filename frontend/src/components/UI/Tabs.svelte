<script>
  import Archmages from "../content/Archmages.svelte";
  import General from "../content/General.svelte";
  import { Tab, TabGroup, TabAnchor } from "@skeletonlabs/skeleton";

  /** @type number */
  let tabSet = 0;

  let tabs = ["General", "Courses", "Stats", "Boss Battle", "Archmages"];

  // TODO: disable highlighting text on tabs, it disrupts clicking
  // - give more margin at the top, in App.svelte
  // - find out why it's not showing up at all
</script>

<main>
  <TabGroup>
    {#each tabs as option, index}
      <Tab bind:group={tabSet} name={option} value={index}>
        <span style="color: hotpink;">{option}</span>
      </Tab>
    {/each}

    <!-- tab panels -->
    <svelte:fragment slot="panel">
      {#if tabSet === 0}
        <General />
      {/if}
      {#if tabSet === 1}
        <Archmages />
      {/if}
      {#if tabSet === 2}
        <Archmages />
      {/if}
      {#if tabSet === 3}
        <Archmages />
      {/if}
      {#if tabSet === 4}
        <Archmages />
      {/if}
    </svelte:fragment>
  </TabGroup>
</main>
