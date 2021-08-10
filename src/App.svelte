<script>
  import { goto, Router } from "@roxi/routify"
  import { routes } from "../.routify/routes"
  import { QueryClient, QueryClientProvider } from "@sveltestack/svelte-query"
  import { t, init, getLocaleFromNavigator } from "svelte-i18n"
  import { loadValidation } from "./validation/index"
  import { initLoad } from "./locales"
  import qs from "qs"

  initLoad()

  init({
    fallbackLocale: "zh-CN",
    initialLocale: getLocaleFromNavigator(),
  })

  loadValidation($t)

  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        onError: (error) => {
          // console.log(error.config, error.response, error.request)
          // TODO add reason and redirect_uri
          if (error.response.status === 401) {
            const reason = qs.stringify({
              reason: error.response.data.error_message,
            })
            window.location.href = `/login?${reason}`
          }
        },
      },
    },
  })
</script>

<!-- <svelte:head>

  <link
    rel="stylesheet"
    href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.1/dist/css/bootstrap.min.css"
  />
  <link
    rel="stylesheet"
    href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.4.1/font/bootstrap-icons.css"
  />
</svelte:head> -->

<QueryClientProvider client={queryClient}>
  <Router {routes} />
</QueryClientProvider>

<style global>
  @import "carbon-components-svelte/css/all.css";

  /* @import "../assets/global.css"; */
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
