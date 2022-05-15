<script lang="ts">
  import { onMount } from "svelte";

  let ref = null;

  function init() {
    console.log(ref);

    ref.contentDocument.querySelector("head").innerHTML += `<style>
                *, ::after, ::before {box-sizing: border-box;}
                :root {tab-size: 4;}
                html {line-height: 1.15;text-size-adjust: 100%;}
                body {margin: 0px; padding: 1rem 0.5rem;}
                body {font-family: system-ui, -apple-system, "Segoe UI", Roboto, Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji";}
                </style>`;
    ref.contentDocument.body.innerHTML += `
                <h1>Hello World!</h1>
                <p>Welcome to the Tailwind WYSIWYG.</p>
                `;

    // ref.contentDocument.designMode = "on";
  }

  function format(cmd, param) {
    ref.contentDocument.execCommand(cmd, !1, param || null);
  }

  onMount(async () => {
    init();
  });
</script>

<div
  class="min-w-screen min-h-screen bg-gray-200 flex items-center justify-center px-5 py-5"
>
  <div
    class="w-full max-w-6xl mx-auto rounded-xl bg-white shadow-lg p-5 text-black"
  >
    <div class="border border-gray-200 overflow-hidden rounded-md">
      <div class="w-full flex border-b border-gray-200 text-xl text-gray-600">
        <button
          class="outline-none focus:outline-none border-r border-gray-200 w-10 h-10 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("bold")}
        >
          <i class="mdi mdi-format-bold" />
        </button>
        <button
          class="outline-none focus:outline-none border-r border-gray-200 w-10 h-10 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("italic")}
        >
          <i class="mdi mdi-format-italic" />
        </button>
        <button
          class="outline-none focus:outline-none border-r border-gray-200 w-10 h-10 mr-1 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("underline")}
        >
          <i class="mdi mdi-format-underline" />
        </button>
        <button
          class="outline-none focus:outline-none border-l border-r border-gray-200 w-10 h-10 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("formatBlock", "P")}
        >
          <i class="mdi mdi-format-paragraph" />
        </button>
        <button
          class="outline-none focus:outline-none border-r border-gray-200 w-10 h-10 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("formatBlock", "H1")}
        >
          <i class="mdi mdi-format-header-1" />
        </button>
        <button
          class="outline-none focus:outline-none border-r border-gray-200 w-10 h-10 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("formatBlock", "H2")}
        >
          <i class="mdi mdi-format-header-2" />
        </button>
        <button
          class="outline-none focus:outline-none border-r border-gray-200 w-10 h-10 mr-1 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("formatBlock", "H3")}
        >
          <i class="mdi mdi-format-header-3" />
        </button>
        <button
          class="outline-none focus:outline-none border-l border-r border-gray-200 w-10 h-10 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("insertUnorderedList")}
        >
          <i class="mdi mdi-format-list-bulleted" />
        </button>
        <button
          class="outline-none focus:outline-none border-r border-gray-200 w-10 h-10 mr-1 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("insertOrderedList")}
        >
          <i class="mdi mdi-format-list-numbered" />
        </button>
        <button
          class="outline-none focus:outline-none border-l border-r border-gray-200 w-10 h-10 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("justifyLeft")}
        >
          <i class="mdi mdi-format-align-left" />
        </button>
        <button
          class="outline-none focus:outline-none border-r border-gray-200 w-10 h-10 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("justifyCenter")}
        >
          <i class="mdi mdi-format-align-center" />
        </button>
        <button
          class="outline-none focus:outline-none border-r border-gray-200 w-10 h-10 hover:text-indigo-500 active:bg-gray-50"
          on:click={() => format("justifyRight")}
        >
          <i class="mdi mdi-format-align-right" />
        </button>
      </div>
      <div class="w-full">
        <iframe bind:this={ref} class="w-full h-96 overflow-y-auto" />
      </div>
    </div>
  </div>
</div>
