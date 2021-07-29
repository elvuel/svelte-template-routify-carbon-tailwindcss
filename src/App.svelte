<script>
  import { Router } from "@roxi/routify";
  import { routes } from "../.routify/routes";
</script>

<svelte:head>
  <link
    rel="stylesheet"
    href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.1/dist/css/bootstrap.min.css"
  />
  <link
    rel="stylesheet"
    href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.4.1/font/bootstrap-icons.css"
  />
</svelte:head>

<Router {routes} />

<style global>
  @import "carbon-components-svelte/css/white.css";

  @import "../assets/global.css";
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
