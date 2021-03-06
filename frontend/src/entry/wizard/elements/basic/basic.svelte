<script lang="ts">
  import type { FieldsStore } from "../../service";
  import * as Elem from "../../service/elem_types";
  import EmojiSelector from "svelte-emoji-selector";

  export let field: object;
  export let data_source: any;
  export let data: any;
  export let field_store: FieldsStore;

  let value = data === undefined ? data_source : data;

  let type = field["type"];
  let name = field["name"];
  let options = field["options"] || {};
  let attrs = field["attrs"] || {};

  const change = (ev) => {
    value = ev.target.value;
    field_store.set_value(name, value);
  };

  const changeBool = (ev) => {
    value = ev.target.checked;
    field_store.set_value(name, value);
  };

  const changeNum = (ev) => {
    value = ev.target.checked;
    field_store.set_value(name, Number(value));
  };

  function onEmoji(ev) {
    value = value ? value + ev.detail : ev.detail;
    field_store.set_value(name, value);
  }
</script>

<label for={name} class="text-base leading-7 uppercase">{name}</label>
{#if type === Elem.BASIC_SHORTTEXT}
  <div class="flex w-full">
    <input
      type="text"
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5 relative"
      placeholder=""
      {...attrs}
    />
    {#if !options["hide_emojipicker"]}
      <EmojiSelector on:emoji={onEmoji} />
    {/if}
  </div>
{:else if type === Elem.BASIC_LONGTEXT}
  <div class="flex w-full">
    <textarea
      on:change={change}
      value={value || ""}
      class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
      placeholder="write something..."
      {...attrs}
    />
    {#if !options["hide_emojipicker"]}
      <EmojiSelector on:emoji={onEmoji} />
    {/if}
  </div>
{:else if type === Elem.BASIC_RANGE}
  <input
    type="range"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_SELECT}
  <div class="flex w-full mt-5">
    <select class="w-full p-2" {value} on:change={change}>
      {#each field["options"] || [] as opt}
        <option value={opt}>{opt}</option>
      {/each}
    </select>
  </div>
{:else if type === Elem.BASIC_MULTI_SELECT}
  <div class="flex flex-col w-full h-full p-1 overflow-auto">
    <div
      class="flex flex-col p-1 space-y-1 border border-dashed rounded-lg bg-gray-100 text-gray-800"
      style="min-height: 2rem;"
      {...attrs}
    >
      {#each field["options"] as opt}
        <label>
          <input type="checkbox" class="form-checkbox h-5 w-5 text-gray-600" />
          {opt}
        </label>
      {/each}
    </div>
  </div>
{:else if type === Elem.BASIC_PHONE}
  <input
    type="tel"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_CHECKBOX}
  <input
    type="checkbox"
    id={name}
    on:change={changeBool}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_COLOR}
  <input
    type="color"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_DATE}
  <input
    type="date"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_DATETIME}
  <input
    type="datetime-local"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_EMAIL}
  <input
    type="email"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_NUMBER}
  <input
    type="number"
    id={name}
    on:change={changeNum}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{/if}
<span class="font-sans text-sm italic">{field["info"] || ""}</span>
