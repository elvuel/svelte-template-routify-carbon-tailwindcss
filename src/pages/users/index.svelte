<script>
  import {
    Grid,
    Row,
    Column,
    DataTable,
    InlineLoading,
    Toolbar,
    ToolbarBatchActions,
    ToolbarContent,
    ToolbarSearch,
    ToolbarMenu,
    ToolbarMenuItem,
    Pagination,
    OverflowMenu,
    OverflowMenuItem,
    Button,
  } from "carbon-components-svelte"
  import Delete16 from "carbon-icons-svelte/lib/Delete16"

  import { useQuery } from "@sveltestack/svelte-query"
  import { users as fetchUsers } from "../../api/index"
  import { url } from "@roxi/routify"
  import { t } from "svelte-i18n"

  const headers = [
    { key: "id", value: "ID" },
    { key: "name", value: "Name" },
    { key: "intro", value: "Intro" },
    { key: "_action", value: "Action" },
  ]

  let queryResult = useQuery("users", fetchUsers, {
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

  let selectedRowIds = []
  let rows = []

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
        rows = $queryResult.data
        break
    }
  }

  $: $queryResult.status, queryStatus()
</script>

<Grid>
  <Row>
    <Column>
      <DataTable
        batchSelection
        zebra
        title="Load balancers"
        description="Your organization's active load balancers."
        bind:selectedRowIds
        {headers}
        {rows}
      >
        <Toolbar>
          <ToolbarBatchActions
            formatTotalSelected={(totalSelected) => {
              return `总计选中: ${totalSelected} 项`
            }}
          >
            <Button
              icon={Delete16}
              on:click={(e) => {
                e.preventDefault()
                console.log(selectedRowIds)
              }}>Destory</Button
            >
          </ToolbarBatchActions>
          <ToolbarContent>
            <ToolbarSearch />
            <ToolbarMenu>
              <ToolbarMenuItem primaryFocus>Restart all</ToolbarMenuItem>
              <ToolbarMenuItem
                href="https://cloud.ibm.com/docs/loadbalancer-service"
              >
                API documentation
              </ToolbarMenuItem>
              <ToolbarMenuItem danger>Stop all</ToolbarMenuItem>
            </ToolbarMenu>
            <Button href={$url("./new")}>New</Button>
          </ToolbarContent>
        </Toolbar>
        {#if message.message !== ""}
          <InlineLoading status={message.type} description={message.message} />
        {/if}
        <span slot="cell" let:cell>
          {#if cell.key === "_action"}
            <OverflowMenu flipped>
              <OverflowMenuItem text="Edit" />
              <OverflowMenuItem href="/users" text="Edit" />
              <OverflowMenuItem danger text="Delete" />
            </OverflowMenu>
          {:else}
            {cell.value}
          {/if}
        </span>
      </DataTable>
      <Pagination totalItems={102} pageSizes={[16, 36, 99]} pageSize={36} />
    </Column>
  </Row>
</Grid>
