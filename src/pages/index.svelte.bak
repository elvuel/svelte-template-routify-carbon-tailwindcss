<script>
  import { getContext } from "svelte";
  import RoutifyIntro from "./example/_components/RoutifyIntro.svelte";
  import { metatags } from "@roxi/routify";

  import { goto } from "@roxi/routify";
  metatags.title = "My Routify app";
  metatags.description = "Description coming soon...";

  const token = getContext("token");
  if (!token) {
    $goto("/login");
  }
</script>

{#if token}
  <RoutifyIntro />
{/if}
