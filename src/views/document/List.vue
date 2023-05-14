<script setup lang="ts">
import { ref, onMounted, type Ref } from 'vue'
import { documentStore } from '@/stores/document'
import { useDark } from '@vueuse/core'
import { userStore } from '@/stores/user'
import { langStore } from '@/stores/lang'
import { ElNotification } from 'element-plus'
import { computed } from 'vue'

// 搜索的文件路径
const searchFilePath = ref('')
const typed = ref(false)

// 当前页面，默认第一页
const currentPage = ref(1)
const pageSize = ref(10)

// 获取当前界面语言
const lang = localStorage.getItem('lang') === undefined ? 'en' : localStorage.getItem('lang') as string
// 获取token
const token = localStorage.getItem('token') === undefined ? 'unknown' : localStorage.getItem('token') as string

// 获取界面文字的对象
const { notificationItems, tableItems } = langStore()
// 获取文档操作的对象
const { getDocuments } = documentStore()
// 获得用户操作的对象
const { userAuth } = userStore()

// 获取界面的文字
const _notificationItems = notificationItems(lang)
const _tableItems = tableItems(lang)

// 获取的表单数据
const formData: Ref<{
    id: number,
    path: string,
    text: string,
    user_name: string
}[] | undefined> = ref()

// 获取表单数据的方法
const getFormData = () => {
    getDocuments(token, (data) => {
        if (data.code === 200) {
            formData.value = data.extra
            return
        }
        if (data.type === "empty") {
            formData.value = undefined
            return
        }
    })
}

// 单页数据
const singlePageData = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value
    const end = currentPage.value * pageSize.value
    if (typed.value) {
        if (searchFilePath.value != '') {
            typed.value = true
            formData.value = formData.value?.filter((data) => data.path.includes(searchFilePath.value))
            return formData.value?.slice(start, end)
        }
        typed.value = false
        getFormData()
    }
    if (searchFilePath.value != '') {
        typed.value = true
        formData.value = formData.value?.filter((data) => data.path.includes(searchFilePath.value))
        return formData.value?.slice(start, end)
    }
    return formData.value?.slice(start, end)
})

// 用户验证登录路由
const userAuthConfirm = () => {
    let token = localStorage.getItem('token')
    if (token === undefined) {
        ElNotification(_notificationItems.auth.error)
        return
    }
    userAuth(token as string, (data) => {
        if (data?.code != 200 || data.code === undefined) {
            ElNotification(_notificationItems.auth.error)
            return
        }
    })
}

// 改变页面数量
const changePageSize = (size: number) => {
    pageSize.value = size
    getFormData()
}

// 改变当前页面
const changeCurrentPage = (pageNumber: number) => {
    currentPage.value = pageNumber
    getFormData()
}

// 挂在时加载
onMounted(() => {
    userAuthConfirm()
    getFormData()
})
</script>

<template>
    <el-container class="container">
        <el-header class="header">
            <el-input class="search" v-model="searchFilePath" :placeholder="_tableItems.document.list.search" />
        </el-header>
        <el-main class="main">
            <el-table class="table" :data="singlePageData" stripe v-if="!useDark().value"
                :header-cell-style="{ background: '#DCDCDC', color: '#000000' }">
                <el-table-column prop="id" :label="_tableItems.document.list.id" />
                <el-table-column prop="path" :label="_tableItems.document.list.path" />
                <el-table-column prop="user_name" :label="_tableItems.document.list.upload_user" />
            </el-table>
            <el-table class="table" :data="singlePageData" stripe v-else>
                <el-table-column prop="id" :label="_tableItems.document.list.id" />
                <el-table-column prop="path" :label="_tableItems.document.list.path" />
                <el-table-column prop="user_name" :label="_tableItems.document.list.upload_user" />
            </el-table>
            <el-pagination background layout="total, sizes, prev, pager, next, jumper" v-model:current-page="currentPage"
                :total="formData?.length" :page-sizes="[5, 10]" :default-page-size="pageSize" @size-change="changePageSize"
                @current-change="changeCurrentPage">
            </el-pagination>
        </el-main>
    </el-container>
</template>

<style scoped>
.search {
    position: absolute;
    width: 15rem;
    right: 1.5rem;
}

.main {
    margin-top: 1rem;
    padding: 0;
}

.container {
    margin: 0;
    padding: 0;
}

.header {
    padding: 0;
    height: .5rem;
    margin-bottom: 1.25rem;
}

.table {
    padding: 0;
    margin-bottom: 1rem;
}
</style>
