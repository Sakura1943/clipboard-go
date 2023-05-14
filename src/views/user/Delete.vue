<script setup lang="ts">
import { userStore } from '@/stores/user'
import { langStore } from '@/stores/lang'
import { onMounted, ref, type Ref, computed } from 'vue'
import { ElNotification } from 'element-plus'
import { useDark } from '@vueuse/core'
import { debounce } from 'lodash-es'
import base from '@/utils/base'
// 当前页面语言
let lang = localStorage.getItem('lang') === undefined ? 'en' : localStorage.getItem('lang') as string

// 获取页面文字
const { userList, userAuth, deleteUser } = userStore()
const { notificationItems, tableItems } = langStore()
const { debounceTime, trailing, leading } = base
const _notificationItems = notificationItems(lang)
const _tableItems = tableItems(lang)

// 当前用户token
const token = localStorage.getItem('token') as string
// 响应式数据
// 搜索的用户名
const searchUserName = ref('')
const typed = ref(false)
// 当前页面数据数量
const pageSize = ref(10)
// 当前页面
const currentPage = ref(1)
// 所有用户信息
const tableUsersInfos: Ref<{
    id: number,
    name: string,
    permission: string
}[] | undefined> = ref([])

// 确认用户是登录
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

// 单页数据
const singlePageInfo = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value
    const end = currentPage.value * pageSize.value
    if (typed.value) {
        if (searchUserName.value != '') {
            typed.value = true
            tableUsersInfos.value = tableUsersInfos.value?.filter((data) => data.name.includes(searchUserName.value))
            return tableUsersInfos.value?.slice(start, end)
        }
        typed.value = false
        getItems()
    }
    if (searchUserName.value != '') {
        typed.value = true
        tableUsersInfos.value = tableUsersInfos.value?.filter((data) => data.name.includes(searchUserName.value))
        return tableUsersInfos.value?.slice(start, end)
    }
    return tableUsersInfos.value?.slice(start, end)
})

// 改变当前页面
const changeCurrentPage = (pageNumber: number) => {
    currentPage.value = pageNumber
    getItems()
}

// 获取用户数据
const getItems = () => {
    userList(token, (data) => {
        if (data.code == 200) {
            tableUsersInfos.value = data.extra
        }
    })
}

// 改变当前的数量
const changePageSize = (size: number) => {
    pageSize.value = size
    getItems()
}

// 删除用户
const _deleteUser = debounce(function (_: number, row: {
    id: string,
    name: string,
    permission: string
}) {
    userAuthConfirm()
    let token = localStorage.getItem('token')
    if (token === undefined) {
        ElNotification(_notificationItems.auth.error)
        return
    }
    let localUser = localStorage.getItem('user')
    if (localUser === undefined) {
        localUser == 'unknown'
    }
    if (row.name == localUser) {
        ElNotification(_notificationItems.deleteSelf.error)
        return
    }
    if (row.name == 'admin') {
        ElNotification(_notificationItems.deleteAdmin.error)
        return
    }
    deleteUser(token as string, row.name, (data) => {
        if (data.code !== 200 || data.code === undefined) {
            ElNotification(_notificationItems.deleteUser.error)
            return
        }
        getItems()
        ElNotification(_notificationItems.deleteUser.success)
    })
}, debounceTime, {
    trailing,
    leading
})

// 挂载后
onMounted(() => {
    userAuthConfirm()
    getItems()
})

</script>

<template>
    <el-container class="container">
        <el-header class="header">
            <el-input v-model="searchUserName" :placeholder="_tableItems.user.add.search" class="search" />
        </el-header>
        <el-main class="main">
            <el-table v-if="!useDark().value" :data="singlePageInfo"
                :header-cell-style="{ background: '#DCDCDC', color: '#000000' }" stripe class="table">
                <el-table-column prop="id" :label="_tableItems.user.delete.id" />
                <el-table-column prop="name" :label="_tableItems.user.delete.name" />
                <el-table-column prop="permission" :label="_tableItems.user.delete.permission" />
                <el-table-column :label="_tableItems.user.delete.operate">
                    <template #default="scope">
                        <el-button type="danger" @click="_deleteUser(scope.$index, scope.row)">
                            {{ _tableItems.user.delete.button.delete }}
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-table v-else :data="singlePageInfo" stripe class="table">
                <el-table-column prop="id" :label="_tableItems.user.delete.id" />
                <el-table-column prop="name" :label="_tableItems.user.delete.name" />
                <el-table-column prop="permission" :label="_tableItems.user.delete.permission" />
                <el-table-column :label="_tableItems.user.delete.operate">
                    <template #default="scope">
                        <el-button type="danger" @click="_deleteUser(scope.$index, scope.row)">
                            {{ _tableItems.user.delete.button.delete }}
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination :page-sizes="[5, 10]" :default-page-size="pageSize" @size-change="changePageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="tableUsersInfos?.length" :current-page="currentPage"
                @current-change="changeCurrentPage" />
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
    margin: 0;
    margin-bottom: 1rem;
}

.search {
    position: absolute;
    right: 1.5rem;
    width: 15rem;
}
</style>