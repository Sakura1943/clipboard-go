<script setup lang="ts">
import { loginFormStore } from '@/stores/auth-form'
import { pingStore } from '@/stores/ping'
import { langStore } from '@/stores/lang'
import { userStore } from '@/stores/user'
import { onBeforeMount, ref } from 'vue'
import router from '@/router/index'
import base from '@/utils/base'
import { ElNotification } from 'element-plus'
import { debounce } from 'lodash-es'
const { ruleFormRef, rules, ruleForm } = loginFormStore()
const { loginFormItems, notificationItems } = langStore()
const { login } = userStore()
const { connect } = pingStore()
const { debounceTime, trailing, leading } = base

// 获取当前界面语言
const lang = localStorage.getItem('lang') === undefined ? 'en' : localStorage.getItem('lang') as string

// 获取界面文字
const _loginFormItems = loginFormItems(lang)
const _notificationItems = notificationItems(lang)

// 响应式数据
const disabled = ref(false)

// 挂载时
onBeforeMount(() => {
    connect((ok) => {
        if (!ok) {
            disabled.value = true
        }
    })
})

// 登录操作
const loginHandle = debounce(function () {
    login(ruleForm.name, ruleForm.pass, (data) => {
        if (data == null) {
            ElNotification(_notificationItems.login.error)
            return
        } else if (data.code != 200 || data.code === undefined) {
            ElNotification(_notificationItems.login.error)
            console.log(data.error)
            return
        }
        else {
            ElNotification(_notificationItems.login.success)
            localStorage.setItem('token', data.extra?.token as string)
        }
        localStorage.setItem('user', data?.extra?.name as string)
        localStorage.setItem('role', data?.extra?.role as string)
        router.push('/navigation')
    })
}, debounceTime, {
    trailing,
    leading
})
</script>

<template>
    <div class="container">
        <div class="box">
            <el-text class="header">{{ _loginFormItems.header }}</el-text>
            <el-form ref="ruleFormRef" class="form" size="large" show-message inline-message :rules="rules" status-icon>
                <el-form-item class="item" prop="name">
                    <el-input v-model="ruleForm.name" type="text" autocomplete="on"
                        :placeholder="_loginFormItems.name_placeholder" :disabled="disabled" @keyup.enter="loginHandle" />
                </el-form-item>
                <el-form-item class="item" prop="password">
                    <el-input v-model="ruleForm.pass" type="password" autocomplete="off"
                        :placeholder="_loginFormItems.password_placeholder" :disabled="disabled"
                        @keyup.enter="loginHandle"></el-input>
                </el-form-item>
                <el-form-item class="button">
                    <el-button type="primary" color="#626aef" @click="loginHandle" :disabled="disabled">{{
                        _loginFormItems.button }}</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<style scoped>
.container {
    height: 100vh;
    width: 100vw;
    display: flex;
    align-items: center;
    justify-content: center;
    background-image: url("@/assets/images/background.jpg");
    background-attachment: fixed;
    background-repeat: no-repeat;
    background-size: cover;
}

.box {
    overflow: hidden;
    height: 25rem;
    width: 28rem;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 20px;
    backdrop-filter: blur(20px);
    background: rgba(255, 255, 255, 0.3);
    box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.2);
    flex-direction: column;
}

.box .form {
    display: flex;
    width: 80%;
    justify-content: center;
    flex-direction: column;
}

.box .form .item {
    margin-bottom: 2rem;
}

.box .form .button {
    width: 100%;
    margin-top: 10px;
    margin-bottom: 10px;
}

.box .form .el-input {
    border-radius: 5px;
    height: 4rem;
    background-color: rgba(255, 255, 255, 0.3);
}

:deep() .el-input__wrapper {
    border: none !important;
    box-shadow: none !important;
}

:deep() .box .form .el-input .el-input__wrapper>* {
    color: rgba(0, 0, 0, 0.5);
}

:deep() .box .form .el-input.is-disabled .el-input__wrapper {
    background-color: rgba(0, 0, 0, 0.08);
}

:deep() .box .form .el-input.is-disabled .el-input__wrapper>* {
    color: rgba(0, 0, 0, 0.5);
}

.box .form .button .el-button {
    border-radius: 5px;
    width: 100%;
    height: 3.5rem;
}

.box .header {
    margin-top: 10px;
    margin-bottom: 30px;
    font-size: 30px;
    color: rgba(0, 0, 0, 0.3);
    user-select: none;
}
</style>