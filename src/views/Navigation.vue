<script setup lang="ts">
import { langStore } from '@/stores/lang'
const { navigationItems } = langStore()
const items = () => {
    const lang = localStorage.getItem('lang') === undefined ? 'en' : localStorage.getItem('lang') as string
    return navigationItems(lang)
}
</script>

<template>
    <el-container class="container">
        <i class="header">{{ items().title }}</i>
        <el-text class="content" v-for="item in items().main_content">{{ item }}</el-text>
    </el-container>
</template>

<style scoped>
.container {
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    height: 100%;
    width: 100%;
    justify-content: center;
    align-items: center;
}

.container .header {
    font-size: 1.5rem;
    margin-bottom: 10px;
}
</style>