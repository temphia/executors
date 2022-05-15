<script lang="ts">
  import Layout from "../core/layout.svelte";
  import Processing from "../core/processing.svelte";
  import Element from "../elements/element.svelte";
  import type { WizardManager } from "../service/wizard";

  export let manager: WizardManager;
  const store = manager._state;
  const fieldstore = manager._fieldsStore;
  const data_sources = $store.data_sources;
</script>

<Layout
  title={manager._wizard_title}
  showButtons={$store.flowState !== "STAGE_PROCESSING"}
  next={manager.stage_next}
>
  {#if $store.flowState === "STAGE_PROCESSING"}
    <Processing />
  {:else}
    {#each $store.fields || [] as field}
      <div class="relative my-4">
        <Element {field} {fieldstore} {data_sources} />
      </div>
    {/each}
  {/if}
</Layout>
