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

  import {
    useQuery,
    useMutation,
    useQueryClient,
  } from "@sveltestack/svelte-query"
  import api from "../../api/index"
  import { url } from "@roxi/routify"
  import { t } from "svelte-i18n"

  const headers = [
    { key: "id", value: $t("model.user.id") },
    { key: "name", value: $t("model.user.name") },
    { key: "intro", value: $t("model.user.intro") },
    { key: "_action", value: $t("page.general.action.default") },
  ]

  const queryClient = useQueryClient()

  let pagination = {
    totalItems: 0,
    totalPages: 0,
    size: 2,
    page: 1,
  }

  let message = {
    type: "active",
    message: "",
  }

  let selectedRowIds = []
  let rows = []

  let queryResult = useQuery(["users", { page: pagination.page }], fetchUsers, {
    retry: (faileCount, error) => {
      if (faileCount > 3) {
        console.log(error)
      }
    },
    retryDelay: (faileCount) => {
      return Math.min(1000 * 2 ** faileCount, 30000)
    },
  })

  function paginate() {
    queryResult = useQuery(["users", { page: pagination.page }], fetchUsers, {
      retry: (faileCount, error) => {
        if (faileCount > 3) {
          console.log(error)
        }
      },
      retryDelay: (faileCount) => {
        return Math.min(1000 * 2 ** faileCount, 30000)
      },
    })
  }

  function fetchUsers({ queryKey }) {
    const [_key, page] = queryKey
    return api.users({ page: page.page, size: pagination.size })
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
        rows = $queryResult.data.items
        pagination.totalItems = $queryResult.data.totalItems
        pagination.totalPages = $queryResult.data.totalPages
        pagination.page = $queryResult.data.page
        pagination.size = $queryResult.data.size
        break
    }
  }
  const deleteMutation = useMutation((id) => api.deleteUser(id))
  const batchDeleteMutation = useMutation((ids) => api.deleteUsers(ids))

  function removeUser(id) {
    $deleteMutation.mutate(id, {
      onSuccess: (data, variables, context) => {
        // I will fire second!
        queryClient.invalidateQueries(["users", { page: pagination.page }])
      },
      onError: (error, variables, context) => {
        // I will fire second!
        setTimeout(() => {
          $deleteMutation.reset()
        }, 5000)
      },
      retry: 0,
    })
  }

  function removeUsers(ids) {
    $batchDeleteMutation.mutate(ids, {
      onSuccess: (data, variables, context) => {
        // I will fire second!
        queryClient.invalidateQueries(["users", { page: pagination.page }])
        selectedRowIds = []
      },
      onError: (error, variables, context) => {
        // I will fire second!
        setTimeout(() => {
          $batchDeleteMutation.reset()
        }, 5000)
      },
      retry: 0,
    })
  }

  $: $queryResult.status, queryStatus()
</script>

<Grid>
  <Row>
    <Column>
      <DataTable
        batchSelection
        zebra
        title={$t("model.user.$name") + $t("page.general.list")}
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
                removeUsers(selectedRowIds)
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
            <Button href={$url("./new")}>{$t("page.general.action.new")}</Button
            >
          </ToolbarContent>
        </Toolbar>
        {#if message.message !== ""}
          <InlineLoading status={message.type} description={message.message} />
        {/if}
        <span slot="cell" let:cell let:row>
          {#if cell.key === "_action"}
            <OverflowMenu flipped>
              <OverflowMenuItem
                href={$url(`./${row.id}`)}
                text={$t("page.general.action.show")}
              />
              <OverflowMenuItem
                href={$url(`./${row.id}/edit`)}
                text={$t("page.general.action.edit")}
              />
              <OverflowMenuItem
                danger
                text={$t("page.general.action.delete")}
                on:click={(e) => {
                  e.preventDefault()
                  removeUser(row.id)
                }}
              />
            </OverflowMenu>
          {:else}
            {cell.value}
          {/if}
        </span>
      </DataTable>
      {#if pagination.totalItems > 0}
        <Pagination
          pageSizeInputDisabled
          totalItems={pagination.totalItems}
          pageSizes={[2, 4, 6, 8]}
          pageSize={pagination.size}
          on:update={(e) => {
            const { pageSize, page } = e.detail
            pagination.page = page
            pagination.size = pageSize
            paginate()
          }}
        />
      {/if}
    </Column>
  </Row>
</Grid>
