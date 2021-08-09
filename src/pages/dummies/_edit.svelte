<script>
  import {
    Grid,
    Row,
    Column,
    Form,
    FormGroup,
    TextInput,
    InlineLoading,
    Button,
  } from "carbon-components-svelte"

  import {
    useQuery,
    useMutation,
    useQueryClient,
  } from "@sveltestack/svelte-query"
  import { newValidator, getValidation } from "../../validation/index"
  import api from "../../api"
  import { t } from "svelte-i18n"
  import { onMount, onDestroy } from "svelte"

  const queryClient = useQueryClient()
  export let data = {}

  const queryResult = useQuery(
    ["dummies", data.id],
    () => api.getDummy(data.id),
    {
      retry: (faileCount, error) => {
        if (faileCount > 3) {
          console.log(error)
        }
      },
      retryDelay: (faileCount) => {
        return Math.min(1000 * 2 ** faileCount, 30000)
      },
    }
  )

  const mutation = useMutation((data) => api.updateDummy(data.id, data.data))

  let message = {
    type: "active",
    message: "",
  }

  let uiCtrl = {
    name: {
      value: "",
      invalid: false,
      invalidText: "",
    },
    intro: {
      value: "",
      invalid: false,
      invalidText: "",
    },
  }

  function resetUiControls() {
    uiCtrl.name.invalid = false
    uiCtrl.name.invalidText = ""
    uiCtrl.name.value = ""

    uiCtrl.intro.invalid = false
    uiCtrl.intro.invalidText = ""
    uiCtrl.intro.value = ""
  }

  function modifyUser(payload, e) {
    $mutation.mutate(payload, {
      onMutate: async (variables) => {
        // A mutation is about to happen!
        message.type = "active"
        message.message = "Saving..."
      },
      onSuccess: (_data, variables, context) => {
        // I will fire second!
        queryClient.invalidateQueries(["dummies", { page: data.page }])
        queryClient.invalidateQueries(["dummies", data.id])
        resetUiControls()
        message.type = "succeed"
        message.message = "Updated"
        e.detail.close()
      },
      onError: (error, variables, context) => {
        message.type = "error"
        message.message = error
        // I will fire second!
        setTimeout(() => {
          $mutation.reset()
        }, 5000)
      },
      onSettled: (data, error, variables, context) => {
        // I will fire second!
      },
      retry: 0,
    })
  }

  function queryStatus() {
    switch ($queryResult.status) {
      case "loading":
        message.type = "active"
        message.message = "Loading..."
        break
      case "idle":
        message.type = "active"
        message.message = "idle"
      case "error":
        message.type = "error"
        message.message = $queryResult.error.message
        break
      case "success":
        message.type = "success"
        message.message = ""
        const item = $queryResult.data
        uiCtrl.name.value = item.name
        uiCtrl.intro.value = item.intro
        break
    }
  }

  function submit(e) {
    const payload = {
      name: uiCtrl.name.value,
      intro: uiCtrl.intro.value,
    }
    const validator = newValidator(getValidation().constraints.user.update, $t)
    validator
      .validate(payload)
      .then(() => {
        modifyUser({ id: data.id, data: payload }, e)
      })
      .catch(({ errors, fields }) => {
        errors.forEach((element) => {
          uiCtrl[element.field].invalid = true
          uiCtrl[element.field].invalidText = element.message
        })
      })
  }

  let me
  onMount(() => {
    window.addEventListener("submit-for-edit", submit)
  })

  onDestroy(() => {
    window.removeEventListener("submit-for-edit", submit)
  })

  $: $queryResult.status, queryStatus()
</script>

<Grid>
  <Row>
    <Column>
      <Form>
        <FormGroup>
          <TextInput
            placeholder="Name"
            labelText="Name"
            bind:value={uiCtrl.name.value}
          />
        </FormGroup>
        <FormGroup>
          <TextInput
            placeholder="Intro"
            labelText="Intro"
            bind:value={uiCtrl.intro.value}
          />
        </FormGroup>
        {#if message.message !== ""}
          <InlineLoading status={message.type} description={message.message} />
        {/if}
      </Form>
    </Column>
  </Row>
</Grid>
