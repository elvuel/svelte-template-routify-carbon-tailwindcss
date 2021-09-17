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
  import { useMutation, useQueryClient } from "@sveltestack/svelte-query"
  import { t } from "svelte-i18n"
  import api from "../../api"
  import { newValidator, getValidation } from "../../validation/index"
  import { goto } from "@roxi/routify"
  import Add16 from "carbon-icons-svelte/lib/Add16"

  const queryClient = useQueryClient()

  const mutation = useMutation((data) => api.createUser(data), {
    // onMutate: async (variables) => {
    //   // A mutation is about to happen!
    // },
    // onSuccess: (data, variables, context) => {
    //   // I will fire first
    // },
    // onError: (error, variables, context) => {
    //   // I will fire first
    // },
    // onSettled: (data, error, variables, context) => {
    //   // I will fire first
    // },
  })

  let message = {
    type: "active",
    message: "",
  }

  function addUser(data) {
    $mutation.mutate(data, {
      onMutate: async (variables) => {
        // A mutation is about to happen!
        message.type = "active"
        message.message = "Saving..."
      },
      onSuccess: (data, variables, context) => {
        // I will fire second!
        queryClient.invalidateQueries(["users", { page: 1 }])
        resetUiControls()
        message.type = "success"
        message.message = "Succeed"
        $goto("/users")
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
</script>

<Grid>
  <Row
    ><Column>
      <Breadcrumb>
        <BreadcrumbItem href="/">{$t("page.dashboard.default")}</BreadcrumbItem>
        <BreadcrumbItem href="/users">{$t("page.user.list")}</BreadcrumbItem>
        <BreadcrumbItem isCurrentPage
          >{$t("page.general.action.new")}</BreadcrumbItem
        >
      </Breadcrumb>
    </Column>
  </Row>
  <Row>
    <Column>
      <Form>
        <FormGroup>
          <TextInput
            placeholder={$t("model.user.name")}
            labelText={$t("model.user.name")}
            bind:value={uiCtrl.name.value}
            invalid={uiCtrl.name.invalid}
            invalidText={uiCtrl.name.invalidText}
          />
        </FormGroup>
        <FormGroup>
          <TextInput
            placeholder={$t("model.user.intro")}
            labelText={$t("model.user.intro")}
            bind:value={uiCtrl.intro.value}
            invalid={uiCtrl.intro.invalid}
            invalidText={uiCtrl.intro.invalidText}
          />
        </FormGroup>
        <Button
          type="button"
          icon={Add16}
          on:click={(e) => {
            e.preventDefault()
            const payload = {
              name: uiCtrl.name.value,
              intro: uiCtrl.intro.value,
            }
            const validator = newValidator(
              getValidation().constraints.user.create,
              $t
            )
            validator
              .validate(payload)
              .then(() => {
                addUser(payload)
              })
              .catch(({ errors, fields }) => {
                errors.forEach((element) => {
                  uiCtrl[element.field].invalid = true
                  uiCtrl[element.field].invalidText = element.message
                })
              })
          }}
          >{$t("page.general.action.create")}
        </Button>
        {#if message.message !== ""}
          <InlineLoading status={message.type} description={message.message} />
        {/if}
      </Form>
    </Column>
  </Row>
</Grid>
