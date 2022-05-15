<script lang="ts">
  const timelines = [
    {
      type: "create_record",
      fields: ["deadline", "target"],
      assoc_user: "user12",
      assoc_version: 0,
      mod_sig: "xyxyxyxyyxyxyxyxyxyxyxyxyxy",
    },
    {
      type: "update_record",
      fields: ["deadline", "target"],
      assoc_user: "user12",
      assoc_version: 1,
      mod_sig: "xyxyxyxyyxyxyxyxyxyxyxyxyxy",
    },
    {
      type: "update_record",
      fields: ["deadline", "target"],
      assoc_user: "user12",
      assoc_version: 2,
      mod_sig: "xyxyxyxyyxyxyxyxyxyxyxyxyxy",
    },
    {
      type: "delete_record",
      fields: ["deadline", "target"],
      assoc_user: "user12",
      assoc_version: 3,
      mod_sig: "xyxyxyxyyxyxyxyxyxyxyxyxyxy",
    },
  ];
</script>

<div class="w-10/12 md:w-7/12 lg:6/12 mx-auto relative py-20">
  <h1 class="text-3xl text-center font-bold text-blue-500">
    Timeline with Tailwindcss
  </h1>
  <div class="border-l-2 mt-10">
    {#each timelines as tl}
      <div
        class="transform border transition cursor-pointer hover:-translate-y-2 ml-10 relative flex items-center px-6 py-4 rounded mb-10 flex-col md:flex-row space-y-4 md:space-y-0 hover:border-purple-600"
      >
        <!-- Dot Follwing the Left Vertical Line -->
        <div
          class="w-5 h-5 bg-blue-600 absolute -left-10 transform -translate-x-2/4 rounded-full z-10 mt-2 md:mt-0"
        />

        <!-- Line that connecting the box with the vertical line -->
        <div class="w-10 h-1 bg-blue-300 absolute -left-10 z-0" />

        <!-- Content that showing in the box -->
        <div class="flex items-center space-x-4">
          <div class="text-blue-200">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-7 w-7"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-6-3a2 2 0 11-4 0 2 2 0 014 0zm-2 4a5 5 0 00-4.546 2.916A5.986 5.986 0 0010 16a5.986 5.986 0 004.546-2.084A5 5 0 0010 11z"
                clip-rule="evenodd"
              />
            </svg>
            <p class="text-sm font-medium text-gray-900 truncate">
              @{tl.assoc_user}
            </p>
          </div>
          <span class="bg-red-300 text-sm p-1 rounded">
            {tl.type}
          </span>

          <a href="#">
            #{tl.mod_sig}
          </a>
        </div>
      </div>
    {/each}
  </div>
</div>
