<script lang="ts">
  import { AdvElements, BasicElement, ViewElements } from "./element";
  import type { FieldsStore } from "../service";

  export let field: object;
  export let fieldstore: FieldsStore;
  export let data_sources: object = {};
  export let prevDatas: object = {};

  let type = field["type"];
  let name = field["name"];
  let data_source = data_sources[field["source"]];
  let data = prevDatas[name];
</script>

{#if type.startsWith("basic.")}
  <BasicElement {data} {data_source} {field} field_store={fieldstore} />
{:else if type.startsWith("adv.")}
  {#if AdvElements[type]}
    <svelte:component
      this={AdvElements[type]}
      {data}
      {data_source}
      {field}
      field_store={fieldstore}
    />
  {:else}
    <div>Adv elem not implemented</div>
  {/if}
{:else if type.startsWith("view.")}
  {#if ViewElements[type]}
    <svelte:component
      this={ViewElements[type]}
      {data}
      {data_source}
      {field}
      field_store={fieldstore}
    />
  {:else}
    <div>View elem not implemented</div>
  {/if}
{:else}
  <div>Not implemented</div>
{/if}
