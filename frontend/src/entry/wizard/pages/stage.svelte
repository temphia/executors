<script lang="ts">
  import Layout from "../core/layout.svelte";
  import Processing from "../core/processing.svelte";
  import Element from "../elements/element.svelte";
  import type { Manager } from "../service/wizard_types";

  export let manager: Manager;
  const store = manager.get_state();
  const data_sources = $store.data_sources;
</script>

<Layout
  title={manager.wizard_title}
  showButtons={$store.flowState !== "STAGE_PROCESSING"}
  next={manager.stage_next}
>
  {#if $store.flowState === "STAGE_PROCESSING"}
    <Processing />
  {:else}
    {#each $store.fields || [] as field}
      <div class="relative my-4">
        <Element
          {field}
          {data_sources}
          fieldstore={manager.get_field_store(field["name"])}
        />
      </div>
    {/each}
  {/if}
</Layout>
