<script>
  import { userQuery } from "@sveltestack/svelte-query";
  import { users as fetchUsers } from "../../api/index";
  import { t } from "svelte-i18n";

  import {
    Grid,
    Row,
    Column,
    DataTable,
    Toolbar,
    ToolbarContent,
    ToolbarSearch,
    ToolbarMenu,
    ToolbarMenuItem,
    Pagination,
    OverflowMenu,
    OverflowMenuItem,
    Button,
  } from "carbon-components-svelte";
</script>

<Grid>
  <Row>
    <Column>
      <DataTable
        on:click:header={(e) => {
          e.preventDefault();
          console.log(e);
        }}
        sortable
        zebra
        title="Load balancers"
        description="Your organization's active load balancers."
        headers={[
          { key: "name", value: "Name" },
          { key: "protocol", value: "Protocol" },
          { key: "port", value: "Port" },
          { key: "rule", value: "Rule" },
          { key: "_action", value: "Action" },
        ]}
        rows={[
          {
            id: "a",
            name: "Load Balancer 3",
            protocol: "HTTP",
            port: 3000,
            rule: "Round robin",
          },
          {
            id: "b",
            name: "Load Balancer 1",
            protocol: "HTTP",
            port: 443,
            rule: "Round robin",
          },
          {
            id: "c",
            name: "Load Balancer 2",
            protocol: "HTTP",
            port: 80,
            rule: "DNS delegation",
          },
          {
            id: "d",
            name: "Load Balancer 6",
            protocol: "HTTP",
            port: 3000,
            rule: "Round robin",
          },
          {
            id: "e",
            name: "Load Balancer 4",
            protocol: "HTTP",
            port: 443,
            rule: "Round robin",
          },
          {
            id: "f",
            name: "Load Balancer 5",
            protocol: "HTTP",
            port: 80,
            rule: "DNS delegation",
          },
        ]}
      >
        <span slot="cell" let:cell>
          {#if cell.key === "_action"}
            <OverflowMenu flipped>
              <OverflowMenuItem text="Manage credentials" />
              <OverflowMenuItem
                href="https://cloud.ibm.com/docs/api-gateway/"
                text="API documentation"
              />
              <OverflowMenuItem danger text="Delete service" />
            </OverflowMenu>
          {:else}{cell.value}{/if}
        </span>

        <Toolbar>
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
            <Button>Create balancer</Button>
          </ToolbarContent>
        </Toolbar>
      </DataTable>
      <Pagination totalItems={102} pageSizes={[16, 36, 99]} pageSize={36} />
    </Column>
  </Row>
</Grid>
