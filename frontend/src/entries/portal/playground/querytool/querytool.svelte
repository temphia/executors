<script lang="ts">
  import { HSplitPane, VSplitPane } from "svelte-split-pane";
</script>

<div class="wrapper w-full h-full">
  <VSplitPane
    updateCallback={() => {
      console.log("VSplitPane Updated!");
    }}
  >
    <top slot="top"> Top Pane </top>
    <down slot="down"> Down Pane </down>
  </VSplitPane>
</div>
