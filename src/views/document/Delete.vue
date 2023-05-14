<script setup lang="ts">
import { computed, ref, type Ref, onMounted } from 'vue'
import { langStore } from '@/stores/lang'
import { documentStore, type PageData } from '@/stores/document'
import { userStore } from '@/stores/user'
import { useDark } from '@vueuse/core'
import { ElMessageBox, ElNotification } from 'element-plus';

// 响应式数据
// 查询的文章路径
const searchFilePath = ref('')
const currentPage = ref(1)
const pageSize = ref(10)
const typed = ref(false)
const visiable = ref(false)

// 获取界面文字的对象
const { notificationItems, tableItems, messageBoxItems } = langStore()
// 获取当前界面语言
const lang = localStorage.getItem('lang') === undefined ? 'en' : localStorage.getItem('lang') as string
// 获取界面的文字
const _tableItems = tableItems(lang)
const _notificationItems = notificationItems(lang)
const _messageBoxItems = messageBoxItems(lang)

// 获取文档操作的对象
const { getDocuments, deleteDocument } = documentStore()
// 获得用户操作的对象
const { userAuth } = userStore()

// 获取token
const token = localStorage.getItem('token') === undefined ? 'unknown' : localStorage.getItem('token') as string
// 获取用户名
const user = localStorage.getItem("user") === undefined ? 'unknown' : localStorage.getItem('user') as string

// 获取的表单数据
const formData: Ref<PageData[] | undefined> = ref()

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

// 改变当前页面
const changeCurrentPage = (pageNumber: number) => {
    currentPage.value = pageNumber
    getFormData()
}

// 改变页面数据数量
const changePageSize = (size: number) => {
    pageSize.value = size
    getFormData()
}

// 删除用户操作
const deleteUserHandler = (_index: number, data: PageData) => {
    if (data.user_name !== user && user !== "admin") {
        ElMessageBox.alert(_messageBoxItems.noPermissionToDeletePage.alert?.warning?.message, _messageBoxItems.noPermissionToDeletePage.alert?.warning?.title, {
            confirmButtonText: _messageBoxItems.noPermissionToDeletePage.alert?.warning?.confirmButtonText,
            type: _messageBoxItems.noPermissionToDeletePage.alert?.warning?.type
        })
        return
    }
    deleteDocument(token, data.path, (data) => {
        if (data === undefined || data.code !== 200) {
            ElNotification(_notificationItems.deleteDocumentFailed.error)
            return
        }
        ElNotification(_notificationItems.deleteDocumentSuccess.success)
        getFormData()
    })
}

// 挂载时
onMounted(() => {
    userAuthConfirm()
    getFormData()
})
</script>

<template>
    <el-container class="container">
        <el-header class="header">
            <el-input class="search" v-model="searchFilePath" :placeholder="_tableItems.document.delete.search" />
        </el-header>
        <el-main class="main">
            <el-table class="table" :data="singlePageData" stripe v-if="!useDark().value"
                :header-cell-style="{ background: '#DCDCDC', color: '#000000' }">
                <el-table-column prop="id" :label="_tableItems.document.delete.id" />
                <el-table-column prop="path" :label="_tableItems.document.delete.path" />
                <el-table-column prop="user_name" :label="_tableItems.document.delete.upload_user" />
                <el-table-column :label="_tableItems.document.delete.operate">
                    <template #default="scope">
                        <el-button type="danger" @click="deleteUserHandler(scope.$index, scope.row)">
                            {{ _tableItems.document.delete.delete }}
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-table class="table" v-else>
                <el-table-column prop="id" :label="_tableItems.document.delete.id" />
                <el-table-column prop="path" :label="_tableItems.document.delete.path" />
                <el-table-column prop="user_name" :label="_tableItems.document.delete.upload_user" />
                <el-table-column :label="_tableItems.document.delete.operate">
                    <template #default="scope">
                        <el-button type="danger" @click="deleteUserHandler(scope.$index, scope.row)">
                            {{ _tableItems.document.delete.delete }}
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination background layout="total, sizes, prev, pager, next, jumper" v-model:current-page="currentPage"
                :total="formData?.length" :page-sizes="[5, 10]" :default-page-size="pageSize" @size-change="changePageSize"
                @current-change="changeCurrentPage">
            </el-pagination>
        </el-main>
    </el-container>
</template>

<style scoped>
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

.search {
    position: absolute;
    width: 15rem;
    right: 1.5rem;
}
</style>