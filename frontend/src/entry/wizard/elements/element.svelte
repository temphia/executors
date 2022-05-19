<script lang="ts">
  import { AllElements, BasicElement } from "./element";
  import type { FieldStore } from "../service/wizard_types";

  export let field: object;
  export let fieldstore: FieldStore;
  export let data_sources: object = {};
  export let prevDatas: object = {};

  let type = field["type"];
  let name = field["name"];
  let data_source = data_sources[field["source"]];
  let data = prevDatas[name];
</script>

{#if type.startsWith("basic.")}
  <BasicElement {data} {data_source} {field} field_store={fieldstore} />
{:else if AllElements[type]}
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
