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

  import { params } from "@roxi/routify"
  import { useQuery } from "@sveltestack/svelte-query"
  import api from "../../../api"

  let queryResult = useQuery(
    ["users", $params.id],
    () => api.getUser($params.id),
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
  <Row
    ><Column>
      <Breadcrumb>
        <BreadcrumbItem href="/">Dashboard</BreadcrumbItem>
        <BreadcrumbItem href="/users">Users</BreadcrumbItem>
        <BreadcrumbItem isCurrentPage>{$params.id}</BreadcrumbItem>
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
