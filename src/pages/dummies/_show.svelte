<script>
  import {
    Grid,
    Row,
    Column,
    Form,
    FormGroup,
    TextInput,
  } from "carbon-components-svelte"

  import { useQuery } from "@sveltestack/svelte-query"
  import api from "../../api"

  export let data = {}

  let queryResult = useQuery(["users", data.id], () => api.getUser(data.id), {
    retry: (faileCount, error) => {
      if (faileCount > 3) {
        console.log(error)
      }
    },
    retryDelay: (faileCount) => {
      return Math.min(1000 * 2 ** faileCount, 30000)
    },
  })

  let message = {
    type: "active",
    message: "",
  }

  let user = {}

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
        user = $queryResult.data
        break
    }
  }

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
            bind:value={user.name}
            readonly
          />
        </FormGroup>
        <FormGroup>
          <TextInput
            placeholder="Intro"
            labelText="Intro"
            readonly
            bind:value={user.intro}
          />
        </FormGroup>
      </Form>
    </Column>
  </Row>
</Grid>
