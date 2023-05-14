<script setup lang="ts">
import { onMounted, onBeforeMount } from 'vue'
import Menu from '@/components/Menu.vue'
import { userStore } from '@/stores/user'
import router from '@/router'
import { ElNotification } from 'element-plus'
import { langStore } from '@/stores/lang'
const { notificationItems } = langStore()
const { userAuth } = userStore()

// 挂载前
onBeforeMount(() => {
    const lang = localStorage.getItem('lang') === undefined ? 'en' : localStorage.getItem('lang') as string
    const items = notificationItems(lang)
    const token = localStorage.getItem('token')
    if (token === undefined) {
        router.push('/auth')
    } else {
        userAuth(token as string, (data) => {
            if (data?.code != 200 || data == null) {
                ElNotification(items.auth.error)
                router.push('/auth')
            }
        })
    }
})

// 挂载后
onMounted(() => {
    if (localStorage.getItem('lang') == undefined) {
        localStorage.setItem('lang', 'en')
    }
})

</script>

<template>
    <Menu />
</template>

<style scoped></style>
