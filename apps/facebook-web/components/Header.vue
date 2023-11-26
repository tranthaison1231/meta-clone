<script setup lang="ts">
import { Gamepad, Home, MonitorPlay, Store, Users } from 'lucide-vue-next'

defineProps<{
  user: {
    avatar: string
  }
}>()
</script>

<template>
  <header class="h-14 shadow-md px-4">
    <div grid="~ cols-7" class="gap-20">
      <div class="col-span-2 p-2 flex items-center gap-2">
        <div> Logo </div>
        <Input class="h-10 w-60 rounded-3xl bg-gray-100" placeholder="Search Facebook" />
      </div>
      <div class="col-span-3 flex justify-between items-center">
        <Home :size="24" class="cursor-pointer" />
        <MonitorPlay :size="24" class="cursor-pointer" />
        <Store :size="24" class="cursor-pointer" />
        <Users :size="24" class="cursor-pointer" />
        <Gamepad :size="24" class="cursor-pointer" />
      </div>
      <div class="col-span-2 flex items-center justify-end gap-2">
        <div> Menu </div>
        <div> Message </div>
        <div> Notification </div>
        <ProfilePopover :avatar="user?.avatar" />
      </div>
    </div>
  </header>
</template>
