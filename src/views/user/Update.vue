<script setup lang="ts">
import { userStore } from '@/stores/user'
import { langStore } from '@/stores/lang'
import { onMounted, ref, type Ref, computed } from 'vue'
import { ElMessageBox, ElNotification } from 'element-plus'
import { useDark } from '@vueuse/core'
import { debounce } from 'lodash-es'
import { userUpdateFormStore } from '@/stores/auth-form'
import base from '@/utils/base'

// 获取基础配置
const { debounceTime, trailing, leading } = base
// 获取用户操作的方法
const { userList, userAuth, updateUser } = userStore()
// 获取界面文字
const { notificationItems, tableItems, messageBoxItems } = langStore()

// 获取用户更新表单的数据的响应式数据和form表单校验规则
const {
    ruleFormRef,
    rules,
    ruleForm
} = userUpdateFormStore()

// 响应式变量
// 搜索的用户名
const searchUserName = ref('')
const typed = ref(false)
// 单个页面的显示数量
const pageSize = ref(10)
// 当前页
const currentPage = ref(1)
// 旧的用户名
const oldName = ref('')
// 对话框可见
const editdialogVisible = ref(false)

// 当前语言
let lang = localStorage.getItem('lang') === undefined ? 'en' : localStorage.getItem('lang') as string
// 当前token
const token = localStorage.getItem('token') as string

// 通知文字
const _notificationItems = notificationItems(lang)
// 消息框文字
const _messageBoxItems = messageBoxItems(lang)
// 表单文字
const _tableItems = tableItems(lang)

// 用户信息文字
const tableUsersInfo: Ref<{
    id: number,
    name: string,
    permission: string
}[] | undefined> = ref([])

// 挂载后
onMounted(() => {
    userAuthConfirm()
    getItems()
})

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

// 获取单页面数据
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

// 改变当前页面数
const changeCurrentPage = (pageNumber: number) => {
    currentPage.value = pageNumber
    getItems()
}

// 获取当前页面用户的所有
const getItems = () => {
    userList(token, (data) => {
        if (data.code == 200) {
            tableUsersInfo.value = data.extra
        }
    })
}

// 改变当前页面数据量
const changePageSize = (size: number) => {
    pageSize.value = size
    getItems()
}

