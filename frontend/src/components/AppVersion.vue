<script setup lang="ts">
const packageVersion = import.meta.env.PACKAGE_VERSION;
</script>
<template>
  <div class="flex flex-initial">
    <div class="p-4 flex flex-auto justify-center items-center">
      <UModal>
        <UButton
          icon="simple-icons:git"
          color="neutral"
          variant="soft"
          size="sm"
          :label="packageVersion"
        />

        <template #content>
          <Placeholder class="h-48 m-4" />
        </template>
      </UModal>
    </div>
  </div>
</template>
