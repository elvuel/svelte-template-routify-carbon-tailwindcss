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
    Button,
  } from "carbon-components-svelte"
  import { useMutation, useQueryClient } from "@sveltestack/svelte-query"
  import { t } from "svelte-i18n"
  import { createUser } from "../../api"
  import { newValidator, getValidation } from "../../validation/index"
  import { goto } from "@roxi/routify"

  const queryClient = useQueryClient()

  const mutation = useMutation((data) => createUser(data), {
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

  function addUser(data) {
    $mutation.mutate(data, {
      onMutate: async (variables) => {
        // A mutation is about to happen!
      },
      onSuccess: (data, variables, context) => {
        // I will fire second!
        queryClient.invalidateQueries("users")
        resetUiControls()
        $goto("/users")
      },
      onError: (error, variables, context) => {
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
        <BreadcrumbItem href="/">Dashboard</BreadcrumbItem>
        <BreadcrumbItem href="/users">Users</BreadcrumbItem>
        <BreadcrumbItem isCurrentPage>New</BreadcrumbItem>
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
            invalid={uiCtrl.name.invalid}
            invalidText={uiCtrl.name.invalidText}
          />
        </FormGroup>
        <FormGroup>
          <TextInput
            placeholder="Intro"
            labelText="intro"
            bind:value={uiCtrl.intro.value}
            invalid={uiCtrl.intro.invalid}
            invalidText={uiCtrl.intro.invalidText}
          />
        </FormGroup>
        <Button
          type="button"
          on:click={(e) => {
            e.preventDefault()
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
          }}>Create</Button
        >
      </Form>
    </Column>
  </Row>
</Grid>