// 用户更新触发
const _updateUser = debounce(function (_: number, row: {
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
    let role = localStorage.getItem('role')

    oldName.value = row.name

    let user = localStorage.getItem('user')
    if (user !== undefined) {
        user === 'unknown'
    }

    ruleForm.permission = row.permission
    ruleForm.name = row.name
    if (role === undefined) {
        role = 'unknown'
    }
    if (user !== 'admin') {
        if (user !== ruleForm.name) {
            ElMessageBox.alert(_messageBoxItems.notAllowChangeOtherUser.alert?.warning?.message, _messageBoxItems.notAllowChangeOtherUser.alert?.warning?.title, {
                confirmButtonText: _messageBoxItems.notAllowChangeOtherUser.alert?.warning?.confirmButtonText,
                type: _messageBoxItems.notAllowChangeOtherUser.alert?.warning?.type
            })
            return
        }
    }
    if (ruleForm.name == '') {
        ElMessageBox.alert(_messageBoxItems.userNameIsEmpty.alert?.warning?.message, _messageBoxItems.userNameIsEmpty.alert?.warning?.title, {
            confirmButtonText: _messageBoxItems.userNameIsEmpty.alert?.warning?.confirmButtonText,
            type: _messageBoxItems.userNameIsEmpty.alert?.warning?.type
        })
        return
    }
    editdialogVisible.value = true
}, debounceTime, {
    trailing,
    leading
})

// 更新用户信息触发
const sendUpdate = () => {
    // 如果当前用户为管理员且更新的用户也是管理员，并且修改其权限，则提示不允许操作
    if (localStorage.getItem('user') === 'admin' && ruleForm.name === 'admin' && ruleForm.permission == 'custom') {
        ElMessageBox.alert(_messageBoxItems.notAllowedChangePermission.alert?.error?.message, _messageBoxItems.notAllowedChangePermission.alert?.error?.title, {
            confirmButtonText: _messageBoxItems.notAllowedChangePermission.alert?.error?.confirmButtonText,
            type: _messageBoxItems.notAllowedChangePermission.alert?.error?.type
        })
        return
    }
    updateUser(token, {
        name: ruleForm.name,
        password: ruleForm.password,
        permission: ruleForm.permission,
        old_name: oldName.value
    }, (data) => {
        let role = localStorage.getItem('role')
        if (role === undefined) {
            role = "unknown"
        }
        let user = localStorage.getItem('user')
        if (user === undefined) {
            user = "unknown"
        }
        if (data.code === undefined) {
            ElNotification(_notificationItems.updateUser.error)
        } else if (data.code !== 200) {
            if (data.type === 'unknown_permission') {
                ElNotification(_notificationItems.unknownPermission.error)
            } else if (data.error === "The current user is not allowed to change their permissions") {
                ElNotification(_notificationItems.notAllowedChangePermission.error)
            } else if (data.type === 'not_allowed_to_change_admin_info') {
                ElNotification(_notificationItems.notAllowedToChangeAdminInfo.error)
            } else {
                ElNotification(_notificationItems.updateUser.error)
            }
        } else {
            ElNotification(_notificationItems.updateUser.success)
            editdialogVisible.value = false
        }
        if (user !== ruleForm.name && user === 'admin') {
            localStorage.setItem('role', ruleForm.permission)
        }
        getItems()
    })
}
</script>

<template>
    <el-container class="container">
        <el-header class="header">
            <el-input v-model="searchUserName" :placeholder="_tableItems.user.add.search" class="search" />
        </el-header>
        <el-main class="main">
            <el-table v-if="!useDark().value" :data="singlePageInfo"
                :header-cell-style="{ background: '#DCDCDC', color: '#000000' }" stripe class="table">
                <el-table-column prop="id" :label="_tableItems.user.update.id" />
                <el-table-column prop="name" :label="_tableItems.user.update.name" />
                <el-table-column prop="permission" :label="_tableItems.user.update.permission" />
                <el-table-column :label="_tableItems.user.update.operate">
                    <template #default="scope">
                        <el-button type="primary" @click="_updateUser(scope.$index, scope.row)">
                            {{ _tableItems.user.update.button.edit }}
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-table v-else :data="singlePageInfo" stripe class="table">
                <el-table-column prop="id" :label="_tableItems.user.update.id" />
                <el-table-column prop="name" :label="_tableItems.user.update.name" />
                <el-table-column prop="permission" :label="_tableItems.user.update.permission" />
                <el-table-column :label="_tableItems.user.update.operate">
                    <template #default="scope">
                        <el-button type="primary" @click="_updateUser(scope.$index, scope.row)">
                            {{ _tableItems.user.update.button.edit }}
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination :page-sizes="[5, 10]" :default-page-size="pageSize" @size-change="changePageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="tableUsersInfo?.length" :current-page="currentPage"
                @current-change="changeCurrentPage" />
        </el-main>
    </el-container>
    <!-- dialog viewer -->
    <el-dialog draggable v-model="editdialogVisible" :title="_tableItems.user.update.edit?.title" width="30rem" center>
        <el-form ref="ruleFormRef" :model="ruleForm" :rules="rules" label-width="auto" label-position="top">
            <el-form-item :label="_tableItems.user.update.edit?.name" prop="name">
                <el-input v-model="ruleForm.name" :placeholder="_tableItems.user.update.edit?.placeHolder.name" />
            </el-form-item>
            <el-form-item :label="_tableItems.user.update.edit?.password">
                <el-input type="password" v-model="ruleForm.password"
                    :placeholder="_tableItems.user.update.edit?.placeHolder.password" />
            </el-form-item>
            <el-form-item :label="_tableItems.user.update.edit?.permission" prop="permission">
                <el-radio-group v-model="ruleForm.permission">
                    <el-radio label="admin" />
                    <el-radio label="custom" />
                </el-radio-group>
            </el-form-item>
        </el-form>
        <template #footer>
            <el-button @click="editdialogVisible = false">{{ _tableItems.user.update.edit?.button.cancel }}</el-button>
            <el-button @click="sendUpdate" type="primary">{{ _tableItems.user.update.edit?.button.confirm }}</el-button>
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