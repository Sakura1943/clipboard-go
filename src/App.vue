<script setup lang="ts">
import { onMounted } from 'vue'
import { RouterView } from 'vue-router'
import { useDark } from '@vueuse/core'
onMounted(() => {
    useDark()
})
</script>

<template>
    <RouterView/>
</template>

<style scoped></style>
