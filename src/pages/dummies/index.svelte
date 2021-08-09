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

  import Modal from "./_modal.svelte"

  const headers = [
    { key: "id", value: "ID" },
    { key: "name", value: "Name" },
    { key: "intro", value: "Intro" },
    { key: "_action", value: "Action" },
  ]

  const queryClient = useQueryClient()

  let modalData = {
    open: false,
    modalHeading: "Modal",
    primaryButtonText: "Confirm",
    secondaryButtonText: "Cancel",
    component: "show",
  }

  let pagination = {
    totalItems: 0,
    totalPages: 0,
    size: 3,
    page: 1,
  }

  let message = {
    type: "active",
    message: "",
  }

  let selectedRowIds = []
  let rows = []

  let queryResult = useQuery(
    ["dummies", { page: pagination.page }],
    fetchDummies,
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

  function paginate() {
    queryResult = useQuery(
      ["dummies", { page: pagination.page }],
      fetchDummies,
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
  }

  function fetchDummies({ queryKey }) {
    const [_key, page] = queryKey
    return api.dummies({ page: page.page, size: pagination.size })
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
  const batchDeleteMutation = useMutation((ids) => api.deleteDummies(ids))

  function removeUser(id) {
    $deleteMutation.mutate(id, {
      onSuccess: (data, variables, context) => {
        // I will fire second!
        queryClient.invalidateQueries(["dummies", { page: pagination.page }])
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

  function removeDummies(ids) {
    $batchDeleteMutation.mutate(ids, {
      onSuccess: (data, variables, context) => {
        // I will fire second!
        queryClient.invalidateQueries(["dummies", { page: pagination.page }])
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
                removeDummies(selectedRowIds)
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
            <Button
              on:click={(e) => {
                e.preventDefault()
                modalData.modalHeading = "New Dummy"
                modalData.component = "new"
                modalData.open = true
              }}>New</Button
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
                text="Show"
                on:click={(e) => {
                  e.preventDefault()
                  modalData.modalHeading = "Show Dummy"
                  modalData.component = "show"
                  modalData.open = true
                  modalData.data = { id: row.id }
                }}
              />
              <OverflowMenuItem
                text="Edit"
                on:click={(e) => {
                  e.preventDefault()
                  modalData.modalHeading = "Edit Dummy"
                  modalData.component = "edit"
                  modalData.open = true
                  modalData.data = { id: row.id, page: pagination.page }
                }}
              />
              <OverflowMenuItem
                danger
                text="Delete"
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
      {#if modalData.open}
        <Modal
          bind:open={modalData.open}
          bind:modalHeading={modalData.modalHeading}
          bind:secondaryButtonText={modalData.secondaryButtonText}
          bind:component={modalData.component}
          bind:data={modalData.data}
        />
      {/if}
    </Column>
  </Row>
</Grid>
