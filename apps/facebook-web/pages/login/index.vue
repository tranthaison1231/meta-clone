<script setup lang="ts">
import { configure } from 'vee-validate'
import type { z } from 'zod'
import { loginSchema } from '~/lib/schema'
import { request } from '~/lib/request'
import { setToken } from '~/lib/storage'

useHead({
  title: 'Facebook - log in or sign up',
})

definePageMeta({
  middleware: 'auth',
})
const router = useRouter()

const { mutateAsync } = useMutation({
  mutationFn: (data: z.infer<typeof loginSchema>) => request.post('/sign-in', data),
})

configure({
  validateOnBlur: true,
})

const { handleSubmit, errors, useFieldModel } = useForm({
  validationSchema: toTypedSchema(loginSchema),
})

const [email, password] = useFieldModel(['email', 'password'])

const onSubmit = handleSubmit(async (values) => {
  const res = await mutateAsync(values)
  setToken(res.data.token)
  router.push('/')
})
</script>

<template>
  <div class="bg-gray-200 h-screen">
    <div class="max-w-5xl w-full mx-auto flex justify-between items-center h-screen">
      <div class="w-1/2">
        <h1 class="text-primary text-6xl font-bold">
          facebook
        </h1>
        <p class="text-3xl mt-4">
          Facebook helps you connect and share with the people in your life.
        </p>
      </div>
      <div>
        <form class="p-4 bg-white shadow-md rounded-xl w-104 h-fit flex flex-col items-center" @submit="onSubmit">
          <Input v-model="email" placeholder="Email address or phone number" />
          <span class="my-1 w-full text-red-500">{{ errors.email }}</span>
          <Input v-model="password" placeholder="Password" class="my-3" type="password" />
          <span class="my-1 w-full text-red-500">{{ errors.password }}</span>
          <Button class="w-full text-xl" type="submit">
            Log in
          </Button>
          <a class="text-primary hover:underline mt-3 block text-center"> Forgotten password? </a>
          <hr class="my-4">
          <Button variant="success" class="w-fit text-base my-2" type="button">
            Create new account
          </Button>
        </form>
        <p class="mt-3 text-center text-sm">
          <span class="font-bold">Create a Page </span> for a celebrity, brand or business.
        </p>
      </div>
    </div>
  </div>
</template>
