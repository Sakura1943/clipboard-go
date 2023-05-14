import type { FormInstance, FormRules } from 'element-plus'
import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'

export const loginFormStore = defineStore('auth-form', () => {
    const ruleFormRef = ref<FormInstance>()

    let lang = localStorage.getItem('lang')
    if (lang === undefined) {
        lang = 'en'
    }

    const checkName = (rule: any, value: any, callback: any) => {
        if (ruleForm.name === '') {
            if (lang === 'zh') {
                return callback(new Error('请输入用户名称'))
            } else {
                return callback(new Error('Please input the user name'))
            }
        } else if (!(/^[0-9a-z]*$/.test(ruleForm.name))) {
            if (lang === 'zh') {
                return callback(new Error('请输入小写字母或数字'))
            } else {
                return callback(new Error('Please enter lowercase letters or integer'))
            }
        } else {
            if (ruleForm.name !== '') {
                if (!ruleFormRef.value) return
                (ruleFormRef.value as unknown as FormInstance).validateField('name', () => null)
            }
        }
    }

    const checkPassword = (rule: any, value: any, callback: any) => {
        if (ruleForm.pass === '') {
            if (lang === 'zh') {
                return callback(new Error('请输入密码'))
            } else {
                return callback(new Error('Please input the password'))
            }
        } else {
            if (ruleForm.pass !== '') {
                if (!ruleFormRef.value) return
                (ruleFormRef.value as unknown as FormInstance).validateField('pass', () => null)
            }
        }
    }

    const ruleForm = reactive({
        name: '',
        pass: ''
    })
    const rules = reactive<FormRules>({
        name: [{ validator: checkName, trigger: 'blur' }],
        password: [{ validator: checkPassword, trigger: 'blur' }]
    })
    return { ruleForm, ruleFormRef, checkName, checkPassword, rules }
})

export const userUpdateFormStore = defineStore('user-update-form', () => {
    const ruleFormRef = ref<FormInstance>()

    let lang = localStorage.getItem('lang')
    if (lang === undefined) {
        lang = 'en'
    }
    const ruleForm = reactive({
        name: '',
        password: '',
        permission: 'admin'
    })
    const validateName = (rule: any, value: any, callback: any) => {
        if (ruleForm.name === '') {
            if (lang === 'zh') {
                return callback(new Error('请输入用户名称'))
            } else {
                return callback(new Error('Please input the user name'))
            }
        } else if (!(/^[0-9a-z]*$/.test(ruleForm.name))) {
            if (lang === 'zh') {
                return callback(new Error('请输入小写字母或数字'))
            } else {
                return callback(new Error('Please enter lowercase letters or integer'))
            }
        } else {
            if (ruleForm.name !== '') {
                if (!ruleFormRef.value) return
                (ruleFormRef.value as unknown as FormInstance).validateField('name', () => null)
            }
        }
    }

    const validatePermission = (rule: any, value: any, callback: any) => {
        if (ruleForm.permission === '') {
            if (lang === 'zh') {
                return callback(new Error('请选择权限'))
            } else {
                return callback(new Error('Please select permission'))
            }
        } else if (ruleForm.permission !== 'admin' && ruleForm.permission !== 'custom') {
            if (lang === 'zh') {
                return callback(new Error('仅允许admin和custom权限'))
            } else {
                return callback(new Error('Only admin and custom permissions are allowed'))
            }
        }
    }
    const rules = reactive<FormRules>({
        name: [{ validator: validateName, trigger: 'blur' }],
        permission: [{ validator: validatePermission, trigger: 'blur' }]
    })

    return { rules, ruleForm, ruleFormRef, validateName, validatePermission }
})

export const addNewUserFormStore = defineStore('add-new-user-form', () => {
    const ruleFormRef = ref<FormInstance>()

    let lang = localStorage.getItem('lang')
    if (lang === undefined) {
        lang = 'en'
    }
    const ruleForm = reactive({
        name: '',
        password: '',
        permission: 'admin'
    })

    const validateName = (rule: any, value: any, callback: any) => {
        if (ruleForm.name === '') {
            if (lang === 'zh') {
                return callback(new Error('请输入用户名称'))
            } else {
                return callback(new Error('Please input the user name'))
            }
        } else if (!(/^[0-9a-z]*$/.test(ruleForm.name))) {
            if (lang === 'zh') {
                return callback(new Error('请输入小写字母或数字'))
            } else {
                return callback(new Error('Please enter lowercase letters or integer'))
            }
        } else {
            if (ruleForm.name !== '') {
                if (!ruleFormRef.value) return
                (ruleFormRef.value as unknown as FormInstance).validateField('name', () => null)
            }
        }
    }

    const validatePermission = (rule: any, value: any, callback: any) => {
        if (ruleForm.permission === '') {
            if (lang === 'zh') {
                return callback(new Error('请选择权限'))
            } else {
                return callback(new Error('Please select permission'))
            }
        } else if (ruleForm.permission !== 'admin' && ruleForm.permission !== 'custom') {
            if (lang === 'zh') {
                return callback(new Error('仅允许admin和custom权限'))
            } else {
                return callback(new Error('Only admin and custom permissions are allowed'))
            }
        }
    }
    const rules = reactive<FormRules>({
        name: [{ validator: validateName, trigger: 'blur' }],
        permission: [{ validator: validatePermission, trigger: 'blur' }]
    })

    return { ruleForm, ruleFormRef, validateName, validatePermission, rules }
})