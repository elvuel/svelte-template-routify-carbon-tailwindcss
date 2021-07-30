<script context="module">
  import { waitLocale, init, getLocaleFromNavigator } from "svelte-i18n";
  import { initLoad } from "./index.js";
  initLoad();

  init({
    fallbackLocale: "zh-CN",
    // initialLocale: getLocaleFromNavigator(),
    initialLocale: "zh-CN",
  });

  export async function preload() {
    // awaits for the loading of the 'en-US' and 'en' dictionaries
    return waitLocale();
  }
</script>
