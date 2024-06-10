<script>
  import { onMount } from "svelte";
  import { Courses, CoursesProgress } from "../../../wailsjs/go/main/App.js";
  const COURSES_IN_ORDER = [
    { UUID: "f9a25dfb-3e00-4727-ac78-36de82315355", slug: "learn-python" },
    {
      UUID: "bc7a07ef-ab87-42ab-80de-e7261f2c58a0",
      slug: "learn-shells-and-terminals",
    },
    { UUID: "933d6dd0-b21a-488e-8ece-469bbef28652", slug: "learn-git" },
    { UUID: "094fd7d4-ec78-4202-96ca-c5f79fc332d2", slug: "build-bookbot" },
    {
      UUID: "f9a48bbc-d1ff-4388-bf0c-23c6e3c60ae0",
      slug: "learn-objct-oriented-programming",
    },
    {
      UUID: "b1459f0c-21eb-41e5-b7f3-562ef69d344c",
      slug: "learn-functional-programming",
    },
    {
      UUID: "d38e78e9-ae52-458e-8494-ec7ecbdab14f",
      slug: "build-static-site-generator",
    },
    { UUID: "884342fc-5469-47b4-8125-8bfc897428a8", slug: "learn-algorithms" },
    {
      UUID: "7bbb53ed-2106-4f6b-b885-e7645c2ff9d8",
      slug: "learn-data-structures",
    },
    {
      UUID: "2b266bb4-2262-49c0-b6d1-75cd8c5e8be8",
      slug: "build-maze-solver-python",
    },
    {
      UUID: "273215af-23f0-4461-a5d8-617c7f19127c",
      slug: "build-personal-project-1",
    },
    { UUID: "2af5c197-21eb-48b4-bd90-b0d59adb311e", slug: "learn-javascript" },
    { UUID: "5d804c54-887a-4c1c-b8c7-b6436f3a132e", slug: "learn-http" },
    {
      UUID: "59fbb2aa-7d67-4e88-bac8-42f49798a9f5",
      slug: "build-web-crawler-javascript",
    },
    {
      UUID: "70c1c7e6-15af-4396-b2bd-3e1731a50bef",
      slug: "build-personal-project-2",
    },
    { UUID: "3b39d0f6-f944-4f1b-832d-a1daba32eda4", slug: "learn-golang" },
    { UUID: "b6ac3462-d76f-453b-bd5d-5d7fe07cdadb", slug: "build-pokedex-cli" },
    { UUID: "81b7293c-60aa-40c7-a158-7c87428f6031", slug: "learn-web-servers" },
    { UUID: "bc0dc34b-025a-4d97-b7a0-382aa21533aa", slug: "learn-sql" },
    {
      UUID: "3a8d6b13-c064-424d-bd09-5e09ceaddfea",
      slug: "build-blog-aggregator",
    },
    { UUID: "2d740eb6-3234-419e-9a23-08ec9e9889b7", slug: "learn-docker" },
    { UUID: "1d057329-0a4a-486f-a8cc-2fccada6b307", slug: "learn-ci-cd" },
    {
      UUID: "25a0e482-fd55-463c-80b6-4fbf789fcef5",
      slug: "build-capstone-project",
    },
    { UUID: "0617f271-09a6-4b1c-ab59-d0ec22f53153", slug: "learn-job-search" },
    { UUID: "6e6236f7-6f6b-45e3-859a-5fd0084754aa", slug: "learn-kubernetes" },
    { UUID: "93174165-cfaf-4201-a5b6-7da1864c9792", slug: "learn-pub-sub" },
    {
      UUID: "6321ddbf-49eb-4748-9737-6bc12e8bb705",
      slug: "learn-cryptography",
    },
    {
      UUID: "aaad49fb-0dc5-43c6-992c-96d3f83ee663",
      slug: "learn-advanced-algorithms",
    },
  ];
  let courses = [];
  let sortedCourses = [];
  let init = false;
  let progress = {};

  // Sort courses by how they appear on the website
  // once the courses array is populated
  // Then set init to true to swap from loading state
  $: if (courses.length > 0) {
    for (let i = 0; i < COURSES_IN_ORDER.length; i++) {
      for (let j = 0; j < courses.length; j++) {
        if (COURSES_IN_ORDER[i].UUID === courses[j].UUID) {
          // For svelte reactivity...look at documentation for alternative
          sortedCourses.push(courses[j]);
        }
      }
    }
  }

  $: if (Object.keys(progress).length > 0) {
    init = true;

    console.log("progress: ", progress.Progress);
  }

  onMount(() => {
    Courses().then((result) => (courses = [...result]));
    CoursesProgress().then((result) => (progress = result));
  });
</script>

<main>
  {#if !init}
    Loading...
  {:else}
    {#each sortedCourses as course (course.UUID)}
      <div>
        <a
          href="https://www.boot.dev/lessons/{course.slug}"
          class="text-primary-500"
        >
          {course.Title}:
        </a>
        {progress.Progress[course.UUID].NumDone}/{progress.Progress[course.UUID]
          .NumMax}
      </div>
    {/each}
  {/if}
</main>
