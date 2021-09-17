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
  import { useQuery } from "@sveltestack/svelte-query"
  import api from "../../../api"
  import { t } from "svelte-i18n"

  let queryResult = useQuery(
    ["users", $params.id],
    () => api.getUser($params.id),
    {
      onError: (error) => {
        console.log(error)
      },
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
        <BreadcrumbItem href="/">{$t("page.dashboard.default")}</BreadcrumbItem>
        <BreadcrumbItem href="/users">{$t("page.user.list")}</BreadcrumbItem>
        <BreadcrumbItem isCurrentPage>{$params.id}</BreadcrumbItem>
      </Breadcrumb>
    </Column>
  </Row>
  <Row>
    {#if message.message !== ""}
      <InlineLoading status={message.type} description={message.message} />
    {/if}
    <Column>
      <Form>
        <FormGroup>
          <TextInput
            placeholder={$t("model.user.name")}
            labelText={$t("model.user.name")}
            bind:value={user.name}
            readonly
          />
        </FormGroup>
        <FormGroup>
          <TextInput
            placeholder={$t("model.user.intro")}
            labelText={$t("model.user.intro")}
            readonly
            bind:value={user.intro}
          />
        </FormGroup>
      </Form>
    </Column>
  </Row>
</Grid>
