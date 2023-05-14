<script setup lang="ts">
import { langStore } from '@/stores/lang'
import router from '@/router/index'
const { errorPageItems } = langStore()

// 获取当前界面语言
const lang = localStorage.getItem('lang') === undefined ? 'en' : localStorage.getItem('lang') as string

// 获取404页面文字
const pageItem = errorPageItems(lang)
</script>

<template>
    <div class="container">
        <el-result icon="warning" :title="pageItem.title" :sub-title="pageItem.sub_title">
            <template #extra>
                <el-button type="primary" @click="router.push('/')">
                    {{ pageItem.button }}
                </el-button>
            </template>
        </el-result>
    </div>
</template>

<style scoped>
.container {
    height: 100vh;
    width: 100vw;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
</style>
