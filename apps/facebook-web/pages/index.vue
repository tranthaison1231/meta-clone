<script setup lang="ts">
defineProps<{
  user: {
    avatar: string
    firstName: string
    lastName: string
  }
}>()

useHead({
  title: 'Facebook',
})

definePageMeta({
  middleware: 'auth',
  layout: 'private',
})
</script>

<template>
  <div grid="~ cols-7" class="gap-20 bg-gray-200 h-[calc(100vh-3.5rem)] p-4">
    <div class="col-span-2">
      <div class="flex gap-2 items-center">
        <img class="h-10 w-10 rounded-full" :src="user?.avatar">
        <p> {{ user?.firstName }} </p>
      </div>
    </div>
    <div class="col-span-3">
      <div> Stories </div>
      <div class="p-4 bg-white rounded-xl mt-4">
        <div class="flex items-center gap-2">
          <img class="h-10 w-10 rounded-full" :src="user?.avatar">
          <Input class="h-10 rounded-3xl bg-gray-100 w-full" placeholder="What's on your mind" />
        </div>
        <hr class="mt-4 mb-6">
        <div class="flex justify-around">
          <div> Live video </div>
          <div> Photo/video </div>
          <div> Feeling/Activity </div>
        </div>
      </div>
      <div class="mt-4">
        Posts
      </div>
    </div>
    <div class="col-span-2">
      <div>Your Pages and profiles</div>
      <hr>
      <div>
        <div>
          <h2> Contacts </h2>
        </div>
      </div>
    </div>
  </div>
</template>
