<script setup lang="ts">
  import { PopoverRoot, PopoverPortal, PopoverTrigger, PopoverContent } from 'radix-vue'
import { removeToken } from '~/lib/storage';
  import { cn } from '~/lib/utils';

  const { avatar, ...props } = defineProps<{
    avatar: string
    class?: string
  }>()

  const router = useRouter()

  const logout = () => {
    removeToken()
    router.push('/login')
  }

</script>

<template>
  <PopoverRoot>
    <PopoverTrigger class="cursor-pointer">
      <img class="h-10 w-10 rounded-full cursor-pointer" :src="avatar" />
    </PopoverTrigger>
    <PopoverPortal>
      <PopoverContent :class="
        cn(
          'z-50 w-72 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2',
          props?.class,
        )
      ">
        <div @click="logout">Logout </div>
      </PopoverContent>
    </PopoverPortal>
  </PopoverRoot>
  
</template>
