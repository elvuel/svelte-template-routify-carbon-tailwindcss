<script>
  import {
    Grid,
    Row,
    Breadcrumb,
    BreadcrumbItem,
    Column,
    Form,
    FormGroup,
    TextInput,
    InlineLoading,
    Button,
  } from "carbon-components-svelte"

  import { params } from "@roxi/routify"
  import {
    useQuery,
    useMutation,
    useQueryClient,
  } from "@sveltestack/svelte-query"
  import { newValidator, getValidation } from "../../../validation/index"
  import { getUser, updateUser } from "../../../api"
  import { t } from "svelte-i18n"

  const queryClient = useQueryClient()

  const queryResult = useQuery(
    ["users", $params.id],
    () => getUser($params.id),
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

  const mutation = useMutation((data) => updateUser(data.id, data.data))

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

  function modifyUser(data) {
    $mutation.mutate(data, {
      onMutate: async (variables) => {
        // A mutation is about to happen!
        message.type = "active"
        message.message = "Saving..."
      },
      onSuccess: (data, variables, context) => {
        // I will fire second!
        queryClient.invalidateQueries("users")
        resetUiControls()
        message.type = "succeed"
        message.message = "Updated"
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
        const user = $queryResult.data
        uiCtrl.name.value = user.name
        uiCtrl.intro.value = user.intro
        break
    }
  }

  $: $queryResult.status, queryStatus()
</script>

<Grid>
  <Row
    ><Column>
      <Breadcrumb>
        <BreadcrumbItem href="/">Dashboard</BreadcrumbItem>
        <BreadcrumbItem href="/users">Users</BreadcrumbItem>
        <BreadcrumbItem isCurrentPage>Edit - {$params.id}</BreadcrumbItem>
      </Breadcrumb>
    </Column>
  </Row>
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
        <Button
          type="button"
          on:click={(e) => {
            e.preventDefault()
            const payload = {
              name: uiCtrl.name.value,
              intro: uiCtrl.intro.value,
            }
            const validator = newValidator(
              getValidation().constraints.user.update,
              $t
            )
            validator
              .validate(payload)
              .then(() => {
                modifyUser({ id: $params.id, data: payload })
              })
              .catch(({ errors, fields }) => {
                errors.forEach((element) => {
                  uiCtrl[element.field].invalid = true
                  uiCtrl[element.field].invalidText = element.message
                })
              })
          }}>Update</Button
        >
        {#if message.message !== ""}
          <InlineLoading status={message.type} description={message.message} />
        {/if}
      </Form>
    </Column>
  </Row>
</Grid>
