import base from '@/utils/base'
import { ElLoading, ElNotification } from 'element-plus'
import { defineStore } from 'pinia'
import { langStore } from '@/stores/lang'

const baseUrl = base.backendUrl

export const pingStore = defineStore('ping', () => {
    const { notificationItems } = langStore()

    let lang = localStorage.getItem('lang')

    if (lang === undefined) {
        lang = 'en'
    }

    const items = notificationItems(lang as 'en' | 'zh')
    type Answer = {
        code: string,
        message: string
    }
    const ping = (callback: (ok: boolean, ans?: string) => void) => {
        fetch(`${baseUrl}/ping`)
            .then(data => data.json())
            .then((data: Answer) => callback(true, data.message))
            .catch(_ => callback(false))
    }

    const connect = (callback?: (ok: boolean) => void) => {
        const { connectTimeout } = base
        const times = connectTimeout / 500
        let loading = ElLoading.service({
            text: "Loading"
        })

        // confirm connected
        ping(ok => {
            if (ok) {
                const conncted = sessionStorage.getItem('connected')
                if (conncted == undefined || conncted == "false") {
                    ElNotification(items.ping.success)
                }
                sessionStorage.setItem('connected', 'true')
                loading.close()
                if (callback != undefined) {
                    callback(true)
                }
                return
            } else {
                let count = 0
                let timer = setInterval(() => {
                    if (count < times) {
                        ping(ok => {
                            if (ok) {
                                const conncted = sessionStorage.getItem('connected')
                                if (conncted == undefined || conncted == "false") {
                                    ElNotification(items.ping.success)
                                }
                                sessionStorage.setItem('connected', 'true')
                                loading.close()
                                if (callback != undefined) {
                                    callback(true)
                                }
                                clearInterval(timer)
                            } else {
                                sessionStorage.setItem('connected', 'false')
                            }
                            count += 1
                        })
                    } else {
                        loading.close()
                        clearInterval(timer)
                        ElNotification(items.ping.error)
                        sessionStorage.setItem('connected', 'false')
                        if (callback != undefined) {
                            callback(false)
                        }
                    }
                }, 500)
            }
        })
    }

    return { ping, connect }
})
