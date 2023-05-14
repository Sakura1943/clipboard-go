<script setup lang="ts">
import { useDark } from '@vueuse/core'
import { langStore } from '../../stores/lang'
import { userStore } from '../../stores/user'
import { addNewUserFormStore } from '../../stores/auth-form'
import { onMounted, ref, type Ref, computed } from 'vue';
import { ElMessageBox, ElNotification } from 'element-plus';

// 获取验证表单的数据
const { ruleForm, ruleFormRef, rules } = addNewUserFormStore()
// 获取界面文字
const { notificationItems, tableItems, messageBoxItems } = langStore()
// 获取用户操作的方法
const { userList, userAuth, addUser } = userStore()

// 搜索的用户名
const searchUserName = ref('')
const typed = ref(false)
// 当前页
const currentPage = ref(1)
// 单个页面的显示数量
const pageSize = ref(10)
// 对话框可见
const dialogVisible = ref(false)
// 用户信息文字
const tableUsersInfo: Ref<{
    id: number,
    name: string,
    permission: string
}[] | undefined> = ref([])

// 获取当前语言，DOM渲染时判断是否存在，不存在则定义为en
let lang = localStorage.getItem('lang') === undefined ? 'en' : localStorage.getItem('lang') as string
// 当前用户
let user = localStorage.getItem('user')
// 当前用户权限
let role = localStorage.getItem('role')
// 获取表单的文字
const _tableItems = tableItems(lang)
// 通知文字
const _notificationItems = notificationItems(lang)
// 获取消息对话框的文字
const _messageBoxItems = messageBoxItems(lang)
// 当前token
const token = localStorage.getItem('token') as string

// 确认用户已登录
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

// 挂载后
onMounted(() => {
    if (user === undefined) {
        user = 'unknown'
    }
    if (role === undefined) {
        role = 'unknown'
    }
    userAuthConfirm()
    getItems()
})

// 获取单页数据
const singlePageInfo = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value
    const end = currentPage.value * pageSize.value
    if (typed.value) {
        if (searchUserName.value != '') {
            typed.value = true
            tableUsersInfo.value = tableUsersInfo.value?.filter((data) => data.name.includes(searchUserName.value))
            return tableUsersInfo.value?.slice(start, end)
        }
        typed.value = false
        getItems()
    }
    if (searchUserName.value != '') {
        typed.value = true
        tableUsersInfo.value = tableUsersInfo.value?.filter((data) => data.name.includes(searchUserName.value))
        return tableUsersInfo.value?.slice(start, end)
    }
    return tableUsersInfo.value?.slice(start, end)
})

// 获取当前页面用户的所有
const getItems = () => {
    userList(token, (data) => {
        if (data.code == 200) {
            tableUsersInfo.value = data.extra
        }
    })
}

// 改变当前页面数据
const changeCurrentPage = (pageNumber: number) => {
    currentPage.value = pageNumber
    getItems()
}

// 改变当前页面数据量
const changePageSize = (size: number) => {
    pageSize.value = size
    getItems()
}

// 打开对话框
const openDialog = () => {
    if (user !== "admin") {
        ElMessageBox.alert(_messageBoxItems.onlyAllowAdminToRegisterNewUser.alert?.warning?.message, _messageBoxItems.onlyAllowAdminToRegisterNewUser.alert?.warning?.title, {
            confirmButtonText: _messageBoxItems.onlyAllowAdminToRegisterNewUser.alert?.warning?.confirmButtonText,
            type: _messageBoxItems.onlyAllowAdminToRegisterNewUser.alert?.warning?.type
        })
        return
    }
    dialogVisible.value = true
}

const submitFormData = () => {
    addUser(token, {
        name: ruleForm.name,
        password: ruleForm.password,
        permission: ruleForm.permission
    }, (data) => {
        if (data === undefined) {
            ElNotification(_notificationItems.createUserFailed.error)
            return
        } else if (data.code !== 200) {
            if (data?.type === "user_exists") {
                ElNotification(_notificationItems.newUserIsExists.error)
                return
            }
            if (data.type === "form_data_parse_error") {
                ElNotification(_notificationItems.MisssingNameFiled.error)
                return
            }
            ElNotification(_notificationItems.createUserFailed.error)
            return
        }
        ElNotification(_notificationItems.createdUser.success)
        dialogVisible.value = false
        getItems()
    })
}

</script>
<template>
    <el-container class="container">
        <el-header class="header">
            <el-button type="primary" @click="openDialog">{{ _tableItems.user.add.button.add }}</el-button>
            <el-input v-model="searchUserName" :placeholder="_tableItems.user.add.search" class="search" />
        </el-header>
        <el-main class="main">
            <el-table v-if="!useDark().value" :data="singlePageInfo"
                :header-cell-style="{ background: '#DCDCDC', color: '#000000' }" stripe class="table">
                <el-table-column prop="id" :label="_tableItems.user.add.id" />
                <el-table-column prop="name" :label="_tableItems.user.add.name" />
                <el-table-column prop="permission" :label="_tableItems.user.add.permission" />
            </el-table>
            <el-table v-else :data="singlePageInfo" stripe class="table">
                <el-table-column prop="id" :label="_tableItems.user.add.id" />
                <el-table-column prop="name" :label="_tableItems.user.add.name" />
                <el-table-column prop="permission" :label="_tableItems.user.add.permission" />
            </el-table>
            <el-pagination :page-sizes="[5, 10]" :default-page-size="pageSize" @size-change="changePageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="tableUsersInfo?.length" :current-page="currentPage"
                @current-change="changeCurrentPage" />
        </el-main>
    </el-container>
    <!-- dialog viewer -->
    <el-dialog draggable v-model="dialogVisible" :title="_tableItems.user.add.add?.title" width="30rem" center>
        <el-form ref="ruleFormRef" :model="ruleForm" :rules="rules" label-width="auto" label-position="top">
            <el-form-item :label="_tableItems.user.add.add?.name" prop="name">
                <el-input v-model="ruleForm.name" :placeholder="_tableItems.user.add.add?.placeHolder.name" />
            </el-form-item>
            <el-form-item :label="_tableItems.user.add.add?.password">
                <el-input type="password" v-model="ruleForm.password"
                    :placeholder="_tableItems.user.add.add?.placeHolder.password" />
            </el-form-item>
            <el-form-item :label="_tableItems.user.add.add?.permission" prop="permission">
                <el-radio-group v-model="ruleForm.permission">
                    <el-radio label="admin" />
                    <el-radio label="custom" />
                </el-radio-group>
            </el-form-item>
        </el-form>
        <template #footer>
            <el-button @click="dialogVisible = false">{{ _tableItems.user.add.add?.button.cancel }}</el-button>
            <el-button @click="submitFormData" type="primary">{{ _tableItems.user.add.add?.button.confirm
            }}</el-button>
        </template>
    </el-dialog>
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
