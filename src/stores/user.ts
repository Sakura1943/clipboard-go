import base from "@/utils/base"
import { defineStore } from "pinia"
import { type Response } from '@/utils/response'

export const userStore = defineStore('user', () => {
    const login = (userName: string, password: string, callback: (data: Response<{
        name: string,
        role: string,
        token: string
    } | undefined> | null) => void) => {
        const url = `${base.backendUrl}/api/login`
        let formData = new FormData()
        formData.append('name', userName)
        formData.append('password', password)
        fetch(url, {
            method: 'POST',
            body: formData
        }).then(data => data.json())
            .then(data => callback(data))
            .catch(_ => callback(null))
    }
    const userAuth = (token: string, callback: (data: Response | null) => void) => {
        const url = `${base.backendUrl}/api/user/auth_confirm`
        fetch(url, {
            method: 'GET',
            headers: {
                token
            },
        }).then(data => data.json())
            .then(data => callback(data))
            .catch(data => callback(data))
    }

    const userList = (token: string, callback: (data: Response<{
        id: number,
        name: string,
        permission: string
    }[] | undefined>) => void) => {
        const url = `${base.backendUrl}/api/user/list`
        fetch(url, {
            method: 'GET',
            headers: {
                token
            }
        }).then(data => data.json())
            .then(data => callback(data))
            .catch(data => callback(data))
    }

    const deleteUser = (token: string, userName: string, callback: (data: Response) => void) => {
        const url = `${base.backendUrl}/api/user/delete/${userName}`
        fetch(url, {
            method: 'DELETE',
            headers: {
                token
            },
        }).then(data => data.json())
            .then(data => callback(data))
            .catch(data => callback(data))
    }

    const updateUser = (token: string, userInfo: {
        name: string,
        password: string,
        permission: string,
        old_name: string
    }, callback: (data: Response<any | undefined>) => void) => {
        const url = `${base.backendUrl}/api/user/update`
        let formData = new FormData()
        formData.append('name', userInfo.name)
        if (userInfo.password !== '') {
            formData.append('password', userInfo.password)
        }

        if (userInfo.permission !== '') {
            formData.append('permission', userInfo.permission)
        }
        formData.append('old_name', userInfo.old_name)
        fetch(url, {
            method: 'PUT',
            headers: {
                token
            },
            body: formData
        }).then(data => data.json())
            .then(data => callback(data))
            .catch(data => callback(data))
    }

    const addUser = (token: string, userInfo: {
        name: string,
        password: string,
        permission: string
    }, callback: (data: Response<any | undefined>) => void) => {
        const url = `${base.backendUrl}/api/user/register`
        let formData = new FormData()
        formData.append('name', userInfo.name)
        formData.append('password', userInfo.password)
        formData.append('permission', userInfo.permission)
        fetch(url, {
            method: 'POST',
            headers: {
                token
            },
            body: formData
        })
            .then(data => data.json())
            .then(data => callback(data))
            .catch(data => callback(data))
    }
    return { login, userAuth, userList, deleteUser, updateUser, addUser }
})
