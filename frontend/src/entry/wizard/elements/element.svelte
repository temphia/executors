<script lang="ts">
  import { AllElements, BasicElement } from "./element";
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
  {#if AllElements[type]}
    <svelte:component
      this={AllElements[type]}
      {data}
      {data_source}
      {field}
      field_store={fieldstore}
    />
  {:else}
    <div>Elem not implemented</div>
  {/if}
{:else}
  <div>Not implemented</div>
{/if}
