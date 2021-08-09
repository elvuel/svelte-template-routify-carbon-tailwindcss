<script>
  import { Modal } from "carbon-components-svelte"

  import Show from "./_show.svelte"
  import Edit from "./_edit.svelte"
  import New from "./_new.svelte"

  export let open = false
  export let modalHeading = "Modal"
  export let primaryButtonText = "Confirm"
  export let secondaryButtonText = "Cancel"
  export let component = "show"

  export let data = undefined

  const components = {
    show: Show,
    edit: Edit,
    new: New,
  }

  function closeModal() {
    open = false
  }

  function openModal() {
    open = true
  }
</script>

<Modal
  bind:open
  bind:modalHeading
  bind:primaryButtonText
  bind:secondaryButtonText
  on:click:button--secondary={() => {
    closeModal()
  }}
  on:open
  on:close
  on:submit={(e) => {
    e.preventDefault()
    const evt = `submit-for-${component}`
    window.dispatchEvent(
      new window.CustomEvent(evt, {
        detail: { close: closeModal, open: openModal },
      })
    )
  }}
>
  <svelte:component this={components[component]} bind:data />
</Modal>
