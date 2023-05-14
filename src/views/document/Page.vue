<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { langStore } from '@/stores/lang'
import { documentStore } from '@/stores/document'
import { ElMessageBox } from 'element-plus'
// import hljs from 'highlight.js'

// 获取当前语言
const lang = localStorage.getItem('lang') === undefined ? 'en' : localStorage.getItem('lang') as string

// 获取提示文字
const { messageBoxItems } = langStore()
const _messageBoxItems = messageBoxItems(lang)

// 代码部分
const code = ref('')

// 获取文档操作的对象
const { getOneDocument } = documentStore()

// 获取路由对象
const router = useRouter()
// 获取路径中的文件路径和文件语言
const fileName = router.currentRoute.value.params.file_name as string
const fileLang = router.currentRoute.value.params.lang as string

onMounted(() => {
    // 强制设置该页面为黑色
    const color = localStorage.getItem('vueuse-color-scheme')
    if (color !== undefined && color === "light") {
        localStorage.setItem('vueuse-color-scheme', 'dark')
    }
    // 获取文档
    getOneDocument(fileName, fileLang, (data) => {
        if (data.code === 200) {
            code.value = data.extra?.text as string
            return
        }
        ElMessageBox.alert(_messageBoxItems.getContentFailed.alert?.error?.message, _messageBoxItems.getContentFailed.alert?.error?.title, {
            confirmButtonText: _messageBoxItems.getContentFailed.alert?.error?.confirmButtonText,
            type: _messageBoxItems.getContentFailed.alert?.error?.type
        })
    })
})

</script>

<template>
    <highlightjs :code="code" />
</template>

<style scoped></style>
