<script>
  import { onMount } from "svelte"
  import { goto } from "@roxi/routify"
  // if (window.location.pathname == "/users/10") {
  //   $goto("/login")
  // }
  // if (window.location.pathname == "/users/10") {
  //   window.location.href = "/login"
  // }
  // onMount(() => {
  //   console.log(location.pathname)
  // })
</script>

<slot />
